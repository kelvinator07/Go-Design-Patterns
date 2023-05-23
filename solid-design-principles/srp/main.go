package main

// Single Responsibility Principle
// Separation of Concerns

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

var entryCount = 0
type Journal struct {
	entries []string
}

func (j *Journal) toString() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s ", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) removeEntry(index int) {
	//...
}

// Separation of concerns

func (j *Journal) Save(filename string) {
	_ =ioutil.WriteFile(filename, []byte(j.toString()), 0644)
}

func (j *Journal) Load(filename string) {
	//...
}

func (j *Journal) LoadFromWeb(url *url.URL) {
	//...
}

func main() {
	j := Journal{}
	j.AddEntry("I solved math today")
	j.AddEntry("I ran 3km today")
	fmt.Print(j.toString())
}
