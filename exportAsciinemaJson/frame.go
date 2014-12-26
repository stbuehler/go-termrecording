package exportAsciinemaJson

import (
	tsm "github.com/stbuehler/go-termrecording/libtsm"
	"strconv"
)

type Frame struct {
	Lines  []Line `json:"lines"`
	Cursor Cursor `json:"cursor"`
}

func MakeFrame(screen tsm.Screen) Frame {
	return Frame{
		Lines:  GetSnapshot(screen),
		Cursor: GetCursor(screen),
	}
}

func (frame Frame) Diff(previousFrame *Frame) interface{} {
	if previousFrame == nil {
		return frame
	}
	m := map[string]interface{}{}
	lines := map[string]interface{}{}
	for y, line := range frame.Lines {
		if !line.Equal(previousFrame.Lines[y]) {
			lines[strconv.Itoa(y)] = line
		}
	}
	if len(lines) > 0 {
		m["lines"] = lines
	}
	if diff := frame.Cursor.Diff(&previousFrame.Cursor); diff != nil {
		m["cursor"] = diff
	}

	if len(m) == 0 {
		return nil
	}
	return m
}
