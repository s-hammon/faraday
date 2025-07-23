package main

import (
	"io"
)

type Decoder struct {
	r       io.Reader
	scanned int64
	err     error
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}
