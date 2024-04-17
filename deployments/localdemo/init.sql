CREATE TABLE IF NOT EXISTS `users`(
    `account_id` INT AUTO_INCREMENT,
    `account` VARCHAR(50) UNIQUE NOT NULL,
    `user_name` VARCHAR(50) NOT NULL,
    `pwd_hash` VARCHAR(100) NOT NULL,
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `status` SMALLINT DEFAULT 0 COMMENT '0: not-in-use, 1: not-active, 2: in-use, 3: black-list',
    `role_bitmap` BIGINT DEFAULT 1,
    PRIMARY KEY (`account_id`),
    UNIQUE KEY `account_key` (`account`),
    UNIQUE KEY `user_name_key` (`user_name`)
) DEFAULT CHARSET=utf8;