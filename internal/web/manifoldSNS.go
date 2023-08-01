package web

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

// var testNewClaim = `{"messageId":"802257d1-31b2-4b3b-9715-a3b23571caa2","receiptHandle":"AQEBAzQYylUjk8ULP+nkS4F7URV+HsEHta3aPXrhPJvZgDwrod78XWioDjbg0fmWgFgQ/978GOu6QhfIWygZXMhKeENKIdISdM4S8sjyPbSn8TnXqwasiYKETDrWASMPSYDPnDJiMl62zujy8PtwV5f1pTC31LXxcAJtVt/Qhu9uMVlYEGTYSWcSYQBc2SJSLJKYscBoyLr3LhaKE5hhx65JBHZyPFYa+SkpgzECu8R0TgiyxPzaSozJ9hHAP4ohO1iTZzQm7BpbK4k37LIzliovyox+hyMzcGzQUCQuAd4jiKU=","body":"{\"creatorAddress\":\"0x0ad99abaa361879c2d5663302a3962a7625be984\",\"link\":\"https://app.manifold.xyz/c/derazetives\",\"networkId\":1}","attributes":{"ApproximateReceiveCount":"1","AWSTraceHeader":"Root=1-64c39e1c-2ea879288afd69ff4bd844ae;Parent=7dc1b39d62ba87de;Sampled=0;Lineage=0d3c07a2:0","SentTimestamp":"1690541598526","SequenceNumber":"18879522722932208128","MessageGroupId":"NewClaim","SenderId":"AIDAYRRVD2ENU4DSO2WBX","MessageDeduplicationId":"Claim/derazetives","ApproximateFirstReceiveTimestamp":"1690541598526"},"messageAttributes":{},"md5OfBody":"d4e6e2196677cfe732e5156c7cd6cb2e","eventSource":"aws:sqs","eventSourceARN":"arn:aws:sqs:eu-central-1:929868421883:NewManifoldInstances.fifo","awsRegion":"eu-central-1"}`

// func TestNewClaim() {
// 	// parse json
// 	var snsMsg ManifoldSNSMessage
// 	if err := json.Unmarshal(body, &snsMsg); err != nil {
// 		log.Warnf("manifoldSNS parsing json failed: %+v | %+v", body, r.Body)

// 		w.WriteHeader(http.StatusBadRequest)

// 		_, err := w.Write([]byte(`{"result": "error", "message": "could not parse request body"}`))
// 		if err != nil {
// 			log.Warnf("manifoldSNS writing response failed: %s", err)
// 		}

// 		log.Warnf("manifoldSNS parsing json failed: %s", r.Body)

// 		return
// 	}
// }

type SNSMessage struct {
	MessageID     string `json:"messageId"`
	ReceiptHandle string `json:"receiptHandle"`
	Body          string `json:"body"`
	// Body          json.RawMessage `json:"body"`
	// Body       ManifoldSNSMessage `json:"body"`
	Attributes struct {
		ApproximateReceiveCount          string `json:"ApproximateReceiveCount"`
		AWSTraceHeader                   string `json:"AWSTraceHeader"`
		SentTimestamp                    string `json:"SentTimestamp"`
		SequenceNumber                   string `json:"SequenceNumber"`
		MessageGroupID                   string `json:"MessageGroupId"`
		SenderID                         string `json:"SenderId"`
		MessageDeduplicationID           string `json:"MessageDeduplicationId"`
		ApproximateFirstReceiveTimestamp string `json:"ApproximateFirstReceiveTimestamp"`
	} `json:"attributes"`
	MessageAttributes struct{} `json:"messageAttributes"`
	Md5OfBody         string   `json:"md5OfBody"`
	EventSource       string   `json:"eventSource"`
	EventSourceARN    string   `json:"eventSourceARN"`
	AwsRegion         string   `json:"awsRegion"`
}

type ManifoldSNSMessage struct {
	Subject                string             `json:"Subject"`
	Message                ManifoldSNSPayload `json:"Message"`
	MessageGroupID         string             `json:"MessageGroupId"`
	MessageDeduplicationID string             `json:"MessageDeduplicationId"`
}

type ManifoldSNSPayload struct {
	CreatorAddress common.Address `json:"creatorAddress"`
	Link           string         `json:"link"`
	NetworkID      int64          `json:"networkId"`
}

func StartmanifoldSNS(gb *gloomberg.Gloomberg) {
	// if viper.GetBool("gloomberg.manifoldSNS.enabled") {

	listenHost := viper.GetString("manifoldSNS.listen")
	port := viper.GetUint16("manifoldSNS.port")
	serverAddress := fmt.Sprintf("%s:%d", listenHost, port)

	postRoute := "manifoldSNS"

	gb.PrModf("web", "starting manifold SNS receiver on https://%s/%s", serverAddress, postRoute)

	// manifold sns handler (received via aws sqs/sns)
	http.HandleFunc("/"+postRoute, HandlerManifoldSNSTopic)

	tlsConfig, err := gloomberg.GetServerTLSConfig()
	if err != nil {
		log.Warn("TLS certificate not found, using insecure connection")
	}

	// create http server
	snsReceiverServer := &http.Server{
		Addr:              serverAddress,
		ReadHeaderTimeout: 2 * time.Second,
		Handler:           nil,
		TLSConfig:         tlsConfig,
	}

	// start http server
	// log.Debugf("starting manifold SNS receiver on %s | %+v", listenOn, snsReceiverServer)
	gb.PrDModf("web", "starting manifold SNS receiver on https://%s/%s", serverAddress, postRoute)

	// go func() {
	if err := snsReceiverServer.ListenAndServeTLS("", ""); err != nil {
		log.Fatalf("manifold sns receiver died: %s", err)
	}
	// }()
}

func HandlerManifoldSNSTopic(w http.ResponseWriter, r *http.Request) {
	// check HTTP method
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte(`{"result": "error", "message": "method not allowed"}`))
		if err != nil {
			log.Warnf("manifoldSNS writing response failed: %s", err)
		}

		log.Warnf("manifoldSNS called with method %s", r.Method)

		return
	}

	// read json body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		_, err := w.Write([]byte(`{"result": "error", "message": "could not read request body"}`))
		if err != nil {
			log.Warnf("manifoldSNS writing response failed: %s", err)
		}

		log.Warnf("manifoldSNS reading body failed: %s", r.Body)

		return
	}

	// parse json
	var snsMsg SNSMessage
	if err := json.Unmarshal(body, &snsMsg); err != nil {
		log.Warnf("manifoldSNS unmarshal json failed: %+v | %+v", err, string(body))

		w.WriteHeader(http.StatusBadRequest)

		_, err := w.Write([]byte(`{"result": "error", "message": "could not parse request body"}`))
		if err != nil {
			log.Warnf("manifoldSNS writing response failed: %+v | %+v", err, string(body))
		}

		log.Warnf("manifoldSNS parsing json failed: %+v | %+v", err, string(body))

		return
	}

	// log.Print("body:")
	// pretty.Println(body)

	// log.Print("⬇  ⬇  ⬇  ⬇  ⬇  ⬇")
	// log.Printf("snsMsg: %#v", snsMsg)
	// log.Print("⬆  ⬆  ⬆  ⬆  ⬆  ⬆")
	// pretty.Println(snsMsg)

	var manifoldSNSPayload ManifoldSNSPayload
	if err := json.Unmarshal([]byte(snsMsg.Body), &manifoldSNSPayload); err != nil {
		log.Warnf("manifoldSNSPayload unmarshal body json failed: %+v | %+v", err, snsMsg.Body)

		// w.WriteHeader(http.StatusBadRequest)

		// _, err := w.Write([]byte(`{"result": "error", "message": "could not parse request body"}`))
		// if err != nil {
		// 	log.Warnf("manifoldSNSPayload writing body response failed: %+v | %+v", err, string(body))
		// }

		// log.Warnf("manifoldSNSPayload parsing body json failed: %+v | %+v", err, unquotedBody)

		// return
	}

	// unquotedBody, err := strconv.Unquote(snsMsg.Body)
	// if err != nil {
	// 	log.Printf("manifoldSNS unquoting body failed: %+v | %+v", err, snsMsg.Body)

	// 	w.WriteHeader(http.StatusBadRequest)

	// 	_, err := w.Write([]byte(`{"result": "error", "message": "could not parse request body"}`))
	// 	if err != nil {
	// 		log.Warnf("manifoldSNS writing response failed: %+v | %+v", err, snsMsg.Body)
	// 	}

	// 	return
	// }

	// log.Print("")
	// log.Print("⬇  ⬇  ⬇  ⬇  ⬇  ⬇")
	// log.Printf(" snsMsg.Body: %#v", snsMsg.Body)
	// log.Printf("unquotedBody: %#v", unquotedBody)
	// log.Print("⬆  ⬆  ⬆  ⬆  ⬆  ⬆")
	// log.Print("")

	// var manifoldSNSMsg ManifoldSNSMessage
	// if err := json.Unmarshal([]byte(unquotedBody), &manifoldSNSMsg); err != nil {
	// 	log.Warnf("manifoldSNS unmarshal body json failed: %+v | %+v", err, unquotedBody)

	// 	// w.WriteHeader(http.StatusBadRequest)

	// 	// _, err := w.Write([]byte(`{"result": "error", "message": "could not parse request body"}`))
	// 	// if err != nil {
	// 	// 	log.Warnf("manifoldSNS writing body response failed: %+v | %+v", err, string(body))
	// 	// }

	// 	// log.Warnf("manifoldSNS parsing body json failed: %+v | %+v", err, unquotedBody)

	// 	// // return
	// }

	// var manifoldSNSPayload ManifoldSNSPayload
	// if err := json.Unmarshal([]byte(unquotedBody), &manifoldSNSPayload); err != nil {
	// 	log.Warnf("manifoldSNSPayload unmarshal body json failed: %+v | %+v", err, unquotedBody)

	// 	// w.WriteHeader(http.StatusBadRequest)

	// 	// _, err := w.Write([]byte(`{"result": "error", "message": "could not parse request body"}`))
	// 	// if err != nil {
	// 	// 	log.Warnf("manifoldSNSPayload writing body response failed: %+v | %+v", err, string(body))
	// 	// }

	// 	// log.Warnf("manifoldSNSPayload parsing body json failed: %+v | %+v", err, unquotedBody)

	// 	// return
	// }

	log.Print("")
	log.Print("⬇  ⬇  ⬇  ⬇  ⬇  ⬇")
	// log.Printf("manifoldSNSMsg: %#v", manifoldSNSMsg)
	log.Printf("manifoldSNSPayload: %#v", manifoldSNSPayload)
	log.Print("⬆  ⬆  ⬆  ⬆  ⬆  ⬆")
	log.Print("")

	message := strings.Builder{}
	message.WriteString(fmt.Sprintf("📢 new manifold claim: %s", manifoldSNSPayload.Link))
	message.WriteString(fmt.Sprintf("\n\n%+v", manifoldSNSPayload))

	// notify.SendMessageViaTelegram(message.String(), viper.GetInt64("notifications.telegram.chat_id"), "", 0, nil) //nolint:contextcheck

	// w.WriteHeader(http.StatusOK)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write([]byte(`{"result": "success"}`))
	if err != nil {
		log.Warnf("manifoldSNS writing response failed: %s", err)
	}
}
