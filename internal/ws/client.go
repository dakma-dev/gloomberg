//nolint:all
package ws

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

// StartWsClient connects to a btv server to receive events from.
func StartWsClient(url string, queueEvents *chan *totra.TokenTransaction) {
	ctx := context.Background()

	gbl.Log.Infof("starting websocket client | url: %s", url)

	// TODO remove sleep and use a proper wait
	time.Sleep(1 * time.Second)

	conn, _, _, err := ws.Dial(ctx, url)
	if err != nil {
		// handle error
		gbl.Log.Errorf("error connecting to server: %s", err)
		return
	}

	gbl.Log.Infof("client conn: %+v | err: %s\n", conn, err)

	go func() {
		defer conn.Close()

		for {
			// msg, op, err := wsutil.ReadServerData(conn)

			// msg, err := wsutil.ReadServerText(conn)
			// m := []wsutil.Message{}
			msg, err := wsutil.ReadServerText(conn)
			if err != nil {
				// handle error
				gbl.Log.Errorf("client msg2: %v | err: %s\n", msg, err)
			}

			var ttx *totra.TokenTransaction
			if err := json.Unmarshal(msg, &ttx); err != nil {
				log.Fatal(err)
			}

			gbl.Log.Infof("client event: %v\n", ttx)

			*queueEvents <- ttx
		}
	}()
}
