export default (function () {
  const data = {
    ROLLEN1: "rollen",
    RIEMEN1: "riemen",
  };

  data.rollenBahn = [
    {
      name: "Pressen Übergangs Tisch",
      engine: {
        count: 16,
        type: data.ROLLEN1,
      },
    },
    {
      name: "Wender",
      engine: {
        count: 26,
        type: data.ROLLEN1,
      },
    },
    {
      name: "Rollen Bahn 1",
      engine: {
        count: 22,
        type: data.ROLLEN1,
      },
    },
    {
      name: "Rollen Bahn 2",
      engine: {
        count: 22,
        type: data.ROLLEN1,
      },
    },
    {
      name: "Rollen Bahn 3",
      engine: {
        count: 22,
        type: data.ROLLEN1,
      },
    },
    {
      name: "Rollen Bahn 4",
      engine: {
        count: 22,
        type: data.ROLLEN1,
      },
    },
    {
      name: "Rollen Bahn 5",
      engine: {
        count: 14,
        type: data.ROLLEN1,
      },
    },
    {
      name: "Ausricter Eingang",
      engine: {
        count: 14,
        type: data.ROLLEN1,
      },
    },
    {
      name: "Trockner Eingang",
      engine: {
        count: 15, // NOTE: 2 rollen form TE sind speciell
        type: data.ROLLEN1,
      },
    },
  ];

  data.trockerAusgang = [
    {
      name: "Gestell Rollen1",
      engine: {
        count: 13, // NOTE: format 60x60
        type: data.ROLLEN1,
      },
    },
    {
      name: "Gestell Rollen 2",
      engine: {
        count: 13, // NOTE: format 60x60
        type: data.ROLLEN1,
      },
    },
  ];

  // Trocknerausgang Riemen
  data.tagr = {
    name: "Gestell Riemen",
    engine: {
      count: 6, // NOTE: 6 riemen
      type: data.RIEMEN1,
    },
  };

  // Proben Tish (P4)
  data.pt = {
    name: "Proben Tisch",
    engine: {
      count: 6, // NOTE: 6 riemen
      type: data.RIEMEN1,
    },
  };

  return data;
})();
