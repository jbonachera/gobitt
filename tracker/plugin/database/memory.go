package database

import (
	"github.com/jbonachera/gobitt/tracker/models"
	"github.com/jbonachera/gobitt/tracker/plugin"
	"log"
	"time"
)

func init() {
	log.Println("Registering memory database plugin")
	plugin.RegisterDatabasePlugin("memory", &MemoryDatabasePlugin{})
}

type MemoryDatabasePlugin struct {
	torrents []models.Torrent
	peers    []models.Peer
}

func (self *MemoryDatabasePlugin) Start() {
	self.torrents = []models.Torrent{}
	self.peers = []models.Peer{}
	go purgePeerRunner(self)

}

func purgePeerRunner(db *MemoryDatabasePlugin) {
	tick := time.Tick(60 * time.Second)
	for {
		select {
		case <-tick:
			db.PurgePeers(3600 * time.Second)
		}
	}
}

func (self *MemoryDatabasePlugin) FindPeerList(limit int, hash string) ([]models.Peer, error) {
	var peers []models.Peer
	counter := 0
	for _, peer := range self.peers {
		if counter >= limit {
			break
		}
		if peer.Hash == hash {
			peers = append(peers, peer)
			counter += 1
		}
	}
	return peers, nil
}

func (self *MemoryDatabasePlugin) FindPeer(limit int, hash string) (models.Peer, error) {
	var peer models.Peer
	for _, current_peer := range self.peers {
		if peer.Hash == hash {
			peer = current_peer
			break
		}
	}
	return peer, nil
}

func (self *MemoryDatabasePlugin) UpsertPeer(peer models.Peer) {
	var found bool
	for index, current_peer := range self.peers {
		if current_peer.Hash == peer.Hash {
			self.peers[index] = peer
			found = true
			break
		}
	}
	if !found {
		self.peers = append(self.peers, peer)
	}

}

func (self *MemoryDatabasePlugin) RemovePeer(peer models.Peer) {
	var index_to_remove int
	for index, current_peer := range self.peers {
		if peer.Hash == current_peer.Hash {
			index_to_remove = index
			break
		}
	}
	self.peers = append(self.peers[:index_to_remove], self.peers[index_to_remove+1:]...)
}

func (self *MemoryDatabasePlugin) FindTorrent(hash string) (models.Torrent, error) {
	var torrent models.Torrent
	for _, current_torrent := range self.torrents {
		if torrent.Hash == hash {
			torrent = current_torrent
			break
		}
	}
	return torrent, nil
}

func (self *MemoryDatabasePlugin) RemoveTorrent(torrent models.Torrent) {
	var index_to_remove int
	for index, current_torrent := range self.torrents {
		if torrent.Hash == current_torrent.Hash {
			index_to_remove = index
			break
		}
	}
	self.torrents = append(self.torrents[:index_to_remove], self.torrents[index_to_remove+1:]...)
}
func (self *MemoryDatabasePlugin) UpsertTorrent(torrent models.Torrent) {
	var found bool
	for index, current_torrent := range self.torrents {
		if current_torrent.Hash == torrent.Hash {
			self.torrents[index] = torrent
			found = true
			break
		}
	}
	if !found {
		self.torrents = append(self.torrents, torrent)
	}
}
func (self *MemoryDatabasePlugin) PurgePeers(maxAge time.Duration) {
	for _, peer := range self.peers {
		if time.Now().After(peer.LastSeen) {
			log.Println("Removing old peer " + peer.PeerId)
			self.RemovePeer(peer)
		}
	}
}
