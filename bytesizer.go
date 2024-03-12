package bytesizer

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type ByteSize int

const (
	Byte ByteSize = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
)

var units = []struct {
	size     ByteSize
	unitName string
}{
	{Byte, "B"}, {KB, "KB"}, {MB, "MB"}, {GB, "GB"}, {TB, "TB"}, {PB, "PB"},
}

// Calc calc the []byte length, trans to ByteSize.
func Calc(b []byte) ByteSize {
	return ByteSize(len(b))
}

// Format method formats the ByteSize value to a string based on the given byte unit.
// It iterates over the predefined units until it matches the given unit,
// then calls formatString to generate the final formatted string.
// If the unit doesn't match any predefined units, it returns the string representation of the ByteSize itself.
func (fs ByteSize) Format(bu ByteSize) string {

	unitVal := float64(fs) / float64(bu)

	for _, u := range units {
		if u.size == bu {
			return formatString(unitVal, u.unitName, 2)
		}
	}

	return fs.String()
}

// String method converts ByteSize to a string with an appropriate unit.
func (fs ByteSize) String() string {
	switch {
	case fs >= PB:
		return formatString(fs.PB(), "PB", 2)
	case fs >= TB:
		return formatString(fs.TB(), "TB", 2)
	case fs >= GB:
		return formatString(fs.GB(), "GB", 2)
	case fs >= MB:
		return formatString(fs.MB(), "MB", 2)
	case fs >= KB:
		return formatString(fs.KB(), "KB", 2)
	}

	return formatString(fs.Byte(), "B", 2)
}

// Byte method returns the ByteSize in bytes as a float64.
func (fs ByteSize) Byte() float64 {
	return float64(fs)
}

// KB method returns the ByteSize in kilobytes as a float64.
func (fs ByteSize) KB() float64 {
	return float64(fs) / float64(KB)
}

// MB method returns the ByteSize in megabytes as a float64.
func (fs ByteSize) MB() float64 {
	return float64(fs) / float64(MB)
}

// GB method returns the ByteSize in gigabytes as a float64.
func (fs ByteSize) GB() float64 {
	return float64(fs) / float64(GB)
}

// TB method returns the ByteSize in terabytes as a float64.
func (fs ByteSize) TB() float64 {
	return float64(fs) / float64(TB)
}

// PB method returns the ByteSize in petabytes as a float64.
func (fs ByteSize) PB() float64 {
	return float64(fs) / float64(PB)
}

// ByteInt method returns the ByteSize in bytes as an integer.
func (fs ByteSize) ByteInt() int {
	return int(fs)
}

// KBInt method returns the ByteSize in kilobytes as an integer.
func (fs ByteSize) KBInt() int {
	return int(fs / KB)
}

// MBInt method returns the ByteSize in megabytes as an integer.
func (fs ByteSize) MBInt() int {
	return int(fs / MB)
}

// GBInt method returns the ByteSize in gigabytes as an integer.
func (fs ByteSize) GBInt() int {
	return int(fs / GB)
}

// TBInt method returns the ByteSize in terabytes as an integer.
func (fs ByteSize) TBInt() int {
	return int(fs / TB)
}

// PBInt method returns the ByteSize in petabytes as an integer.
func (fs ByteSize) PBInt() int {
	return int(fs / TB)
}

// parse a string s in bytes, kilobytes, megabytes, gigabytes,
// terabytes or petabytes format and converts it into ByteSize, a datatype representing byte sizes.
// accepts a string s like "10B", "10KB", "10MB", "10GB", "10TB", "10PB" and returns the corresponding ByteSize.
// returns an error if the format of s is invalid or if an invalid size unit is found.
//
// Example usage:
//
//	size, err := Parse("10KB")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(size)
//
// Output: 10240 // Bytes equivalent of 10KB
func Parse(s string) (ByteSize, error) {
	if len(s) == 0 {
		return 0, fmt.Errorf("empty size string")
	}

	units := map[string]ByteSize{
		"B":  Byte,
		"KB": KB,
		"MB": MB,
		"GB": GB,
		"TB": TB,
		"PB": PB,
	}

	var unitName string
	var valueStr string

	if len(s) > 2 && strings.Contains("KMGTP", s[len(s)-2:len(s)-1]) {
		unitName = s[len(s)-2:]
		valueStr = s[:len(s)-2]
	} else {
		unitName = s[len(s)-1:]
		valueStr = s[:len(s)-1]
	}

	unit, exists := units[strings.ToUpper(unitName)]
	if !exists {
		return 0, fmt.Errorf("invalid size unit: %v", unitName)
	}

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return 0, err
	}

	return ByteSize(value * float64(unit)), nil
}

// formatString. format value in a proper way
// v: value, unit: unit string, decimalCount optional decimal count (default 2)
// example formatString(1.00, "MB") => 1MB (if equal int, then without decimal part)
// example formatString(1.011, "MB") => 1.01MB
// example formatString(1.001, "MB") => 1.00MB
func formatString(v float64, unit string, maxDecimalCount ...int) string {
	decimals := decimalPlaces(v)
	if len(maxDecimalCount) > 0 && maxDecimalCount[0] >= 0 {
		if decimals > maxDecimalCount[0] {
			decimals = maxDecimalCount[0]
		}
	}

	// rounding
	multiper := math.Pow(10, float64(decimals))
	n := math.Round(v*multiper) / multiper

	formatString := fmt.Sprintf("%%.%df%%s", decimals)
	return fmt.Sprintf(formatString, n, unit)
}

// decimalPlaces counts the decimal places in a float64.
// E.g1: decimalPlaces(1.23) returns 2
// E.g2: decimalPlaces(100.456) returns 3.
// E.g3: decimalPlaces(10.100) returns 1.
func decimalPlaces(f float64) (n int) {
	const epsilon = 1e-10
	for math.Abs(f-math.Floor(f)) > epsilon {
		f *= 10
		n++
	}
	return n
}
