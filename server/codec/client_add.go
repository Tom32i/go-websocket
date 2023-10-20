package codec

import (
    "bytes"
)

type ClientAddCodec struct {
    idCodec Int8Codec
    nameCodec StringCodec
}

type ClientAddMessage struct {
    Id uint8
    Name string
}

func (c ClientAddCodec) encode(buffer *bytes.Buffer, data any) {
    message := data.(ClientAddMessage)
    c.idCodec.encode(buffer, message.Id)
    c.nameCodec.encode(buffer, message.Name)
}

func (c ClientAddCodec) decode(buffer *bytes.Buffer) any {
    return ClientAddMessage{
        Id: c.idCodec.decode(buffer).(uint8),
        Name: c.nameCodec.decode(buffer).(string),
    }
}

func CreateClientAddCodec() ClientAddCodec {
    return ClientAddCodec{
        idCodec: Int8Codec{},
        nameCodec: StringCodec{},
    }
}
