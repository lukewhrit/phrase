# phrase

[![Go Reference](https://pkg.go.dev/badge/github.com%2Flukewhrit%2Fphrase.svg)](https://pkg.go.dev/github.com%2Flukewhrit%2Fphrase) [![Go Report Card](https://goreportcard.com/badge/github.com/lukewhrit/phrase)](https://goreportcard.com/report/github.com/lukewhrit/phrase)

phrase is a simple library that generates [passphrases](https://www.explainxkcd.com/wiki/index.php/936:_Password_Strength) using pure Go and a customizable dictionary.

**The default dictionary that ships with phrase is not large; some 2300 words. This is probably not a big enough data set for secure passwords. Because of that, you should consider implementing a [custom dictionary](#dictionary) before using this for passwords.**

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

A dictionary is an array of strings with one additional function (which itself has another function), `Generate()`. Generate takes one parameter and returns an array of randomly-generated strings. The parameter being `words`, an integer, which dictates how long the returned array will be.

This library provides a default dictionary which is good enough for most (unsecure) purposes, it is available under the `Default` name (`phrase.Default`). From this dictionary, all normal functions and behaviors are available; no special conditions.

To create a custom dictionary (which you should be doing if you're generating passwords), simply define a new variable using the `phrase.Dictionary` type! A dictionary is a plain-old array of strings (literally `[]string`). If you need to programmatically create this array, say if you want to read from a file, you can easily populate it using any function (as long as it's called before `Generate()` is used), and the `append` built-in.

Something along the lines of:

```go
package main

var p Dictionary

func init() {
	p = append(p, "a")
	p = append(p, "b")
	p = append(p, "c")
}

func main() {
	Default.Generate(2) // => ["a", "c", "b"]
	Default.Generate(2).String() // => "b-a-c"
}
```

## Credits and License

MIT License (c) 2022 Luke Whritenour. See [`LICENSE`](LICENSE).

The world list is graciously taken from [correcthorse by Nicolás Bevacqua](https://github.com/bevacqua/correcthorse), also under the MIT license.
