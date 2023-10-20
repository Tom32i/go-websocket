package codec

import (
    "bytes"
)

type StringCodec struct {
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
