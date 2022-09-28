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

export default class Game {
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
      console.log(`[DEBUG] pointerdown`);
      this.view.pX = ev.x;
      this.pointer = true;
    };

    /** @type {null|((ev: PointerEvent) => any)} */
    this.onpointermove = null;
    /** @type {null|((ev: PointerEvent) => any)} */
    this._pointermove = (ev) => {
      if (!this.pointer) return;
      console.log(`[DEBUG] pointermove`);

      this.view.x -= this.view.pX - ev.x;
      this.view.y = this.view.canvas.height / 2 - 312 / 2;

      // update pX value
      this.view.pX = ev.x;
      this._viewChanged;
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

    this._canvas.width = window.innerWidth - 2;
    this._canvas.height = window.innerHeight - 4;

    window.onresize = () => {
      this._canvas.width = window.innerWidth - 2;
      this._canvas.height = window.innerHeight - 4;
      // TODO: redraw
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
      let engine = this.createEngine(Data.rb[index], lastX, this.view.y);
      this.engines.push(engine);

      lastX += engine.width;
    }
  }

  /**
   * @param {EngineRollenBahn} engine
   * @param {number} x
   * @param {number} y
   * @returns {EngineRollenBahn}
   */
  updateEngine(engine, x, y) {
    // TODO: removed and replaced (see your notes)
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

      if (this._viewChanged) {
        // TODO: if view changed => updatePosition

        this._viewChanged = false;
      }

      this.ctx.clearRect(0, 0, this._canvas.width, this._canvas.height);
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
