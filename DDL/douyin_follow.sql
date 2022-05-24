-- ----------------------------
-- Table structure for douyin_follow
-- ----------------------------
DROP TABLE IF EXISTS `douyin_follow`;
CREATE TABLE `douyin_follow`
(
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `followed_id` int NULL DEFAULT NULL,
  `follower_id` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  FOREIGN KEY (`followed_id`) REFERENCES `douyin_user`(`id`),
  FOREIGN KEY (`follower_id`) REFERENCES `douyin_user`(`id`),
  UNIQUE INDEX `uniq_idx`(`follower_id` , `followed_id`) USING BTREE,
  INDEX `idx_douyin_follow_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;