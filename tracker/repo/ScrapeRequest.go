package repo

import (
	"github.com/jbonachera/gobitt/tracker/models"
	"net/http"
)

// NewScrapeRequest  returns a ScrapeRequest object from a raw hash parameter.
func NewScrapeRequest(
	info_hash string,
) *models.ScrapeRequest {
	return &models.ScrapeRequest{info_hash}
}

// NewScrapeRequest takes an http request object, extracts parameters from it,
// and creates a ScrapeRequest object.
func NewScrapeRequestFromHTTPRequest(r *http.Request) (*models.ScrapeRequest, error) {
	hash, protocol_error := validateArg(r.URL.Query().Get("info_hash"), 20, models.ErrInvalidInfoHash, models.ErrMissingInfoHash)
	if protocol_error != nil {
		return nil, protocol_error
	}
	return NewScrapeRequest(hash), nil
}
