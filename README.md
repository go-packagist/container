# Container

[![Go Version](https://badgen.net/github/release/go-packagist/container/stable)](https://github.com/go-packagist/container/releases)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-packagist/container)](https://pkg.go.dev/github.com/go-packagist/container)
[![codecov](https://codecov.io/gh/go-packagist/container/branch/master/graph/badge.svg?token=5TWGQ9DIRU)](https://codecov.io/gh/go-packagist/container)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-packagist/container)](https://goreportcard.com/report/github.com/go-packagist/container)
[![tests](https://github.com/go-packagist/container/actions/workflows/go.yml/badge.svg)](https://github.com/go-packagist/container/actions/workflows/go.yml)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

## Installation

```bash
go get github.com/go-packagist/container
```

## Usage

```go
package main

import "github.com/go-packagist/container"

func main() {
	c := container.NewContainer()

	c.Instance("aa", "aaa")
	c.Bind("bb", func(c *container.Container) interface{} {
		return "bbb"
	}, false)
	c.Singleton("cc", func(c *container.Container) interface{} {
		return "ccc"
	})
	c.Singleton("dd", func(c *container.Container) interface{} {
		return func(hello string) string {
			return "hello " + hello
		}
	})

	var aa, bb, cc string
	c.Make("aa", &aa)
	c.Make("bb", &bb)
	c.Make("cc", &cc)
	println(aa, bb, cc) // aaa bbb ccc

	var dd func(name string) string
	c.Make("dd", &dd)
	println(dd("world")) // hello world
}
```

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.