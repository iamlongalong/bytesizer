package bytesizer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByteSizeString(t *testing.T) {
	tests := []struct {
		name string
		size ByteSize
		want string
	}{
		{"Byte", 1, "1B"},
		{"Kilobyte", KB, "1KB"},
		{"Megabyte", MB, "1MB"},
		{"Gigabyte", GB, "1GB"},
		{"Terabyte", TB, "1TB"},
		{"Petabyte", PB, "1PB"},
		{"Multiple Bytes", 532, "532B"},
		{"Multiple Kilobytes", 1025 * KB, "1.00MB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.size.String())
		})
	}
}

// Test format function with different units
func TestByteSizeFormat(t *testing.T) {
	tests := []struct {
		name     string
		size     ByteSize
		unit     ByteSize
		expected string
	}{
		{"Format as KB", KB, KB, "1KB"},
		{"Format as MB", 5 * MB, MB, "5MB"},
		{"Format as GB", 2 * GB, GB, "2GB"},
		{"Format as TB", 3 * TB, TB, "3TB"},
		{"Format as PB", 4 * PB, PB, "4PB"},
		{"Format with no unit set", 512 * MB, 0, "512MB"},                          // Default to MB if no unit is set
		{"Format as B", 500, Byte, "500B"},                                         // Test with Byte unit
		{"Format as KB with decimal", ByteSize(1.5 * float64(KB)), KB, "1.5KB"},    // Test KB with decimal
		{"Format as MB with decimal", ByteSize(2.75 * float64(MB)), MB, "2.75MB"},  // Test MB with decimal
		{"Format as GB with decimal", ByteSize(3.125 * float64(GB)), GB, "3.13GB"}, // Test GB with decimal
		{"Format as TB with decimal", ByteSize(3.875 * float64(TB)), TB, "3.88TB"}, // Test TB with decimal
		{"Format as PB with decimal", ByteSize(4.5 * float64(PB)), PB, "4.5PB"},    // Test PB with decimal
	}

	for _, tt := range tests {
		if tt.name == "Format as TB" {
			fmt.Println("")
		}
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.size.Format(tt.unit))
		})
	}
}

// Test conversion functions to various units
func TestConversions(t *testing.T) {
	const byteValue ByteSize = 1024
	tests := []struct {
		name     string
		size     ByteSize
		unit     ByteSize
		expected string
	}{
		{"To Bytes", byteValue, Byte, "1024B"},
		{"To KB", byteValue, KB, "1KB"},
		{"To MB", byteValue * KB, MB, "1MB"},
		{"To GB", byteValue * MB, GB, "1GB"},
		{"To TB", byteValue * GB, TB, "1TB"},
		{"To PB", byteValue * TB, PB, "1PB"},

		{"To Half KB", byteValue / 2, KB, "0.5KB"},
		{"To Half MB", (byteValue * KB) / 2, MB, "0.5MB"},
		{"To Half GB", (byteValue * MB) / 2, GB, "0.5GB"},
		{"To Half TB", (byteValue * GB) / 2, TB, "0.5TB"},
		{"To Half PB", (byteValue * TB) / 2, PB, "0.5PB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tval := tt.size.Format(tt.unit)
			assert.Equal(t, tt.expected, tval)
		})
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		name      string
		sizeStr   string
		expectErr bool
		expected  ByteSize
	}{
		{"Valid Parse B", "1024B", false, 1024},
		{"Valid Parse KB", "1KB", false, KB},
		{"Valid Parse MB", "1MB", false, MB},
		{"Valid Parse GB", "1GB", false, GB},
		{"Valid Parse TB", "1TB", false, TB},
		{"Valid Parse PB", "1PB", false, PB},
		{"Invalid Unit", "1XB", true, 0},
		{"Invalid Format", "OneKB", true, 0},
		{"Empty String", "", true, 0},

		{"Valid Parse B with decimal", "1024.5B", false, 1024},
		{"Valid Parse KB with decimal", "1.5KB", false, ByteSize(1.5 * float64(KB))},
		{"Valid Parse MB with decimal", "1.5MB", false, ByteSize(1.5 * float64(MB))},
		{"Valid Parse GB with decimal", "1.5GB", false, ByteSize(1.5 * float64(GB))},
		{"Valid Parse TB with decimal", "1.5TB", false, ByteSize(1.5 * float64(TB))},
		{"Valid Parse PB with decimal", "1.5PB", false, ByteSize(1.5 * float64(PB))},

		// Invalid case with floating point value
		{"Invalid Format with MA", "5MA", true, 0},
		{"Invalid Format with float", "One.5KB", true, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			size, err := Parse(tt.sizeStr)

			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, size)
			}
		})
	}
}

func TestFormatString(t *testing.T) {
	tests := []struct {
		name      string
		inputV    float64
		inputUnit string
		inputDec  []int
		expected  string
	}{
		{
			name:      "No decimal",
			inputV:    1.00,
			inputUnit: "MB",
			inputDec:  nil,
			expected:  "1MB",
		},
		{
			name:      "Decimal length 2",
			inputV:    1.011,
			inputUnit: "MB",
			inputDec:  []int{2},
			expected:  "1.01MB",
		},
		{
			name:      "Decimal length 3 and equal",
			inputV:    1.001,
			inputUnit: "MB",
			inputDec:  []int{3},
			expected:  "1.001MB",
		},
		{
			name:      "Decimal length 3 not equal",
			inputV:    1.0011,
			inputUnit: "MB",
			inputDec:  []int{3},
			expected:  "1.001MB",
		},
		{
			name:      "Decimal length negative",
			inputV:    1.0011,
			inputUnit: "MB",
			inputDec:  []int{-1},
			expected:  "1.0011MB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := formatString(tt.inputV, tt.inputUnit, tt.inputDec...)
			if output != tt.expected {
				t.Errorf("formatString(%f, %s, %v) = %s; want %s", tt.inputV, tt.inputUnit, tt.inputDec, output, tt.expected)
			}
		})
	}
}

func TestByteSizeMethods(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		fs                                         ByteSize
		byteInt, kbInt, mbInt, gbInt, tbInt, pbInt int
	}{
		{2048 * MB, 2048 * int(MB), 2048 * int(MB) / int(KB), 2048, int(2), 0, 0},
	}

	for _, test := range tests {
		assert.Equal(test.byteInt, test.fs.ByteInt(), "They should be equal")
		assert.Equal(test.kbInt, test.fs.KBInt(), "They should be equal")
		assert.Equal(test.mbInt, test.fs.MBInt(), "They should be equal")
		assert.Equal(test.gbInt, test.fs.GBInt(), "They should be equal")
		assert.Equal(test.tbInt, test.fs.TBInt(), "They should be equal")
		assert.Equal(test.pbInt, test.fs.PBInt(), "They should be equal")
	}
}
