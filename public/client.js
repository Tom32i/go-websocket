const { Client, BinaryEncoder, Int8Codec, StringCodec } = netcode;
const events = [
    ['id', new Int8Codec()],
    ['name', new StringCodec()],
    ['say', new StringCodec()],
];
const client = new Client('ws://localhost:8032/ws', new BinaryEncoder(events))
//const name = (Math.random() * 1000).toString(16);
const name = 'tom';

function coucou() {
    console.log('coucou 2');
    client.send('say', 'Hello again');
}

client.on('open', () => {
    console.log('open');
    client.on('id', event => {
        console.log(`My id is ${event.detail}.`);
        client.send('name', name);
        console.log(`My name is ${name}.`)
    });
    //client.send('say', 'Hello world!');
    //setTimeout(coucou, 2000);
    //client.send('say', 42);
});
