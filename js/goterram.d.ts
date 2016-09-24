declare module "@fusebot/goterram" {
  export interface ITerramGame {
    Start(): void;
    Stop(): void;
  }

  export interface ITerramJsGlobal {
    BuildTerramGame(frontend: IFrontend): ITerramGame;
  }

  export interface INetEntity {
    Id: Number;
    ParentId?: Number;
  }

  export interface IFrontendComponent {
    init?(): void;
    initLate?(): void;
    update?(): void;
    destroy?(): void;
  }

  export interface IFrontendEntity {
    // List of all components
    components: { [id: number]: IFrontendComponent };

    // Called when entity is created, no components
    init?(): void;
    // Called when adding components
    addComponent?(component_id: number): IFrontendComponent | void;
    // Called when all components have been added.
    initLate?(): void;
    // Called when the entity is removed.
    destroy?(): void;
  }

  export enum EGameOperatingMode {
    LOCAL = 0,
    REMOTE,
  }

  export var EGameOperatingModeString: { [id: number]: string };

  export interface IFrontendGameRules {
    init?(): void;
    setGameOperatingMode?(op_mode: EGameOperatingMode);
    destroy?(): void;
  }

  /* Frontend interface */
  export interface IFrontend {
    init?(): IFrontendGameRules | void;
    // Called when an entity is added.
    addEntity?(entity: INetEntity): IFrontendEntity | void;
    destroy?(): void;
  }

  export var TerramBuilder: ITerramJsGlobal;
}
