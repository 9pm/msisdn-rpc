package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

// Country : model of country for extraction codes and so on
type Country struct {
	Name          string `json:"name"`
	CC1           string `json:"cc1"`
	CC2           string `json:"cc2"`
	CC            string `json:"cc"`
	Region        string `json:"region"`
	SubRegion     string `json:"sub_region"`
	RegionCode    string `json:"region_code"`
	SubRegionCode string `json:"sub_region_code"`
}

// FindCountry : return one Coutry by coutry code
func FindCountry(cc string) Country {
	csvFile, _ := os.Open("data/all.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var coutry Country

	fmt.Println(cc)

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		if cc == line[3] {
			coutry = Country{
				Name:          line[0],
				CC1:           line[1],
				CC2:           line[2],
				CC:            line[3],
				Region:        line[4],
				SubRegion:     line[5],
				RegionCode:    line[6],
				SubRegionCode: line[7],
			}
		}

	}
	fmt.Printf("Country: %s\n", coutry.Name)
	return coutry
}
