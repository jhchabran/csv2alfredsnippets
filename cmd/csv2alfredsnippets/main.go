package main

import (
	"archive/zip"
	"encoding/csv"
	"encoding/json"
	"errors"
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

	var entries []*alfredEntry
	for _, record := range records {
		entry := csvToAlfred(record)
		entries = append(entries, entry)
	}

	err = createJsonEntries("test.alfredsnippets", entries)
	if err != nil {
		return
	}
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

	raw, err := json.Marshal(struct {
		AlfredSnippet *alfredEntry `json:"alfredsnippet"`
	}{entry})
	if err != nil {
		return nil, err
	}

	return raw, nil
}

func createJsonEntries(filepath string, entries []*alfredEntry) (err error) {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}

	w := zip.NewWriter(file)

	for _, entry := range entries {
		name := entry.Name + " " + entry.Uid + ".json"
		f, err := w.Create(name)
		if err != nil {
			return err
		}

		b, err := dumpEntry(entry)
		if err != nil {
			return err
		}

		_, err = f.Write(b)
		if err != nil {
			return err
		}
	}

	return w.Close()
}
