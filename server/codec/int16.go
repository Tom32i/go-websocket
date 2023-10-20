package codec

import (
    //"log"
    "bytes"
    "encoding/binary"
)

type Int16Codec struct {
}

func (c Int16Codec) encode(buffer *bytes.Buffer, data any) {
    value := data.(uint16)
    //log.Printf("value: %T %v", value, value)
    b := make([]byte, 2)
    //log.Printf("b: %T %v", b, b)
    binary.BigEndian.PutUint16(b[0:], value)
    //log.Printf("b: %T %v", b, b)
    buffer.Write(b)
}

func (c Int16Codec) decode(buffer *bytes.Buffer) any {
    //log.Printf("decode: %v", buffer.Bytes())
    b := buffer.Next(2)
    //log.Printf("decode: %v", b)

    return binary.BigEndian.Uint16(b[0:])
}
