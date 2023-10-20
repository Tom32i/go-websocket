package codec

import (
    "bytes"
)

type ClientNameCodec struct {
    idCodec Int8Codec
    nameCodec StringCodec
}

type ClientNameMessage struct {
    Id uint8
    Name string
}

func (c ClientNameCodec) encode(buffer *bytes.Buffer, data any) {
    message := data.(ClientNameMessage)
    c.idCodec.encode(buffer, message.Id)
    c.nameCodec.encode(buffer, message.Name)
}

func (c ClientNameCodec) decode(buffer *bytes.Buffer) any {
    return ClientNameMessage{
        Id: c.idCodec.decode(buffer).(uint8),
        Name: c.nameCodec.decode(buffer).(string),
    }
}

func CreateClientNameCodec() ClientNameCodec {
    return ClientNameCodec{
        idCodec: Int8Codec{},
        nameCodec: StringCodec{},
    }
}
