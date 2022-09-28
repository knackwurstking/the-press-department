import Data from "./data";
import { EngineRollenBahn } from "./engine";

/**
 * @typedef Assets
 * @type {{
 *  rolleLeft?: HTMLImageElement,
 *  rolleRight?: HTMLImageElement,
 *  rbAluBlockLeft?: HTMLImageElement,
 *  rbAluBlockRight?: HTMLImageElement,
 *  rbRiemen290x5?: HTMLImageElement,
 *  rbRiemen270x5?: HTMLImageElement,
 * }}
 */

/**
 * @typedef View
 * @type {{
 * canvas: HTMLCanvasElement,
 * pX: number,
 * x: number,
 * y: null | number,
 * getY: () => number
 * }}
 */

export class Game {
  /**
   * @param {HTMLCanvasElement} canvas
   * @param {CanvasRenderingContext2D} ctx
   * @param {number} hz
   */
  constructor(canvas, ctx, hz) {
    this._canvas = canvas;
    this.ctx = ctx;

    this.updateHz(hz);
    this._lastFrame = 0 - this._fps;
    this._engineFrame = -1;

    /** @type {Assets} */
    this.assets = {};

    this._viewChanged = false;

    /** @type {View} */
    this.view = {
      pX: 0, // pointerdown x position
      x: 0,
      y: null,
      canvas: this._canvas,

      getY() {
        return this.y === null ? this.canvas.height / 2 - 321 / 2 : this.y;
      },
    };

    /** @type {EngineRollenBahn[]} */
    this.engines = [];

    // touch event handlers
    this.pointer = false;

    /** @type {null|((ev: PointerEvent) => any)} */
    this.onpointerdown = null;
    /** @type {null|((ev: PointerEvent) => any)} */
    this._pointerdown = (ev) => {
      this.view.pX = ev.x;
      this.pointer = true;
    };

    /** @type {null|((ev: PointerEvent) => any)} */
    this.onpointermove = null;
    /** @type {null|((ev: PointerEvent) => any)} */
    this._pointermove = (ev) => {
      if (!this.pointer) return;

      this.view.x -= this.view.pX - ev.x;

      // update pX value
      this.view.pX = ev.x;
      this._viewChanged = true;
    };

    /** @type {null|((ev: PointerEvent) => any)} */
    this.onpointerup = null;
    /** @type {null|((ev: PointerEvent) => any)} */
    this._pointerup = () => {
      if (!this.pointer) return;
      this.pointer = false;
    };

    /** @type {null|((ev: PointerEvent) => any)} */
    this.onpointercancel = null;
    /** @type {null|((ev: PointerEvent) => any)} */
    this._pointercancel = (ev) => {
      if (!this.pointer) return;
      if (this._pointerup) this._pointerup(ev);
    };

    /** @type {null|((ev: PointerEvent) => any)} */
    this.onpointerout = null;
    /** @type {null|((ev: PointerEvent) => any)} */
    this._pointerout = (ev) => {
      if (!this.pointer) return;
      if (this._pointerup) this._pointerup(ev);
    };

    this._canvas.width = window.innerWidth - 2;
    this._canvas.height = window.innerHeight - 4;

    window.onresize = () => {
      this._canvas.width = window.innerWidth - 2;
      this._canvas.height = window.innerHeight - 4;
      // TODO: redraw
    };
  }

  /** @param {number} hz */
  updateHz(hz) {
    this._fps = 600 / hz;
  }

  /**
   * @param {{
   *  name: string,
   *  engine: {
   *    count: number,
   *    type: number,
   *    side: "left"|"right",
   *  },
   * }} data
   * @param {number} x
   * @param {number} y
   * @returns {EngineRollenBahn}
   */
  createEngine(data, x, y) {
    let width = data.engine.count * this.assets.rbAluBlockLeft.width;
    let height = this._canvas.height;

    let engine = new EngineRollenBahn(
      data.name,
      data.engine.side,
      data.engine.count,
      x,
      y,
      width,
      height,
      this.assets
    );
    return engine;
  }

  buildEngines() {
    this.engines = [];
    let lastX = this.view.x;
    for (let index = 0; index < Data.rb.length; index++) {
      let engine = this.createEngine(Data.rb[index], lastX, this.view.getY());
      this.engines.push(engine);

      lastX += engine.width;
    }
  }

  updatePosition() {
    let lastX = this.view.x;
    for (let index = 0; index < this.engines.length; index++) {
      let engine = this.engines[index];
      engine.x = lastX;
      engine.y = this.view.getY();
      lastX += engine.width;
    }
  }

  handleUserInput() {
    this._canvas.onpointerdown = this._pointerdown;
    this._canvas.onpointermove = this._pointermove;
    this._canvas.onpointerup = this._pointerup;
    this._canvas.onpointercancel = this._pointercancel;
    this._canvas.onpointerout = this._pointerout;
  }

  /**
   * @param {number} frame
   */
  draw(frame) {
    if (frame - this._lastFrame >= this._fps) {
      this._engineFrame += 1;
      this._lastFrame = frame;

      if (this._engineFrame > 5) {
        this._engineFrame = 0;
      }

      this.ctx.clearRect(0, 0, this._canvas.width, this._canvas.height);

      if (this._viewChanged) {
        this.updatePosition();
        this._viewChanged = false;
      }

      for (let engine of this.engines) {
        engine.draw(this.ctx, this._engineFrame);
      }
    }
  }

  async start() {
    this.buildEngines();
    this.handleUserInput();

    const animate = (/** @type {number} */ frame) => {
      this.draw(frame);
      requestAnimationFrame(animate);
    };

    animate(0);
  }
}

/**
 * @param {HTMLCanvasElement} canvas
 * @param {number} rbHz
 * @returns {Promise<Game>}
 */
export async function animate(canvas, rbHz) {
  const ctx = canvas.getContext("2d");
  const game = new Game(canvas, ctx, rbHz);

  // loading assets before runninng the game loop
  const queue = new Set();
  for (let asset of Data.assets) {
    game.assets[asset.name] = new Image(asset.width, asset.height);

    console.log("[DEBUG] load image:", asset.src);
    game.assets[asset.name].src = asset.src;
    game.assets[asset.name].onerror = (ev) => {
      console.warn("[WARNING] load game asset failed:", ev.target.src);
    };

    queue.add(game.assets[asset.name].src);
    game.assets[asset.name].onload = (ev) => {
      queue.delete(ev.target.src);
      if (!queue.size) game.start();
    };
  }

  return game;
}
