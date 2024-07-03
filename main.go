package main

import (
	"flag"
	"fmt"
	"os"
)

//  //  //

const DISCORD_WEBHOOKS_URL string = "https://discord.com/api/webhooks/"

// discordwh [-u username] [-a avatar-url] [--dry-run] channel "message"
func main() {
	if len(os.Args) == 1 {
		showHelp()
		return
	}

	msg := messageValues{}
	dryRun := false
	args := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	args.StringVar(&msg.username, "u", "", "Username")
	args.Var(&msg.avatarUrl, "a", "Avatar URL")
	args.BoolVar(&dryRun, "dry-run", false, "Perform a trial run without send the message")

	if err := args.Parse(os.Args[1:]); err != nil {
		exitByInvalidArguments(err.Error())
	} else if err := msg.channel.Set(args.Arg(0), DISCORD_WEBHOOKS_URL); err != nil {
		exitByInvalidArguments(err.Error())
	} else if msg.message = args.Arg(1); msg.message == "" {
		exitByInvalidArguments("message required")
	} else if dryRun {
		showInfoMessage(msg)
	} else if err := sendMessage(msg); err != nil {
		fmt.Fprintf(os.Stderr, "Error to send message: %s.\n", err)
		os.Exit(1)
	}
}

func exitByInvalidArguments(error string) {
	fmt.Fprintf(os.Stderr, "Error: %s.\nType \"discordwh\" without parameters for help.\n", error)
	os.Exit(2)
}

func showHelp() {
	fmt.Print(`discordwh [-u username] [-a avatar-url] [--dry-run] channel message

Send a message using Discord Webhooks. The channel parameter maybe channel ID,
an alias or a webhook URL (partial or full).

Channel ID or alias must be defined in "discordwh.conf" file.

Use the parameter "--dry-run" to perform a trial run without send the message.

`)
}

func showInfoMessage(msg messageValues) {
	username, avatarUrl := "<not defined>", "<not defined>"

	if msg.username != "" {
		username = msg.username
	}

	if msg.avatarUrl != "" {
		avatarUrl = msg.avatarUrl.String()
	}

	fmt.Printf(
		"Discord channel: %s\nMessage: %s\nUsername: %s\nAvatar URL: %s\n",
		msg.channel,
		msg.message,
		username,
		avatarUrl,
	)
}
