-- ジャンル情報
CREATE TABLE `genre` (
    `id`          BIGINT       NOT NULL AUTO_INCREMENT COMMENT "ジャンルID",
    `genre_key`    VARCHAR(20)  NOT NULL                COMMENT "ジャンルKEY",
    `name`        VARCHAR(50)  NOT NULL                COMMENT "ジャンル名",
    `description` VARCHAR(191) NOT NULL                COMMENT "説明",
    `type`        VARCHAR(191) NOT NULL                COMMENT "種別",
    `created_at`  TIMESTAMP    NOT NULL                COMMENT "作成日時",
    `updated_at`  TIMESTAMP    NOT NULL                COMMENT "更新日時",
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
