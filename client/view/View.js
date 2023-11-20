import Client from '@client/view/Client';
import Loop from '@client/view/Loop';

export default class View {
    constructor(list, counter) {
        this.list = list;
        this.counter = counter;
        this.clients = new Map();
        this.loop = new Loop();
        this.me = null;

        this.addClient = this.addClient.bind(this);
        this.removeClient = this.removeClient.bind(this);
        this.render = this.render.bind(this);

        this.loop.setCallback(this.render);
        this.loop.start();
    }

    setMe(id) {
        if (this.clients.has(id)) {
            const client = this.clients.get(id);
            client.setMe();
            this.me = client;
        }
    }

    setClientName(id, name) {
        if (this.clients.has(id)) {
            const client = this.clients.get(id);
            client.setName(name);
            console.log(`Client #${id} is named "${name}".`);
        }
    }

    setClientPosition(id, x, y) {
        if (this.clients.has(id)) {
            this.clients.get(id).setPosition(x, y);
        }
    }

    addClient(id, name) {
        const element = document.createElement('li');
        const client = new Client(id, name, element);

        this.list.appendChild(element);
        this.clients.set(id, client);
        this.counter.innerText = this.clients.size;

        console.log(`New client #${id}: "${name}".`);
    }

    removeClient(id) {
        if (this.clients.has(id)) {
            const client = this.clients.get(id);
            client.element.remove();
            this.clients.delete(id);
            this.counter.innerText = this.clients.size;
            console.log(`Client #${id} left.`);
        }
    }

    render() {
        this.clients.forEach(client => client.render());
    }
}
