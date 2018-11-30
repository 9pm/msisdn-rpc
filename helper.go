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
	Name          string
	CC1           string
	CC2           string
	CC            string
	Region        string
	SubRegion     string
	RegionCode    string
	SubRegionCode string
}

// FindCountry : decription...
func FindCountry(cc int) Country {
	csvFile, _ := os.Open("data/all.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var coutry Country

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		if ToInt(line[3]) == cc {
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
	fmt.Println(coutry)
	return coutry
}
