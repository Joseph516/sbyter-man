-- douyin.douyin_user definition

CREATE TABLE `douyin_user`
(
    `id`             int          NOT NULL AUTO_INCREMENT,
    `user_name`      varchar(100) NOT NULL,
    `password`       varchar(100) NOT NULL,
    `follow_count`   int          NOT NULL,
    `follower_count` int          NOT NULL,
    `created_at`     int          NOT NULL,
    `updated_at`     int          NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `douyin_user_UN` (`user_name`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
