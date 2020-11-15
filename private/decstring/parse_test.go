package decstring

import (
	"strings"
	"testing"
)

func combineInput(s string) []string {
	ss := []string{
		s,
		" " + s,
		" " + s + " ",
		s + " ",
		strings.ToUpper(s),
	}
	if strings.Contains(s, "+") {
		ss = append(ss, strings.Replace(s, "+", " +", 1))
		ss = append(ss, strings.Replace(s, "+", " + ", 1))
		ss = append(ss, strings.Replace(s, "+", "+ ", 1))
	}
	if strings.Contains(s, "-") {
		ss = append(ss, strings.Replace(s, "-", " -", 1))
		ss = append(ss, strings.Replace(s, "-", " - ", 1))
		ss = append(ss, strings.Replace(s, "-", "- ", 1))
	}
	return ss
}

func TestPositiveInf(t *testing.T) {
	testruns := []string{"inf", "+inf", "infinity", "+infinity"}
	for _, tinput := range testruns {
		for _, txt := range combineInput(tinput) {
			d := Parse(txt)
			if !d.IsInf() {
				t.Errorf("Parse(%s): expected=Inf, got=%v", txt, d)
			}
			if d.DecParts() != nil {
				t.Errorf("Parse(%s): expected nil DecParts, got=%v", txt, d.DecParts())
			}
		}
	}
}

func TestNegativeInf(t *testing.T) {
	testruns := []string{"-inf", "-infinity"}
	for _, tinput := range testruns {
		for _, txt := range combineInput(tinput) {
			d := Parse(txt)
			if !d.IsMinusInf() {
				t.Errorf("Parse(%s): expected=-Inf, got=%v", txt, d)
			}
			if d.DecParts() != nil {
				t.Errorf("Parse(%s): expected nil DecParts, got=%v", txt, d.DecParts())
			}
		}
	}
}

func TestPositiveNaN(t *testing.T) {
	testruns := []string{"nan", "+nan", "nan 123"}
	for _, tinput := range testruns {
		for _, txt := range combineInput(tinput) {
			d := Parse(txt)
			if !d.IsNaN() {
				t.Errorf("Parse(%s): expected=NaN, got=%v", txt, d)
			}
			if d.DecParts() != nil {
				t.Errorf("Parse(%s): expected nil DecParts, got=%v", txt, d.DecParts())
			}
		}
	}
}

func TestNegativeNaN(t *testing.T) {
	testruns := []string{"-NaN", "-nan"}
	for _, tinput := range testruns {
		for _, txt := range combineInput(tinput) {
			d := Parse(txt)
			if !d.IsMinusNaN() {
				t.Errorf("Parse(%s): expected=-NaN, got=%v", txt, d)
			}
			if d.DecParts() != nil {
				t.Errorf("Parse(%s): expected nil DecParts, got=%v", txt, d.DecParts())
			}
		}
	}
}

func TestPositiveSNaN(t *testing.T) {
	testruns := []string{"abc", "snan", "+snan", "snan 123"}
	for _, tinput := range testruns {
		for _, txt := range combineInput(tinput) {
			d := Parse(txt)
			if !d.IsSNaN() {
				t.Errorf("Parse(%s): expected=SNaN, got=%v", txt, d)
			}
			if d.DecParts() != nil {
				t.Errorf("Parse(%s): expected nil DecParts, got=%v", txt, d.DecParts())
			}
		}
	}
}

func TestNegativeSNaN(t *testing.T) {
	testruns := []string{"-snan", "-snan 123"}
	for _, tinput := range testruns {
		for _, txt := range combineInput(tinput) {
			d := Parse(txt)
			if !d.IsMinusSNaN() {
				t.Errorf("Parse(%s): expected=-SNaN, got=%v", txt, d)
			}
			if d.DecParts() != nil {
				t.Errorf("Parse(%s): expected nil DecParts, got=%v", txt, d.DecParts())
			}
		}
	}
}

func TestOne(t *testing.T) {
	testruns := []string{"1", "+1", "1.0", "+1.0", "1.", "+1."}
	for _, tinput := range testruns {
		for _, txt := range combineInput(tinput) {
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
}
