package main

import (
	"context"

	"github.com/sirupsen/logrus"
	"ttp.sh/spotigo/internal/app/spotify"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	srv := spotify.New()

	srv.Connect(context.Background())

	// srv.Authenticate()
}
