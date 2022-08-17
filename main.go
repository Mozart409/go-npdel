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

	if folders == nil {
		log.Fatal("No node_modules found")
	}
	fmt.Println(folders)

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

	/* for _, folder := range folders {

		fmt.Println(folder)
	} */

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

	/* prompt2 := promptui.Select{
		Label: "Select Folder",
		Items: options,
		Size:  8,
	}

	_, result, err := prompt2.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)

	}

	fmt.Printf("You choose %q\n", result)

	return result, err */

	return result
}
