package main

import (
	"github.com/spudfkc/gotify"
)

type User struct {
	Radio       *Radio
	SpotifyUser *gotify.User
}

type Radio struct {
	Seeds []Seed
}

type Seed struct {
	Content Content
}
