-- ゲーム情報
CREATE TABLE `game` (
    `id`          BIGINT       NOT NULL AUTO_INCREMENT COMMENT "ゲームID",
    `game_key`    VARCHAR(20)  NOT NULL UNIQUE KEY     COMMENT "ジャンルKEY",
    `genre_key`   VARCHAR(20)  NOT NULL                COMMENT "ゲームKEY",
    `name`        VARCHAR(50)  NOT NULL                COMMENT "ゲーム名",
    `description` VARCHAR(191) NOT NULL                COMMENT "説明",
    `type`        VARCHAR(191) NOT NULL                COMMENT "種別",
    `created_at`  TIMESTAMP    NOT NULL                COMMENT "作成日時",
    `updated_at`  TIMESTAMP    NOT NULL                COMMENT "更新日時",
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
