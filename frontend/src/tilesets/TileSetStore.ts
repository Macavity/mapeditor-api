import { defineStore } from "pinia";
import { TileSetFactory } from "@/tilesets/TileSetFactory";

export const useTileSetStore = defineStore({
  id: "tileSetStore",
  state: () => ({
    tileSetEntries: [] as ITileSet[],
  }),
  getters: {
    tileSets: (state) => state.tileSetEntries,
  },
  actions: {
    loadTileSets() {
      this.tileSets.push(
        TileSetFactory.create(
          "000-Types",
          "tilesets/000-Types.png",
          1,
          9,
          32,
          32,
          96,
          96,
          0,
          0
        )
      );
      this.tileSets.push(
        TileSetFactory.create(
          "mage-city",
          "tilesets/mage-city.png",
          10,
          368,
          32,
          32,
          1450,
          256,
          0,
          0
        )
      );
      this.tileSets.push(
        TileSetFactory.create(
          "castle_exterior_mc",
          "tilesets/castle_exterior_mc.png",
          378,
          1544,
          32,
          32,
          6176,
          256,
          0,
          0
        )
      );
    },
  },
});
