/*
 * @author Daniel Popov <lalabuy9948@gmail.com>
 * @copyright <Do whatever you want>
 */

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

// Operator : model for operator
type Operator struct {
	MCCMNC   string `json:"mccmnc"`
	MCC      string `json:"mcc"`
	MNC      string `json:"mnc"`
	Country  string `json:"country"`
	Operator string `json:"operator"`
}

// Dialing : model for dialing code
type Dialing struct {
	CCAlpha string `json:"ccalpha"`
	Code    string `json:"code"`
}

// FindCountry : return one Coutry by coutry code
func FindCountry(cc string) Country {
	csvFile, err := os.Open("data/countries.csv")
	if err != nil {
		fmt.Println(err)
	}
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

// FindMNO : function which searching MNO identifiers in dataset
func FindMNO(countryName string) []string {
	csvFile, err := os.Open("data/operator-list.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var mnos []string

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		if countryName == line[3] {
			if line[0] != "" {
				mnos = append(mnos, line[0])
			} else {
				mnos = append(mnos, "NA")
			}
		}
	}

	return mnos
}

//GetDialing : find dialing code by country code alpha
func GetDialing(ccalpha string) string {
	csvFile, err := os.Open("data/dialing-codes.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var dialingCode string

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		if ccalpha == line[0] {
			if line[1] != "" {
				dialingCode = line[1]
			} else {
				dialingCode = "NA"
			}
		}
	}

	return dialingCode
}
