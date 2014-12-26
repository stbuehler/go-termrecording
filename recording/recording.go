package recording

import (
	"github.com/stbuehler/go-termrecording/rawrecording"
	"io"
)

func Execute(recordingFile io.Writer, command string, args ...string) error {
	term := NewPty()

	stdout := rawrecording.NewRecordingWriter(recordingFile)
	term.ResizeCallback = func(columns uint32, rows uint32) {
		stdout.Resize(columns, rows)
	}

	if err := term.Record(stdout, command, args...); err != nil {
		return err
	}

	if err := stdout.Close(); err != nil {
		return err
	}

	return nil
}
