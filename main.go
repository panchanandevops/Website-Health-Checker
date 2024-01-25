// main.go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	// Define the command-line application
	app := &cli.App{
		Name:  "Healthchecker",
		Usage: "A tiny tool that checks the given domain is down.",
		Flags: []cli.Flag{
			// Define the domain flag
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "Domain name to check.",
				Required: true,
			},
			// Define the port flag
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "Port number to check.",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			// Extract values from flags or use defaults
			port := c.String("port")
			if c.String("port") == "" {
				port = "80"
			}
			// Call the Check function and print the result
			status := Check(c.String("domain"), port)
			fmt.Println(status)
			return nil
		},
	}

	// Run the CLI application
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
