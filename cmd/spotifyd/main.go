package main

import (
	"context"

	"github.com/ThatTomPerson/spotigo/internal/app/spotify"
)

func main() {
	srv := spotify.New()

	srv.Connect(context.Background())

	// srv.Authenticate()
}
