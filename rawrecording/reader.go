package rawrecording

import (
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
)

type Reader struct {
	currentPack     *FramePack
	currentFrameNdx int // index in currentPack
	currentPackNdx  int

	packs            *io.SectionReader
	indexTimeOffsets []float64
	indexFileOffsets []int64

	Meta Meta
}

func NewReaderFromSectionReader(source *io.SectionReader) (*Reader, error) {
	var err error

	sourceSize := source.Size()
	if sourceSize < 8 {
		return nil, fmt.Errorf("file too small (size: %d), couldn't find meta data", sourceSize)
	}

	source.Seek(sourceSize-8, 0)

	var metaSize int64
	if err := binary.Read(source, binary.LittleEndian, &metaSize); err != nil {
		return nil, fmt.Errorf("couldn't read meta data size: %s", err)
	}
	if metaSize < 0 {
		return nil, fmt.Errorf("meta data size negative: %d", metaSize)
	}

	if sourceSize-8 < metaSize {
		return nil, fmt.Errorf("file too small (size: %d), can't contain meta data of size %d", sourceSize, metaSize)
	}

	var metaReader *gzip.Reader
	if metaReader, err = gzip.NewReader(io.NewSectionReader(source, sourceSize-8-metaSize, metaSize)); err != nil {
		return nil, fmt.Errorf("couldn't read meta data: %s", err)
	}

	var metaBytes []byte
	if metaBytes, err = ioutil.ReadAll(metaReader); err != nil {
		return nil, fmt.Errorf("couldn't read meta data: %s", err)
	}

	result := &Reader{
		currentFrameNdx: -1,
		currentPackNdx:  -1,
	}

	if err = (&result.Meta).Unmarshal(metaBytes); err != nil {
		return nil, fmt.Errorf("couldn't parse meta data: %s", err)
	}

	result.indexTimeOffsets = make([]float64, len(result.Meta.PackIndex.Entries))
	result.indexFileOffsets = make([]int64, len(result.Meta.PackIndex.Entries)+1)
	fileOffset := int64(0)
	for i, entry := range result.Meta.PackIndex.Entries {
		result.indexTimeOffsets[i] = entry.Offset
		result.indexFileOffsets[i] = fileOffset
		fileOffset += int64(entry.PackSize)
	}
	result.indexFileOffsets[len(result.Meta.PackIndex.Entries)] = fileOffset

	result.packs = io.NewSectionReader(source, 0, sourceSize-8-metaSize)

	return result, nil
}

func (reader *Reader) readPack(packNdx int) error {
	if reader.currentPackNdx == packNdx {
		// already have pack
		return nil
	}

	packReader, err := gzip.NewReader(io.NewSectionReader(
		reader.packs,
		reader.indexFileOffsets[packNdx],
		reader.indexFileOffsets[packNdx+1]-reader.indexFileOffsets[packNdx]))
	if err != nil {
		return fmt.Errorf("couldn't read pack at index %d: %s", packNdx, err)
	}

	packBytes, err := ioutil.ReadAll(packReader)
	if err != nil {
		return fmt.Errorf("couldn't unzip pack at index %d: %s", packNdx, err)
	}

	var framePack FramePack
	if err := (&framePack).Unmarshal(packBytes); err != nil {
		return fmt.Errorf("couldn't parse frame pack at index %d: %s", packNdx, err)
	}

	if 0 == len(framePack.Frames) {
		return fmt.Errorf("frame pack at index %d is empty", packNdx)
	}

	reader.currentPackNdx = packNdx
	reader.currentPack = &framePack
	reader.currentFrameNdx = 0

	return nil
}

func (reader *Reader) Seek(offset float64) error {
	if 0 == len(reader.indexTimeOffsets) {
		return nil
	}

	// search the lowest index so that the first offset in the next pack is
	// larger than the one we seek
	packNdx := sort.Search(len(reader.indexTimeOffsets)-1, func(i int) bool {
		return reader.indexTimeOffsets[i+1] > offset
	})

	if err := reader.readPack(packNdx); err != nil {
		return err
	}

	curFrames := reader.currentPack.Frames
	reader.currentFrameNdx = sort.Search(len(curFrames)-1, func(i int) bool {
		return curFrames[i+1].Offset > offset
	})

	return nil
}

// returns nil as Frame without error if there are no more frames
func (reader *Reader) ReadFrame() (*Frame, error) {
	if -1 == reader.currentPackNdx {
		if err := reader.readPack(0); err != nil {
			return nil, err
		}
	}
	if reader.currentFrameNdx >= len(reader.currentPack.Frames) {
		if reader.currentPackNdx+1 >= len(reader.indexTimeOffsets) {
			return nil, nil
		}
		if err := reader.readPack(reader.currentPackNdx + 1); err != nil {
			return nil, err
		}
	}

	result := &reader.currentPack.Frames[reader.currentFrameNdx]
	reader.currentFrameNdx++
	return result, nil
}
