CREATE TABLE `tasks` (
	`id` BINARY(16) NOT NULL,
	`title` VARCHAR(255) NOT NULL,
	`detail` TEXT NULL,
	`completed` BOOLEAN NOT NULL,
	`deadline` DATE NULL,
	`created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
	`updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
	`version` BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID())),
	PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;