-- 連携ユーザー
CREATE TABLE `game_user` (
    `id`            BIGINT       NOT NULL AUTO_INCREMENT COMMENT "ゲームユーザーID",
    `game_user_key` VARCHAR(20)  NOT NULL UNIQUE KEY     COMMENT "ゲームユーザーKEY",
    `user_key`      VARCHAR(20)  NOT NULL                COMMENT "ユーザーKEY",
    `game_key`      VARCHAR(20)  NOT NULL                COMMENT "連携ゲームKEY",
    `deleted_at`    TIMESTAMP    DEFAULT NULL            COMMENT "削除日時",
    `created_at`    TIMESTAMP    NOT NULL                COMMENT "作成日時",
    `updated_at`    TIMESTAMP    NOT NULL                COMMENT "更新日時",
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
