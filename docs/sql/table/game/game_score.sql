-- 連携ゲームスコア
CREATE TABLE `game_score` (
    `id`                    BIGINT       NOT NULL AUTO_INCREMENT COMMENT "連携ゲームID",
    `game_score_key`        VARCHAR(20)  NOT NULL                COMMENT "連携ゲームスコアKEY",
    `game_key`              VARCHAR(20)  NOT NULL                COMMENT "連携ゲームKEY",
    `user_key`              VARCHAR(20)  NOT NULL                COMMENT "ユーザーKEY",
    `game_score`            INT          NOT NULL                COMMENT "スコア",
    `game_combo_score`      INT          NOT NULL                COMMENT "コンボスコア",
    `game_rank`             INT          NOT NULL                COMMENT "ランク",
    `game_play_time`        INT          NOT NULL                COMMENT "プレイ時間",
    `game_score_image_path` VARCHAR(191) NOT NULL                COMMENT "画像パス",
    `created_at`            TIMESTAMP    NOT NULL                COMMENT "作成日時",
    `updated_at`            TIMESTAMP    NOT NULL                COMMENT "更新日時",
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
