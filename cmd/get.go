package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"tuitune/config"

	"github.com/lrstanley/go-ytdlp"
)


var cmdGet = cobra.Command{
	Use: "get url",
	Short: "Downloads the media which the given url points to",
	Long: "TODO",
	Aliases: []string{"g"},
}

func init() {
	cmdGet.Run = func(cmd *cobra.Command, args []string) {
		for _, url := range args {
			// TODO: quality argument
			get(url, 0)
		}

	}

}

func get(url string, quality uint) {
	dl := ytdlp.New().ProgressFunc(
		250 * time.Millisecond,
		func(update ytdlp.ProgressUpdate) {
			fmt.Printf(
				"\rDownloading \"%s\" @ %s\033[K",
				*update.Info.Title,
				update.PercentString(),
			)
		},
	).Paths(config.MusicDirectory).ExtractAudio().AudioQuality(fmt.Sprintf("%d", quality))



	res, err := dl.Run(context.Background(), url)
	fmt.Println()
	if err != nil {
		panic(err)
	}

	infos, err := res.GetExtractedInfo()
	if err != nil {
		panic(err)
	}

	if res.ExitCode == 0 {
		for _, info := range infos {
			fmt.Printf("Successfully downloaded \"%s\"\n", *info.Title)
		}
	} else {
		for _, info := range infos {
			fmt.Printf("Failed to download \"%s\"\n", *info.Title)
		}

		fmt.Printf("Exited with code %d\n", res.ExitCode)

	}

}

