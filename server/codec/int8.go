package codec

import (
    "bytes"
)

type Int8Codec struct {
}

func (c Int8Codec) encode(buffer *bytes.Buffer, data any) {
    buffer.Write([]byte{data.(uint8)})
}

func (c Int8Codec) decode(buffer *bytes.Buffer) any {
    return buffer.Next(1)[0]
}
