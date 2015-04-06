package models

// Torrent represents an announced file on the current tracker.
// It possesses a unique Hash, and a "completed" counter, which is
// incremented each time a peer send a "completed" event. This can be used
// to make statistics.
type Torrent struct {
	Hash      string
	Completed int
}

// NewTorrent returns a new "torrent" object from a string representing
// the file's hash. The completed counter is set to 0, as the file was just
// created.
func NewTorrent(hash string) *Torrent {
	return &Torrent{hash, 0}
}
