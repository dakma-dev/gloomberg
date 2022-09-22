package server

import (
	"fmt"
	"os"
	"strings"

	"github.com/benleb/gloomberg/internal/server/node"

	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/gbnode"
	"github.com/benleb/gloomberg/internal/server/subscribe"
	"github.com/benleb/gloomberg/internal/server/ws"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/viper"
)

type ChainWatcher struct {
	// mu sync.RWMutex

	Nodes node.Nodes

	OwnCollections *collections.Collections

	queueLogs   *chan types.Log
	queueEvents *chan *collections.Event
	queueOutWS  *chan *collections.Event

	WebsocketsServer *ws.WebsocketsServer
}

func New() *ChainWatcher {
	// initialize our node connections
	nodes := GetNodes()

	// get the successfully connected nodes
	var nodesOnline node.Nodes = nodes["online"]

	// create a queue/channel for the received logs
	queueLogs := make(chan types.Log, 1024)
	// create a queue/channel for events to be sent out via ws
	queueOutWS := make(chan *collections.Event, 1024)

	return &ChainWatcher{
		OwnCollections: collections.New(),
		Nodes:          nodesOnline,

		queueLogs:  &queueLogs,
		queueOutWS: &queueOutWS,
	}
}

func (cs *ChainWatcher) GetNodesAsList() []*node.Node {
	var nodes []*node.Node

	for _, cNode := range cs.Nodes {
		nodes = append(nodes, cNode)
	}

	return nodes
}

func (cs *ChainWatcher) GetNodesAsGBNodeCollection() *gbnode.NodeCollection {
	var nodes gbnode.NodeCollection

	for _, cNode := range cs.Nodes {
		cNode := gbnode.ChainNode{
			NodeID:             cNode.NodeID,
			Name:               cNode.Name,
			Client:             cNode.Client,
			Marker:             cNode.Marker,
			WebsocketsEndpoint: cNode.WebsocketsEndpoint,
			ReceivedMessages:   cNode.ReceivedMessages,
			KillTimer:          cNode.KillTimer,
			Error:              cNode.Error,
		}
		nodes = append(nodes, &cNode)
	}

	return &nodes
}

func (cs *ChainWatcher) Subscribe(queueEvents *chan *collections.Event) {
	// channel to push the parsed events to
	cs.queueEvents = queueEvents

	for _, cNode := range cs.Nodes {
		gbl.Log.Infof("%s: subscribing to chain events | cs.QueueEvents: %d", cNode.Name, len(*cs.queueEvents))

		// subscribe to all events where first topic is the "Transfer" topic
		if _, err := cNode.SubscribeToTransfers(*cs.queueLogs); err != nil {
			gbl.Log.Warnf("Transfers subscribe to %s failed: %s", cNode.WebsocketsEndpoint, err)
		}
		// subscribe to all events where first topic is the "SingleTransfer" topic
		if _, err := cNode.SubscribeToSingleTransfers(*cs.queueLogs); err != nil {
			gbl.Log.Warnf("SingleTransfers subscribe to %s failed: %s", cNode.WebsocketsEndpoint, err)
		}

		// create a defined number of workers/handlers per cNode to receive and process incoming events/logs
		for workerID := 1; workerID <= viper.GetInt("server.workers.subscription_logs"); workerID++ {
			go subscribe.WorkerLogsQueue(workerID, cNode, cs.Nodes, cs.OwnCollections, cs.queueLogs, cs.queueEvents, cs.queueOutWS)
		}
	}
}

func GetNodes() map[string][]*node.Node {
	nodesOnline := make([]*node.Node, 0)
	nodesOffline := make([]*node.Node, 0)

	nodesSpinner := style.GetSpinner("setting up n connections...")
	_ = nodesSpinner.Start()

	for idx, ep := range viper.Get("endpoints").([]interface{}) {
		config := make(map[string]string, 0)
		localNode := false

		switch nodeConfig := ep.(type) {
		case string:
			config["endpoint"] = nodeConfig

		case map[string]interface{}:
			for k, v := range nodeConfig {
				switch setting := v.(type) {
				case string:
					config[k] = setting
				case bool:
					if k == "local" {
						localNode = setting
					}
				}
			}
		}

		if config["marker"] == "" {
			config["marker"] = fmt.Sprint(" ", idx)
		}

		if config["color"] != "" {
			config["marker"] = lipgloss.NewStyle().Foreground(lipgloss.Color(config["color"])).Render(config["marker"])
		}

		if config["name"] == "" {
			config["name"] = fmt.Sprint("cNode", idx)
		}

		if config["endpoint"] == "" {
			fmt.Printf("endpoint missin config: %+v\n\n", config)
			continue
		}

		cNode, err := node.New(idx, config["name"], config["marker"], config["endpoint"], localNode)

		if err == nil {
			nodesOnline = append(nodesOnline, cNode)

			gbl.Log.Infof("✅ successfully connected to %s", style.BoldStyle.Render(config["endpoint"]))
		} else {
			nodesOffline = append(nodesOffline, cNode)

			gbl.Log.Warnf("❌ failed to connect to %s | %s:", config["name"], config["endpoint"])
			gbl.Log.Warnf("%s %s", style.PinkBoldStyle.PaddingLeft(3).Render("→"), err)
		}
	}

	numOnlineNodes := len(nodesOnline)

	if numOnlineNodes == 0 {
		gbl.Log.Error("No n providers found")

		nodesSpinner.StopMessage(fmt.Sprint(style.BoldStyle.Render("no n providers found")))
		_ = nodesSpinner.Stop()

		os.Exit(1)
	}

	nodeNames := make([]string, 0)
	for _, n := range nodesOnline {
		nodeNames = append(nodeNames, style.BoldStyle.Render(n.Name))
	}

	nodesSpinner.StopMessage(
		fmt.Sprint(style.BoldStyle.Render(fmt.Sprint(numOnlineNodes)), " nodes connected: ", strings.Join(nodeNames, ", ")) + "\n",
	)

	_ = nodesSpinner.Stop()

	return map[string][]*node.Node{
		"online":  nodesOnline,
		"offline": nodesOffline,
	}
}
