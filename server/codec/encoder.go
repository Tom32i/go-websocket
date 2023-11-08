package codec

import (
	"bytes"
)

type Codec interface {
	encode(buffer *bytes.Buffer, data any)
	decode(buffer *bytes.Buffer) any
}

type BinaryEncoder struct {
	idCodec      Codec
	codecsByName map[string]RegisteredCodec
	codecsById   map[uint8]RegisteredCodec
}

type RegisteredCodec struct {
	Id      uint8
	Name    string
	Handler Codec
}

type Message struct {
	Name string
	Data any
}

func CreateBinaryEncoder(codecs []RegisteredCodec, idCodec Codec) BinaryEncoder {
	codecsByName := make(map[string]RegisteredCodec)
	codecsById := make(map[uint8]RegisteredCodec)

	for index, codec := range codecs {
		codec.Id = uint8(index)
		codecsById[codec.Id] = codec
		codecsByName[codec.Name] = codec
	}

	return BinaryEncoder{
		idCodec:      idCodec,
		codecsByName: codecsByName,
		codecsById:   codecsById,
	}
}

func (e *BinaryEncoder) Encode(name string, data any) []byte {
	codec := e.codecsByName[name]

	var buffer bytes.Buffer

	e.idCodec.encode(&buffer, codec.Id)
	codec.Handler.encode(&buffer, data)

	return buffer.Bytes()
}

func (e *BinaryEncoder) Decode(data []byte) Message {
	var buffer = bytes.NewBuffer(data)
	id := e.idCodec.decode(buffer)
	codec := e.codecsById[id.(uint8)]

	return Message{
		Name: codec.Name,
		Data: codec.Handler.decode(buffer),
	}
}
