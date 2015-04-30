package models

import "time"

// Peer represents a Bitorrent peer: a unique PeerID+IP+Port combo,
// requesting a peer list for a unique file hash.
// Peer can announce statistics:
//   * the number of Downloaded and Uploaded bytes
//   * How much data is left before having a full file
// Peer also announce how many peers does he want, and this number
// has a limit of 50, hardcoded.
// A Peer also have a 'LastSeen' parameter, which can be used to expire
// old row from the persistent storage.
type Peer struct {
	PeerId     string
	Hash       string
	Port       string
	Downloaded int
	Uploaded   int
	Left       int
	Ip         []string
	NumWant    int
	LastSeen   time.Time
}
