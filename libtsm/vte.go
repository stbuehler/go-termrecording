package libtsm

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

// #cgo pkg-config: libtsm
// #include <stdlib.h>
// #include <libtsm.h>
// #include "callback_wrapper.h"
import "C"

type VteWriteCallback func(data string)

type vteRef struct {
	ptr           *C.struct_tsm_vte
	writeCallback VteWriteCallback
}

type Vte struct {
	ref *vteRef
}

func finalizeVte(vteRef *vteRef) {
	if nil != vteRef.ptr {
		C.tsm_vte_unref(vteRef.ptr)
		vteRef.ptr = nil
	}
}

//export go_wrap_tsm_vte_write_cb
func go_wrap_tsm_vte_write_cb(u8 *C.char, len C.size_t, data unsafe.Pointer) {
	vte := (*vteRef)(data)
	if vte.writeCallback != nil {
		vte.writeCallback(C.GoStringN(u8, C.int(len)))
	}
}

func NewVte(screen Screen, writeCallback VteWriteCallback) (vte Vte, err error) {
	vte.ref = &vteRef{
		writeCallback: writeCallback,
	}
	runtime.SetFinalizer(vte.ref, finalizeVte)
	failed, _ := C.go_tsm_vte_new(&vte.ref.ptr, screen.ref.ptr, unsafe.Pointer(vte.ref))
	if failed != 0 {
		err = fmt.Errorf("tsm_vte_new failed: %s", tsmErrorString(failed))
	}
	return
}

const (
	PaletteDefault        = ""
	PaletteSolarized      = "solarized"
	PaletteSolarizedBlack = "solarized-black"
	PaletteSolarizedWhite = "solarized-white"
)

func (vte Vte) SetPalette(palette string) error {
	palettePtr := C.CString(palette)
	defer C.free(unsafe.Pointer(palettePtr))

	failed, _ := C.tsm_vte_set_palette(vte.ref.ptr, palettePtr)
	if failed != 0 {
		return fmt.Errorf("tsm_vte_set_palette('%s') failed: %s", palette, tsmErrorString(failed))
	}
	return nil
}

func (vte Vte) Reset() {
	C.tsm_vte_reset(vte.ref.ptr)
}
func (vte Vte) HardReset() {
	C.tsm_vte_hard_reset(vte.ref.ptr)
}

func (vte Vte) Input(data string) {
	dataHdr := *(*reflect.StringHeader)(unsafe.Pointer(&data))
	C.tsm_vte_input(vte.ref.ptr, (*C.char)(unsafe.Pointer(dataHdr.Data)), C.size_t(dataHdr.Len))
}

func (vte Vte) InputBytes(data []byte) {
	dataHdr := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	C.tsm_vte_input(vte.ref.ptr, (*C.char)(unsafe.Pointer(dataHdr.Data)), C.size_t(dataHdr.Len))
}

func (vte Vte) HandleKeyboard(keysym uint32, ascii uint32, mods VteModifier, unicode rune) bool {
	result, _ := C.tsm_vte_handle_keyboard(vte.ref.ptr, C.uint32_t(keysym), C.uint32_t(ascii), C.uint(mods), C.uint32_t(unicode))
	return bool(result)
}
