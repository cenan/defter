CREATE TABLE IF NOT EXISTS `pages` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `title` VARCHAR(255) NULL,
    `content` TEXT,
    `updated_at` INTEGER
);
