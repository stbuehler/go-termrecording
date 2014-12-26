package rawrecording

import (
	"io"
	"time"
)

type Recording struct {
	writer    FrameWriter
	startTime time.Time
}

func NewRecording(writer FrameWriter) *Recording {
	return &Recording{
		writer:    writer,
		startTime: time.Now(),
	}
}

func NewRecordingWriter(target io.Writer) *Recording {
	return &Recording{
		writer:    NewWriter(target),
		startTime: time.Now(),
	}
}

func copyBytes(src []byte) []byte {
	dst := make([]byte, len(src))
	copy(dst, src)
	return dst
}

func (s *Recording) Resize(columns uint32, rows uint32) error {
	return s.writer.Resize(columns, rows)
}

func (s *Recording) Write(p []byte) (int, error) {
	frame := Frame{
		Offset: time.Now().Sub(s.startTime).Seconds(),
		Data:   copyBytes(p),
	}
	if err := s.writer.Write(frame); err != nil {
		return 0, err
	}
	return len(p), nil
}

// Closes the underlying FrameWriter
func (s *Recording) Close() error {
	return s.writer.Close()
}
