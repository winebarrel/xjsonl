package xjsonl

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEachJsonLine(t *testing.T) {
	assert := assert.New(t)

	stdin := bytes.NewBufferString("foo,bar\nbar,zoo")
	stdout := &bytes.Buffer{}

	_ = EachJsonLine(stdin, ",", []string{}, false, func(line string) {
		fmt.Fprintln(stdout, line)
	})

	assert.Equal(`["foo","bar"]
["bar","zoo"]
`, stdout.String())
}

func TestEachJsonLineWithSep(t *testing.T) {
	assert := assert.New(t)

	stdin := bytes.NewBufferString("foo\tbar\nbar\tzoo")
	stdout := &bytes.Buffer{}

	_ = EachJsonLine(stdin, "\t", []string{}, false, func(line string) {
		fmt.Fprintln(stdout, line)
	})

	assert.Equal(`["foo","bar"]
["bar","zoo"]
`, stdout.String())
}

func TestEachJsonLineWithoutSep(t *testing.T) {
	assert := assert.New(t)

	stdin := bytes.NewBufferString("foo,bar\nbar,zoo")
	stdout := &bytes.Buffer{}

	_ = EachJsonLine(stdin, "", []string{}, false, func(line string) {
		fmt.Fprintln(stdout, line)
	})

	assert.Equal(`["foo,bar"]
["bar,zoo"]
`, stdout.String())
}

func TestEachJsonLineWithKeys(t *testing.T) {
	assert := assert.New(t)

	stdin := bytes.NewBufferString("foo,bar\nbar,zoo")
	stdout := &bytes.Buffer{}

	_ = EachJsonLine(stdin, ",", []string{"a", "b"}, false, func(line string) {
		fmt.Fprintln(stdout, line)
	})

	assert.Equal(`{"a":"foo","b":"bar"}
{"a":"bar","b":"zoo"}
`, stdout.String())
}

func TestEachJsonLineWithHeader(t *testing.T) {
	assert := assert.New(t)

	stdin := bytes.NewBufferString("foo,bar\nzoo,baz\n1,2")
	stdout := &bytes.Buffer{}

	_ = EachJsonLine(stdin, ",", []string{}, true, func(line string) {
		fmt.Fprintln(stdout, line)
	})

	assert.Equal(`{"foo":"zoo","bar":"baz"}
{"foo":"1","bar":"2"}
`, stdout.String())
}
