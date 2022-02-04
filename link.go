package main

import "github.com/mmcdole/gofeed"

// getLink επιστρέφει το link της είδησης από το εκάστοτε website, αν υπάρχει.
func getLink(item *gofeed.Item) (link string) {
	if len(item.Links) > 0 {
		link = item.Links[0]
	}

	return link
}
