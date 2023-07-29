-- ルーム情報
CREATE TABLE `room` (
    `id` INT NOT NULL AUTO_INCREMENT COMMENT "ルームID",
    `room_key` VARCHAR(12) NOT NULL UNIQUE KEY COMMENT "ルームKey",
    `user_id` INT NOT NULL COMMENT "ユーザーID",
    `name` VARCHAR(50) NOT NULL COMMENT "ルーム名",
    `explanation` VARCHAR(50) NOT NULL COMMENT "説明",
    `image_path` VARCHAR(50) NOT NULL COMMENT "画像パス",
    `user_count` INT NOT NULL COMMENT "ユーザー数",
    `status` VARCHAR(255) NOT NULL COMMENT "状態",
    `created_at` TIMESTAMP NOT NULL COMMENT "作成日時",
    `updated_at` TIMESTAMP NOT NULL COMMENT "更新日時",
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
