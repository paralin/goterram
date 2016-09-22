declare module "@fusebot/goterram" {
  let ITerramGame: any;
  declare interface ITerramJsGlobal {
    BuildTerramGame(frontend: IFrontend): ITerramGame;
  }

  declare interface INetEntity {
    id: Number;
    parent_id?: Number;
  }

  declare interface IFrontendComponent {
    init?(): void;
    destroy?(): void;
  }

  declare interface IFrontendEntity {
    // Called when entity is created, no components
    init?(): void;
    // Called when adding components
    addComponent?(component_id: number): IFrontendComponent | void;
    // Called when all components have been added.
    initLate?(): void;
    // Called when the entity is removed.
    destroy?(): void;
  }

  /* Frontend interface */
  declare interface IFrontend {
    init?(): void;
    // Called when an entity is added.
    addEntity?(entity: INetEntity): IFrontendEntity | void;
    destroy?(): void;
  }

  export = ITerramJsGlobal;
}
