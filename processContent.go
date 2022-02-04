package main

import "fmt"

// createContent επεξεργάζεται και φτιάχνει το κείμενο που θα σταλθεί στο Discord.
func createContent(rssURL string, description string, source string) (content string) {
	// Φτιάξε το κείμενο που θα σταλθεί
	// Αν είναι από το Astronomy.com τότε μην βάλεις Description γιατί είναι η ίδια με τον τίτλο
	if rssURL == astronomyNews {
		description = ""
	}

	content = fmt.Sprintf("%s\nAπό το %s\n", description, source)

	return content
}
