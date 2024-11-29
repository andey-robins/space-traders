-- Create "Starship" table
CREATE TABLE `Starship` (
    `Id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    `Name` text NOT NULL
);
-- Create "Agent" table
CREATE TABLE `Agent` (
    `Id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    `Symbol` text NOT NULL,
    `Faction` text NOT NULL,
    `Token` text NULL
);
-- Create "Config" table
CREATE TABLE `Config` (
    `Id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    `Key` text NOT NULL,
    `Value` text NOT NULL
);