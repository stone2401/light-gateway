/*Generated by xorm 2023-07-16 20:57:42, from mysql to mysql*/

SET sql_mode='NO_BACKSLASH_ESCAPES';
CREATE TABLE IF NOT EXISTS `gateway_admin` (`id` BIGINT(20) PRIMARY KEY AUTO_INCREMENT NOT NULL COMMENT '自增id', `user_name` VARCHAR(255) NOT NULL COMMENT '用户名', `salt` VARCHAR(255) NOT NULL COMMENT '盐', `password` VARCHAR(255) NOT NULL COMMENT '密码', `create_at` DATETIME NOT NULL COMMENT '创建时间', `update_at` DATETIME NOT NULL COMMENT '更新时间', `delete_at` DATETIME NULL COMMENT '删除时间', `is_delete` INT NULL) ENGINE=InnoDB DEFAULT CHARSET utf8;
CREATE UNIQUE INDEX `UQE_gateway_admin_user_name` ON `gateway_admin` (`user_name`);
INSERT INTO `gateway_admin` (`id`, `user_name`, `salt`, `password`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (1,'admin','2823d896e9822c0833d41d4904f0c00756d718570fce49b9a379a62c804689d3','admin','2023-07-16 16:20:08','2023-07-16 16:20:08',NULL,0);

CREATE TABLE IF NOT EXISTS `gateway_app` (`id` BIGINT(20) PRIMARY KEY AUTO_INCREMENT NOT NULL COMMENT '自增id', `app_id` VARCHAR(255) DEFAULT '' NOT NULL COMMENT '租户id', `name` VARCHAR(255) DEFAULT '租户名称' NOT NULL, `secret` VARCHAR(255) DEFAULT '' NOT NULL COMMENT '密钥', `white_ips` VARCHAR(1000) DEFAULT '' NOT NULL COMMENT 'ip白名单，支持前缀匹配', `qpd` BIGINT(20) DEFAULT 0 NOT NULL COMMENT '日请求量限制', `qps` BIGINT(20) DEFAULT 0 NOT NULL COMMENT '每秒请求量限制', `create_at` DATETIME NULL COMMENT '添加时间', `update_at` DATETIME NULL COMMENT '更新时间', `delete_at` DATETIME NULL COMMENT '删除时间', `is_delete` INT DEFAULT 0 NOT NULL COMMENT '是否删除 1=删除') ENGINE=InnoDB DEFAULT CHARSET utf8;
INSERT INTO `gateway_app` (`id`, `app_id`, `name`, `secret`, `white_ips`, `qpd`, `qps`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (1,'app_id_a','租户A','449441eb5e72dca9c42a12f3924ea3a2','',100000,100,'2023-07-16 16:20:08','2023-07-16 16:20:08',NULL,0);
INSERT INTO `gateway_app` (`id`, `app_id`, `name`, `secret`, `white_ips`, `qpd`, `qps`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (2,'app_id_b','租户B','8d7b11ec9be0e59a36b52f32366c09cb','',0,0,'2023-07-16 16:20:08','2023-07-16 16:20:08',NULL,0);
INSERT INTO `gateway_app` (`id`, `app_id`, `name`, `secret`, `white_ips`, `qpd`, `qps`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (3,'app_id','租户名称','','',0,0,'2023-07-16 16:20:08','2023-07-16 16:20:08',NULL,0);
INSERT INTO `gateway_app` (`id`, `app_id`, `name`, `secret`, `white_ips`, `qpd`, `qps`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (4,'app_id45','名称','07d980f8a49347523ee1d5c1c41aec02','',0,0,'2023-07-16 16:20:08','2023-07-16 16:20:08',NULL,0);

CREATE TABLE IF NOT EXISTS `gateway_service_access_control` (`id` BIGINT(20) PRIMARY KEY AUTO_INCREMENT NOT NULL COMMENT '自增主键', `service_id` BIGINT(20) DEFAULT 0 NOT NULL COMMENT '服务id', `open_auth` TINYINT DEFAULT 0 NOT NULL COMMENT '是否开启权限 1=开启', `black_list` VARCHAR(1000) DEFAULT '' NOT NULL COMMENT '黑名单', `white_list` VARCHAR(1000) DEFAULT '' NOT NULL COMMENT '白名单', `white_host_name` VARCHAR(1000) DEFAULT '' NOT NULL COMMENT '白名单主机', `clientip_flow_limit` INT DEFAULT 0 NOT NULL COMMENT '客户端ip限流', `service_flow_limit` INT DEFAULT 0 NOT NULL COMMENT '服务的限流') ENGINE=InnoDB DEFAULT CHARSET utf8;
INSERT INTO `gateway_service_access_control` (`id`, `service_id`, `open_auth`, `black_list`, `white_list`, `white_host_name`, `clientip_flow_limit`, `service_flow_limit`) VALUES (162,35,1,'','','',0,0);
INSERT INTO `gateway_service_access_control` (`id`, `service_id`, `open_auth`, `black_list`, `white_list`, `white_host_name`, `clientip_flow_limit`, `service_flow_limit`) VALUES (165,34,0,'','','',0,0);
INSERT INTO `gateway_service_access_control` (`id`, `service_id`, `open_auth`, `black_list`, `white_list`, `white_host_name`, `clientip_flow_limit`, `service_flow_limit`) VALUES (167,36,0,'','','',0,0);
INSERT INTO `gateway_service_access_control` (`id`, `service_id`, `open_auth`, `black_list`, `white_list`, `white_host_name`, `clientip_flow_limit`, `service_flow_limit`) VALUES (168,38,1,'111.11','22.33','11.11',12,12);
INSERT INTO `gateway_service_access_control` (`id`, `service_id`, `open_auth`, `black_list`, `white_list`, `white_host_name`, `clientip_flow_limit`, `service_flow_limit`) VALUES (169,41,1,'111.11','22.33','11.11',12,12);
INSERT INTO `gateway_service_access_control` (`id`, `service_id`, `open_auth`, `black_list`, `white_list`, `white_host_name`, `clientip_flow_limit`, `service_flow_limit`) VALUES (170,42,1,'111.11','22.33','11.11',12,12);
INSERT INTO `gateway_service_access_control` (`id`, `service_id`, `open_auth`, `black_list`, `white_list`, `white_host_name`, `clientip_flow_limit`, `service_flow_limit`) VALUES (171,43,0,'111.11','22.33','11.11',12,12);
INSERT INTO `gateway_service_access_control` (`id`, `service_id`, `open_auth`, `black_list`, `white_list`, `white_host_name`, `clientip_flow_limit`, `service_flow_limit`) VALUES (172,44,0,'','','',0,0);
INSERT INTO `gateway_service_access_control` (`id`, `service_id`, `open_auth`, `black_list`, `white_list`, `white_host_name`, `clientip_flow_limit`, `service_flow_limit`) VALUES (173,45,0,'','','',0,0);
INSERT INTO `gateway_service_access_control` (`id`, `service_id`, `open_auth`, `black_list`, `white_list`, `white_host_name`, `clientip_flow_limit`, `service_flow_limit`) VALUES (174,46,0,'','','',0,0);
INSERT INTO `gateway_service_access_control` (`id`, `service_id`, `open_auth`, `black_list`, `white_list`, `white_host_name`, `clientip_flow_limit`, `service_flow_limit`) VALUES (175,47,0,'','','',0,0);
INSERT INTO `gateway_service_access_control` (`id`, `service_id`, `open_auth`, `black_list`, `white_list`, `white_host_name`, `clientip_flow_limit`, `service_flow_limit`) VALUES (176,48,0,'','','',0,0);
INSERT INTO `gateway_service_access_control` (`id`, `service_id`, `open_auth`, `black_list`, `white_list`, `white_host_name`, `clientip_flow_limit`, `service_flow_limit`) VALUES (177,49,0,'','','',0,0);
INSERT INTO `gateway_service_access_control` (`id`, `service_id`, `open_auth`, `black_list`, `white_list`, `white_host_name`, `clientip_flow_limit`, `service_flow_limit`) VALUES (178,50,0,'','','',0,0);
INSERT INTO `gateway_service_access_control` (`id`, `service_id`, `open_auth`, `black_list`, `white_list`, `white_host_name`, `clientip_flow_limit`, `service_flow_limit`) VALUES (179,51,0,'','','',0,0);
INSERT INTO `gateway_service_access_control` (`id`, `service_id`, `open_auth`, `black_list`, `white_list`, `white_host_name`, `clientip_flow_limit`, `service_flow_limit`) VALUES (180,52,0,'','','',0,0);
INSERT INTO `gateway_service_access_control` (`id`, `service_id`, `open_auth`, `black_list`, `white_list`, `white_host_name`, `clientip_flow_limit`, `service_flow_limit`) VALUES (181,53,0,'','','',0,0);
INSERT INTO `gateway_service_access_control` (`id`, `service_id`, `open_auth`, `black_list`, `white_list`, `white_host_name`, `clientip_flow_limit`, `service_flow_limit`) VALUES (182,54,1,'127.0.0.3','127.0.0.2','',11,12);
INSERT INTO `gateway_service_access_control` (`id`, `service_id`, `open_auth`, `black_list`, `white_list`, `white_host_name`, `clientip_flow_limit`, `service_flow_limit`) VALUES (183,55,1,'127.0.0.2','127.0.0.1','',45,34);
INSERT INTO `gateway_service_access_control` (`id`, `service_id`, `open_auth`, `black_list`, `white_list`, `white_host_name`, `clientip_flow_limit`, `service_flow_limit`) VALUES (184,56,0,'192.168.1.0','','',0,0);
INSERT INTO `gateway_service_access_control` (`id`, `service_id`, `open_auth`, `black_list`, `white_list`, `white_host_name`, `clientip_flow_limit`, `service_flow_limit`) VALUES (185,57,0,'','127.0.0.1,127.0.0.2','',0,0);
INSERT INTO `gateway_service_access_control` (`id`, `service_id`, `open_auth`, `black_list`, `white_list`, `white_host_name`, `clientip_flow_limit`, `service_flow_limit`) VALUES (186,58,1,'','','',0,0);
INSERT INTO `gateway_service_access_control` (`id`, `service_id`, `open_auth`, `black_list`, `white_list`, `white_host_name`, `clientip_flow_limit`, `service_flow_limit`) VALUES (187,59,1,'127.0.0.1','','',2,3);
INSERT INTO `gateway_service_access_control` (`id`, `service_id`, `open_auth`, `black_list`, `white_list`, `white_host_name`, `clientip_flow_limit`, `service_flow_limit`) VALUES (188,60,1,'','','',0,0);
INSERT INTO `gateway_service_access_control` (`id`, `service_id`, `open_auth`, `black_list`, `white_list`, `white_host_name`, `clientip_flow_limit`, `service_flow_limit`) VALUES (189,61,0,'','','',0,0);

CREATE TABLE IF NOT EXISTS `gateway_service_grpc_rule` (`id` BIGINT(20) PRIMARY KEY AUTO_INCREMENT NOT NULL COMMENT '自增主键', `service_id` BIGINT(20) DEFAULT 0 NOT NULL, `port` INT NOT NULL COMMENT '端口', `header_transfor` VARCHAR(5000) DEFAULT '' NOT NULL COMMENT 'header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔') ENGINE=InnoDB DEFAULT CHARSET utf8;
CREATE UNIQUE INDEX `UQE_gateway_service_grpc_rule_service_id` ON `gateway_service_grpc_rule` (`service_id`);
INSERT INTO `gateway_service_grpc_rule` (`id`, `service_id`, `port`, `header_transfor`) VALUES (171,53,8009,'');
INSERT INTO `gateway_service_grpc_rule` (`id`, `service_id`, `port`, `header_transfor`) VALUES (172,54,8002,'add metadata1 datavalue,edit metadata2 datavalue2');
INSERT INTO `gateway_service_grpc_rule` (`id`, `service_id`, `port`, `header_transfor`) VALUES (173,58,8012,'add meta_name meta_value');

CREATE TABLE IF NOT EXISTS `gateway_service_http_rule` (`id` BIGINT(20) PRIMARY KEY AUTO_INCREMENT NOT NULL COMMENT '自增主键', `service_id` BIGINT(20) NOT NULL COMMENT '服务id', `rule_type` INT DEFAULT 0 NOT NULL COMMENT '匹配类型 0=url前缀url_prefix 1=域名domain', `rule` VARCHAR(255) DEFAULT '' NOT NULL COMMENT 'type=domain表示域名，type=url_prefix时表示url前缀', `need_https` INT DEFAULT 0 NOT NULL COMMENT '支持https 1=支持', `need_strip_url` TINYINT DEFAULT 0 NOT NULL COMMENT '启用strip_uri 1=启用', `need_websocket` TINYINT DEFAULT 0 NOT NULL COMMENT '是否支持websocket 1=支持', `url_rewrite` VARCHAR(5000) DEFAULT '' NOT NULL COMMENT 'url重写功能 格式：^/gatekeeper/test_service(.*) $1 多个逗号间隔', `header_transfor` VARCHAR(5000) DEFAULT '' NOT NULL COMMENT 'header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔') ENGINE=InnoDB DEFAULT CHARSET utf8;
CREATE UNIQUE INDEX `UQE_gateway_service_http_rule_service_id` ON `gateway_service_http_rule` (`service_id`);
INSERT INTO `gateway_service_http_rule` (`id`, `service_id`, `rule_type`, `rule`, `need_https`, `need_strip_url`, `need_websocket`, `url_rewrite`, `header_transfor`) VALUES (165,35,1,'',0,0,0,'','');
INSERT INTO `gateway_service_http_rule` (`id`, `service_id`, `rule_type`, `rule`, `need_https`, `need_strip_url`, `need_websocket`, `url_rewrite`, `header_transfor`) VALUES (168,34,0,'',0,0,0,'','');
INSERT INTO `gateway_service_http_rule` (`id`, `service_id`, `rule_type`, `rule`, `need_https`, `need_strip_url`, `need_websocket`, `url_rewrite`, `header_transfor`) VALUES (170,36,0,'',0,0,0,'','');
INSERT INTO `gateway_service_http_rule` (`id`, `service_id`, `rule_type`, `rule`, `need_https`, `need_strip_url`, `need_websocket`, `url_rewrite`, `header_transfor`) VALUES (171,38,0,'/abc',1,0,1,'^/abc $1','add head1 value1');
INSERT INTO `gateway_service_http_rule` (`id`, `service_id`, `rule_type`, `rule`, `need_https`, `need_strip_url`, `need_websocket`, `url_rewrite`, `header_transfor`) VALUES (172,43,0,'/usr',1,1,0,'^/afsaasf $1,^/afsaasf $1','');
INSERT INTO `gateway_service_http_rule` (`id`, `service_id`, `rule_type`, `rule`, `need_https`, `need_strip_url`, `need_websocket`, `url_rewrite`, `header_transfor`) VALUES (173,44,1,'www.test.com',1,1,1,'','');
INSERT INTO `gateway_service_http_rule` (`id`, `service_id`, `rule_type`, `rule`, `need_https`, `need_strip_url`, `need_websocket`, `url_rewrite`, `header_transfor`) VALUES (174,47,1,'www.test.com',1,1,1,'','');
INSERT INTO `gateway_service_http_rule` (`id`, `service_id`, `rule_type`, `rule`, `need_https`, `need_strip_url`, `need_websocket`, `url_rewrite`, `header_transfor`) VALUES (175,48,1,'www.test.com',1,1,1,'','');
INSERT INTO `gateway_service_http_rule` (`id`, `service_id`, `rule_type`, `rule`, `need_https`, `need_strip_url`, `need_websocket`, `url_rewrite`, `header_transfor`) VALUES (176,49,1,'www.test.com',1,1,1,'','');
INSERT INTO `gateway_service_http_rule` (`id`, `service_id`, `rule_type`, `rule`, `need_https`, `need_strip_url`, `need_websocket`, `url_rewrite`, `header_transfor`) VALUES (177,56,0,'/test_http_service',1,1,1,'^/test_http_service/abb/{.*} /test_http_service/bba/$1','add header_name header_value');
INSERT INTO `gateway_service_http_rule` (`id`, `service_id`, `rule_type`, `rule`, `need_https`, `need_strip_url`, `need_websocket`, `url_rewrite`, `header_transfor`) VALUES (178,59,1,'test.com',0,1,1,'','add headername headervalue');
INSERT INTO `gateway_service_http_rule` (`id`, `service_id`, `rule_type`, `rule`, `need_https`, `need_strip_url`, `need_websocket`, `url_rewrite`, `header_transfor`) VALUES (179,60,0,'/test_strip_uri',0,1,0,'^/aaa/{.*} /bbb/$1','');
INSERT INTO `gateway_service_http_rule` (`id`, `service_id`, `rule_type`, `rule`, `need_https`, `need_strip_url`, `need_websocket`, `url_rewrite`, `header_transfor`) VALUES (180,61,0,'/test_https_server',1,1,0,'','');

CREATE TABLE IF NOT EXISTS `gateway_service_info` (`id` BIGINT(20) PRIMARY KEY AUTO_INCREMENT NOT NULL COMMENT '自增主键', `load_type` TINYINT DEFAULT 0 NOT NULL COMMENT '负载类型 0=http 1=tcp 2=grpc', `service_name` VARCHAR(255) NOT NULL COMMENT '服务名称 6-128 数字字母下划线', `service_desc` VARCHAR(255) DEFAULT '' NULL COMMENT '服务描述', `create_at` DATETIME NULL COMMENT '创建时间', `update_at` DATETIME NULL COMMENT '更新时间', `delete_at` DATETIME NULL COMMENT '删除时间', `is_delete` INT DEFAULT 0 NOT NULL COMMENT '是否删除 1=删除') ENGINE=InnoDB DEFAULT CHARSET utf8;
CREATE UNIQUE INDEX `UQE_gateway_service_info_service_name` ON `gateway_service_info` (`service_name`);
INSERT INTO `gateway_service_info` (`id`, `load_type`, `service_name`, `service_desc`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (34,0,'websocket_test','websocket_test','2023-07-16 16:27:41','2023-07-16 16:27:41','2023-07-16 16:27:41',1);
INSERT INTO `gateway_service_info` (`id`, `load_type`, `service_name`, `service_desc`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (35,1,'test_grpc','test_grpc','2023-07-16 16:27:41','2023-07-16 16:27:41','2023-07-16 16:27:41',1);
INSERT INTO `gateway_service_info` (`id`, `load_type`, `service_name`, `service_desc`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (36,2,'test_httpe','test_httpe','2023-07-16 16:27:41','2023-07-16 16:27:41','2023-07-16 16:27:41',1);
INSERT INTO `gateway_service_info` (`id`, `load_type`, `service_name`, `service_desc`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (38,0,'service_name','11111','2023-07-16 16:27:41','2023-07-16 16:27:41','2023-07-16 16:27:41',1);
INSERT INTO `gateway_service_info` (`id`, `load_type`, `service_name`, `service_desc`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (41,0,'service_name_tcp','11111','2023-07-16 16:27:41','2023-07-16 16:27:41','2023-07-16 16:27:41',1);
INSERT INTO `gateway_service_info` (`id`, `load_type`, `service_name`, `service_desc`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (42,0,'service_name_tcp2','11111','2023-07-16 16:27:41','2023-07-16 16:27:41','2023-07-16 16:27:41',1);
INSERT INTO `gateway_service_info` (`id`, `load_type`, `service_name`, `service_desc`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (43,1,'service_name_tcp4','service_name_tcp4','2023-07-16 16:27:41','2023-07-16 16:27:41','2023-07-16 16:27:41',1);
INSERT INTO `gateway_service_info` (`id`, `load_type`, `service_name`, `service_desc`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (44,0,'websocket_service','websocket_service','2023-07-16 16:27:41','2023-07-16 16:27:41','2023-07-16 16:27:41',1);
INSERT INTO `gateway_service_info` (`id`, `load_type`, `service_name`, `service_desc`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (45,1,'tcp_service','tcp_desc','2023-07-16 16:27:41','2023-07-16 16:27:41','2023-07-16 16:27:41',1);
INSERT INTO `gateway_service_info` (`id`, `load_type`, `service_name`, `service_desc`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (46,1,'grpc_service','grpc_desc','2023-07-16 16:27:41','2023-07-16 16:27:41','2023-07-16 16:27:41',1);
INSERT INTO `gateway_service_info` (`id`, `load_type`, `service_name`, `service_desc`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (47,0,'testsefsafs','werrqrr','2023-07-16 16:27:41','2023-07-16 16:27:41','2023-07-16 16:27:41',1);
INSERT INTO `gateway_service_info` (`id`, `load_type`, `service_name`, `service_desc`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (48,0,'testsefsafs1','werrqrr','2023-07-16 16:27:41','2023-07-16 16:27:41','2023-07-16 16:27:41',1);
INSERT INTO `gateway_service_info` (`id`, `load_type`, `service_name`, `service_desc`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (49,0,'testsefsafs1222','werrqrr','2023-07-16 16:27:41','2023-07-16 16:27:41','2023-07-16 16:27:41',1);
INSERT INTO `gateway_service_info` (`id`, `load_type`, `service_name`, `service_desc`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (50,2,'grpc_service_name','grpc_service_desc','2023-07-16 16:27:41','2023-07-16 16:27:41','2023-07-16 16:27:41',1);
INSERT INTO `gateway_service_info` (`id`, `load_type`, `service_name`, `service_desc`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (51,2,'gresafsf','wesfsf','2023-07-16 16:27:41','2023-07-16 16:27:41','2023-07-16 16:27:41',1);
INSERT INTO `gateway_service_info` (`id`, `load_type`, `service_name`, `service_desc`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (52,2,'gresafsf11','wesfsf','2023-07-16 16:27:41','2023-07-16 16:27:41','2023-07-16 16:27:41',1);
INSERT INTO `gateway_service_info` (`id`, `load_type`, `service_name`, `service_desc`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (53,2,'tewrqrw111','123313','2023-07-16 16:27:41','2023-07-16 16:27:41','2023-07-16 16:27:41',1);
INSERT INTO `gateway_service_info` (`id`, `load_type`, `service_name`, `service_desc`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (54,2,'test_grpc_service1','test_grpc_service1','2023-07-16 16:27:41','2023-07-16 16:27:41','2023-07-16 16:27:41',1);
INSERT INTO `gateway_service_info` (`id`, `load_type`, `service_name`, `service_desc`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (55,1,'test_tcp_service1','redis服务代理','2023-07-16 16:27:41','2023-07-16 16:27:41','2023-07-16 16:27:41',1);
INSERT INTO `gateway_service_info` (`id`, `load_type`, `service_name`, `service_desc`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (56,0,'test_http_service','测试HTTP代理','2023-07-16 16:27:41','2023-07-16 16:27:41',NULL,0);
INSERT INTO `gateway_service_info` (`id`, `load_type`, `service_name`, `service_desc`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (57,1,'test_tcp_service','测试TCP代理','2023-07-16 16:27:41','2023-07-16 16:27:41',NULL,0);
INSERT INTO `gateway_service_info` (`id`, `load_type`, `service_name`, `service_desc`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (58,2,'test_grpc_service','测试GRPC服务','2023-07-16 16:27:41','2023-07-16 16:27:41',NULL,0);
INSERT INTO `gateway_service_info` (`id`, `load_type`, `service_name`, `service_desc`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (59,0,'test.com:8080','测试域名接入','2023-07-16 16:27:41','2023-07-16 16:27:41',NULL,0);
INSERT INTO `gateway_service_info` (`id`, `load_type`, `service_name`, `service_desc`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (60,0,'test_strip_uri','测试路径接入','2023-07-16 16:27:41','2023-07-16 16:27:41',NULL,0);
INSERT INTO `gateway_service_info` (`id`, `load_type`, `service_name`, `service_desc`, `create_at`, `update_at`, `delete_at`, `is_delete`) VALUES (61,0,'test_https_server','测试https服务','2023-07-16 16:27:41','2023-07-16 16:27:41',NULL,0);

CREATE TABLE IF NOT EXISTS `gateway_service_load_balance` (`id` BIGINT(20) PRIMARY KEY AUTO_INCREMENT NOT NULL COMMENT '自增主键', `service_id` BIGINT(20) NOT NULL COMMENT '服务id', `check_method` INT DEFAULT 0 NOT NULL COMMENT '检查方法 0=tcpchk,检测端口是否握手成功', `check_timeout` INT DEFAULT 0 NOT NULL COMMENT 'check超时时间,单位s', `check_interval` INT DEFAULT 0 NOT NULL COMMENT '检查间隔, 单位s', `round_type` INT DEFAULT 0 NOT NULL COMMENT '轮询方式 0=random 1=round-robin 2=weight_round-robin 3=ip_hash', `ip_list` VARCHAR(2000) DEFAULT '' NOT NULL COMMENT 'ip列表', `weight_list` VARCHAR(2000) DEFAULT '' NOT NULL COMMENT '权重列表', `forbid_list` VARCHAR(2000) DEFAULT '' NOT NULL COMMENT '禁用ip列表', `upstream_connect_timeout` INT DEFAULT 0 NOT NULL COMMENT '建立连接超时, 单位s', `upstream_header_timeout` INT DEFAULT 0 NOT NULL COMMENT '获取header超时, 单位s', `upstream_idle_timeout` INT DEFAULT 0 NOT NULL COMMENT '链接最大空闲时间, 单位s', `upstream_max_idle` INT DEFAULT 0 NOT NULL COMMENT '最大空闲链接数') ENGINE=InnoDB DEFAULT CHARSET utf8;
CREATE UNIQUE INDEX `UQE_gateway_service_load_balance_service_id` ON `gateway_service_load_balance` (`service_id`);
INSERT INTO `gateway_service_load_balance` (`id`, `service_id`, `check_method`, `check_timeout`, `check_interval`, `round_type`, `ip_list`, `weight_list`, `forbid_list`, `upstream_connect_timeout`, `upstream_header_timeout`, `upstream_idle_timeout`, `upstream_max_idle`) VALUES (162,35,0,2000,5000,2,'127.0.0.1:50051','100','',10000,0,0,0);
INSERT INTO `gateway_service_load_balance` (`id`, `service_id`, `check_method`, `check_timeout`, `check_interval`, `round_type`, `ip_list`, `weight_list`, `forbid_list`, `upstream_connect_timeout`, `upstream_header_timeout`, `upstream_idle_timeout`, `upstream_max_idle`) VALUES (165,34,0,2000,5000,2,'100.90.164.31:8072,100.90.163.51:8072,100.90.163.52:8072,100.90.165.32:8072','50,50,50,80','',20000,20000,10000,100);
INSERT INTO `gateway_service_load_balance` (`id`, `service_id`, `check_method`, `check_timeout`, `check_interval`, `round_type`, `ip_list`, `weight_list`, `forbid_list`, `upstream_connect_timeout`, `upstream_header_timeout`, `upstream_idle_timeout`, `upstream_max_idle`) VALUES (167,36,0,2000,5000,2,'100.90.164.31:8072,100.90.163.51:8072,100.90.163.52:8072,100.90.165.32:8072','50,50,50,80','100.90.164.31:8072,100.90.163.51:8072',10000,10000,10000,100);
INSERT INTO `gateway_service_load_balance` (`id`, `service_id`, `check_method`, `check_timeout`, `check_interval`, `round_type`, `ip_list`, `weight_list`, `forbid_list`, `upstream_connect_timeout`, `upstream_header_timeout`, `upstream_idle_timeout`, `upstream_max_idle`) VALUES (168,38,0,0,0,1,'111:111,22:111','11,11','111',1111,111,222,333);
INSERT INTO `gateway_service_load_balance` (`id`, `service_id`, `check_method`, `check_timeout`, `check_interval`, `round_type`, `ip_list`, `weight_list`, `forbid_list`, `upstream_connect_timeout`, `upstream_header_timeout`, `upstream_idle_timeout`, `upstream_max_idle`) VALUES (169,41,0,0,0,1,'111:111,22:111','11,11','111',0,0,0,0);
INSERT INTO `gateway_service_load_balance` (`id`, `service_id`, `check_method`, `check_timeout`, `check_interval`, `round_type`, `ip_list`, `weight_list`, `forbid_list`, `upstream_connect_timeout`, `upstream_header_timeout`, `upstream_idle_timeout`, `upstream_max_idle`) VALUES (170,42,0,0,0,1,'111:111,22:111','11,11','111',0,0,0,0);
INSERT INTO `gateway_service_load_balance` (`id`, `service_id`, `check_method`, `check_timeout`, `check_interval`, `round_type`, `ip_list`, `weight_list`, `forbid_list`, `upstream_connect_timeout`, `upstream_header_timeout`, `upstream_idle_timeout`, `upstream_max_idle`) VALUES (171,43,0,2,5,1,'111:111,22:111','11,11','',1111,2222,333,444);
INSERT INTO `gateway_service_load_balance` (`id`, `service_id`, `check_method`, `check_timeout`, `check_interval`, `round_type`, `ip_list`, `weight_list`, `forbid_list`, `upstream_connect_timeout`, `upstream_header_timeout`, `upstream_idle_timeout`, `upstream_max_idle`) VALUES (172,44,0,2,5,2,'127.0.0.1:8076','50','',0,0,0,0);
INSERT INTO `gateway_service_load_balance` (`id`, `service_id`, `check_method`, `check_timeout`, `check_interval`, `round_type`, `ip_list`, `weight_list`, `forbid_list`, `upstream_connect_timeout`, `upstream_header_timeout`, `upstream_idle_timeout`, `upstream_max_idle`) VALUES (173,45,0,2,5,2,'127.0.0.1:88','50','',0,0,0,0);
INSERT INTO `gateway_service_load_balance` (`id`, `service_id`, `check_method`, `check_timeout`, `check_interval`, `round_type`, `ip_list`, `weight_list`, `forbid_list`, `upstream_connect_timeout`, `upstream_header_timeout`, `upstream_idle_timeout`, `upstream_max_idle`) VALUES (174,46,0,2,5,2,'127.0.0.1:8002','50','',0,0,0,0);
INSERT INTO `gateway_service_load_balance` (`id`, `service_id`, `check_method`, `check_timeout`, `check_interval`, `round_type`, `ip_list`, `weight_list`, `forbid_list`, `upstream_connect_timeout`, `upstream_header_timeout`, `upstream_idle_timeout`, `upstream_max_idle`) VALUES (175,47,0,2,5,2,'12777:11','11','',0,0,0,0);
INSERT INTO `gateway_service_load_balance` (`id`, `service_id`, `check_method`, `check_timeout`, `check_interval`, `round_type`, `ip_list`, `weight_list`, `forbid_list`, `upstream_connect_timeout`, `upstream_header_timeout`, `upstream_idle_timeout`, `upstream_max_idle`) VALUES (176,48,0,2,5,2,'12777:11','11','',0,0,0,0);
INSERT INTO `gateway_service_load_balance` (`id`, `service_id`, `check_method`, `check_timeout`, `check_interval`, `round_type`, `ip_list`, `weight_list`, `forbid_list`, `upstream_connect_timeout`, `upstream_header_timeout`, `upstream_idle_timeout`, `upstream_max_idle`) VALUES (177,49,0,2,5,2,'12777:11','11','',0,0,0,0);
INSERT INTO `gateway_service_load_balance` (`id`, `service_id`, `check_method`, `check_timeout`, `check_interval`, `round_type`, `ip_list`, `weight_list`, `forbid_list`, `upstream_connect_timeout`, `upstream_header_timeout`, `upstream_idle_timeout`, `upstream_max_idle`) VALUES (178,50,0,2,5,2,'127.0.0.1:8001','50','',0,0,0,0);
INSERT INTO `gateway_service_load_balance` (`id`, `service_id`, `check_method`, `check_timeout`, `check_interval`, `round_type`, `ip_list`, `weight_list`, `forbid_list`, `upstream_connect_timeout`, `upstream_header_timeout`, `upstream_idle_timeout`, `upstream_max_idle`) VALUES (179,51,0,2,5,2,'1212:11','50','',0,0,0,0);
INSERT INTO `gateway_service_load_balance` (`id`, `service_id`, `check_method`, `check_timeout`, `check_interval`, `round_type`, `ip_list`, `weight_list`, `forbid_list`, `upstream_connect_timeout`, `upstream_header_timeout`, `upstream_idle_timeout`, `upstream_max_idle`) VALUES (180,52,0,2,5,2,'1212:11','50','',0,0,0,0);
INSERT INTO `gateway_service_load_balance` (`id`, `service_id`, `check_method`, `check_timeout`, `check_interval`, `round_type`, `ip_list`, `weight_list`, `forbid_list`, `upstream_connect_timeout`, `upstream_header_timeout`, `upstream_idle_timeout`, `upstream_max_idle`) VALUES (181,53,0,2,5,2,'1111:11','111','',0,0,0,0);
INSERT INTO `gateway_service_load_balance` (`id`, `service_id`, `check_method`, `check_timeout`, `check_interval`, `round_type`, `ip_list`, `weight_list`, `forbid_list`, `upstream_connect_timeout`, `upstream_header_timeout`, `upstream_idle_timeout`, `upstream_max_idle`) VALUES (182,54,0,2,5,1,'127.0.0.1:80','50','',0,0,0,0);
INSERT INTO `gateway_service_load_balance` (`id`, `service_id`, `check_method`, `check_timeout`, `check_interval`, `round_type`, `ip_list`, `weight_list`, `forbid_list`, `upstream_connect_timeout`, `upstream_header_timeout`, `upstream_idle_timeout`, `upstream_max_idle`) VALUES (183,55,0,2,5,3,'127.0.0.1:81','50','',0,0,0,0);
INSERT INTO `gateway_service_load_balance` (`id`, `service_id`, `check_method`, `check_timeout`, `check_interval`, `round_type`, `ip_list`, `weight_list`, `forbid_list`, `upstream_connect_timeout`, `upstream_header_timeout`, `upstream_idle_timeout`, `upstream_max_idle`) VALUES (184,56,0,2,5,2,'127.0.0.1:2003,127.0.0.1:2004','50,50','',0,0,0,0);
INSERT INTO `gateway_service_load_balance` (`id`, `service_id`, `check_method`, `check_timeout`, `check_interval`, `round_type`, `ip_list`, `weight_list`, `forbid_list`, `upstream_connect_timeout`, `upstream_header_timeout`, `upstream_idle_timeout`, `upstream_max_idle`) VALUES (185,57,0,2,5,2,'127.0.0.1:6379','50','',0,0,0,0);
INSERT INTO `gateway_service_load_balance` (`id`, `service_id`, `check_method`, `check_timeout`, `check_interval`, `round_type`, `ip_list`, `weight_list`, `forbid_list`, `upstream_connect_timeout`, `upstream_header_timeout`, `upstream_idle_timeout`, `upstream_max_idle`) VALUES (186,58,0,2,5,2,'127.0.0.1:50055','50','',0,0,0,0);
INSERT INTO `gateway_service_load_balance` (`id`, `service_id`, `check_method`, `check_timeout`, `check_interval`, `round_type`, `ip_list`, `weight_list`, `forbid_list`, `upstream_connect_timeout`, `upstream_header_timeout`, `upstream_idle_timeout`, `upstream_max_idle`) VALUES (187,59,0,2,5,2,'127.0.0.1:2003,127.0.0.1:2004','50,50','',0,0,0,0);
INSERT INTO `gateway_service_load_balance` (`id`, `service_id`, `check_method`, `check_timeout`, `check_interval`, `round_type`, `ip_list`, `weight_list`, `forbid_list`, `upstream_connect_timeout`, `upstream_header_timeout`, `upstream_idle_timeout`, `upstream_max_idle`) VALUES (188,60,0,2,5,2,'127.0.0.1:2003,127.0.0.1:2004','50,50','',0,0,0,0);
INSERT INTO `gateway_service_load_balance` (`id`, `service_id`, `check_method`, `check_timeout`, `check_interval`, `round_type`, `ip_list`, `weight_list`, `forbid_list`, `upstream_connect_timeout`, `upstream_header_timeout`, `upstream_idle_timeout`, `upstream_max_idle`) VALUES (189,61,0,2,5,2,'127.0.0.1:3003,127.0.0.1:3004','50,50','',0,0,0,0);

CREATE TABLE IF NOT EXISTS `gateway_service_tcp_rule` (`id` BIGINT(20) PRIMARY KEY AUTO_INCREMENT NOT NULL COMMENT '自增主键', `service_id` BIGINT(20) DEFAULT 0 NOT NULL, `port` INT NOT NULL COMMENT '端口') ENGINE=InnoDB DEFAULT CHARSET utf8;
CREATE UNIQUE INDEX `UQE_gateway_service_tcp_rule_service_id` ON `gateway_service_tcp_rule` (`service_id`);
INSERT INTO `gateway_service_tcp_rule` (`id`, `service_id`, `port`) VALUES (171,41,8002);
INSERT INTO `gateway_service_tcp_rule` (`id`, `service_id`, `port`) VALUES (172,42,8003);
INSERT INTO `gateway_service_tcp_rule` (`id`, `service_id`, `port`) VALUES (173,43,8004);
INSERT INTO `gateway_service_tcp_rule` (`id`, `service_id`, `port`) VALUES (174,38,8004);
INSERT INTO `gateway_service_tcp_rule` (`id`, `service_id`, `port`) VALUES (175,45,8001);
INSERT INTO `gateway_service_tcp_rule` (`id`, `service_id`, `port`) VALUES (176,46,8005);
INSERT INTO `gateway_service_tcp_rule` (`id`, `service_id`, `port`) VALUES (177,50,8006);
INSERT INTO `gateway_service_tcp_rule` (`id`, `service_id`, `port`) VALUES (178,51,8007);
INSERT INTO `gateway_service_tcp_rule` (`id`, `service_id`, `port`) VALUES (179,52,8008);
INSERT INTO `gateway_service_tcp_rule` (`id`, `service_id`, `port`) VALUES (180,55,8010);
INSERT INTO `gateway_service_tcp_rule` (`id`, `service_id`, `port`) VALUES (181,57,8011);