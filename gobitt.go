// Gobitt is an implementation of a file tracker, for the bittorent
// p2p protocol.
package gobitt

import (
	"github.com/jbonachera/gobitt/tracker"
	"github.com/jbonachera/gobitt/tracker/config"
	"github.com/jbonachera/gobitt/tracker/context"
	"github.com/jbonachera/gobitt/tracker/plugin"
	_ "github.com/jbonachera/gobitt/tracker/plugin/database"
	"log"
	"net/http"
)

type contextFunc func(c context.ApplicationContext, w http.ResponseWriter, r *http.Request)

func (h contextHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.f(h.c, w, r)
}

type contextHandler struct {
	c context.ApplicationContext
	f contextFunc
}

func Start() {
	cfg := config.GetConfig()
	context := context.ApplicationContext{}
	context.Database = plugin.GetDatabasePlugin("mongodb")
	context.Database.Start(cfg)

	log.Print("Running on: " + cfg.Server.BindAddress + ":" + cfg.Server.Port)
	http.Handle("/announce", contextHandler{context, tracker.AnnounceHandler})
	http.Handle("/scrape", contextHandler{context, tracker.ScrapeHandler})
	http.ListenAndServe(cfg.Server.BindAddress+":"+cfg.Server.Port, nil)
}
