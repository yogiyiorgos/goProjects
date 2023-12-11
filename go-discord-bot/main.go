package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "os/signal"
    "strings"
    "syscall"

    "github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

const KuteGoAPIURL = "https://kutego-api-xxxxx.ew.a.run.app"

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token") // Package `flag` implements command line flag parsing
	flag.Parse()
}

func main() {
	// Create a new Discord session using the provided bot token
	dg, err := discordgo.New("Bot" + Token)
	if err != nil {
		fmt.Println("Error creating Discord session, ", err)
		return
	}

	// Register the messageCreate func as a callback for messageCreate events
	dg.AddHandler(messageCreate)

	// In this case, we are only interested in receiving message events
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection, ", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1) // Create a new channel of type os.Signal with a buffer size 1
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session
	dg.Close()
}

type Gopher struct {
	Name string `json: "name"`
}

// This function will be called (AddHandler above) every time a new message is created
// or any channel that the authenticated bot has access to
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Good practice: ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!gopher" {
		// Call the KuteGo API and retrieve Dr. Who Gopher
		response, err := http.Get(KuteGoAPIURL + "/gopher/" + "dr-who")
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			_, err = s.ChannelFileSend((m.ChannelID, "dr-who.png", response.Body)
			if err
		}
	}
}
