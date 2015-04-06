package models

// AnnounceAnswer represents the data a tracker must send back to a peer,
// using the "dict" mode (aka the non-compact mode).
// This object can then be encoded using BEncode, and sent back to the peer.
type AnnounceAnswer struct {
	Interval    int           `bencode:"interval"`
	MinInterval int           `bencode:"min interval"`
	Peers       []interface{} `bencode:"peers"`
}
