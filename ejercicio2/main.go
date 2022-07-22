package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2"

func main() {

	// id pokemon
	id := 1

	nombre, err := pokemonByID(id)
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}

	// respuesta
	fmt.Println("Nombre del pokemon: ", nombre)
}

func pokemonByID(id int) (nombre string, err error) {

	// consulta Get http
	client := &http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf(baseURL+"/pokemon-form/%d", id)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	// body a bytes
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// volcando datos json a nuestra estructura
	posts := PokemonForm{}
	err = json.Unmarshal(body, &posts)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return posts.Name, nil
}
