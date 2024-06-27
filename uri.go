package main

import (
	"net/url"
)

//  //  //

type URI string

func (uri *URI) Set(value string) error {
	if _, err := url.ParseRequestURI(value); err != nil {
		return err
	}

	*uri = URI(value)

	return nil
}

func (uri *URI) String() string {
	return string(*uri)
}
