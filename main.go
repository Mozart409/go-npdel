package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/urfave/cli/v2"
)

func main() {

	folders := findNodeModules()

	if len(folders) == 0 {
		log.Fatal("No node_modules found")
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print only the version",
	}

	app := &cli.App{
		Name:     "go-npdel",
		Version:  "v0.0.1",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "mozart409",
				Email: "amadeus@mozart409.com",
			},
		},
		Copyright: "MIT License",
		Usage:     "delete node_modules folder",
		Action: func(*cli.Context) error {
			//fmt.Println("cli ", folders)
			result := selectUi(folders)

			start := time.Now()

			for _, folder := range result {
				fmt.Println("Deleting ", folder)
				os.RemoveAll(folder)
			}

			timeElapsed := time.Since(start).Seconds()
			fmt.Printf("Deletion took %f", timeElapsed)

			return nil
		},
	}

	app.Suggest = true

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func findNodeModules() []string {

	folders, err := filepath.Glob("**/node_modules")
	//folders, err := filepathx.Glob("**/node_modules")

	if err != nil {
		log.Fatal(err)
	}

	return folders
}

func selectUi(options []string) []string {
	result := []string{}

	prompt := &survey.MultiSelect{
		Message:  "Select folders:",
		Options:  options,
		PageSize: 10,
	}
	survey.AskOne(prompt, &result)

	return result
}
