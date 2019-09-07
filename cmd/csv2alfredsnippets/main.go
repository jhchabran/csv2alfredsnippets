package main

import (
	"archive/zip"
	"encoding/csv"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: csv2alfedsnippets input.csv output.alfredsnippets\n")
	}

	flag.Parse()
	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(1)
	}

	file, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

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

	err = createJsonEntries(flag.Arg(1), entries)
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
