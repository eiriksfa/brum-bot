package brum

import (
	"brum-bot/internal/app/civ"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"text/tabwriter"
	"text/template"

	"github.com/bwmarrin/discordgo"
	"github.com/urfave/cli/v2"
)

type customWriter struct {
	s *discordgo.Session
	m *discordgo.MessageCreate
	w bytes.Buffer
}

func (e customWriter) Write(p []byte) (int, error) {
	n, err := e.w.Write(p)
	outStr := "```" + e.w.String() + "```"
	e.s.ChannelMessageSend(e.m.ChannelID, outStr)

	if err != nil {
		return n, err
	}
	if n != len(p) {
		return n, io.ErrShortWrite
	}
	return len(p), nil
}

func Brum(Token string) {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	msg := strings.Split(m.Content, " ")

	if msg[0] != "!brum" && msg[0] != "!b" {
		return
	}

	cli.HelpPrinter = func(w io.Writer, templ string, data interface{}) {
		funcMap := template.FuncMap{
			"join": strings.Join,
		}
		t := template.Must(template.New("help").Funcs(funcMap).Parse(templ))

		buf := new(bytes.Buffer)
		out := tabwriter.NewWriter(buf, 1, 8, 2, ' ', 0)
		t.Execute(out, data)

		outStr := "```" + buf.String() + "```"
		s.ChannelMessageSend(m.ChannelID, outStr)
	}

	cli.OsExiter = func(code int) {

	}

	w := &customWriter{s: s, m: m, w: bytes.Buffer{}}
	cli.ErrWriter = w

	app := &cli.App{
		Name:      "BrumBot",
		HelpName:  "contrive",
		Usage:     "bils discord bot",
		UsageText: "!brum (or !b) [global options] command [command options] [arguments...]",
		Commands: []*cli.Command{
			{
				Name:  "pong",
				Usage: "sends ping",
				Action: func(c *cli.Context) error {
					s.ChannelMessageSend(m.ChannelID, "ping")
					return nil
				},
			},
			{
				Name:  "ping",
				Usage: "sends pong",
				Action: func(c *cli.Context) error {
					s.ChannelMessageSend(m.ChannelID, "pong")
					return nil
				},
			},
			{
				Name:      "civ",
				Aliases:   []string{"c"},
				Usage:     "civlization related commands",
				UsageText: "!brum (or !b) civ [global options] command [command options] [arguments...]",
				Subcommands: []*cli.Command{
					{
						Name:    "assign",
						Usage:   "assigns each player a nation",
						Aliases: []string{"a"},
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "count",
								Value:   "1",
								Aliases: []string{"c"},
								Usage:   "number of nations for each player",
							},
							&cli.StringFlag{
								Name:    "ranks",
								Value:   "AB",
								Aliases: []string{"r"},
								Usage:   "Ranks to pull nations from",
							},
						},
						Action: func(c *cli.Context) error {
							count := c.Int("count")
							s.ChannelMessageSend(m.ChannelID, civ.Assign(c.Args().Slice(), c.String("ranks"), count)) //
							return nil
						},
					},
					{
						Name:  "rankings",
						Usage: "Print out a list of leader rankings",
						Action: func(c *cli.Context) error {
							s.ChannelMessageSend(m.ChannelID, civ.Rankings()) //
							return nil
						},
					},
					{
						Name:  "leaders",
						Usage: "Print out a list of leaders",
						Action: func(c *cli.Context) error {
							s.ChannelMessageSend(m.ChannelID, civ.Leaders(c.Args().Slice())) //
							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(msg)
	if err != nil {
		log.Println(err)
	}

}
