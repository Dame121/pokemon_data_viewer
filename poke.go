package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Pokemon struct {
	Name   string
	Height int
	Weight int
}

func getPokemon(pokemon string) {
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while connecting to pokemon api")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Error to fetch pokemon")
	}
	//decoding the json response
	var p Pokemon
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&p)
	if err != nil { //if error is not nil
		fmt.Println("Error decoding ")
	}

	fmt.Printf("Name: %s\nHeight: %.1fm \nWeight: %.1fkg\n", p.Name, float32(p.Height)/10, float32(p.Weight)/10)

}

func main() {
	var pokemon string
	fmt.Print("Enter the pokemon you want:")
	fmt.Scan(&pokemon)

	getPokemon(pokemon)
}
