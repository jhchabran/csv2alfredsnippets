package main

import (
	"archive/zip"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "usage: csv2alfedsnippets input.csv output.alfredsnippets\n")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var entries []*alfredEntry
	for _, record := range records {
		entry := csvToAlfred(record)
		entries = append(entries, entry)
	}

	err = createJsonEntries(os.Args[2], entries)
	if err != nil {
		log.Fatal(err)
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
		Snippet: entry[2],
		Uid:     uid.String(),
		Name:    entry[0],
		Keyword: entry[1],
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
	defer file.Close()

	w := zip.NewWriter(file)
	defer func() {
		e := w.Close()
		if err == nil {
			err = e
		}
	}()

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

	return nil
}
