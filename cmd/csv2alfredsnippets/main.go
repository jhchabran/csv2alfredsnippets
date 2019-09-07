package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"
)

func main() {
	file, err := os.Open("/Users/tech/Downloads/snippets.csv")
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	entry := csvToAlfred(records[0])
	raw, err := dumpEntry(entry)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(raw))
}

type alfredEntry struct {
	Snippet string `json:"snippet"`
	Uid     string `json:"uid"`
	Name    string `json:"name"`
	Keyword string `json:"keyword"`
}

func csvToAlfred(entry []string) *alfredEntry {
	uid := uuid.New()

	return &alfredEntry{
		Snippet: entry[0],
		Uid:     uid.String(),
		Name:    entry[1],
		Keyword: entry[2],
	}
}

func dumpEntry(entry *alfredEntry) ([]byte, error) {
	if entry == nil {
		return nil, errors.New("nil alfred entry")
	}

	raw, err := json.Marshal(entry)
	if err != nil {
		return nil, err
	}

	return raw, nil
}
