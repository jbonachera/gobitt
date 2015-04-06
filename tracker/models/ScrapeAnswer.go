package models

// ScrapeAnswer represents the data which must be sent to the peer
// which summited a "scrape" request (a scrape request is just basically
// asking for some stats about a file).
// ScrapeAnswer represents a list of files, and can be Bencoded.
type ScrapeAnswer struct {
	Files map[string]interface{} `bencode:"files"`
}

// ScrapeFile represents statistics about a file, which has to be included in a
// ScrapeAnswer object, when a peer is asking for statistics.
type ScrapeFile struct {
	Complete   int `bencode:"complete"`
	Downloaded int `bencode:"downloaded"`
	Incomplete int `bencode:"incomplete"`
}
