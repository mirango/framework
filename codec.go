package framework

import (
	"io"
)

type Encoder interface {
	EncodingType() string
	Encode(interface{}) ([]byte, error)
	EncodeTo(io.Writer, interface{}) error
}

type Encoders []Encoder

func (e Encoders) Get(encodingType string) Encoder {
	for _, ce := range e {
		if ce.EncodingType() == encodingType {
			return ce
		}
	}
	return nil
}

func (e *Encoders) Append(nc ...Encoder) {
	if nc == nil {
		return
	}
	for _, nce := range nc {
		found := false
		for i, ce := range *e {
			if ce.EncodingType() == nce.EncodingType() {
				(*e)[i] = nce
				found = true
				break
			}
		}
		if !found {
			*e = append(*e, nce)
		}
	}
}

type Decoder interface {
	DecodingType() string
	Decode([]byte, interface{}) error
	DecodeFrom(io.Reader, interface{}) error
}

type Decoders []Decoder

func (d Decoders) Get(decodingType string) Decoder {
	for _, ce := range d {
		if ce.DecodingType() == decodingType {
			return ce
		}
	}
	return nil
}

func (d *Decoders) Append(nc ...Decoder) {
	if nc == nil {
		return
	}
	for _, nce := range nc {
		found := false
		for i, ce := range *d {
			if ce.DecodingType() == nce.DecodingType() {
				(*d)[i] = nce
				found = true
				break
			}
		}
		if !found {
			*d = append(*d, nce)
		}
	}
}

type Codec interface {
	MIMEType() string
	Encoder
	Decoder
}

type Encodable interface {
	Encode() ([]byte, error)
	EncodeTo(io.Writer) error
}

type Codecs []Codec

func (c Codecs) Get(mimeType string) Codec {
	for _, ce := range c {
		if ce.MIMEType() == mimeType {
			return ce
		}
	}
	return nil
}

func (c *Codecs) Append(nc ...Codec) {
	if nc == nil {
		return
	}
	for _, nce := range nc {
		found := false
		for i, ce := range *c {
			if ce.MIMEType() == nce.MIMEType() {
				(*c)[i] = nce
				found = true
				break
			}
		}
		if !found {
			*c = append(*c, nce)
		}
	}
}
