package repo

import (
	"github.com/jbonachera/gobitt/tracker/models"
	"gopkg.in/mgo.v2/bson"
)

// NewTorrent returns a new "torrent" object from a string representing
// the file's hash. The completed counter is set to 0, as the file was just
// created.
func NewTorrent(c models.ApplicationContext, hash string) models.Torrent {
	return models.Torrent{hash, NewPeerListFromHash(c, hash), 0}
}

func GetTorrent(c models.ApplicationContext, hash string) models.Torrent {
	var t models.Torrent
	sess := c.Session.Copy()
	defer sess.Close()
	db_files := sess.DB("tracker").C("files")
	if err := db_files.Find(bson.M{"hash": hash}).One(&t); err != nil {
		t = NewTorrent(c, hash)
	} else {
		t.Peers = NewPeerListFromHash(c, hash)
	}
	return t
}

func SaveTorrent(c models.ApplicationContext, t models.Torrent) {
	sess := c.Session.Copy()
	defer sess.Close()
	db_files := sess.DB("tracker").C("files")
	db_files.Upsert(bson.M{"hash": t.Hash}, &t)

}
