This is a temporary fork for adding support for parsing XML nil values `xsi:nil="true"` inside elements with the `xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"` attribute.

Original issue: https://github.com/danilobuerger/null/issues/1

[![Build Status](https://travis-ci.org/danilobuerger/null.svg?branch=master)](https://travis-ci.org/danilobuerger/null) [![Coverage Status](https://coveralls.io/repos/github/danilobuerger/null/badge.svg?branch=master)](https://coveralls.io/github/danilobuerger/null?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/danilobuerger/null)](https://goreportcard.com/report/github.com/danilobuerger/null)

# null

## Types

Nullable types for Go. All types implement:

* `sql.Scanner`
* `driver.Valuer`
* `json.Marshaler`
* `json.Unmarshaler`
* `encoding.TextMarshaler`
* `encoding.TextUnmarshaler`
* `xml.Marshaler`
* `xml.Unmarshaler`

When marshaling xml, the attribute `xsi:nil="true"` will be added for null values.
When unmarshaling xml, the value will be null if `xsi:nil="true"` is set.

### Float64

`null.Float64` is based on `sql.NullFloat64` and exports the fields `Float64` and `Valid`.
`Valid` is true if `Float64` is not null.

### Int64

`null.Int64` is based on `sql.NullInt64` and exports the fields `Int64` and `Valid`.
`Valid` is true if `Int64` is not null.

### String

`null.String` is based on `sql.NullString` and exports the fields `String` and `Valid`.
`Valid` is true if `String` is not null.

When using `encoding.TextUnmarshaler`, an empty text will be valid.
