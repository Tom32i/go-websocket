import Codec from 'netcode/src/encoder/codec/Codec';
import Int16Codec from 'netcode/src/encoder/codec/Int16Codec';
import StringCodec from 'netcode/src/encoder/codec/StringCodec';

export default class PositionCodec extends Codec {
    constructor() {
        super();

        this.intCodec = new Int16Codec();
    }

    /**
     * @type {Number}
     */
    getByteLength(data) {
        return this.intCodec.getByteLength() * 2;
    }

    /**
     * {@inheritdoc}
     */
    encode(buffer, offset, data) {
        const { x, y } = data;

        this.intCodec.encode(buffer, offset, x);
        this.intCodec.encode(buffer, offset + this.intCodec.getByteLength(), y);
    }

    /**
     * {@inheritdoc}
     */
    decode(buffer, offset) {
        const x = this.intCodec.decode(buffer, offset);
        const y = this.intCodec.decode(buffer, offset + this.intCodec.getByteLength());

        return { x, y };
    }
}
