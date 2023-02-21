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

func validateFraction(f *big.Rat, b, p, q int) error {
	// Validate numerator.
	err := validateNumerator(f, b, p, q)
	if err != nil {
		return err
	}
	// Validate denominator.
	err = validateDenominator(f, b, p)
	if err != nil {
		return err
	}
	return nil
}

func validateNumerator(f *big.Rat, b, p, q int) error {
	// Variable conversion.
	bFloat := float64(b)
	// (equation 1) eq1 = b - 1 / 2
	eq1 := (bFloat - 1.0) / 2.0
	// (equation 2) eq2 = b^(q - p + 1) / b - 1
	e := float64(q-p) + 1.0
	eq2 := math.Pow(bFloat, e) / (bFloat - 1.0)
	// (round up) ru = ceil(eq1)
	ru := math.Ceil(eq1)
	// (round down) rd = floor(eq1)
	rd := math.Floor(eq2)
	// (Lower bound) lb = -1 * ru * eq2
	lb := int64(-ru * eq2)
	// (Upper bound) ub = rd * eq2
	ub := int64(rd * eq2)
	// Check if numerator is (lb <= numerator <= ub).
	// Therefore, we can check if numerator is (numerator < lb or ub < numerator).
	n := f.Num()
	numeratorIsLessThanLowerBound := n.Cmp(big.NewInt(lb)) == -1
	numeratorIsGreaterThanUpperBound := n.Cmp(big.NewInt(ub)) == 1
	if numeratorIsLessThanLowerBound || numeratorIsGreaterThanUpperBound {
		return ErrNumeratorIsNotInTheMessageSpaceRange
	}
	return nil
}

func validateDenominator(f *big.Rat, b, p int) error {
	d := f.Denom()
	absP := math.Abs(float64(p))
	bPowP := math.Pow(float64(b), absP)
	// if d.Cmp(big.NewInt(int64(bPowP))) == 0 then denominator == b^(|p|).
	if d.Cmp(big.NewInt(int64(bPowP))) != 0 {
		return ErrDenominatorIsNotEqualToBToThePowerOfP
	}
	return nil
}

func validateEncodingParameters(f *big.Rat, b, p, q, d int) error {
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
	err = validateFraction(f, b, p, q)
	if err != nil {
		return err
	}
	// Validate denominator.
	err = validateDenominator(f, b, p)
	if err != nil {
		return err
	}
	// Validate numerator.
	err = validateNumerator(f, b, p, q)
	if err != nil {
		return err
	}
	return nil
}
