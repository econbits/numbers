package ctxt

// A Context defines the rules that adjust how the operations are performed.
// There are 2 main components: the rounding mode and the errors to mute.
type Ctxt struct {
	rm RoundingMode
	me MutedError
}

// Returns a new Ctxt with the given rounding mode. It mutes the Inexact
// error.
func New(rm RoundingMode) Ctxt {
	return NewWithMutedErrors(rm, MutedErrorInexact)
}

// Returns a new Ctxt with the given rounding mode and muted errors
func NewWithMutedErrors(rm RoundingMode, me MutedError) Ctxt {
	return Ctxt{rm: rm, me: me}
}

// Returns the Ctxt that is used by default. It uses the rounding mode
// Half Even and mutes the error Inexact.
func Default() Ctxt {
	return New(RoundingModeHalfEven)
}

// Compares 2 Contexts
func (ctx Ctxt) IsEqual(ctx2 Ctxt) bool {
	return ctx.rm == ctx2.rm && ctx.me == ctx2.me
}

// Returns the Rounding Mode for the Context
func (ctx Ctxt) RoundingMode() RoundingMode {
	return ctx.rm
}

// Creates a new Context with the selected Rounding Mode, but keeping the
// Muted Errors as previously configured.
func (ctx Ctxt) CopyWithRoundingMode(rm RoundingMode) Ctxt {
	return NewWithMutedErrors(rm, ctx.me)
}

// Creates a new Context with the selected Muted Errors, but keeping the
// Rounding Mode as previously configured.
func (ctx Ctxt) CopyWithMutedErrors(me MutedError) Ctxt {
	return NewWithMutedErrors(ctx.rm, me)
}

// Returns true if the Clamped error is muted
func (ctx Ctxt) IsClampedMuted() bool {
	return (ctx.me & MutedErrorClamped) == MutedErrorClamped
}

// Returns true if the Division by Zero error is muted
func (ctx Ctxt) IsDivByZeroMuted() bool {
	return (ctx.me & MutedErrorDivByZero) == MutedErrorDivByZero
}

// Returns true if the Inexact error is muted
func (ctx Ctxt) IsInexactMuted() bool {
	return (ctx.me & MutedErrorInexact) == MutedErrorInexact
}

// Returns true if the Invalid Operation error is muted
func (ctx Ctxt) IsInvalidOpMuted() bool {
	return (ctx.me & MutedErrorInvalidOp) == MutedErrorInvalidOp
}

// Returns true if the Overflow error is muted
func (ctx Ctxt) IsOverflowMuted() bool {
	return (ctx.me & MutedErrorOverflow) == MutedErrorOverflow
}

// Returns true if the Rounded error is muted
func (ctx Ctxt) IsRoundedMuted() bool {
	return (ctx.me & MutedErrorRounded) == MutedErrorRounded
}

// Returns true if the Subnormal error is muted
func (ctx Ctxt) IsSubnormalMuted() bool {
	return (ctx.me & MutedErrorSubnormal) == MutedErrorSubnormal
}

// Returns true if the Underflow error is muted
func (ctx Ctxt) IsUnderflowMuted() bool {
	return (ctx.me & MutedErrorUnderflow) == MutedErrorUnderflow
}
