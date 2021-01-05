package main

import (
	"encoding/json"
	"fmt"
	"log"

	"database/sql"

	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
	_ "github.com/mattn/go-sqlite3"
)

// handleMessages handles messages
func handleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "changeUrl":
		{
			fmt.Println("changing url from go")
			// Unmarshal payload
			var path string
			if len(m.Payload) > 0 {
				// Unmarshal payload
				if err = json.Unmarshal(m.Payload, &path); err != nil {
					payload = err.Error()
					return
				}
			}

			// Explore
			w.ExecuteJavaScript("window.location.href='" + path + "'")
			payload = path
			break
		}
	case "historyNav":
		{
			// Unmarshal payload
			var action string
			if len(m.Payload) > 0 {
				// Unmarshal payload
				if err = json.Unmarshal(m.Payload, &action); err != nil {
					payload = err.Error()
					return
				}
			}

			// Explore
			if action == "back" {
				w.ExecuteJavaScript("window.history.back()")
			} else {
				w.ExecuteJavaScript("window.history.forward()")
			}
			payload = action
			break
		}
	case "historyPush":
		{
			// Unmarshal payload
			var action string
			if len(m.Payload) > 0 {
				// Unmarshal payload
				if err = json.Unmarshal(m.Payload, &action); err != nil {
					payload = err.Error()
					return
				}
			}

			// Explore
			fmt.Println("add to history " + action)
			payload = action
			break
		}
	case "getUsers":
		{

			db, err := sql.Open("sqlite3", "./resources/app/main.db")
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()

			rows, err := db.Query("select * from korisnici")
			if err != nil {
				log.Fatal(err)
			}

			defer rows.Close()
			arr := []Korisnik{}
			for rows.Next() {
				var id int
				var korisnickoime string
				var datum string
				err = rows.Scan(&id, &korisnickoime, &datum)
				if err != nil {
					log.Fatal(err)
				}
				elem := Korisnik{
					Id:            id,
					Korisnickoime: korisnickoime,
					Datum:         datum,
				}
				arr = append(arr, elem)

			}
			err = rows.Err()
			if err != nil {
				log.Fatal(err)
			}
			payload = arr

			break
		}
	case "fetchUserData":
		{
			var action string
			if len(m.Payload) > 0 {
				// Unmarshal payload
				if err = json.Unmarshal(m.Payload, &action); err != nil {
					payload = err.Error()
					return
				}
			}

			db, err := sql.Open("sqlite3", "./resources/app/main.db")
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()

			rows, err := db.Query("select id, url, vremenskistambilj from povijest where korisnici_id = " + action + "")
			if err != nil {
				log.Fatal(err)
			}

			rowsBook, err := db.Query("select p.id,p.ime,p.url,k.ime from knjizneoznake p left join kategorije k on p.kategorije_id = k.id where korisnici_id = " + action + "")
			if err != nil {
				log.Fatal(err)
			}

			rowsExt, err := db.Query("select p.id, p.ime, p.opis from prosirenja p join korisnici_prosirenja_veza v on p.id = v.prosirenja_id join korisnici k on k.id = v.korisnici_id WHERE v.korisnici_id = " + action + "")
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()
			defer rowsBook.Close()
			defer rowsExt.Close()

			povijestData := []PovijestData{}
			for rows.Next() {
				var id int
				var url string
				var vremenskistambilj string
				err = rows.Scan(&id, &url, &vremenskistambilj)
				if err != nil {
					log.Fatal(err)
				}
				povijestData = append(povijestData, PovijestData{
					Id:                id,
					Url:               url,
					Vremenskistambilj: vremenskistambilj,
				})

			}

			knjizneoznakeData := []KnjiznaOznakaData{}
			for rowsBook.Next() {
				var id int
				var ime string
				var url string
				var kategorija string
				err = rowsBook.Scan(&id, &ime, &url, &kategorija)
				if err != nil {
					log.Fatal(err)
				}
				knjizneoznakeData = append(knjizneoznakeData, KnjiznaOznakaData{
					Id:         id,
					Ime:        ime,
					Url:        url,
					Kategorija: kategorija,
				})

			}

			prosirenjaData := []ProsirenjeData{}
			for rowsExt.Next() {
				var id int
				var ime string
				var opis string
				err = rowsExt.Scan(&id, &ime, &opis)
				if err != nil {
					log.Fatal(err)
				}
				prosirenjaData = append(prosirenjaData, ProsirenjeData{
					Id:   id,
					Ime:  ime,
					Opis: opis,
				})

			}

			err = rows.Err()
			if err != nil {
				log.Fatal(err)
			}
			err = rowsBook.Err()
			if err != nil {
				log.Fatal(err)
			}
			err = rowsExt.Err()
			if err != nil {
				log.Fatal(err)
			}
			arr := KorisnikData{
				Povijest:      povijestData,
				KnjizneOznake: knjizneoznakeData,
				Prosirenja:    prosirenjaData,
			}
			payload = arr

			break
		}

	}
	return
}

type KorisnikData struct {
	Povijest      []PovijestData      `json:"povijestdata"`
	KnjizneOznake []KnjiznaOznakaData `json:"knjizneoznakedata"`
	Prosirenja    []ProsirenjeData    `json:"prosirenjadata"`
}

type PovijestData struct {
	Id                int    `json:"id"`
	Url               string `json:"url"`
	Vremenskistambilj string `json:"vremenskistambilj"`
}

type KnjiznaOznakaData struct {
	Id         int    `json:"id"`
	Ime        string `json:"ime"`
	Url        string `json:"url"`
	Kategorija string `json:"kategorija"`
}

type ProsirenjeData struct {
	Id   int    `json:"id"`
	Ime  string `json:"ime"`
	Opis string `json:"opis"`
}

type Korisnik struct {
	Id            int    `json:"id"`
	Korisnickoime string `json:"korisnickoime"`
	Datum         string `json:"datum"`
}
