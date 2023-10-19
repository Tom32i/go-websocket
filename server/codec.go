package server

import (
    // "log"
    "bytes"
)

type BinaryEncoder struct {
	idCodec Codec
	codecsByName map[string]Codec
	codecsById map[uint8]Codec
}

func createBinaryEncoder(codecs []Codec, idCodec Codec) BinaryEncoder {
	var codecsByName map[string]Codec
	codecsByName = make(map[string]Codec)
	var codecsById map[uint8]Codec
	codecsById = make(map[uint8]Codec)

	for _, codec := range codecs {
		codecsById[codec.getId()] = codec
		codecsByName[codec.getName()] = codec
	}

	return BinaryEncoder{
		idCodec: idCodec,
		codecsByName: codecsByName,
		codecsById: codecsById,
	}
}

func (e BinaryEncoder) encode(name string, data any) []byte {
	handler := e.codecsByName[name]

	var buffer bytes.Buffer

	e.idCodec.encode(&buffer, handler.getId())
	handler.encode(&buffer, data)

	return buffer.Bytes()
}

func (e BinaryEncoder) decode(data []byte) Message {
	var buffer = bytes.NewBuffer(data)
	id := e.idCodec.decode(buffer);
	handler := e.codecsById[id.(uint8)]

	return Message {
		name: handler.getName(),
		data: handler.decode(buffer),
	}
}

type Codec interface {
	getId() uint8
	getName() string
	encode(buffer *bytes.Buffer, data any)
	decode(buffer *bytes.Buffer) any
}

type BaseCodec struct {
	id uint8
	name string
}

func (c BaseCodec) getId() uint8 {
	return c.id
}


func (c BaseCodec) getName() string {
	return c.name
}

type Int8Codec struct {
	BaseCodec
}

func (c Int8Codec) encode(buffer *bytes.Buffer, data any) {
	var b byte = data.(uint8)
	buffer.Write([]byte{b})
}

func (c Int8Codec) decode(buffer *bytes.Buffer) any {
	b := buffer.Next(1)

	return b[0]
}

type StringCodec struct {
	BaseCodec
}

func (c StringCodec) encode(buffer *bytes.Buffer, data any) {
	length := len(data.(string))
	var b byte = uint8(length)
	buffer.Write([]byte{b})
	buffer.WriteString(data.(string))
}

func (c StringCodec) decode(buffer *bytes.Buffer) any {
	length := int(buffer.Next(1)[0])
	value := buffer.Next(length)

	return string(value)
}
