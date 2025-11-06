package main

import (
	"context"
	"flag"
	"fmt"
	"strings"

	mb "go.uploadedlobster.com/musicbrainzws2"
)

var Version = "0.0"

func main() {

	flag.Parse()

	args := flag.Args()

	for i := range args {
		fmt.Printf("%s\n", args[i])
	}

	if (len(args) == 0) {
		return
	}

	if (args[0] == "search" || args[0] == "s") {
		query := strings.Join(args[1:], " ")
		search(query)
	}

}

func search(query string) {

	appInfo := mb.AppInfo{
		Name: "tuitune",
		Version: Version,
		URL: "https://github.com/ajoyfuljester/tuitune",
	}


	client := mb.NewClient(appInfo)
	defer client.Close()

	ctx := context.Background()

	searchFilter := mb.SearchFilter{
		Query: query,
	}

	paginator := mb.Paginator{
		Limit: 12,
	}

	res, err := client.SearchWorks(ctx, searchFilter, paginator)
	if err != nil {
		panic(err)
	}



}
