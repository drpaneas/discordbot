package main

import (
	"log"
)

// postContentToDiscord δημοσιεύει την είδηση στον Discord server
// και κάνει print αν τα κατάφερε ή όχι.
func postContentToDiscord(link string, title string, content string, image string) {
	if err := postDiscord(link, title, content, image); err != nil {
		log.Printf("%s\t%s\t%v\n", failed, title, err)
	} else {
		log.Printf("%s\t%s\n", succeed, title)
	}
}
