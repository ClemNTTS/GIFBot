package GIFBot

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	CommandIDs = []string{}
	HelloGifs  = []string{
		"https://tenor.com/view/dog-cute-hello-hellothere-gif-20141152",
		"https://tenor.com/view/hi-hello-gif-1314135106863776295",
		"https://tenor.com/view/hello-gif-24076043",
		"https://tenor.com/view/hello-hi-minion-gif-13004117825953885603",
	}
	popularGIFs = []string{
		"https://tenor.com/view/dog-lol-gif-14866323",
		"https://tenor.com/view/sad-pikachu-cry-gif-13276698",
		"https://tenor.com/view/angry-mad-frustrated-rage-gif-17414662",
		"https://tenor.com/view/crying-rain-tears-flood-funny-gif-15007517",
		"https://tenor.com/view/spongebob-rainbow-imagination-gif-12141797",
		"https://tenor.com/view/facepalm-disappointed-star-trek-gif-15481725",
		"https://tenor.com/view/salute-yes-sir-soldier-respect-gif-15342227",
		"https://tenor.com/view/cat-dance-dance-cat-cat-ai-cat-ai-cat-dance-gif-16943432931705998786",
		"https://tenor.com/view/wow-gif-20343480",
		"https://tenor.com/view/shhh-shush-silence-nose-gif-17895433",
		"https://tenor.com/view/shut-up-gif-25474198",
		"https://tenor.com/view/danny-devito-fuck-you-angry-annoyed-raging-gif-4276067",
		"https://tenor.com/view/the-moves-gif-13395927",
		"https://tenor.com/view/rickroll-roll-rick-never-gonna-give-you-up-never-gonna-gif-22954713",
		"https://tenor.com/view/hazbik-mortal-combat-gif-3795006771601208806",
		"https://tenor.com/view/cat-wif-hat-catwifhat-cwif-gif-17193762039885744356",
		"https://tenor.com/view/meme-gif-21111852",
		"https://tenor.com/view/satan-dancing-gif-20743835",
		"https://tenor.com/view/take-me-lucifer-lucifer-hell-take-me-take-gif-17353217",
		"https://giphy.com/gifs/oliver-jamie-sausage-xbgwMAYvi7PtC",
	}
	fuckGifs = []string{
		"https://tenor.com/view/fuck-fuck-you-middle-finger-middle-finger-gif-15294280",
		"https://tenor.com/view/baby-girl-middle-finger-mood-screw-you-leave-me-alone-gif-10174031",
		"https://tenor.com/view/blah-shut-up-whatever-not-listening-nonsense-gif-16883389",
		"https://tenor.com/view/shutupangry-gif-26036331",
		"https://tenor.com/view/shut-up-gif-25474198",
		"https://tenor.com/view/shut-up-picard-murphs33-star-trek-tng-gif-27039542",
		"https://tenor.com/view/shhh-quiet-annoyed-keep-quiet-ishaa-saha-gif-12899289",
	}
	byeGifs = []string{
		"https://tenor.com/view/bye-baby-bad-teefs-gif-9520448450563971390",
		"https://tenor.com/view/bye-teletubbies-gif-14425062827418128030",
		"https://tenor.com/view/forest-gump-wave-hi-hello-howdy-gif-17500258755908763204",
		"https://tenor.com/view/edds-world-goodbye-chat-eddsworld-gif-900911743961647749",
		"https://tenor.com/view/veeunus-spongebob-spongebob-squarepants-squidward-done-gif-24674884",
		"https://tenor.com/view/bye-im-out-offline-ghosting-good-night-gif-16784070",
		"https://giphy.com/gifs/spongebob-squarepants-leave-smooth-48FhEMYGWji8",
		"https://giphy.com/gifs/despicable-me-minions-goodbye-9eM1SWnqjrc40",
		"https://giphy.com/gifs/bye-good-go-kart-8J1ijiQV76fGm9Y1bO",
	}
	reactGifs = []string{
		"https://tenor.com/view/wow-oh-my-god-omg-shocked-surprised-gif-15089848",
		"https://tenor.com/view/owen-wilson-wow-marley-and-me-smooth-hd-hq-gif-977282450591946462",
		"https://tenor.com/view/ohhhahhhhh-wow-whoa-amazed-gif-14254886",
		"https://tenor.com/view/surprised-sorprendido-shaquille-oneal-gif-23222312",
		"https://tenor.com/view/excited-hockey-kid-yeah-gif-19976923",
		"https://tenor.com/view/thanks-awesome-gif-14611487",
		"https://tenor.com/view/borat-borat-very-nice-verynice-thumbs-up-gif-25080066",
		"https://tenor.com/view/nice-the-rock-sus-the-rock-sus-meme-louis-thrane-gif-25854498",
	}

	chokedGifs = []string{
		"https://tenor.com/view/shocked-joey-friends-gasp-shook-gif-7187491",
		"https://tenor.com/view/crying-cry-emotional-gif-13202027",
		"https://tenor.com/view/arifn13-candiace-dillard-shock-surprised-shocked-gif-24030455",
		"https://tenor.com/view/shocked-surprised-gasp-what-cat-shock-gif-635629308990545194",
		"https://tenor.com/view/mike-smith-trailer-park-boys-bubbles-glasses-blink-gif-4011183380204339784",
		"https://tenor.com/view/dog-dog-glasses-dog-glasses-wtf-wtf-dog-wtf-gif-11114902363326536726",
	}
)

// launch the bot
func Launch() {
	godotenv.Load(".env")
	key := os.Getenv("GIFBot")
	session, err := discordgo.New("Bot " + key)
	if err != nil {
		fmt.Println("Session creation failed : ", err)
		return
	}
	session.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages | discordgo.IntentsMessageContent

	session.AddHandler(ready)
	session.AddHandler(messageCreate)

	err = session.Open()
	if err != nil {
		fmt.Println("Error at opening session : ", err)
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc

	session.Close()
}

// handle messages
func messageCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.ApplicationCommandData().Name {
	case "hello":
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: HelloGifs[rand.Intn(4)],
			},
		})
	case "chocked":
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: chokedGifs[rand.Intn(len(chokedGifs))],
			},
		})
	case "bye":
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: byeGifs[rand.Intn(len(byeGifs))],
			},
		})
	case "react":
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: reactGifs[rand.Intn(len(reactGifs))],
			},
		})
	case "random":
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: popularGIFs[rand.Intn(len(popularGIFs)-1)],
			},
		})
	case "fuck-you":
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fuckGifs[rand.Intn(len(fuckGifs)-1)],
			},
		})
	case "gif":
		message := i.ApplicationCommandData().Options[0].StringValue()
		fmt.Println(message)
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: GifRequest(MessageToBot(message)),
			},
		})
	}
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	fmt.Println("Bot Ready")

	HelloCommand := &discordgo.ApplicationCommand{
		Name:        "hello",
		Description: "Send an Hello Gif",
	}

	ChockedCommand := &discordgo.ApplicationCommand{
		Name:        "chocked",
		Description: "If you are chocked",
	}

	byeCommand := &discordgo.ApplicationCommand{
		Name:        "hello",
		Description: "Say bye with a Gif",
	}

	reactCommand := &discordgo.ApplicationCommand{
		Name:        "react",
		Description: "React to a nice tchat with a gif",
	}

	RandomCommand := &discordgo.ApplicationCommand{
		Name:        "random",
		Description: "Send a random popular Gif",
	}

	fuck_you_Command := &discordgo.ApplicationCommand{
		Name:        "fuck-you",
		Description: "Insult your friend",
	}

	GifCommand := &discordgo.ApplicationCommand{
		Name:        "gif",
		Description: "Answer with a gif related with your message",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "message",
				Description: "Le message à interpréter",
				Required:    true,
			},
		},
	}

	commands := []*discordgo.ApplicationCommand{HelloCommand, ChockedCommand, GifCommand, RandomCommand, fuck_you_Command, byeCommand, reactCommand}
	for _, cmd := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, s.State.Application.GuildID, cmd)
		if err != nil {
			log.Fatalf("Impossible de créer la commande '%s' : %v", cmd.Name, err)
		}
	}
}
