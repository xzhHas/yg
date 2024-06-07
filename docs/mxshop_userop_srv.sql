/*
 Navicat Premium Data Transfer

 Source Server         : mysql80
 Source Server Type    : MySQL
 Source Server Version : 80013
 Source Host           : localhost:3306
 Source Schema         : mxshop_userop_srv

 Target Server Type    : MySQL
 Target Server Version : 80013
 File Encoding         : 65001

 Date: 07/06/2024 10:46:42
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for address
-- ----------------------------
DROP TABLE IF EXISTS `address`;
CREATE TABLE `address`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_deleted` tinyint(1) NULL DEFAULT NULL,
  `user` int(11) NULL DEFAULT NULL,
  `province` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `city` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `district` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `address` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `signer_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `signer_mobile` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_address_user`(`user` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of address
-- ----------------------------
INSERT INTO `address` VALUES (1, '2024-06-01 11:00:27.677', '2024-06-01 11:39:27.913', '2024-06-01 11:39:33.823', 0, 14, '北京市', '北京市', '东城区', '1144区', '西瓜', '18782222220');
INSERT INTO `address` VALUES (2, '2024-06-01 11:40:21.430', '2024-06-01 11:40:21.430', NULL, 0, 14, '北京市', '北京市', '东城区', '1244区', '西瓜', '18737735619');
INSERT INTO `address` VALUES (3, '2024-06-01 11:40:44.950', '2024-06-01 11:40:44.950', NULL, 0, 14, '天津市', '天津市', '南开区', '12312344区', '苹果', '18737735619');

-- ----------------------------
-- Table structure for leavingmessages
-- ----------------------------
DROP TABLE IF EXISTS `leavingmessages`;
CREATE TABLE `leavingmessages`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_deleted` tinyint(1) NULL DEFAULT NULL,
  `user` int(11) NULL DEFAULT NULL,
  `message_type` int(11) NULL DEFAULT NULL COMMENT '留言类型: 1(留言),2(投诉),3(询问),4(售后),5(求购)',
  `subject` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `message` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `file` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_leavingmessages_user`(`user` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of leavingmessages
-- ----------------------------
INSERT INTO `leavingmessages` VALUES (1, '2024-06-01 21:09:33.556', '2024-06-01 21:09:33.556', NULL, 0, 14, 1, '西瓜', '我想吃新鲜的西瓜，请问你们是否可以尽快出售新鲜的西瓜吗', '');

-- ----------------------------
-- Table structure for userfav
-- ----------------------------
DROP TABLE IF EXISTS `userfav`;
CREATE TABLE `userfav`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_deleted` tinyint(1) NULL DEFAULT NULL,
  `user` int(11) NULL DEFAULT NULL,
  `goods` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_user_goods`(`user` ASC, `goods` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of userfav
-- ----------------------------
INSERT INTO `userfav` VALUES (5, '2024-06-01 19:52:44.486', '2024-06-01 19:52:44.486', NULL, 0, 14, 17);
INSERT INTO `userfav` VALUES (7, '2024-06-02 17:47:07.020', '2024-06-02 17:47:07.020', NULL, 0, 14, 15);

SET FOREIGN_KEY_CHECKS = 1;
