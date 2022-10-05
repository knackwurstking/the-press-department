export class Engine {
  /**
   * @param {string} name
   * @param {"left"|"right"} side
   * @param {number} count
   * @param {number} x
   * @param {number} y
   * @param {number} width
   * @param {number} height
   * @param {import("./game").Assets} assets
   */
  constructor(name, side, count, x, y, width, height, assets) {
    this.name = name;
    this.assets = assets;
    this.side = side;
    this.count = count;
    this.x = x;
    this.y = y;
    this.width = width;
    this.height = height;

    /** @type {number} */
    this._frameNumber = 0;

    this.hz = 12;
    this.lastFrame = 0 - 600 / this.hz;
  }

  /**
   * @param {CanvasRenderingContext2D} ctx
   * @param {number} dX
   */
  drawRiemen(ctx, dX) {
    // NOTE: ${count - aluBlock.width - padding-left/right}
    const image = this.assets[`rbRiemen${this.count * 20 - 10}x5`];
    if (image) {
      ctx.drawImage(
        image,
        0, // sX
        5 * (this._frameNumber % 3), // sY
        image.width, // sWidth
        image.height / 3, // sHeight
        this.x + dX,
        this.side === "left"
          ? this.y + 11
          : this.y + (this.assets.rolleLeft.height + (10 - 4 - 5)), // dY
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
    ctx.drawImage(image, this.x + dX, this.y, image.width, image.height);
  }

  /**
   * @param {CanvasRenderingContext2D} ctx
   * @param {number} dX
   */
  drawAluBlockRight(ctx, dX) {
    let image = this.assets.rbAluBlockRight;
    ctx.drawImage(
      image,
      this.x + dX,
      this.y +
        (this.assets.rbAluBlockLeft.height +
          (this.assets.rolleLeft.height - 4)),
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
    let dY = this.y + 8;
    let dWidth = rolle.width / 6;
    let dHeight = rolle.height;

    ctx.drawImage(
      rolle,
      sX,
      sY,
      sWidth,
      sHeight,
      this.x + dX,
      dY,
      dWidth,
      dHeight
    );
  }

  /**
   * @param {CanvasRenderingContext2D} ctx
   * @param {number} frame
   */
  draw(ctx, frame) {
    if (frame - this.lastFrame >= (600 / this.hz)) {
      this.lastFrame = frame;
      let backup = this._frameNumber;

      try {
        this.drawRiemen(ctx, 5);
      } catch (error) {
        this._frameNumber = backup;
      }

      let index = -1;
      for (let x = 0; x < this.count; x++) {
        index += 1;

        let posX = index * this.assets.rbAluBlockLeft.width;
        this.drawAluBlockLeft(ctx, posX);
        this.drawAluBlockRight(ctx, posX);
        this.drawRolle(ctx, posX + 7);
      }

      this._frameNumber += 1;
      if (this._frameNumber > 5) this._frameNumber = 0
    }
  }
}
