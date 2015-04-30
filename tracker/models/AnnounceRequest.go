package models

// AnnounceRequest represents all the data a peer can include when submitting
// an announce request to the tracker.
type AnnounceRequest struct {
	Info_hash  string
	Peer_id    string
	Ip         []string
	Port       string
	Downloaded int
	Uploaded   int
	Left       int
	Event      string
	NumWant    int
	No_peer_id string
	Compact    string
}
