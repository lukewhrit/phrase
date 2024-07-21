# phrase

[![Go Reference](https://pkg.go.dev/badge/github.com%2Flukewhrit%2Fphrase.svg)](https://pkg.go.dev/github.com%2Flukewhrit%2Fphrase) [![Go Report Card](https://goreportcard.com/badge/github.com/lukewhrit/phrase)](https://goreportcard.com/report/github.com/lukewhrit/phrase)

phrase is a simple library that generates [passphrases](https://www.explainxkcd.com/wiki/index.php/936:_Password_Strength) using pure Go and a customizable dictionary.

> [!WARNING]
> The default dictionary shipped with *phrase* is small and should not be used to generate secure passwords. Consider [defining your own dictionary](#dictionary) with a much larger selection of words before generating passwords.

## Documentation

### Installation

```
go get github.com/lukewhrit/phrase
```

### Basic Example

```go
package main

import "github.com/lukewhrit/phrase"

// Use the default dictionary
var p = phrase.Default

func main() {
    p.Generate(3) // => ["enrage", "juice", "await"]
    p.Generate(3).String() // => "trimmer-gothic-uncle"
}
```

### Dictionary

A dictionary is an array of strings that is used to generate the phrases. You can call the `Dictionary.Generate()` function to create a passphrase. The function requires a "length" parameter and will return an array of random words. A helper function is provided on the array that allows you to convert it to a string.

A default dictionary is provided, containing ~2300 words, to generate passphrases with. It is available under the name `phrase.Default`. **We strongly advise against using this library for generating secure passwords. A larger dictionary should be used instead.**

You can define a custom dictionary by instantiating your own `Dictionary` object. For example:

```go
package main

import "github.com/lukewhrit/phrase"

var p = phrase.Dictionary{"a","c","b"}

func main() {
	p.Generate(3) // => ["a", "c", "b"]
	p.Generate(3).String() // => "b-a-c"
}
```

## Credits and License

MIT License (c) 2022 Luke Whritenour. See [`LICENSE`](LICENSE).

The world list is graciously taken from [correcthorse by Nicolás Bevacqua](https://github.com/bevacqua/correcthorse), also under the MIT license.
