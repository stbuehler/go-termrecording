package libtsm

// #cgo pkg-config: libtsm
// #include <libtsm.h>
// #include "bitfields.h"
import "C"

type RGB struct {
	Red, Green, Blue uint8
}

type Color struct {
	Code int8 /* color code or <0 for RGB */
	RGB
}

type ScreenAttr struct {
	Foreground Color
	Background Color
	Bold       bool /* bold character */
	Underline  bool /* underlined character */
	Inverse    bool /* inverse colors */
	Protect    bool /* cannot be erased */
	Blink      bool /* blinking character */
}

func selectFlag(flag bool, mask uint32) uint32 {
	if flag {
		return mask
	} else {
		return 0
	}
}

func putScreenAttr(attr *ScreenAttr) *C.struct_tsm_screen_attr {
	if attr == nil {
		return nil
	}
	cattr := C.struct_tsm_screen_attr{
		fccode: C.int8_t(attr.Foreground.Code),
		bccode: C.int8_t(attr.Background.Code),
		fr:     C.uint8_t(attr.Foreground.Red),
		fg:     C.uint8_t(attr.Foreground.Green),
		fb:     C.uint8_t(attr.Foreground.Blue),
		br:     C.uint8_t(attr.Background.Red),
		bg:     C.uint8_t(attr.Background.Green),
		bb:     C.uint8_t(attr.Background.Blue),
	}
	bits := (0 |
		selectFlag(attr.Bold, screenAttrBold) |
		selectFlag(attr.Underline, screenAttrUnderline) |
		selectFlag(attr.Inverse, screenAttrInverse) |
		selectFlag(attr.Protect, screenAttrProtect) |
		selectFlag(attr.Blink, screenAttrBlink) |
		0)
	C.go_set_tsm_screen_attr_bitfields(&cattr, C.uint32_t(bits))
	return &cattr
}

func getScreenAttr(cattr *C.struct_tsm_screen_attr) *ScreenAttr {
	if cattr == nil {
		return nil
	}
	bits := uint32(C.go_get_tsm_screen_attr_bitfields(cattr))
	attr := ScreenAttr{
		Foreground: Color{
			Code: int8(cattr.fccode),
			RGB: RGB{
				Red:   uint8(cattr.fr),
				Green: uint8(cattr.fg),
				Blue:  uint8(cattr.fb),
			},
		},
		Background: Color{
			Code: int8(cattr.bccode),
			RGB: RGB{
				Red:   uint8(cattr.br),
				Green: uint8(cattr.bg),
				Blue:  uint8(cattr.bb),
			},
		},
		Bold:      0 != bits&screenAttrBold,
		Underline: 0 != bits&screenAttrUnderline,
		Inverse:   0 != bits&screenAttrInverse,
		Protect:   0 != bits&screenAttrProtect,
		Blink:     0 != bits&screenAttrBlink,
	}
	return &attr
}
