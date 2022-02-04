package main

//nolint
import (
	"github.com/mmcdole/gofeed"
	"html"
)

// getTitle επιστρέφει τον τίτλο της είδησης.
func getTitle(item *gofeed.Item) string {
	title := html.UnescapeString(item.Title) //  αποκωδικοποίησε τα unicode HTML chars, αν υπάρχουν

	return title
}
