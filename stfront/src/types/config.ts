import type { Base } from "./base";

export type ConfigId = number;

export type Config = {
    key: string;
    value: string;
} & Base