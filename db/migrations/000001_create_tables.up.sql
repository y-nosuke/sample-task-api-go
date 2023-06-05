CREATE TABLE IF NOT EXISTS `tasks` (
	`id` BINARY(16) NOT NULL,
	`title` VARCHAR(255) NOT NULL,
	`detail` TEXT NULL,
	`completed` BOOLEAN NOT NULL,
	`deadline` DATE NULL,
	`created_by` BINARY(16) NOT NULL,
	`created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
	`updated_by` BINARY(16) NOT NULL,
	`updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
	`version` BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID())),
	PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `task_registrations` (
	`task_id` BINARY(16) NOT NULL,
	`title` VARCHAR(255) NOT NULL,
	`detail` TEXT NULL,
	`deadline` DATE NULL,
	`created_by` BINARY(16) NOT NULL,
	`created_at` DATETIME(6) NOT NULL,
	PRIMARY KEY (`task_id`),
	FOREIGN KEY `FK_task_registrations_tasks`(`task_id`) REFERENCES tasks(id) ON DELETE CASCADE
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `task_updates` (
	`task_id` BINARY(16) NOT NULL,
	`title` VARCHAR(255) NOT NULL,
	`detail` TEXT NULL,
	`deadline` DATE NULL,
	`created_by` BINARY(16) NOT NULL,
	`created_at` DATETIME(6) NOT NULL,
	PRIMARY KEY (`task_id`, `created_at`),
	FOREIGN KEY `FK_task_updates_tasks`(`task_id`) REFERENCES tasks(id) ON DELETE CASCADE
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `task_completions` (
	`task_id` BINARY(16) NOT NULL,
	`created_by` BINARY(16) NOT NULL,
	`created_at` DATETIME(6) NOT NULL,
	PRIMARY KEY (`task_id`, `created_at`),
	FOREIGN KEY `FK_task_completions_tasks`(`task_id`) REFERENCES tasks(id) ON DELETE CASCADE
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `task_uncompletions` (
	`task_id` BINARY(16) NOT NULL,
	`created_by` BINARY(16) NOT NULL,
	`created_at` DATETIME(6) NOT NULL,
	PRIMARY KEY (`task_id`, `created_at`),
	FOREIGN KEY `FK_task_uncompletions_tasks`(`task_id`) REFERENCES tasks(id) ON DELETE CASCADE
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;