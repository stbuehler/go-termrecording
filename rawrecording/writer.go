package rawrecording

import (
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"io"
)

type countingWriter struct {
	writer io.Writer
	count  int64
}

func (cw *countingWriter) Write(p []byte) (n int, err error) {
	n, err = cw.writer.Write(p)
	cw.count = cw.count + int64(n)
	return
}

func newCountingWriter(writer io.Writer) *countingWriter {
	return &countingWriter{writer: writer, count: 0}
}

type FrameWriter interface {
	io.Closer
	Write(frame Frame) error
	Resize(columns uint32, rows uint32) error
}

// maximum size of a pack that includes more than one frame
// (a single frame is never split)
const MaximumPackSize int = 1 << 24 // 16 MByte

type Writer struct {
	err                 error
	target              io.Writer
	meta                Meta
	maxPackSize         int
	currentPackSize     int // 0 when currentPackFrames is empty
	currentPackFrames   []Frame
	currentTerminalSize *TerminalSize
}

func NewWriter(target io.Writer) FrameWriter {
	return NewWriterWithMaxPackSize(target, 256*1024)
}

// maxPackSize is always limited by MaximumPackSize
func NewWriterWithMaxPackSize(target io.Writer, maxPackSize int) *Writer {
	if maxPackSize > MaximumPackSize {
		maxPackSize = MaximumPackSize
	}
	return &Writer{
		target:          target,
		maxPackSize:     maxPackSize,
		currentPackSize: 0,
	}
}

func (w *Writer) storeError(err error) error {
	w.err = err
	return err
}

func (w *Writer) FlushPack() error {
	var err error

	if w.err != nil {
		return w.err
	}

	if 0 == len(w.currentPackFrames) {
		return nil
	}

	framePack := &FramePack{Frames: w.currentPackFrames}

	cWriter := newCountingWriter(w.target)
	gzipWriter := gzip.NewWriter(cWriter)
	var packBytes []byte
	if packBytes, err = framePack.Marshal(); err != nil {
		return w.storeError(fmt.Errorf("couldn't serialize pack: %s", err))
	}
	if _, err = gzipWriter.Write(packBytes); err != nil {
		return w.storeError(fmt.Errorf("couldn't write gzipped pack: %s", err))
	}
	if err = gzipWriter.Close(); err != nil {
		return w.storeError(fmt.Errorf("couldn't close gzipped pack: %s", err))
	}

	w.meta.PackIndex.Entries = append(w.meta.PackIndex.Entries, PackIndexEntry{
		Offset:   w.currentPackFrames[0].Offset,
		PackSize: uint32(cWriter.count),
	})

	w.currentPackFrames = nil
	w.currentTerminalSize = nil
	w.currentPackSize = 0

	return nil
}

func (w *Writer) Resize(columns uint32, rows uint32) error {
	if err := w.FlushPack(); err != nil {
		return err
	}

	w.currentTerminalSize = &TerminalSize{
		Columns: columns,
		Rows:    rows,
	}
	if columns > w.meta.MaxTerminalSize.Columns {
		w.meta.MaxTerminalSize.Columns = columns
	}
	if rows > w.meta.MaxTerminalSize.Rows {
		w.meta.MaxTerminalSize.Rows = rows
	}
	return nil
}

func (w *Writer) Write(frame Frame) error {
	if w.err != nil {
		return w.err
	}

	if 0 != len(w.currentPackFrames) && w.currentPackSize+len(frame.Data) > w.maxPackSize {
		if err := w.FlushPack(); err != nil {
			return err
		}
	}

	// try appending
	newFrames := append(w.currentPackFrames, frame)
	newPackSize := (&FramePack{Frames: newFrames}).Size()

	if 0 != len(w.currentPackFrames) && newPackSize > w.maxPackSize {
		// flush old frames; new frame will be the only one remaining
		if err := w.FlushPack(); err != nil {
			return err
		}

		w.currentPackFrames = []Frame{frame}
		w.currentPackSize = (&FramePack{Frames: w.currentPackFrames}).Size()
	} else {
		w.currentPackFrames = newFrames
		w.currentPackSize = newPackSize
	}

	if w.currentPackSize > w.maxPackSize {
		// new frame is too large to append more to it, flush it
		if err := w.FlushPack(); err != nil {
			return err
		}
	}

	return nil
}

// does NOT close the underlying io.Writer
func (w *Writer) Close() error {
	var err error
	if w.err != nil {
		return w.err
	}

	if err = w.FlushPack(); err != nil {
		return err
	}

	cWriter := newCountingWriter(w.target)
	gzipWriter := gzip.NewWriter(cWriter)
	var metaBytes []byte
	if metaBytes, err = (&w.meta).Marshal(); err != nil {
		return w.storeError(fmt.Errorf("couldn't serialize meta data: %s", err))
	}
	if _, err = gzipWriter.Write(metaBytes); err != nil {
		return w.storeError(fmt.Errorf("couldn't write gzipped meta data: %s", err))
	}
	if err = gzipWriter.Close(); err != nil {
		return w.storeError(fmt.Errorf("couldn't close gzipped meta data: %s", err))
	}
	if err = binary.Write(w.target, binary.LittleEndian, int64(cWriter.count)); err != nil {
		return w.storeError(fmt.Errorf("couldn't write meta data size: %s", err))
	}

	_ = w.storeError(fmt.Errorf("writer already closed"))

	return nil
}
