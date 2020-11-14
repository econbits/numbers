package decstring

// Exponent is composed of a sign (+, -) and a value representing an unsigned integer
type DecExp struct {
	positive bool
	value    string
}

func (de DecExp) IsPositive() bool {
	return de.positive
}

func (de DecExp) Value() string {
	return de.value
}

// The Decimal Parts are form by a sign (+, -), a coefficient (unsigned int) and an exponent
type DecParts struct {
	positive bool
	coeff    string
	exp      DecExp
}

func (dp DecParts) IsPositive() bool {
	return dp.positive
}

func (dp DecParts) Coeff() string {
	return dp.coeff
}

func (dp DecParts) Exp() DecExp {
	return dp.exp
}

// Special Values are NaN, -NaN, SNaN, -SNaN, Inf, -Inf
type specialValue byte

const (
	specialValueNaN       = 1
	specialValueMinusNaN  = 2
	specialValueSNaN      = 3
	specialValueMinusSNaN = 4
	specialValueInf       = 5
	specialValueMinusInf  = 6
)

// Parsed Decimal String, a.k.a. Numeric String
type DecString struct {
	sv specialValue
	dp *DecParts
}

func (ds DecString) DecParts() *DecParts {
	return ds.dp
}

func (ds DecString) IsNaN() bool {
	return ds.sv == specialValueNaN
}

func (ds DecString) IsMinusNaN() bool {
	return ds.sv == specialValueMinusNaN
}

func (ds DecString) IsSNaN() bool {
	return ds.sv == specialValueSNaN
}

func (ds DecString) IsMinusSNaN() bool {
	return ds.sv == specialValueMinusSNaN
}

func (ds DecString) IsInf() bool {
	return ds.sv == specialValueInf
}

func (ds DecString) IsMinusInf() bool {
	return ds.sv == specialValueMinusInf
}
