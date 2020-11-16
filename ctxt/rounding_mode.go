package ctxt

// Rounding is applied when a result coefficient has more significant digits
// than the value of precision; in this case the result coefficient is shortened
// to precision digits and may then be incremented by one (which may require a
// further shortening), depending on the rounding algorithm selected and the
// remaining digits of the original coefficient. The exponent is adjusted to
// compensate for any shortening.
type RoundingMode uint8

// Rounding modes.
const (
	// (Round toward 0; truncate.) The discarded digits are ignored;
	// the result is unchanged.
	RoundingModeDown = 0x01

	// If the discarded digits represent greater than or equal to half (0.5) of
	// the value of a one in the next left position then the result coefficient
	// should be incremented by 1 (rounded up). Otherwise the discarded digits
	// are ignored.
	RoundingModeHalfUp = 0x02

	// If the discarded digits represent greater than half (0.5) the value of a
	// one in the next left position then the result coefficient should be
	// incremented by 1 (rounded up). If they represent less than half, then
	// the result coefficient is not adjusted (that is, the discarded digits are
	// ignored).
	// Otherwise (they represent exactly half) the result coefficient is unaltered
	// if its rightmost digit is even, or incremented by 1 (rounded up) if its
	// rightmost digit is odd (to make an even digit).
	RoundingModeHalfEven = 0x04

	// (Round toward +∞.) If all of the discarded digits are zero or if the sign is
	// 1 the result is unchanged. Otherwise, the result coefficient should be
	// incremented by 1 (rounded up).
	RoundingModeCeiling = 0x08

	// (Round toward -∞.) If all of the discarded digits are zero or if the sign is
	// 0 the result is unchanged. Otherwise, the sign is 1 and the result coefficient
	// should be incremented by 1.
	RoundingModeFloor = 0x10

	// If the discarded digits represent greater than half (0.5) of the value of a one
	// in the next left position then the result coefficient should be incremented by
	// 1 (rounded up). Otherwise (the discarded digits are 0.5 or less) the discarded
	// digits are ignored.
	RoundingModeHalfDown = 0x20

	// (Round away from 0.) If all of the discarded digits are zero the result is
	// unchanged. Otherwise, the result coefficient should be incremented by 1
	// (rounded up).
	RoundingModeUp = 0x40

	// (Round zero or five away from 0.) The same as round-up, except that rounding up
	// only occurs if the digit to be rounded up is 0 or 5, and after overflow the
	// result is the same as for round-down.
	RoundingMode05Up = 0x80
)
