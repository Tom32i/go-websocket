//import '@css/style.scss';
import Client from 'netcode/src/client/Client';
import BinaryEncoder from 'netcode/src/encoder/BinaryEncoder';
import Int8Codec from 'netcode/src/encoder/codec/Int8Codec';
import StringCodec from 'netcode/src/encoder/codec/StringCodec';
import ClientAddCodec from '@client/codec/ClientAddCodec';
import ClientNameCodec from '@client/codec/ClientNameCodec';
import View from '@client/view/View';

const client = new Client(
    'ws://localhost:8032/ws',
    new BinaryEncoder([
        ['me:id', new Int8Codec()],
        ['me:name', new StringCodec()],
        ['client:add', new ClientAddCodec()],
        ['client:remove', new Int8Codec()],
        ['client:name', new ClientNameCodec()],
        ['say', new StringCodec()],
    ])
);

const view = new View(
    document.getElementById('clients')
);

const name = (Math.random() + 1).toString(36).substring(7);
// const name = 'tom32i€';

client.on('open', () => {
    client.on('me:id', event => {
        console.log(`My id is ${event.detail}, my name is ${name}.`);
        view.setId(event.detail);
        client.send('me:name', name);
    });
    client.on('client:add', ({ detail }) => view.addClient(detail.id, detail.name));
    client.on('client:remove', ({ detail }) => view.removeClient(detail));
    client.on('client:name', ({ detail }) => view.setName(detail.id, detail.name));
    client.on('say', event => {
        console.log(`Server says "${event.detail}".`);
    });
});
