package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	mb "go.uploadedlobster.com/musicbrainzws2"

	"tuitune/config"
)


var cmdSearch = cobra.Command{
	Use: "search [-a artist] [-e] title",
	Short: "Searches musicbrainz.org for data about recordings",
	Long: "TODO",
	Aliases: []string{"s"},
}

func init() {
	var pArtist string
	var pLimit uint
	var pExact bool

	cmdSearch.Run = func(cmd *cobra.Command, args []string) {
		search(SearchParams{
			title: strings.Join(args, " "),
			artist: pArtist,
		}, pLimit, pExact)
	}


	cmdSearch.Flags().StringVarP(&pArtist, "artist", "a", "", "artist name that is credited in the recording")
	cmdSearch.Flags().UintVarP(&pLimit, "limit", "l", 12, "maximum number of results that will be showed, 25 is probably maximum")
	// i think it works word by word
	// TODO: write this in long help
	cmdSearch.Flags().BoolVarP(&pExact, "exact", "e", false, "should the search by title try to be an exact match")
}

type SearchParams struct {
	artist string;
	title string;
}

func search(params SearchParams, limit uint, isExact bool) {

	query := params.title
	if !isExact {
		query += "~"
	}

	if params.artist != "" {
		query += fmt.Sprintf(" AND artist:%s", params.artist)
	}


	fmt.Printf("Searching for \"%s\"...\n", query)




	client := mb.NewClient(config.DefaultAppInfo)
	defer client.Close()

	ctx := context.Background()

	searchFilter := mb.SearchFilter{
		Query: query,
	}


	paginator := mb.Paginator{
		Limit: int(limit),
	}

	res, err := client.SearchRecordings(ctx, searchFilter, paginator)
	if err != nil {
		panic(err)
	}


	for i, r := range res.Recordings {
		fmt.Printf("%d: %s by %s (%s)\n", i, r.Title, r.ArtistCredit.String(), r.Length.String())
	}



}

