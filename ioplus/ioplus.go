package ioplus

import (
	"io"
	"io/ioutil"
)

// is ioutil.Discard
var Discard = ioutil.Discard

// call ioutil.NopCloser
func NopCloser(r io.Reader) io.ReadCloser {
	return ioutil.NopCloser(r)
}

// call ioutil.ReadAll
func ReadAll(r io.Reader) ([]byte, error) {
	return ioutil.ReadAll(r)
}
