package main

//nolint: gofumpt, goimports
import (
	"github.com/mmcdole/gofeed" //nolint: gci, gofumpt
	"time"                      //nolint: gci, gofumpt
)

// getDate επιστρέφει την ημερομηνία που δημοσιεύτηκε η είδηση.
func getDate(item *gofeed.Item) *time.Time {
	publishedDate := item.PublishedParsed

	return publishedDate
}
