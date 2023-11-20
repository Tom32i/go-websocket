export default class MoveHandler {
    constructor(callback) {
        this.callback = callback;

        this.handleFistTouch = this.handleFistTouch.bind(this);
        this.handleStart = this.handleStart.bind(this);
        this.handleMove = this.handleMove.bind(this);
        this.handleMouse = this.handleMouse.bind(this);

        window.addEventListener('mousemove',  this.handleMouse);
        window.addEventListener('touchstart', this.handleFistTouch);
    }

    handleFistTouch() {
        window.removeEventListener('mousemove',  this.handleMouse);
        window.removeEventListener('touchstart', this.handleFistTouch);
        window.addEventListener('touchstart', this.handleStart);
        window.addEventListener('touchmove', this.handleMove);
    }

    handleStart(event) {
        const { clientX: x, clientY: y } = event.touches[0];
        this.callback(Math.round(x), Math.round(y));
    }

    handleMove(event) {
        const { clientX: x, clientY: y } = event.touches[0];
        this.callback(Math.round(x), Math.round(y));
    }

    handleMouse(event) {
        const { clientX: x, clientY: y } = event;
        this.callback(x, y);
    }
}
