/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80032 (8.0.32)
 Source Host           : 127.0.0.1:3306
 Source Schema         : auth-gateway

 Target Server Type    : MySQL
 Target Server Version : 80032 (8.0.32)
 File Encoding         : 65001

 Date: 13/06/2024 16:38:27
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `organization_id` int UNSIGNED NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `mobile` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `role` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `role_id` bigint NULL DEFAULT NULL,
  `token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `user_id` int UNSIGNED NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_token`(`token` ASC) USING BTREE,
  UNIQUE INDEX `uni_mobile`(`mobile` ASC) USING BTREE,
  UNIQUE INDEX `uni_name`(`name` ASC) USING BTREE,
  INDEX `organization_id`(`organization_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin
-- ----------------------------
INSERT INTO `admin` VALUES (1, 1, 'admin', 'admin', 'e64b78fc3bc91bcbc7dc232ba8ec59e0', '超级管理员', 2, '38c9d42c40e6a988b5758afc9f341865', 2, '2023-04-03 16:02:16', '2024-06-13 16:15:45');
INSERT INTO `admin` VALUES (16, 1, 'admin1', 'admin1', 'e64b78fc3bc91bcbc7dc232ba8ec59e0', '123123', 2, 'e5fce01b1b1316e3c0b08ddd38ec8359', 0, '2024-06-13 15:25:54', '2024-06-13 16:14:55');

-- ----------------------------
-- Table structure for admin_menu
-- ----------------------------
DROP TABLE IF EXISTS `admin_menu`;
CREATE TABLE `admin_menu`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `parent_id` int NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `label` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `type` enum('页面','接口') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '页面',
  `show` enum('true','false') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT 'true',
  `sort` smallint NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 155 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_menu
-- ----------------------------
INSERT INTO `admin_menu` VALUES (9, 0, 'home', '主页', '/', 'home-icon', '页面', 'true', 0, '2023-08-10 15:44:04', '2024-04-08 11:23:39');
INSERT INTO `admin_menu` VALUES (20, 0, 'limit', '权限设置', '/limit', 'limit-icon', '页面', 'true', 7, '2023-08-10 15:51:57', '2024-04-02 09:35:34');
INSERT INTO `admin_menu` VALUES (21, 20, 'limit', '菜单', '/limit', '', '页面', 'true', 2, '2023-08-10 15:52:48', '2023-08-10 16:08:02');
INSERT INTO `admin_menu` VALUES (22, 20, 'role', '角色', '/role', '', '页面', 'true', 1, '2023-08-10 15:53:07', '2023-08-10 15:53:07');
INSERT INTO `admin_menu` VALUES (23, 20, 'systemAdmin', '管理员', '/systemAdmin', '', '页面', 'true', 1, '2023-08-10 16:07:57', '2023-08-10 16:07:57');
INSERT INTO `admin_menu` VALUES (24, 20, 'recordList', '操作记录', '/recordList', '', '页面', 'true', 4, '2023-08-11 17:08:30', '2024-03-22 15:33:43');
INSERT INTO `admin_menu` VALUES (36, 24, '', '列表', '/admin/RecordList', '', '接口', 'true', 1, '2024-03-22 11:36:23', '2024-03-22 11:36:23');
INSERT INTO `admin_menu` VALUES (37, 21, '', 'tree', '/admin/MenuTree', '', '接口', 'true', 1, '2024-03-22 11:37:01', '2024-03-22 11:37:01');
INSERT INTO `admin_menu` VALUES (38, 21, '', '编辑', '/admin/MenuEdit', '', '接口', 'true', 2, '2024-03-22 11:37:29', '2024-03-22 11:37:29');
INSERT INTO `admin_menu` VALUES (39, 21, '', '删除', '/admin/MenuDelete', '', '接口', 'true', 3, '2024-03-22 11:37:46', '2024-03-22 11:37:46');
INSERT INTO `admin_menu` VALUES (40, 22, '', '列表', '/admin/RoleList', '', '接口', 'true', 1, '2024-03-22 11:38:24', '2024-03-22 11:38:24');
INSERT INTO `admin_menu` VALUES (41, 22, '', '详情', '/admin/RoleDetail', '', '接口', 'true', 2, '2024-03-22 11:38:36', '2024-03-22 11:38:36');
INSERT INTO `admin_menu` VALUES (42, 22, '', '编辑', '/admin/RoleEdit', '', '接口', 'true', 3, '2024-03-22 11:38:54', '2024-03-22 11:38:54');
INSERT INTO `admin_menu` VALUES (43, 22, '', '操作权限编辑', '/admin/RolePermissionEdit', '', '接口', 'true', 4, '2024-03-22 11:39:17', '2024-03-29 11:06:29');
INSERT INTO `admin_menu` VALUES (44, 22, '', '删除', '/admin/RoleDelete', '', '接口', 'true', 5, '2024-03-22 11:39:29', '2024-03-22 11:39:29');
INSERT INTO `admin_menu` VALUES (45, 23, '', '我的列表', '/admin/MyList', '', '接口', 'true', 1, '2024-03-22 11:40:57', '2024-04-03 10:13:39');
INSERT INTO `admin_menu` VALUES (46, 23, '', '编辑', '/admin/Edit', '', '接口', 'true', 2, '2024-03-22 11:41:14', '2024-03-22 11:41:14');
INSERT INTO `admin_menu` VALUES (47, 23, '', '删除', '/admin/Delete', '', '接口', 'true', 3, '2024-03-22 11:41:30', '2024-03-22 11:41:30');
INSERT INTO `admin_menu` VALUES (93, 0, 'xxxx', '个人信息', 'xx', 'user-icon', '页面', 'false', 10, '2024-04-03 09:22:31', '2024-05-15 09:19:56');
INSERT INTO `admin_menu` VALUES (94, 93, 'xxxx', '常用接口', 'xxxx', '', '页面', 'false', 1, '2024-04-03 09:30:14', '2024-04-03 09:30:14');
INSERT INTO `admin_menu` VALUES (95, 94, '', '登录', '/admin/Login', '', '接口', 'true', 1, '2024-04-03 09:30:48', '2024-05-15 10:14:49');
INSERT INTO `admin_menu` VALUES (96, 94, '', '个人信息', '/admin/Info', '', '接口', 'true', 0, '2024-04-03 09:47:11', '2024-04-03 09:47:11');
INSERT INTO `admin_menu` VALUES (98, 94, '', '退出登陆', '/admin/Logout', '', '接口', 'true', 0, '2024-04-07 16:20:51', '2024-04-07 16:20:51');
INSERT INTO `admin_menu` VALUES (99, 22, '', '全部列表', '/admin/RoleAllList', '', '接口', 'true', 0, '2024-04-07 16:53:06', '2024-04-07 16:53:06');
INSERT INTO `admin_menu` VALUES (148, 94, '', '修改密码', '/admin/ChangePassword', '', '接口', 'true', 4, '2024-05-15 10:14:39', '2024-05-15 10:14:39');
INSERT INTO `admin_menu` VALUES (152, 0, '/limit112', '测试菜单', '/limit12', 'user-icon', '页面', 'true', 3, '2024-06-13 16:17:50', '2024-06-13 16:37:35');
INSERT INTO `admin_menu` VALUES (153, 152, '测试菜单子', '测试菜单子', '/limit12', '', '页面', 'true', 0, '2024-06-13 16:18:18', '2024-06-13 16:37:45');
INSERT INTO `admin_menu` VALUES (154, 153, '', '测试权限', '/article_history/list', '', '接口', 'true', 1, '2024-06-13 16:18:54', '2024-06-13 16:18:54');

-- ----------------------------
-- Table structure for admin_record
-- ----------------------------
DROP TABLE IF EXISTS `admin_record`;
CREATE TABLE `admin_record`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `admin_id` int NULL DEFAULT NULL,
  `admin_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `api_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2275 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_record
-- ----------------------------

-- ----------------------------
-- Table structure for admin_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_role`;
CREATE TABLE `admin_role`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `admin_menu_ids` json NULL,
  `admin_menu_api_paths` json NULL,
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_role
-- ----------------------------
INSERT INTO `admin_role` VALUES (1, '站点管理员', '[9, 20, 23, 45, 46, 47, 99, 40, 37, 93, 94, 96, 98, 95, 148]', '[\"/admin/MenuTree\", \"/admin/RoleList\", \"/admin/MyList\", \"/admin/Edit\", \"/admin/Delete\", \"/admin/Login\", \"/admin/Info\", \"/admin/Logout\", \"/admin/RoleAllList\", \"/admin/ChangePassword\"]', '2023-08-10 08:48:01', '2024-06-13 16:16:22');
INSERT INTO `admin_role` VALUES (2, '超级管理员', '[9, 152, 153, 154, 20, 23, 45, 46, 47, 22, 99, 40, 41, 42, 43, 44, 21, 37, 38, 39, 24, 36, 93, 94, 98, 96, 95, 148]', '[\"/admin/RecordList\", \"/admin/MenuTree\", \"/admin/MenuEdit\", \"/admin/MenuDelete\", \"/admin/RoleList\", \"/admin/RoleDetail\", \"/admin/RoleEdit\", \"/admin/RolePermissionEdit\", \"/admin/RoleDelete\", \"/admin/MyList\", \"/admin/Edit\", \"/admin/Delete\", \"/admin/Login\", \"/admin/Info\", \"/admin/Logout\", \"/admin/RoleAllList\", \"/admin/ChangePassword\", \"/article_history/list\"]', '2023-08-10 08:48:46', '2024-06-13 16:19:15');

SET FOREIGN_KEY_CHECKS = 1;
