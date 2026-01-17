package main

import (
	"bufio"
	"context"
	"log/slog"
	"os"
	"os/signal"
	"strings"

	"github.com/catouberos/icool/config"
	"github.com/catouberos/icool/icool"
	"github.com/kkdai/youtube/v2"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	client, err := icool.Dial(ctx, cfg.Host, cfg.Room)
	if err != nil {
		slog.Error("failed to create icool client", "error", err, "host", cfg.Host)
		os.Exit(1)
	}

	ytClient := youtube.Client{}

	input := make(chan string)

	go read(input)

	for {
		select {
		case <-ctx.Done():
			return
		case id := <-input:
			video, err := ytClient.GetVideo(id)
			if err != nil {
				slog.Error("cannot get video info", "error", err)
			}

			var thumbnail = "" // default thumbnail
			if len(video.Thumbnails) > 0 {
				thumbnail = video.Thumbnails[0].URL
			}

			err = client.Queue(video.Title, thumbnail, video.ID)
			if err != nil {
				slog.Error("failed to queue", "error", err)
			}
		}
	}
}

func read(input chan string) {
	for {
		reader := bufio.NewReader(os.Stdin)
		id, err := reader.ReadString('\n')
		if err != nil {
			slog.Error("cannot read input", "error", err)
			close(input)
			return
		}

		id = strings.TrimSuffix(id, "\n")
		input <- id
	}
}
