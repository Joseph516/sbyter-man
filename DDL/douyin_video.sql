CREATE DATABASE IF NOT EXISTS douyin DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

USE douyin;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- douyin.douyin_video definition

CREATE TABLE `douyin_video` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '编号id',
  `author_id` int NOT NULL COMMENT '作者id',
  `play_url` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '视频播放地址',
  `cover_url` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '视频封面地址',
  `favorite_count` int DEFAULT '0' COMMENT '视频的点赞总数',
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '视频标题',
  `comment_count` int DEFAULT '0' COMMENT '视频的评论总数',
  `publish_date` timestamp NULL DEFAULT NULL COMMENT '发布时期',
  `created_at` timestamp DEFAULT NULL,
  `updated_at` timestamp DEFAULT NULL,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;