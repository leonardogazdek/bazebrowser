package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
)

var (
	AppName            string
	BuiltAt            string
	VersionAstilectron string
	VersionElectron    string
)

var (
	fs           = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	debug        = fs.Bool("d", true, "enables the debug mode")
	w            *astilectron.Window
	wAddr        *astilectron.Window
	selectedUser = 1
	curUrl       = ""
)

func main() {
	l := log.New(log.Writer(), log.Prefix(), log.Flags())

	fs.Parse(os.Args[1:])

	l.Printf("Running app built at %s\n", BuiltAt)
	if err := bootstrap.Run(bootstrap.Options{

		Asset:    Asset,
		AssetDir: AssetDir,
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			AppIconDarwinPath:  "resources/icon.icns",
			AppIconDefaultPath: "resources/icon.png",
			SingleInstance:     true,
			VersionAstilectron: VersionAstilectron,
			VersionElectron:    VersionElectron,
		},
		Debug:  *debug,
		Logger: l,
		MenuOptions: []*astilectron.MenuItemOptions{{
			Label: astikit.StrPtr("File"),
			SubMenu: []*astilectron.MenuItemOptions{
				{Role: astilectron.MenuItemRoleClose},
			},
		}},
		OnWait: func(_ *astilectron.Astilectron, ws []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			w = ws[0]
			wAddr = ws[1]
			w.On(astilectron.EventNameWindowEventWillNavigate, func(e astilectron.Event) (deleteListener bool) {
				curUrl = e.URL
				wAddr.ExecuteJavaScript("document.getElementById('url').value='" + e.URL + "'")
				db, err := sql.Open("sqlite3", "./resources/app/main.db")
				if err != nil {
					log.Fatal(err)
				}
				defer db.Close()

				tx, err := db.Begin()
				if err != nil {
					log.Fatal(err)
				}
				stmt, err := tx.Prepare("insert into povijest (url, korisnici_id) values(?, ?)")
				if err != nil {
					log.Fatal(err)
				}
				defer stmt.Close()

				_, err = stmt.Exec(e.URL, selectedUser)
				if err != nil {
					log.Fatal(err)
				}
				tx.Commit()
				return
			})
			return nil
		},
		RestoreAssets: RestoreAssets,
		Windows: []*bootstrap.Window{
			{
				Homepage:       "index.html",
				MessageHandler: handleMessages,
				Options: &astilectron.WindowOptions{
					BackgroundColor: astikit.StrPtr("#333"),
					Center:          astikit.BoolPtr(true),
					Height:          astikit.IntPtr(700),
					Width:           astikit.IntPtr(700),
				},
			},
			{
				Homepage:       "address.html",
				MessageHandler: handleMessages,
				Options: &astilectron.WindowOptions{
					BackgroundColor: astikit.StrPtr("#333"),
					X:               astikit.IntPtr(100),
					Y:               astikit.IntPtr(0),
					Height:          astikit.IntPtr(200),
					Width:           astikit.IntPtr(300),
					AlwaysOnTop:     astikit.BoolPtr(true),
					Closable:        astikit.BoolPtr(false),
					Frame:           astikit.BoolPtr(true),
					Transparent:     astikit.BoolPtr(true),
				},
			},
		},
	}); err != nil {
		l.Fatal(fmt.Errorf("running bootstrap failed: %w", err))
	}

}
