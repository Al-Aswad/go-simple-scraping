package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/foolin/pagser"
)

type PageData struct {
	Perkara      []string `pagser:"#tablePerkaraAll > tbody > tr > td:nth-child(4)"`
	NomorPerkara []string `pagser:"#tablePerkaraAll > tbody > tr > td:nth-child(2)"`
	Pihak        []string `pagser:"#tablePerkaraAll > tbody > tr > td:nth-child(5)"`
	Status       []string `pagser:"#tablePerkaraAll > tbody > tr > td:nth-child(6)"`
}

type Perkara struct {
	Perkara      string
	NomorPerkara string
	Pihak        string
	Status       string
}

func main() {

	url := "http://sipp.pn-enrekang.go.id/list_perkara"
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	log.Println(res.StatusCode)

	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	//New default config
	p := pagser.New()

	//data parser model
	var data PageData
	//parse html data
	err = p.ParseReader(&data, res.Body)
	//check error
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(toJson(data.Perkara))

	rows := make([]Perkara, 0)
	row := new(Perkara)

	for i := range data.Perkara {
		row.Perkara = data.Perkara[i]
		row.NomorPerkara = data.NomorPerkara[i]
		row.Pihak = data.Pihak[i]
		row.Status = data.Status[i]

		rows = append(rows, *row)
	}

	log.Println("Tes", toJson(rows))

}

func toJson(v interface{}) string {
	data, _ := json.MarshalIndent(v, "", "\t")
	return string(data)
}
