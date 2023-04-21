package web

//
//  gloomberg websockets server & web ui
//
//  created with help of this awesome articly by @percybolmer 👌 thanks!
//  https://programmingpercy.tech/blog/mastering-websockets-with-go/
//

import (
	"html/template"
	"log"
	"net/http"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/totra"
)

func StartWebUI(queueWsOutTokenTransactions chan *totra.TokenTransaction) {
	// Create a Manager instance used to handle WebSocket Connections
	hub := NewHub(queueWsOutTokenTransactions)

	// load index template
	// tmpl := template.Must(template.ParseFiles("www/index.html"))

	tmplFiles := []string{"www/index.html", "www/style.html", "www/javascript.html"}

	tmpl, err := template.ParseFiles(tmplFiles...)
	if err != nil {
		gbl.Log.Error(err)
	}

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

	http.HandleFunc("/ws", hub.serveWS)

	// Serve on port :8080, fudge yeah hardcoded port
	log.Fatal(http.ListenAndServeTLS(":8080", "./home.benleb.de.crt", "./home.benleb.de.key", nil))
}
