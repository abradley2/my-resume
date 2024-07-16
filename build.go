package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"text/template"

	"github.com/BurntSushi/toml"
)

type ResumeData struct {
	Name                    string
	Email                   string
	Phone                   string
	Address                 string
	LinkedIn                string
	Github                  string
	EmploymentHistory       []EmploymentHistory
	OpenSourceContributions []OpenSourceContribution
	Technologies            []Technology
}

type EmploymentHistory struct {
	Title       string
	Dates       string
	JobTitle    string
	Description string
}

type OpenSourceContribution struct {
	Link        string
	Language    string
	Description string
}

type Technology struct {
	Ranking string
	Items   string
}

func build() error {
	var err error

	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("getting working directory: %s", err)
	}

	tplcontent, err := os.ReadFile(
		path.Join(wd, "index.template.html"),
	)

	if err != nil {
		return fmt.Errorf("reading template file: %s", err)
	}

	datacontent, err := os.ReadFile(
		path.Join(wd, "data.toml"),
	)

	if err != nil {
		return fmt.Errorf("reading data file: %s", err)
	}

	tpl := template.New("resume")
	tpl, err = tpl.Parse(string(tplcontent))

	if err != nil {
		return fmt.Errorf("parsing template: %s", err)
	}

	destf, err := os.OpenFile(path.Join(wd, "index.html"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		return fmt.Errorf("opening destination file: %s", err)
	}

	defer destf.Close()

	data := ResumeData{}

	err = toml.Unmarshal(datacontent, &data)

	if err != nil {
		return fmt.Errorf("unmarshalling data from datacontent: %s", err)
	}

	err = tpl.Execute(destf, data)

	if err != nil {
		return fmt.Errorf("executing template: %s", err)
	}

	return err
}

func main() {
	if err := build(); err != nil {
		log.Fatalf("Failed with error: %v", err)
	}
}
