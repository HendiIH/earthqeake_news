package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GempaResponse struct {
	Infogempa struct {
		Gempa struct {
			Tanggal   string `json:"Tanggal"`
			Jam       string `json:"Jam"`
			Magnitude string `json:"Magnitude"`
			Kedalaman string `json:"Kedalaman"`
			Wilayah   string `json:"Wilayah"`
		} `json:"gempa"`
	} `json:"Infogempa"`
}

func main() {
	resp, err := http.Get("https://data.bmkg.go.id/DataMKG/TEWS/autogempa.json")
	if err != nil {
		fmt.Println("No response", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		fmt.Println("Error no response", err)
		return
	}

	var GempResp GempaResponse
	err = json.Unmarshal(body, &GempResp)
	if err != nil {
		fmt.Println("Error decoding JSON", err)
		return
	}
	tanggal := GempResp.Infogempa.Gempa.Tanggal
	jam := GempResp.Infogempa.Gempa.Jam
	magnitude := GempResp.Infogempa.Gempa.Magnitude
	kedalaman := GempResp.Infogempa.Gempa.Kedalaman
	wilayah := GempResp.Infogempa.Gempa.Wilayah

	fmt.Println("Tanngal:", tanggal)
	fmt.Println("Jam:", jam)
	fmt.Println("Magnitude:", magnitude)
	fmt.Println("Kedalaman:", kedalaman)
	fmt.Println("Wilayah:", wilayah)
}
