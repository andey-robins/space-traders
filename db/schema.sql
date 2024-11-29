-- Create "ConfigData" table
CREATE TABLE `ConfigData` (
    `Id` integer NOT NULL,
    `Key` text NOT NULL,
    `Value` text NOT NULL,
    PRIMARY KEY (`Id`)
);
-- Create "Starship" table
CREATE TABLE `Starship` (
    `Id` integer NOT NULL,
    `Name` text NOT NULL,
    PRIMARY KEY (`Id`)
);