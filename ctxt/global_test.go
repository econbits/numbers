package ctxt

import (
	"testing"
)

func TestGlobalDefault(t *testing.T) {
	expected := Default()
	got := GlobalCtxt()
	if !expected.IsEqual(got) {
		t.Errorf("GlobalCtxt(): expected=%v, got=%v", expected, got)
	}
}

func TestChangeRoundingModeInGlobalCtxt(t *testing.T) {
	ctx := GlobalCtxt()
	ctx = ctx.CopyWithRoundingMode(RoundingModeDown)
	if ctx.RoundingMode() != RoundingModeDown {
		t.Errorf("Expected Rounding Mode: %v; got: %v", RoundingModeDown, ctx.RoundingMode())
	}
	if ctx.IsEqual(GlobalCtxt()) {
		t.Errorf("GlobalCtxt(): global state was updated by accident")
	}
	SetGlobalCtxt(ctx)
	if Default().IsEqual(GlobalCtxt()) {
		t.Errorf("GlobalCtxt() was the same as Default() after updating it")
	}
	ResetGlobalCtxt()
	if !Default().IsEqual(GlobalCtxt()) {
		t.Errorf("GlobalCtxt() was the _not_ same as Default() after resetting it")
	}
}

func TestChangeMutedErrorInGlobalCtxt(t *testing.T) {
	ctx := GlobalCtxt()
	ctx = ctx.CopyWithMutedErrors(MutedErrorNone)
	if ctx.IsInexactMuted() {
		t.Errorf("Expected Inexact Error to be muted")
	}
	if ctx.IsEqual(GlobalCtxt()) {
		t.Errorf("GlobalCtxt(): global state was updated by accident")
	}
	SetGlobalCtxt(ctx)
	if Default().IsEqual(GlobalCtxt()) {
		t.Errorf("GlobalCtxt() was the same as Default() after updating it")
	}
	ResetGlobalCtxt()
	if !Default().IsEqual(GlobalCtxt()) {
		t.Errorf("GlobalCtxt() was the _not_ same as Default() after resetting it")
	}
}
