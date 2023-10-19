package server

import (
    "bytes"
)

type ClientAddCodec struct {
    BaseCodec
    idCodec Int8Codec
    nameCodec StringCodec
}

type ClientAddMessage struct {
    id uint8
    name string
}

func (c ClientAddCodec) encode(buffer *bytes.Buffer, data any) {
    message := data.(ClientAddMessage)
    c.idCodec.encode(buffer, message.id)
    c.nameCodec.encode(buffer, message.name)
}

func (c ClientAddCodec) decode(buffer *bytes.Buffer) any {
    return ClientAddMessage{
        id: c.idCodec.decode(buffer).(uint8),
        name: c.nameCodec.decode(buffer).(string),
    }
    /*b := buffer.Next(1)
    t := buffer.Next(int(b[0]) * 2)

    return string(t)*/
}

func createClientAddCodec(id uint8, name string) ClientAddCodec {
    return ClientAddCodec{
        BaseCodec: BaseCodec{
            id: id,
            name: name,
        },
        idCodec: Int8Codec{},
        nameCodec: StringCodec{},
    }
}
