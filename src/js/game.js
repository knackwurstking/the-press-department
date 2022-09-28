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

export default class Game {
  /**
   * @param {HTMLCanvasElement} canvas
   * @param {CanvasRenderingContext2D} ctx
   * @param {number} hz
   */
  constructor(canvas, ctx, hz) {
    this.canvas = canvas;
    this.ctx = ctx;

    this.updateHz(hz);
    this._lastFrame = 0 - this._fps;
    this._engineFrame = -1;

    /** @type {Assets} */
    this.assets = {};

    /** @type {EngineRollenBahn[]} */
    this.engines = [];

    // touch event handlers
    this.pointer = false;

    /** @type {null|((ev: PointerEvent) => any)} */
    this.onpointerdown = null;
    /** @type {null|((ev: PointerEvent) => any)} */
    this._pointerdown = (ev) => {
      console.log(`[DEBUG] pointerdown`);
      this.view.sX = ev.x;
      this.pointer = true;
    };

    /** @type {null|((ev: PointerEvent) => any)} */
    this.onpointermove = null;
    /** @type {null|((ev: PointerEvent) => any)} */
    this._pointermove = (ev) => {
      if (!this.pointer) return;
      console.log(`[DEBUG] pointermove`);

      this.view.x -= this.view.sX - ev.x;
      this.view.y = this.canvas.height / 2 - 312 / 2;

      this.view.sX = ev.x;
      this.ctx.clearRect(0, 0, this.canvas.width, this.canvas.height);

      this.moveView(this.view);
    };

    /** @type {null|((ev: PointerEvent) => any)} */
    this.onpointerup = null;
    /** @type {null|((ev: PointerEvent) => any)} */
    this._pointerup = () => {
      if (!this.pointer) return;
      this.pointer = false;
      console.log(`[DEBUG] pointerup`);
    };

    /** @type {null|((ev: PointerEvent) => any)} */
    this.onpointercancel = null;
    /** @type {null|((ev: PointerEvent) => any)} */
    this._pointercancel = (ev) => {
      if (!this.pointer) return;
      console.log(`[DEBUG] pointercancel`);
      if (this._pointerup) this._pointerup(ev);
    };

    /** @type {null|((ev: PointerEvent) => any)} */
    this.onpointerout = null;
    /** @type {null|((ev: PointerEvent) => any)} */
    this._pointerout = (ev) => {
      if (!this.pointer) return;
      console.log(`[DEBUG] pointerout`);
      if (this._pointerup) this._pointerup(ev);
    };

    this.canvas.width = window.innerWidth - 2;
    this.canvas.height = window.innerHeight - 4;

    window.onresize = () => {
      this.canvas.width = window.innerWidth - 2;
      this.canvas.height = window.innerHeight - 4;
      this.view.y = this.canvas.height / 2 - 312 / 2;
      this.moveView(this.view);
    };

    this.view = {
      sX: 0,
      x: 0,
      y: this.canvas.height / 2 - 312 / 2,
    };
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
    // create
    let width = data.engine.count * this.assets.rbAluBlockLeft.width;
    let height = this.canvas.height;

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

  /**
   * @param {EngineRollenBahn} engine
   * @param {number} x
   * @param {number} y
   * @returns {EngineRollenBahn}
   */
  updateEngine(engine, x, y) {
    engine.x = x;
    engine.y = y;
    engine.height = this.canvas.height;
    return engine;
  }

  /**
   * @param {{ sX: number, x: number, y: number }} view
   */
  moveView(view) {
    // TODO: just move the background and engines
    let lastX = view.x;
    for (let index = 0; index < Data.rb.length; index++) {
      let engine = this.engines[index];
      if (engine) {
        engine = this.updateEngine(engine, lastX, view.y);
      } else {
        engine = this.createEngine(Data.rb[index], lastX, view.y);
        this.engines.push(engine);
      }

      lastX += engine.width;
    }
  }

  handleUserInput() {
    this.canvas.onpointerdown = this._pointerdown;
    this.canvas.onpointermove = this._pointermove;
    this.canvas.onpointerup = this._pointerup;
    this.canvas.onpointercancel = this._pointercancel;
    this.canvas.onpointerout = this._pointerout;
  }

  /** @param {number} hz */
  updateHz(hz) {
    this._fps = 600 / hz;
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

      for (let engine of this.engines) {
        engine.draw(this.ctx, this._engineFrame);
      }
    }
  }

  drawBackground() {
    // TODO: handle canvas background
  }

  async start() {
    this.moveView(this.view);
    this.handleUserInput();
    const animate = (/** @type {number} */ frame) => {
      this.draw(frame);
      requestAnimationFrame(animate);
    };

    animate(0);
  }
}
