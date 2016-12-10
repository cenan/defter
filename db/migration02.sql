CREATE TABLE IF NOT EXISTS `attachments` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `data` BLOB,
    `mime` VARCHAR(255),
    `version` INTEGER,
    `created_at` INTEGER
);
