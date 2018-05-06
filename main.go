package main

import (
	"io/ioutil"
	"log"
	"os"
	"text/template"

	// External --------------------
	"gopkg.in/yaml.v2"
)

const (
	EXIT_OK = iota
	EXIT_ERROR
	EXIT_HELP
)

type Story struct {
	Title              string   `yaml:"Title"`
	IdealSchedule      []string `yaml:"Ideal Schedule"`
	Goal               string   `yaml:"Goal"`
	Who                string   `yaml:"Who"`
	What               string   `yaml:"What"`
	Why                string   `yaml:"Why"`
	AcceptanceCriteria []string `yaml:"Acceptance Criteria"`
	Issue              *string  `yaml:"Issue"`
	Labels             []string `yaml:"Labels"`
}

/**
 * Simple utility for error check
 */
func panicIfError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("This program need one file name")
		os.Exit(EXIT_ERROR)
	}

	stat, err := os.Stat(os.Args[1])
	if err != nil {
		log.Fatalf("Faced an error on checking file status")
		log.Fatalf("Err: %v", err)
		os.Exit(EXIT_ERROR)
	}

	if !stat.Mode().IsRegular() {
		log.Fatalf("Specified file seems not a regular file")
		os.Exit(EXIT_ERROR)
	}

	raw, err := ioutil.ReadFile(stat.Name())

	var stories []Story
	err = yaml.Unmarshal(raw, &stories)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
		os.Exit(EXIT_ERROR)
	}

	tmpl, err := template.ParseFiles("md.tmpl")
	panicIfError(err)
	for _, s := range stories {
		err = tmpl.Execute(os.Stdout, s)
		panicIfError(err)
		return
	}

}
