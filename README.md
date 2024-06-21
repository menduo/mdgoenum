# mdgoenum

[![GoDoc](https://godoc.org/github.com/menduo/mdgoenum?status.svg)](https://godoc.org/github.com/menduo/mdgoenum)
[![Build Status](https://travis-ci.org/menduo/mdgoenum.svg?branch=master)](https://travis-ci.org/menduo/mdgoenum)

`mdgoenum` is a Go module designed to provide a flexible and type-safe enumeration mechanism for both string and integer
values. It allows developers to define and manage enumerations with ease, ensuring consistency and reliability across
your codebase.

## Features

- Define and manage string and integer enumerations.
- Easy to use API for adding, retrieving, and validating enum members.
- JSON marshaling and unmarshaling support for enum members.
- Thread-safe operations for concurrent applications.
- Flexible configuration options for strict mode and other behaviors.

## Installation

To install the package, use the following command:

```sh
go get github.com/menduo/mdgoenum
```

## Usage

### String Enum

```go
package main

import (
	"fmt"
	"github.com/menduo/mdgoenum"
)

func main() {
	enum := mdgoenum.NewStringEnum()
	enum.Add("Active", "Represents an active state")
	enum.Add("Inactive", "Represents an inactive state")

	member, err := enum.Get("Active")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Value:", member.GetValue(), "Description:", member.GetDesc())
	}
}
```

### Int Enum

```go
package main

import (
	"fmt"
	"github.com/menduo/mdgoenum"
)

func main() {
	enum := mdgoenum.NewIntEnum()
	enum.Add(1, "One")
	enum.Add(2, "Two")

	member, err := enum.Get(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Value:", member.GetValue(), "Description:", member.GetDesc())
	}
}
```

## MustAdd and MustGet

`MustAdd` and `MustGet` are convenience methods that panic if an error occurs. They are useful for initialization code
where you want to ensure that the enum members are added successfully.

### MustAdd Example

```go
package main

import (
	"github.com/menduo/mdgoenum"
)

func main() {
	enum := mdgoenum.NewIntEnum()
	enum.Add(1, "One")
	enum.MustAdd(1, "One") // Panic: duplicate value
}

```

### MustGet Example

```go
package main

import (
	"github.com/menduo/mdgoenum"
)

func main() {
	enum := mdgoenum.NewIntEnum()
	enum.Add(1, "One")
	enum.MustGet(2) // Panic: member not found
}

```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
