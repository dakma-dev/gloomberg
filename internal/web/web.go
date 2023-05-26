package web

//
//  gloomberg websockets server & web ui
//
//  created with help of this awesome articly by @percybolmer 👌 thanks!
//  https://programmingpercy.tech/blog/mastering-websockets-with-go/
//

import (
	"crypto/tls"
	"html/template"
	"net"
	"net/http"
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

func StartWebUI(queueWsOutTokenTransactions chan *totra.TokenTransaction) {
	// Create a Manager instance used to handle WebSocket Connections
	hub := NewHub(queueWsOutTokenTransactions)

	listenOn := &net.TCPAddr{IP: net.ParseIP(viper.GetString("web.host")), Port: viper.GetInt("web.port")}
	certPath := viper.GetString("web.tls.certificate")
	keyPath := viper.GetString("web.tls.key")

	// load index template
	tmplFiles := []string{"www/index.html", "www/style.html", "www/javascript.html"}
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

	// websocket endpoint
	http.HandleFunc("/ws", hub.serveWS)

	// load tls certificate
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		log.Fatal(err)
	}

	// create http server
	server := &http.Server{
		Addr:              listenOn.AddrPort().String(),
		ReadHeaderTimeout: 2 * time.Second,
		Handler:           nil,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
			MinVersion:   tls.VersionTLS12,
			MaxVersion:   0,
		},
	}

	// start http server
	log.Debugf("starting web ui on %s | %+v", listenOn, server)
	log.Fatal(server.ListenAndServeTLS("", ""))
}
