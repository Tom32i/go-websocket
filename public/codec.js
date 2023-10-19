const { Client, BinaryEncoder, Int8Codec, StringCodec, Codec } = netcode;

class ClientAddCodec extends Codec {
    constructor() {
        super();

        this.intCodec = new Int8Codec();
        this.stringCodec = new StringCodec();
    }

    /**
     * @type {Number}
     */
    getByteLength(data) {
        return this.intCodec.getByteLength() + this.stringCodec.getByteLength(data.name);
    }

    /**
     * {@inheritdoc}
     */
    encode(buffer, offset, data) {
        const { id, name } = data;

        this.intCodec.encode(buffer, offset, id);
        this.stringCodec.encode(buffer, offset + this.intCodec.getByteLength(), name);
    }

    /**
     * {@inheritdoc}
     */
    decode(buffer, offset) {
        const id = this.intCodec.decode(buffer, offset);
        const name = this.stringCodec.decode(buffer, offset + this.intCodec.getByteLength());

        return { id, name };
    }
}
