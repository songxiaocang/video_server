CREATE TABLE `comments` (
  `id` varchar(64) NOT NULL COMMENT '主键',
  `video_id` varchar(255) DEFAULT NULL COMMENT '视频id',
  `author_id` int(10) DEFAULT NULL COMMENT '作者id',
  `content` text COMMENT '内容',
  `time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `sessions` (
  `session_id` varchar(64) NOT NULL COMMENT 'session_id',
  `TTL` tinytext COMMENT 'TTL权限',
  `login_name` text COMMENT '用户名',
  PRIMARY KEY (`session_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `login_name` varchar(64) DEFAULT NULL COMMENT '用户名',
  `pwd` text COMMENT '密码',
  PRIMARY KEY (`id`),
  UNIQUE KEY `login_name` (`login_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `video_del_rec` (
  `video_id` varchar(64) NOT NULL,
  PRIMARY KEY (`video_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `video_info` (
  `id` varchar(64) NOT NULL COMMENT '主键',
  `author_id` int(10) DEFAULT NULL COMMENT '作者id',
  `name` text COMMENT '名字',
  `display_ctime` text COMMENT '展示时间',
  `create_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

