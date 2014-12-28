package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"github.com/stbuehler/go-termrecording/exportAsciinemaJson"
	"github.com/stbuehler/go-termrecording/recording"
	"io"
	"io/ioutil"
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
	opts := flag.NewFlagSet("recording options", flag.PanicOnError)
	outputBaseName := opts.String("out", "recording", "recording output name (used as basename for .html and -stdout.json files); defaults to 'recording'")

	err := opts.Parse(os.Args[1:])
	if err != nil {
		return
	}

	command := opts.Args()

	if len(command) == 0 {
		shell := os.Getenv("SHELL")
		if shell == "" {
			shell = "/bin/sh"
		}
		command = []string{shell}
	}

	recordingFile, err := ioutil.TempFile("", "term-recording")
	if err != nil {
		panic("couldn't create temporary recording file")
	}
	defer recordingFile.Close()
	defer os.Remove(recordingFile.Name())

	stdoutJsonName := *outputBaseName + "-stdout.json"
	htmlName := *outputBaseName + ".html"

	stdoutJsonFile, err := os.OpenFile(
		stdoutJsonName,
		os.O_RDWR|os.O_CREATE|os.O_EXCL,
		0644)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	htmlFile, err := os.OpenFile(
		htmlName,
		os.O_RDWR|os.O_CREATE|os.O_EXCL,
		0644)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	println("Recording into " + htmlName)

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

	err = exportAsciinemaJson.MakeFilm(stdoutJsonName, stdoutJsonFile, htmlFile, fileSecReader)
	if err != nil {
		println("couldn't write recording as JSON/HTML: " + err.Error())
		return
	}

	println("Recording finished successfully")
}
