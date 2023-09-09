-- 連携ゲーム
CREATE TABLE `game` (
    `id`               BIGINT       NOT NULL AUTO_INCREMENT COMMENT "連携ゲームID",
    `game_key`         VARCHAR(20)  NOT NULL UNIQUE KEY     COMMENT "連携ゲームKEY",
    `genre_key`         VARCHAR(50)  NOT NULL                COMMENT "ジャンルKEY",
    `admin_user_key`   VARCHAR(20)  NOT NULL                COMMENT "企業ユーザーKEY",
    `api_key`          VARCHAR(20)  NOT NULL UNIQUE KEY     COMMENT "連携APIKEY",    
    `game_title`       VARCHAR(50)  NOT NULL                COMMENT "ゲームタイトル",
    `game_image_path`  VARCHAR(191) NOT NULL                COMMENT "画像パス",
    `deleted`          TINYINT      NOT NULL DEFAULT "0"    COMMENT "削除状態",
    `created_at`       TIMESTAMP    NOT NULL                COMMENT "作成日時",
    `updated_at`       TIMESTAMP    NOT NULL                COMMENT "更新日時",
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
