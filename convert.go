package xjsonl

import (
	"bufio"
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
	if len(keys) == 0 {
		serializer = func(cols []string) string {
			vals := make([]string, len(cols))

			for i, c := range cols {
				val, _ := json.Marshal(c)
				vals[i] = string(val)
			}

			return "[" + strings.Join(vals, ",") + "]"
		}
	} else {
		keysLen := len(keys)
		jsonKeys := make([]string, keysLen)

		for i, k := range keys {
			key, _ := json.Marshal(k)
			jsonKeys[i] = string(key)
		}

		serializer = func(cols []string) string {
			n := len(cols)

			if keysLen < n {
				n = keysLen
			}

			keyVals := make([]string, n)

			for i := 0; i < n; i++ {
				val, _ := json.Marshal(cols[i])
				keyVals[i] = jsonKeys[i] + ":" + string(val)
			}

			return "{" + strings.Join(keyVals, ",") + "}"
		}
	}

	return
}
