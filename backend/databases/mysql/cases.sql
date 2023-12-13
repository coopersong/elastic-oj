-- +migrate Up
CREATE TABLE `cases` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `case_id` VARCHAR(64) NOT NULL,
    `problem_id` VARCHAR(64) NOT NULL,
    `docs` VARCHAR(4096) NOT NULL,
    `version` bigint NOT NULL DEFAULT 0,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_problem_id` (`case_id`),
    KEY `index_created_at` (`created_at`),
    KEY `index_updated_at` (`updated_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE IF EXISTS `cases`;