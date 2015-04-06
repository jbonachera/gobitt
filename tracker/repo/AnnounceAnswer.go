package repo

import (
	"github.com/jbonachera/gobitt/tracker/models"
)

// NewAnnounceAnswer creates a new announceAnswer object from 3 parameters:
//  * interval
//  * minInterval
//  * a peer array, which is the list of all the peers sharing the hash wanted
//    by the current peer.
func NewAnnounceAnswer(interval, minInterval int, peers []models.Peer) *models.AnnounceAnswer {
	peers_clean := make([]interface{}, len(peers))
	for index, item := range peers {
		peer := make(map[string]interface{}, 3)
		peer["id"] = item.PeerId
		peer["ip"] = item.Ip
		peer["port"] = item.Port
		peers_clean[index] = peer
	}
	return &models.AnnounceAnswer{interval, minInterval, peers_clean}
}
