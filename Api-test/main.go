package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ResponseData struct {
	Total int                  `json:"total"`
	Data  []ResponseDataEmiten `json:"data"`
}

type ResponseDataEmiten struct {
	CodeEmiten string `json:"code_emiten"`
}

func main() {
	url := "https://tulabi.com:3801/v3.7.1/devidend"

	response := getData(url)

	res := ResponseData{}

	if err := json.Unmarshal(response, &res); err != nil {
		fmt.Printf("Couuld not unmarshal response bytes %v", err)
	}
	fmt.Println(res.Data)

}

func getData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)

	if err != nil {
		log.Printf("couldn't request a dadjoke. %v", err)
	}

	var bearer = "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjMwMjE4MSwiZGF0YSI6eyJpZCI6MzAyMTgxLCJ1dWlkIjoiNWNmNWFkOWQtMzUyOC00ODRiLWIzNTYtOWY3NmJjZGY0ZjNmIiwiZW1haWwiOiJuYW5kaWFwbGF5c3RvcmVAZ21haWwuY29tIiwiY3JlYXRlZF9hdCI6IjIwMjEtMTAtMDIgMTI6Mjk6MDQiLCJ1cGRhdGVkX2F0IjoiMjAyMi0wNi0yOSAxOTozNDoxMCIsImRlbGV0ZWRfYXQiOm51bGwsInJvbGVfaWQiOjIsImlzX3ZlcmlmaWVkIjoxLCJ0d29fZmFjdG9yX2F1dGgiOjAsInR3b19mYWN0b3Jfc2VjcmV0IjpudWxsLCJpc19sb2dnZWRfaW4iOjEsImlzX2RlbGV0ZWQiOjAsImNyZWF0ZWRfYnkiOm51bGwsInVwZGF0ZWRfYnkiOm51bGwsImlzX290cCI6MSwiYXR0ZW1wdCI6MiwiYXR0ZW1wdF9lbWFpbCI6MCwiZmluZ2VyX3ByaW50IjoxLCJhdHRlbXB0X290cCI6MCwiYXR0ZW1wdF9waW4iOjAsIm5hbWUiOiJOYW5kaWEgUmFoYWRpYW4gTnVncmFoYSIsInRyYWRlcl90eXBlIjoicGVyc29uYWwiLCJpc19yZXNldF9wYXNzd29yZCI6MSwiaXNfdmVyaWZpZWRfa3ljIjoyLCJyb2xlX25hbWUiOiJUcmFkZXIifSwiaWF0IjoxNjU2NTA2MzExfQ.alHe5RtDxBxYsZuDTWKsb5EPn-P9n0_paF3uG84BlMo"

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "application/json")
	request.Header.Add("Authorization", bearer)

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
