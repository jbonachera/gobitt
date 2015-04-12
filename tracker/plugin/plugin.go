package plugin

import "github.com/jbonachera/gobitt/tracker/config"
import "github.com/jbonachera/gobitt/tracker/models"

var plugin = make(map[string]DatabasePlugin)

func RegisterDatabasePlugin(name string, newPlugin DatabasePlugin) {
	plugin[name] = newPlugin
}

func GetDatabasePlugin(name string) DatabasePlugin {
	return plugin[name]
}

type DatabasePlugin interface {
	// This should initialize database session, store it in the context, and eventually managed the database schema.
	Start(cfg config.Config)
	FindPeerList(limit int, hash string) ([]models.Peer, error)
	FindPeer(limit int, hash string) (models.Peer, error)
	UpsertPeer(peer models.Peer)
	RemovePeer(peer models.Peer)

	FindTorrent(hash string) (models.Torrent, error)
	UpsertTorrent(t models.Torrent)
	RemoveTorrent(t models.Torrent)
}
