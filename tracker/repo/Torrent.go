package repo

import (
	"github.com/jbonachera/gobitt/tracker/models"
)

// NewTorrent returns a new "torrent" object from a string representing
// the file's hash. The completed counter is set to 0, as the file was just
// created.
func NewTorrent(c models.ApplicationContext, hash string) models.Torrent {
	return models.Torrent{hash, NewPeerListFromHash(c, hash), 0, 0, 0}
}

func GetTorrent(c models.ApplicationContext, hash string) models.Torrent {
	var t models.Torrent
	if t, err := c.Database.FindTorrent(hash); err != nil {
		t = NewTorrent(c, hash)
	} else {
		t.Peers = NewPeerListFromHash(c, hash)
	}
	complete, incomplete := 0, len(t.Peers)
	for _, peer := range t.Peers {
		if peer.Left == 0 {
			complete += 1
			incomplete -= 1
		}
	}

	return t
}

func SaveTorrent(c models.ApplicationContext, t models.Torrent) {
	c.Database.UpsertTorrent(t)
}
