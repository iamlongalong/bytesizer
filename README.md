# bytesizer

`bytesizer` is a Go package that provides a simple and intuitive way to work with byte counts. It allows you to easily convert between different units of byte sizes, such as bytes, kilobytes, megabytes, gigabytes, terabytes, and petabytes. Similar to how `time.Duration` works in Go, `bytesizer` helps in representing byte sizes with appropriate types and methods.

## Features

- Conversion between byte units and bytes
- Formatting of byte sizes into human-readable strings
- Parsing of strings into byte sizes with support for different units
- Fetching byte sizes as integer or floating-point numbers for precision work

## Installation

To install `bytesizer`, you can use the following Go command:

```bash
go get github.com/iamlongalong/bytesizer
```

Replace `iamlongalong` with your actual GitHub username where the package is hosted.

## Usage

Below you'll find the methods provided by the `bytesizer` package and some examples of how to use them.

### Constants

`bytesizer` defines the following byte size units:

```go
const (
	Byte ByteSize = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
)
```

### Methods

#### Calc
Calculate the length of a byte slice and return it as a `ByteSize`:

```go
size := bytesizer.Calc([]byte("Hello World!"))
```

#### Format
Format a `ByteSize` value to a string according to a specified unit:

```go
formattedSize := size.Format(bytesizer.KB) // returns string like "1KB"
```

#### String
Convert a `ByteSize` to a string with an automatic unit:

```go
sizeString := size.String() // returns string like "11B"
```

#### Byte, KB, MB, GB, TB, PB
Get the byte size as different units (returns `float64`):

```go
bytes := size.Byte()
kilobytes := size.KB()
// ... and so on for MB, GB, TB, PB
```

#### ByteInt, KBInt, MBInt, GBInt, TBInt, PBInt
Get the byte size as different units (returns `int`):

```go
bytesInt := size.ByteInt()
kilobytesInt := size.KBInt()
// ... and so on for MBInt, GBInt, TBInt, PBInt
```

#### Parse
Parse a string representation of a byte size into a `ByteSize` object:

```go
size, err := bytesizer.Parse("10KB")
if err != nil {
    log.Fatal(err)
}
fmt.Println(size) // Output: 10240 (Bytes equivalent of 10KB)
```

## Contributing

Contributions to `bytesizer` are welcome! Feel free to report issues or submit pull requests on our GitHub repository.

## License

This package is licensed under the MIT License - see the LICENSE.md file for details.
