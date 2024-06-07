/*
 Navicat Premium Data Transfer

 Source Server         : mysql80
 Source Server Type    : MySQL
 Source Server Version : 80013
 Source Host           : localhost:3306
 Source Schema         : mxshop_user_srv

 Target Server Type    : MySQL
 Target Server Version : 80013
 File Encoding         : 65001

 Date: 07/06/2024 10:46:34
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_deleted` tinyint(1) NULL DEFAULT NULL,
  `mobile` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `nick_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `birthday` datetime NULL DEFAULT NULL,
  `gender` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'male' COMMENT 'female表示女, male表示男',
  `role` int(11) NULL DEFAULT 1 COMMENT '1表示普通用户, 2表示管理员',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `mobile`(`mobile` ASC) USING BTREE,
  INDEX `idx_mobile`(`mobile` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (14, '2024-05-23 19:58:35.698', '2024-06-01 19:48:06.015', NULL, 0, '18737735619', '$pbkdf2-sha512$hPwGrx6Xx1tsDFmm$a2abf731121c5f739574d1bce0ad5c7005cd7f31ed6a4e7c85c879708ff8f820', 'xiwanli', '1970-01-01 00:00:00', 'male', 2);

SET FOREIGN_KEY_CHECKS = 1;
