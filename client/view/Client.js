export default class Client {
    constructor(id, name, element) {
        this.id = id;
        this.name = name;
        this.element = element;
        this.me = false;
        this.x = 0;
        this.y = 0;

        this.element.id = `client-${id}`
        this.element.innerText = name;
    }

    setPosition(x, y) {
        this.x = x;
        this.y = y;
    }

    setName(name) {
        this.name = name;
    }

    setMe() {
        this.me = true;
        this.element.className = "me";
    }

    render() {
        this.element.style.transform = `translate3d(${this.x}px, ${this.y}px, 0px)`
        this.element.innerText = `${this.name} (${this.x},${this.y})`;
    }
}
