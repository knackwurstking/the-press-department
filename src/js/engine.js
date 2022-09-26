export class EngineRollenBahn {
  /**
   * @param {import("./game").Assets} assets
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

    /** @type {number} */
    this._frameNumber;
  }

  /**
   * @param {CanvasRenderingContext2D} ctx
   * @param {number} dX
   */
  drawRiemen(ctx, dX) {
    const image = this.assets[`rbRiemen${this.count * 10}x5`];
    if (image) {
      ctx.drawImage(
        image,
        0, // sX
        5 * (2 - (this._frameNumber % 3)), // sY: backwards
        image.width, // sWidth
        image.height / 3, // sHeight
        dX,
        this.side === "left" ? 11 : this.assets.rolleLeft.height + (10 - 4 - 5), // dY
        image.width, // dWidth
        5 // dHeight
      );
    } else {
      throw `missing assets: "rbRiemen${this.count * 10}x5"`;
    }
  }

  /**
   * @param {CanvasRenderingContext2D} ctx
   * @param {number} dX
   */
  drawAluBlockLeft(ctx, dX) {
    let image = this.assets.rbAluBlockLeft;
    ctx.drawImage(image, dX, 0, image.width, image.height);
  }

  /**
   * @param {CanvasRenderingContext2D} ctx
   * @param {number} dX
   */
  drawAluBlockRight(ctx, dX) {
    let image = this.assets.rbAluBlockRight;
    ctx.drawImage(
      image,
      dX,
      this.assets.rbAluBlockLeft.height + (this.assets.rolleLeft.height - 4),
      image.width,
      image.height
    );
  }

  /**
   * @param {CanvasRenderingContext2D} ctx
   * @param {number} dX
   */
  drawRolle(ctx, dX) {
    let rolle;
    if (this.side === "left") {
      rolle = this.assets.rolleLeft;
    } else {
      rolle = this.assets.rolleRight;
    }

    let sX = 6 * this._frameNumber;
    let sY = 0;
    let sWidth = 6;
    let sHeight = rolle.height;
    let dY = 8;
    let dWidth = rolle.width / 6;
    let dHeight = rolle.height;

    ctx.drawImage(rolle, sX, sY, sWidth, sHeight, dX, dY, dWidth, dHeight);
  }

  /**
   * @param {CanvasRenderingContext2D} ctx
   * @param {number} frameNumber - draw a frame (1-6)
   */
  draw(ctx, frameNumber) {
    this._frameNumber = frameNumber;

    this.drawRiemen(ctx, this.sX + 5);

    let index = -1;
    for (let x = 0; x < this.count; x++) {
      index += 1;

      let posX = this.sX + index * this.assets.rbAluBlockLeft.width;
      this.drawAluBlockLeft(ctx, posX);
      this.drawAluBlockRight(ctx, posX);
      this.drawRolle(ctx, posX + 7);
    }
  }
}
