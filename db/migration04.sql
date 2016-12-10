CREATE TABLE `notebooks` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `name` VARCHAR(255) NULL
);

ALTER TABLE `pages` ADD `notebook_id` INTEGER;
