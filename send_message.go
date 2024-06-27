package main

import (
	"fmt"
	"strings"

	"github.com/gtuk/discordwebhook"
)

//  //  //

const CONFIG_FILENAME string = "discordwh.conf"

func sendMessage(msg messageValues) error {
	var (
		channelURI string
		err        error
	)

	if msg.channel.valueType == customChannelURI {
		channelURI = msg.channel.channel
	} else if channelURI, err = getChannelURIByType(msg.channel); err != nil {
		return err
	}

	msgd := discordwebhook.Message{
		Content: &msg.message,
	}

	if msg.username != "" {
		msgd.Username = &msg.username
	}

	if msg.avatarUrl != "" {
		msgd.AvatarUrl = (*string)(&msg.avatarUrl)
	}

	return discordwebhook.SendMessage(channelURI, msgd)
}

func getChannelURIByType(cch customChannel) (string, error) {
	cfg, err := loadConfig(CONFIG_FILENAME)

	if err != nil {
		return "", err
	}

	if cch.valueType == customChannelAlias {
		for alias, webhookURI := range cfg.Aliases {
			if alias == cch.channel {
				return webhookURI, nil
			}
		}
	} else if cch.valueType == customChannelID {
		for _, webhookURI := range cfg.Webhooks {
			if strings.Contains(webhookURI, cch.channel) {
				return webhookURI, nil
			}
		}

		for _, webhookURI := range cfg.Aliases {
			if strings.Contains(webhookURI, cch.channel) {
				return webhookURI, nil
			}
		}
	}

	return "", fmt.Errorf("webhook not found by alias or ID (%s)", cch.channel)
}
