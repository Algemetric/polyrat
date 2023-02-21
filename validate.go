package sim2d

import (
	"math"
	"math/big"
)

func validateModulo(b int) error {
	// b >= 2.
	if b < 2 {
		return ErrBIsLessThan2
	}
	return nil
}

func validateSmallestPowerOfExpansion(p, q int) error {
	// p < q.
	if p >= q {
		return ErrPIsLessThanQ
	}
	// p < 0.
	if p >= 0 {
		return ErrPIsGreaterThanOrEqualToZero
	}
	return nil
}

func validateGreatestPowerOfExpansion(q int) error {
	// q > 0.
	if q <= 0 {
		return ErrQIsLessThanOrEqualToZero
	}
	return nil
}

func validateDegree(p, q, d int) error {
	// d >= 1.
	if d < 1 {
		return ErrDIsLessThanOne
	}
	// d is a power of 2.
	// Log base 2 of d.
	floatD := float64(d)
	logD := math.Log2(floatD)
	// Integer part of log base 2 of d.
	intLog := math.Round(logD)
	// Recalculated d.
	d2 := math.Pow(2.0, intLog)
	if floatD != d2 {
		return ErrDIsNotAPowerOf2
	}
	// d > q + |p|.
	absP := math.Abs(float64(p))
	if d <= (q + int(absP)) {
		return ErrDIsLessThanOrEqualToQPlusP
	}
	return nil
}

func validateFraction(num, den *big.Int, b, p, q int) error {
	// Validate numerator.
	err := validateNumerator(num, b, p, q)
	if err != nil {
		return err
	}
	// Validate denominator.
	err = validateDenominator(den, b, p)
	if err != nil {
		return err
	}
	return nil
}

func validateNumerator(n *big.Int, b, p, q int) error {
	num1 := big.NewRat(1, 1)
	num1.SetInt(n)
	// (equation 1) eq1 = b - 1 / 2
	eq1 := big.NewRat(int64(b-1), 2)
	// (equation 2) eq2 = b^(q - p + 1) / b - 1
	e := big.NewInt(int64(q - p + 1))
	num := big.NewInt(int64(b))
	num.Exp(num, e, nil)
	den := big.NewInt(int64(b - 1))
	eq2 := big.NewRat(1, 1)
	eq2.SetFrac(num, den)
	// (round up) ru = ceil(eq1)
	// ru := math.Ceil(eq1)

	// (round down) rd = floor(eq1)
	// rd := math.Floor(eq2)

	// (Lower bound) lb = -1 * ru * eq2
	lb := big.NewRat(-1, 1)
	lb.Mul(lb, eq1)
	lb.Mul(lb, eq2)
	// (Upper bound) ub = rd * eq2
	ub := big.NewRat(1, 1)
	ub.Mul(ub, eq1)
	ub.Mul(ub, eq2)
	// Check if numerator is (lb <= numerator <= ub).
	// Therefore, we can check if numerator is (numerator < lb or ub < numerator).
	numeratorIsLessThanLowerBound := num1.Cmp(lb) == -1
	numeratorIsGreaterThanUpperBound := num1.Cmp(ub) == 1
	if numeratorIsLessThanLowerBound || numeratorIsGreaterThanUpperBound {
		return ErrNumeratorIsNotInTheMessageSpaceRange
	}
	return nil
}

func validateDenominator(d *big.Int, b, p int) error {
	absP := math.Abs(float64(p))
	bPowP := big.NewInt(int64(math.Pow(float64(b), absP)))
	// Original condition: denominator == b^(|p|).
	denominatorIsNotEqualToBToThePowerOfP := d.Cmp(bPowP) != 0
	if denominatorIsNotEqualToBToThePowerOfP {
		return ErrDenominatorIsNotEqualToBToThePowerOfP
	}
	return nil
}

func validateEncodingParameters(num, den *big.Int, b, p, q, d int) error {
	// Generate fraction.
	f := big.NewRat(1, 1)
	f.SetFrac(num, den)
	// Error variable.
	var err error
	// Validate modulo b.
	err = validateModulo(b)
	if err != nil {
		return err
	}
	// Validade smallest power of expansion.
	err = validateSmallestPowerOfExpansion(p, q)
	if err != nil {
		return err
	}
	err = validateGreatestPowerOfExpansion(q)
	if err != nil {
		return err
	}
	// Validate degree.
	err = validateDegree(p, q, d)
	if err != nil {
		return err
	}
	// Validate fraction.
	err = validateFraction(num, den, b, p, q)
	if err != nil {
		return err
	}
	return nil
}

func validateDegreeOfCode(code []int64) error {
	// Degree of code.
	d := len(code)
	// Check if d is a power of 2.
	// Log base 2 of d.
	floatD := float64(d)
	logD := math.Log2(floatD)
	// Integer part of log base 2 of d.
	intLog := math.Round(logD)
	// Recalculated d.
	d2 := math.Pow(2.0, intLog)
	if floatD != d2 {
		return ErrCodeDegreeIsNotAPowerOf2
	}
	return nil
}

func validateDecodingParameters(code []int64, b, p, q int) error {
	var err error
	// Validate modulo b.
	err = validateModulo(b)
	if err != nil {
		return err
	}
	// Validade smallest power of expansion.
	err = validateSmallestPowerOfExpansion(p, q)
	if err != nil {
		return err
	}
	err = validateGreatestPowerOfExpansion(q)
	if err != nil {
		return err
	}
	// Validate if degree of code is a power of 2.
	err = validateDegreeOfCode(code)
	if err != nil {
		return err
	}
	return nil
}
