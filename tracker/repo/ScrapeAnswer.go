package repo

import (
	"github.com/jbonachera/gobitt/tracker/models"
	"github.com/zeebo/bencode"
	"log"
)

// NewScrapeAnswer produces a ScrapeAnswer object from a string representing a
// file's hash, and a ScrapeFile object, representing the file.
// This function can not be used to produce a ScrapeAnswer representing more
// than one file: this is not something I need right now.
func NewScrapeAnswer(hash string, scrapedata *models.ScrapeFile) *models.ScrapeAnswer {
	var answer models.ScrapeAnswer
	files := make(map[string]interface{}, 1)
	files[hash] = scrapedata
	answer.Files = files
	return &answer
}

func NewScrapeAnswerString(hash string, scrapedata *models.ScrapeFile) string {
	s := NewScrapeAnswer(hash, scrapedata)
	data, err := bencode.EncodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

// NewScrapeFile returns a ScrapeFile object from 3 arguments, representing
// statistics about a file.
func NewScrapeFile(complete, downloaded, incomplete int) *models.ScrapeFile {
	return &models.ScrapeFile{complete, downloaded, incomplete}
}
