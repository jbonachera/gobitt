package models

import "github.com/jbonachera/gobitt/tracker/config"

type DatabasePlugin interface {
	// This should initialize database session, store it in the context, and eventually managed the database schema.
	Start(cfg config.Config)
	FindPeerList(limit int, hash string) ([]Peer, error)
	FindPeer(limit int, hash string) (Peer, error)
	UpsertPeer(peer Peer)
	RemovePeer(peer Peer)

	FindTorrent(hash string) (Torrent, error)
	UpsertTorrent(t Torrent)
	RemoveTorrent(t Torrent)
}
