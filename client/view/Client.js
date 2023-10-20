export default class Client {
    constructor(id, name, element) {
        this.id = id;
        this.name = name;
        this.element = element;
        this.me = false;

        this.element.id = `client-${id}`
        this.element.innerText = name;
    }

    setName(name) {
        this.name = name;
        this.element.innerText = name;
    }

    setMe() {
        this.me = true;
        this.element.className = "me";
    }
}
