-- 連携ゲーム情報
CREATE TABLE `link_game` (
    `id`               BIGINT       NOT NULL AUTO_INCREMENT COMMENT "連携ゲームID",
    `link_game_key`    VARCHAR(20)  NOT NULL UNIQUE KEY     COMMENT "連携ゲームKEY",
    `api_key`          VARCHAR(20)  NOT NULL UNIQUE KEY     COMMENT "連携APIKEY",    
    `game_title`       VARCHAR(50)  NOT NULL                COMMENT "ゲームタイトル",
    `game_image_path`  VARCHAR(191) NOT NULL                COMMENT "画像パス",
    `game_genre`       VARCHAR(50)  NOT NULL                COMMENT "ゲームジャンル",
    `created_at`       TIMESTAMP    NOT NULL                COMMENT "作成日時",
    `updated_at`       TIMESTAMP    NOT NULL                COMMENT "更新日時",
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
