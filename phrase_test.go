/*
 * MIT License
 *
 * Copyright (c) 2022-2024 Luke Whritenour
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

package phrase_test

import (
	"reflect"
	"testing"

	"github.com/lukewhrit/phrase"
)

func TestGenerate(t *testing.T) {
	p := phrase.Default.Generate(2)

	if len(p) != 2 {
		t.Errorf("Length of phrase is wrong; is %d, should be 2", len(p))
	}
}

func TestString(t *testing.T) {
	p := phrase.Default.Generate(2)

	if reflect.TypeOf(p.String()) != reflect.TypeOf("") {
		t.Errorf("phrase is not in string form")
	}

	if len(p) != 2 {
		t.Errorf("Length of phrase is wrong; is %d, should be 2", len(p))
	}
}
