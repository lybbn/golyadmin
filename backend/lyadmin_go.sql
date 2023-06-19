# Host: localhost  (Version: 5.7.26)
# Date: 2023-06-17 01:18:45
# Generator: MySQL-Front 5.3  (Build 4.234)

/*!40101 SET NAMES utf8 */;

#
# Structure for table "lyadmin_users"
#

CREATE TABLE `lyadmin_users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `uuid` varchar(50) DEFAULT NULL COMMENT 'uuid',
  `username` varchar(50) NOT NULL COMMENT '用户名',
  `password` varchar(128) DEFAULT NULL COMMENT '密码',
  `nickname` varchar(20) DEFAULT '系统用户' COMMENT '昵称',
  `mobile` char(25) DEFAULT NULL COMMENT '手机号',
  `email` varchar(100) DEFAULT NULL COMMENT '邮箱',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `gender` varchar(10) DEFAULT NULL COMMENT '性别',
  `dept_id` mediumint(9) DEFAULT NULL COMMENT '部门',
  `post_id` mediumint(9) DEFAULT NULL COMMENT '岗位',
  `role_id` mediumint(9) DEFAULT NULL COMMENT '角色ID',
  `is_staff` tinyint(1) DEFAULT '1' COMMENT '是否可登录后台',
  `is_superuser` tinyint(1) DEFAULT '0' COMMENT '是否超管',
  `is_active` tinyint(1) DEFAULT '1' COMMENT '状态(1正常、2冻结)',
  `identity` tinyint(4) DEFAULT '1' COMMENT '身份(1后台、2前台)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  KEY `idx_lyadmin_users_username` (`username`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

#
# Data for table "lyadmin_users"
#

REPLACE INTO `lyadmin_users` VALUES (1,'2023-06-17 00:33:57.627','2023-06-17 00:33:57.627','74407f959b6340faa23ed48b4235aa8c','superadmin','$2a$10$dheIRWaaI1BmjzH8c0XQU.JWF..otoPqqOKwgIBggj5ZPV6aMOsN6','系统用户','','','','',0,0,0,1,0,1,1);
