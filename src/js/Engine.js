export class EngineRollenBahn {
  /**
   * @param {{
   *  rolleLeft: HTMLImageElement,
   *  rolleRight: HTMLImageElement,
   * }} assets
   * @param {"left"|"right"} side
   * @param {number} count
   * @param {number} sX
   * @param {number} sY
   * @param {number} width
   * @param {number} height
   */
  constructor(assets, side, count, sX, sY, width, height) {
    this.assets = assets;
    this.side = side;
    this.count = count;
    this.sX = sX;
    this.sY = sY;
    this.width = width;
    this.height = height;
  }

  /**
   * @param {CanvasRenderingContext2D} ctx
   * @param {number} frameNumber - draw a frame (1-6)
   */
  draw(ctx, frameNumber) {
    let index = -1;
    for (let x = 0; x < this.count; x++) {
      index += 1;

      let posX = this.sX + index * 10;

      let image;
      if (this.side === "right") {
        image = this.assets.rolleRight;
      } else if (this.side === "left") {
        image = this.assets.rolleLeft;
      } else {
        continue;
      }

      ctx.drawImage(
        image,
        // game asset
        6 * frameNumber,
        0,
        6,
        image.height,
        // canvas position
        posX,
        2,
        6,
        image.height
      );
    }
  }
}
