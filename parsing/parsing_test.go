package parsing_test

import (
	"math/big"
	"testing"

	"github.com/smartcontractkit/crec-sdk-ext-dta/parsing"
	"github.com/stretchr/testify/require"
)

func TestScientificNotationToBigInt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string // string representation of expected big.Int
		wantOk   bool
	}{
		// Regular decimal numbers
		{
			name:     "Simple integer",
			input:    "123456789",
			expected: "123456789",
			wantOk:   true,
		},
		{
			name:     "Zero",
			input:    "0",
			expected: "0",
			wantOk:   true,
		},
		{
			name:     "Large integer",
			input:    "123456789012345678901234567890",
			expected: "123456789012345678901234567890",
			wantOk:   true,
		},

		// Scientific notation - positive exponents
		{
			name:     "Simple scientific notation",
			input:    "1e18",
			expected: "1000000000000000000",
			wantOk:   true,
		},
		{
			name:     "Scientific notation with decimal",
			input:    "1.2e+21",
			expected: "1200000000000000000000",
			wantOk:   true,
		},
		{
			name:     "Scientific notation uppercase E",
			input:    "5E20",
			expected: "500000000000000000000",
			wantOk:   true,
		},
		{
			name:     "Scientific notation with plus sign",
			input:    "3.14e+5",
			expected: "314000",
			wantOk:   true,
		},
		{
			name:     "Large scientific notation",
			input:    "1.23456789e+30",
			expected: "1234567890000000000000000000000",
			wantOk:   true,
		},

		// Scientific notation - negative exponents (should truncate to integer)
		{
			name:     "Scientific notation negative exponent small",
			input:    "1.5e-1",
			expected: "0", // truncated to integer
			wantOk:   true,
		},
		{
			name:     "Scientific notation negative exponent zero result",
			input:    "5e-10",
			expected: "0", // truncated to integer
			wantOk:   true,
		},
		{
			name:     "Scientific notation that becomes integer",
			input:    "1.23e2",
			expected: "123",
			wantOk:   true,
		},

		// Edge cases
		{
			name:     "Very large scientific notation",
			input:    "1e100",
			expected: "10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			wantOk:   true,
		},
		{
			name:     "Scientific notation with many decimal places",
			input:    "1.23456789123456789e+20",
			expected: "123456789123456789000",
			wantOk:   true,
		},

		// Decimal numbers without scientific notation
		{
			name:     "Decimal number with zeros",
			input:    "600000000000000000000.000000",
			expected: "600000000000000000000",
			wantOk:   true,
		},
		{
			name:     "Simple decimal number with non-zero fractional part",
			input:    "123.456",
			expected: "",
			wantOk:   false, // Should fail because fractional part is not all zeros
		},
		{
			name:     "Decimal number with integer zero but non-zero fractional part",
			input:    "0.123456",
			expected: "",
			wantOk:   false, // Should fail because fractional part is not all zeros
		},
		{
			name:     "Large decimal number with non-zero fractional part",
			input:    "999999999999999999999.999999999",
			expected: "",
			wantOk:   false, // Should fail because fractional part is not all zeros
		},
		{
			name:     "Decimal with single digit",
			input:    "5.0",
			expected: "5",
			wantOk:   true,
		},

		// Invalid inputs
		{
			name:     "Invalid format",
			input:    "not-a-number",
			expected: "",
			wantOk:   false,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
			wantOk:   false,
		},
		{
			name:     "Invalid scientific notation",
			input:    "1e",
			expected: "",
			wantOk:   false,
		},
		{
			name:     "Multiple decimal points",
			input:    "1.2.3e10",
			expected: "",
			wantOk:   false,
		},
		{
			name:     "Invalid exponent",
			input:    "1eabc",
			expected: "",
			wantOk:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, ok := parsing.ScientificNotationToBigInt(tt.input)

			require.Equal(t, tt.wantOk, ok, "ScientificNotationToBigInt(%q) ok = %v, want %v", tt.input, ok, tt.wantOk)

			if tt.wantOk && ok {
				expected := new(big.Int)
				expected.SetString(tt.expected, 10)
				require.Equal(t, expected, result, "ScientificNotationToBigInt(%q) = %v, want %v", tt.input, result, expected)
			}
		})
	}
}

func TestScientificNotationToUint64(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    uint64
		expectError bool
	}{
		// Regular decimal numbers
		{
			name:        "Simple integer",
			input:       "123456",
			expected:    123456,
			expectError: false,
		},
		{
			name:        "Zero",
			input:       "0",
			expected:    0,
			expectError: false,
		},
		{
			name:        "Max uint64",
			input:       "18446744073709551615",
			expected:    18446744073709551615,
			expectError: false,
		},

		// Scientific notation - positive exponents
		{
			name:        "Simple scientific notation",
			input:       "1e6",
			expected:    1000000,
			expectError: false,
		},
		{
			name:        "Scientific notation with decimal",
			input:       "1.5e+3",
			expected:    1500,
			expectError: false,
		},
		{
			name:        "Scientific notation uppercase E",
			input:       "5E4",
			expected:    50000,
			expectError: false,
		},
		{
			name:        "Large scientific notation",
			input:       "1.23e+15",
			expected:    1230000000000000,
			expectError: false,
		},

		// Scientific notation - negative exponents (truncated to integer)
		{
			name:        "Scientific notation negative exponent small",
			input:       "1.5e-1",
			expected:    0, // truncated to integer
			expectError: false,
		},
		{
			name:        "Scientific notation that becomes integer",
			input:       "1.23e2",
			expected:    123,
			expectError: false,
		},

		// Error cases
		{
			name:        "Invalid format",
			input:       "not-a-number",
			expected:    0,
			expectError: true,
		},
		{
			name:        "Empty string",
			input:       "",
			expected:    0,
			expectError: true,
		},
		{
			name:        "Negative number",
			input:       "-123",
			expected:    0,
			expectError: true,
		},
		{
			name:        "Value too large for uint64",
			input:       "1e100",
			expected:    0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parsing.ScientificNotationToUint64(tt.input)

			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestScientificNotationToUint8(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    uint8
		expectError bool
	}{
		// Regular decimal numbers
		{
			name:        "Simple integer",
			input:       "123",
			expected:    123,
			expectError: false,
		},
		{
			name:        "Zero",
			input:       "0",
			expected:    0,
			expectError: false,
		},
		{
			name:        "Max uint8",
			input:       "255",
			expected:    255,
			expectError: false,
		},

		// Scientific notation - positive exponents
		{
			name:        "Simple scientific notation",
			input:       "1e2",
			expected:    100,
			expectError: false,
		},
		{
			name:        "Scientific notation with decimal",
			input:       "2.5e+1",
			expected:    25,
			expectError: false,
		},
		{
			name:        "Scientific notation uppercase E",
			input:       "1E1",
			expected:    10,
			expectError: false,
		},

		// Scientific notation - negative exponents (truncated to integer)
		{
			name:        "Scientific notation negative exponent small",
			input:       "1.5e-1",
			expected:    0, // truncated to integer
			expectError: false,
		},
		{
			name:        "Scientific notation that becomes integer",
			input:       "1.23e1",
			expected:    12, // truncated
			expectError: false,
		},

		// Error cases
		{
			name:        "Invalid format",
			input:       "not-a-number",
			expected:    0,
			expectError: true,
		},
		{
			name:        "Empty string",
			input:       "",
			expected:    0,
			expectError: true,
		},
		{
			name:        "Negative number",
			input:       "-1",
			expected:    0,
			expectError: true,
		},
		{
			name:        "Value too large for uint8",
			input:       "256",
			expected:    0,
			expectError: true,
		},
		{
			name:        "Scientific notation too large",
			input:       "1e3",
			expected:    0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parsing.ScientificNotationToUint8(tt.input)

			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, result)
			}
		})
	}
}

