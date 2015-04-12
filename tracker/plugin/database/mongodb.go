package database

import (
	"github.com/jbonachera/gobitt/tracker/config"
	"github.com/jbonachera/gobitt/tracker/models"
	"github.com/jbonachera/gobitt/tracker/plugin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func init() {
	log.Println("Registering mongodb database plugin")
	plugin.RegisterDatabasePlugin("mongodb", &MongoDBDatabasePlugin{})
}

type MongoDBDatabasePlugin struct {
	dbSession *mgo.Session
}

func getDatabaseConfString(cfg config.Config) string {
	confString := "?"
	if cfg.Database.Connect != "" {
		confString += "connect=" + cfg.Database.Connect + "&"
	}
	if cfg.Database.AuthSource != "" {
		confString += "authSource=" + cfg.Database.AuthSource + "&"
	}

	if cfg.Database.AuthMecanism != "" {
		confString += "authMecanism=" + cfg.Database.AuthMecanism + "&"
	}
	if cfg.Database.GSSAPIService != "" {
		confString += "gssapiService=" + cfg.Database.GSSAPIService + "&"
	}
	if cfg.Database.MaxPoolSize != "" {
		confString += "maxPoolSize=" + cfg.Database.MaxPoolSize + "&"
	}
	return confString
}

func getDatabaseString(cfg config.Config) string {
	var auth_string string
	if cfg.Database.User != "" && cfg.Database.Password != "" {
		auth_string = cfg.Database.User + ":" + cfg.Database.Password + "@"
	} else {
		auth_string = ""
	}
	return auth_string + cfg.Database.Host + ":" + cfg.Database.Port + getDatabaseConfString(cfg)
}
func getDatabase(cfg config.Config) *mgo.Session {
	log.Printf("Initiating connection to MongoDB")
	session, err := mgo.Dial(getDatabaseString(cfg))
	if err != nil {
		if err := session.Ping(); err != nil {
			log.Fatal("Error while opening connexion to MongDB!")
		}
	}
	log.Println("Sucessfully connected to MongoDB!")
	return session
}

func (self *MongoDBDatabasePlugin) Start(cfg config.Config) {
	self.dbSession = getDatabase(cfg)
}

func (self *MongoDBDatabasePlugin) FindPeerList(limit int, hash string) ([]models.Peer, error) {
	var peers []models.Peer
	table := "peers"
	sess := self.dbSession.Copy()
	defer sess.Close()
	db := sess.DB("tracker").C(table)

	err := db.Find(bson.M{"hash": hash}).All(&peers)
	return peers, err
}

func (self *MongoDBDatabasePlugin) FindPeer(limit int, hash string) (models.Peer, error) {
	var peer models.Peer
	table := "peers"
	sess := self.dbSession.Copy()
	defer sess.Close()
	db := sess.DB("tracker").C(table)

	err := db.Find(bson.M{"hash": hash}).One(&peer)
	return peer, err
}

func (self *MongoDBDatabasePlugin) UpsertPeer(peer models.Peer) {
	table := "peers"
	sess := self.dbSession.Copy()
	defer sess.Close()
	db := sess.DB("tracker").C(table)
	db.Upsert(bson.M{"ip": peer.Ip, "port": peer.Port, "peerid": peer.PeerId, "hash": peer.Hash}, &peer)
}

func (self *MongoDBDatabasePlugin) RemovePeer(peer models.Peer) {
	table := "peers"
	sess := self.dbSession.Copy()
	defer sess.Close()
	db := sess.DB("tracker").C(table)
	db.Remove(bson.M{"hash": peer.Hash, "peerid": peer.PeerId, "ip": peer.Ip, "port": peer.Port})
}

func (self *MongoDBDatabasePlugin) FindTorrent(hash string) (models.Torrent, error) {
	var t models.Torrent
	table := "files"
	sess := self.dbSession.Copy()
	defer sess.Close()
	db := sess.DB("tracker").C(table)
	err := db.Find(bson.M{"hash": hash}).One(&t)
	return t, err
}

func (self *MongoDBDatabasePlugin) RemoveTorrent(t models.Torrent) {
	table := "files"
	sess := self.dbSession.Copy()
	defer sess.Close()
	db := sess.DB("tracker").C(table)
	db.Remove(bson.M{"hash": t.Hash})
}

func (self *MongoDBDatabasePlugin) UpsertTorrent(t models.Torrent) {
	table := "files"
	sess := self.dbSession.Copy()
	defer sess.Close()
	db := sess.DB("tracker").C(table)
	db.Upsert(bson.M{"hash": t.Hash}, &t)
}
