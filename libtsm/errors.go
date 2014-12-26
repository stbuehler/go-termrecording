package libtsm

// #include <string.h>
import "C"

// tsm error return values are negative errno values
func tsmErrorString(result C.int) string {
	return C.GoString(C.strerror(-result))
}
