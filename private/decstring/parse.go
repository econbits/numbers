package decstring

import (
	"regexp"
)

// Parses a Decimal String, following this pattern:
//   sign ::= '+' | '-'
//   digit ::= '0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9'
//   indicator ::=  'e' | 'E'
//   digits ::= digit [digit]...
//   decimal-part ::= digits '.' [digits] | ['.'] digits
//   exponent-part ::= indicator [sign] digits
//   infinity ::= 'Infinity' | 'Inf'
//   nan ::= 'NaN' [digits] | 'sNaN' [digits]
//   numeric-value ::= decimal-part [exponent-part] | infinity
//   numeric-string ::= [sign] numeric-value | [sign] nan
// Caveats:
//
// - NaN and sNaN digits are not taken into account, meaning that NaN 111 is equivalent to NaN
func Parse(ds string) DecString {
	d := parseSpecialValue(ds)
	if d.sv != 0 {
		return d
	}
	d = parseNumeric(ds)
	if d.dp != nil {
		return d
	}
	d.sv = specialValueSNaN
	return d
}

// sign ::= '+' | '-'
// digit ::= '0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9'
// indicator ::=  'e' | 'E'
// digits ::= digit [digit]...
// decimal-part ::= digits '.' [digits] | ['.'] digits
// exponent-part ::= indicator [sign] digits
// numeric-value ::= decimal-part [exponent-part]
// ds ::= [sign] numeric-value
var reNumeric = regexp.MustCompile(`(?i)^(?P<decsign>[+\-]{0,1})(?P<coeff>[0-9]*\.{0,1}[0-9]*)(e(?P<expsign>[+\-]{0,1})(?P<exp>[0-9]+)){0,1}$`)

func parseNumeric(ds string) DecString {
	dp := DecParts{positive: true, coeff: "", exp: DecExp{positive: true, value: ""}}

	groupNames := reNumeric.SubexpNames()
	matches := reNumeric.FindStringSubmatch(ds)

	for i, value := range matches {
		if i == 0 {
			// this is the whole str
			continue
		}
		if groupNames[i] == "decsign" {
			dp.positive = (value == "+")
		} else if groupNames[i] == "coeff" {
			dp.coeff = value
		} else if groupNames[i] == "expsign" {
			dp.exp.positive = (value == "+")
		} else if groupNames[i] == "exp" {
			dp.exp.value = value
		}
	}
	d := DecString{sv: 0, dp: nil}
	if dp.coeff != "" {
		d.dp = &dp
	}
	return d
}

// sign ::= '+' | '-'
// infinity ::= 'Infinity' | 'Inf'
// ds ::= [sign] infinity
var reInf = regexp.MustCompile(`(?i)^\s*\+{0,1}\s*Inf(inity){0,1}\s*$`)
var reMinusInf = regexp.MustCompile(`(?i)^\s*\-\s*Inf(inity){0,1}\s*$`)

// sign ::= '+' | '-'
// nan ::= 'NaN' [digits] | 'sNaN' [digits]
// ds ::= [sign] nan
var reNaN = regexp.MustCompile(`(?i)^\s*\+{0,1}\s*nan(\s+[0-9]+){0,1}\s*$`)
var reMinusNaN = regexp.MustCompile(`(?i)^\-nan(\s+[0-9]+){0,1}$`)
var reSNaN = regexp.MustCompile(`(?i)^\+{0,1}snan\s+[0-9]*$`)
var reMinusSNaN = regexp.MustCompile(`(?i)^\-snan\s+[0-9]*$`)

var specialValues = []struct {
	re *regexp.Regexp
	sv specialValue
}{
	{re: reInf, sv: specialValueInf},
	{re: reMinusInf, sv: specialValueMinusInf},
	{re: reNaN, sv: specialValueNaN},
	{re: reMinusNaN, sv: specialValueMinusNaN},
	{re: reSNaN, sv: specialValueSNaN},
	{re: reMinusSNaN, sv: specialValueMinusSNaN},
}

func parseSpecialValue(ds string) DecString {
	d := DecString{sv: 0, dp: nil}

	for _, resv := range specialValues {
		if resv.re.MatchString(ds) {
			d.sv = resv.sv
			break
		}
	}
	return d
}
