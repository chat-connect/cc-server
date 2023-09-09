-- 連携ゲーム設定
CREATE TABLE `game_setting` (
    `id`                    BIGINT      NOT NULL AUTO_INCREMENT COMMENT "連携ゲームID",
    `game_key`              VARCHAR(20) NOT NULL UNIQUE KEY     COMMENT "連携ゲームKEY",
    `admin_user_key`        VARCHAR(20) NOT NULL                COMMENT "企業ユーザーKEY",
    `game_score`            TINYINT     NOT NULL DEFAULT "0"    COMMENT "スコア",
    `game_combo_score`      TINYINT     NOT NULL DEFAULT "0"    COMMENT "コンボスコア",
    `game_rank`             TINYINT     NOT NULL DEFAULT "0"    COMMENT "ランク",
    `game_play_time`        TINYINT     NOT NULL DEFAULT "0"    COMMENT "プレイ時間",
    `game_score_image_path` TINYINT     NOT NULL DEFAULT "0"    COMMENT "画像パス",
    `deleted`               TINYINT      NOT NULL DEFAULT "0"    COMMENT "削除状態",
    `created_at`            TIMESTAMP   NOT NULL                COMMENT "作成日時",
    `updated_at`            TIMESTAMP   NOT NULL                COMMENT "更新日時",
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
