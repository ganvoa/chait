package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ganvoa/chait/internal"
	"github.com/ganvoa/chait/pkg/logger"
)

func main() {
	l, err := logger.New()
	if err != nil {
		panic(fmt.Errorf("error logger init %v", err))
	}

	ml := l.WithName("MAIN")

	apiKey := os.Getenv("OPENAI_API_KEY")

	if apiKey == "" {
		ml.Error(errors.New("must set env variable OPENAI_API_KEY"), "error reading api key")
		os.Exit(1)
	}

	var configFile string
	var outputFile string
	flag.StringVar(&configFile, "config", "", "config file")
	flag.StringVar(&outputFile, "output", "", "output file")
	flag.Parse()

	if configFile == "" {
		ml.Error(errors.New("must specify a config file"), "error reading config file")
		os.Exit(1)
	}

	var outputFileW *os.File

	if outputFile != "" {
		dir, _ := filepath.Split(outputFile)

		if dir != "" {
			err := os.MkdirAll(dir, 0755)
			if err != nil {
				ml.Error(err, "error trying to create output folder")
				os.Exit(1)
			}
		}

		outputFileW, err = os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			ml.Error(err, "error trying to create output file")
			os.Exit(1)
		}
		defer outputFileW.Close()
	}

	handle, err := os.Open(configFile)

	if err != nil {
		ml.Error(err, "error reading config file")
		os.Exit(1)
	}
	defer handle.Close()

	config, err := internal.NewConfig(handle)
	if err != nil {
		ml.Error(err, "error parsing config")
		os.Exit(1)
	}

	p1 := internal.NewAiParticipant("U1", config.Chait.Rol1, apiKey, l.WithName("U1"))
	p2 := internal.NewAiParticipant("U2", config.Chait.Rol2, apiKey, l.WithName("U2"))

	ml.Info("init new chat room", "replies", config.Chait.Replies)
	cr, err := internal.NewChatRoom(config.Chait.Replies, p1, p2)
	if err != nil {
		ml.Error(err, "error initializing chat room")
		os.Exit(1)
	}

	ml.Info("start conversation")
	err = cr.StartConversation()
	if err != nil {
		ml.Error(err, "error calling StartConversation")
		os.Exit(1)
	}

	ml.Info("finished conversation")

	tr := internal.NewTableRenderer(cr.Conversation, os.Stdout)
	err = tr.Render()
	if err != nil {
		ml.Error(err, "error rendering conversation")
		os.Exit(1)
	}

	if outputFile != "" {
		tr := internal.NewTableRenderer(cr.Conversation, outputFileW)
		err = tr.Render()
		if err != nil {
			ml.Error(err, "error rendering the conversation to the output file")
			os.Exit(1)
		}
	}
}
