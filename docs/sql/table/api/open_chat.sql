-- オープンチャット
CREATE TABLE `open_chat` (
    `id`            BIGINT       NOT NULL AUTO_INCREMENT COMMENT "オープンチャットID",
    `open_chat_key` VARCHAR(20)  NOT NULL UNIQUE KEY     COMMENT "オープンチャットKEY",
    `user_key`      VARCHAR(20)  NOT NULL                COMMENT "ユーザーKEY",
    `user_name`     VARCHAR(50)  NOT NULL                COMMENT "ユーザー名",
    `content`       VARCHAR(191) NOT NULL                COMMENT "コンテンツ",
    `image_path`    VARCHAR(191) NOT NULL                COMMENT "画像パス",
    `posted_at`     TIMESTAMP    NOT NULL                COMMENT "投稿日時",
    `created_at`    TIMESTAMP    NOT NULL                COMMENT "作成日時",
    `updated_at`    TIMESTAMP    NOT NULL                COMMENT "更新日時",
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
