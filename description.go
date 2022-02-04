package main

import ( //nolint
	"github.com/k3a/html2text" // nolint
	"github.com/mmcdole/gofeed"
	"html"    //nolint
	"strings" //nolint
)

// getDescription επιστρέφει τη περιγραφή της είδησης, αν υπάρχει στο RSS.
func getDescription(item *gofeed.Item) string {
	description := html2text.HTML2Text(html.UnescapeString(item.Description))

	// Δες αν υπάρχει κενή γραμμή στην περιγραφή και αφαίρεσε το κείμενα μετά από αυτή.
	// Αυτό το κάνουμε γιατί συνήθως σε τέτοιες περιπτώσεις το κείμενο που ακολουθεί είναι άσχετο με την είδηση.
	if strings.Contains(description, "\n") {
		description = strings.SplitAfter(description, "\n")[0]
	}

	return description
}
