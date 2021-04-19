package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/slienlaffa/OtraReferenciaStarWars/communications"
)

type Satellite struct {
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}
type Satellites struct {
	Satellites []Satellite `json:"satellites"`
}
type Coordinates struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}
type Response struct {
	Position Coordinates `json:"position"`
	Message  string      `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`Post - /topsecret<br> Post - /topsecret/{satellite_name}`))
}
func handlerTopSecret(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data Satellites
	if r.Method == "POST" {
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if len(data.Satellites) != 3 {
			http.Error(w, "must be data for 3 satellites", http.StatusBadRequest)
			return
		}
		var resp Response
		resp.Position.X, resp.Position.Y = communications.GetLocation(data.Satellites[0].Distance, data.Satellites[1].Distance, data.Satellites[2].Distance)
		resp.Message = communications.GetMessage(data.Satellites[0].Message, data.Satellites[1].Message, data.Satellites[2].Message)
		if strings.Trim(resp.Message, " ") == "" {
			http.Error(w, "404", http.StatusNotFound)
			return
		}
		js, err := json.Marshal(resp)
		if err != nil {
			log.Print("ERROR: Can't encode to json: ", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write(js)
	}
}

var sat_split = make(map[string]Satellite)

func handlerTopSecretSplit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	title := r.URL.Path[len("/topsecret_split/"):]
	switch r.Method {
	case "GET":
		if len(sat_split) == 3 {
			var resp Response
			resp.Position.X, resp.Position.Y = communications.GetLocation(sat_split["kenobi"].Distance, sat_split["skywalker"].Distance, sat_split["sato"].Distance)
			resp.Message = communications.GetMessage(sat_split["kenobi"].Message, sat_split["skywalker"].Message, sat_split["sato"].Message)
			if strings.Trim(resp.Message, " ") == "" {
				http.Error(w, "404", http.StatusNotFound)
				return
			}
			js, err := json.Marshal(resp)
			if err != nil {
				log.Print("ERROR: Can't encode to json: ", err)
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.Write(js)
			return
		} else {
			http.Error(w, "No hay suficiente información aún", http.StatusNotFound)
		}
	case "POST":
		title = strings.ToLower(title)
		if title == "kenobi" || title == "skywalker" || title == "sato" {
			var data Satellite
			err := json.NewDecoder(r.Body).Decode(&data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			sat_split[title] = data
		} else {
			title = "wrong name"
		}
	}
	w.Write([]byte(title))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/topsecret", handlerTopSecret)
	http.HandleFunc("/topsecret_split/", handlerTopSecretSplit)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
