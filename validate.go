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
	// Components for equations.
	// a = b^(q-p+1) - 1.
	e := big.NewInt(int64(q - p + 1))
	a := big.NewInt(int64(b))
	a.Exp(a, e, nil)
	a.Sub(a, big.NewInt(1))
	// Check parity (if b is even or odd).
	if b%2 == 0 {
		// Even:
		// First operand: b/2. Later we can just multiply by -1 or subtract by -1 to achieve the lower and upper bound.
		fo := big.NewRat(int64(b), 2)
		// Second operand: ((b^(q-p+1) - 1) / (b-1)).
		so := big.NewRat(1, int64(b)-1)
		so.SetFrac(a, so.Denom())

		// Lower bound is -b/2 x (b^(q-p+1) - 1) / (b-1).
		lb := big.NewRat(-1, 1)
		// -b/2.
		lb.Mul(lb, fo)
		// -b/2 x (b^(q-p+1) - 1) / (b-1).
		lb.Mul(lb, so)

		// Upper bound: (b/2 - 1/1) x ((b^(q-p+1) - 1) / 1)
		// b/2 - 1/1 = (b-2)/2.
		ub := big.NewRat(int64(b-2), 2)
		// (b-2)/2 x ((b^(q-p+1) - 1) / 1).
		so.SetFrac(so.Num(), big.NewInt(1))
		ub.Mul(ub, so)

		// Check if numerator is (lb <= numerator <= ub).
		// Therefore, we can check if numerator is (numerator < lb or ub < numerator).
		nf := big.NewRat(1, 1)
		nf.SetInt(n)
		numeratorIsLessThanLowerBound := nf.Cmp(lb) == -1
		numeratorIsGreaterThanUpperBound := nf.Cmp(ub) == 1
		if numeratorIsLessThanLowerBound || numeratorIsGreaterThanUpperBound {
			return ErrNumeratorIsNotInTheMessageSpaceRange
		}
	} else {
		// Odd:
		// Lower bound: (1 - b^(q-p+1)) / 2 x (b-1)
		// Upper bound: (b^(q-p+1) - 1) / 2 x (b-1)
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
