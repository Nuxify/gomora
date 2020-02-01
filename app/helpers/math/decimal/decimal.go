package decimal

import "gopkg.in/inf.v0"

// AddDecimals - add two decimal inputs with decimal result
func AddDecimals(first, second *inf.Dec) *inf.Dec {
	d := new(inf.Dec)

	sum := d.Add(first, second)
	return sum
}

// DivRoundDownDecimals - divide decimals using round down
func DivRoundDownDecimals(first, second *inf.Dec) *inf.Dec {
	d := new(inf.Dec)

	quotient := d.QuoRound(first, second, 13, inf.RoundDown)
	return quotient
}

// MulDecimals - multiple decimals
func MulDecimals(first, second *inf.Dec) *inf.Dec {
	d := new(inf.Dec)

	prod := d.Mul(first, second)
	return prod
}

// SubDecimals - subtract two decimal numbers
func SubDecimals(first, second *inf.Dec) *inf.Dec {
	d := new(inf.Dec)

	diff := d.Sub(first, second)
	return diff
}

// ToDecimal - converts string to *inf.Dec data tyoe
func ToDecimal(s string) *inf.Dec {
	d := new(inf.Dec)

	res, _ := d.SetString(s)
	return res
}
