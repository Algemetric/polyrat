# Single Instruction Much More Data Codec

# Parameters

Both the encoding and decoding (i.e., codec functions) procedures will use a set of parameters, that need to be created beforehand. The set of parameters given to these functions is comprised by:

- Base (b);
- Higher power (q);
- Lower power (p);
- Degree (d).

Parameters will be created before using the encoding and/or decoding functions, because the instantiation of parameters will perform a check on the relationships between values. Errors will be returned if such values do not hold the proper relationships. To create a new set of parameters we use the function:

```golang
func NewParameters(p int, q int, d int) (*Parameters, error) {...}
```

This function receives `p`, `q`, and `d` and returns a reference to a structure containing the chosen parameters or an error showing why values were not defined correctly. The base `b` is already defined internally to be base `10`.

If we want to create a structure with `p = -4`, `q = 11`, and `d = 16`, then we instantiate the structure by issuing

```golang
params, err := sim2d.NewParameters(-4, 11, 16)
```

The error variable (i.e., `err`) must be checked for any returned error. If no error occurred, then the `params` variable can be passed to the encoding or decoding function.

# Encode

The `Encode` function encodes a rational number into a set of polynomial degrees. The function accepts as input a 64-bits rational number (float64) and bounds the precision by the lower power `p`. If a number exceeds the precision given by `p`, then such number will be truncaded. The function is defined as

```golang
func Encode(rat float64, params *Parameters) ([]int64, error) {...}
```

To encode a rational number `r`, such as `98123.45`, we first create the parameters' structure `params` and execute the function such that a code (i.e., encoded number) will be generated or an error will be returned:

```golang
c, err := sim2d.Encode(r, params)
```

The code `c` returned by the `Encode` function is a set of 64-bits integers. To make possible to decode the same original rational, the same set of parameters should be used.

# Decode

The `Decode` function will basically reverse the code generated by `Encode` into the original rational, as soon as the same set of parameters are used. The `Decode` definition is

```golang
func Decode(code []int64, params *Parameters) (float64, error) {...}
```

The returned value `r` is the original rational, along with an error variable `err` showing if the procedure was carried without errors.

```golang
r, err := sim2d.Decode(c, params)
```