package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/jinzhu/copier"
	"github.com/spf13/viper"
)

func StartWebsocketServer(queueWS *chan *collections.Event) {
	listenHost := viper.GetString("server.host")
	listenPort := viper.GetUint("server.port")
	listenOn := fmt.Sprint(listenHost) + ":" + fmt.Sprint(listenPort)

	gbl.Log.Infof("starting websocket server on %s\n", listenOn)

	http.ListenAndServe(listenOn, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			// handle error
			gbl.Log.Error(err)
		}

		go func() {
			defer conn.Close()

			for event := range *queueWS {

				var publishEvent collections.Event

				copier.Copy(&publishEvent, &event)

				// salira := publishEvent.Collection.SaLiRa
				// publishEvent.Collection.Sal = salira.Value()
				// publishEvent.Collection.SaLiRa = nil

				// wannabeFloor := publishEvent.Collection.ArtificialFloor
				// publishEvent.Collection.MovingAverageValue = wannabeFloor.Value()
				// publishEvent.Collection.MovingAverage = nil

				// msg := fmt.Sprintf("%+v", event)
				marshalledEvent, _ := json.Marshal(publishEvent)
				err = wsutil.WriteServerText(conn, marshalledEvent)
				if err != nil {
					// handle error
					gbl.Log.Error(err)
				}

			}

			// msg, op, err := wsutil.ReadClientData(conn)
			// if err != nil {
			// 	// handle error
			// }
			// fmt.Printf("server msg: %s | op: %s | err: %s\n", string(msg), string(op), err)
		}()
	}))
}
