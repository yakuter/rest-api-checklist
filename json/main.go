package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	// read data.json file
	content, err := os.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	// parse json
	var data Starship
	err = json.Unmarshal(content, &data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", data)

	var buf bytes.Buffer
	err = json.Indent(&buf, content, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(buf.String())

}

type Starship struct {
	Name                 string        `json:"name"`
	Model                string        `json:"model"`
	Manufacturer         string        `json:"manufacturer"`
	CostInCredits        string        `json:"cost_in_credits"`
	Length               string        `json:"length"`
	MaxAtmospheringSpeed string        `json:"max_atmosphering_speed"`
	Crew                 string        `json:"crew"`
	Passengers           string        `json:"passengers"`
	CargoCapacity        string        `json:"cargo_capacity"`
	Consumables          string        `json:"consumables"`
	HyperdriveRating     string        `json:"hyperdrive_rating"`
	Mglt                 string        `json:"MGLT"`
	StarshipClass        string        `json:"starship_class"`
	Pilots               []interface{} `json:"pilots"`
	Films                []string      `json:"films"`
	Created              time.Time     `json:"created"`
	Edited               time.Time     `json:"edited"`
	URL                  string        `json:"url"`
}
