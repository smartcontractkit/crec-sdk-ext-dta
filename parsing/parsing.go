// Package parsing provides common parsing utilities for DTA event decoding.
// It handles scientific notation and decimal number parsing for blockchain event data.
package parsing

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

// ScientificNotationToBigInt converts scientific notation strings to big.Int.
// Handles formats like "1.2e+21", "1e18", etc. that big.Int.SetString cannot parse directly.
// Also handles decimal numbers like "600000000000000000000.000000".
// Returns nil and false if parsing fails.
func ScientificNotationToBigInt(value string) (*big.Int, bool) {
	// First try direct parsing in case it's already a regular integer
	if result, ok := new(big.Int).SetString(value, 10); ok {
		return result, true
	}

	// Handle scientific notation
	lowerValue := strings.ToLower(value)
	if strings.Contains(lowerValue, "e") {
		// Split on 'e' to get mantissa and exponent
		parts := strings.Split(lowerValue, "e")
		if len(parts) != 2 {
			return nil, false
		}

		mantissaStr := parts[0]
		exponentStr := parts[1]

		// Remove optional '+' from exponent
		exponentStr = strings.TrimPrefix(exponentStr, "+")

		// Parse exponent as integer
		exponent, err := strconv.Atoi(exponentStr)
		if err != nil {
			return nil, false
		}

		// Handle negative exponents (fractional results truncated to integer)
		if exponent < 0 {
			// For negative exponents, we need to check if the result would be < 1
			// If so, truncate to 0 (integer part)
			mantissaFloat, err := strconv.ParseFloat(mantissaStr, 64)
			if err != nil {
				return nil, false
			}

			// Calculate the actual value to see if it's < 1
			actualValue := mantissaFloat * pow10(exponent)
			if actualValue < 1.0 {
				return big.NewInt(0), true
			}

			// If >= 1, we need to handle it properly
			// Convert to string without scientific notation and truncate decimal part
			decimalStr := fmt.Sprintf("%.0f", actualValue)
			if result, ok := new(big.Int).SetString(decimalStr, 10); ok {
				return result, true
			}
			return nil, false
		}

		// For positive exponents, handle manually to avoid precision loss
		var mantissaBig *big.Int

		// Check if mantissa has decimal point
		if strings.Contains(mantissaStr, ".") {
			// Split mantissa into integer and fractional parts
			decimalParts := strings.Split(mantissaStr, ".")
			if len(decimalParts) != 2 {
				return nil, false
			}

			integerPart := decimalParts[0]
			fractionalPart := decimalParts[1]

			// Combine integer and fractional parts
			combinedStr := integerPart + fractionalPart

			// Parse as big integer
			var ok bool
			mantissaBig, ok = new(big.Int).SetString(combinedStr, 10)
			if !ok {
				return nil, false
			}

			// Adjust exponent to account for the fractional digits
			exponent -= len(fractionalPart)
		} else {
			// No decimal point, parse directly
			var ok bool
			mantissaBig, ok = new(big.Int).SetString(mantissaStr, 10)
			if !ok {
				return nil, false
			}
		}

		// Multiply by 10^exponent
		if exponent > 0 {
			// Multiply by 10^exponent
			multiplier := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(exponent)), nil)
			result := new(big.Int).Mul(mantissaBig, multiplier)
			return result, true
		} else if exponent < 0 {
			// Divide by 10^(-exponent) and truncate to integer
			divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(-exponent)), nil)
			result := new(big.Int).Div(mantissaBig, divisor)
			return result, true
		} else {
			// exponent == 0
			return mantissaBig, true
		}
	}

	// Handle decimal numbers without scientific notation (like "600000000000000000000.000000")
	if strings.Contains(value, ".") {
		// Split into integer and fractional parts
		decimalParts := strings.Split(value, ".")
		if len(decimalParts) != 2 {
			return nil, false
		}

		integerPart := decimalParts[0]
		fractionalPart := decimalParts[1]

		// Check if fractional part contains only zeros
		allZeros := true
		for _, digit := range fractionalPart {
			if digit != '0' {
				allZeros = false
				break
			}
		}

		// If fractional part is not all zeros, we cannot safely convert to big.Int
		if !allZeros {
			return nil, false
		}

		// Parse the integer part directly
		result, ok := new(big.Int).SetString(integerPart, 10)
		if !ok {
			return nil, false
		}

		// For integer conversion, we can safely truncate the all-zero decimal part
		return result, true
	}

	return nil, false
}

// ScientificNotationToUint64 converts scientific notation strings to uint64.
// Handles formats like "1.2e+21", "1e18", etc. that strconv.ParseUint cannot parse directly.
// Returns an error if the value cannot be parsed or is too large for uint64.
func ScientificNotationToUint64(value string) (uint64, error) {
	// First try direct parsing in case it's already a regular integer
	if result, err := strconv.ParseUint(value, 10, 64); err == nil {
		return result, nil
	}

	// Handle scientific notation using the big.Int parser and then convert
	bigIntResult, ok := ScientificNotationToBigInt(value)
	if !ok {
		return 0, fmt.Errorf("unable to parse scientific notation: %s", value)
	}

	// Check if the result fits in uint64
	if !bigIntResult.IsUint64() {
		return 0, fmt.Errorf("value too large for uint64: %s", value)
	}

	return bigIntResult.Uint64(), nil
}

// ScientificNotationToUint8 converts scientific notation strings to uint8.
// Handles formats like "1e2", "2.5e+1", etc. that strconv.ParseUint cannot parse directly.
// Returns an error if the value cannot be parsed or is out of range for uint8.
func ScientificNotationToUint8(value string) (uint8, error) {
	// First try direct parsing in case it's already a regular integer
	if result, err := strconv.ParseUint(value, 10, 8); err == nil {
		return uint8(result), nil
	}

	// Handle scientific notation using the big.Int parser and then convert
	bigIntResult, ok := ScientificNotationToBigInt(value)
	if !ok {
		return 0, fmt.Errorf("unable to parse scientific notation: %s", value)
	}

	// Check if the result fits in uint8 (0-255)
	if bigIntResult.Sign() < 0 || bigIntResult.Cmp(big.NewInt(255)) > 0 {
		return 0, fmt.Errorf("value out of range for uint8: %s", value)
	}

	return uint8(bigIntResult.Uint64()), nil
}

// pow10 calculates 10^exp for small exponents using float64 arithmetic.
func pow10(exp int) float64 {
	if exp == 0 {
		return 1.0
	}
	if exp > 0 {
		result := 1.0
		for i := 0; i < exp; i++ {
			result *= 10.0
		}
		return result
	} else {
		result := 1.0
		for i := 0; i < -exp; i++ {
			result /= 10.0
		}
		return result
	}
}

