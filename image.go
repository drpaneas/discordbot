package main

import (
	"github.com/badoux/goscraper"
	"github.com/mmcdole/gofeed"
)

// getImage επιστρέφει τη featured εικόνα που αντιπροσωπεύει την είδηση, αν τη βρει.
func getImage(item *gofeed.Item, rssURL string) (image string) {
	switch rssURL {
	// Μερικά από τα sites παρέχουν την εικόνα μέσω του RSS Feed
	case physicsgg:
		if value, ok := item.Extensions["media"]; ok {
			image = value["thumbnail"][0].Attrs["url"]
		}
	case newScientist:
		if value, ok := item.Extensions["media"]; ok {
			image = value["thumbnail"][0].Attrs["url"]
		}
	case spaceCom:
		if len(item.Enclosures) > 0 {
			image = item.Enclosures[0].URL
		}
	case universeToday:
		if value, ok := item.Extensions["media"]; ok {
			image = value["content"][0].Attrs["url"]
		}
	// Για τα υπόλοιπα sites που δεν παρέχουν την εικόνα μέσω RSS, προσπάθησε να τη "ψαρέψεις" μέσα από τη HTML
	default:
		image = getImageFromHTML(item)
	}

	return image
}

// getFeaturedImage επιστρέφει τη feature εικόνα διαβάζοντας την HTML.
func getFeaturedImage(link string) string {
	s, _ := goscraper.Scrape(link, 5) //nolint: gomnd
	if len(s.Preview.Images) > 0 {
		return s.Preview.Images[0]
	}

	return ""
}
