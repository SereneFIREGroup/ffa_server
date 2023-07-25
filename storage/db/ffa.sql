CREATE TABLE `family` (
    `uuid` varchar(8) COLLATE latin1_bin NOT NULL,
    `create_time` bigint(20) NOT NULL DEFAULT '0',
    `status` tinyint(4) NOT NULL DEFAULT '-1',
    `name` varchar(128) CHARACTER SET utf8mb4 NOT NULL,
    `owner` varchar(8) COLLATE latin1_bin DEFAULT NULL,
    `fire_goal` DECIMAL(10, 2),
    PRIMARY KEY (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

CREATE TABLE `user` (
    `family_uuid` varchar(8) CHARACTER SET latin1 COLLATE latin1_bin NOT NULL DEFAULT '',
    `uuid` varchar(8) CHARACTER SET latin1 COLLATE latin1_bin NOT NULL DEFAULT '',
    `name` varchar(128) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '姓名',
    `avatar` varchar(255) CHARACTER SET latin1 COLLATE latin1_bin DEFAULT '' COMMENT '头像',
    `phone` varchar(32) CHARACTER SET latin1 COLLATE latin1_bin DEFAULT NULL COMMENT '手机号码',
    `password` varchar(60) CHARACTER SET latin1 COLLATE latin1_bin DEFAULT '' COMMENT '密码',
    `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '1.正常 2.删除的 3.待激活 4.禁用的（被管理员禁用）',
    `create_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建时间',
    PRIMARY KEY (`family_uuid`,`uuid`),
    UNIQUE KEY `unq_phone` (`phone`),
    KEY `idx_uuid` (`uuid`),
    KEY `idx_phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `session` (
    `user_uuid` varchar(8) COLLATE latin1_bin NOT NULL,
    `token` varchar(64) COLLATE latin1_bin NOT NULL,
    PRIMARY KEY (`user_uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

create table `verify_code` (
    `verify_type` tinyint(4) NOT NULL COMMENT '1.email,2.phone',
    `number` varchar(255) COLLATE latin1_bin NOT NULL,
    `code` varchar(32) COLLATE latin1_bin NOT NULL,
    `create_time` bigint(20) NOT NULL DEFAULT '0',
    `ip` varchar(32) COLLATE latin1_bin NOT NULL,
    `status` tinyint(4) NOT NULL DEFAULT 0,
    PRIMARY KEY (`verify_type`,`number`,`code`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_bin;