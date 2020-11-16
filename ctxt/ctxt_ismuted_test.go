package ctxt

import (
	"testing"
)

func TestNonMutedErrors(t *testing.T) {
	ctx := NewWithMutedErrors(RoundingModeHalfEven, MutedErrorNone)
	if ctx.IsClampedMuted() {
		t.Errorf("Clamped Error is muted")
	}
	if ctx.IsDivByZeroMuted() {
		t.Errorf("Division by Zero Error is muted")
	}
	if ctx.IsInexactMuted() {
		t.Errorf("Inexact Error is muted")
	}
	if ctx.IsInvalidOpMuted() {
		t.Errorf("Invalida Operation Error is muted")
	}
	if ctx.IsOverflowMuted() {
		t.Errorf("Overflow Error is muted")
	}
	if ctx.IsRoundedMuted() {
		t.Errorf("Rounded Error is muted")
	}
	if ctx.IsSubnormalMuted() {
		t.Errorf("Subnormal Error is muted")
	}
	if ctx.IsUnderflowMuted() {
		t.Errorf("Underflow Error is muted")
	}
}

func TestMutedErrors(t *testing.T) {
	ctx := NewWithMutedErrors(RoundingModeHalfEven, MutedErrorAll)
	if !ctx.IsClampedMuted() {
		t.Errorf("Clamped Error is not muted")
	}
	if !ctx.IsDivByZeroMuted() {
		t.Errorf("Division by Zero Error is not muted")
	}
	if !ctx.IsInexactMuted() {
		t.Errorf("Inexact Error is not muted")
	}
	if !ctx.IsInvalidOpMuted() {
		t.Errorf("Invalida Operation Error is not muted")
	}
	if !ctx.IsOverflowMuted() {
		t.Errorf("Overflow Error is not muted")
	}
	if !ctx.IsRoundedMuted() {
		t.Errorf("Rounded Error is not muted")
	}
	if !ctx.IsSubnormalMuted() {
		t.Errorf("Subnormal Error is not muted")
	}
	if !ctx.IsUnderflowMuted() {
		t.Errorf("Underflow Error is not muted")
	}
}
