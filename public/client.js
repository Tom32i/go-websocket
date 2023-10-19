
const events = [
    ['id', new Int8Codec()],
    ['me:name', new StringCodec()],
    ['client:add', new ClientAddCodec()],
    ['say', new StringCodec()],
];
const client = new Client('ws://localhost:8032/ws', new BinaryEncoder(events))
const name = Math.round(Math.random() * 10000).toString(16);
// const name = 'tom32i€';

function coucou() {
    console.log('coucou 2');
    client.send('say', 'Hello again');
}

client.on('open', () => {
    console.log('open');
    client.on('id', event => {
        console.log(`My id is ${event.detail}, my name is ${name}.`);
        client.send('me:name', name);
    });
    client.on('client:add', event => {
        const { id, name } = event.detail;
        console.log(`New client #${id}: "${name}".`);
        client.send('client:add', { id, name });
    });
    client.on('say', event => {
        console.log(`Server says "${event.detail}".`);
    });
    //client.send('say', 'Hello world!');
    //setTimeout(coucou, 2000);
    //client.send('say', 42);
});
