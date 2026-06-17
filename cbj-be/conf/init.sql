CREATE DATABASE `cbj-launchpad` DEFAULT CHARSET utf8mb4;

CREATE TABLE `product_contract` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` varchar(80) NOT NULL COMMENT '项目名称',
  `description` longtext COMMENT '项目描述',
  `img` varchar(500) DEFAULT NULL COMMENT '项目图标地址',
  `twitter_name` varchar(40) DEFAULT NULL,
  `status` int NOT NULL DEFAULT '0' COMMENT '项目状态',
  `amount` varchar(40) DEFAULT NULL COMMENT '当前币种质押个数',
  `sale_contract_address` varchar(42) DEFAULT NULL COMMENT 'sale合约地址',
  `token_address` varchar(42) DEFAULT NULL COMMENT 'bre合约地址',
  `payment_token` varchar(42) DEFAULT NULL COMMENT '支付地址',
  `follower` int NOT NULL DEFAULT '0' COMMENT 'ins或推特的follow数',
  `tge` datetime DEFAULT NULL COMMENT 'tge、时间',
  `project_website` varchar(500) DEFAULT NULL COMMENT 'projectWebsite',
  `about_html` varchar(5000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT 'about_html',
  `registration_time_starts` datetime DEFAULT NULL COMMENT '开始时间',
  `registration_time_ends` datetime DEFAULT NULL COMMENT '结束时间',
  `sale_start` datetime DEFAULT NULL COMMENT 'sale开始时间',
  `sale_end` datetime DEFAULT NULL COMMENT 'sale结束时间',
  `max_participation` varchar(40) DEFAULT NULL COMMENT '硬顶',
  `token_price_in_PT` varchar(40) DEFAULT NULL COMMENT 'Token price',
  `total_tokens_sold` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '所有已卖的token个数',
  `amount_of_tokens_to_sell` varchar(60) DEFAULT NULL COMMENT '质押币种',
  `total_raised` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '质押币种单位',
  `symbol` varchar(60) DEFAULT NULL COMMENT '币种单位（缩写）',
  `decimals` int DEFAULT '8' COMMENT '数位',
  `unlock_time` datetime DEFAULT NULL COMMENT '解锁时间',
  `medias` varchar(200) DEFAULT NULL COMMENT '媒体链接',
  `number_of_registrants` int DEFAULT NULL COMMENT '注册人数',
  `vesting` varchar(40) DEFAULT NULL,
  `tricker` varchar(40) DEFAULT NULL,
  `token_name` varchar(20) DEFAULT NULL COMMENT 'token名',
  `roi` varchar(20) DEFAULT NULL COMMENT 'roi',
  `vesting_portions_unlock_time` varchar(60) DEFAULT NULL,
  `vesting_percent_per_portion` varchar(60) DEFAULT NULL,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `type` int DEFAULT NULL COMMENT '项目类型',
  `card_link` varchar(200) DEFAULT NULL COMMENT '主页卡片跳转链接',
  `tweet_id` varchar(40) DEFAULT NULL COMMENT '转推任务的推文ID',
  `chain_id` int DEFAULT '0' COMMENT '链ChainID',
  `payment_token_decimals` int DEFAULT '8',
  `current_price` bigint DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `product_participation` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `account_id` varchar(64) NOT NULL COMMENT '用户钱包地址',
  `product_id` bigint unsigned NOT NULL COMMENT '项目ID',
  `amount_bought` varchar(80) NOT NULL DEFAULT '0' COMMENT '购买的代币数量(wei)',
  `amount_paid` varchar(80) NOT NULL DEFAULT '0' COMMENT '支付的ETH数量(wei)',
  `tx_hash` varchar(80) DEFAULT NULL COMMENT '认购交易哈希',
  `block_number` bigint unsigned DEFAULT '0' COMMENT '认购所在区块',
  `participated_at` datetime DEFAULT NULL COMMENT '认购时间',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态:1-有效',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_account_product` (`account_id`,`product_id`) COMMENT '同一用户同一项目只能认购一次',
  KEY `idx_product_id` (`product_id`),
  KEY `idx_account_id` (`account_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户项目认购明细表';

CREATE TABLE `product_registration` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `account_id` varchar(64) NOT NULL COMMENT '用户钱包地址',
  `product_id` bigint unsigned NOT NULL COMMENT '项目ID',
  `referral_code` varchar(64) DEFAULT NULL COMMENT '推荐码(可空)',
  `stake_amount` varchar(80) DEFAULT '0' COMMENT '注册时质押量(wei,字符串存大数)',
  `allocation` varchar(80) DEFAULT '0' COMMENT '分配的认购额度(wei)',
  `tx_hash` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '链上注册交易哈希',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态:1-已注册,0-无效',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_account_product` (`account_id`,`product_id`) COMMENT '同一用户同一项目只能注册一次',
  KEY `idx_product_id` (`product_id`),
  KEY `idx_account_id` (`account_id`)
) ENGINE=InnoDB AUTO_INCREMENT=64 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户项目注册表';


CREATE TABLE `sync_state` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(64) NOT NULL COMMENT '同步任务标识,如 register_listener',
  `last_block` bigint unsigned NOT NULL DEFAULT '0' COMMENT '最后处理的区块高度',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_name` (`name`) COMMENT '任务标识唯一'
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='链上事件同步状态表';