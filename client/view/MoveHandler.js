export default class MoveHandler {
    constructor(callback) {
        this.callback = callback;

        this.touch = false;
        this.handleFistTouch = this.handleFistTouch.bind(this);
        this.handleStart = this.handleStart.bind(this);
        this.handleMove = this.handleMove.bind(this);
        this.handleMouse = this.handleMouse.bind(this);

        this.x = 0;
        this.y = 0;

        window.addEventListener('mousemove',  this.handleMouse);
        window.addEventListener('touchstart', this.handleFistTouch);
    }

    setPosition(sx, sy) {
        const x = Math.round(sx);
        const y = Math.round(sy);

        if (this.x !== x || this.y !== y) {
            this.x = x;
            this.y = y;
            this.callback(this.x, this.y);
        }
    }

    handleFistTouch() {
        this.touch = true;
        window.removeEventListener('mousemove',  this.handleMouse);
        window.removeEventListener('touchstart', this.handleFistTouch);
        window.addEventListener('touchstart', this.handleStart);
        window.addEventListener('touchmove', this.handleMove);
    }

    handleStart(event) {
        const { clientX: x, clientY: y } = event.touches[0];
        this.setPosition(x, y);
    }

    handleMove(event) {
        const { clientX: x, clientY: y } = event.touches[0];
        this.setPosition(x, y);
    }

    handleMouse(event) {
        const { clientX: x, clientY: y } = event;
        this.callback(x, y);
    }
}
