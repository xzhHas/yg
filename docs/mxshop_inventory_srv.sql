/*
 Navicat Premium Data Transfer

 Source Server         : mysql80
 Source Server Type    : MySQL
 Source Server Version : 80013
 Source Host           : localhost:3306
 Source Schema         : mxshop_inventory_srv

 Target Server Type    : MySQL
 Target Server Version : 80013
 File Encoding         : 65001

 Date: 07/06/2024 10:46:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for inventory
-- ----------------------------
DROP TABLE IF EXISTS `inventory`;
CREATE TABLE `inventory`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_deleted` tinyint(1) NULL DEFAULT NULL,
  `goods` int(11) NULL DEFAULT NULL,
  `stocks` int(11) NULL DEFAULT NULL,
  `version` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_inventory_goods`(`goods` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of inventory
-- ----------------------------
INSERT INTO `inventory` VALUES (1, '2024-05-28 14:48:36.334', '2024-05-28 14:48:36.334', NULL, 0, 14, 90, 0);
INSERT INTO `inventory` VALUES (2, '2024-05-28 14:50:23.485', '2024-05-28 14:50:23.485', NULL, 0, 15, 1000, 0);
INSERT INTO `inventory` VALUES (3, '2024-05-29 21:05:12.598', '2024-05-29 21:05:12.598', NULL, 0, 16, 123, 0);
INSERT INTO `inventory` VALUES (4, '2024-05-29 21:06:11.616', '2024-06-07 10:43:00.614', NULL, 0, 17, 2306, 0);
INSERT INTO `inventory` VALUES (5, '2024-05-29 21:07:38.778', '2024-05-29 21:07:38.778', NULL, 0, 18, 123123, 0);
INSERT INTO `inventory` VALUES (6, '2024-05-29 21:10:47.991', '2024-05-29 21:10:47.991', NULL, 0, 19, 123123, 0);
INSERT INTO `inventory` VALUES (7, '2024-05-29 21:11:11.992', '2024-05-29 21:11:11.992', NULL, 0, 20, 123123, 0);
INSERT INTO `inventory` VALUES (8, '2024-05-29 21:11:47.676', '2024-05-29 21:11:47.676', NULL, 0, 21, 3333, 0);
INSERT INTO `inventory` VALUES (9, '2024-05-29 21:12:12.645', '2024-05-29 21:12:12.645', NULL, 0, 22, 333, 0);
INSERT INTO `inventory` VALUES (10, '2024-05-29 21:12:46.634', '2024-05-29 21:12:46.634', NULL, 0, 23, 2, 0);
INSERT INTO `inventory` VALUES (11, '2024-05-29 21:13:38.953', '2024-05-29 21:13:38.953', NULL, 0, 24, 2, 0);

-- ----------------------------
-- Table structure for stockselldetail
-- ----------------------------
DROP TABLE IF EXISTS `stockselldetail`;
CREATE TABLE `stockselldetail`  (
  `order_sn` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `status` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `detail` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  UNIQUE INDEX `idx_order_sn`(`order_sn` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of stockselldetail
-- ----------------------------
INSERT INTO `stockselldetail` VALUES ('imooc-bobby', '1', '[{\"Goods\":1,\"Num\":2},{\"Goods\":2,\"Num\":3}]');
INSERT INTO `stockselldetail` VALUES ('20246710435196044001471', '1', '[{\"Goods\":17,\"Num\":17}]');

SET FOREIGN_KEY_CHECKS = 1;
