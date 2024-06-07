/*
 Navicat Premium Data Transfer

 Source Server         : mysql80
 Source Server Type    : MySQL
 Source Server Version : 80013
 Source Host           : localhost:3306
 Source Schema         : mxshop_order_srv

 Target Server Type    : MySQL
 Target Server Version : 80013
 File Encoding         : 65001

 Date: 07/06/2024 10:46:26
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for ordergoods
-- ----------------------------
DROP TABLE IF EXISTS `ordergoods`;
CREATE TABLE `ordergoods`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_deleted` tinyint(1) NULL DEFAULT NULL,
  `order` int(11) NULL DEFAULT NULL,
  `goods` int(11) NULL DEFAULT NULL,
  `goods_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `goods_image` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_german2_ci NULL DEFAULT NULL,
  `goods_price` float NULL DEFAULT NULL,
  `nums` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_ordergoods_goods_name`(`goods_name` ASC) USING BTREE,
  INDEX `idx_ordergoods_order`(`order` ASC) USING BTREE,
  INDEX `idx_ordergoods_goods`(`goods` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ordergoods
-- ----------------------------

-- ----------------------------
-- Table structure for orderinfo
-- ----------------------------
DROP TABLE IF EXISTS `orderinfo`;
CREATE TABLE `orderinfo`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_deleted` tinyint(1) NULL DEFAULT NULL,
  `user` int(11) NULL DEFAULT NULL,
  `order_sn` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `pay_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT 'alipay(支付宝)， wechat(微信)',
  `status` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT 'PAYING(待支付), TRADE_SUCCESS(成功)， TRADE_CLOSED(超时关闭), WAIT_BUYER_PAY(交易创建), TRADE_FINISHED(交易结束)',
  `trade_no` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '交易号',
  `order_mount` float NULL DEFAULT NULL,
  `pay_time` datetime NULL DEFAULT NULL,
  `address` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `signer_name` varchar(20) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `singer_mobile` varchar(11) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `post` varchar(20) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_orderinfo_user`(`user` ASC) USING BTREE,
  INDEX `idx_orderinfo_order_sn`(`order_sn` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of orderinfo
-- ----------------------------

-- ----------------------------
-- Table structure for shoppingcart
-- ----------------------------
DROP TABLE IF EXISTS `shoppingcart`;
CREATE TABLE `shoppingcart`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_deleted` tinyint(1) NULL DEFAULT NULL,
  `user` int(11) NULL DEFAULT NULL,
  `goods` int(11) NULL DEFAULT NULL,
  `nums` int(11) NULL DEFAULT NULL,
  `checked` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_shoppingcart_user`(`user` ASC) USING BTREE,
  INDEX `idx_shoppingcart_goods`(`goods` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of shoppingcart
-- ----------------------------
INSERT INTO `shoppingcart` VALUES (1, '2024-05-29 21:28:08.267', '2024-05-29 21:28:15.905', '2024-05-29 21:28:18.006', 0, 14, 16, 1, 0);
INSERT INTO `shoppingcart` VALUES (2, '2024-05-29 21:28:22.849', '2024-05-29 21:28:39.157', '2024-05-29 21:28:41.325', 0, 14, 16, 1, 1);
INSERT INTO `shoppingcart` VALUES (3, '2024-05-30 15:26:17.508', '2024-05-30 15:26:32.454', '2024-05-30 15:26:37.288', 0, 14, 24, 1, 1);
INSERT INTO `shoppingcart` VALUES (4, '2024-05-30 16:07:43.487', '2024-05-30 16:07:43.487', '2024-05-30 16:08:01.410', 0, 14, 17, 13, 0);
INSERT INTO `shoppingcart` VALUES (5, '2024-06-01 21:11:19.882', '2024-06-02 18:04:13.567', '2024-06-02 18:04:54.842', 0, 14, 17, 9, 1);
INSERT INTO `shoppingcart` VALUES (6, '2024-06-02 18:05:08.332', '2024-06-06 21:31:16.026', '2024-06-06 21:31:18.297', 0, 14, 17, 1, 0);
INSERT INTO `shoppingcart` VALUES (7, '2024-06-06 21:31:33.285', '2024-06-07 10:39:50.112', '2024-06-07 10:39:53.934', 0, 14, 17, 1, 0);
INSERT INTO `shoppingcart` VALUES (8, '2024-06-07 10:40:02.760', '2024-06-07 10:40:56.854', NULL, 0, 14, 17, 17, 1);

SET FOREIGN_KEY_CHECKS = 1;
