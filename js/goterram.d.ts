declare module "@fusebot/goterram" {
  let ITerramGame: any;
  declare interface ITerramJsGlobal {
    BuildTerramGame(): ITerramGame;
  }

  export = ITerramJsGlobal;
}
