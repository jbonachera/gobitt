package tracker

import (
	"fmt"
	"github.com/jbonachera/gobitt/tracker/models"
	"github.com/jbonachera/gobitt/tracker/repo"
	"log"
	"net/http"
)

func ScrapeHandler(c models.ApplicationContext, w http.ResponseWriter, r *http.Request) {
	log.Print("New scrape request from " + r.RemoteAddr)
	w.Header().Set("Content-type", "text/plain")

	req, protocol_error := repo.NewScrapeRequestFromHTTPRequest(r)
	if protocol_error != nil {
		http.Error(w, protocol_error.Error(), http.StatusBadRequest)
		log.Fatal("Peer " + r.RemoteAddr + " sent an invalid request")
	}

	file := repo.GetTorrent(c, req.Info_hash)
	scrapeFile := repo.NewScrapeFileFromTorrent(file)
	scrapeAnswer := repo.NewScrapeAnswerString(req.Info_hash, scrapeFile)
	fmt.Fprintf(w, scrapeAnswer)
}
