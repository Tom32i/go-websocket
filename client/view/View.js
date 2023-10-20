import Client from '@client/view/Client';

export default class View {
    constructor(element) {
        this.element = element;
        this.clients = new Map();
        this.me = null;

        this.addClient = this.addClient.bind(this);
        this.removeClient = this.removeClient.bind(this);
    }

    setId(id) {
        this.me = id;

        if (this.clients.has(id)) {
            this.clients.get(id).setMe();
        }
    }

    setName(id, name) {
        if (this.clients.has(id)) {
            const client = this.clients.get(id);
            client.setName(name);
            console.log(`Client #${id} is named "${name}".`);
        }
    }

    addClient(id, name) {
        const element = document.createElement('li');
        const client = new Client(id, name, element);
        console.log(this.me, id, this.me === id);

        this.element.appendChild(element);

        this.clients.set(id, client);

        console.log(`New client #${id}: "${name}".`);
    }

    removeClient(id) {
        if (this.clients.has(id)) {
            const client = this.clients.get(id);
            client.element.remove();
            this.clients.delete(id);
            console.log(`Client #${id} left.`);
        }
    }
}
