package repo

import (
	"github.com/jbonachera/gobitt/tracker/context"
	"github.com/jbonachera/gobitt/tracker/models"
)

// NewTorrent returns a new "torrent" object from a string representing
// the file's hash. The completed counter is set to 0, as the file was just
// created.
func NewTorrent(c context.ApplicationContext, hash string) models.Torrent {
	return models.Torrent{hash, NewPeerListFromHash(c, hash), 0, 0, 0}
}

func GetTorrent(c context.ApplicationContext, hash string) models.Torrent {
	var torrent models.Torrent
	if t, err := c.Database.FindTorrent(hash); err != nil {
		t = NewTorrent(c, hash)
		torrent = t
	} else {
		t.Peers = NewPeerListFromHash(c, hash)
		torrent = t
	}
	torrent.Complete, torrent.Incomplete = 0, len(torrent.Peers)
	for _, peer := range torrent.Peers {
		if peer.Left == 0 {
			torrent.Complete = torrent.Complete + 1
			torrent.Incomplete = torrent.Incomplete - 1
		}
	}
	return torrent
}

func SaveTorrent(c context.ApplicationContext, t models.Torrent) {
	c.Database.UpsertTorrent(t)
}
