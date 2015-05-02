package tracker

import (
	hex "encoding/hex"
	"fmt"
	. "github.com/gorilla/feeds"
	"github.com/jbonachera/gobitt/tracker/context"
	"github.com/jbonachera/gobitt/tracker/repo"
	"net/http"
	"time"
)

func RSSHandler(c context.ApplicationContext, w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	feed := &Feed{
		Title:       "VX-Labs tracker",
		Link:        &Link{Href: "http://www.vx-labs.net/"},
		Description: "torrents RSS",
		Author:      &Author{},
		Created:     now,
		Copyright:   "",
	}
	torrents := repo.ListTorrents(c)
	for _, torrent := range torrents {
		pretty_hash := hex.EncodeToString([]byte(torrent.Hash))
		feed.Items = append(feed.Items, &Item{
			Title:   pretty_hash,
			Link:    &Link{Href: c.Config.RSS.RSSBaseDownloadURL + pretty_hash + ".torrent"},
			Created: now,
		})
	}
	atom, _ := feed.ToAtom()
	fmt.Fprintf(w, atom)
}
