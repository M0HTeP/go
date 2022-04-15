package main

import (
	"context"
	"fmt"
	"github.com/shomali11/slacker"
	"log"
	"os"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3398622637444-3420057801264-83ngRBGlfJOOceesO2TOG8nz")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03B7L3NV55-3389640585750-9324c170eece11c47d83bf80048b3ce5af5fad4c025df06391de3b1a0378342a")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
}
