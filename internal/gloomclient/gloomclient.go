package gloomclient

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

// ConnectToServer connects to a btv server to receive events from
func ConnectToServer(connectHost string, queueEvents *chan *collections.Event) {
	ctx := context.Background()

	gbl.Log.Infof("starting websocket client")

	conn, _, _, err := ws.Dial(ctx, connectHost)
	if err != nil {
		// handle error
		gbl.Log.Errorf("error connecting to server: %s", err)
		return
	}

	fmt.Printf("client conn: %+v | err: %s\n", conn, err)

	go func() {
		defer conn.Close()

		for {
			// msg, op, err := wsutil.ReadServerData(conn)

			// fmt.Printf("client msg1 msg: %+v\n", msg)

			// if err != nil {
			// 	// handle error
			// 	fmt.Printf("client msg1: %s | op: %s | err: %s\n", string(msg), string(op), err)
			// }
			// fmt.Printf("client msg1: %s\n", string(msg))

			// msg, err := wsutil.ReadServerText(conn)
			// m := []wsutil.Message{}
			msg, err := wsutil.ReadServerText(conn)
			if err != nil {
				// handle error
				fmt.Printf("client msg2: %v | err: %s\n", msg, err)
			}

			var event *collections.Event
			if err := json.Unmarshal(msg, &event); err != nil {
				log.Fatal(err)
			}

			gbl.Log.Infof("client event: %v\n", event)

			*queueEvents <- event
		}
	}()
}
