package exportAsciinemaJson

import (
	tsm "github.com/stbuehler/go-termrecording/libtsm"
)

func GetSnapshot(screen tsm.Screen) []Line {
	var lines []Line
	var currentLine *Line
	var currentLineNum uint
	var currentCell []rune
	var currentCellNulls int
	var currentAttr *tsm.ScreenAttr
	var currentBrush Brush

	// remember nulls, but not actually write them
	putNull := func() {
		currentCellNulls++
	}

	// write all remembered null characters (as spaces)
	syncNulls := func() {
		for i := 0; i < currentCellNulls; i++ {
			currentCell = append(currentCell, ' ')
		}
		currentCellNulls = 0
	}

	putCharacters := func(runes ...rune) {
		syncNulls()
		currentCell = append(currentCell, runes...)
	}

	closeCell := func() {
		if currentAttr != nil {
			syncNulls()
			if len(currentCell) > 0 {
				currentLine.Cells = append(currentLine.Cells, Cell{
					Text:  string(currentCell),
					Brush: currentBrush,
				})
				currentCell = []rune{}
			}
			currentAttr = nil
		}
	}

	openCell := func(attr *tsm.ScreenAttr) {
		if currentAttr == nil || *currentAttr != *attr {
			brush := MakeBrush(attr)
			if currentAttr == nil || brush != currentBrush {
				closeCell()
				currentAttr = attr
				currentBrush = brush
			}
		}
	}

	closeLine := func() {
		// drop all trailing nulls
		currentCellNulls = 0
		closeCell()
		if currentLine != nil {
			lines = append(lines, *currentLine)
			currentLine = nil
		}
	}

	useLine := func(y uint) {
		if currentLine == nil || currentLineNum != y {
			closeLine()
			currentLine = &Line{
				Cells: []Cell{},
			}
			currentLineNum = y
		}
	}

	drawCb := func(id uint32, character string, width uint, x uint, y uint, attr *tsm.ScreenAttr, age tsm.Age) bool {
		useLine(y)
		openCell(attr)
		if 32 == id {
			// tsm reports blanks as empty character
			putCharacters(' ')
		} else if len(character) == 0 {
			putNull()
		} else {
			// asciinema only takes the first rune
			putCharacters([]rune(character)[0])
			//putCharacters([]rune(character)...)
		}
		return true
	}

	screen.Draw(drawCb)
	closeLine()

	return lines
}
