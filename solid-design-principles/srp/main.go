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
	entryCount--
	j.entries = append(j.entries[:index], j.entries[index+1:]...)
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

var LineSeperator  = "\n"
func SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, LineSeperator)), 0644)
}

type Persistence struct {
	lineSeperator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeperator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("I solved math today")
	j.AddEntry("I ran 3km today")
	fmt.Print(j.toString())

	fmt.Println("\n==")

	j.removeEntry(1)
	fmt.Print(j.toString())

	SaveToFile(&j, "journal.txt")

	p := Persistence{"\r\n"}
	p.SaveToFile(&j, "journal.txt")

}
