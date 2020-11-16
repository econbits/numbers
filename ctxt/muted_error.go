package ctxt

// bitmap that mutes specific errors.
type MutedError uint8

// Flags to mute an error.
const (
	// Mutes ErrClamped errors.
	MutedErrorClamped = 0x01

	// Mutes ErrDivByZero errors.
	MutedErrorDivByZero = 0x02

	// Mutes ErrInexact errors.
	MutedErrorInexact = 0x04

	// Mutes ErrInvalidOp errors.
	MutedErrorInvalidOp = 0x08

	// Mutes ErrOverflow errors.
	MutedErrorOverflow = 0x10

	// Mutes ErrRounded errors.
	MutedErrorRounded = 0x20

	// Mutes ErrSubnormal errors.
	MutedErrorSubnormal = 0x40

	// Mutes ErrUnderflow errors.
	MutedErrorUnderflow = 0x80

	// Mutes ALL errors.
	MutedErrorAll = 0xff

	// Mutes NO errors.
	MutedErrorNone = 0x00
)
