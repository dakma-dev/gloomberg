package web

//
//  gloomberg websockets server & web ui
//
//  created with help of this awesome articly by @percybolmer 👌 thanks!
//  https://programmingpercy.tech/blog/mastering-websockets-with-go/
//

import (
	"html/template"
	"net"
	"net/http"
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

func StartWebUI(gb *gloomberg.Gloomberg) (*WsHub, error) {
	// Create a Manager instance used to handle WebSocket Connections
	hub := NewHub(gb)

	listenOn := &net.TCPAddr{IP: net.ParseIP(viper.GetString("web.host")), Port: viper.GetInt("web.port")}

	// load index template
	tmplFiles := []string{"www/index.tpl.html", "www/style.tpl.html", "www/javascript.tpl.html"}
	tmpl, err := template.ParseFiles(tmplFiles...)
	if err != nil {
		gbl.Log.Error(err)
	}

	// index page
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		data := map[string]string{
			"Title": "gloomberg | " + internal.GloombergVersion,
		}

		if err := tmpl.Execute(w, data); err != nil {
			gbl.Log.Error("Error executing template: ", err)
		}
	})

	// static js files (the stripping feels a bit weird...)
	http.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("./www/js"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts", http.FileServer(http.Dir("./www/fonts"))))

	// websocket endpoint
	http.HandleFunc("/ws", hub.serveWS)

	tlsConfig, err := gloomberg.GetServerTLSConfig()
	if err != nil {
		log.Warn("TLS certificate not found, using insecure connection")
	}

	// create http server
	hub.server = &http.Server{
		Addr:              listenOn.AddrPort().String(),
		ReadHeaderTimeout: 2 * time.Second,
		Handler:           nil,
		TLSConfig:         tlsConfig,
	}

	// start http server
	log.Debugf("starting web ui on %s | %+v", listenOn, hub.server)
	go func() {
		if err := hub.server.ListenAndServeTLS("", ""); err != nil {
			log.Fatal(err)
		}
	}()

	return hub, nil
}
