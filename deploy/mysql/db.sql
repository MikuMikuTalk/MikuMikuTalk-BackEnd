-- --------------------------------------------------------
-- 主机:                           192.168.3.249
-- 服务器版本:                        8.0.40 - MySQL Community Server - GPL
-- 服务器操作系统:                      Linux
-- HeidiSQL 版本:                  12.8.0.6908
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- 导出 im_server_db 的数据库结构
CREATE DATABASE IF NOT EXISTS `im_server_db` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `im_server_db`;

-- 导出  表 im_server_db.chat_models 结构
CREATE TABLE IF NOT EXISTS `chat_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `send_user_id` bigint unsigned DEFAULT NULL,
  `rev_user_id` bigint unsigned DEFAULT NULL,
  `msg_type` tinyint DEFAULT NULL,
  `msg_preview` varchar(64) COLLATE utf8mb4_bin DEFAULT NULL,
  `msg` longtext COLLATE utf8mb4_bin,
  `system_msg` longtext COLLATE utf8mb4_bin,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=57 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- 正在导出表  im_server_db.chat_models 的数据：~34 rows (大约)
INSERT INTO `chat_models` (`id`, `created_at`, `updated_at`, `send_user_id`, `rev_user_id`, `msg_type`, `msg_preview`, `msg`, `system_msg`) VALUES
	(23, '2025-01-03 03:02:47.920', '2025-01-03 03:02:47.920', 1, 2, 1, 'helloworld', '{"type":1,"textMsg":{"content":"helloworld"},"imageMsg":null,"videoMsg":null,"fileMsg":null,"voiceMsg":null,"voiceCallMsg":null,"videoCallMsg":null,"withdrawMsg":null,"replyMsg":null,"quoteMsg":null,"atMsg":null,"tipMsg":null}', NULL),
	(24, '2025-01-03 05:40:40.108', '2025-01-03 05:40:40.108', 2, 1, 1, 'helloworld', '{"type":1,"textMsg":{"content":"helloworld"},"imageMsg":null,"videoMsg":null,"fileMsg":null,"voiceMsg":null,"voiceCallMsg":null,"videoCallMsg":null,"withdrawMsg":null,"replyMsg":null,"quoteMsg":null,"atMsg":null,"tipMsg":null}', NULL),
	(25, '2025-01-03 05:40:52.620', '2025-01-03 05:40:52.620', 1, 2, 1, 'helloworld222', '{"type":1,"textMsg":{"content":"helloworld222"},"imageMsg":null,"videoMsg":null,"fileMsg":null,"voiceMsg":null,"voiceCallMsg":null,"videoCallMsg":null,"withdrawMsg":null,"replyMsg":null,"quoteMsg":null,"atMsg":null,"tipMsg":null}', NULL),
	(26, '2025-01-03 05:40:56.903', '2025-01-03 05:40:56.903', 1, 2, 1, 'helloworld222', '{"type":1,"textMsg":{"content":"helloworld222"},"imageMsg":null,"videoMsg":null,"fileMsg":null,"voiceMsg":null,"voiceCallMsg":null,"videoCallMsg":null,"withdrawMsg":null,"replyMsg":null,"quoteMsg":null,"atMsg":null,"tipMsg":null}', NULL),
	(27, '2025-01-03 05:55:10.499', '2025-01-03 05:55:10.499', 1, 1, 1, 'hi', '{"type":1,"textMsg":{"content":"hi"},"imageMsg":null,"videoMsg":null,"fileMsg":null,"voiceMsg":null,"voiceCallMsg":null,"videoCallMsg":null,"withdrawMsg":null,"replyMsg":null,"quoteMsg":null,"atMsg":null,"tipMsg":null}', NULL),
	(28, '2025-01-03 05:56:13.285', '2025-01-03 05:56:13.285', 1, 2, 1, 'hi', '{"type":1,"textMsg":{"content":"hi"},"imageMsg":null,"videoMsg":null,"fileMsg":null,"voiceMsg":null,"voiceCallMsg":null,"videoCallMsg":null,"withdrawMsg":null,"replyMsg":null,"quoteMsg":null,"atMsg":null,"tipMsg":null}', NULL),
	(29, '2025-01-03 06:37:19.762', '2025-01-03 06:37:19.762', 1, 2, 1, 'hi', '{"type":1,"textMsg":{"content":"hi"},"imageMsg":null,"videoMsg":null,"fileMsg":null,"voiceMsg":null,"voiceCallMsg":null,"videoCallMsg":null,"withdrawMsg":null,"replyMsg":null,"quoteMsg":null,"atMsg":null,"tipMsg":null}', NULL),
	(30, '2025-01-03 06:37:30.893', '2025-01-03 06:37:30.893', 2, 1, 1, 'helloworld', '{"type":1,"textMsg":{"content":"helloworld"},"imageMsg":null,"videoMsg":null,"fileMsg":null,"voiceMsg":null,"voiceCallMsg":null,"videoCallMsg":null,"withdrawMsg":null,"replyMsg":null,"quoteMsg":null,"atMsg":null,"tipMsg":null}', NULL),
	(31, '2025-01-03 06:38:05.714', '2025-01-03 06:38:05.714', 1, 2, 1, '你在干什么？', '{"type":1,"textMsg":{"content":"你在干什么？"},"imageMsg":null,"videoMsg":null,"fileMsg":null,"voiceMsg":null,"voiceCallMsg":null,"videoCallMsg":null,"withdrawMsg":null,"replyMsg":null,"quoteMsg":null,"atMsg":null,"tipMsg":null}', NULL),
	(32, '2025-01-03 06:38:23.082', '2025-01-03 06:38:23.082', 2, 1, 1, '吃饭', '{"type":1,"textMsg":{"content":"吃饭"},"imageMsg":null,"videoMsg":null,"fileMsg":null,"voiceMsg":null,"voiceCallMsg":null,"videoCallMsg":null,"withdrawMsg":null,"replyMsg":null,"quoteMsg":null,"atMsg":null,"tipMsg":null}', NULL),
	(33, '2025-01-03 07:07:30.046', '2025-01-03 07:07:30.046', 1, 1, 1, '你在干什么？', '{"type":1,"textMsg":{"content":"你在干什么？"},"imageMsg":null,"videoMsg":null,"fileMsg":null,"voiceMsg":null,"voiceCallMsg":null,"videoCallMsg":null,"withdrawMsg":null,"replyMsg":null,"quoteMsg":null,"atMsg":null,"tipMsg":null}', NULL),
	(34, '2025-01-04 03:35:58.119', '2025-01-04 03:35:58.119', 1, 2, 1, 'helloworld', '{"type":1,"textMsg":{"content":"helloworld"}}', NULL),
	(35, '2025-01-04 03:36:10.189', '2025-01-04 03:36:10.189', 2, 2, 1, 'helloworld', '{"type":1,"textMsg":{"content":"helloworld"}}', NULL),
	(36, '2025-01-04 03:39:10.056', '2025-01-04 03:39:10.056', 2, 1, 1, 'helloworld', '{"type":1,"textMsg":{"content":"helloworld"}}', NULL),
	(37, '2025-01-04 03:40:35.227', '2025-01-04 03:40:35.227', 2, 1, 1, 'helloworld', '{"type":1,"textMsg":{"content":"helloworld"}}', NULL),
	(38, '2025-01-04 03:40:36.463', '2025-01-04 03:40:36.463', 2, 1, 1, 'helloworld', '{"type":1,"textMsg":{"content":"helloworld"}}', NULL),
	(39, '2025-01-04 03:41:20.820', '2025-01-04 03:41:20.820', 2, 1, 1, 'helloworld', '{"type":1,"textMsg":{"content":"helloworld"}}', NULL),
	(40, '2025-01-04 03:41:22.559', '2025-01-04 03:41:22.559', 2, 1, 1, 'helloworld', '{"type":1,"textMsg":{"content":"helloworld"}}', NULL),
	(41, '2025-01-04 03:41:23.590', '2025-01-04 03:41:23.590', 2, 1, 1, 'helloworld', '{"type":1,"textMsg":{"content":"helloworld"}}', NULL),
	(42, '2025-01-04 03:41:59.726', '2025-01-04 03:41:59.726', 1, 2, 1, 'helloworld', '{"type":1,"textMsg":{"content":"helloworld"}}', NULL),
	(43, '2025-01-04 03:42:26.957', '2025-01-04 03:42:26.957', 1, 2, 1, 'helloworld', '{"type":1,"textMsg":{"content":"helloworld"}}', NULL),
	(44, '2025-01-04 03:42:28.226', '2025-01-04 03:42:28.226', 1, 2, 1, 'helloworld', '{"type":1,"textMsg":{"content":"helloworld"}}', NULL),
	(45, '2025-01-04 03:48:50.429', '2025-01-04 03:48:50.429', 1, 2, 1, 'helloworld', '{"type":1,"textMsg":{"content":"helloworld"}}', NULL),
	(46, '2025-01-04 03:48:59.391', '2025-01-04 03:48:59.391', 1, 2, 1, 'helloworld', '{"type":1,"textMsg":{"content":"helloworld"}}', NULL),
	(47, '2025-01-04 03:48:59.770', '2025-01-04 03:48:59.770', 1, 2, 1, 'helloworld', '{"type":1,"textMsg":{"content":"helloworld"}}', NULL),
	(48, '2025-01-04 03:49:25.440', '2025-01-04 03:49:25.440', 1, 2, 1, 'helloworld', '{"type":1,"textMsg":{"content":"helloworld"}}', NULL),
	(49, '2025-01-04 03:49:25.595', '2025-01-04 03:49:25.595', 1, 2, 1, 'helloworld', '{"type":1,"textMsg":{"content":"helloworld"}}', NULL),
	(50, '2025-01-04 03:49:25.758', '2025-01-04 03:49:25.758', 1, 2, 1, 'helloworld', '{"type":1,"textMsg":{"content":"helloworld"}}', NULL),
	(51, '2025-01-04 03:49:25.904', '2025-01-04 03:49:25.904', 1, 2, 1, 'helloworld', '{"type":1,"textMsg":{"content":"helloworld"}}', NULL),
	(52, '2025-01-04 03:49:26.043', '2025-01-04 03:49:26.043', 1, 2, 1, 'helloworld', '{"type":1,"textMsg":{"content":"helloworld"}}', NULL),
	(53, '2025-01-04 03:49:26.205', '2025-01-04 03:49:26.205', 1, 2, 1, 'helloworld', '{"type":1,"textMsg":{"content":"helloworld"}}', NULL),
	(54, '2025-01-04 04:02:48.530', '2025-01-04 04:02:48.530', 2, 1, 1, 'helloworld', '{"type":1,"textMsg":{"content":"helloworld"}}', NULL),
	(55, '2025-01-04 04:03:46.825', '2025-01-04 04:03:46.825', 2, 1, 1, 'helloworld', '{"type":1,"textMsg":{"content":"helloworld"}}', NULL),
	(56, '2025-01-04 04:05:00.167', '2025-01-04 04:05:00.167', 2, 1, 1, 'helloworld', '{"type":1,"textMsg":{"content":"helloworld"}}', NULL);

-- 导出  表 im_server_db.file_models 结构
CREATE TABLE IF NOT EXISTS `file_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `file_name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '文件名',
  `hash` varchar(32) COLLATE utf8mb4_bin NOT NULL COMMENT '文件md5值',
  `uid` longtext COLLATE utf8mb4_bin,
  `user_id` bigint unsigned DEFAULT NULL,
  `size` bigint DEFAULT NULL,
  `path` longtext COLLATE utf8mb4_bin,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_filename_md5` (`file_name`,`hash`)
) ENGINE=InnoDB AUTO_INCREMENT=45 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- 正在导出表  im_server_db.file_models 的数据：~10 rows (大约)
INSERT INTO `file_models` (`id`, `created_at`, `updated_at`, `file_name`, `hash`, `uid`, `user_id`, `size`, `path`) VALUES
	(35, '2025-01-02 10:23:54.324', '2025-01-02 10:23:54.324', 'lgoo.jpg', '37e3fe54921e5f1c21f20d696ffa32cf', '061ea4ee-3db0-48dd-959a-6d5447f82315', 1, 17814, 'uploads/avatar/lgoo.jpg'),
	(36, '2025-01-02 10:24:08.072', '2025-01-02 10:24:08.072', 'dfasdfdsasdfs.png', '04fd33a51a9a76c1069d5bdc6e2429e4', '9fe75182-e81e-4886-9b97-952661920e71', 1, 855068, 'uploads/avatar/dfasdfdsasdfs.png'),
	(37, '2025-01-03 10:31:37.621', '2025-01-03 10:31:37.621', 'images.png', '08a4fce8f1969e61cd958c30c98d9fd9', 'c9b649d6-8726-4f7e-97ea-dcfa5854015d', 1, 230974, 'uploads/avatar/images.png'),
	(38, '2025-01-03 10:37:17.766', '2025-01-03 10:37:17.766', 'imagesnDIqrepl.png', '733b482443e1bc2a977725e047709e5f', 'dadc3e7b-ec14-4512-bde1-a4a11b126357', 1, 454674, 'uploads/avatar/imagesnDIqrepl.png'),
	(39, '2025-01-03 10:38:04.748', '2025-01-03 10:38:04.748', 'imagesMAf8oyJo.png', '375ebeb22624a1f307445ed74356c19c', '58d372b4-36c6-4993-bdba-9d233f01027d', 1, 199735, 'uploads/avatar/imagesMAf8oyJo.png'),
	(40, '2025-01-03 10:39:06.368', '2025-01-03 10:39:06.368', 'imagescYyF4C3X.png', 'fb606c6c470f1302f61f2472da3711ae', '75167f6f-b189-449a-8767-9b6fadd27956', 1, 320646, 'uploads/avatar/imagescYyF4C3X.png'),
	(41, '2025-01-03 12:38:12.953', '2025-01-03 12:38:12.953', 'imagescXiN3ioU.png', 'a45ef3eefce6e688ce2fbc5738c2e124', '97dee549-f574-4da9-be7c-d6a84f15657f', 1, 953171, 'uploads/avatar/imagescXiN3ioU.png'),
	(42, '2025-01-04 01:53:01.674', '2025-01-04 01:53:01.674', 'images.png', '10c65cfe27c3153b659075a5a39d10ba', 'c841b55a-5de1-4c83-b9ed-8d17b034e956', 1, 270520, 'uploads/avatar/images.png'),
	(43, '2025-01-04 02:30:59.511', '2025-01-04 02:30:59.511', 'imagesbg6tMqkF.png', 'f86b3826562749b00fbf04ea1620be16', '192df253-3aa8-4971-ba5b-7d7615678792', 2, 628379, 'uploads/avatar/imagesbg6tMqkF.png'),
	(44, '2025-01-04 02:43:26.929', '2025-01-04 02:43:26.929', 'imageszE7YqNRB.png', 'd883ff06d2a672b16b0981cdf8ff9aa6', 'b7ef8354-7cfe-4910-9d29-08d7a1fb6597', 2, 299392, 'uploads/avatar/imageszE7YqNRB.png');

-- 导出  表 im_server_db.friend_models 结构
CREATE TABLE IF NOT EXISTS `friend_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `send_user_id` bigint unsigned DEFAULT NULL COMMENT '''发起验证方ID''',
  `rev_user_id` bigint unsigned DEFAULT NULL COMMENT '''接收验证方ID''',
  `sen_user_notice` varchar(128) COLLATE utf8mb4_bin DEFAULT NULL,
  `rev_user_notice` varchar(128) COLLATE utf8mb4_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_friend_models_send_user_model` (`send_user_id`),
  KEY `fk_friend_models_rev_user_model` (`rev_user_id`),
  CONSTRAINT `fk_friend_models_rev_user_model` FOREIGN KEY (`rev_user_id`) REFERENCES `user_models` (`id`),
  CONSTRAINT `fk_friend_models_send_user_model` FOREIGN KEY (`send_user_id`) REFERENCES `user_models` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- 正在导出表  im_server_db.friend_models 的数据：~0 rows (大约)
INSERT INTO `friend_models` (`id`, `created_at`, `updated_at`, `send_user_id`, `rev_user_id`, `sen_user_notice`, `rev_user_notice`) VALUES
	(1, '2024-12-19 10:38:31.556', '2024-12-19 10:38:31.556', 1, 2, '乌龟', '🐱');

-- 导出  表 im_server_db.friend_verify_models 结构
CREATE TABLE IF NOT EXISTS `friend_verify_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `send_user_id` bigint unsigned DEFAULT NULL COMMENT '''发起验证方用户ID''',
  `rev_user_id` bigint unsigned DEFAULT NULL COMMENT '''接收验证方用户ID''',
  `send_status` tinyint DEFAULT NULL COMMENT '发送方状态 4 删除',
  `rev_status` tinyint DEFAULT NULL COMMENT ' 接收方状态 0 未操作 1 同意 2 拒绝 3 忽略 4 删除 ',
  `status` tinyint DEFAULT '0' COMMENT '''状态: 0 未操作, 1 同意, 2 拒绝, 3 忽略 4 删除''',
  `additional_messages` varchar(128) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '''附加消息''',
  `verification_question` json DEFAULT NULL COMMENT '''验证问题, 仅验证方式为3或4时需要''',
  PRIMARY KEY (`id`),
  KEY `idx_friend_verify_models_send_user_id` (`send_user_id`),
  KEY `idx_friend_verify_models_rev_user_id` (`rev_user_id`),
  CONSTRAINT `fk_friend_verify_models_rev_user_model` FOREIGN KEY (`rev_user_id`) REFERENCES `user_models` (`id`),
  CONSTRAINT `fk_friend_verify_models_send_user_model` FOREIGN KEY (`send_user_id`) REFERENCES `user_models` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- 正在导出表  im_server_db.friend_verify_models 的数据：~0 rows (大约)
INSERT INTO `friend_verify_models` (`id`, `created_at`, `updated_at`, `send_user_id`, `rev_user_id`, `send_status`, `rev_status`, `status`, `additional_messages`, `verification_question`) VALUES
	(1, '2024-12-19 10:37:35.582', '2024-12-19 10:38:31.561', 1, 2, 0, 1, 0, '', NULL);

-- 导出  表 im_server_db.group_member_models 结构
CREATE TABLE IF NOT EXISTS `group_member_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `group_id` bigint unsigned DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  `member_nickname` varchar(32) COLLATE utf8mb4_bin DEFAULT NULL,
  `role` tinyint DEFAULT NULL,
  `prohibition_time` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_group_models_member_list` (`group_id`),
  CONSTRAINT `fk_group_member_models_group_model` FOREIGN KEY (`group_id`) REFERENCES `group_models` (`id`),
  CONSTRAINT `fk_group_models_member_list` FOREIGN KEY (`group_id`) REFERENCES `group_models` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- 正在导出表  im_server_db.group_member_models 的数据：~0 rows (大约)

-- 导出  表 im_server_db.group_models 结构
CREATE TABLE IF NOT EXISTS `group_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `title` longtext COLLATE utf8mb4_bin,
  `abstract` longtext COLLATE utf8mb4_bin,
  `avatar` longtext COLLATE utf8mb4_bin,
  `creator` bigint unsigned DEFAULT NULL,
  `is_search` tinyint(1) DEFAULT NULL,
  `verification` tinyint DEFAULT NULL,
  `verification_question` longtext COLLATE utf8mb4_bin,
  `is_invite` tinyint(1) DEFAULT NULL,
  `is_temporary_session` tinyint(1) DEFAULT NULL,
  `is_prohibition` tinyint(1) DEFAULT NULL,
  `size` bigint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- 正在导出表  im_server_db.group_models 的数据：~0 rows (大约)
INSERT INTO `group_models` (`id`, `created_at`, `updated_at`, `title`, `abstract`, `avatar`, `creator`, `is_search`, `verification`, `verification_question`, `is_invite`, `is_temporary_session`, `is_prohibition`, `size`) VALUES
	(1, '2025-01-04 13:19:57.479', '2025-01-04 13:19:57.479', 'MeowRain的小群', '', 'M', 1, 1, 2, NULL, 0, 0, 0, 100);

-- 导出  表 im_server_db.group_msg_models 结构
CREATE TABLE IF NOT EXISTS `group_msg_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `group_id` bigint unsigned DEFAULT NULL,
  `send_user_id` bigint unsigned DEFAULT NULL,
  `msg_type` tinyint DEFAULT NULL,
  `msg_preview` varchar(64) COLLATE utf8mb4_bin DEFAULT NULL,
  `msg` longtext COLLATE utf8mb4_bin,
  `system_msg` longtext COLLATE utf8mb4_bin,
  `group_member_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_group_msg_models_group_model` (`group_id`),
  KEY `fk_group_msg_models_group_member_model` (`group_member_id`),
  CONSTRAINT `fk_group_msg_models_group_member_model` FOREIGN KEY (`group_member_id`) REFERENCES `group_member_models` (`id`),
  CONSTRAINT `fk_group_msg_models_group_model` FOREIGN KEY (`group_id`) REFERENCES `group_models` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- 正在导出表  im_server_db.group_msg_models 的数据：~0 rows (大约)

-- 导出  表 im_server_db.group_user_msg_delete_models 结构
CREATE TABLE IF NOT EXISTS `group_user_msg_delete_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  `msg_id` bigint unsigned DEFAULT NULL,
  `group_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- 正在导出表  im_server_db.group_user_msg_delete_models 的数据：~0 rows (大约)

-- 导出  表 im_server_db.group_user_top_models 结构
CREATE TABLE IF NOT EXISTS `group_user_top_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  `group_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- 正在导出表  im_server_db.group_user_top_models 的数据：~0 rows (大约)

-- 导出  表 im_server_db.group_verify_models 结构
CREATE TABLE IF NOT EXISTS `group_verify_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `group_id` bigint unsigned DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  `status` tinyint DEFAULT NULL,
  `additional_messages` varchar(32) COLLATE utf8mb4_bin DEFAULT NULL,
  `verification_question` longtext COLLATE utf8mb4_bin,
  `type` tinyint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_group_verify_models_group_model` (`group_id`),
  CONSTRAINT `fk_group_verify_models_group_model` FOREIGN KEY (`group_id`) REFERENCES `group_models` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- 正在导出表  im_server_db.group_verify_models 的数据：~0 rows (大约)

-- 导出  表 im_server_db.log_models 结构
CREATE TABLE IF NOT EXISTS `log_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `log_type` tinyint DEFAULT NULL,
  `ip` varchar(32) COLLATE utf8mb4_bin DEFAULT NULL,
  `addr` varchar(64) COLLATE utf8mb4_bin DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  `user_nickname` varchar(64) COLLATE utf8mb4_bin DEFAULT NULL,
  `user_avatar` varchar(256) COLLATE utf8mb4_bin DEFAULT NULL,
  `level` varchar(12) COLLATE utf8mb4_bin DEFAULT NULL,
  `title` varchar(32) COLLATE utf8mb4_bin DEFAULT NULL,
  `content` longtext COLLATE utf8mb4_bin,
  `service` varchar(32) COLLATE utf8mb4_bin DEFAULT NULL,
  `is_read` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- 正在导出表  im_server_db.log_models 的数据：~0 rows (大约)

-- 导出  表 im_server_db.top_user_models 结构
CREATE TABLE IF NOT EXISTS `top_user_models` (
  `user_id` bigint unsigned DEFAULT NULL,
  `top_user_id` bigint unsigned DEFAULT NULL,
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- 正在导出表  im_server_db.top_user_models 的数据：~0 rows (大约)

-- 导出  表 im_server_db.user_chat_delete_models 结构
CREATE TABLE IF NOT EXISTS `user_chat_delete_models` (
  `user_id` bigint unsigned DEFAULT NULL,
  `chat_id` bigint unsigned DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- 正在导出表  im_server_db.user_chat_delete_models 的数据：~0 rows (大约)

-- 导出  表 im_server_db.user_conf_models 结构
CREATE TABLE IF NOT EXISTS `user_conf_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL COMMENT '''用户ID''',
  `recall_message` varchar(32) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '''撤回消息的提示内容''',
  `friend_online` tinyint(1) DEFAULT '1' COMMENT '''好友上线提醒''',
  `enable_sound` tinyint(1) DEFAULT '1' COMMENT '''是否开启声音提醒''',
  `secure_link` tinyint(1) DEFAULT '0' COMMENT '''安全链接是否开启''',
  `save_pwd` tinyint(1) DEFAULT '0' COMMENT '''是否保存密码''',
  `search_user` tinyint DEFAULT '0' COMMENT '''别人查找到你的方式: 0 不允许, 1 用户号, 2 昵称''',
  `verification` tinyint DEFAULT '2' COMMENT '''验证类型: 0 不允许任何人, 1 允许任何人, 2 验证消息, 3 回答问题, 4 正确回答问题''',
  `verification_question` json DEFAULT NULL COMMENT '''验证问题''',
  `online` tinyint(1) DEFAULT '0' COMMENT '''是否在线''',
  `curtail_chat` tinyint(1) DEFAULT NULL,
  `curtail_add_user` tinyint(1) DEFAULT NULL,
  `curtail_create_group` tinyint(1) DEFAULT NULL,
  `curtail_in_group_chat` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_user_models_user_conf_model` (`user_id`),
  CONSTRAINT `fk_user_models_user_conf_model` FOREIGN KEY (`user_id`) REFERENCES `user_models` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- 正在导出表  im_server_db.user_conf_models 的数据：~2 rows (大约)
INSERT INTO `user_conf_models` (`id`, `created_at`, `updated_at`, `user_id`, `recall_message`, `friend_online`, `enable_sound`, `secure_link`, `save_pwd`, `search_user`, `verification`, `verification_question`, `online`, `curtail_chat`, `curtail_add_user`, `curtail_create_group`, `curtail_in_group_chat`) VALUES
	(1, '2024-12-19 10:11:30.831', '2024-12-19 10:11:30.831', 1, NULL, 1, 1, 0, 0, 2, 2, NULL, 1, NULL, NULL, NULL, NULL),
	(2, '2024-12-19 10:36:12.381', '2024-12-19 10:36:12.381', 2, NULL, 1, 1, 0, 0, 2, 2, NULL, 1, NULL, NULL, NULL, NULL);

-- 导出  表 im_server_db.user_models 结构
CREATE TABLE IF NOT EXISTS `user_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `pwd` varchar(64) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '''密码''',
  `nickname` varchar(32) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '''用户名''',
  `abstract` varchar(128) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '''简介''',
  `avatar` varchar(256) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '''头像''',
  `ip` varchar(32) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '''ip地址''',
  `addr` varchar(64) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '''地址''',
  `role` tinyint DEFAULT NULL COMMENT '''角色 1是管理员 2是普通用户''',
  `register_source` varchar(16) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '''注册来源''',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_user_models_nickname` (`nickname`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- 正在导出表  im_server_db.user_models 的数据：~2 rows (大约)
INSERT INTO `user_models` (`id`, `created_at`, `updated_at`, `pwd`, `nickname`, `abstract`, `avatar`, `ip`, `addr`, `role`, `register_source`) VALUES
	(1, '2024-12-19 10:11:30.828', '2025-01-04 01:53:01.914', '$2a$04$t9tjHkWKCYBkjUnb/eY3JeoGBBOLcf.YUL0vsP/4gms/J5E/649XS', 'meowrain', '努力生活', '/api/file/c841b55a-5de1-4c83-b9ed-8d17b034e956', '', '', 2, '账户密码注册'),
	(2, '2024-12-19 10:36:12.377', '2025-01-04 02:43:27.562', '$2a$04$DfZNKCQn58HGsnawSuJdgeXDFzfAkc.kSIlUOr3FJqQIzbdQEpo6i', 'vedal987', 'Fight for neurosama', '/api/file/b7ef8354-7cfe-4910-9d29-08d7a1fb6597', '', '', 2, '账户密码注册');

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
