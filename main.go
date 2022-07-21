package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type Article struct {
	Title       string
	Url         string
	Description string
}

func main() {

	for i := 1; i < 5; i++ {

		// println(i)

		res, err := http.Get("https://www.detik.com/search/searchall?query=teknologi&siteid=" + strconv.Itoa(i))
		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()
		log.Println(res.StatusCode)

		if res.StatusCode != 200 {
			log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)

		if err != nil {
			log.Fatal(err)
		}

		rows := make([]Article, 0)

		doc.Find(".list-berita").Children().Each(func(i int, sel *goquery.Selection) {
			row := new(Article)

			row.Title = sel.Find(".title").Text()
			row.Url = sel.Find("a").AttrOr("href", "")
			row.Description = sel.Find("p").Text()

			rows = append(rows, *row)
		})

		bts, err := json.MarshalIndent(rows, "", " ")

		if err != nil {
			log.Fatal(err)
		}

		log.Println(string(bts))

	}

}
