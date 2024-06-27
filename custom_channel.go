package main

import (
	"errors"
	"net/url"
	"regexp"
	"strings"
)

//  //  //

type customChannelType uint8

const (
	customChannelUnkown customChannelType = iota
	customChannelURI
	customChannelID
	customChannelAlias
)

type customChannel struct {
	channel   string
	valueType customChannelType
}

func (cch customChannel) String() string {
	return cch.channel
}

func (cch *customChannel) Set(ch string, defWebhooksURL string) error {
	ch = strings.TrimSpace(ch)

	if ch == "" {
		return errors.New("channel required")
	}

	digitCheck := regexp.MustCompile(`^[0-9]+$`)
	aliasCheck := regexp.MustCompile(`^[a-z\-_]+$`)

	if digitCheck.MatchString(ch) {
		// Channel ID
		cch.channel = ch
		cch.valueType = customChannelID
	} else if aliasCheck.MatchString(ch) {
		// Alias of webhook URL
		cch.channel = ch
		cch.valueType = customChannelAlias
	} else if _, err := url.ParseRequestURI(ch); err == nil {
		cch.channel = ch
		cch.valueType = customChannelURI
	} else {
		cch.channel, _ = url.JoinPath(defWebhooksURL, ch)
		cch.valueType = customChannelURI
	}

	return nil
}
