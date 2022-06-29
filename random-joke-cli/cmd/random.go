/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "get-joke",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		jokeTerm, _ := cmd.Flags().GetString("term")

		if jokeTerm != "" {
			getRandomJokeWithTerm(jokeTerm)
		} else {
			getRandomJoke()
		}
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
	randomCmd.PersistentFlags().String("term", "", "A search term for a dad joke.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

type SearchResult struct {
	Results    json.RawMessage `json:"results"`
	SearchTerm string          `json:"search_term"`
	Status     int             `json:"status"`
	TotalJokes int             `json:"total_jokes"`
}

func getRandomJoke() {
	fmt.Println("Get random joke :P")
	url := "https://icanhazdadjoke.com/"
	responseBytes := getJokeData(url)

	joke := Joke{}

	if err := json.Unmarshal(responseBytes, &joke); err != nil {
		fmt.Printf("Couuld not unmarshal response bytes %v", err)
	}
	// fmt.Println(joke)
	fmt.Println(string(joke.Joke))
}

func getJokeData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)

	if err != nil {
		log.Printf("couldn't request a dadjoke. %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("couldn't not make request . %v", err)

	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("couldn't read response body. %v", err)
	}

	return responseBytes
}

func getRandomJokeWithTerm(jokeTerm string) {
	log.Printf("You searched for a joke with the term: %v", jokeTerm)
	total, result := getJokeDataWithTerm(jokeTerm)
	randomiseJokeList(total, result)
	fmt.Println(result)
}

func getJokeDataWithTerm(jokeTerm string) (totalJokes int, jokeList []Joke) {
	url := fmt.Sprintf("https://icanhazdadjoke.com/search?term=%s", jokeTerm)
	responseBytes := getJokeData(url)

	jokeListRaw := SearchResult{}

	if err := json.Unmarshal(responseBytes, &jokeListRaw); err != nil {
		log.Printf("Could not unmarshal reponseBytes. %v", err)
	}

	jokes := []Joke{}
	if err := json.Unmarshal(jokeListRaw.Results, &jokes); err != nil {
		log.Printf("Could not unmarshal reponseBytes. %v", err)
	}
	fmt.Println(string(jokeListRaw.Results))

	return jokeListRaw.TotalJokes, jokes
}

func randomiseJokeList(length int, jokeList []Joke) {
	rand.Seed(time.Now().UnixNano())

	min := 0
	max := length - 1

	if length <= 0 {
		err := fmt.Errorf("no jokes found this term")
		fmt.Printf(err.Error())
	} else {
		randomNum := min - rand.Intn(max-min)
		fmt.Println(jokeList[randomNum].Joke)
	}
}
