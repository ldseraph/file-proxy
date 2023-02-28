package main

import (
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

var routeFlag string
var dirFlag string
var portFlag string

func init() {
	rand.Seed(time.Now().Unix())
	root.PersistentFlags().StringVarP(&routeFlag, "route", "r", "/static", "proxy route")
	root.PersistentFlags().StringVarP(&dirFlag, "dir", "d", "./static", "proxy dir")
	root.PersistentFlags().StringVarP(&portFlag, "port", "p", ":3000", "proxy port")
}

var root = &cobra.Command{
	Short: "file-proxy",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	app.Static(routeFlag, dirFlag)
	app.Listen(portFlag)
}

func main() {
	root.Execute()
}
