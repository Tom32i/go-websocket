package codec

import (
    "bytes"
    "encoding/binary"
)

type Int16Codec struct {
}

func (c *Int16Codec) encode(buffer *bytes.Buffer, data any) {
    b := make([]byte, 2)
    binary.BigEndian.PutUint16(b[0:], data.(uint16))
    buffer.Write(b)
}

func (c *Int16Codec) decode(buffer *bytes.Buffer) any {
    return binary.BigEndian.Uint16(buffer.Next(2)[0:])
}
