package main

import (
	"log"
	urlib "net/url"
	"strings"
)

// getBaseURL επιστρέφει το αρχικό (base URL) του website από το οποίο προέρχεται η είδηση.
// στη μορφή "site.com" -- αφαιρώντας δηλαδή το 'http(s)://www.'.
func getBaseURL(rssURL string) (source string) {
	if link, err := urlib.Parse(rssURL); err != nil {
		log.Fatal(err)
	} else {
		source = link.Host

		// Αφαίρεσε το http://www αν υπάρχει, ψάχνοντας για 'www'
		if strings.Contains(source, "www.") {
			source = strings.SplitAfter(source, "www.")[1]
		}
	}

	return source
}
