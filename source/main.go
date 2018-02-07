package main

/*
Bot-Bot V0.7
Author: Dylan Blanchard
*/

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"github.com/bwmarrin/discordgo"
	"github.com/tsukinai/CommanD-Bot"
)

// Current BotSession //
var botSession *discordgo.Session

// Pre-load and start bot //
// - Loads maps
// - Create new BotSession
func init() {
	// Create new BotSession //
	botSession = CommanD_Bot.New("Bot MzU3OTUwMTc3OTQ1OTc2ODM5.DOYtIQ.oa9Fqrl8RlhyunioLrmfItnpBkE")
}

// Program entry point //
// - Waits for CTRL-C or terminate signal to close program
// - On terminate signal close bot session
func main() {
	// Wait until CTRL-C or other terminate signal is received. //
	fmt.Println("Bot-Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Close bot session //
	err := botSession.Close()
	if err != nil {
		log.Println(err)
	}

	// End of main //
	log.Println("End of main.")
	return
}
