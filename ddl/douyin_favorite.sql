CREATE TABLE IF NOT EXISTS `douyin_favorite`
(
    `id`          int NOT NULL AUTO_INCREMENT,
    `user_id`     int NOT NULL,
    `video_id`    int NOT NULL,
    `created_at`  int NOT NULL,
    `updated_at`  int NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `user_video_un` (`user_id` ASC, `video_id` ASC),
    KEY `video_id_index` (`video_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 11
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;