package main

//nolint
import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"log"
	"time" //nolint
)

// postNewsPerFeed "παρσάρει" το RSS και συλλέγει όλες τις απαραίτητες πληροφορίες που χρειαζόμαστε γύρω από την είδηση.
func postNewsPerFeed(fp *gofeed.Parser, oneHourAgo time.Time, rssURL string) {
	url, err := fp.ParseURL(rssURL)
	if err != nil {
		log.Println("Δεν μπορεί να φτιάξει το URL του RSS από το", rssURL)
		return
	}
	feed, _ := url, err

	// Πάρε το βασικό website link από το οποίο προέρχεται το RSS feed
	source := getBaseURL(rssURL)

	// Βρόχος επανάληψης, στον οποίο κάθε φορά επιλέγεται μία καινούρια αστρο-είδηση
	for _, item := range feed.Items {
		title := getTitle(item)
		description := getDescription(item)
		link := getLink(item)
		publishedDate := getDate(item)
		image := getImage(item, rssURL)

		// 	Αν η είδηση έχει δημοσιευτεί εντός των προηγούμενων 60 λεπτών, τότε στείλε τη στο Discord
		if publishedDate.After(oneHourAgo) {
			content := createContent(rssURL, description, source)

			// 	Για να σταλθεί η είδηση στο discord, θα πρέπει να υπάρχει η μεταβλητή "WEBHOOK_DISCORD" στο περιβάλλον
			//  εκτέλεσης του προγράμματος. Διαφορετικά, θα εμφανίσει τις ειδήσεις, χωρίς (!) να τις στείλει κάπου.
			if webhook == "" {
				fmt.Printf("Τίτλος: %s\nΠεριγραφή: %s\nLink: %s\nΕικόνα %s\n---\n", title, content, link, image) //nolint: forbidigo
			} else {
				postContentToDiscord(link, title, content, image)
			}

			// Περίμενε γιατί αν είναι πολλά τα νέα, παίζει να φας ban απο τον server για spam
			time.Sleep(10 * time.Second) //nolint: gomnd
		}
	}
}
