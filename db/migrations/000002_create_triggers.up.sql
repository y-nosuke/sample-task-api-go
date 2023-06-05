CREATE TRIGGER IF NOT EXISTS `on_insert_tasks`
AFTER
INSERT
    ON `tasks` FOR EACH ROW
BEGIN
    INSERT INTO
        `task_registrations` (
            `task_id`,
            `title`,
            `detail`,
            `deadline`,
            `created_by`,
            `created_at`
        )
    VALUES
        (
            NEW.`id`,
            NEW.`title`,
            NEW.`detail`,
            NEW.`deadline`,
            NEW.`created_by`,
            NEW.`created_at`
        );
END;

CREATE TRIGGER IF NOT EXISTS `on_update_tasks`
AFTER
UPDATE
    ON `tasks` FOR EACH ROW
BEGIN
    IF NEW.`completed` = OLD.`completed` THEN
        INSERT INTO
            `task_updates` (
                `task_id`,
                `title`,
                `detail`,
                `deadline`,
                `created_by`,
                `created_at`
            )
        VALUES
            (
                NEW.`id`,
                NEW.`title`,
                NEW.`detail`,
                NEW.`deadline`,
                NEW.`updated_by`,
                NEW.`updated_at`
            );
    END IF;

    IF OLD.`completed` = false AND NEW.`completed` = true THEN
        INSERT INTO
            `task_completions` (
                `task_id`,
                `created_by`,
                `created_at`
            )
        VALUES
            (
                NEW.`id`,
                NEW.`updated_by`,
                NEW.`updated_at`
            );
    END IF;

    IF OLD.`completed` = true AND NEW.`completed` = false THEN
        INSERT INTO
            `task_uncompletions` (
                `task_id`,
                `created_by`,
                `created_at`
            )
        VALUES
            (
                NEW.`id`,
                NEW.`updated_by`,
                NEW.`updated_at`
            );
    END IF;
END;
