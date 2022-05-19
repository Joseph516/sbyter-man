CREATE TABLE `douyin_comment` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int  NOT NULL,
  `video_id` int  NOT NULL,
  `content` varchar(100) NOT NULL,
  `created_at` int NOT NULL,
  `updated_at` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;