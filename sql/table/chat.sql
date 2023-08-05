-- チャット情報
CREATE TABLE `chat` (
    `id`          BIGINT       NOT NULL AUTO_INCREMENT COMMENT "チャットID",
    `chat_key`    VARCHAR(20)  NOT NULL UNIQUE KEY     COMMENT "チャットKEY",
    `channel_key` VARCHAR(20)  NOT NULL                COMMENT "チャンネルKEY",
    `user_key`    VARCHAR(20)  NOT NULL                COMMENT "ユーザーKEY",
    `user_name`   VARCHAR(50)  NOT NULL                COMMENT "ユーザー名",
    `content`     VARCHAR(191) NOT NULL                COMMENT "コンテンツ",
    `created_at`  TIMESTAMP    NOT NULL                COMMENT "作成日時",
    `updated_at`  TIMESTAMP    NOT NULL                COMMENT "更新日時",
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
