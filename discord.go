package main

import (
	"log"
	"os"
	"strings"

	"github.com/ecnepsnai/discord"
)

var webhook = getWebhook() //nolint: gochecknoglobals

// getWebhook διαβάζει την μεταβλητή WEBHOOK_DISCORD env var και επιστρέφει την τιμή της.
func getWebhook() string {
	webhook, present := os.LookupEnv("WEBHOOK_DISCORD")
	if !present {
		log.Println("WEBHOOK_DISCORD env variable does NOT exists")
	}

	return strings.TrimSpace(webhook)
}

func postDiscord(link, title, desc, imageLink string) (err error) {
	var pic discord.Image
	pic.URL = imageLink
	discord.WebhookURL = webhook

	// nolint
	return discord.Post(discord.PostOptions{
		Embeds: []discord.Embed{
			{
				//nolint: gomnd
				Color:       16777215,
				URL:         link,
				Title:       title,
				Description: desc,
				Thumbnail:   &pic,
			},
		},
	})
}
