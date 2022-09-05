package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/bradsec/gophercises/link"
)

func main() {
	urlFlag := flag.String("URL", "https://gophercises.com", "the URL that you want to build a site map for.")
	flag.Parse()

	fmt.Println(*urlFlag)
	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//io.Copy(os.Stdout, resp.Body)

	links, _ := link.Parse(resp.Body)
}
