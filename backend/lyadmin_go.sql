# Host: localhost  (Version: 5.7.26)
# Date: 2023-07-09 22:34:21
# Generator: MySQL-Front 5.3  (Build 4.234)

/*!40101 SET NAMES utf8 */;

#
# Structure for table "lyadmin_button"
#

CREATE TABLE `lyadmin_button` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `name` varchar(256) DEFAULT NULL COMMENT '按钮名称',
  `value` varchar(256) DEFAULT NULL COMMENT '按钮值',
  `create_by` bigint(20) unsigned DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) unsigned DEFAULT NULL COMMENT '更新者',
  `belong_dept` bigint(20) unsigned DEFAULT NULL COMMENT '数据归属部门',
  PRIMARY KEY (`id`),
  KEY `idx_lyadmin_button_create_by` (`create_by`),
  KEY `idx_lyadmin_button_update_by` (`update_by`),
  KEY `idx_lyadmin_button_belong_dept` (`belong_dept`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;

#
# Data for table "lyadmin_button"
#

REPLACE INTO `lyadmin_button` VALUES (1,'2023-06-26 23:06:46.625','2023-06-26 23:06:46.625','新增','Create',1,0,0),(2,'2023-06-26 23:07:00.289','2023-06-26 23:07:00.289','编辑','Update',1,0,0),(3,'2023-06-26 23:07:12.712','2023-06-26 23:07:12.712','删除','Delete',1,0,0),(4,'2023-06-26 23:07:59.467','2023-06-26 23:07:59.467','查询','Search',1,0,0),(5,'2023-07-01 00:04:00.989','2023-07-01 00:04:00.989','修改密码','Changepassword',1,0,0),(6,'2023-07-01 13:14:00.914','2023-07-01 13:14:00.914','详情','Detail',1,0,0),(7,'2023-07-01 13:38:23.509','2023-07-01 13:38:23.509','保存','Save',1,0,0);

#
# Structure for table "lyadmin_dept"
#

CREATE TABLE `lyadmin_dept` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `parent_id` bigint(20) unsigned DEFAULT NULL COMMENT '上级部门',
  `name` varchar(256) DEFAULT NULL COMMENT '部门名称',
  `sort` bigint(20) DEFAULT '1' COMMENT '显示顺序',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态',
  `owner` varchar(100) DEFAULT NULL COMMENT '负责人',
  `phone` varchar(32) DEFAULT NULL COMMENT '手机',
  `email` varchar(64) DEFAULT NULL COMMENT '邮箱',
  `create_by` bigint(20) unsigned DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) unsigned DEFAULT NULL COMMENT '更新者',
  `belong_dept` bigint(20) unsigned DEFAULT NULL COMMENT '数据归属部门',
  PRIMARY KEY (`id`),
  KEY `idx_lyadmin_dept_create_by` (`create_by`),
  KEY `idx_lyadmin_dept_update_by` (`update_by`),
  KEY `idx_lyadmin_dept_belong_dept` (`belong_dept`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

#
# Data for table "lyadmin_dept"
#

REPLACE INTO `lyadmin_dept` VALUES (1,'2023-06-26 11:06:52.000','2023-07-01 13:00:15.138',0,'golyadmin团队',1,1,'','','',0,1,0),(2,'2023-06-27 10:42:20.776','2023-07-01 12:59:38.480',1,'财务部门',1,1,'','','',1,1,0);

#
# Structure for table "lyadmin_jwt_blacklist"
#

CREATE TABLE `lyadmin_jwt_blacklist` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `jwt` text COMMENT 'jwt',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4;

#
# Data for table "lyadmin_jwt_blacklist"
#

REPLACE INTO `lyadmin_jwt_blacklist` VALUES (1,'2023-06-26 09:57:21.417','2023-06-26 09:57:21.417','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTBhYzZiMzQ5NTMwNGI0YzlhZDAxZWMxYTQzN2MyMmQiLCJJRCI6MSwiVXNlcm5hbWUiOiJzdXBlcmFkbWluIiwiTmlja25hbWUiOiLotoXnuqfnrqHnkIblkZgiLCJJZGVudGl0eSI6MCwiRGVwdElkIjowLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoibHliYm4iLCJhdWQiOlsibHlhZG1pbiJdLCJleHAiOjE2ODgzNDkwODcsIm5iZiI6MTY4Nzc0NDI4NywiaWF0IjoxNjg3NzQ0Mjg3fQ.hAus8hNqfUejfD2c6Up-xgul3nTTb4eaZsa1vwVPlQE'),(2,'2023-06-26 16:30:24.380','2023-06-26 16:30:24.380','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTBhYzZiMzQ5NTMwNGI0YzlhZDAxZWMxYTQzN2MyMmQiLCJJRCI6MSwiVXNlcm5hbWUiOiJzdXBlcmFkbWluIiwiTmlja25hbWUiOiLotoXnuqfnrqHnkIblkZgiLCJJZGVudGl0eSI6MSwiRGVwdElkIjowLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoibHliYm4iLCJhdWQiOlsibHlhZG1pbiJdLCJleHAiOjE2ODgzNDk0NDEsIm5iZiI6MTY4Nzc0NDY0MSwiaWF0IjoxNjg3NzQ0NjQxfQ.eiPfvqMTcAvxW5d3dDflSvz9DwoQ6jkdEuJJ0Q4Mneg'),(3,'2023-06-26 16:32:01.673','2023-06-26 16:32:01.673','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTBhYzZiMzQ5NTMwNGI0YzlhZDAxZWMxYTQzN2MyMmQiLCJJRCI6MSwiVXNlcm5hbWUiOiJzdXBlcmFkbWluIiwiTmlja25hbWUiOiLotoXnuqfnrqHnkIblkZgiLCJJZGVudGl0eSI6MSwiRGVwdElkIjowLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoibHliYm4iLCJhdWQiOlsibHlhZG1pbiJdLCJleHAiOjE2ODgzNzMwMjQsIm5iZiI6MTY4Nzc2ODIyNCwiaWF0IjoxNjg3NzY4MjI0fQ.6JKoUmNVXuzvvkL2gcZ0nQ_rBhGnqRKJ9zPxiurVgm4'),(4,'2023-06-26 16:38:39.727','2023-06-26 16:38:39.727','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTBhYzZiMzQ5NTMwNGI0YzlhZDAxZWMxYTQzN2MyMmQiLCJJRCI6MSwiVXNlcm5hbWUiOiJzdXBlcmFkbWluIiwiTmlja25hbWUiOiLotoXnuqfnrqHnkIblkZgiLCJJZGVudGl0eSI6MSwiRGVwdElkIjowLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoibHliYm4iLCJhdWQiOlsibHlhZG1pbiJdLCJleHAiOjE2ODgzNzMxMjEsIm5iZiI6MTY4Nzc2ODMyMSwiaWF0IjoxNjg3NzY4MzIxfQ.2fM0yc3_MT7HzibtXtIxzdaU-YgQPl-GMCHrWZZy-ys'),(5,'2023-06-26 16:39:47.972','2023-06-26 16:39:47.972','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTBhYzZiMzQ5NTMwNGI0YzlhZDAxZWMxYTQzN2MyMmQiLCJJRCI6MSwiVXNlcm5hbWUiOiJzdXBlcmFkbWluIiwiTmlja25hbWUiOiLotoXnuqfnrqHnkIblkZgiLCJJZGVudGl0eSI6MSwiRGVwdElkIjowLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoibHliYm4iLCJhdWQiOlsibHlhZG1pbiJdLCJleHAiOjE2ODgzNzM1MTksIm5iZiI6MTY4Nzc2ODcxOSwiaWF0IjoxNjg3NzY4NzE5fQ.bROTT7PmrCvDfWvMHx6CkJNCgGbdg0EAGusc63tU_dY'),(6,'2023-06-26 16:58:28.358','2023-06-26 16:58:28.358','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTBhYzZiMzQ5NTMwNGI0YzlhZDAxZWMxYTQzN2MyMmQiLCJJRCI6MSwiVXNlcm5hbWUiOiJzdXBlcmFkbWluIiwiTmlja25hbWUiOiLotoXnuqfnrqHnkIblkZgiLCJJZGVudGl0eSI6MSwiRGVwdElkIjowLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoibHliYm4iLCJhdWQiOlsibHlhZG1pbiJdLCJleHAiOjE2ODgzNzM1ODcsIm5iZiI6MTY4Nzc2ODc4NywiaWF0IjoxNjg3NzY4Nzg3fQ.RyGXHNHEgXkUfD6RlNt1w7jf4Nm_QpLE7c3kzNyO64U'),(7,'2023-06-26 17:02:29.099','2023-06-26 17:02:29.099','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTBhYzZiMzQ5NTMwNGI0YzlhZDAxZWMxYTQzN2MyMmQiLCJJRCI6MSwiVXNlcm5hbWUiOiJzdXBlcmFkbWluIiwiTmlja25hbWUiOiLotoXnuqfnrqHnkIblkZgiLCJJZGVudGl0eSI6MSwiRGVwdElkIjowLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoibHliYm4iLCJhdWQiOlsibHlhZG1pbiJdLCJleHAiOjE2ODgzNzQ3MDgsIm5iZiI6MTY4Nzc2OTkwOCwiaWF0IjoxNjg3NzY5OTA4fQ.Sk9YAGnQnsIerf3aTFaNpUVepU_aIbs77THvdSukF8U'),(8,'2023-06-26 17:11:24.589','2023-06-26 17:11:24.589','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTBhYzZiMzQ5NTMwNGI0YzlhZDAxZWMxYTQzN2MyMmQiLCJJRCI6MSwiVXNlcm5hbWUiOiJzdXBlcmFkbWluIiwiTmlja25hbWUiOiLotoXnuqfnrqHnkIblkZgiLCJJZGVudGl0eSI6MSwiRGVwdElkIjowLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoibHliYm4iLCJhdWQiOlsibHlhZG1pbiJdLCJleHAiOjE2ODgzNzQ5NDksIm5iZiI6MTY4Nzc3MDE0OSwiaWF0IjoxNjg3NzcwMTQ5fQ.hnjgx15r-WS8RvpQfnx6mxU-2l7hgwvV118jBrSIu1c'),(9,'2023-06-26 18:20:24.796','2023-06-26 18:20:24.796','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTBhYzZiMzQ5NTMwNGI0YzlhZDAxZWMxYTQzN2MyMmQiLCJJRCI6MSwiVXNlcm5hbWUiOiJzdXBlcmFkbWluIiwiTmlja25hbWUiOiLotoXnuqfnrqHnkIblkZgiLCJJZGVudGl0eSI6MSwiRGVwdElkIjowLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoibHliYm4iLCJhdWQiOlsibHlhZG1pbiJdLCJleHAiOjE2ODgzNzU0ODQsIm5iZiI6MTY4Nzc3MDY4NCwiaWF0IjoxNjg3NzcwNjg0fQ._1S7Bf3N9LgMKBUzEZF0Eh8Q9SKRW5PiMC6mDGjL2-4'),(10,'2023-06-26 20:58:40.213','2023-06-26 20:58:40.213','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZGYwMzMyNzJjZmVmNDhjNzkyYmEyYmQ5OTM1NDVlYWIiLCJJRCI6MSwiVXNlcm5hbWUiOiJzdXBlcmFkbWluIiwiTmlja25hbWUiOiLotoXnuqfnrqHnkIblkZgiLCJJZGVudGl0eSI6MCwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6Imx5YmJuIiwiYXVkIjpbImx5YWRtaW4iXSwiZXhwIjoxNjg4MjA0NjgwLCJuYmYiOjE2ODc1OTk4ODAsImlhdCI6MTY4NzU5OTg4MH0.dvfU4M8NSEyU6uzdP8PJmOpLOt-qzQuJ_uFo0N9VHcQ'),(11,'2023-06-26 21:05:19.321','2023-06-26 21:05:19.321','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTBhYzZiMzQ5NTMwNGI0YzlhZDAxZWMxYTQzN2MyMmQiLCJJRCI6MSwiVXNlcm5hbWUiOiJzdXBlcmFkbWluIiwiTmlja25hbWUiOiLotoXnuqfnrqHnkIblkZgiLCJJZGVudGl0eSI6MSwiRGVwdElkIjowLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoibHliYm4iLCJhdWQiOlsibHlhZG1pbiJdLCJleHAiOjE2ODgzODkxMjAsIm5iZiI6MTY4Nzc4NDMyMCwiaWF0IjoxNjg3Nzg0MzIwfQ.ErK9Jaw0g7HITHmnB0U8ClJoRxX3kWDv43h6AGX6lUs'),(12,'2023-06-26 21:22:52.769','2023-06-26 21:22:52.769','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTBhYzZiMzQ5NTMwNGI0YzlhZDAxZWMxYTQzN2MyMmQiLCJJRCI6MSwiVXNlcm5hbWUiOiJzdXBlcmFkbWluIiwiTmlja25hbWUiOiLotoXnuqfnrqHnkIblkZgiLCJJZGVudGl0eSI6MSwiRGVwdElkIjowLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoibHliYm4iLCJhdWQiOlsibHlhZG1pbiJdLCJleHAiOjE2ODgzODk1MTksIm5iZiI6MTY4Nzc4NDcxOSwiaWF0IjoxNjg3Nzg0NzE5fQ.R9LZ1MhbeF9bzrdpZmfGD31A8fHdFCArsHHZp3Agvv0'),(13,'2023-06-26 22:02:12.362','2023-06-26 22:02:12.362','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTBhYzZiMzQ5NTMwNGI0YzlhZDAxZWMxYTQzN2MyMmQiLCJJRCI6MSwiVXNlcm5hbWUiOiJzdXBlcmFkbWluIiwiTmlja25hbWUiOiLotoXnuqfnrqHnkIblkZgiLCJJZGVudGl0eSI6MSwiRGVwdElkIjowLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoibHliYm4iLCJhdWQiOlsibHlhZG1pbiJdLCJleHAiOjE2ODgzOTA1NzIsIm5iZiI6MTY4Nzc4NTc3MiwiaWF0IjoxNjg3Nzg1NzcyfQ.CCKQrpqbWCfE45ef0TUS3gaF8UIWmfFMwdb7TzTxlvU'),(14,'2023-06-27 10:38:19.697','2023-06-27 10:38:19.697','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTBhYzZiMzQ5NTMwNGI0YzlhZDAxZWMxYTQzN2MyMmQiLCJJRCI6MSwiVXNlcm5hbWUiOiJzdXBlcmFkbWluIiwiTmlja25hbWUiOiLotoXnuqfnrqHnkIblkZgiLCJJZGVudGl0eSI6MSwiRGVwdElkIjowLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoibHliYm4iLCJhdWQiOlsibHlhZG1pbiJdLCJleHAiOjE2ODg0MzIyMTEsIm5iZiI6MTY4NzgyNzQxMSwiaWF0IjoxNjg3ODI3NDExfQ.yHP0EXTGf0SbBoj53kZr5c3dYIqVAWtaD9guTwoeN1U'),(15,'2023-06-27 12:13:29.903','2023-06-27 12:13:29.903','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTBhYzZiMzQ5NTMwNGI0YzlhZDAxZWMxYTQzN2MyMmQiLCJJRCI6MSwiVXNlcm5hbWUiOiJzdXBlcmFkbWluIiwiTmlja25hbWUiOiLotoXnuqfnrqHnkIblkZgiLCJJZGVudGl0eSI6MSwiRGVwdElkIjowLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoibHliYm4iLCJhdWQiOlsibHlhZG1pbiJdLCJleHAiOjE2ODg0MzgyOTksIm5iZiI6MTY4NzgzMzQ5OSwiaWF0IjoxNjg3ODMzNDk5fQ.pdDcE4VAmWY9mNknBnmDn6oPxuy9TBB6yEQQUFKwQNc'),(16,'2023-06-27 14:06:39.730','2023-06-27 14:06:39.730','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTBhYzZiMzQ5NTMwNGI0YzlhZDAxZWMxYTQzN2MyMmQiLCJJRCI6MSwiVXNlcm5hbWUiOiJzdXBlcmFkbWluIiwiTmlja25hbWUiOiLotoXnuqfnrqHnkIblkZgiLCJJZGVudGl0eSI6MSwiRGVwdElkIjowLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoibHliYm4iLCJhdWQiOlsibHlhZG1pbiJdLCJleHAiOjE2ODg0NDQwMDksIm5iZiI6MTY4NzgzOTIwOSwiaWF0IjoxNjg3ODM5MjA5fQ.lzxFN5po8GEajeRdiKjLAN0bt04jXKl6sEwF1h7F8CY'),(17,'2023-06-29 20:29:56.693','2023-06-29 20:29:56.693','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTQwMWMxN2U3MDE4NDc2NGIzMjdiZTg5NDJhMWFjODAiLCJJRCI6MiwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tuYW1lIjoi566h55CG5ZGYIiwiSWRlbnRpdHkiOjIsIkRlcHRJZCI6MSwiUm9sZUlkcyI6WzFdLCJSb2xlRGVwdElkcyI6WzFdLCJSb2xlRGF0YVNjb3BlcyI6WzBdLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoibHliYm4iLCJhdWQiOlsibHlhZG1pbiJdLCJleHAiOjE2ODg2NDY1NDYsIm5iZiI6MTY4ODA0MTc0NiwiaWF0IjoxNjg4MDQxNzQ2fQ.1s-Z4-Y5YrdHPxApTuh-xyZ2qKVaLvshvB6LNCwzI3I'),(18,'2023-06-29 21:29:00.767','2023-06-29 21:29:00.767','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTBhYzZiMzQ5NTMwNGI0YzlhZDAxZWMxYTQzN2MyMmQiLCJJRCI6MSwiVXNlcm5hbWUiOiJzdXBlcmFkbWluIiwiTmlja25hbWUiOiLotoXnuqfnrqHnkIblkZgiLCJJZGVudGl0eSI6MSwiRGVwdElkIjowLCJSb2xlSWRzIjpudWxsLCJSb2xlRGVwdElkcyI6W10sIlJvbGVEYXRhU2NvcGVzIjpbXSwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6Imx5YmJuIiwiYXVkIjpbImx5YWRtaW4iXSwiZXhwIjoxNjg4NjQ2NTYyLCJuYmYiOjE2ODgwNDE3NjIsImlhdCI6MTY4ODA0MTc2Mn0.fsZuE6f62FEZ2f-4XMNvbR3SuUwf6pC2Q936qmSgPo4'),(19,'2023-06-29 21:30:07.729','2023-06-29 21:30:07.729','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTQwMWMxN2U3MDE4NDc2NGIzMjdiZTg5NDJhMWFjODAiLCJJRCI6MiwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tuYW1lIjoi566h55CG5ZGYIiwiSWRlbnRpdHkiOjIsIkRlcHRJZCI6MSwiUm9sZUlkcyI6WzFdLCJSb2xlRGVwdElkcyI6WzFdLCJSb2xlRGF0YVNjb3BlcyI6WzBdLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoibHliYm4iLCJhdWQiOlsibHlhZG1pbiJdLCJleHAiOjE2ODg2NDY1OTYsIm5iZiI6MTY4ODA0MTc5NiwiaWF0IjoxNjg4MDQxNzk2fQ.aY8tI5A1z4yFn9wNKeDCxMNWjYUowdNRd5_EP2w5gik'),(20,'2023-06-29 21:35:41.316','2023-06-29 21:35:41.316','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTQwMWMxN2U3MDE4NDc2NGIzMjdiZTg5NDJhMWFjODAiLCJJRCI6MiwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tuYW1lIjoi566h55CG5ZGYIiwiSWRlbnRpdHkiOjIsIkRlcHRJZCI6MSwiUm9sZUlkcyI6WzFdLCJSb2xlRGVwdElkcyI6WzEsMiwzLDQsNSw2XSwiUm9sZURhdGFTY29wZXMiOls0XSwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6Imx5YmJuIiwiYXVkIjpbImx5YWRtaW4iXSwiZXhwIjoxNjg4NjUwMjA3LCJuYmYiOjE2ODgwNDU0MDcsImlhdCI6MTY4ODA0NTQwN30.xkh6xxi_GuHZlfsLlSzbuSYurD23Ookv649VjS5D40M'),(21,'2023-06-29 21:37:08.390','2023-06-29 21:37:08.390','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTBhYzZiMzQ5NTMwNGI0YzlhZDAxZWMxYTQzN2MyMmQiLCJJRCI6MSwiVXNlcm5hbWUiOiJzdXBlcmFkbWluIiwiTmlja25hbWUiOiLotoXnuqfnrqHnkIblkZgiLCJJZGVudGl0eSI6MSwiRGVwdElkIjowLCJSb2xlSWRzIjpudWxsLCJSb2xlRGVwdElkcyI6W10sIlJvbGVEYXRhU2NvcGVzIjpbXSwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6Imx5YmJuIiwiYXVkIjpbImx5YWRtaW4iXSwiZXhwIjoxNjg4NjUwMTQwLCJuYmYiOjE2ODgwNDUzNDAsImlhdCI6MTY4ODA0NTM0MH0.NQqm80k7S87RG-KwgOEQvLA2ZGctaD6wVDyYHOLSoww'),(22,'2023-06-29 21:37:40.363','2023-06-29 21:37:40.363','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTQwMWMxN2U3MDE4NDc2NGIzMjdiZTg5NDJhMWFjODAiLCJJRCI6MiwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tuYW1lIjoi566h55CG5ZGYIiwiSWRlbnRpdHkiOjIsIkRlcHRJZCI6MSwiUm9sZUlkcyI6WzFdLCJSb2xlRGVwdElkcyI6WzEsMiwzLDQsNSw2XSwiUm9sZURhdGFTY29wZXMiOls0XSwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6Imx5YmJuIiwiYXVkIjpbImx5YWRtaW4iXSwiZXhwIjoxNjg4NjUwNTQxLCJuYmYiOjE2ODgwNDU3NDEsImlhdCI6MTY4ODA0NTc0MX0.pl09rwmQFrj4V9WPyB59fdvcGCbDZcfA5h_WYDcmAt0'),(23,'2023-06-29 21:49:12.009','2023-06-29 21:49:12.009','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTBhYzZiMzQ5NTMwNGI0YzlhZDAxZWMxYTQzN2MyMmQiLCJJRCI6MSwiVXNlcm5hbWUiOiJzdXBlcmFkbWluIiwiTmlja25hbWUiOiLotoXnuqfnrqHnkIblkZgiLCJJZGVudGl0eSI6MSwiRGVwdElkIjowLCJSb2xlSWRzIjpudWxsLCJSb2xlRGVwdElkcyI6W10sIlJvbGVEYXRhU2NvcGVzIjpbXSwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6Imx5YmJuIiwiYXVkIjpbImx5YWRtaW4iXSwiZXhwIjoxNjg4NjUwNjI4LCJuYmYiOjE2ODgwNDU4MjgsImlhdCI6MTY4ODA0NTgyOH0.WoLZnFBJNWT2MmjVRO4oa2OaG-Z74AB9r1vWQp_u-xw'),(24,'2023-06-29 23:28:06.240','2023-06-29 23:28:06.240','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTQwMWMxN2U3MDE4NDc2NGIzMjdiZTg5NDJhMWFjODAiLCJJRCI6MiwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tuYW1lIjoi566h55CG5ZGYIiwiSWRlbnRpdHkiOjIsIkRlcHRJZCI6MSwiUm9sZUlkcyI6WzFdLCJSb2xlRGVwdElkcyI6WzEsMiwzLDQsNSw2XSwiUm9sZURhdGFTY29wZXMiOlsyXSwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6Imx5YmJuIiwiYXVkIjpbImx5YWRtaW4iXSwiZXhwIjoxNjg4NjUwNjYwLCJuYmYiOjE2ODgwNDU4NjAsImlhdCI6MTY4ODA0NTg2MH0.RUH2Kia5AFFQe07IEZmDdKBLoJheFfVSinh0M4usA9M'),(25,'2023-06-29 23:51:05.965','2023-06-29 23:51:05.965','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTBhYzZiMzQ5NTMwNGI0YzlhZDAxZWMxYTQzN2MyMmQiLCJJRCI6MSwiVXNlcm5hbWUiOiJzdXBlcmFkbWluIiwiTmlja25hbWUiOiLotoXnuqfnrqHnkIblkZgiLCJJZGVudGl0eSI6MSwiRGVwdElkIjowLCJSb2xlSWRzIjpudWxsLCJSb2xlRGVwdElkcyI6W10sIlJvbGVEYXRhU2NvcGVzIjpbXSwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6Imx5YmJuIiwiYXVkIjpbImx5YWRtaW4iXSwiZXhwIjoxNjg4NjUxMzUxLCJuYmYiOjE2ODgwNDY1NTEsImlhdCI6MTY4ODA0NjU1MX0.xfZm5kny8jh7Vst5vaqLBChDYxGQKfs57KUcobOxMSc'),(26,'2023-07-01 00:05:07.536','2023-07-01 00:05:07.536','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTBhYzZiMzQ5NTMwNGI0YzlhZDAxZWMxYTQzN2MyMmQiLCJJRCI6MSwiVXNlcm5hbWUiOiJzdXBlcmFkbWluIiwiTmlja25hbWUiOiLotoXnuqfnrqHnkIblkZgiLCJJZGVudGl0eSI6MSwiRGVwdElkIjowLCJSb2xlSWRzIjpudWxsLCJSb2xlRGVwdElkcyI6W10sIlJvbGVEYXRhU2NvcGVzIjpbXSwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6Imx5YmJuIiwiYXVkIjpbImx5YWRtaW4iXSwiZXhwIjoxNjg4NzQ1ODYzLCJuYmYiOjE2ODgxNDEwNjMsImlhdCI6MTY4ODE0MTA2M30.RaavoYJ-gK9Z5Why0YfJMIalWjKubfDGjj6QmDM6hJE'),(27,'2023-07-01 00:08:15.604','2023-07-01 00:08:15.604','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTBhYzZiMzQ5NTMwNGI0YzlhZDAxZWMxYTQzN2MyMmQiLCJJRCI6MSwiVXNlcm5hbWUiOiJzdXBlcmFkbWluIiwiTmlja25hbWUiOiLotoXnuqfnrqHnkIblkZgiLCJJZGVudGl0eSI6MSwiRGVwdElkIjowLCJSb2xlSWRzIjpudWxsLCJSb2xlRGVwdElkcyI6W10sIlJvbGVEYXRhU2NvcGVzIjpbXSwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6Imx5YmJuIiwiYXVkIjpbImx5YWRtaW4iXSwiZXhwIjoxNjg4NzQ1OTA3LCJuYmYiOjE2ODgxNDExMDcsImlhdCI6MTY4ODE0MTEwN30.-VdoJSEA8KjrRy_7Qzy7PwYdUyByLaTc3u_f4kSgzB8'),(28,'2023-07-01 00:13:52.117','2023-07-01 00:13:52.117','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTBhYzZiMzQ5NTMwNGI0YzlhZDAxZWMxYTQzN2MyMmQiLCJJRCI6MSwiVXNlcm5hbWUiOiJzdXBlcmFkbWluIiwiTmlja25hbWUiOiLotoXnuqfnrqHnkIblkZgiLCJJZGVudGl0eSI6MSwiRGVwdElkIjowLCJSb2xlSWRzIjpudWxsLCJSb2xlRGVwdElkcyI6W10sIlJvbGVEYXRhU2NvcGVzIjpbXSwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6Imx5YmJuIiwiYXVkIjpbImx5YWRtaW4iXSwiZXhwIjoxNjg4NzQ2MDk1LCJuYmYiOjE2ODgxNDEyOTUsImlhdCI6MTY4ODE0MTI5NX0.PZKcbJsDOKvRswJk2ivCVD8VzSiumi8Z8-ZgQ5R7jbI'),(29,'2023-07-01 00:23:54.487','2023-07-01 00:23:54.487','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTBhYzZiMzQ5NTMwNGI0YzlhZDAxZWMxYTQzN2MyMmQiLCJJRCI6MSwiVXNlcm5hbWUiOiJzdXBlcmFkbWluIiwiTmlja25hbWUiOiLotoXnuqfnrqHnkIblkZgiLCJJZGVudGl0eSI6MSwiRGVwdElkIjowLCJSb2xlSWRzIjpudWxsLCJSb2xlRGVwdElkcyI6W10sIlJvbGVEYXRhU2NvcGVzIjpbXSwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6Imx5YmJuIiwiYXVkIjpbImx5YWRtaW4iXSwiZXhwIjoxNjg4NzQ2NDMyLCJuYmYiOjE2ODgxNDE2MzIsImlhdCI6MTY4ODE0MTYzMn0.KbjUSBEER0pe1HqNDEmJ9moS7PLbQzWhMpAWlspaCXc'),(30,'2023-07-01 09:48:43.432','2023-07-01 09:48:43.432','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTBhYzZiMzQ5NTMwNGI0YzlhZDAxZWMxYTQzN2MyMmQiLCJJRCI6MSwiVXNlcm5hbWUiOiJzdXBlcmFkbWluIiwiTmlja25hbWUiOiLotoXnuqfnrqHnkIblkZgiLCJJZGVudGl0eSI6MSwiRGVwdElkIjowLCJSb2xlSWRzIjpudWxsLCJSb2xlRGVwdElkcyI6W10sIlJvbGVEYXRhU2NvcGVzIjpbXSwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6Imx5YmJuIiwiYXVkIjpbImx5YWRtaW4iXSwiZXhwIjoxNjg4NzQ3MDM0LCJuYmYiOjE2ODgxNDIyMzQsImlhdCI6MTY4ODE0MjIzNH0.2xlFkWox4TguLosVmF81JjgQyWw6V2sxoxvPE-5nnbg'),(31,'2023-07-01 10:33:59.252','2023-07-01 10:33:59.252','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTBhYzZiMzQ5NTMwNGI0YzlhZDAxZWMxYTQzN2MyMmQiLCJJRCI6MSwiVXNlcm5hbWUiOiJzdXBlcmFkbWluIiwiTmlja25hbWUiOiLotoXnuqfnrqHnkIblkZgiLCJJZGVudGl0eSI6MSwiRGVwdElkIjowLCJSb2xlSWRzIjpudWxsLCJSb2xlRGVwdElkcyI6W10sIlJvbGVEYXRhU2NvcGVzIjpbXSwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6Imx5YmJuIiwiYXVkIjpbImx5YWRtaW4iXSwiZXhwIjoxNjg4NzgwOTIzLCJuYmYiOjE2ODgxNzYxMjMsImlhdCI6MTY4ODE3NjEyM30._KqfowD1vP1R-D7ZBhBcwgrNw915I3jXSf9D9VMXU4I'),(32,'2023-07-01 10:36:05.256','2023-07-01 10:36:05.256','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTBhYzZiMzQ5NTMwNGI0YzlhZDAxZWMxYTQzN2MyMmQiLCJJRCI6MSwiVXNlcm5hbWUiOiJzdXBlcmFkbWluIiwiTmlja25hbWUiOiLotoXnuqfnrqHnkIblkZgiLCJJZGVudGl0eSI6MSwiRGVwdElkIjowLCJSb2xlSWRzIjpudWxsLCJSb2xlRGVwdElkcyI6W10sIlJvbGVEYXRhU2NvcGVzIjpbXSwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6Imx5YmJuIiwiYXVkIjpbImx5YWRtaW4iXSwiZXhwIjoxNjg4NzgzNjM5LCJuYmYiOjE2ODgxNzg4MzksImlhdCI6MTY4ODE3ODgzOX0.nly44jLNkecxHaBmH4HX9u52ea7gUO5DCyyx_1GpJPU'),(33,'2023-07-01 12:35:23.418','2023-07-01 12:35:23.418','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTBhYzZiMzQ5NTMwNGI0YzlhZDAxZWMxYTQzN2MyMmQiLCJJRCI6MSwiVXNlcm5hbWUiOiJzdXBlcmFkbWluIiwiTmlja25hbWUiOiLotoXnuqfnrqHnkIblkZgiLCJJZGVudGl0eSI6MSwiRGVwdElkIjowLCJSb2xlSWRzIjpudWxsLCJSb2xlRGVwdElkcyI6W10sIlJvbGVEYXRhU2NvcGVzIjpbXSwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6Imx5YmJuIiwiYXVkIjpbImx5YWRtaW4iXSwiZXhwIjoxNjg4NzgzNzY1LCJuYmYiOjE2ODgxNzg5NjUsImlhdCI6MTY4ODE3ODk2NX0.UtqGIViLvDig141KFfVfLh2ydWVLXhbkyaIrDjFPn0c');

#
# Structure for table "lyadmin_menu"
#

CREATE TABLE `lyadmin_menu` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `parent_id` bigint(20) unsigned DEFAULT NULL COMMENT '父菜单ID',
  `name` varchar(256) DEFAULT NULL COMMENT '菜单名称',
  `icon` varchar(256) DEFAULT NULL COMMENT '菜单图标',
  `web_path` varchar(256) DEFAULT NULL COMMENT '路由地址',
  `is_link` tinyint(1) DEFAULT '0' COMMENT '是否外链',
  `visible` tinyint(1) DEFAULT '1' COMMENT '是否显示菜单',
  `component` varchar(256) DEFAULT NULL COMMENT '对应前端文件路径',
  `component_name` varchar(256) DEFAULT NULL COMMENT '对应前端文件名称',
  `sort` bigint(20) DEFAULT '1' COMMENT '显示顺序',
  `is_catalog` tinyint(1) DEFAULT '0' COMMENT '是否目录',
  `keep_alive` tinyint(1) DEFAULT '0' COMMENT '是否缓存页面',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态',
  `create_by` bigint(20) unsigned DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) unsigned DEFAULT NULL COMMENT '更新者',
  `belong_dept` bigint(20) unsigned DEFAULT NULL COMMENT '数据归属部门',
  PRIMARY KEY (`id`),
  KEY `idx_lyadmin_menu_create_by` (`create_by`),
  KEY `idx_lyadmin_menu_update_by` (`update_by`),
  KEY `idx_lyadmin_menu_belong_dept` (`belong_dept`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4;

#
# Data for table "lyadmin_menu"
#

REPLACE INTO `lyadmin_menu` VALUES (2,'2023-06-26 15:57:36.000','2023-07-09 16:36:27.176',0,'管理员管理','avatar','adminManage',0,1,'','',20,0,0,1,NULL,1,NULL),(3,'2023-06-26 15:57:40.000','2023-07-09 16:37:04.440',0,'用户管理CRUD','UserFilled','userManageCrud',0,1,'','',30,0,0,1,NULL,1,NULL),(4,'2023-06-26 15:57:40.000','2023-06-29 23:53:56.118',0,'系统管理','tools','',0,1,'','',990,1,0,1,NULL,1,NULL),(5,'2023-06-26 15:57:40.000','2023-07-09 22:16:40.987',4,'菜单管理','Menu','menuManage',0,1,'','',2,0,0,1,NULL,1,NULL),(6,'2023-06-26 15:57:40.000','2023-07-09 22:18:34.306',4,'部门管理','Collection','departmentManage',0,1,'','',1,0,0,1,NULL,1,NULL),(7,'2023-06-27 12:09:22.049','2023-06-27 12:09:22.049',4,'操作日志','InfoFilled','journalManage',0,1,'','',99,0,0,1,1,0,0),(8,'2023-06-27 12:12:34.511','2023-07-09 22:16:17.860',4,'角色管理','Key','roleManage',0,1,'','',5,0,0,1,1,1,0),(9,'2023-06-28 21:53:41.125','2023-07-09 22:19:11.712',4,'权限管理','Lock','authorityManage',0,1,'','',7,0,0,1,1,1,0),(10,'2023-06-29 21:47:17.157','2023-07-09 16:38:50.064',0,'个人中心','Place','personalCenter',0,1,'','',866,0,0,1,2,1,1),(11,'2023-07-09 16:35:41.147','2023-07-09 16:36:15.662',0,'DashBoard','DataLine','analysis',0,1,'','',1,0,0,1,1,1,0);

#
# Structure for table "lyadmin_menu_button"
#

CREATE TABLE `lyadmin_menu_button` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `menu_id` bigint(20) unsigned DEFAULT NULL COMMENT '关联菜单ID',
  `name` varchar(256) DEFAULT NULL COMMENT '名称',
  `value` varchar(256) DEFAULT NULL COMMENT '权限值',
  `api` varchar(256) DEFAULT NULL COMMENT '接口地址',
  `method` varchar(256) DEFAULT NULL COMMENT '接口请求方法',
  `create_by` bigint(20) unsigned DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) unsigned DEFAULT NULL COMMENT '更新者',
  `belong_dept` bigint(20) unsigned DEFAULT NULL COMMENT '数据归属部门',
  PRIMARY KEY (`id`),
  KEY `idx_lyadmin_menu_button_create_by` (`create_by`),
  KEY `idx_lyadmin_menu_button_update_by` (`update_by`),
  KEY `idx_lyadmin_menu_button_belong_dept` (`belong_dept`)
) ENGINE=InnoDB AUTO_INCREMENT=45 DEFAULT CHARSET=utf8mb4;

#
# Data for table "lyadmin_menu_button"
#

REPLACE INTO `lyadmin_menu_button` VALUES (1,'2023-06-26 10:00:38.000','2023-06-26 10:00:38.000',1,'查询','Search','/api/system/user','GET',NULL,NULL,NULL),(2,'2023-06-26 10:00:38.000','2023-07-09 20:31:55.276',2,'查询','Search','/api/system/user/getAdminUserList','GET',0,1,0),(3,'2023-06-26 10:00:38.000','2023-07-09 20:31:55.276',5,'查询','Search','/api/system/menu/menu','GET',0,0,0),(4,'2023-06-26 10:00:38.000','2023-07-09 20:31:55.276',5,'新增','Create','/api/system/menu/menu','POST',0,0,0),(5,'2023-06-26 10:00:38.000','2023-07-09 20:31:55.276',5,'编辑','Update','/api/system/menu/menu/:id','PUT',0,0,0),(6,'2023-06-26 10:00:38.000','2023-07-09 20:31:55.276',5,'删除','Delete','/api/system/menu/menu/:id','DELETE',0,0,0),(7,'2023-06-27 09:56:05.167','2023-07-09 20:31:55.276',6,'查询','Search','/api/system/dept/dept','GET',1,1,0),(8,'2023-06-27 10:37:32.191','2023-07-09 20:31:55.276',6,'删除','Delete','/api/system/dept/dept/:id','DELETE',1,0,0),(9,'2023-06-27 10:37:51.367','2023-07-09 20:31:55.276',6,'编辑','Update','/api/system/dept/dept/:id','PUT',1,0,0),(10,'2023-06-27 10:37:59.918','2023-07-09 20:31:55.276',6,'新增','Create','/api/system/dept/dept','POST',1,0,0),(12,'2023-06-27 12:09:22.058','2023-07-09 20:31:55.276',7,'删除','Delete','/api/system/operation_log/log/:id','DELETE',1,1,0),(14,'2023-06-27 12:09:22.058','2023-07-09 20:31:55.276',7,'查询','Search','/api/system/operation_log/loglist','GET',1,1,0),(16,'2023-06-27 12:12:34.514','2023-07-09 20:31:55.276',8,'新增','Create','/api/system/role/role','POST',1,1,0),(17,'2023-06-27 12:12:34.514','2023-07-09 20:31:55.276',8,'删除','Delete','/api/system/role/role/:id','DELETE',1,1,0),(18,'2023-06-27 12:12:34.514','2023-07-09 20:31:55.276',8,'编辑','Update','/api/system/role/role/:id','PUT',1,1,0),(19,'2023-06-27 12:12:34.514','2023-07-09 20:31:55.276',8,'查询','Search','/api/system/role/roleList','GET',1,1,0),(21,'2023-06-28 21:53:41.138','2023-07-09 20:31:55.276',9,'保存','Save','/api/system/role/permission','PUT',1,1,0),(24,'2023-06-28 21:53:41.138','2023-07-09 20:31:55.276',9,'查询','Search','/api/system/role/role','GET',1,1,0),(25,'2023-06-28 21:53:41.138','2023-07-09 20:31:55.276',9,'详情','Detail','/api/system/role/role_id_to_menu/:id','GET',1,1,0),(28,'2023-06-29 21:47:17.160','2023-07-09 20:31:55.276',10,'编辑','Update','/api/system/user/setUserInfo','POST',2,1,1),(29,'2023-06-29 21:47:17.160','2023-07-09 20:31:55.276',10,'查询','Search','/api/system/user/getUserInfo','GET',2,1,1),(31,'2023-07-01 00:07:57.490','2023-07-09 20:31:55.276',10,'修改密码','Changepassword','/api/system/user/changePassword','POST',1,1,0),(35,'2023-07-09 16:35:41.155','2023-07-09 20:31:55.276',11,'查询','Search','','GET',1,0,0),(36,'2023-07-09 16:35:41.155','2023-07-09 20:31:55.276',11,'详情','Detail','','GET',1,0,0),(37,'2023-07-09 16:43:49.010','2023-07-09 20:31:55.276',2,'新增','Create','/api/system/user/adminUser','POST',1,0,0),(38,'2023-07-09 16:44:36.787','2023-07-09 20:31:55.276',2,'编辑','Update','/api/system/user/adminUser/:id','PUT',1,0,0),(39,'2023-07-09 16:44:47.303','2023-07-09 20:31:55.276',2,'删除','Delete','/api/system/user/adminUser/:id','DELETE',1,0,0),(40,'2023-07-09 20:30:20.229','2023-07-09 20:31:55.276',3,'新增','Create','/api/user/user/users','POST',1,0,0),(41,'2023-07-09 20:30:37.089','2023-07-09 20:31:55.276',3,'编辑','Update','/api/user/user/users/:id','PUT',1,0,0),(42,'2023-07-09 20:30:44.485','2023-07-09 20:31:55.276',3,'删除','Delete','/api/user/user/users/:id','DELETE',1,0,0),(43,'2023-07-09 20:31:02.410','2023-07-09 20:31:55.276',3,'查询','Search','/api/user/user/getUserList','GET',1,0,0),(44,'2023-07-09 20:31:35.822','2023-07-09 20:31:55.276',3,'详情','Detail','/api/user/user/users/:id','GET',1,0,0);

#
# Structure for table "lyadmin_operation_log"
#

CREATE TABLE `lyadmin_operation_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `ip` varchar(50) DEFAULT NULL COMMENT '请求ip',
  `method` varchar(8) DEFAULT NULL COMMENT '请求方法',
  `path` varchar(256) DEFAULT NULL COMMENT '请求路径',
  `code` varchar(32) DEFAULT NULL COMMENT '请求状态',
  `latency` bigint(20) DEFAULT NULL COMMENT '延迟',
  `agent` varchar(256) DEFAULT NULL COMMENT 'UserAgent代理',
  `msg` varchar(256) DEFAULT NULL COMMENT '返回信息',
  `body` text COMMENT '请求Body',
  `resp` text COMMENT '响应Body',
  `user_id` bigint(20) DEFAULT NULL COMMENT '用户id',
  `create_by` bigint(20) unsigned DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) unsigned DEFAULT NULL COMMENT '更新者',
  `belong_dept` bigint(20) unsigned DEFAULT NULL COMMENT '数据归属部门',
  PRIMARY KEY (`id`),
  KEY `idx_lyadmin_operation_log_create_by` (`create_by`),
  KEY `idx_lyadmin_operation_log_update_by` (`update_by`),
  KEY `idx_lyadmin_operation_log_belong_dept` (`belong_dept`)
) ENGINE=InnoDB AUTO_INCREMENT=1269 DEFAULT CHARSET=utf8mb4;

#
# Data for table "lyadmin_operation_log"
#

REPLACE INTO `lyadmin_operation_log` VALUES (1268,'2023-07-09 22:19:18.859','2023-07-09 22:19:18.859','127.0.0.1','DELETE','/api/system/operation_log/deletealllogs','200',9977100,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36','','{}','{\"code\":2000,\"data\":null,\"msg\":\"清空成功\"}',1,0,0,0);

#
# Structure for table "lyadmin_post"
#

CREATE TABLE `lyadmin_post` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `name` varchar(256) DEFAULT NULL COMMENT '岗位名称',
  `code` varchar(100) DEFAULT NULL COMMENT '岗位编码',
  `sort` bigint(20) DEFAULT '1' COMMENT '显示顺序',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态',
  `create_by` bigint(20) unsigned DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) unsigned DEFAULT NULL COMMENT '更新者',
  `belong_dept` bigint(20) unsigned DEFAULT NULL COMMENT '数据归属部门',
  PRIMARY KEY (`id`),
  KEY `idx_lyadmin_post_update_by` (`update_by`),
  KEY `idx_lyadmin_post_belong_dept` (`belong_dept`),
  KEY `idx_lyadmin_post_create_by` (`create_by`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

#
# Data for table "lyadmin_post"
#


#
# Structure for table "lyadmin_role"
#

CREATE TABLE `lyadmin_role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `name` varchar(256) DEFAULT NULL COMMENT '角色名称',
  `key` varchar(256) DEFAULT NULL COMMENT '权限字符',
  `sort` bigint(20) DEFAULT '1' COMMENT '显示顺序',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态',
  `create_by` bigint(20) unsigned DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) unsigned DEFAULT NULL COMMENT '更新者',
  `belong_dept` bigint(20) unsigned DEFAULT NULL COMMENT '数据归属部门',
  `data_range` bigint(20) DEFAULT '0' COMMENT '数据权限范围',
  PRIMARY KEY (`id`),
  UNIQUE KEY `key` (`key`),
  UNIQUE KEY `key_2` (`key`),
  KEY `idx_lyadmin_role_update_by` (`update_by`),
  KEY `idx_lyadmin_role_belong_dept` (`belong_dept`),
  KEY `idx_lyadmin_role_key` (`key`),
  KEY `idx_lyadmin_role_create_by` (`create_by`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

#
# Data for table "lyadmin_role"
#

REPLACE INTO `lyadmin_role` VALUES (1,'2023-06-26 10:51:03.000','2023-07-09 20:31:55.275','管理员','admin',1,1,0,1,0,3);

#
# Structure for table "lyadmin_role_dept"
#

CREATE TABLE `lyadmin_role_dept` (
  `lyadmin_role_id` bigint(20) NOT NULL COMMENT '主键',
  `lyadmin_dept_id` bigint(20) NOT NULL COMMENT '主键',
  PRIMARY KEY (`lyadmin_role_id`,`lyadmin_dept_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

#
# Data for table "lyadmin_role_dept"
#


#
# Structure for table "lyadmin_role_menu"
#

CREATE TABLE `lyadmin_role_menu` (
  `lyadmin_role_id` bigint(20) NOT NULL COMMENT '主键',
  `lyadmin_menu_id` bigint(20) NOT NULL COMMENT '主键',
  PRIMARY KEY (`lyadmin_role_id`,`lyadmin_menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

#
# Data for table "lyadmin_role_menu"
#

REPLACE INTO `lyadmin_role_menu` VALUES (1,2),(1,3),(1,4),(1,5),(1,6),(1,7),(1,8),(1,9),(1,10),(1,11);

#
# Structure for table "lyadmin_role_menubutton"
#

CREATE TABLE `lyadmin_role_menubutton` (
  `lyadmin_role_id` bigint(20) NOT NULL COMMENT '主键',
  `lyadmin_menu_button_id` bigint(20) NOT NULL COMMENT '主键',
  PRIMARY KEY (`lyadmin_role_id`,`lyadmin_menu_button_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

#
# Data for table "lyadmin_role_menubutton"
#

REPLACE INTO `lyadmin_role_menubutton` VALUES (1,2),(1,3),(1,4),(1,5),(1,6),(1,7),(1,8),(1,9),(1,10),(1,11),(1,12),(1,13),(1,14),(1,15),(1,16),(1,17),(1,18),(1,19),(1,20),(1,21),(1,22),(1,23),(1,24),(1,25),(1,26),(1,27),(1,28),(1,29),(1,30),(1,31),(1,35),(1,36),(1,37),(1,38),(1,39),(1,40),(1,41),(1,42),(1,43),(1,44);

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
  `name` varchar(20) DEFAULT NULL COMMENT '姓名',
  `nickname` varchar(20) DEFAULT NULL COMMENT '昵称',
  `mobile` char(25) DEFAULT NULL COMMENT '手机号',
  `email` varchar(100) DEFAULT NULL COMMENT '邮箱',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `gender` varchar(10) DEFAULT '男' COMMENT '性别',
  `dept_id` bigint(20) DEFAULT NULL COMMENT '部门',
  `is_staff` tinyint(1) DEFAULT '1' COMMENT '是否可登录后台',
  `is_superuser` tinyint(1) DEFAULT '0' COMMENT '是否超管',
  `is_active` tinyint(1) DEFAULT '1' COMMENT '状态(1正常、0冻结)',
  `identity` tinyint(4) DEFAULT '2' COMMENT '身份(1 超级管理员 、2后台、3前台)',
  `create_by` bigint(20) unsigned DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) unsigned DEFAULT NULL COMMENT '更新者',
  `belong_dept` bigint(20) unsigned DEFAULT NULL COMMENT '数据归属部门',
  `is_delete` tinyint(1) DEFAULT '0' COMMENT '是否删除(1删除、0正常)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  KEY `idx_lyadmin_users_username` (`username`),
  KEY `idx_lyadmin_users_create_by` (`create_by`),
  KEY `idx_lyadmin_users_update_by` (`update_by`),
  KEY `idx_lyadmin_users_belong_dept` (`belong_dept`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

#
# Data for table "lyadmin_users"
#

REPLACE INTO `lyadmin_users` VALUES (1,'2023-06-26 08:37:21.068','2023-07-01 00:25:47.008','10ac6b3495304b4c9ad01ec1a437c22d','superadmin','$2a$10$8bFSiG0THdgR2al7yfQBu.kPhS5NGKfZo/C5J2DId8KY5CmpPzjga','超级管理员','超级管理员','18000000000','','','男',0,1,1,1,1,0,0,0,0),(2,'2023-06-26 10:44:48.690','2023-07-09 16:14:08.828','1401c17e70184764b327be8942a1ac80','admin','$2a$10$aWM7YczX8hq5htpe1yh2v.6TxIGbZhTsPOg1h4U4qS9.XJBiiPyU.','管理员','管理员','18000000000','','','男',1,1,0,1,2,0,1,0,0),(3,'2023-07-09 19:56:41.488','2023-07-09 20:29:43.863','653505e254fe4a2eaf9179a82d80a695','test','$2a$10$HGpnhxNuLBwqt1HJ5gbWnOVof4b6tVIFRDN3/VSQw43hwnwmnIer2','','测试前端用户','18000000000','','','男',0,1,0,1,3,1,1,0,0);

#
# Structure for table "lyadmin_users_post"
#

CREATE TABLE `lyadmin_users_post` (
  `lyadmin_users_id` bigint(20) NOT NULL COMMENT '主键',
  `lyadmin_post_id` bigint(20) NOT NULL COMMENT '主键',
  PRIMARY KEY (`lyadmin_users_id`,`lyadmin_post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

#
# Data for table "lyadmin_users_post"
#


#
# Structure for table "lyadmin_users_role"
#

CREATE TABLE `lyadmin_users_role` (
  `lyadmin_users_id` bigint(20) NOT NULL COMMENT '主键',
  `lyadmin_role_id` bigint(20) NOT NULL COMMENT '主键',
  PRIMARY KEY (`lyadmin_users_id`,`lyadmin_role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

#
# Data for table "lyadmin_users_role"
#

REPLACE INTO `lyadmin_users_role` VALUES (2,1);
