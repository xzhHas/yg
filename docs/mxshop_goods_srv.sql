/*
 Navicat Premium Data Transfer

 Source Server         : mysql80
 Source Server Type    : MySQL
 Source Server Version : 80013
 Source Host           : localhost:3306
 Source Schema         : mxshop_goods_srv

 Target Server Type    : MySQL
 Target Server Version : 80013
 File Encoding         : 65001

 Date: 07/06/2024 10:46:08
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for banner
-- ----------------------------
DROP TABLE IF EXISTS `banner`;
CREATE TABLE `banner`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_deleted` tinyint(1) NULL DEFAULT NULL,
  `image` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `index` int(11) NOT NULL DEFAULT 1,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of banner
-- ----------------------------
INSERT INTO `banner` VALUES (1, '2024-05-27 16:16:01.117', '2024-05-27 16:16:01.117', NULL, 0, 'http://mxx-shop.oss-cn-beijing.aliyuncs.com/下载 (1).jfif', 'https://blog.csdn.net/m0_73337964?type=blog', 1);
INSERT INTO `banner` VALUES (2, '2024-05-27 16:16:16.973', '2024-05-27 16:16:16.973', NULL, 0, 'http://mxx-shop.oss-cn-beijing.aliyuncs.com/下载.jfif', 'https://blog.csdn.net/m0_73337964?type=blog', 2);
INSERT INTO `banner` VALUES (3, '2024-05-27 16:16:31.613', '2024-05-27 16:16:31.613', NULL, 0, 'http://mxx-shop.oss-cn-beijing.aliyuncs.com/下载.png', 'https://blog.csdn.net/m0_73337964?type=blog', 3);

-- ----------------------------
-- Table structure for brands
-- ----------------------------
DROP TABLE IF EXISTS `brands`;
CREATE TABLE `brands`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_deleted` tinyint(1) NULL DEFAULT NULL,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '\'品牌名称\'',
  `logo` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '\'品牌图标\'',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of brands
-- ----------------------------
INSERT INTO `brands` VALUES (1, '2024-05-27 20:52:44.363', '2024-05-27 20:52:44.363', NULL, 0, '华为', 'http://mxx-shop.oss-cn-beijing.aliyuncs.com/a1.jpg');
INSERT INTO `brands` VALUES (2, '2024-05-27 20:52:56.191', '2024-05-27 20:52:56.191', NULL, 0, '小米', 'http://mxx-shop.oss-cn-beijing.aliyuncs.com/下载 (2).png');
INSERT INTO `brands` VALUES (6, '2024-05-27 21:20:31.601', '2024-05-27 21:20:31.601', '2024-05-27 21:20:34.881', 0, '百花奖', 'http://mxx-shop.oss-cn-beijing.aliyuncs.com/下载 (1).jfif');
INSERT INTO `brands` VALUES (7, '2024-05-27 21:31:59.492', '2024-05-27 21:31:59.492', '2024-05-27 21:34:36.681', 0, '阿斯顿发生df', 'http://mxx-shop.oss-cn-beijing.aliyuncs.com/下载.png');
INSERT INTO `brands` VALUES (8, '2024-05-28 15:17:21.299', '2024-05-28 15:17:21.299', '2024-05-28 15:17:25.064', 0, '123123', 'http://mxx-shop.oss-cn-beijing.aliyuncs.com/QQ截图20240524122230.png');
INSERT INTO `brands` VALUES (9, '2024-05-29 12:03:40.471', '2024-05-29 12:03:40.471', NULL, 0, 'yg', 'http://mxx-shop.oss-cn-beijing.aliyuncs.com/scenere.jpg');

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_deleted` tinyint(1) NULL DEFAULT NULL,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `parent_category_id` int(11) NULL DEFAULT NULL,
  `level` int(11) NOT NULL DEFAULT 1 COMMENT '\'1为1级类目，2为2级...\'',
  `is_tab` tinyint(1) NOT NULL DEFAULT 0 COMMENT '\'能否展示在Tab栏\'',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_category_parent_category`(`parent_category_id` ASC) USING BTREE,
  CONSTRAINT `fk_category_parent_category` FOREIGN KEY (`parent_category_id`) REFERENCES `category` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_category_sub_category` FOREIGN KEY (`parent_category_id`) REFERENCES `category` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of category
-- ----------------------------
INSERT INTO `category` VALUES (1, NULL, NULL, '2024-05-27 21:36:11.307', NULL, '小米', NULL, 1, 1);
INSERT INTO `category` VALUES (2, NULL, NULL, '2024-05-27 21:36:13.792', NULL, '智能手机', NULL, 1, 1);
INSERT INTO `category` VALUES (3, NULL, NULL, NULL, NULL, '无人汽车', NULL, 1, 1);
INSERT INTO `category` VALUES (4, NULL, NULL, NULL, NULL, '智能手机', NULL, 1, 1);
INSERT INTO `category` VALUES (5, NULL, NULL, NULL, NULL, '超级平板', NULL, 1, 1);
INSERT INTO `category` VALUES (6, NULL, NULL, NULL, NULL, '平板电脑', NULL, 1, 1);
INSERT INTO `category` VALUES (7, NULL, NULL, NULL, NULL, '星耀', 4, 2, 1);
INSERT INTO `category` VALUES (8, NULL, NULL, '2024-05-27 21:38:00.967', NULL, '星耀7pro', 4, 2, 0);
INSERT INTO `category` VALUES (9, NULL, NULL, NULL, NULL, '星耀7pro', 7, 3, 1);
INSERT INTO `category` VALUES (10, NULL, NULL, NULL, NULL, '华为', 6, 2, 1);
INSERT INTO `category` VALUES (11, NULL, NULL, NULL, NULL, '曲面屏电脑', 10, 3, 1);
INSERT INTO `category` VALUES (12, NULL, NULL, NULL, NULL, '星耀plus', 7, 3, 1);
INSERT INTO `category` VALUES (13, NULL, NULL, NULL, NULL, '二级车', 3, 2, 1);
INSERT INTO `category` VALUES (14, NULL, NULL, NULL, NULL, '1', 13, 3, 1);
INSERT INTO `category` VALUES (15, NULL, NULL, NULL, NULL, '二级车省级', 3, 2, 1);
INSERT INTO `category` VALUES (16, NULL, NULL, NULL, NULL, '1', 15, 3, 1);
INSERT INTO `category` VALUES (17, NULL, NULL, NULL, NULL, '2', 15, 3, 1);
INSERT INTO `category` VALUES (18, NULL, NULL, NULL, NULL, '3', 13, 3, 1);

-- ----------------------------
-- Table structure for goods
-- ----------------------------
DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_deleted` tinyint(1) NULL DEFAULT NULL,
  `category_id` int(11) NOT NULL,
  `brands_id` int(11) NOT NULL,
  `on_sale` tinyint(1) NOT NULL DEFAULT 0 COMMENT '\'是否上架\'',
  `ship_free` tinyint(1) NOT NULL DEFAULT 0 COMMENT '\'是否免运费\'',
  `is_new` tinyint(1) NOT NULL DEFAULT 0 COMMENT '\'是否新品\'',
  `is_hot` tinyint(1) NOT NULL DEFAULT 0 COMMENT '\'是否热卖商品\'',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `goods_sn` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '\'商家的内部编号\'',
  `click_num` int(11) NOT NULL DEFAULT 0 COMMENT '\'点击数\'',
  `sold_num` int(11) NOT NULL DEFAULT 0 COMMENT '\'销售量\'',
  `fav_num` int(11) NOT NULL DEFAULT 0 COMMENT '\'收藏数\'',
  `market_price` float NOT NULL COMMENT '\'商品价格\'',
  `shop_price` float NOT NULL COMMENT '\'实际价格\'',
  `goods_brief` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '\'商品简介\'',
  `images` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '\'商品图片\'',
  `desc_images` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '\'商品图片\'',
  `goods_front_image` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '\'商品展示图\'',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_goods_category`(`category_id` ASC) USING BTREE,
  INDEX `fk_goods_brands`(`brands_id` ASC) USING BTREE,
  CONSTRAINT `fk_goods_brands` FOREIGN KEY (`brands_id`) REFERENCES `brands` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_goods_category` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 25 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of goods
-- ----------------------------
INSERT INTO `goods` VALUES (1, '2024-05-28 13:33:14.226', '2024-05-28 13:33:14.226', '2024-05-28 14:47:38.418', 0, 9, 2, 0, 1, 1, 1, '星耀手机', '121212', 0, 0, 0, 9888, 8888, '超级手机', '[\"http://mxx-shop.oss-cn-beijing.aliyuncs.com/a.png\"]', '[\"http://mxx-shop.oss-cn-beijing.aliyuncs.com/a.png\"]', 'http://mxx-shop.oss-cn-beijing.aliyuncs.com/a.png');
INSERT INTO `goods` VALUES (2, '2024-05-28 13:34:03.405', '2024-05-28 13:34:03.405', '2024-05-28 14:47:40.127', 0, 11, 1, 0, 1, 0, 0, '无敌电脑', '12312312', 0, 0, 0, 19999, 9999, '无敌电脑', '[\"http://mxx-shop.oss-cn-beijing.aliyuncs.com/下载.png\"]', '[\"http://mxx-shop.oss-cn-beijing.aliyuncs.com/下载.png\"]', 'http://mxx-shop.oss-cn-beijing.aliyuncs.com/下载.png');
INSERT INTO `goods` VALUES (14, '2024-05-28 14:48:36.214', '2024-05-28 20:02:42.843', NULL, 0, 11, 1, 0, 1, 1, 1, '超级电脑', 'a1', 0, 222, 0, 1988, 2211, '超级电脑', '[\"http://mxx-shop.oss-cn-beijing.aliyuncs.com/下载.png\"]', '[\"http://mxx-shop.oss-cn-beijing.aliyuncs.com/下载.png\"]', 'http://mxx-shop.oss-cn-beijing.aliyuncs.com/下载.png');
INSERT INTO `goods` VALUES (15, '2024-05-28 14:50:23.454', '2024-05-28 20:02:41.392', NULL, 0, 9, 2, 0, 1, 1, 1, '无敌手机', 'b1', 0, 12, 0, 9888, 8888, '这是一款超级无敌的手机', '[\"http://mxx-shop.oss-cn-beijing.aliyuncs.com/a.png\"]', '[\"http://mxx-shop.oss-cn-beijing.aliyuncs.com/a.png\"]', 'http://mxx-shop.oss-cn-beijing.aliyuncs.com/a.png');
INSERT INTO `goods` VALUES (16, '2024-05-29 21:05:12.483', '2024-05-30 16:35:40.558', NULL, 0, 14, 1, 0, 1, 0, 0, '123', '123', 0, 0, 0, 123, 123, '123', '[\"http://mxx-shop.oss-cn-beijing.aliyuncs.com/images.jfif\"]', '[\"http://mxx-shop.oss-cn-beijing.aliyuncs.com/images.jfif\"]', 'http://mxx-shop.oss-cn-beijing.aliyuncs.com/images.jfif');
INSERT INTO `goods` VALUES (17, '2024-05-29 21:06:11.596', '2024-05-29 21:06:11.596', NULL, 0, 14, 1, 1, 1, 1, 1, '王子牛奶', '20', 1000, 2000, 1233, 3, 2, '超级好喝的牛奶', '[\"http://mxx-shop.oss-cn-beijing.aliyuncs.com/images.jfif\"]', '[\"http://mxx-shop.oss-cn-beijing.aliyuncs.com/images.jfif\"]', 'http://mxx-shop.oss-cn-beijing.aliyuncs.com/images.jfif');
INSERT INTO `goods` VALUES (18, '2024-05-29 21:07:38.759', '2024-05-30 16:35:42.012', '2024-06-02 17:41:37.883', 0, 11, 1, 0, 1, 0, 0, '13123123', '12313123', 0, 0, 0, 123123000, 123123, '123123123', '[\"http://mxx-shop.oss-cn-beijing.aliyuncs.com/QQ截图20240525173422.png\"]', '[\"http://mxx-shop.oss-cn-beijing.aliyuncs.com/QQ截图20240525173422.png\"]', 'http://mxx-shop.oss-cn-beijing.aliyuncs.com/QQ截图20240525173422.png');
INSERT INTO `goods` VALUES (19, '2024-05-29 21:10:47.927', '2024-05-30 16:35:35.616', NULL, 0, 16, 9, 0, 1, 0, 0, 'asdfasdf', 'asdfasdf', 0, 0, 0, 12312, 12312, '123123', '[\"http://mxx-shop.oss-cn-beijing.aliyuncs.com/images.jfif\"]', '[\"http://mxx-shop.oss-cn-beijing.aliyuncs.com/images.jfif\"]', 'http://mxx-shop.oss-cn-beijing.aliyuncs.com/images.jfif');
INSERT INTO `goods` VALUES (20, '2024-05-29 21:11:11.900', '2024-05-30 16:32:43.582', NULL, 0, 14, 1, 0, 1, 0, 0, 'adfasdfa', '1222', 0, 0, 0, 12312, 1233, 'asdfasdfasd', '[\"http://mxx-shop.oss-cn-beijing.aliyuncs.com/images.jfif\"]', '[\"http://mxx-shop.oss-cn-beijing.aliyuncs.com/images.jfif\"]', 'http://mxx-shop.oss-cn-beijing.aliyuncs.com/images.jfif');
INSERT INTO `goods` VALUES (21, '2024-05-29 21:11:47.648', '2024-05-29 21:11:47.648', NULL, 0, 14, 2, 0, 1, 0, 0, 'asdffffff', '223123', 0, 0, 0, 2233, 123, '123123', '[\"http://mxx-shop.oss-cn-beijing.aliyuncs.com/images.jfif\"]', '[\"http://mxx-shop.oss-cn-beijing.aliyuncs.com/images.jfif\"]', 'http://mxx-shop.oss-cn-beijing.aliyuncs.com/images.jfif');
INSERT INTO `goods` VALUES (22, '2024-05-29 21:12:12.604', '2024-05-29 21:12:12.604', NULL, 0, 17, 1, 0, 1, 0, 0, 'sfdfds', '332', 0, 0, 0, 3232330, 1, '123333', '[\"http://mxx-shop.oss-cn-beijing.aliyuncs.com/images.jfif\"]', '[\"http://mxx-shop.oss-cn-beijing.aliyuncs.com/images.jfif\"]', 'http://mxx-shop.oss-cn-beijing.aliyuncs.com/images.jfif');
INSERT INTO `goods` VALUES (23, '2024-05-29 21:12:46.579', '2024-05-29 21:12:46.579', NULL, 0, 17, 9, 0, 1, 0, 0, 'sa', '2', 0, 0, 0, 2, 2, '32323', '[\"http://mxx-shop.oss-cn-beijing.aliyuncs.com/images.jfif\"]', '[\"http://mxx-shop.oss-cn-beijing.aliyuncs.com/images.jfif\"]', 'http://mxx-shop.oss-cn-beijing.aliyuncs.com/images.jfif');
INSERT INTO `goods` VALUES (24, '2024-05-29 21:13:38.877', '2024-05-29 21:13:38.877', NULL, 0, 14, 9, 0, 1, 0, 0, '3213', '2', 0, 0, 0, 231, 2, '2', '[\"http://mxx-shop.oss-cn-beijing.aliyuncs.com/images.jfif\"]', '[\"http://mxx-shop.oss-cn-beijing.aliyuncs.com/images.jfif\"]', 'http://mxx-shop.oss-cn-beijing.aliyuncs.com/images.jfif');

-- ----------------------------
-- Table structure for goodscategorybrand
-- ----------------------------
DROP TABLE IF EXISTS `goodscategorybrand`;
CREATE TABLE `goodscategorybrand`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_deleted` tinyint(1) NULL DEFAULT NULL,
  `category_id` int(11) NULL DEFAULT NULL,
  `brands_id` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_category_brand`(`category_id` ASC, `brands_id` ASC) USING BTREE,
  INDEX `fk_goodscategorybrand_brands`(`brands_id` ASC) USING BTREE,
  CONSTRAINT `fk_goodscategorybrand_brands` FOREIGN KEY (`brands_id`) REFERENCES `brands` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_goodscategorybrand_category` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 27 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of goodscategorybrand
-- ----------------------------
INSERT INTO `goodscategorybrand` VALUES (1, '2024-05-27 21:39:43.442', '2024-05-27 21:39:43.442', NULL, 0, 9, 2);
INSERT INTO `goodscategorybrand` VALUES (12, '2024-05-29 11:45:38.437', '2024-05-29 11:45:38.437', NULL, 0, 11, 1);
INSERT INTO `goodscategorybrand` VALUES (19, '2024-05-29 21:04:17.109', '2024-05-29 21:04:17.109', NULL, 0, 14, 1);
INSERT INTO `goodscategorybrand` VALUES (20, '2024-05-29 21:04:24.482', '2024-05-29 21:04:24.482', NULL, 0, 17, 1);
INSERT INTO `goodscategorybrand` VALUES (21, '2024-05-29 21:04:31.257', '2024-05-29 21:04:31.257', NULL, 0, 16, 9);
INSERT INTO `goodscategorybrand` VALUES (22, '2024-05-29 21:04:37.299', '2024-05-29 21:04:37.299', NULL, 0, 14, 2);
INSERT INTO `goodscategorybrand` VALUES (25, '2024-05-29 21:10:18.339', '2024-05-29 21:10:18.339', NULL, 0, 17, 9);
INSERT INTO `goodscategorybrand` VALUES (26, '2024-05-29 21:13:10.005', '2024-05-29 21:13:10.005', NULL, 0, 14, 9);

SET FOREIGN_KEY_CHECKS = 1;
