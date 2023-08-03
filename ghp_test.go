package main

import (
	"log"
	"os"
	"testing"

	"hello/components"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

var ghPagePrefix string

// The main function is the entry point where the app is configured and started.
// It is executed in 2 different environments: A client (the web browser) and a
// server.
func _main() {
	// The first thing to do is to associate the hello component with a path.
	//
	// This is done by calling the Route() function,  which tells go-app what
	// component to display for a given path, on both client and server-side.
	app.Route("/", &components.Hello{})

	// Once the routes set up, the next thing to do is to either launch the app
	// or the server that serves the app.
	//
	// When executed on the server-side, RunWhenOnBrowser() does nothing, which
	// lets room for server implementation without the need for precompiling
	// instructions.
	app.RunWhenOnBrowser()

	err := app.GenerateStaticWebsite("docs", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
		Resources:   app.GitHubPages(ghPagePrefix),
		Image:       "/web/icon.png",
		Icon: app.Icon{
			Default:    "/web/icon.png",
			Large:      "/web/icon.png",
			AppleTouch: "/web/icon.png",
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}

func TestPublish(t *testing.T) {
	ghPagePrefix = os.Getenv("ghPagePrefix")
	t.Logf(ghPagePrefix)
	_main()
	t.Logf("Site built")
}
