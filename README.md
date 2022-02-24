# Hyperscale Locale [![Last release](https://img.shields.io/github/release/hyperscale-stack/locale.svg)](https://github.com/hyperscale-stack/locale/releases/latest) [![Documentation](https://godoc.org/github.com/hyperscale-stack/locale?status.svg)](https://godoc.org/github.com/hyperscale-stack/locale)

[![Go Report Card](https://goreportcard.com/badge/github.com/hyperscale-stack/locale)](https://goreportcard.com/report/github.com/hyperscale-stack/locale)

| Branch | Status                                                                                                                                                                     | Coverage                                                                                                                                               |
| ------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| master | [![Build Status](https://github.com/hyperscale-stack/locale/workflows/Go/badge.svg?branch=master)](https://github.com/hyperscale-stack/locale/actions?query=workflow%3AGo) | [![Coveralls](https://img.shields.io/coveralls/hyperscale-stack/locale/master.svg)](https://coveralls.io/github/hyperscale-stack/locale?branch=master) |

The Hyperscale locale library provides a simple locale manager

## Example

```go
package main

import (
    "fmt"

    "github.com/hyperscale-stack/locale"
)

func main() {
    ctx := context.Background()

    ctx = locale.ToContext(ctx, language.French)

    tag := locale.FromContext(ctx) // return language.French or locale.DefaultLocale if
}

```

## License

Hyperscale Locale is licensed under [the MIT license](LICENSE.md).
