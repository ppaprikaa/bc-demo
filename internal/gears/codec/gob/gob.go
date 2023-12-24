package gob

import (
	"bytes"
	"encoding/gob"
)

func SliceToBytes[T any](slice []T) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	err := enc.Encode(slice)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func BytesToSlice[T any](data []byte) ([]T, error) {
	var slice []T
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	err := dec.Decode(&slice)
	if err != nil {
		return nil, err
	}

	return slice, nil
}
