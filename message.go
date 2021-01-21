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
	case "insertBookmark":
		{
			var bookmark KnjiznaOznakaUpdate
			if len(m.Payload) > 0 {
				// Unmarshal payload
				if err = json.Unmarshal(m.Payload, &bookmark); err != nil {
					payload = err.Error()
					return
				}
			}

			fmt.Println("bookmark set ", bookmark.Ime, bookmark.Url, bookmark.Kategorije_id, bookmark.Korisnici_id)

			db, err := sql.Open("sqlite3", "./resources/app/main.db")
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()

			tx, err := db.Begin()
			if err != nil {
				log.Fatal(err)
			}
			stmt, err := tx.Prepare("insert into knjizneoznake (ime, url, korisnici_id, kategorije_id) values(?, ?, ?, ?)")
			if err != nil {
				log.Fatal(err)
			}
			defer stmt.Close()

			_, err = stmt.Exec(bookmark.Ime, bookmark.Url, bookmark.Korisnici_id, bookmark.Kategorije_id)
			if err != nil {
				log.Fatal(err)
			}
			tx.Commit()

			payload = bookmark
			break
		}
	case "updateBookmark":
		{
			var bookmark KnjiznaOznakaUpdate
			if len(m.Payload) > 0 {
				// Unmarshal payload
				if err = json.Unmarshal(m.Payload, &bookmark); err != nil {
					payload = err.Error()
					return
				}
			}

			fmt.Println("bookmark set ", bookmark.Ime, bookmark.Url, bookmark.Kategorije_id, bookmark.Korisnici_id)

			db, err := sql.Open("sqlite3", "./resources/app/main.db")
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()

			tx, err := db.Begin()
			if err != nil {
				log.Fatal(err)
			}
			stmt, err := tx.Prepare("UPDATE knjizneoznake SET ime = ?, url = ?, kategorije_id = ? WHERE id = ?")
			if err != nil {
				log.Fatal(err)
			}
			defer stmt.Close()

			_, err = stmt.Exec(bookmark.Ime, bookmark.Url, bookmark.Kategorije_id, bookmark.Id)
			if err != nil {
				log.Fatal(err)
			}
			tx.Commit()

			payload = bookmark
			break
		}
	case "insertExtensions":
		{
			var ext ProsirenjeData
			if len(m.Payload) > 0 {
				// Unmarshal payload
				if err = json.Unmarshal(m.Payload, &ext); err != nil {
					payload = err.Error()
					return
				}
			}

			db, err := sql.Open("sqlite3", "./resources/app/main.db")
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()

			tx, err := db.Begin()
			if err != nil {
				log.Fatal(err)
			}
			stmt, err := tx.Prepare("insert into prosirenja (ime, opis) values(?, ?)")
			if err != nil {
				log.Fatal(err)
			}
			defer stmt.Close()

			_, err = stmt.Exec(ext.Ime, ext.Opis)
			if err != nil {
				log.Fatal(err)
			}
			tx.Commit()

			payload = ext
			break
		}
	case "updateExtensions":
		{
			var ext ProsirenjeData
			if len(m.Payload) > 0 {
				// Unmarshal payload
				if err = json.Unmarshal(m.Payload, &ext); err != nil {
					payload = err.Error()
					return
				}
			}

			db, err := sql.Open("sqlite3", "./resources/app/main.db")
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()

			tx, err := db.Begin()
			if err != nil {
				log.Fatal(err)
			}
			stmt, err := tx.Prepare("UPDATE prosirenja SET ime = ?, opis = ? WHERE id = ?")
			if err != nil {
				log.Fatal(err)
			}
			defer stmt.Close()

			_, err = stmt.Exec(ext.Ime, ext.Opis, ext.Id)
			if err != nil {
				log.Fatal(err)
			}
			tx.Commit()

			payload = ext
			break
		}

	case "insertTab":
		{
			var tab OtvoreneKarticeData
			if len(m.Payload) > 0 {
				// Unmarshal payload
				if err = json.Unmarshal(m.Payload, &tab); err != nil {
					payload = err.Error()
					return
				}
			}

			db, err := sql.Open("sqlite3", "./resources/app/main.db")
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()

			tx, err := db.Begin()
			if err != nil {
				log.Fatal(err)
			}
			stmt, err := tx.Prepare("insert into otvorene_kartice (url, korisnici_id) values(?, ?)")
			if err != nil {
				log.Fatal(err)
			}
			defer stmt.Close()

			_, err = stmt.Exec(tab.Url, tab.Korisnici_id)
			if err != nil {
				log.Fatal(err)
			}
			tx.Commit()

			payload = tab
			break
		}
	case "updateTab":
		{
			var tab OtvoreneKarticeData
			if len(m.Payload) > 0 {
				// Unmarshal payload
				if err = json.Unmarshal(m.Payload, &tab); err != nil {
					payload = err.Error()
					return
				}
			}

			db, err := sql.Open("sqlite3", "./resources/app/main.db")
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()

			tx, err := db.Begin()
			if err != nil {
				log.Fatal(err)
			}
			stmt, err := tx.Prepare("UPDATE otvorene_kartice SET url = ? WHERE id = ?")
			if err != nil {
				log.Fatal(err)
			}
			defer stmt.Close()

			_, err = stmt.Exec(tab.Url, tab.Id)
			if err != nil {
				log.Fatal(err)
			}
			tx.Commit()

			payload = tab
			break
		}

	case "insertSettings":
		{
			var set PostavkeData
			if len(m.Payload) > 0 {
				// Unmarshal payload
				if err = json.Unmarshal(m.Payload, &set); err != nil {
					payload = err.Error()
					return
				}
			}

			db, err := sql.Open("sqlite3", "./resources/app/main.db")
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()

			tx, err := db.Begin()
			if err != nil {
				log.Fatal(err)
			}
			stmt, err := tx.Prepare("insert into postavke (ime, vrijednost) values(?, ?)")
			if err != nil {
				log.Fatal(err)
			}
			defer stmt.Close()

			_, err = stmt.Exec(set.Ime, set.Vrijednost)
			if err != nil {
				log.Fatal(err)
			}
			tx.Commit()

			payload = set
			break
		}
	case "updateSettings":
		{
			var set PostavkeData
			if len(m.Payload) > 0 {
				// Unmarshal payload
				if err = json.Unmarshal(m.Payload, &set); err != nil {
					payload = err.Error()
					return
				}
			}

			db, err := sql.Open("sqlite3", "./resources/app/main.db")
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()

			tx, err := db.Begin()
			if err != nil {
				log.Fatal(err)
			}
			stmt, err := tx.Prepare("UPDATE postavke SET ime = ?, vrijednost = ? WHERE id = ?")
			if err != nil {
				log.Fatal(err)
			}
			defer stmt.Close()

			_, err = stmt.Exec(set.Ime, set.Vrijednost, set.Id)
			if err != nil {
				log.Fatal(err)
			}
			tx.Commit()

			payload = set
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

			fmt.Println("getting povijest")
			rows, err := db.Query("select id, url, vremenskistambilj from povijest where korisnici_id = " + action + "")
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("getting knjizneoznake")
			rowsBook, err := db.Query("select p.id,p.ime,p.url,k.ime,p.kategorije_id from knjizneoznake p left join kategorije k on p.kategorije_id = k.id where korisnici_id = " + action + "")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("getting prosirenja")
			rowsExt, err := db.Query("select p.id, p.ime, p.opis from prosirenja p join korisnici_prosirenja_veza v on p.id = v.prosirenja_id join korisnici k on k.id = v.korisnici_id WHERE v.korisnici_id = " + action + "")
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("getting otvorene_kartice")
			rowsTabs, err := db.Query("select id, url, korisnici_id FROM otvorene_kartice WHERE korisnici_id = " + action + "")
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("getting postavke")
			rowsSettings, err := db.Query("select id, ime, vrijednost FROM postavke")
			if err != nil {
				log.Fatal(err)
			}

			defer rows.Close()
			defer rowsBook.Close()
			defer rowsExt.Close()
			defer rowsTabs.Close()
			defer rowsSettings.Close()
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
				var kategorije_id int
				err = rowsBook.Scan(&id, &ime, &url, &kategorija, &kategorije_id)
				if err != nil {
					log.Fatal(err)
				}
				knjizneoznakeData = append(knjizneoznakeData, KnjiznaOznakaData{
					Id:            id,
					Ime:           ime,
					Url:           url,
					Kategorija:    kategorija,
					Kategorije_id: kategorije_id,
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

			otvoreneKarticeData := []OtvoreneKarticeData{}
			for rowsTabs.Next() {
				var id int
				var url string
				var korisnici_id int
				err = rowsTabs.Scan(&id, &url, &korisnici_id)
				if err != nil {
					log.Fatal(err)
				}
				otvoreneKarticeData = append(otvoreneKarticeData, OtvoreneKarticeData{
					Id:  id,
					Url: url,
					Korisnici_id: korisnici_id,
				})

			}

			postavkeData := []PostavkeData{}
			for rowsSettings.Next() {
				var id int
				var ime string
				var vrijednost string
				err = rowsSettings.Scan(&id, &ime, &vrijednost)
				if err != nil {
					log.Fatal(err)
				}
				postavkeData = append(postavkeData, PostavkeData{
					Id:  id,
					Ime: ime,
					Vrijednost: vrijednost,
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

			err = rowsTabs.Err()
			if err != nil {
				log.Fatal(err)
			}

			err = rowsSettings.Err()
			if err != nil {
				log.Fatal(err)
			}

			arr := KorisnikData{
				Povijest:        povijestData,
				KnjizneOznake:   knjizneoznakeData,
				Prosirenja:      prosirenjaData,
				OtvoreneKartice: otvoreneKarticeData,
				Postavke: 		 postavkeData,
			}
			payload = arr

			break
		}

	}
	return
}

type KorisnikData struct {
	Povijest        []PovijestData        `json:"povijestdata"`
	KnjizneOznake   []KnjiznaOznakaData   `json:"knjizneoznakedata"`
	Prosirenja      []ProsirenjeData      `json:"prosirenjadata"`
	OtvoreneKartice []OtvoreneKarticeData `json:"otvorenekarticedata"`
	Postavke		[]PostavkeData		  `json:"postavkedata"`
}

type PovijestData struct {
	Id                int    `json:"id"`
	Url               string `json:"url"`
	Vremenskistambilj string `json:"vremenskistambilj"`
}

type KnjiznaOznakaData struct {
	Id            int    `json:"id"`
	Ime           string `json:"ime"`
	Url           string `json:"url"`
	Kategorija    string `json:"kategorija"`
	Kategorije_id int    `json:"kategorije_id"`
}

type ProsirenjeData struct {
	Id   int    `json:"id"`
	Ime  string `json:"ime"`
	Opis string `json:"opis"`
}

type OtvoreneKarticeData struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
	Korisnici_id int `json:"korisnici_id"`
}

type Korisnik struct {
	Id            int    `json:"id"`
	Korisnickoime string `json:"korisnickoime"`
	Datum         string `json:"datum"`
}

type KnjiznaOznakaUpdate struct {
	Id            int    `json: "id"`
	Ime           string `json:"ime"`
	Url           string `json:"url"`
	Kategorije_id int    `json:"kategorije_id"`
	Korisnici_id  int    `json:"korisnici_id"`
}

type PostavkeData struct {
	Id int `json:"id"`
	Ime string `json:"ime"`
	Vrijednost string `json:"vrijednost"`
}