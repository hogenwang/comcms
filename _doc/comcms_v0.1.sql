/*
Navicat MySQL Data Transfer

Source Server         : mysql
Source Server Version : 50544
Source Host           : localhost:3306
Source Database       : comcms

Target Server Type    : MYSQL
Target Server Version : 50544
File Encoding         : 65001

Date: 2015-12-24 23:40:28
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(20) DEFAULT NULL,
  `pass_word` varchar(255) DEFAULT NULL,
  `nick_name` varchar(255) DEFAULT NULL,
  `role_id` bigint(20) DEFAULT NULL,
  `last_login_time` datetime DEFAULT NULL,
  `last_login_i_p` varchar(255) DEFAULT NULL,
  `this_login_time` datetime DEFAULT NULL,
  `this_login_i_p` varchar(255) DEFAULT NULL,
  `is_lock` tinyint(1) DEFAULT NULL,
  `editor_id` bigint(20) DEFAULT NULL,
  `notes` varchar(255) DEFAULT NULL,
  `login_count` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UQE_admin_user_name` (`user_name`),
  KEY `IDX_admin_login_count` (`login_count`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of admin
-- ----------------------------
INSERT INTO `admin` VALUES ('1', 'admin', '21232f297a57a5a743894a0e4a801fc3', 'admin', '1', '2015-12-24 22:46:54', '127.0.0.1', '2015-12-24 23:37:09', '127.0.0.1', '0', '0', null, '27');

-- ----------------------------
-- Table structure for admin_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_role`;
CREATE TABLE `admin_role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `role_name` varchar(255) DEFAULT NULL,
  `role_description` varchar(255) DEFAULT NULL,
  `is_super_admin` tinyint(1) DEFAULT NULL,
  `stars` bigint(20) DEFAULT NULL,
  `color` varchar(255) DEFAULT NULL,
  `not_allow_del` tinyint(1) DEFAULT NULL,
  `rank` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UQE_admin_role_role_name` (`role_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of admin_role
-- ----------------------------

-- ----------------------------
-- Table structure for ads
-- ----------------------------
DROP TABLE IF EXISTS `ads`;
CREATE TABLE `ads` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `kid` bigint(20) DEFAULT NULL,
  `tid` bigint(20) DEFAULT NULL,
  `title` varchar(250) DEFAULT NULL,
  `rank` bigint(20) DEFAULT NULL,
  `description` varchar(250) DEFAULT NULL,
  `content` text,
  `is_hide` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `IDX_ads_rank` (`rank`),
  KEY `IDX_ads_kid` (`kid`),
  KEY `IDX_ads_tid` (`tid`),
  KEY `IDX_ads_title` (`title`),
  KEY `IDX_ads_is_hide` (`is_hide`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `kid` bigint(20) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `content` text,
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  `author_id` bigint(20) DEFAULT NULL,
  `origin` varchar(250) DEFAULT NULL,
  `origin_url` varchar(250) DEFAULT NULL,
  `rank` bigint(20) DEFAULT NULL,
  `keyword` varchar(250) DEFAULT NULL,
  `description` varchar(250) DEFAULT NULL,
  `link_url` varchar(250) DEFAULT NULL,
  `title_color` varchar(20) DEFAULT NULL,
  `pic` varchar(250) DEFAULT NULL,
  `tag` varchar(250) DEFAULT NULL,
  `template_file` varchar(250) DEFAULT NULL,
  `file_name` varchar(250) DEFAULT NULL,
  `views` bigint(20) DEFAULT NULL,
  `is_pass` bigint(20) DEFAULT NULL,
  `is_recommend` bigint(20) DEFAULT NULL,
  `is_top` bigint(20) DEFAULT NULL,
  `is_best` bigint(20) DEFAULT NULL,
  `is_new` bigint(20) DEFAULT NULL,
  `is_del` bigint(20) DEFAULT NULL,
  `is_member` bigint(20) DEFAULT NULL,
  `is_hide` bigint(20) DEFAULT NULL,
  `comment_count` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `IDX_article_title` (`title`),
  KEY `IDX_article_origin` (`origin`),
  KEY `IDX_article_rank` (`rank`),
  KEY `IDX_article_views` (`views`),
  KEY `IDX_article_is_del` (`is_del`),
  KEY `IDX_article_updated` (`updated`),
  KEY `IDX_article_is_best` (`is_best`),
  KEY `IDX_article_is_member` (`is_member`),
  KEY `IDX_article_kid` (`kid`),
  KEY `IDX_article_created` (`created`),
  KEY `IDX_article_origin_url` (`origin_url`),
  KEY `IDX_article_pic` (`pic`),
  KEY `IDX_article_tag` (`tag`),
  KEY `IDX_article_file_name` (`file_name`),
  KEY `IDX_article_is_pass` (`is_pass`),
  KEY `IDX_article_is_top` (`is_top`),
  KEY `IDX_article_comment_count` (`comment_count`),
  KEY `IDX_article_is_hide` (`is_hide`),
  KEY `IDX_article_is_new` (`is_new`),
  KEY `IDX_article_is_recommend` (`is_recommend`),
  KEY `IDX_article_author_id` (`author_id`),
  KEY `IDX_article_template_file` (`template_file`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `pid` bigint(20) DEFAULT NULL,
  `ctype` bigint(20) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `content` text,
  `page_title` varchar(255) DEFAULT NULL,
  `rank` bigint(20) DEFAULT NULL,
  `level` bigint(20) DEFAULT NULL,
  `keyword` varchar(250) DEFAULT NULL,
  `description` varchar(250) DEFAULT NULL,
  `link_url` varchar(250) DEFAULT NULL,
  `title_color` varchar(20) DEFAULT NULL,
  `template_file` varchar(250) DEFAULT NULL,
  `detail_template_file` varchar(250) DEFAULT NULL,
  `is_list` bigint(20) DEFAULT NULL,
  `page_size` bigint(20) DEFAULT NULL,
  `is_lock` bigint(20) DEFAULT NULL,
  `is_del` bigint(20) DEFAULT NULL,
  `is_hide` bigint(20) DEFAULT NULL,
  `is_disabled` bigint(20) DEFAULT NULL,
  `is_comment` bigint(20) DEFAULT NULL,
  `is_header_nav` bigint(20) DEFAULT NULL,
  `is_footer_nav` bigint(20) DEFAULT NULL,
  `counts` bigint(20) DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  `catalog_id` bigint(20) DEFAULT NULL,
  `pic` varchar(250) DEFAULT NULL,
  `ads_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `IDX_category_created` (`created`),
  KEY `IDX_category_pic` (`pic`),
  KEY `IDX_category_is_footer_nav` (`is_footer_nav`),
  KEY `IDX_category_counts` (`counts`),
  KEY `IDX_category_page_title` (`page_title`),
  KEY `IDX_category_rank` (`rank`),
  KEY `IDX_category_is_del` (`is_del`),
  KEY `IDX_category_title` (`title`),
  KEY `IDX_category_level` (`level`),
  KEY `IDX_category_is_lock` (`is_lock`),
  KEY `IDX_category_is_header_nav` (`is_header_nav`),
  KEY `IDX_category_ads_id` (`ads_id`),
  KEY `IDX_category_is_comment` (`is_comment`),
  KEY `IDX_category_updated` (`updated`),
  KEY `IDX_category_is_list` (`is_list`),
  KEY `IDX_category_page_size` (`page_size`),
  KEY `IDX_category_is_hide` (`is_hide`),
  KEY `IDX_category_catalog_id` (`catalog_id`),
  KEY `IDX_category_is_disabled` (`is_disabled`),
  KEY `IDX_category_pid` (`pid`),
  KEY `IDX_category_ctype` (`ctype`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for config
-- ----------------------------
DROP TABLE IF EXISTS `config`;
CREATE TABLE `config` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `site_name` varchar(100) DEFAULT NULL,
  `site_u_r_l` varchar(100) DEFAULT NULL,
  `site_logo` varchar(200) DEFAULT NULL,
  `i_c_p` varchar(200) DEFAULT NULL,
  `site_email` varchar(50) DEFAULT NULL,
  `copyright` text,
  `is_close_site` bigint(20) DEFAULT NULL,
  `close_reason` text,
  `keyword` varchar(250) DEFAULT NULL,
  `description` varchar(250) DEFAULT NULL,
  `site_title` varchar(250) DEFAULT NULL,
  `search_min_time` bigint(20) DEFAULT NULL,
  `online_q_q` varchar(250) DEFAULT NULL,
  `online_skype` varchar(250) DEFAULT NULL,
  `online_wang_wang` varchar(250) DEFAULT NULL,
  `skin` varchar(50) DEFAULT NULL,
  `last_update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of config
-- ----------------------------
INSERT INTO `config` VALUES ('1', 'COMCMS', 'http://www.comcms.com', null, null, 'master@comcms.com', '<p style=\"text-align: center;\"><span style=\"color: rgb(255, 0, 0);\">版权所有 2010-2016 COMCMS</span></p>', null, null, 'CMS', '幻之角CMS', null, null, '466364748', null, null, 'default', '2015-12-24 23:19:33');

-- ----------------------------
-- Table structure for guestbook
-- ----------------------------
DROP TABLE IF EXISTS `guestbook`;
CREATE TABLE `guestbook` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `kid` bigint(20) DEFAULT NULL,
  `title` varchar(250) DEFAULT NULL,
  `content` text,
  `uid` bigint(20) DEFAULT NULL,
  `user_name` varchar(50) DEFAULT NULL,
  `user_img` varchar(250) DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  `is_verify` bigint(20) DEFAULT NULL,
  `is_read` bigint(20) DEFAULT NULL,
  `is_del` bigint(20) DEFAULT NULL,
  `i_p` varchar(20) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `tel` varchar(200) DEFAULT NULL,
  `q_q` varchar(20) DEFAULT NULL,
  `skype` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `IDX_guestbook_kid` (`kid`),
  KEY `IDX_guestbook_created` (`created`),
  KEY `IDX_guestbook_is_verify` (`is_verify`),
  KEY `IDX_guestbook_is_del` (`is_del`),
  KEY `IDX_guestbook_email` (`email`),
  KEY `IDX_guestbook_user_name` (`user_name`),
  KEY `IDX_guestbook_q_q` (`q_q`),
  KEY `IDX_guestbook_title` (`title`),
  KEY `IDX_guestbook_uid` (`uid`),
  KEY `IDX_guestbook_skype` (`skype`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of guestbook
-- ----------------------------

-- ----------------------------
-- Table structure for link
-- ----------------------------
DROP TABLE IF EXISTS `link`;
CREATE TABLE `link` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `kid` bigint(20) DEFAULT NULL,
  `title` varchar(250) DEFAULT NULL,
  `rank` bigint(20) DEFAULT NULL,
  `url` varchar(250) DEFAULT NULL,
  `description` varchar(250) DEFAULT NULL,
  `is_hide` bigint(20) DEFAULT NULL,
  `logo` varchar(250) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `IDX_link_kid` (`kid`),
  KEY `IDX_link_title` (`title`),
  KEY `IDX_link_rank` (`rank`),
  KEY `IDX_link_is_hide` (`is_hide`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of link
-- ----------------------------
INSERT INTO `link` VALUES ('1', '0', 'COMCMS', '999', 'http://www.comcms.com', '官网链接', '0', '');

-- ----------------------------
-- Table structure for link_kind
-- ----------------------------
DROP TABLE IF EXISTS `link_kind`;
CREATE TABLE `link_kind` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL,
  `rank` bigint(20) DEFAULT NULL,
  `keyword` varchar(250) DEFAULT NULL,
  `description` varchar(250) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `IDX_link_kind_rank` (`rank`),
  KEY `IDX_link_kind_title` (`title`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of link_kind
-- ----------------------------
