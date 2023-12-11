## What does the bot do?
  - Displays a cute Gopher, when the user enters `!gopher` in Discord server(s)
  - Dislpays the list of available Gophers, when the user enters `!gophers`.
  - Display a random Gopher, when the user enters `!random`.

To achieve that, a Client needs to interact with Go servers. 
We will use [DiscordGo](https://github.com/bwmarrin/discordgo) library for this.

```bash
go get github.com/bwmarrin/discordgo
`
