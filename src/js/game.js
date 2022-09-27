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
    this.touchmove = false;

    /** @type {null|((ev: TouchEvent) => any)} */
    this.ontouchstart = null;
    /** @type {null|((ev: TouchEvent) => any)} */
    this._touchstart = () => {
      //console.log(`[DEBUG] canvas: "touchstart" event`);
      this.touchmove = true;
    };

    /** @type {null|((ev: TouchEvent) => any)} */
    this.ontouchmove = null;
    /** @type {null|((ev: TouchEvent) => any)} */
    this._touchmove = () => {
      if (!this.touchmove) return;
      //console.log(`[DEBUG] canvas: "touchmove" event`);
    };

    /** @type {null|((ev: TouchEvent) => any)} */
    this.ontouchend = null;
    /** @type {null|((ev: TouchEvent) => any)} */
    this._touchend = () => {
      //console.log(`[DEBUG] canvas: "touchend" event`);
      this.touchmove = false;
    };

    /** @type {null|((ev: TouchEvent) => any)} */
    this.ontouchcancel = null;
    /** @type {null|((ev: TouchEvent) => any)} */
    this._touchcancel = (ev) => {
      if (!this.touchmove) return;
      //console.log(`[DEBUG] canvas: "touchcancel" event`);
      if (this._touchend) this._touchend(ev);
    };

    // mouse event handlers
    this.mousemove = false;

    /** @type {null|((ev: MouseEvent) => any)} */
    this.onmousedown = null;
    /** @type {null|((ev: MouseEvent) => any)} */
    this._mousedown = () => {
      //console.log(`[DEBUG] canvas: "mousedown" event`);
      this.mousemove = true;
    };

    /** @type {null|((ev: MouseEvent) => any)} */
    this.onmousemove = null;
    /** @type {null|((ev: MouseEvent) => any)} */
    this._mousemove = () => {
      if (!this.mousemove) return;
      //console.log(`[DEBUG] canvas: "mousemove" event`);
    };

    /** @type {null|((ev: MouseEvent) => any)} */
    this.onmouseup = null;
    /** @type {null|((ev: MouseEvent) => any)} */
    this._mouseup = () => {
      //console.log(`[DEBUG] canvas: "mouseup" event`);
      this.mousemove = false;
    };

    /** @type {null|((ev: MouseEvent) => any)} */
    this.onmouseover = null;
    /** @type {null|((ev: MouseEvent) => any)} */
    this._mouseover = null;

    /** @type {null|((ev: MouseEvent) => any)} */
    this.onmouseout = null;
    /** @type {null|((ev: MouseEvent) => any)} */
    this._mouseout = (ev) => {
      if (!this.mousemove) return;
      //console.log(`[DEBUG] canvas: "mouseout" event`);
      if (this._mouseup) this._mouseup(ev);
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
    // handle touch events
    this.canvas.ontouchstart = this._touchstart;
    this.canvas.ontouchmove = this._touchmove;
    this.canvas.ontouchend = this._touchend;
    this.canvas.ontouchcancel = this._touchcancel;
    // handle mouse events
    this.canvas.onmousedown = this._mousedown;
    this.canvas.onmousemove = this._mousemove;
    this.canvas.onmouseup = this._mouseup;
    this.canvas.onmouseover = this._mouseover;
    this.canvas.onmouseout = this._mouseout;
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
