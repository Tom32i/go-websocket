package codec

import (
    "bytes"
)

type PositionCodec struct {
    intCodec Int16Codec
}
type Position struct {
    X uint16
    Y uint16
}

func (c PositionCodec) encode(buffer *bytes.Buffer, data any) {
    position := data.(Position)
    c.intCodec.encode(buffer, position.X)
    c.intCodec.encode(buffer, position.Y)
}

func (c PositionCodec) decode(buffer *bytes.Buffer) any {
    return Position{
        X: c.intCodec.decode(buffer).(uint16),
        Y: c.intCodec.decode(buffer).(uint16),
    }
}

func CreatePositionCodec() PositionCodec {
    return PositionCodec{
        intCodec: Int16Codec{},
    }
}
