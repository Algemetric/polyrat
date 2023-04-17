package polyrat

// Encode encodes a rational number into a set of polynomial degrees.
// The function accepts as input a 64-bit rational number (float64) and bounds the precision by the lower power p.
// If a number exceeds the precision given by p, then such number will be truncaded.
func Encode(rat float64, params *Parameters) ([]int64, error) {
	// Transforms a rational number into an integer.
	n := parseRational(rat, params)
	// Input validation.
	if inputIsInvalid(n, params) {
		return nil, ErrNumeratorIsNotInTheMessageSpaceRange
	}
	// Calculate expansion.
	e := expansion(n, params)
	// Generate encoding (code).
	c := generateCode(e, params)
	// return code.
	return c, nil
}

func generateCode(exp []int64, params *Parameters) []int64 {
	// Polynomial length.
	pl := polynomialLength(params)
	// Initiate code with members of the expansion from -p to the end of the vector.
	code := exp[-params.MinPower():]
	// d - l gives the number of zeros to be concatenated in the vector.
	for i := 0; i < (params.Degree() - pl); i++ {
		code = append(code, 0)
	}
	// Concatenate to the code, from 0 to (-p) - 1, the elements of the expansion.
	for i := 0; i < -params.MinPower(); i++ {
		code = append(code, -int64(exp[i]))
	}
	// Return code.
	return code
}
