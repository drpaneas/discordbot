package main

import (
	"strings"

	"github.com/mmcdole/gofeed"
	"mvdan.cc/xurls/v2"
)

func getImageFromHTML(item *gofeed.Item) (image string) {
	rxStrict := xurls.Strict()
	images := rxStrict.FindAllString(item.Content, -1)

	// Ψάξε σε όλα τα links να δεις αν κάποιο έχει '.jpg'
	image = getLinkWithSuffix(images, ".jpg")

	// Αν το image είναι ακόμα κενό, τότε δε βρήκε κάποιο.
	// Σε αυτή την περίπτωση, ψάξε όλα τα links για να δεις αν κάποιο έχει ίσως το 'png'
	if image == "" {
		image = getLinkWithSuffix(images, ".png")
	}

	// Αν το image είναι ακόμα κενό, τότε δε βρήκε κάποιο.
	// Σε αυτή την περίπτωση, ψάξε όλα τα links για να δεις αν κάποιο έχει ίσως το 'jpeg'
	if image == "" {
		image = getLinkWithSuffix(images, ".jpeg")
	}

	// Αν το image είναι ακόμα κενό, τότε δε βρήκε κάποιο.
	// Σε αυτή την περίπτωση, ψάξε όλα τα links για να δεις αν κάποιο έχει ίσως το 'gif'
	if image == "" {
		image = getLinkWithSuffix(images, ".gif")
	}

	if image == "" {
		images = rxStrict.FindAllString(item.Description, -1)

		// Ψάξε σε όλα τα links να δεις αν κάποιο έχει '.jpg'
		image = getLinkWithSuffix(images, ".jpg")

		// Αν το image είναι ακόμα κενό, τότε δε βρήκε κάποιο.
		// Σε αυτή την περίπτωση, ψάξε όλα τα links για να δεις αν κάποιο έχει ίσως το 'png'
		if image == "" {
			image = getLinkWithSuffix(images, ".png")
		}

		// Αν το image είναι ακόμα κενό, τότε δε βρήκε κάποιο.
		// Σε αυτή την περίπτωση, ψάξε όλα τα links για να δεις αν κάποιο έχει ίσως το 'jpeg'
		if image == "" {
			image = getLinkWithSuffix(images, ".jpeg")
		}

		// Αν το image είναι ακόμα κενό, τότε δε βρήκε κάποιο.
		// Σε αυτή την περίπτωση, ψάξε όλα τα links για να δεις αν κάποιο έχει ίσως το 'gif'
		if image == "" {
			image = getLinkWithSuffix(images, ".gif")
		}
	}

	if image == "" {
		image = getFeaturedImage(item.Link)
	}

	return image
}

func getLinkWithSuffix(images []string, suffix string) (image string) {
	for _, v := range images {
		if strings.Contains(v, suffix) {
			image = strings.SplitAfter(v, suffix)[0]

			break
		}
	}

	return image
}
