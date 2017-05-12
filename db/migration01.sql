CREATE TABLE IF NOT EXISTS `pages` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `title` VARCHAR(255) NULL,
    `notebook_id` INTEGER,
    `content` TEXT,
    `updated_at` INTEGER
);
