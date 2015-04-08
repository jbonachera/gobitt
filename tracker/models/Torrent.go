package models

// Torrent represents an announced file on the current tracker.
// It possesses a unique Hash, and a "completed" counter, which is
// incremented each time a peer send a "completed" event. This can be used
// to make statistics.
type Torrent struct {
	Hash       string
	Peers      []Peer `bson:"-"` // This is generated on the fly
	Downloaded int
	Incomplete int `bson:"-"` // This is generated on the fly
	Complete   int `bson:"-"` // This is generated on the fly
}
