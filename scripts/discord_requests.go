package GIFBot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// launch the bot
func Launch() {

	if err != nil {
		fmt.Println("Session creation failed : ", err)
		return
	}
	session.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages | discordgo.IntentsMessageContent

	session.AddHandler(messageCreate)

	err = session.Open()
	if err != nil {
		fmt.Println("Error at opening session : ", err)
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	session.Close()
}

// handle messages
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	answer := GifRequest(MessageToBot(m.Content))
	s.ChannelMessageSend(m.ChannelID, answer)
}
