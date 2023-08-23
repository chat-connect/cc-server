-- ルームのユーザー
CREATE TABLE `room_user` (
    `id`            BIGINT       NOT NULL AUTO_INCREMENT COMMENT "ルームユーザーID",
    `room_user_key` VARCHAR(20)  NOT NULL UNIQUE KEY     COMMENT "ルームユーザーKEY",
    `room_key`      VARCHAR(20)  NOT NULL                COMMENT "ルームKEY",
    `user_key`      VARCHAR(20)  NOT NULL                COMMENT "ユーザーKEY",
    `host`          TINYINT      NOT NULL DEFAULT "0"    COMMENT "ホストユーザー",
    `status`        VARCHAR(191) NOT NULL                COMMENT "状態",
    `created_at`    TIMESTAMP    NOT NULL                COMMENT "作成日時",
    `updated_at`    TIMESTAMP    NOT NULL                COMMENT "更新日時",
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
