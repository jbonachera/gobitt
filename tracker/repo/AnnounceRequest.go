package repo

import (
	"github.com/jbonachera/gobitt/tracker/models"
	"net"
	"net/http"
	"strconv"
)

// NewAnnounceRequest returns an AnnounceRequest from raw data.
func NewAnnounceRequest(
	info_hash,
	peer_id,
	ip,
	port string,
	downloaded int,
	uploaded int,
	left int,
	event string,
	numwant int,
	no_peer_id,
	compact string,
) *models.AnnounceRequest {
	return &models.AnnounceRequest{info_hash, peer_id, ip, port, downloaded,
		uploaded, left, event, numwant, no_peer_id, compact}
}

// validateArg is a private helper to validate arguments passed to
// NewAnnounceRequestFromHTTPRequest
func validateArg(a string, size int,
	sizeError, missingError error) (string, error) {
	var err error
	if a == "" {
		err = missingError
	} else if size > 0 && len(a) != size {
		err = sizeError
	}
	return a, err
}

// NewAnnounceRequestFromHTTPRequest returns a AnnounceRequest from an HTTP
// request object, extracting the data from the URL GET parameters.
// It also checks the parameters the peer sent, and return a gobitt error
// when a mandatory argument is missing.
// It also currently hardcode the numwant parameter to a maximum value of 50
// instead of returning an error code the client. This may change in the future.
func NewAnnounceRequestFromHTTPRequest(r *http.Request) (*models.AnnounceRequest, error) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		panic(err)
	}
	hash, protocol_error := validateArg(r.URL.Query().Get("info_hash"), 20, models.ErrInvalidInfoHash, models.ErrMissingInfoHash)
	if protocol_error != nil {
		return nil, protocol_error
	}
	peerID, protocol_error := validateArg(r.URL.Query().Get("peer_id"), 20, models.ErrInvalidPeerID, models.ErrMissingPeerID)
	if protocol_error != nil {
		return nil, protocol_error
	}
	port, protocol_error := validateArg(r.URL.Query().Get("port"), -1, models.ErrMissingPort, models.ErrMissingPort)
	if protocol_error != nil {
		return nil, protocol_error
	}
	downloaded, err := strconv.Atoi(r.URL.Query().Get("downloaded"))
	if err != nil {
		downloaded = 0
	}
	uploaded, err := strconv.Atoi(r.URL.Query().Get("uploaded"))
	if err != nil {
		uploaded = 0
	}
	left, err := strconv.Atoi(r.URL.Query().Get("left"))
	if err != nil {
		left = 1
	}
	numwant, err := strconv.Atoi(r.URL.Query().Get("numwant"))
	if err != nil || numwant > 50 {
		numwant = 50
	}
	return NewAnnounceRequest(
		hash,
		peerID,
		ip,
		port,
		downloaded,
		uploaded,
		left,
		r.URL.Query().Get("event"),
		numwant,
		r.URL.Query().Get("no_peer_id"),
		r.URL.Query().Get("compact"),
	), protocol_error
}
