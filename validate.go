package sim2d

import (
	"math"
	"math/big"
)

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

func validateNumerator(num *big.Int, b, p, q int) error {
	// Define a common component of all bounds: b^(q-p+1) - 1.
	// Exponent: b^(q-p+1).
	e := big.NewInt(int64(q - p + 1))
	// b^(q-p+1) - 1.
	commonEquation := big.NewInt(int64(b))
	commonEquation.Exp(commonEquation, e, nil)
	commonEquation.Sub(commonEquation, big.NewInt(1))
	// Define variables for lower and upper bound.
	lowerBound := big.NewRat(-1, 1)
	upperBound := big.NewRat(1, 1)
	// Define variables for first and second operands.
	firstOperand := big.NewRat(1, 1)
	secondOperand := big.NewRat(1, 1)
	// Check parity (if b is even or odd).
	// We define the lower and upper bounds by defining the equations in separated parts.
	// Lower bound is -b/2 x (b^(q-p+1) - 1) / (b-1).
	// Upper bound: (b/2 - 1/1) x ((b^(q-p+1) - 1) / 1)
	if b%2 == 0 {
		// Even:
		// First operand: b/2. Later we can just multiply by -1 or subtract by -1 to achieve the lower and upper bound.
		firstOperand.Mul(firstOperand, big.NewRat(int64(b), 2))
		// Second operand: ((b^(q-p+1) - 1) / (b-1)).
		secondOperand.Mul(secondOperand, big.NewRat(1, int64(b)-1))
		secondOperand.SetFrac(commonEquation, secondOperand.Denom())

		// Lower bound is -b/2 x (b^(q-p+1) - 1) / (b-1).
		// -b/2.
		lowerBound.Mul(lowerBound, firstOperand)
		// -b/2 x (b^(q-p+1) - 1) / (b-1).
		lowerBound.Mul(lowerBound, secondOperand)

		// Upper bound: (b/2 - 1/1) x ((b^(q-p+1) - 1) / 1)
		// b/2 - 1/1 = (b-2)/2.
		upperBound.Mul(upperBound, big.NewRat(int64(b-2), 2))
		// (b-2)/2 x ((b^(q-p+1) - 1) / 1).
		secondOperand.SetFrac(secondOperand.Num(), big.NewInt(1))
		upperBound.Mul(upperBound, secondOperand)
	} else {
		// Odd:
		// First operand: 1/2. Later we can just multiply the numerator by -1 to achieve the upper bound.
		firstOperand.Mul(firstOperand, big.NewRat(1, 2))
		// Second operand: (b^(q-p+1) - 1) / 1.
		secondOperand.SetFrac(commonEquation, secondOperand.Denom())

		// Lower bound: -1/2 x (b^(q-p+1)/1).
		// -1/2.
		lowerBound.Mul(lowerBound, firstOperand)
		// -1/2 x (b^(q-p+1)/1).
		lowerBound.Mul(lowerBound, secondOperand)

		// Upper bound: 1/2 x (b^(q-p+1)/1).
		// 1/2.
		upperBound.Mul(upperBound, firstOperand)
		// 1/2 x (b^(q-p+1)/1).
		upperBound.Mul(upperBound, secondOperand)
	}
	// Check if numerator is (lb <= numerator <= ub).
	// Therefore, we can check if numerator is (numerator < lb or ub < numerator).
	n := big.NewRat(1, 1)
	n.SetInt(num)
	numeratorIsLessThanLowerBound := n.Cmp(lowerBound) == -1
	numeratorIsGreaterThanUpperBound := n.Cmp(upperBound) == 1
	if numeratorIsLessThanLowerBound || numeratorIsGreaterThanUpperBound {
		return ErrNumeratorIsNotInTheMessageSpaceRange
	}
	return nil
}

func validateDenominator(den *big.Int, b, p int) error {
	// Check if denominator is zero.
	if den.Cmp(big.NewInt(0)) == 0 {
		return ErrDenominatorIsZero
	}
	absP := math.Abs(float64(p))
	bPowP := big.NewInt(int64(math.Pow(float64(b), absP)))
	// Original condition: denominator == b^(|p|).
	denominatorIsNotEqualToBToThePowerOfP := den.Cmp(bPowP) != 0
	if denominatorIsNotEqualToBToThePowerOfP {
		return ErrDenominatorIsNotEqualToBToThePowerOfP
	}
	return nil
}

func validateDegreeOfCode(code []int64, d int) error {
	// Code degree.
	cd := len(code)
	// Check if d is a power of 2.
	// Log base 2 of d.
	floatD := float64(cd)
	logD := math.Log2(floatD)
	// Integer part of log base 2 of d.
	intLog := math.Round(logD)
	// Code degree recalculated.
	cdr := math.Pow(2.0, intLog)
	// Check if code degree is a power of 2.
	if floatD != cdr {
		return ErrCodeDegreeIsNotAPowerOfTwo
	}
	// Check if code degree is the same as the given degree.
	if d != cd {
		return ErrCodeDegreeIsDifferentFromDegree
	}
	return nil
}

func validateEncodingParameters(num, den *big.Int, b, p, q, d int) error {
	// Validate fraction.
	err := validateFraction(num, den, b, p, q)
	if err != nil {
		return err
	}
	return nil
}

func validateDecodingParameters(code []int64, b, p, q, d int) error {
	// Validate if degree of code is a power of 2.
	err := validateDegreeOfCode(code, d)
	if err != nil {
		return err
	}
	return nil
}
