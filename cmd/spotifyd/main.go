package main

import (
	"context"

	"gitlab.com/ThatTomPerson/spotify-go/internal/app/spotify"
)

func main() {
	srv := spotify.New()

	srv.Connect(context.Background())

	// srv.Authenticate()
}
