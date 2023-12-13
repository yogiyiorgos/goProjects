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

// KuteGO API is a REST API created specifically for the author's Gophers repo
// It can be locally installed and run by replacing the URL below with localhost:PORT
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
			if err != nil {
				fmt.Println(err)
			} 
		} else {
			fmt.Println("Error: Can't get dr-who Gopher! :-")
		}
	}

	if m.Content == "!random" {
		// Call the KuteGo API and retrieve a random Gopher
		response, err := http.Get(KuteGoAPIURL + "/gopher/random/")
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			_, err = s.ChannelFileSend(m.ChannelID, "random-gopher.png", response.body)
			if err != nil {
				fmt.Println("err")
			}
		} else {
			fmt.Println("Error: Can't get random Gopher! :(")
		}
	}

	if m.Content == "!gophers" {
		// Call the KuteGo API and display the list of available Gophers
		response, err := http.Get(KuteGoAPIURL + "/gophers/")
		if err != nil {
			fmt.Println(err)
		} 
		defer response.Body.Close()

		if response.StatusCode == 200 {
			// Transform our response to []byte
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Println(err)
			}

			// Put only necessary info of the JSON document in the Gopher array
			var data []Gopher
			err = json.Unmarshal(body, &data) // Unmarshal turns JSON to a Go Struct
			if err != nil {
				fmt.Println(err)
			}

			// Create a string with all of the Gopher's name and a blank line as separator
			var gopher strings.Builder
			for _, gopher := range data {
				gophers.WriteString(gopher.Name + "\n")	
			}

			// Send a text message with the list of Gophers
			_, err = s.ChannelMessageSend(m.ChannelID, gophers.String())
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Error: Can't get list of Gophers! :-(")
		}
	}
}
