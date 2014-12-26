package libtsm

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

// #cgo pkg-config: libtsm
// #include <libtsm.h>
// #include "callback_wrapper.h"
import "C"

type Age uint32

type screenReference struct {
	ptr *C.struct_tsm_screen
}

type Screen struct {
	ref *screenReference
}

func finalizeScreen(screenRef *screenReference) {
	if nil != screenRef.ptr {
		C.tsm_screen_unref(screenRef.ptr)
		screenRef.ptr = nil
	}
}

func NewScreen() (screen Screen, err error) {
	screen.ref = &screenReference{}
	runtime.SetFinalizer(screen.ref, finalizeScreen)
	failed, _ := C.tsm_screen_new(&screen.ref.ptr, nil, nil)
	if failed != 0 {
		err = fmt.Errorf("tsm_screen_new failed")
	}
	return
}

func (screen Screen) GetSize() (uint, uint) {
	width, _ := C.tsm_screen_get_width(screen.ref.ptr)
	height, _ := C.tsm_screen_get_height(screen.ref.ptr)
	return uint(width), uint(height)
}

func (screen Screen) Resize(x uint, y uint) error {
	failed, _ := C.tsm_screen_resize(screen.ref.ptr, C.uint(x), C.uint(y))
	if failed != 0 {
		return fmt.Errorf("tsm_screen_resize failed: %s", tsmErrorString(failed))
	}
	return nil
}
func (screen Screen) SetMargins(top uint, bottom uint) error {
	failed, _ := C.tsm_screen_set_margins(screen.ref.ptr, C.uint(top), C.uint(bottom))
	if failed != 0 {
		return fmt.Errorf("tsm_screen_set_margins failed: %s", tsmErrorString(failed))
	}
	return nil
}
func (screen Screen) SetMaxScrollbuffer(max uint) {
	C.tsm_screen_set_max_sb(screen.ref.ptr, C.uint(max))
}
func (screen Screen) ClearScrollbuffer() {
	C.tsm_screen_clear_sb(screen.ref.ptr)
}
func (screen Screen) ScrollbufferUp(num uint) {
	C.tsm_screen_sb_up(screen.ref.ptr, C.uint(num))
}
func (screen Screen) ScrollbufferDown(num uint) {
	C.tsm_screen_sb_down(screen.ref.ptr, C.uint(num))
}
func (screen Screen) ScrollbufferPageUp(num uint) {
	C.tsm_screen_sb_page_up(screen.ref.ptr, C.uint(num))
}
func (screen Screen) ScrollbufferPageDown(num uint) {
	C.tsm_screen_sb_page_down(screen.ref.ptr, C.uint(num))
}
func (screen Screen) ScrollbufferReset() {
	C.tsm_screen_sb_reset(screen.ref.ptr)
}
func (screen Screen) SetDefAttr(attr *ScreenAttr) {
	C.tsm_screen_set_def_attr(screen.ref.ptr, putScreenAttr(attr))
}
func (screen Screen) Reset() {
	C.tsm_screen_reset(screen.ref.ptr)
}
func (screen Screen) SetFlags(flags ScreenFlags) {
	C.tsm_screen_set_flags(screen.ref.ptr, C.uint(flags))
}
func (screen Screen) ResetFlags(flags ScreenFlags) {
	C.tsm_screen_reset_flags(screen.ref.ptr, C.uint(flags))
}
func (screen Screen) GetFlags() ScreenFlags {
	flags, _ := C.tsm_screen_get_flags(screen.ref.ptr)
	return ScreenFlags(flags)
}
func (screen Screen) GetCursor() (uint, uint) {
	x, _ := C.tsm_screen_get_cursor_x(screen.ref.ptr)
	y, _ := C.tsm_screen_get_cursor_y(screen.ref.ptr)
	return uint(x), uint(y)
}
func (screen Screen) SetTabstop() {
	C.tsm_screen_set_tabstop(screen.ref.ptr)
}
func (screen Screen) ResetTabstop() {
	C.tsm_screen_reset_tabstop(screen.ref.ptr)
}
func (screen Screen) ResetAllTabstops() {
	C.tsm_screen_reset_all_tabstops(screen.ref.ptr)
}
func (screen Screen) Write(ch rune, attr *ScreenAttr) {
	C.tsm_screen_write(screen.ref.ptr, C.tsm_symbol_t(ch), putScreenAttr(attr))
}
func (screen Screen) Newline() {
	C.tsm_screen_newline(screen.ref.ptr)
}
func (screen Screen) ScrollUp(num uint) {
	C.tsm_screen_scroll_up(screen.ref.ptr, C.uint(num))
}
func (screen Screen) ScrollDown(num uint) {
	C.tsm_screen_scroll_down(screen.ref.ptr, C.uint(num))
}
func (screen Screen) MoveTo(x uint, y uint) {
	C.tsm_screen_move_to(screen.ref.ptr, C.uint(x), C.uint(y))
}
func (screen Screen) MoveUp(num uint, scroll bool) {
	C.tsm_screen_move_up(screen.ref.ptr, C.uint(num), C.bool(scroll))
}
func (screen Screen) MoveDown(num uint, scroll bool) {
	C.tsm_screen_move_down(screen.ref.ptr, C.uint(num), C.bool(scroll))
}
func (screen Screen) MoveLeft(num uint) {
	C.tsm_screen_move_left(screen.ref.ptr, C.uint(num))
}
func (screen Screen) MoveRight(num uint) {
	C.tsm_screen_move_right(screen.ref.ptr, C.uint(num))
}
func (screen Screen) MoveLineEnd() {
	C.tsm_screen_move_line_end(screen.ref.ptr)
}
func (screen Screen) MoveLineHome() {
	C.tsm_screen_move_line_home(screen.ref.ptr)
}
func (screen Screen) TabRight(num uint) {
	C.tsm_screen_tab_right(screen.ref.ptr, C.uint(num))
}
func (screen Screen) TabLeft(num uint) {
	C.tsm_screen_tab_left(screen.ref.ptr, C.uint(num))
}
func (screen Screen) InsertLines(num uint) {
	C.tsm_screen_insert_lines(screen.ref.ptr, C.uint(num))
}
func (screen Screen) DeleteLines(num uint) {
	C.tsm_screen_delete_lines(screen.ref.ptr, C.uint(num))
}
func (screen Screen) InsertChars(num uint) {
	C.tsm_screen_insert_chars(screen.ref.ptr, C.uint(num))
}
func (screen Screen) DeleteChars(num uint) {
	C.tsm_screen_delete_chars(screen.ref.ptr, C.uint(num))
}
func (screen Screen) EraseCursor() {
	C.tsm_screen_erase_cursor(screen.ref.ptr)
}
func (screen Screen) EraseChars(num uint) {
	C.tsm_screen_erase_chars(screen.ref.ptr, C.uint(num))
}
func (screen Screen) EraseCursorToEnd(protect bool) {
	C.tsm_screen_erase_cursor_to_end(screen.ref.ptr, C.bool(protect))
}
func (screen Screen) EraseHomeToCursor(protect bool) {
	C.tsm_screen_erase_home_to_cursor(screen.ref.ptr, C.bool(protect))
}
func (screen Screen) EraseCurrentLine(protect bool) {
	C.tsm_screen_erase_current_line(screen.ref.ptr, C.bool(protect))
}
func (screen Screen) EraseScreenToCursor(protect bool) {
	C.tsm_screen_erase_screen_to_cursor(screen.ref.ptr, C.bool(protect))
}
func (screen Screen) EraseCursorToScreen(protect bool) {
	C.tsm_screen_erase_cursor_to_screen(screen.ref.ptr, C.bool(protect))
}
func (screen Screen) EraseScreen(protect bool) {
	C.tsm_screen_erase_screen(screen.ref.ptr, C.bool(protect))
}
func (screen Screen) SelectionReset() {
	C.tsm_screen_selection_reset(screen.ref.ptr)
}
func (screen Screen) SelectionStart(posx uint, posy uint) {
	C.tsm_screen_selection_start(screen.ref.ptr, C.uint(posx), C.uint(posy))
}
func (screen Screen) SelectionTarget(posx uint, posy uint) {
	C.tsm_screen_selection_target(screen.ref.ptr, C.uint(posx), C.uint(posy))
}

func (screen Screen) SelectionCopy() (string, error) {
	var buf *C.char
	len, _ := C.tsm_screen_selection_copy(screen.ref.ptr, &buf)
	if len < 0 {
		return "", fmt.Errorf("tsm_screen_selection_copy failed: %s",
			tsmErrorString(len))
	}
	return C.GoStringN(buf, len), nil
}

type ScreenDrawCallback func(id uint32, character string, width uint, x uint, y uint, attr *ScreenAttr, age Age) bool

type go_wrap_tsm_screen_draw_cb_data struct {
	callback ScreenDrawCallback
}

//export go_wrap_tsm_screen_draw_cb
func go_wrap_tsm_screen_draw_cb(id C.uint32_t,
	ch *C.uint32_t,
	chLen C.size_t,
	width C.uint,
	posx C.uint,
	posy C.uint,
	attr *C.struct_tsm_screen_attr,
	age C.tsm_age_t,
	data unsafe.Pointer) C.int {

	chHdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(ch)),
		Len:  int(chLen),
		Cap:  int(chLen),
	}
	chSlice := *(*[]C.uint32_t)(unsafe.Pointer(&chHdr))
	chRunes := make([]rune, len(chSlice))
	for i, r := range chSlice {
		chRunes[i] = rune(r)
	}
	// a character can consist of multiple combined runes, or none at all
	character := string(chRunes)

	myData := (*go_wrap_tsm_screen_draw_cb_data)(data)
	success := myData.callback(uint32(id), character, uint(width), uint(posx), uint(posy), getScreenAttr(attr), Age(age))
	if success {
		return 0
	} else {
		return 1
	}
}

func (screen Screen) Draw(callback ScreenDrawCallback) Age {
	data := go_wrap_tsm_screen_draw_cb_data{
		callback: callback,
	}
	age, _ := C.go_tsm_screen_draw(screen.ref.ptr, unsafe.Pointer(&data))
	return Age(age)
}
