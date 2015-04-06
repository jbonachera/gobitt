package models

// compactannounceAnswer represents a "compact" answer to an announce request.
// It means that it contains 2 strings: peers, and peers_ipv6, which are a
// binary list of ip addresses and ports of all the peers sharing the
// hash wanted by the current peer.
// This object must be encoded using BEncode, and sent to the peera in the
// announce answer.
type CompactAnnounceAnswer struct {
	Interval    int    `bencode:"interval"`
	MinInterval int    `bencode:"min interval"`
	Peers       string `bencode:"peers"`
	PeersIPv6   string `bencode:"peers_ipv6"`
}
