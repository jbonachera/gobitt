package tracker

import (
	"fmt"
	"github.com/jbonachera/gobitt/tracker/context"
	"github.com/jbonachera/gobitt/tracker/models"
	"github.com/jbonachera/gobitt/tracker/repo"
	"log"
	"net/http"
)

func processAnnounceStarted(c context.ApplicationContext, w http.ResponseWriter, torrent models.Torrent) {
	// We get a list of all peers seeding this file
	answer := repo.NewCompactAnnounceAnswerString(600, 100, // I should move that to a tracker.Configuration file.
		repo.NewPeerListFromHash(c, torrent.Hash))
	repo.SaveTorrent(c, torrent)
	fmt.Fprintf(w, answer)
}

func processAnnounceCompleted(c context.ApplicationContext, torrent models.Torrent) {
	torrent.Downloaded += 1
	repo.SaveTorrent(c, torrent)
}

func processAnnounceStopped(c context.ApplicationContext, peer models.Peer) {
	repo.RemovePeer(c, peer)
}

func AnnounceHandler(c context.ApplicationContext, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/plain")
	// Ensure that the announce request is conform to the protocol spec
	announceRequest, protocol_error := repo.NewAnnounceRequestFromHTTPRequest(r)
	if protocol_error != nil {
		http.Error(w, protocol_error.Error(), http.StatusBadRequest)
		log.Println("Peer " + r.RemoteAddr + " sent an invalid request")
		return
	}
	// If the announce is correct, we create a peer object from the request
	peer := repo.NewPeerFromAnnounceRequest(announceRequest)
	repo.SavePeer(c, peer)
	torrent := repo.GetTorrent(c, peer.Hash)

	// Process event
	switch announceRequest.Event {
	case "completed":
		log.Print(peer.Ip[0] + " completed a download")
		processAnnounceCompleted(c, torrent)
	case "stopped":
		log.Print(peer.Ip[0] + " stopped announcing")
		processAnnounceStopped(c, peer)
	default: //"Started" is managed in the default case
		log.Print("New announce from " + peer.Ip[0])
		processAnnounceStarted(c, w, torrent)
	}
}
