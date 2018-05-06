package main

import (
	"bytes"
	"fmt"
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
	template           *template.Template
}

func (s *Story) toMarkdown() []byte {
	var err error
	if s.template == nil {
		s.template, err = template.New("md.tmpl").ParseFiles("md.tmpl")
		panicIfError(err)
	}

	var b bytes.Buffer
	err = s.template.Execute(&b, s)
	panicIfError(err)

	return b.Bytes()
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

	for _, s := range stories {
		fmt.Println(string(s.toMarkdown()))
	}

}
