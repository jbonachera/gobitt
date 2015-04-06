package models

import "errors"

// Those vars represents all the error messages possible
// in the Bitorrent protocol. These messages should be
// along some specific error code, but those overlaps with HTTP
// error code (ex: using "101" to announce an invalid request).
var (
	ErrInvalidRequestType = errors.New("Invalid request type: client request was not a HTTP GET")
	ErrMissingInfoHash    = errors.New("Missing info_hash")
	ErrMissingPeerID      = errors.New("Missing peer_id")
	ErrMissingPort        = errors.New("Missing port")
	ErrInvalidInfoHash    = errors.New("Invalid infohash: infohash is not 20 bytes long")
	ErrInvalidPeerID      = errors.New("Invalid peerid: peerid is not 20 bytes long")
	ErrInvalidNumWant     = errors.New("Invalid numwant. Client requested more peers than allowed by tracker")
	ErrEventlessRequest   = errors.New("Client sent an eventless request before the specified time")
	ErrGeneric            = errors.New("Unknown error")
)
