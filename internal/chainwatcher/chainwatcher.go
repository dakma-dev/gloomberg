package chainwatcher

import (
	"github.com/benleb/gloomberg/internal/chainwatcher/subscribe"
	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/nodes"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/benleb/gloomberg/internal/ws"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/viper"
)

type ChainWatcher struct {
	// mu sync.RWMutex

	Nodes nodes.Nodes

	OwnCollections *collections.Collections

	queueLogs *chan types.Log
	// queueEvents *chan *collections.Event
	queueOutWS *chan *collections.Event

	WebsocketsServer *ws.WebsocketsServer
}

func New(nodes []*nodes.Node) *ChainWatcher {
	// create a queue/channel for the received logs
	queueLogs := make(chan types.Log, 1024)
	// create a queue/channel for events to be sent out via ws
	queueOutWS := make(chan *collections.Event, 1024)

	return &ChainWatcher{
		OwnCollections: collections.New(),
		Nodes:          nodes,

		queueLogs:  &queueLogs,
		queueOutWS: &queueOutWS,
	}
}

func (cs *ChainWatcher) GetNodesAsList() []*nodes.Node {
	var nodes []*nodes.Node

	for _, cNode := range cs.Nodes {
		nodes = append(nodes, cNode)
	}

	return nodes
}

func (cs *ChainWatcher) SubscribeToSales(queueEvents *chan *collections.Event) {
	// channel to push the parsed events to
	// cs.queueEvents = queueEvents
	for _, cNode := range cs.Nodes {
		gbl.Log.Infof("%s: subscribing to chain events | QueueEvents: %d", cNode.Name, len(*queueEvents))

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
			go subscribe.WorkerLogsQueue(workerID, cNode, cs.Nodes, cs.OwnCollections, cs.queueLogs, queueEvents)
		}
	}
}

//
//func GetNodes() map[string][]*nodes.Node {
//	nodesOnline := make([]*nodes.Node, 0)
//	nodesOffline := make([]*nodes.Node, 0)
//
//	nodesSpinner := style.GetSpinner("setting up n connections...")
//	_ = nodesSpinner.Start()
//
//	for idx, ep := range viper.Get("endpoints").([]interface{}) {
//		config := make(map[string]string, 0)
//		localNode := false
//
//		switch nodeConfig := ep.(type) {
//		case string:
//			config["endpoint"] = nodeConfig
//
//		case map[string]interface{}:
//			for k, v := range nodeConfig {
//				switch setting := v.(type) {
//				case string:
//					config[k] = setting
//				case bool:
//					if k == "local" {
//						localNode = setting
//					}
//				}
//			}
//		}
//
//		if config["marker"] == "" {
//			config["marker"] = fmt.Sprint(" ", idx)
//		}
//
//		if config["color"] != "" {
//			config["marker"] = lipgloss.NewStyle().Foreground(lipgloss.Color(config["color"])).Render(config["marker"])
//		}
//
//		if config["name"] == "" {
//			config["name"] = fmt.Sprint("cNode", idx)
//		}
//
//		if config["endpoint"] == "" {
//			fmt.Printf("endpoint missin config: %+v\n\n", config)
//			continue
//		}
//
//		cNode, err := nodes.New(idx, config["name"], config["marker"], config["endpoint"], localNode)
//
//		if err == nil {
//			nodesOnline = append(nodesOnline, cNode)
//
//			gbl.Log.Infof("✅ successfully connected to %s", style.BoldStyle.Render(config["endpoint"]))
//		} else {
//			nodesOffline = append(nodesOffline, cNode)
//
//			gbl.Log.Warnf("❌ failed to connect to %s | %s:", config["name"], config["endpoint"])
//			gbl.Log.Warnf("%s %s", style.PinkBoldStyle.PaddingLeft(3).Render("→"), err)
//		}
//	}
//
//	numOnlineNodes := len(nodesOnline)
//
//	if numOnlineNodes == 0 {
//		gbl.Log.Error("No n providers found")
//
//		nodesSpinner.StopMessage(fmt.Sprint(style.BoldStyle.Render("no n providers found")))
//		_ = nodesSpinner.Stop()
//
//		os.Exit(1)
//	}
//
//	nodeNames := make([]string, 0)
//	for _, n := range nodesOnline {
//		nodeNames = append(nodeNames, style.BoldStyle.Render(n.Name))
//	}
//
//	nodesSpinner.StopMessage(
//		fmt.Sprint(style.BoldStyle.Render(fmt.Sprint(numOnlineNodes)), " nodes connected: ", strings.Join(nodeNames, ", ")) + "\n",
//	)
//
//	_ = nodesSpinner.Stop()
//
//	return map[string][]*nodes.Node{
//		"online":  nodesOnline,
//		"offline": nodesOffline,
//	}
//}
