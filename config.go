package main

import apicaller "github.com/jman2476/go-kedex/internal/api-caller"

type config struct {
	client   apicaller.Client
	Next     *string
	Previous *string
}
