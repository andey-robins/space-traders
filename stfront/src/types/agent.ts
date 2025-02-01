import type { Base } from "./base";

export type AgentId = number;

export type Agent = {
    symbol: string;
    faction: string;
    token?: string;
} & Base