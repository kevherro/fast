package math

import (
	stdMath "math"
	"unsafe"
)

// QRsqrt approximates the inverse square root of x using a fast method
// from the Quake III Arena source code, which involves bit manipulation
// and two iterations of Newton's method.
func QRsqrt(x float32) float32 {
	if x == 0 {
		return float32(stdMath.Inf(1))
	}

	x2 := x * 0.5

	// Reinterpret the bits of the floating point number as an int32.
	i := *(*int32)(unsafe.Pointer(&x))

	// This is a magic constant which provides a good
	// first approximation of the inverse square root.
	i = 0x5f3759df - (i >> 1)

	// Convert our int32 back to a floating point number.
	y := *(*float32)(unsafe.Pointer(&i))

	// Perform two iterations of Newton's method to
	// refine our estimate.
	y = y * (1.5 - (x2 * y * y))
	y = y * (1.5 - (x2 * y * y)) // This can be deleted.

	return y
}
