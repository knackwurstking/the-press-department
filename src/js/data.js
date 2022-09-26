/**
 * @typedef Engine
 * @type {{
 *   count: number,
 *   type: "rollen"|"riemen",
 * }}
 */

export default (function () {
  const data = {
    ROLLEN_ROUND: 1,
    ROLLEN_GRIP: 2,
    RIEMEN_ROUND: 3,
    RIEMEN_GRIP: 4,

    /** 17m, 1700cm */
    rbLength: 1700,

    /** 1.5m, 150cm */
    rbWidth: 150,

    rbMaxCount: 173,

    /** 1.5m, 150cm -- 6 Rollen nicht sichtbar von oben (Eingang und Ausgang) */
    trockerLength1: 150,

    /** 1.6m, 160cm */
    trocknerAusgangLength: 160,

    /** Keep it simple for now */
    glRiemenCount: 4,
  };

  data.assets = [
    {
      name: "rbAluBlockLeft",
      width: 20,
      height: 10,
      src: "assets/RollenBahnAluBlockLeft_20x10.png",
    },
    {
      name: "rbAluBlockRight",
      width: 20,
      height: 10,
      src: "assets/RollenBahnAluBlockRight_20x10.png",
    },
    {
      name: "rolleLeft",
      width: 6 * 6,
      height: 296,
      src: "assets/RolleLeft_6x296.png",
    },
    {
      name: "rolleRight",
      width: 6 * 6,
      height: 296,
      src: "assets/RolleRight_6x296.png",
    },
    {
      name: "rbRiemen290x5",
      width: 290,
      height: 5 * 3,
      src: "assets/Riemen-290x5.png",
    },
    {
      name: "rbRiemen270x5",
      width: 290,
      height: 5 * 3,
      src: "assets/Riemen-270x5.png",
    },
  ];

  data.rb = [
    {
      name: "Trockner Eingang",
      engine: {
        /** 2 Rollen for dem Trocker Eingang sind speziell */
        count: 15,
        type: data.ROLLEN_GRIP,
        side: "left",
      },
    },
    {
      name: "Ausrichter Eingang",
      engine: {
        count: 14,
        type: data.ROLLEN_GRIP,
        side: "right",
      },
    },
    {
      name: "Rollen Bahn 5",
      engine: {
        count: 14,
        type: data.ROLLEN_GRIP,
        side: "left",
      },
    },
    {
      name: "Rollen Bahn 4",
      engine: {
        count: 22,
        type: data.ROLLEN_GRIP,
        side: "right",
      },
    },
    {
      name: "Rollen Bahn 3",
      engine: {
        count: 22,
        type: data.ROLLEN_GRIP,
        side: "left",
      },
    },
    {
      name: "Rollen Bahn 2",
      engine: {
        count: 22,
        type: data.ROLLEN_GRIP,
        side: "right",
      },
    },
    {
      name: "Rollen Bahn 1",
      engine: {
        count: 22,
        type: data.ROLLEN_GRIP,
        side: "left",
      },
    },
    {
      name: "Wender",
      engine: {
        count: 26,
        type: data.ROLLEN_GRIP,
        side: "right",
      },
    },
    {
      name: "Pressen Übergangs Tisch",
      engine: {
        count: 16,
        type: data.ROLLEN_GRIP,
        side: "right",
      },
    },
  ];

  data.ta = [
    {
      name: "Gestell Rollen1",
      engine: {
        /** Format 60x60 -- Der Antriebsriemen ist ca. 1 Rollen lang */
        count: 13,
        type: data.ROLLEN_GRIP,
      },
    },
    {
      name: "Gestell Rollen 2",
      engine: {
        /** Format 60x60 -- Der Antriebsriemen ist ca. 11 Rollen lang */
        count: 9,
        type: data.ROLLEN_GRIP,
      },
    },
  ];

  // Trocknerausgang Riemen
  data.trockerAusgangRiemen = {
    name: "Gestell Riemen",
    engine: {
      count: 6,
      type: data.RIEMEN_ROUND,
    },
  };

  // Proben Tisch (P4)
  data.pressenTisch = {
    name: "Proben Tisch",
    engine: {
      count: 6,
      type: data.RIEMEN_GRIP,
    },
  };

  return data;
})();
