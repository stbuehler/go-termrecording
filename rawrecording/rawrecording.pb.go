// Code generated by protoc-gen-gogo.
// source: rawrecording.proto
// DO NOT EDIT!

/*
	Package rawrecording is a generated protocol buffer package.

	It is generated from these files:
		rawrecording.proto

	It has these top-level messages:
		Frame
		TerminalSize
		FramePack
		PackIndexEntry
		PackIndex
		Meta
*/
package rawrecording

import proto "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto/gogo.pb"

import io "io"
import math1 "math"
import fmt "fmt"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"

import fmt1 "fmt"
import strings "strings"
import reflect "reflect"

import math2 "math"

import fmt2 "fmt"
import strings1 "strings"
import github_com_gogo_protobuf_proto1 "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect1 "reflect"

import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// single raw frame
type Frame struct {
	// time offset from the beginning of the recording
	Offset float64 `protobuf:"fixed64,1,req,name=offset" json:"offset"`
	// recorded data
	Data []byte `protobuf:"bytes,2,req,name=data" json:"data"`
}

func (m *Frame) Reset()      { *m = Frame{} }
func (*Frame) ProtoMessage() {}

func (m *Frame) GetOffset() float64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Frame) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type TerminalSize struct {
	Columns uint32 `protobuf:"varint,1,req,name=columns" json:"columns"`
	Rows    uint32 `protobuf:"varint,2,req,name=rows" json:"rows"`
}

func (m *TerminalSize) Reset()      { *m = TerminalSize{} }
func (*TerminalSize) ProtoMessage() {}

func (m *TerminalSize) GetColumns() uint32 {
	if m != nil {
		return m.Columns
	}
	return 0
}

func (m *TerminalSize) GetRows() uint32 {
	if m != nil {
		return m.Rows
	}
	return 0
}

// list of raw frames
type FramePack struct {
	Frames       []Frame       `protobuf:"bytes,1,rep,name=frames" json:"frames"`
	TerminalSize *TerminalSize `protobuf:"bytes,2,opt,name=terminalSize" json:"terminalSize,omitempty"`
}

func (m *FramePack) Reset()      { *m = FramePack{} }
func (*FramePack) ProtoMessage() {}

func (m *FramePack) GetFrames() []Frame {
	if m != nil {
		return m.Frames
	}
	return nil
}

func (m *FramePack) GetTerminalSize() *TerminalSize {
	if m != nil {
		return m.TerminalSize
	}
	return nil
}

type PackIndexEntry struct {
	// time offset from the beginning of the recording of the first frame
	// in a pack
	Offset float64 `protobuf:"fixed64,1,req,name=offset" json:"offset"`
	// size of serialized pack
	PackSize uint32 `protobuf:"varint,2,req,name=pack_size" json:"pack_size"`
}

func (m *PackIndexEntry) Reset()      { *m = PackIndexEntry{} }
func (*PackIndexEntry) ProtoMessage() {}

func (m *PackIndexEntry) GetOffset() float64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *PackIndexEntry) GetPackSize() uint32 {
	if m != nil {
		return m.PackSize
	}
	return 0
}

// describe a list of Packs; packs must be splitted when they get too big.
type PackIndex struct {
	Entries []PackIndexEntry `protobuf:"bytes,1,rep,name=entries" json:"entries"`
}

func (m *PackIndex) Reset()      { *m = PackIndex{} }
func (*PackIndex) ProtoMessage() {}

func (m *PackIndex) GetEntries() []PackIndexEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

// footer contains meta data
type Meta struct {
	PackIndex       PackIndex    `protobuf:"bytes,1,req,name=packIndex" json:"packIndex"`
	MaxTerminalSize TerminalSize `protobuf:"bytes,2,req,name=maxTerminalSize" json:"maxTerminalSize"`
}

func (m *Meta) Reset()      { *m = Meta{} }
func (*Meta) ProtoMessage() {}

func (m *Meta) GetPackIndex() PackIndex {
	if m != nil {
		return m.PackIndex
	}
	return PackIndex{}
}

func (m *Meta) GetMaxTerminalSize() TerminalSize {
	if m != nil {
		return m.MaxTerminalSize
	}
	return TerminalSize{}
}

func init() {
}
func (m *Frame) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			var v uint64
			i := index + 8
			if i > l {
				return io.ErrUnexpectedEOF
			}
			index = i
			v = uint64(data[i-8])
			v |= uint64(data[i-7]) << 8
			v |= uint64(data[i-6]) << 16
			v |= uint64(data[i-5]) << 24
			v |= uint64(data[i-4]) << 32
			v |= uint64(data[i-3]) << 40
			v |= uint64(data[i-2]) << 48
			v |= uint64(data[i-1]) << 56
			m.Offset = math1.Float64frombits(v)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, data[index:postIndex]...)
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := github_com_gogo_protobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			index += skippy
		}
	}
	return nil
}
func (m *TerminalSize) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Columns", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				m.Columns |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rows", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				m.Rows |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := github_com_gogo_protobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			index += skippy
		}
	}
	return nil
}
func (m *FramePack) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Frames", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Frames = append(m.Frames, Frame{})
			m.Frames[len(m.Frames)-1].Unmarshal(data[index:postIndex])
			index = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TerminalSize", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TerminalSize == nil {
				m.TerminalSize = &TerminalSize{}
			}
			if err := m.TerminalSize.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := github_com_gogo_protobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			index += skippy
		}
	}
	return nil
}
func (m *PackIndexEntry) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			var v uint64
			i := index + 8
			if i > l {
				return io.ErrUnexpectedEOF
			}
			index = i
			v = uint64(data[i-8])
			v |= uint64(data[i-7]) << 8
			v |= uint64(data[i-6]) << 16
			v |= uint64(data[i-5]) << 24
			v |= uint64(data[i-4]) << 32
			v |= uint64(data[i-3]) << 40
			v |= uint64(data[i-2]) << 48
			v |= uint64(data[i-1]) << 56
			m.Offset = math1.Float64frombits(v)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PackSize", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				m.PackSize |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := github_com_gogo_protobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			index += skippy
		}
	}
	return nil
}
func (m *PackIndex) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entries = append(m.Entries, PackIndexEntry{})
			m.Entries[len(m.Entries)-1].Unmarshal(data[index:postIndex])
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := github_com_gogo_protobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			index += skippy
		}
	}
	return nil
}
func (m *Meta) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PackIndex", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PackIndex.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTerminalSize", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxTerminalSize.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := github_com_gogo_protobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			index += skippy
		}
	}
	return nil
}
func (this *Frame) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Frame{`,
		`Offset:` + fmt1.Sprintf("%v", this.Offset) + `,`,
		`Data:` + fmt1.Sprintf("%v", this.Data) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TerminalSize) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TerminalSize{`,
		`Columns:` + fmt1.Sprintf("%v", this.Columns) + `,`,
		`Rows:` + fmt1.Sprintf("%v", this.Rows) + `,`,
		`}`,
	}, "")
	return s
}
func (this *FramePack) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&FramePack{`,
		`Frames:` + strings.Replace(strings.Replace(fmt1.Sprintf("%v", this.Frames), "Frame", "Frame", 1), `&`, ``, 1) + `,`,
		`TerminalSize:` + strings.Replace(fmt1.Sprintf("%v", this.TerminalSize), "TerminalSize", "TerminalSize", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PackIndexEntry) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PackIndexEntry{`,
		`Offset:` + fmt1.Sprintf("%v", this.Offset) + `,`,
		`PackSize:` + fmt1.Sprintf("%v", this.PackSize) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PackIndex) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PackIndex{`,
		`Entries:` + strings.Replace(strings.Replace(fmt1.Sprintf("%v", this.Entries), "PackIndexEntry", "PackIndexEntry", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Meta) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Meta{`,
		`PackIndex:` + strings.Replace(strings.Replace(this.PackIndex.String(), "PackIndex", "PackIndex", 1), `&`, ``, 1) + `,`,
		`MaxTerminalSize:` + strings.Replace(strings.Replace(this.MaxTerminalSize.String(), "TerminalSize", "TerminalSize", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringRawrecording(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt1.Sprintf("*%v", pv)
}
func (m *Frame) Size() (n int) {
	var l int
	_ = l
	n += 9
	l = len(m.Data)
	n += 1 + l + sovRawrecording(uint64(l))
	return n
}

func (m *TerminalSize) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovRawrecording(uint64(m.Columns))
	n += 1 + sovRawrecording(uint64(m.Rows))
	return n
}

func (m *FramePack) Size() (n int) {
	var l int
	_ = l
	if len(m.Frames) > 0 {
		for _, e := range m.Frames {
			l = e.Size()
			n += 1 + l + sovRawrecording(uint64(l))
		}
	}
	if m.TerminalSize != nil {
		l = m.TerminalSize.Size()
		n += 1 + l + sovRawrecording(uint64(l))
	}
	return n
}

func (m *PackIndexEntry) Size() (n int) {
	var l int
	_ = l
	n += 9
	n += 1 + sovRawrecording(uint64(m.PackSize))
	return n
}

func (m *PackIndex) Size() (n int) {
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
			l = e.Size()
			n += 1 + l + sovRawrecording(uint64(l))
		}
	}
	return n
}

func (m *Meta) Size() (n int) {
	var l int
	_ = l
	l = m.PackIndex.Size()
	n += 1 + l + sovRawrecording(uint64(l))
	l = m.MaxTerminalSize.Size()
	n += 1 + l + sovRawrecording(uint64(l))
	return n
}

func sovRawrecording(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRawrecording(x uint64) (n int) {
	return sovRawrecording(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func NewPopulatedFrame(r randyRawrecording, easy bool) *Frame {
	this := &Frame{}
	this.Offset = r.Float64()
	if r.Intn(2) == 0 {
		this.Offset *= -1
	}
	v1 := r.Intn(100)
	this.Data = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.Data[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedTerminalSize(r randyRawrecording, easy bool) *TerminalSize {
	this := &TerminalSize{}
	this.Columns = r.Uint32()
	this.Rows = r.Uint32()
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedFramePack(r randyRawrecording, easy bool) *FramePack {
	this := &FramePack{}
	if r.Intn(10) != 0 {
		v2 := r.Intn(10)
		this.Frames = make([]Frame, v2)
		for i := 0; i < v2; i++ {
			v3 := NewPopulatedFrame(r, easy)
			this.Frames[i] = *v3
		}
	}
	if r.Intn(10) != 0 {
		this.TerminalSize = NewPopulatedTerminalSize(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedPackIndexEntry(r randyRawrecording, easy bool) *PackIndexEntry {
	this := &PackIndexEntry{}
	this.Offset = r.Float64()
	if r.Intn(2) == 0 {
		this.Offset *= -1
	}
	this.PackSize = r.Uint32()
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedPackIndex(r randyRawrecording, easy bool) *PackIndex {
	this := &PackIndex{}
	if r.Intn(10) != 0 {
		v4 := r.Intn(10)
		this.Entries = make([]PackIndexEntry, v4)
		for i := 0; i < v4; i++ {
			v5 := NewPopulatedPackIndexEntry(r, easy)
			this.Entries[i] = *v5
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedMeta(r randyRawrecording, easy bool) *Meta {
	this := &Meta{}
	v6 := NewPopulatedPackIndex(r, easy)
	this.PackIndex = *v6
	v7 := NewPopulatedTerminalSize(r, easy)
	this.MaxTerminalSize = *v7
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyRawrecording interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneRawrecording(r randyRawrecording) rune {
	res := rune(r.Uint32() % 1112064)
	if 55296 <= res {
		res += 2047
	}
	return res
}
func randStringRawrecording(r randyRawrecording) string {
	v8 := r.Intn(100)
	tmps := make([]rune, v8)
	for i := 0; i < v8; i++ {
		tmps[i] = randUTF8RuneRawrecording(r)
	}
	return string(tmps)
}
func randUnrecognizedRawrecording(r randyRawrecording, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldRawrecording(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldRawrecording(data []byte, r randyRawrecording, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateRawrecording(data, uint64(key))
		v9 := r.Int63()
		if r.Intn(2) == 0 {
			v9 *= -1
		}
		data = encodeVarintPopulateRawrecording(data, uint64(v9))
	case 1:
		data = encodeVarintPopulateRawrecording(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateRawrecording(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateRawrecording(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateRawrecording(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateRawrecording(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *Frame) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Frame) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x9
	i++
	i = encodeFixed64Rawrecording(data, i, uint64(math2.Float64bits(m.Offset)))
	data[i] = 0x12
	i++
	i = encodeVarintRawrecording(data, i, uint64(len(m.Data)))
	i += copy(data[i:], m.Data)
	return i, nil
}

func (m *TerminalSize) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *TerminalSize) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintRawrecording(data, i, uint64(m.Columns))
	data[i] = 0x10
	i++
	i = encodeVarintRawrecording(data, i, uint64(m.Rows))
	return i, nil
}

func (m *FramePack) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *FramePack) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Frames) > 0 {
		for _, msg := range m.Frames {
			data[i] = 0xa
			i++
			i = encodeVarintRawrecording(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.TerminalSize != nil {
		data[i] = 0x12
		i++
		i = encodeVarintRawrecording(data, i, uint64(m.TerminalSize.Size()))
		n1, err := m.TerminalSize.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *PackIndexEntry) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *PackIndexEntry) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x9
	i++
	i = encodeFixed64Rawrecording(data, i, uint64(math2.Float64bits(m.Offset)))
	data[i] = 0x10
	i++
	i = encodeVarintRawrecording(data, i, uint64(m.PackSize))
	return i, nil
}

func (m *PackIndex) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *PackIndex) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for _, msg := range m.Entries {
			data[i] = 0xa
			i++
			i = encodeVarintRawrecording(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Meta) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Meta) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintRawrecording(data, i, uint64(m.PackIndex.Size()))
	n2, err := m.PackIndex.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	data[i] = 0x12
	i++
	i = encodeVarintRawrecording(data, i, uint64(m.MaxTerminalSize.Size()))
	n3, err := m.MaxTerminalSize.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	return i, nil
}

func encodeFixed64Rawrecording(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Rawrecording(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintRawrecording(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (this *Frame) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings1.Join([]string{`&rawrecording.Frame{` +
		`Offset:` + fmt2.Sprintf("%#v", this.Offset),
		`Data:` + fmt2.Sprintf("%#v", this.Data) + `}`}, ", ")
	return s
}
func (this *TerminalSize) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings1.Join([]string{`&rawrecording.TerminalSize{` +
		`Columns:` + fmt2.Sprintf("%#v", this.Columns),
		`Rows:` + fmt2.Sprintf("%#v", this.Rows) + `}`}, ", ")
	return s
}
func (this *FramePack) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings1.Join([]string{`&rawrecording.FramePack{` +
		`Frames:` + strings1.Replace(fmt2.Sprintf("%#v", this.Frames), `&`, ``, 1),
		`TerminalSize:` + fmt2.Sprintf("%#v", this.TerminalSize) + `}`}, ", ")
	return s
}
func (this *PackIndexEntry) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings1.Join([]string{`&rawrecording.PackIndexEntry{` +
		`Offset:` + fmt2.Sprintf("%#v", this.Offset),
		`PackSize:` + fmt2.Sprintf("%#v", this.PackSize) + `}`}, ", ")
	return s
}
func (this *PackIndex) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings1.Join([]string{`&rawrecording.PackIndex{` +
		`Entries:` + strings1.Replace(fmt2.Sprintf("%#v", this.Entries), `&`, ``, 1) + `}`}, ", ")
	return s
}
func (this *Meta) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings1.Join([]string{`&rawrecording.Meta{` +
		`PackIndex:` + strings1.Replace(this.PackIndex.GoString(), `&`, ``, 1),
		`MaxTerminalSize:` + strings1.Replace(this.MaxTerminalSize.GoString(), `&`, ``, 1) + `}`}, ", ")
	return s
}
func valueToGoStringRawrecording(v interface{}, typ string) string {
	rv := reflect1.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect1.Indirect(rv).Interface()
	return fmt2.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringRawrecording(e map[int32]github_com_gogo_protobuf_proto1.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings1.Join(ss, ",") + "}"
	return s
}
func (this *Frame) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Frame)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Offset != that1.Offset {
		return false
	}
	if !bytes.Equal(this.Data, that1.Data) {
		return false
	}
	return true
}
func (this *TerminalSize) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*TerminalSize)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Columns != that1.Columns {
		return false
	}
	if this.Rows != that1.Rows {
		return false
	}
	return true
}
func (this *FramePack) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*FramePack)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Frames) != len(that1.Frames) {
		return false
	}
	for i := range this.Frames {
		if !this.Frames[i].Equal(&that1.Frames[i]) {
			return false
		}
	}
	if !this.TerminalSize.Equal(that1.TerminalSize) {
		return false
	}
	return true
}
func (this *PackIndexEntry) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*PackIndexEntry)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Offset != that1.Offset {
		return false
	}
	if this.PackSize != that1.PackSize {
		return false
	}
	return true
}
func (this *PackIndex) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*PackIndex)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Entries) != len(that1.Entries) {
		return false
	}
	for i := range this.Entries {
		if !this.Entries[i].Equal(&that1.Entries[i]) {
			return false
		}
	}
	return true
}
func (this *Meta) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Meta)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.PackIndex.Equal(&that1.PackIndex) {
		return false
	}
	if !this.MaxTerminalSize.Equal(&that1.MaxTerminalSize) {
		return false
	}
	return true
}
