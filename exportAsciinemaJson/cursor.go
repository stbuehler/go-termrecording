package exportAsciinemaJson

import (
	tsm "github.com/stbuehler/go-termrecording/libtsm"
)

type Cursor struct {
	X       uint `json:"x"`
	Y       uint `json:"y"`
	Visible bool `json:"visible"`
}

func GetCursor(screen tsm.Screen) Cursor {
	x, y := screen.GetCursor()
	flags := screen.GetFlags()
	return Cursor{
		X:       x,
		Y:       y,
		Visible: 0 != (flags & tsm.ScreenHideCursor),
	}
}

func (cursor Cursor) Diff(previous *Cursor) interface{} {
	if previous == nil {
		return cursor
	}
	m := map[string]interface{}{}
	if previous.X != cursor.X {
		m["x"] = cursor.X
	}
	if previous.Y != cursor.Y {
		m["y"] = cursor.Y
	}
	if previous.Visible != cursor.Visible {
		m["visible"] = cursor.Visible
	}
	if len(m) == 0 {
		return nil
	}
	return m
}
