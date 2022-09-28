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
    this.engines;

    // touch event handlers
    this.pointer = false;

    /** @type {null|((ev: PointerEvent) => any)} */
    this.onpointerdown = null;
    /** @type {null|((ev: PointerEvent) => any)} */
    this._pointerdown = () => {
      console.log(`[DEBUG] pointerdown`);
      this.pointer = true;
    };

    /** @type {null|((ev: PointerEvent) => any)} */
    this.onpointermove = null;
    /** @type {null|((ev: PointerEvent) => any)} */
    this._pointermove = () => {
      if (!this.pointer) return;
      console.log(`[DEBUG] pointermove`);
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
  }

  initialize() {
    this.canvas.width = window.innerWidth - 2;
    this.canvas.height = window.innerHeight - 4;
    window.onresize = () => {
      this.initialize();
    };

    this.engines = []; // left to right

    let lastX = 0;
    for (let section of Data.rb) {
      let x = lastX;
      let y = this.canvas.height / 2 - 312 / 2;
      let width = section.engine.count * this.assets.rbAluBlockLeft.width;
      let height = this.canvas.height;
      lastX += width;

      this.engines.push(
        new EngineRollenBahn(
          section.name,
          section.engine.side,
          section.engine.count,
          x,
          y,
          width,
          height,
          this.assets
        )
      );
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

  async start() {
    this.initialize();
    this.handleUserInput();
    const animate = (/** @type {number} */ frame) => {
      this.draw(frame);
      requestAnimationFrame(animate);
    };

    animate(0);
  }
}
