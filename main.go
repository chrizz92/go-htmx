package main

import (
	"github.com/a-h/templ"
	"github.com/chrizz92/go-htmx/templates"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/phayes/freeport"
	webview "github.com/webview/webview_go"
	"log"
	"strconv"
)

func getPort() string {
	port, err := freeport.GetFreePort()
	if err != nil {
		log.Fatal(err)
	}
	return ":" + strconv.Itoa(port)
}

func runServer(port string) {
	app := fiber.New()
	app.Static("/assets", "./assets", fiber.Static{Browse: true})
	app.Get("/", adaptor.HTTPHandler(templ.Handler(templates.Root("World"))))
	app.Post("/", adaptor.HTTPHandler(templ.Handler(templates.Clicked())))
	log.Fatal(app.Listen(port))
}

func main() {
	port := getPort()
	go runServer(port)
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Go-HTMX-Webview-Example")
	w.SetSize(480, 320, webview.HintNone)
	w.Navigate("http://localhost" + port + "/")
	w.Run()
}
