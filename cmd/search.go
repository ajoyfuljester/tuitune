package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	mb "go.uploadedlobster.com/musicbrainzws2"

	"tuitune/config"
)

var pArtist string
var pTitle string

var cmdSearch = cobra.Command{
	Use: "search [-a artist] [-t title] the rest of the query",
	Short: "Searches musicbrainz.org for data about recordings",
	Long: "TODO",
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		search(SearchParams{
			query: strings.Join(args, " "),
			artist: pArtist,
			title: pTitle,
		})
	},
}

func init() {
	cmdSearch.Flags().StringVarP(&pArtist, "artist", "a", "", "artist name that is credited in the recording")
	cmdSearch.Flags().StringVarP(&pTitle, "title", "t", "", "title of the recording")
}

type SearchParams struct {
	query string;
	artist string;
	title string;
}

func search(params SearchParams) {

	query := params.query

	if params.title != "" {
		query += fmt.Sprintf(" AND recording:\"%s\"", params.title)
	}

	if params.artist != "" {
		query += fmt.Sprintf(" AND artist:\"%s\"", params.artist)
	}


	fmt.Printf("Searching for \"%s\"...\n", query)




	client := mb.NewClient(config.DefaultAppInfo)
	defer client.Close()

	ctx := context.Background()

	searchFilter := mb.SearchFilter{
		Query: query,
	}


	paginator := mb.Paginator{
		Limit: 25,
	}

	res, err := client.SearchRecordings(ctx, searchFilter, paginator)
	if err != nil {
		panic(err)
	}


	for i, r := range res.Recordings {
		fmt.Printf("%d: %s by %s (%s)\n", i, r.Title, r.ArtistCredit.String(), r.Length.String())
	}



}

