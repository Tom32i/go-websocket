package server

import (
    "log"
    "bytes"
    // "encoding/binary"
    // "encoding/gob"
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
		//codec.setId(id)
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
	log.Printf("data: %v", data)
	var buffer = bytes.NewBuffer(data)
	//log.Printf("next 1: %v", buffer.Next(1))
	//log.Printf("next 1: %v", buffer.Next(1))
	//log.Printf("next 3: %v", buffer.Next(3))
	id := e.idCodec.decode(buffer);
	log.Printf("id: %v", id)
	handler := e.codecsById[id.(uint8)]
	log.Printf("handler: %v", handler)

	return Message {
		Name: handler.getName(),
		Data: handler.decode(buffer),
	}
}

type Codec interface {
	//setId(id int)
	getId() uint8
	getName() string
	//getByteLength(data any) int
	encode(buffer *bytes.Buffer, data any)
	decode(buffer *bytes.Buffer) any
}

type BaseCodec struct {
	id uint8
	name string
}

/*func (c *BaseCodec) setId(id int) {
	c.id = id
}*/

func (c BaseCodec) getId() uint8 {
	return c.id
}


func (c BaseCodec) getName() string {
	return c.name
}

type Int8Codec struct {
	BaseCodec
}

/*func (c Int8Codec) getByteLength(data any) int {
	return 1
}*/

func (c Int8Codec) encode(buffer *bytes.Buffer, data any) {
	var b byte = data.(uint8)
	buffer.Write([]byte{b})
	//log.Printf("encode.data: %v", data)
	//log.Printf("encode.buffer: %v", buffer)
	//log.Printf("encode.bytes: %v", buffer.Bytes())
}

func (c Int8Codec) decode(buffer *bytes.Buffer) any {
	b := buffer.Next(1)
	log.Printf("b: %T %v", b, b)
	return b[0]
}

type StringCodec struct {
	BaseCodec
}

/*func (c StringCodec) getByteLength(data any) int {
	return 2
}*/

func (c StringCodec) encode(buffer *bytes.Buffer, data any) {
	/*length := len(data.(string))
	log.Printf('')
	var b byte = .(uint8)
	buffer.Write([]byte{b})*/
	buffer.WriteString(data.(string))
	//log.Printf("encode.data: %v", data)
	//log.Printf("encode.buffer: %v", buffer)
	//log.Printf("encode.bytes: %v", buffer.Bytes())
}

func (c StringCodec) decode(buffer *bytes.Buffer) any {
	b := buffer.Next(1)
	t := buffer.Next(int(b[0]) * 2)

	return string(t)
}
