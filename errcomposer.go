// The MIT License (MIT)

// Copyright © 2017 Dmitry Moskowski

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
package errcomposer

// ComposedError is an error that aggregates multiple
// errors.
type ComposedError struct {
	errors []error
}

// Error returns the aggregation of the composed error messages.
func (ce *ComposedError) Error() string {
	var errStr string
	for k, err := range ce.errors {
		if k > 0 {
			errStr += " > "
		}
		errStr += err.Error()
	}
	return errStr
}

// Compose multiple errors.
func Compose(errors ...error) error {
	return &ComposedError{errors}
}

// Decompose the composed error
// It tries to decompose the error and if it can
// will return a slice of errors and true, if it failed
// to decompose the error it will return nil and false.
// Just like the map value retrieval by the key.
func Decompose(err error) ([]error, bool) {
	if composed, ok := err.(*ComposedError); ok {
		return composed.errors, true
	}
	return nil, false
}

// Has checks if err is inside composedErr
// If err is one of composedErr members then it will return true.
// If composedErr is not a ComposedError then it will compare
// composedErr with err returning the result.
// All composed errors processsed recursively, it is ok to have
// nested composed errors.
func Has(composedErr, err error) bool {
	if errs, ok := Decompose(composedErr); ok {
		for _, e := range errs {
			if Has(e, err) {
				return true
			}
		}
		return false
	}
	return composedErr == err
}
