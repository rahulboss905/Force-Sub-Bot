package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Abishnoi69/Force-Sub-Bot/FallenSub/config"
	"github.com/Abishnoi69/Force-Sub-Bot/FallenSub/modules"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/webhook"
)

func main() {
	token := config.Token
	publicURL := os.Getenv("WEBHOOK_URL") // e.g. https://your-app.onrender.com

	b, err := gotgbot.NewBot(token, nil)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	updater := ext.NewUpdater(modules.Dispatcher, nil)

	mux := http.NewServeMux()
	wh := webhook.New(b, updater, "/bot"+b.Token)
	mux.Handle("/bot"+b.Token, wh)

	go func() {
		log.Println("Starting webhook server on port 10000...")
		err := http.ListenAndServe(":10000", mux)
		if err != nil {
			log.Fatalf("Failed to start webhook server: %v", err)
		}
	}()

	_, err = b.SetWebhook(publicURL+"/bot"+b.Token, nil)
	if err != nil {
		log.Fatalf("Failed to set webhook: %v", err)
	}

	log.Println("Webhook set. Bot is live.")
	_, _ = b.SendMessage(config.LoggerId, "Bot started with webhook ✔️", nil)

	updater.Idle()
}
