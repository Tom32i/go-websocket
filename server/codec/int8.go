package codec

import (
    "bytes"
)

type Int8Codec struct {
}

func (c Int8Codec) encode(buffer *bytes.Buffer, data any) {
    var b byte = data.(uint8)
    buffer.Write([]byte{b})
}

func (c Int8Codec) decode(buffer *bytes.Buffer) any {
    b := buffer.Next(1)

    return b[0]
}
