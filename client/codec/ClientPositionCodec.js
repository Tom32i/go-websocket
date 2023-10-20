import Codec from 'netcode/src/encoder/codec/Codec';
import Int8Codec from 'netcode/src/encoder/codec/Int8Codec';
import Int16Codec from 'netcode/src/encoder/codec/Int16Codec';
import StringCodec from 'netcode/src/encoder/codec/StringCodec';
import PositionCodec from '@client/codec/PositionCodec';

export default class ClientPositionCodec extends Codec {
    constructor() {
        super();

        this.idCodec = new Int8Codec();
        this.positionCodec = new PositionCodec();
    }

    /**
     * @type {Number}
     */
    getByteLength(data) {
        return this.idCodec.getByteLength() + this.positionCodec.getByteLength();
    }

    /**
     * {@inheritdoc}
     */
    encode(buffer, offset, data) {
        const { id, x, y } = data;

        this.idCodec.encode(buffer, offset, id);
        this.positionCodec.encode(buffer, offset + this.idCodec.getByteLength(), { x, y });
    }

    /**
     * {@inheritdoc}
     */
    decode(buffer, offset) {
        const id = this.idCodec.decode(buffer, offset);
        const { x, y } = this.positionCodec.decode(buffer, offset + this.idCodec.getByteLength());

        return { id, x, y };
    }
}
