package models

// ScrapeRequest represents the parameters a peer can send when submitting
// a scrape request: a hash, corresponding to a torrent on the tracker.
type ScrapeRequest struct {
	Info_hash string
}
