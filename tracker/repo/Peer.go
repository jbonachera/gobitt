package repo

import (
	"github.com/jbonachera/gobitt/tracker/context"
	"github.com/jbonachera/gobitt/tracker/models"
	"log"
	"time"
)

// NewPeer returns a new Peer, and initializes the LastSeen parameter
// to the current time.
func NewPeer(peerId, hash, port string,
	downloaded, uploaded, left int,
	ip []string,
	numwant int) models.Peer {
	return models.Peer{peerId,
		hash,
		downloaded,
		uploaded,
		left,
		ip,
		numwant,
		time.Now()}
}

// NewPeerFromAnnounceRequest takes an AnnounceRequest object, and extracts data
// from it, to instanciates a NewPeer object.
func NewPeerFromAnnounceRequest(a *models.AnnounceRequest) models.Peer {
	return NewPeer(a.Peer_id,
		a.Info_hash,
		a.Port,
		a.Downloaded,
		a.Uploaded,
		a.Left,
		a.Ip,
		a.NumWant)
}

func NewPeerListFromHash(c context.ApplicationContext, hash string) []models.Peer {
	var peers []models.Peer
	var err error
	// We get a list of all peers seeding this file
	if peers, err = c.Database.FindPeerList(-1, hash); err != nil {
		log.Fatal(err)
	}
	return peers
}

func RemovePeer(c context.ApplicationContext, peer models.Peer) {
	c.Database.RemovePeer(peer)
}

func SavePeer(c context.ApplicationContext, peer models.Peer) {
	c.Database.UpsertPeer(peer)
}
