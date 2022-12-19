/*
 * MIT License
 *
 * Copyright (c) 2022 Luke Whritenour
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package phrase

import (
	"math/rand"
	"strings"
	"time"
)

type Dictionary []string
type Phrase []string

// Generate populates an array of length `words` with random words from `d`
func (d Dictionary) Generate(words int) Phrase {
	rand.Seed(time.Now().UnixNano())

	phrase := Phrase{}

	for i := 0; i < words; i++ {
		number := rand.Intn(words) + words
		phrase = append(phrase, d[number])
	}

	return phrase
}

func (p Phrase) String() string {
	return strings.Join(p, "-")
}
