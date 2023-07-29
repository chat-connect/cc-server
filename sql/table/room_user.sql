-- ルームのユーザー情報
CREATE TABLE `room_user` (
    `id` INT NOT NULL AUTO_INCREMENT COMMENT "ルームID",
    `room_user_key` VARCHAR(12) NOT NULL UNIQUE KEY COMMENT "ルームユーザーKey",
    `room_id` INT NOT NULL COMMENT "ルームID",
    `user_id` INT NOT NULL COMMENT "ユーザーID",
    `host`  TINYINT NOT NULL DEFAULT "0" COMMENT "ホストユーザー",
    `status` VARCHAR(255) NOT NULL COMMENT "状態",
    `created_at` TIMESTAMP NOT NULL COMMENT "作成日時",
    `updated_at` TIMESTAMP NOT NULL COMMENT "更新日時",
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
