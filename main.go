package main

import (
	"encoding/xml"
	"os"
	"text/template"
)

// Testsuites was generated 2023-09-19 12:11:27 by https://xml-to-go.github.io/ in Ukraine.
type Testsuites struct {
	XMLName    xml.Name `xml:"testsuites"`
	Text       string   `xml:",chardata"`
	Tests      string   `xml:"tests,attr"`
	Failures   string   `xml:"failures,attr"`
	TestSuites []struct {
		Text      string `xml:",chardata"`
		Name      string `xml:"name,attr"`
		Tests     string `xml:"tests,attr"`
		Failures  string `xml:"failures,attr"`
		Errors    string `xml:"errors,attr"`
		ID        string `xml:"id,attr"`
		Hostname  string `xml:"hostname,attr"`
		Time      string `xml:"time,attr"`
		Timestamp string `xml:"timestamp,attr"`
		TestCases []struct {
			Text      string `xml:",chardata"`
			Name      string `xml:"name,attr"`
			Classname string `xml:"classname,attr"`
			Time      string `xml:"time,attr"`
			Failure   struct {
				Text    string `xml:",chardata"`
				Message string `xml:"message,attr"`
			} `xml:"failure"`
			SystemOut string `xml:"system-out"`
		} `xml:"testcase"`
	} `xml:"testsuite"`
}

type files struct {
}

func main() {
	argsWithoutProg := os.Args[1:]
	var tmplFile = "template.tmpl"
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	processFiles(tmpl, argsWithoutProg)
}

func processFiles(tmpl *template.Template, paths []string) {
	for _, path := range paths {
		dat, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}
		var ts Testsuites
		err = xml.Unmarshal(dat, &ts)
		if err != nil {
			panic(err)
		}

		err = tmpl.Execute(os.Stdout, ts)
		if err != nil {
			panic(err)
		}
	}

}
