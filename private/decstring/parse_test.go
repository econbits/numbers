package decstring

import (
	"testing"
)

func TestPositiveInf(t *testing.T) {
	testruns := []string{"Inf", "inf", "+inf", "+Inf", "Infinity", "infinity", "+Infinity", "+infinity"}
	for _, txt := range testruns {
		d := Parse(txt)
		if !d.IsInf() {
			t.Errorf("Parse(%s): expected=Inf, got=%v", txt, d)
		}
		if d.DecParts() != nil {
			t.Errorf("Parse(%s): expected nil DecParts, got=%v", txt, d.DecParts())
		}
	}
}

func TestNegativeInf(t *testing.T) {
	testruns := []string{"-Inf", "-inf", "-Infinity", "-infinity"}
	for _, txt := range testruns {
		d := Parse(txt)
		if !d.IsMinusInf() {
			t.Errorf("Parse(%s): expected=-Inf, got=%v", txt, d)
		}
		if d.DecParts() != nil {
			t.Errorf("Parse(%s): expected nil DecParts, got=%v", txt, d.DecParts())
		}
	}
}

func TestPositiveNaN(t *testing.T) {
	testruns := []string{"NaN", "nan", "+nan", "nan 123"}
	for _, txt := range testruns {
		d := Parse(txt)
		if !d.IsNaN() {
			t.Errorf("Parse(%s): expected=NaN, got=%v", txt, d)
		}
		if d.DecParts() != nil {
			t.Errorf("Parse(%s): expected nil DecParts, got=%v", txt, d.DecParts())
		}
	}
}

func TestNegativeNaN(t *testing.T) {
	testruns := []string{"-NaN", "-nan"}
	for _, txt := range testruns {
		d := Parse(txt)
		if !d.IsMinusNaN() {
			t.Errorf("Parse(%s): expected=-NaN, got=%v", txt, d)
		}
		if d.DecParts() != nil {
			t.Errorf("Parse(%s): expected nil DecParts, got=%v", txt, d.DecParts())
		}
	}
}

func TestNotADecimal(t *testing.T) {
	testruns := []string{"abc"}
	for _, txt := range testruns {
		d := Parse(txt)
		if !d.IsSNaN() {
			t.Errorf("Parse(%s): expected=SNaN, got=%v", txt, d)
		}
		if d.DecParts() != nil {
			t.Errorf("Parse(%s): expected nil DecParts, got=%v", txt, d.DecParts())
		}
	}
}

func TestOne(t *testing.T) {
	testruns := []string{"1", "+1", "1.0", "+1.0", "1.", "+1."}
	for _, txt := range testruns {
		d := Parse(txt)
		if d.DecParts() == nil {
			t.Errorf("Parse(%s): expected DecParts, got=%v", txt, d)
		}
		if d.IsInf() {
			t.Errorf("Parse(%s): unexpected Inf", txt)
		}
		if d.IsMinusInf() {
			t.Errorf("Parse(%s): unexpected -Inf", txt)
		}
		if d.IsNaN() {
			t.Errorf("Parse(%s): unexpected NaN", txt)
		}
		if d.IsMinusNaN() {
			t.Errorf("Parse(%s): unexpected -NaN", txt)
		}
		if d.IsSNaN() {
			t.Errorf("Parse(%s): unexpected SNaN", txt)
		}
		if d.IsMinusSNaN() {
			t.Errorf("Parse(%s): unexpected -SNaN", txt)
		}
	}
}
