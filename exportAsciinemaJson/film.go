package exportAsciinemaJson

import (
	"encoding/json"
	tsm "github.com/stbuehler/go-termrecording/libtsm"
	"github.com/stbuehler/go-termrecording/rawrecording"
	"io"
)

type FilmFrame struct {
	Delay float64
	Diff  interface{}
}

type FilmMaker struct {
	FirstFrameOffset float64
	LastFrameOffset  float64
	PreviousFrame    *Frame
	Frames           []FilmFrame
}

func (filmFrame FilmFrame) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{filmFrame.Delay, filmFrame.Diff})
}

func (film FilmMaker) MarshalJSON() ([]byte, error) {
	return json.Marshal(film.Frames)
}

func (film *FilmMaker) AddFrame(timeOffset float64, screen tsm.Screen) {
	frame := MakeFrame(screen)
	frameDiff := frame.Diff(film.PreviousFrame)

	if frameDiff != nil {
		delay := timeOffset - film.LastFrameOffset
		if len(film.Frames) == 0 {
			delay = 0 // first frame always starts at offset 0
			film.FirstFrameOffset = timeOffset
		}

		film.Frames = append(film.Frames, FilmFrame{
			Delay: delay,
			Diff:  frameDiff,
		})
		film.PreviousFrame = &frame
		film.LastFrameOffset = timeOffset
	}
}

func MakeFilm(jsonFileName string, jsonFile io.Writer, htmlFile io.Writer, fileReader *io.SectionReader) error {
	reader, err := rawrecording.NewReaderFromSectionReader(fileReader)
	if err != nil {
		return err
	}

	film := FilmMaker{}

	screen, err := tsm.NewScreen()
	if err != nil {
		return err
	}

	if err := screen.Resize(uint(reader.Meta.MaxTerminalSize.Columns), uint(reader.Meta.MaxTerminalSize.Rows)); err != nil {
		return err
	}

	vte, err := tsm.NewVte(screen, nil)
	if err != nil {
		return err
	}

	// println(Stringify(reader.Meta))

	for {
		frame, _ := reader.ReadFrame()
		if nil == frame {
			break
		}
		vte.InputBytes(frame.Data)
		// println(Stringify(frame))
		film.AddFrame(frame.Offset, screen)
	}

	totalTime := film.LastFrameOffset - film.FirstFrameOffset

	filmJson, err := json.Marshal(film)
	if err != nil {
		return err
	}

	if _, err := jsonFile.Write(filmJson); err != nil {
		return err
	}

	snapshot := MakeFrame(screen)

	err = WriteHTML(htmlFile, reader.Meta.MaxTerminalSize, &snapshot, jsonFileName, totalTime)
	return err
}
