package server

import (
    "log"
    "bytes"
    "encoding/binary"
)

type BinaryEncoder struct {
	idCodec Codec
	codecsByName map[string]Codec
	codecsById map[int]Codec
}

func createBinaryEncoder(codecs map[string]Codec, idCodec Codec) BinaryEncoder {
	var codecsByName map[string]Codec
	codecsByName = make(map[string]Codec)
	var codecsById map[int]Codec
	codecsById = make(map[int]Codec)
	id := 0

	for name, codec := range codecs {
		id++
		//codec.setId(id)
		codecsById[id] = codec
		codecsByName[name] = codec
	}

	log.Printf("codecsByName: %v", codecs)
	log.Printf("codecsByName: %v", codecsById)

	return BinaryEncoder{
		idCodec: idCodec,
		codecsByName: codecsByName,
		codecsById: codecsById,
	}
}

func (e BinaryEncoder) encode(name string, data interface{}) []byte {
	log.Printf("encode: %v %v", name, data)
	handler := e.codecsByName[name]
	log.Printf("handler: #%v %v", handler.getId(), handler)
	//buffer := [e.idCodec.getByteLength() + handler.getByteLength(data)]byte{}
	//buffer := new(bytes.Buffer)
	var buffer bytes.Buffer

	e.idCodec.encode(buffer, handler.getId())
	handler.encode(buffer, data)
	//buffer = append(buffer, e.idCodec.encode(handler.getId())...)
	//buffer = append(buffer, handler.encode(data)...)
	log.Printf("buffer: %v", buffer)
	log.Printf("bytes: %v", buffer.Bytes())
	return buffer.Bytes()
}

type Codec interface {
	//setId(id int)
	getId() int
	//getByteLength(data interface{}) int
	encode(buffer bytes.Buffer, data interface{})
}

type BaseCodec struct {
	id int
}

/*func (c *BaseCodec) setId(id int) {
	c.id = id
}*/

func (c BaseCodec) getId() int {
	return c.id
}

type Int8Codec struct {
	BaseCodec
}

/*func (c Int8Codec) getByteLength(data interface{}) int {
	return 1
}*/

func (c Int8Codec) encode(buffer bytes.Buffer, data interface{}) {
    //binary.Write(&buffer, binary.LittleEndian, data)
    //binary.AppendVarint(&buffer, data)

	log.Printf("encode.data: %v", data)
	log.Printf("encode.buffer: %v", buffer)
	log.Printf("encode.bytes: %v", buffer.Bytes())
}

type StringCodec struct {
	BaseCodec
}

/*func (c StringCodec) getByteLength(data interface{}) int {
	return 2
}*/

func (c StringCodec) encode(buffer bytes.Buffer, data interface{}) {
    binary.Write(&buffer, binary.LittleEndian, data)
    //return []byte(data.(string))
}
