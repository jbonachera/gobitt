package repo

import (
	"github.com/jbonachera/gobitt/tracker/models"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

// NewPeer returns a new Peer, and initializes the LastSeen parameter
// to the current time.
func NewPeer(peerId, hash, port string,
	downloaded, uploaded, left int,
	ip string,
	numwant int) models.Peer {
	return models.Peer{peerId,
		hash,
		port,
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

func NewPeerListFromHash(c models.ApplicationContext, hash string) []models.Peer {
	var peers []models.Peer
	sess := c.Session.Copy()
	defer sess.Close()
	db_peers := sess.DB("tracker").C("peers")
	// We get a list of all peers seeding this file
	if err := db_peers.Find(bson.M{"hash": hash}).All(&peers); err != nil {
		log.Fatal(err)
	}
	return peers
}

func RemovePeer(c models.ApplicationContext, peer models.Peer) {
	sess := c.Session.Copy()
	defer sess.Close()
	db_peers := sess.DB("tracker").C("peers")
	db_peers.Remove(bson.M{"hash": peer.Hash, "peerid": peer.PeerId, "ip": peer.Ip, "port": peer.Port})
}

func SavePeer(c models.ApplicationContext, peer models.Peer) {
	sess := c.Session.Copy()
	defer sess.Close()
	db_peers := sess.DB("tracker").C("peers")
	db_peers.Upsert(bson.M{"ip": peer.Ip, "port": peer.Port, "peerid": peer.PeerId, "hash": peer.Hash}, &peer)
}
