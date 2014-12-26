package libtsm

// #cgo pkg-config: libtsm
// #include <libtsm.h>
// #include "bitfields.h"
import "C"

type ScreenFlags uint32

const (
	ScreenInsertMode ScreenFlags = C.TSM_SCREEN_INSERT_MODE
	ScreenAutoWrap   ScreenFlags = C.TSM_SCREEN_AUTO_WRAP
	ScreenRelOrigin  ScreenFlags = C.TSM_SCREEN_REL_ORIGIN
	ScreenInverse    ScreenFlags = C.TSM_SCREEN_INVERSE
	ScreenHideCursor ScreenFlags = C.TSM_SCREEN_HIDE_CURSOR
	ScreenFixedPos   ScreenFlags = C.TSM_SCREEN_FIXED_POS
	ScreenAlternate  ScreenFlags = C.TSM_SCREEN_ALTERNATE
)

type VteModifier uint32

const (
	ModifierShift   VteModifier = C.TSM_SHIFT_MASK
	ModifierLock    VteModifier = C.TSM_LOCK_MASK
	ModifierControl VteModifier = C.TSM_CONTROL_MASK
	ModifierAlt     VteModifier = C.TSM_ALT_MASK
	ModifierLogo    VteModifier = C.TSM_LOGO_MASK
)

// internal bitfield conversion

const (
	screenAttrBold      uint32 = C.GO_TSM_SCREEN_ATTR_BOLD
	screenAttrUnderline uint32 = C.GO_TSM_SCREEN_ATTR_UNDERLINE
	screenAttrInverse   uint32 = C.GO_TSM_SCREEN_ATTR_INVERSE
	screenAttrProtect   uint32 = C.GO_TSM_SCREEN_ATTR_PROTECT
	screenAttrBlink     uint32 = C.GO_TSM_SCREEN_ATTR_BLINK
)
