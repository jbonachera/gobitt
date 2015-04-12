package database

import (
	"code.google.com/p/gcfg"
	"github.com/jbonachera/gobitt/tracker/models"
	"github.com/jbonachera/gobitt/tracker/plugin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type dBConfig struct {
	MongoDB struct {
		Host          string
		Port          string
		User          string
		Password      string
		Connect       string
		AuthSource    string
		AuthMecanism  string
		GSSAPIService string
		MaxPoolSize   string
	}
}

func init() {
	log.Println("Registering mongodb database plugin")
	plugin.RegisterDatabasePlugin("mongodb", &MongoDBDatabasePlugin{})
}

type MongoDBDatabasePlugin struct {
	dbSession *mgo.Session
}

func getDatabaseConfString(cfg dBConfig) string {
	confString := "?"
	if cfg.MongoDB.Connect != "" {
		confString += "connect=" + cfg.MongoDB.Connect + "&"
	}
	if cfg.MongoDB.AuthSource != "" {
		confString += "authSource=" + cfg.MongoDB.AuthSource + "&"
	}

	if cfg.MongoDB.AuthMecanism != "" {
		confString += "authMecanism=" + cfg.MongoDB.AuthMecanism + "&"
	}
	if cfg.MongoDB.GSSAPIService != "" {
		confString += "gssapiService=" + cfg.MongoDB.GSSAPIService + "&"
	}
	if cfg.MongoDB.MaxPoolSize != "" {
		confString += "maxPoolSize=" + cfg.MongoDB.MaxPoolSize + "&"
	}
	return confString
}

func getDatabaseString(cfg dBConfig) string {
	var auth_string string
	if cfg.MongoDB.User != "" && cfg.MongoDB.Password != "" {
		auth_string = cfg.MongoDB.User + ":" + cfg.MongoDB.Password + "@"
	} else {
		auth_string = ""
	}
	return auth_string + cfg.MongoDB.Host + ":" + cfg.MongoDB.Port + getDatabaseConfString(cfg)
}
func getDatabase(cfg dBConfig) *mgo.Session {
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

func (self *MongoDBDatabasePlugin) Start() {
	var cfg dBConfig
	gcfg.ReadFileInto(&cfg, "mongodb.ini")
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
