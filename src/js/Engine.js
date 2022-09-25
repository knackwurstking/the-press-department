export class EngineRollenBahn {
  /**
   * @param {{
   *  rolleLeft: HTMLImageElement,
   *  rolleRight: HTMLImageElement,
   *  rbGestellAluBlockLeft: HTMLImageElement,
   *  rbGestellAluBlockRight: HTMLImageElement,
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

      let rolle;
      if (this.side === "right") {
        rolle = this.assets.rolleRight;
      } else if (this.side === "left") {
        rolle = this.assets.rolleLeft;
      } else {
        continue;
      }

      let aluBlockLeft = this.assets.rbGestellAluBlockLeft;
      let aluBlockRight = this.assets.rbGestellAluBlockRight;
      aluBlockRight.style.transform = "rotate(180deg)";

      let posX = this.sX + index * 10;
      let sX = 6 * frameNumber;
      let sY = 0;
      let sWidth = 6;
      let sHeight = rolle.height;
      let dX = posX + 2;
      let dY = 8;
      let dWidth = 6;
      let dHeight = sHeight;

      ctx.drawImage(aluBlockLeft, posX, 0, 10, aluBlockRight.height);

      ctx.drawImage(
        aluBlockRight,
        posX,
        aluBlockLeft.height + rolle.height - dY / 2,
        10,
        aluBlockRight.height
      );

      ctx.drawImage(rolle, sX, sY, sWidth, sHeight, dX, dY, dWidth, dHeight);
    }
  }
}
