-- MySQL dump 10.13  Distrib 5.5.62, for Win64 (AMD64)
--
-- Host: 106.14.89.192    Database: douyin
-- ------------------------------------------------------
-- Server version	8.0.26

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `douyin_comment`
--

DROP TABLE IF EXISTS `douyin_comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `douyin_comment` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `video_id` int NOT NULL,
  `content` varchar(100) NOT NULL,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `douyin_comment`
--

LOCK TABLES `douyin_comment` WRITE;
/*!40000 ALTER TABLE `douyin_comment` DISABLE KEYS */;
INSERT INTO `douyin_comment` VALUES (24,45,35,'Hello！','2022-05-29 07:02:21','2022-05-29 07:02:21',NULL),(25,39,28,'啊啊啊','2022-05-31 08:37:45','2022-05-31 08:37:45','2022-06-04 14:39:11'),(26,39,28,'为什么还是0评论','2022-06-04 14:38:54','2022-06-04 14:38:54','2022-06-04 14:39:13'),(27,39,28,'为什么还是0评论','2022-06-04 14:38:55','2022-06-04 14:38:56','2022-06-04 14:39:14'),(28,39,28,'为什么还是0评论','2022-06-04 14:38:56','2022-06-04 14:38:56','2022-06-04 14:39:16'),(29,39,28,'为什么还是0评论','2022-06-04 14:38:57','2022-06-04 14:38:57','2022-06-04 14:39:18'),(30,39,28,'评论为-1','2022-06-04 14:39:39','2022-06-04 14:39:39',NULL),(31,39,29,'啊','2022-06-04 14:40:55','2022-06-04 14:40:55',NULL),(32,39,29,'嗷嗷','2022-06-04 14:41:10','2022-06-04 14:41:10',NULL),(33,39,29,'嗷嗷嗷','2022-06-04 14:44:22','2022-06-04 14:44:22',NULL);
/*!40000 ALTER TABLE `douyin_comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `douyin_favorite`
--

DROP TABLE IF EXISTS `douyin_favorite`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `douyin_favorite` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `video_id` int NOT NULL,
  `created_at` int NOT NULL,
  `updated_at` int NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_video_un` (`user_id`,`video_id`),
  KEY `video_id_index` (`video_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `douyin_favorite`
--

LOCK TABLES `douyin_favorite` WRITE;
/*!40000 ALTER TABLE `douyin_favorite` DISABLE KEYS */;
/*!40000 ALTER TABLE `douyin_favorite` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `douyin_follow`
--

DROP TABLE IF EXISTS `douyin_follow`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `douyin_follow` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` int DEFAULT NULL,
  `updated_at` int DEFAULT NULL,
  `followed_id` int DEFAULT NULL,
  `follower_id` int DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uniq_idx` (`follower_id`,`followed_id`) USING BTREE,
  KEY `followed_id` (`followed_id`),
  CONSTRAINT `douyin_follow_ibfk_1` FOREIGN KEY (`followed_id`) REFERENCES `douyin_user` (`id`),
  CONSTRAINT `douyin_follow_ibfk_2` FOREIGN KEY (`follower_id`) REFERENCES `douyin_user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `douyin_follow`
--

LOCK TABLES `douyin_follow` WRITE;
/*!40000 ALTER TABLE `douyin_follow` DISABLE KEYS */;
INSERT INTO `douyin_follow` VALUES (38,1654353471,1654353471,38,39),(39,1654480549,1654480549,38,42),(40,1654480588,1654480588,40,42),(41,1654480928,1654480928,42,42);
/*!40000 ALTER TABLE `douyin_follow` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `douyin_tag`
--

DROP TABLE IF EXISTS `douyin_tag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `douyin_tag` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '标签id',
  `name` varchar(100) NOT NULL COMMENT '标签名',
  PRIMARY KEY (`id`),
  UNIQUE KEY `douyin_tag_UN` (`name`),
  KEY `douyin_tag_name_IDX` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='视频的一些标签类别';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `douyin_tag`
--

LOCK TABLES `douyin_tag` WRITE;
/*!40000 ALTER TABLE `douyin_tag` DISABLE KEYS */;
INSERT INTO `douyin_tag` VALUES (4,'人体部位'),(3,'人物'),(2,'其他'),(6,'动植物'),(5,'哺乳动物'),(1,'宠物'),(11,'手机游戏'),(8,'王者'),(12,'腾讯'),(7,'视频');
/*!40000 ALTER TABLE `douyin_tag` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `douyin_user`
--

DROP TABLE IF EXISTS `douyin_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `douyin_user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_name` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `follow_count` int NOT NULL,
  `follower_count` int NOT NULL,
  `created_at` int NOT NULL,
  `updated_at` int NOT NULL,
  `avatar` varchar(100) NOT NULL,
  `signature` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `background_image` varchar(100) NOT NULL,
  `login_ip` varchar(100) DEFAULT NULL COMMENT '最近登录的ip',
  `total_favorited` bigint NOT NULL DEFAULT '0' COMMENT '被赞总次数',
  `favorite_count` bigint NOT NULL DEFAULT '0' COMMENT '喜欢的总数量',
  PRIMARY KEY (`id`),
  UNIQUE KEY `douyin_user_UN` (`user_name`)
) ENGINE=InnoDB AUTO_INCREMENT=50 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `douyin_user`
--

LOCK TABLES `douyin_user` WRITE;
/*!40000 ALTER TABLE `douyin_user` DISABLE KEYS */;
INSERT INTO `douyin_user` VALUES (38,'502725171@qq.com','$2a$14$9NRy2/PZGHfCDQOLoqAiqO/c0zlietQ6y7.yAJzQ.iDNfBA4cwrz2',0,0,1653447438,1653983972,'https://api.multiavatar.com/502725171@qq.com.png?apikey=A5wbsoJPETy1uk','从前日色变得慢，车，马，邮件都慢。一生只够爱一个人。','https://tuapi.eees.cc/fengjing/img86037463021.jpg','192.168.1.108',0,0),(39,'1206027926@qq.com','$2a$14$r4IKUpyrmd6eo9KfIwqFJuGblPsNxYuOpK66.eTx1LqzhcoR4IjbW',0,0,1653459497,1654352369,'https://api.multiavatar.com/1206027926@qq.com.png?apikey=A5wbsoJPETy1uk','你看，耳机是一对，我们也是一对','https://tuapi.eees.cc/fengjing/img66732284291.jpg','192.168.185.79',0,0),(40,'1506182898','$2a$14$SN6/b0euPwpz2xMKVynWt.Cg9U/gKAf/AedDpdcXQkFlR24C/J6cO',0,0,1653460870,1653983973,'https://api.multiavatar.com/1506182898.png?apikey=A5wbsoJPETy1uk','不知道为啥你要隔三差五发张自拍，我真的无语，要发就天天发，这是在拯救世界','https://tuapi.eees.cc/fengjing/img66338436291.jpg',NULL,0,0),(41,'abc','$2a$14$4JauTyQngVyoxJIOPnv2MejZvm62L3toVcaKTNtYKDnOpHysrNyCe',0,0,1653472672,1653472672,'https://api.multiavatar.com/abc.png?apikey=A5wbsoJPETy1uk','再冷的天你一笑我就暖了','https://tuapi.eees.cc/fengjing/img56934825421.jpg','192.168.56.1',0,0),(42,'123456','$2a$14$9fALx748mdFUdlf9ASzVeOwQo6FEiI1UcUgyx11eKjHBLMMEE1kQi',0,0,1653555622,1653983973,'https://api.multiavatar.com/123456.png?apikey=A5wbsoJPETy1uk','你爱的人不一定爱你','https://tuapi.eees.cc/fengjing/img96335630981.jpg','192.168.56.1',0,0),(43,'2664006323@qq.com','$2a$14$eTkxtfPZpRcUKDP.iA/oFeDFZ5vOwk9cA9jkaWIoHqvt6RMxM13aS',0,0,1653666817,1653923165,'https://api.multiavatar.com/2664006323@qq.com.png?apikey=A5wbsoJPETy1uk','你怎么长成这样？符合我全部想象。','https://tuapi.eees.cc/fengjing/img26730116441.jpg','192.168.1.171',1,1),(44,'532369157@qq.com','$2a$14$4fuzHm7j15HG.nsnIzVdnexuL7wE4ZS7FZe0OmByqIVh2QaaC94Mi',0,0,1653806361,1653983974,'https://api.multiavatar.com/532369157@qq.com.png?apikey=A5wbsoJPETy1uk','一天，我网购了一个魔镜，我对魔镜说:“魔镜，魔镜，谁才是世界上最帅的的人。”。魔镜回答说:“我的主人，您是世上最帅的人。”听完这话我一脚把它踹碎，大吼说:“在场点赞的哪个不比我帅','https://tuapi.eees.cc/fengjing/img36737728861.jpg','192.168.3.38',0,0),(45,'1007187668@qq.com','$2a$14$uL9bBpmnS.vOW6JDc58GlutWrmOL9AAI4ERxh/PVvOzDyldWvB7cW',0,0,1653807708,1653807708,'https://api.multiavatar.com/1007187668@qq.com.png?apikey=A5wbsoJPETy1uk','月亮很亮，亮也没用，没用也亮，喜欢你，喜欢也没用，没用也喜欢','https://tuapi.eees.cc/fengjing/img56434149311.jpg',NULL,0,0),(47,'abcdef','$2a$14$Zh8tt5Ue9XVHEUt5GBBdNu.IoednedjIJzl8uaOBmaC9aQiEADaI6',0,0,1653982318,1653983972,'https://api.multiavatar.com/abcdef.png?apikey=A5wbsoJPETy1uk','我总在意流云与星群的坐标，一朵花开谢的时日，原野与平川一望无垠是否平展。哪知你轻轻抬起眉梢，竟赐予我一生大好河山。','https://tuapi.eees.cc/fengjing/img16637670491.jpg','10.11.74.139',0,0);
/*!40000 ALTER TABLE `douyin_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `douyin_video`
--

DROP TABLE IF EXISTS `douyin_video`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `douyin_video` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '编号id',
  `author_id` int NOT NULL COMMENT '作者id',
  `play_url` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '视频播放地址',
  `cover_url` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '视频封面地址',
  `favorite_count` int DEFAULT '0' COMMENT '视频的点赞总数',
  `comment_count` int DEFAULT '0' COMMENT '视频的评论总数',
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '视频标题',
  `tags` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '视频标签（解析自标题）',
  `publish_date` timestamp NULL DEFAULT NULL COMMENT '发布时期',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=61 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `douyin_video`
--

LOCK TABLES `douyin_video` WRITE;
/*!40000 ALTER TABLE `douyin_video` DISABLE KEYS */;
INSERT INTO `douyin_video` VALUES (28,38,'http://172.27.124.6/38/9a241bb955d9266e083d02ecccb302ed.mp4','https://c-ssl.dtstatic.com/uploads/item/201803/13/20180313083933_olurq.thumb.1000_0.jpg',0,0,'zltest',NULL,'2022-05-25 02:58:01','2022-05-25 02:58:01','2022-05-25 02:58:01',NULL),(29,38,'http://192.168.1.107/38/5b42e2f825b61b553e09ea40e66b2fd8.mp4','https://c-ssl.dtstatic.com/uploads/item/201803/13/20180313083933_olurq.thumb.1000_0.jpg',0,0,'zltest2',NULL,'2022-05-25 03:00:59','2022-05-25 03:00:59','2022-05-25 03:00:59',NULL),(31,40,'http://192.168.56.1:8000/static/40/18a40d1b9f65fcc7ed273dedf8a6d0fa.mp4','https://c-ssl.dtstatic.com/uploads/item/201803/13/20180313083933_olurq.thumb.1000_0.jpg',0,0,'蹭个热门！',NULL,'2022-05-25 06:48:52','2022-05-25 06:48:52','2022-05-25 06:48:52',NULL),(32,40,'http://192.168.56.1:8000/static/40/9a9bd75b024f9c7f0c975628dce0a7bb.mp4','https://c-ssl.dtstatic.com/uploads/item/201803/13/20180313083933_olurq.thumb.1000_0.jpg',0,0,'for fun',NULL,'2022-05-25 06:54:07','2022-05-25 06:54:07','2022-05-25 06:54:07',NULL),(33,38,'http://192.168.1.107/38/914286c12f26509b8b10378de8fdad3e.mp4','https://c-ssl.dtstatic.com/uploads/item/201803/13/20180313083933_olurq.thumb.1000_0.jpg',0,0,'zltest02',NULL,'2022-05-26 07:20:20','2022-05-26 07:20:20','2022-05-26 07:20:20',NULL),(34,38,'http://192.168.1.107:8000/38/f81084fc4c6f9f041b4befa8744889ae.mp4','https://c-ssl.dtstatic.com/uploads/item/201803/13/20180313083933_olurq.thumb.1000_0.jpg',0,0,'zltest01',NULL,'2022-05-26 07:36:36','2022-05-26 07:36:36','2022-05-26 07:36:36',NULL),(35,42,'http://172.27.124.6/42/6840f4561dd0e24166bf67d99951becf.mp4','https://c-ssl.dtstatic.com/uploads/item/201803/13/20180313083933_olurq.thumb.1000_0.jpg',0,0,'wan e zhi yuan',NULL,'2022-05-26 09:00:59','2022-05-26 09:00:59','2022-05-26 09:00:59',NULL),(36,38,'http://192.168.1.107/38/4b31656f4081dec7d6b2149fec4719b1.mp4','https://c-ssl.dtstatic.com/uploads/item/201803/13/20180313083933_olurq.thumb.1000_0.jpg',0,0,'zhtest3',NULL,'2022-05-27 02:53:25','2022-05-27 02:53:25','2022-05-27 02:53:25',NULL),(37,38,'http://192.168.1.107/38/a7a2239da13896a86650d58f91aaa8e6.mp4','https://c-ssl.dtstatic.com/uploads/item/201803/13/20180313083933_olurq.thumb.1000_0.jpg',0,0,'zhtest04',NULL,'2022-05-27 03:00:41','2022-05-27 03:00:41','2022-05-27 03:00:41',NULL),(38,38,'http://192.168.1.107:8000/38/87a9baa38152edb5e242b22815761ebd.mp4','https://c-ssl.dtstatic.com/uploads/item/201803/13/20180313083933_olurq.thumb.1000_0.jpg',0,0,'zltest6',NULL,'2022-05-27 03:05:34','2022-05-27 03:05:34','2022-05-27 03:05:34',NULL),(39,38,'http://192.168.1.107:8000/static/38/94eae28e8b087f66e373079b7284af7d.mp4','https://c-ssl.dtstatic.com/uploads/item/201803/13/20180313083933_olurq.thumb.1000_0.jpg',0,0,'zltest7',NULL,'2022-05-27 03:08:27','2022-05-27 03:08:27','2022-05-27 03:08:27',NULL),(43,39,'http://222.20.74.16:8000/static/39/e93d1df6dfa4f678840043bc387f6405.mp4','http://222.20.74.16:8000/static/39/e93d1df6dfa4f678840043bc387f6405.png',0,0,'猫猫',NULL,'2022-05-28 07:43:04','2022-05-28 07:43:04','2022-05-28 07:43:04',NULL),(44,44,'http://192.168.3.7:8000/static/44/7d37448732330ea4364ff7c4a1262c9e.mp4','http://192.168.3.7:8000/static/44/7d37448732330ea4364ff7c4a1262c9e.png',0,0,'王者荣耀视频','7;8','2022-05-29 06:40:15','2022-05-29 06:40:15','2022-05-29 06:40:15',NULL),(51,44,'http://192.168.3.7:8000/static/44/51f173fbad11e7434760de896a89fbab.mp4','http://192.168.3.7:8000/static/44/51f173fbad11e7434760de896a89fbab.png',0,0,'我的王者','8','2022-05-29 07:14:01','2022-05-29 07:14:01','2022-05-29 07:14:01',NULL),(52,44,'http://192.168.3.7:8000/static/44/5a4525ca2f4e05306458288d3f5c2cd4.mp4','http://192.168.3.7:8000/static/44/5a4525ca2f4e05306458288d3f5c2cd4.png',0,0,'百里守约无敌','','2022-05-29 07:14:25','2022-05-29 07:14:25','2022-05-29 07:14:25',NULL),(53,44,'http://192.168.3.7:8000/static/44/197ab90ac5e8b6d985c9e61f0701d71d.mp4','http://192.168.3.7:8000/static/44/197ab90ac5e8b6d985c9e61f0701d71d.png',0,0,'腾讯手机游戏','11;12','2022-05-29 07:17:13','2022-05-29 07:17:13','2022-05-29 07:17:14',NULL),(54,39,'http://222.20.74.16/39/f0b1fc74c9737acc866b294f366fe9d1.mp4','http://222.20.74.16/39/f0b1fc74c9737acc866b294f366fe9d1.png',0,0,'冲冲冲',NULL,'2022-05-31 08:37:02','2022-05-31 08:37:02','2022-05-31 08:37:02',NULL),(55,39,'http://192.168.185.209:8000/39/61a47bdf94ad06bdd8b668c169c95cc7.mp4','http://192.168.185.209:8000/storage/uploads/39/61a47bdf94ad06bdd8b668c169c95cc7.png',0,0,'要快乐',NULL,'2022-06-04 14:19:57','2022-06-04 14:19:57','2022-06-04 14:19:57',NULL),(56,39,'http://192.168.185.209/39/6858be02f962dac9a618b535ff25c272.mp4','http://192.168.185.209/storage/uploads/39/6858be02f962dac9a618b535ff25c272.png',0,0,'test\n\n',NULL,'2022-06-04 14:31:53','2022-06-04 14:31:53','2022-06-04 14:31:53',NULL),(57,39,'http://192.168.185.79/39/a02e2670da7111c785255bf59fc1becc.mp4','http://192.168.185.79/storage/uploads/39/a02e2670da7111c785255bf59fc1becc.png',0,0,'test',NULL,'2022-06-04 14:32:40','2022-06-04 14:32:40','2022-06-04 14:32:40',NULL),(58,39,'http://192.168.185.79/39/63e61ef64f9f1611b52836061ceae0aa.mp4','http://192.168.185.79/storage/uploads/39/63e61ef64f9f1611b52836061ceae0aa.png',0,0,'test\n',NULL,'2022-06-04 14:42:51','2022-06-04 14:42:51','2022-06-04 14:42:51',NULL),(59,39,'http://192.168.185.209/39/1175efde4ea385db9333b359c05fb822.mp4','http://192.168.185.209/storage/uploads/39/1175efde4ea385db9333b359c05fb822.png',0,0,'test\n',NULL,'2022-06-04 14:50:37','2022-06-04 14:50:37','2022-06-04 14:50:37',NULL),(60,39,'http://10.21.38.17/39/8a97bbe235a7226478a6e1686a054ea8.mp4','http://10.21.38.17/storage/uploads/39/8a97bbe235a7226478a6e1686a054ea8.png',0,0,'test',NULL,'2022-06-04 15:00:51','2022-06-04 15:00:51','2022-06-04 15:00:51',NULL);
/*!40000 ALTER TABLE `douyin_video` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `test_table`
--

DROP TABLE IF EXISTS `test_table`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `test_table` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `count` bigint DEFAULT NULL,
  `key_word` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_test_table_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `test_table`
--

LOCK TABLES `test_table` WRITE;
/*!40000 ALTER TABLE `test_table` DISABLE KEYS */;
INSERT INTO `test_table` VALUES (1,'2022-05-23 18:47:42.543','2022-05-23 18:49:43.987',NULL,10,'video');
/*!40000 ALTER TABLE `test_table` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'douyin'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-06-06 10:08:33
