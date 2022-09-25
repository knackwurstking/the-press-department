/**
 * @typedef Engine
 * @type {{
 *   count: number,
 *   type: "rollen"|"riemen",
 * }}
 */

export default (function () {
  const data = {
    ROLLEN_1: "rollen",
    RIEMEN_1: "riemen",

    /** 17m, 1700cm */
    ROLLEN_BAHN_LENGTH: 1700,

    /** 1.5m, 150cm */
    ROLLEN_BAHN_WIDTH: 150,

    ROLLEN_BAHN_MAX_COUNT: 173,

    /** 1.5m, 150cm -- 6 Rollen nicht sichtbar von oben (Eingang und Ausgang) */
    TROCKER_LENGTH_1: 150,

    /** Make a cut here */
    TROCKER_LENGTH_2: -1,

    /** 1.6m, 160cm */
    TROCKNER_AUSGANG_LENGTH: 160,

    /** Keep it simple for now */
    GLASUR_RIEMEN: 4,
  };

  data.rollenBahn = [
    {
      name: "Pressen Übergangs Tisch",
      engine: {
        count: 16,
        type: data.ROLLEN_1,
      },
    },
    {
      name: "Wender",
      engine: {
        count: 26,
        type: data.ROLLEN_1,
      },
    },
    {
      name: "Rollen Bahn 1",
      engine: {
        count: 22,
        type: data.ROLLEN_1,
      },
    },
    {
      name: "Rollen Bahn 2",
      engine: {
        count: 22,
        type: data.ROLLEN_1,
      },
    },
    {
      name: "Rollen Bahn 3",
      engine: {
        count: 22,
        type: data.ROLLEN_1,
      },
    },
    {
      name: "Rollen Bahn 4",
      engine: {
        count: 22,
        type: data.ROLLEN_1,
      },
    },
    {
      name: "Rollen Bahn 5",
      engine: {
        count: 14,
        type: data.ROLLEN_1,
      },
    },
    {
      name: "Ausricter Eingang",
      engine: {
        count: 14,
        type: data.ROLLEN_1,
      },
    },
    {
      name: "Trockner Eingang",
      engine: {
        /** 2 Rollen for dem Trocker Eingang sind speziell */
        count: 15,
        type: data.ROLLEN_1,
      },
    },
  ];

  data.trockerAusgang = [
    {
      name: "Gestell Rollen1",
      engine: {
        /** Format 60x60 -- Der Antriebsriemen ist ca. 1 Rollen lang */
        count: 13,
        type: data.ROLLEN_1,
      },
    },
    {
      name: "Gestell Rollen 2",
      engine: {
        /** Format 60x60 -- Der Antriebsriemen ist ca. 11 Rollen lang */
        count: 9,
        type: data.ROLLEN_1,
      },
    },
  ];

  // Trocknerausgang Riemen
  data.trockerAusgangRiemen = {
    name: "Gestell Riemen",
    engine: {
      count: 6,
      type: data.RIEMEN_1,
    },
  };

  // Proben Tisch (P4)
  data.pressenTisch = {
    name: "Proben Tisch",
    engine: {
      count: 6,
      type: data.RIEMEN_1,
    },
  };

  return data;
})();
