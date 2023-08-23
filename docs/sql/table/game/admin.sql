-- 企業
CREATE TABLE `admin_user` (
    `id`               BIGINT       NOT NULL AUTO_INCREMENT COMMENT "企業ID",
    `admin_user_key`   VARCHAR(20)  NOT NULL UNIQUE KEY     COMMENT "企業KEY",
    `name`             VARCHAR(50)  NOT NULL                COMMENT "企業名",
    `email`            VARCHAR(191) NOT NULL                COMMENT "メールアドレス",
    `password`         VARCHAR(191) NOT NULL                COMMENT "パスワード",
    `token`            VARCHAR(255) NOT NULL                COMMENT "アクセストークン",
    `status`           VARCHAR(191) NOT NULL                COMMENT "状態",
    `created_at`       TIMESTAMP    NOT NULL                COMMENT "作成日時",
    `updated_at`       TIMESTAMP    NOT NULL                COMMENT "更新日時",
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
