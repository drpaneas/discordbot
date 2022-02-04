package main

//nolint
import (
	"github.com/mmcdole/gofeed" // nolint
	"time"                      //nolint
)

// Σημείωση: Βάλτε τo πρόγραμμα να τρέχει κάθε 60 λεπτά σε κάποιο CI/Scheduler

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

func main() {
	myTime := time.Now().UTC()
	oneHourAgo := myTime.Add(-time.Hour * 1)
	fp := gofeed.NewParser()

	// Επισκέψου κάθε RSS Feed και βάλε τις ειδήσεις του στο Discord, αν έχουν δημοσιευτεί εντός της ώρας
	for _, feed := range feeds {
		postNewsPerFeed(fp, oneHourAgo, feed)
	}
}
