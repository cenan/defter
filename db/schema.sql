CREATE TABLE `pages` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `title` VARCHAR(255) NULL,
    `content` TEXT,
	`notebook_id` INTEGER,
    `updated_at` INTEGER
);

CREATE TABLE `attachments` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `data` BLOB,
    `mime` VARCHAR(255),
    `version` INTEGER,
    `created_at` INTEGER
);

CREATE TABLE `page_attachments` (
    `page_id` INTEGER,
    `attachment_id` INTEGER
);

CREATE TABLE `notebooks` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `name` VARCHAR(255) NULL
);
