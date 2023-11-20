//import '@css/style.scss';
import Client from 'netcode/src/client/Client';
import MoveHandler from '@client/view/MoveHandler';
import BinaryEncoder from 'netcode/src/encoder/BinaryEncoder';
import Int8Codec from 'netcode/src/encoder/codec/Int8Codec';
import Int16Codec from 'netcode/src/encoder/codec/Int16Codec';
import StringCodec from 'netcode/src/encoder/codec/StringCodec';
import ClientAddCodec from '@client/codec/ClientAddCodec';
import ClientNameCodec from '@client/codec/ClientNameCodec';
import PositionCodec from '@client/codec/PositionCodec';
import ClientPositionCodec from '@client/codec/ClientPositionCodec';
import View from '@client/view/View';

const protocol = location.protocol === 'https:' ? 'wss' : 'ws';
const client = new Client(
    `${protocol}://${location.host}/ws`,
    new BinaryEncoder([
        ['me:id', new Int8Codec()],
        ['me:name', new StringCodec()],
        ['me:position', new PositionCodec()],
        ['client:add', new ClientAddCodec()],
        ['client:remove', new Int8Codec()],
        ['client:name', new ClientNameCodec()],
        ['client:position', new ClientPositionCodec()],
        ['say', new StringCodec()],
    ])
);

const view = new View(document.getElementById('clients'));
const nameInput = document.getElementById('name');
nameInput.value = (Math.random() + 1).toString(36).substring(7);
nameInput.addEventListener('change', () => client.send('me:name', nameInput.value));

function handleMove(x, y) {
    if (view.me !== null) {
        view.me.setPosition(x, y);
        client.send('me:position', { x, y });
    }
}

const move = new MoveHandler(handleMove);

client.on('open', () => {
    client.on('me:id', event => {
        console.log(`My id is ${event.detail}, my name is ${nameInput.value}.`);
        view.setMe(event.detail);
        client.send('me:name', nameInput.value);
    });
    client.on('client:add', ({ detail }) => view.addClient(detail.id, detail.name));
    client.on('client:remove', ({ detail }) => view.removeClient(detail));
    client.on('client:name', ({ detail }) => view.setClientName(detail.id, detail.name));
    client.on('client:position', ({ detail }) => view.setClientPosition(detail.id, detail.x, detail.y));
    client.on('say', event => {
        console.log(`Server says "${event.detail}".`);
    });
});
