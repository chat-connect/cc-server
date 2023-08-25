-- フォロー
CREATE TABLE `follow` (
    `id`                 BIGINT       NOT NULL AUTO_INCREMENT COMMENT "フォローID",
    `follow_key`         VARCHAR(20)  NOT NULL UNIQUE KEY     COMMENT "フォローKEY",
    `user_key`           VARCHAR(50)  NOT NULL                COMMENT "ユーザーKEY",
    `following_user_key` VARCHAR(50)  NOT NULL                COMMENT "フォローされているユーザーKEY",
    `mutual`             TINYINT      NOT NULL DEFAULT "0"    COMMENT "相互状態",
    `mutual_follow_key`  VARCHAR(20)  NOT NULL                COMMENT "相互フォローKEY",
    `created_at`         TIMESTAMP    NOT NULL                COMMENT "作成日時",
    `updated_at`         TIMESTAMP    NOT NULL                COMMENT "更新日時",
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
