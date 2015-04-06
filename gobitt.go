// Gotracker is an implementation of a file tracker, for the bittorent
// p2p protocol.
package gobitt

import (
	"github.com/jbonachera/gobitt/tracker"
	"github.com/jbonachera/gobitt/tracker/config"
	"github.com/jbonachera/gobitt/tracker/models"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

type contextFunc func(c models.ApplicationContext, w http.ResponseWriter, r *http.Request)

func (h contextHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.f(h.c, w, r)
}

type contextHandler struct {
	c models.ApplicationContext
	f contextFunc
}

func getDatabaseString(cfg config.Config) string {
	var auth_string string
	if cfg.Database.User != "" && cfg.Database.Password != "" {
		auth_string = cfg.Database.User + ":" + cfg.Database.Password + "@"
	} else {
		auth_string = ""
	}
	return auth_string + cfg.Database.Host + ":" + cfg.Database.Port
}

func getDatabase(cfg config.Config) *mgo.Session {
	session, err := mgo.Dial(getDatabaseString(cfg))
	if err != nil {
		log.Fatal("Error while opening database")
	}
	return session
}

func Start() {
	cfg := config.GetConfig()
	session := getDatabase(cfg)
	defer session.Close()
	context := models.ApplicationContext{session}

	log.Print("Running on: " + cfg.Server.BindAddress + ":" + cfg.Server.Port)
	http.Handle("/announce", contextHandler{context, tracker.AnnounceHandler})
	http.Handle("/scrape", contextHandler{context, tracker.ScrapeHandler})
	http.ListenAndServe(cfg.Server.BindAddress+":"+cfg.Server.Port, nil)
}
