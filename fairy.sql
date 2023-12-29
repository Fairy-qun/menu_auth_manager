/*
 Navicat Premium Data Transfer

 Source Server         : gin_user
 Source Server Type    : MySQL
 Source Server Version : 80030
 Source Host           : localhost:3306
 Source Schema         : fairy

 Target Server Type    : MySQL
 Target Server Version : 80030
 File Encoding         : 65001

 Date: 29/12/2023 10:57:08
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_casbin_rule`(`ptype` ASC, `v0` ASC, `v1` ASC, `v2` ASC, `v3` ASC, `v4` ASC, `v5` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (1, 'g', 'admin', 'admin', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (2, 'g', 'xiaomai', 'user', NULL, NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (3, 'g', 'xiaomai1', 'user', NULL, NULL, NULL, NULL);

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `parent_id` int NULL DEFAULT NULL,
  `sort` int NULL DEFAULT NULL,
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `type` int NULL DEFAULT NULL,
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `role` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (1, '主控台', 0, 0, '/sys/home', 2, 'House', 'user');
INSERT INTO `sys_menu` VALUES (2, '系统管理', 0, 1, NULL, 1, 'Setting', 'admin');
INSERT INTO `sys_menu` VALUES (3, '用户管理', 2, 1, '/sys/user', 2, 'User', 'admin');
INSERT INTO `sys_menu` VALUES (4, '角色管理', 2, 2, '/sys/role', 2, 'Compass', 'admin');
INSERT INTO `sys_menu` VALUES (5, '菜单管理', 2, 3, '/sys/menu', 2, 'House', 'admin');
INSERT INTO `sys_menu` VALUES (6, '站外链接', 0, 2, NULL, 1, 'Discount', 'user');
INSERT INTO `sys_menu` VALUES (7, 'golang官网', 6, 1, '/site/golang', 2, 'MessageBox', 'user');
INSERT INTO `sys_menu` VALUES (8, 'php官网', 6, 2, '/site/php', 2, 'Crop', 'user');
INSERT INTO `sys_menu` VALUES (9, '关于我们', 0, 3, '/about/me', 2, 'PieChart', 'user');

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (2, 'xiaomai', '$2a$10$dcH2./y1oz51WCggriB2..3Ze8bt0CSNRsQmzq5KXEJHDpIUya1zq');
INSERT INTO `sys_user` VALUES (3, 'xiaomai1', '$2a$10$67hr/FhXltkZ47nXtCl3DOzqY7yxy3y/XVN048qWb26YmfEYh.gkW');
INSERT INTO `sys_user` VALUES (4, 'admin', '$2a$10$HsBQXhLc8J7Oxxmsdkl8keQSf2sFPAy9gbSMddnNW.9Wzb1Rlk5wO');

SET FOREIGN_KEY_CHECKS = 1;
