package xjsonl

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"strings"
)

var ReadLineBufSize = 4096

func EachJsonLine(file io.Reader, sep string, keys []string, cb func(string)) error {
	reader := bufio.NewReader(file)
	serialize := newSerializer(keys)

	for {
		line, err := readLine(reader)

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		cols := strings.Split(line, sep)
		json := serialize(cols)
		cb(json)
	}

	return nil
}

func readLine(reader *bufio.Reader) (string, error) {
	buf := make([]byte, 0, ReadLineBufSize)
	var err error

	for {
		line, isPrefix, e := reader.ReadLine()
		err = e

		if len(line) > 0 {
			buf = append(buf, line...)
		}

		if !isPrefix || err != nil {
			break
		}
	}

	return string(buf), err
}

func newSerializer(keys []string) (serializer func([]string) string) {
	marshal := newMarshaller()

	if len(keys) == 0 {
		serializer = func(cols []string) string {
			vals := make([]string, len(cols))

			for i, c := range cols {
				vals[i] = marshal(c)
			}

			return "[" + strings.Join(vals, ",") + "]"
		}
	} else {
		keysLen := len(keys)
		jsonKeys := make([]string, keysLen)

		for i, k := range keys {
			jsonKeys[i] = marshal(k)
		}

		serializer = func(cols []string) string {
			n := len(cols)

			if keysLen < n {
				n = keysLen
			}

			keyVals := make([]string, n)

			for i := 0; i < n; i++ {
				keyVals[i] = jsonKeys[i] + ":" + marshal(cols[i])
			}

			return "{" + strings.Join(keyVals, ",") + "}"
		}
	}

	return
}

func newMarshaller() func(string) string {
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)

	return func(s string) string {
		defer buf.Reset()
		_ = encoder.Encode(s)
		return strings.TrimRight(buf.String(), "\n")
	}
}
