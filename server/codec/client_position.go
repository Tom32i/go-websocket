package codec

import (
    "bytes"
)

type ClientPositionCodec struct {
    idCodec *Int8Codec
    positionCodec *PositionCodec
}

type ClientPosition struct {
    Id uint8
    Position Position
}

func (c *ClientPositionCodec) encode(buffer *bytes.Buffer, data any) {
    message := data.(ClientPosition)
    c.idCodec.encode(buffer, message.Id)
    c.positionCodec.encode(buffer, message.Position)
}

func (c *ClientPositionCodec) decode(buffer *bytes.Buffer) any {
    return ClientPosition{
        Id: c.idCodec.decode(buffer).(uint8),
        Position: c.positionCodec.decode(buffer).(Position),
    }
}

func CreateClientPositionCodec() *ClientPositionCodec {
    return &ClientPositionCodec{
        idCodec: &Int8Codec{},
        positionCodec: &PositionCodec{},
    }
}
