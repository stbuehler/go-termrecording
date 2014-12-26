package main

import (
	"bytes"
	"encoding/json"
	"github.com/stbuehler/go-termrecording/exportAsciinemaJson"
	"github.com/stbuehler/go-termrecording/recording"
	"io"
	"os"
)

func bytesSectionReader(b []byte) *io.SectionReader {
	return io.NewSectionReader(bytes.NewReader(b), 0, int64(len(b)))
}

func fileSectionReader(f *os.File) (*io.SectionReader, error) {
	if stat, err := f.Stat(); err != nil {
		return nil, err
	} else {
		return io.NewSectionReader(f, 0, stat.Size()), nil
	}
}

func stringify(v interface{}) string {
	str, err := json.Marshal(v)
	if err != nil {
		return err.Error()
	}
	return string(str)
}

func main() {
	recordingFile, err := os.OpenFile(
		"term.rec",
		os.O_RDWR|os.O_CREATE|os.O_TRUNC,
		0600)
	if err != nil {
		panic("couldn't open term.rec: " + err.Error())
	}

	command := os.Args[1:]
	if len(command) == 0 {
		command = []string{"/bin/sh"}
	}

	err = recording.Execute(recordingFile, command[0], command[1:]...)
	if err != nil {
		panic("record failed: " + err.Error())
	}
	err = recordingFile.Sync()
	if err != nil {
		panic("writing recording failed: " + err.Error())
	}

	// println("parse recording")

	fileSecReader, err := fileSectionReader(recordingFile)
	if err != nil {
		panic("couldn't read recording: " + err.Error())
	}

	err = exportAsciinemaJson.MakeFilm("term-stdout.json", "term.html", fileSecReader)
	if err != nil {
		panic("couldn't write recording as JSON/HTML: " + err.Error())
	}
}
