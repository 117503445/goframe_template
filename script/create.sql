SET NAMES utf8mb4;

SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles`  (
                          `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
                          `created_at` datetime(3) NULL DEFAULT NULL,
                          `updated_at` datetime(3) NULL DEFAULT NULL,
                          `deleted_at` datetime(3) NULL DEFAULT NULL,
                          `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
                          PRIMARY KEY (`id`) USING BTREE,
                          INDEX `idx_roles_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of roles
-- ----------------------------
INSERT INTO `roles` VALUES (1, '2021-01-23 00:22:09.266', '2021-01-23 00:22:09.266', NULL, 'admin');
INSERT INTO `roles` VALUES (2, '2021-01-23 00:22:09.266', '2021-01-23 00:22:09.266', NULL, 'user');

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role`  (
                              `user_id` bigint UNSIGNED NOT NULL,
                              `role_id` bigint UNSIGNED NOT NULL,
                              PRIMARY KEY (`user_id`, `role_id`) USING BTREE,
                              INDEX `fk_user_role_role`(`role_id`) USING BTREE,
                              CONSTRAINT `fk_user_role_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
                              CONSTRAINT `fk_user_role_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_role
-- ----------------------------
INSERT INTO `user_role` VALUES (1, 1);
INSERT INTO `user_role` VALUES (1, 2);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
                          `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
                          `created_at` datetime(3) NULL DEFAULT NULL,
                          `updated_at` datetime(3) NULL DEFAULT NULL,
                          `deleted_at` datetime(3) NULL DEFAULT NULL,
                          `username` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
                          `password` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
                          `avatar` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
                          PRIMARY KEY (`id`) USING BTREE,
                          INDEX `idx_users_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, '2021-01-23 00:22:09.264', '2021-01-23 00:22:09.264', NULL, 'admin', '$2a$12$XLiKy7M77cY56.1aE9IxDeKONAHbz1Z0pE7IOmzRfpjTHKiMZYsjG', '');

SET FOREIGN_KEY_CHECKS = 1;
