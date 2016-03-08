/*
Navicat MySQL Data Transfer

Source Server         : magicid
Source Server Version : 50624
Source Host           : localhost:3306
Source Database       : magicid_db

Target Server Type    : MYSQL
Target Server Version : 50624
File Encoding         : 65001

Date: 2016-03-08 23:19:45
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `article`
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` tinytext NOT NULL,
  `content` text NOT NULL,
  `author` tinyint(4) NOT NULL,
  `createdate` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of article
-- ----------------------------
INSERT INTO `article` VALUES ('1', '   测试标题', '&lt;div class=&#34;card_try&#34; style=&#34;font-family: SimSun; line-height: 18px; color: rgb(51, 51, 51); margin-bottom: 20px; background-color: rgb(249, 249, 249);&#34;&gt;&lt;p style=&#34;border: 0px; border-collapse: collapse; border-spacing: 0px; list-style: none; margin: 0px; padding: 0px 0px 5px; font-family: &#39;Microsoft YaHei&#39;; font-size: 14px; line-height: 19px;&#34;&gt;为您提供优质的海外文献检索服务，从此省去大海捞针的烦恼，立即体验：&lt;/p&gt;&lt;ol class=&#34;card_try_list1&#34; style=&#34;border: 0px; border-collapse: collapse; border-spacing: 0px; list-style: none; margin: 0px; padding: 0px; font-family: Arial;&#34;&gt;&lt;li style=&#34;border: 0px; border-collapse: collapse; border-spacing: 0px; list-style: none; margin: 0px; padding: 0px;&#34;&gt;&lt;a href=&#34;http://cn.bing.com/academic/search?q=Speaker-independent+phone+recognition+using+hidden+Markov+models&amp;amp;form=QBRE&#34; style=&#34;color: rgb(96, 0, 144); text-decoration: none;&#34;&gt;Speaker-independent phone recognition using hidden Markov models&lt;/a&gt;&lt;/li&gt;&lt;li style=&#34;border: 0px; border-collapse: collapse; border-spacing: 0px; list-style: none; margin: 0px; padding: 0px;&#34;&gt;&lt;a href=&#34;http://cn.bing.com/academic/search?q=Mining+frequent+patterns+without+candidate+generation&amp;amp;qs=n&amp;amp;form=QBRE&#34; style=&#34;color: rgb(96, 0, 144); text-decoration: none;&#34;&gt;Mining frequent patterns without candidate generation&lt;/a&gt;&lt;/li&gt;&lt;/ol&gt;&lt;/div&gt;&lt;div class=&#34;card_try&#34; style=&#34;font-family: SimSun; line-height: 18px; color: rgb(51, 51, 51); margin-bottom: 20px; background-color: rgb(249, 249, 249);&#34;&gt;&lt;p style=&#34;border: 0px; border-collapse: collapse; border-spacing: 0px; list-style: none; margin: 0px; padding: 0px 0px 5px; font-family: &#39;Microsoft YaHei&#39;; font-size: 14px; line-height: 19px;&#34;&gt;好奇某个学术领域？研究范围、期刊会议、重要学者......统统都到碗里来：&lt;/p&gt;&lt;ol class=&#34;card_try_list2&#34; style=&#34;border: 0px; border-collapse: collapse; border-spacing: 0px; list-style: none; margin: 0px; padding: 0px; font-family: Arial;&#34;&gt;&lt;li style=&#34;border: 0px; border-collapse: collapse; border-spacing: 0px; list-style: none; margin: 0px; padding: 0px; display: inline;&#34;&gt;&lt;a href=&#34;http://cn.bing.com/academic/search?q=Deep+Learning&amp;amp;form=QBRE&#34; style=&#34;color: rgb(96, 0, 144); text-decoration: none;&#34;&gt;Deep Learning&lt;/a&gt;&lt;/li&gt;&lt;i&gt;&amp;nbsp;·&amp;nbsp;&lt;/i&gt;&lt;li style=&#34;border: 0px; border-collapse: collapse; border-spacing: 0px; list-style: none; margin: 0px; padding: 0px; display: inline;&#34;&gt;&lt;a href=&#34;http://cn.bing.com/academic/search?q=Artificial+Intelligence&amp;amp;qs=n&amp;amp;form=QBRE&#34; style=&#34;color: rgb(96, 0, 144); text-decoration: none;&#34;&gt;Artificial Intelligence&lt;/a&gt;&lt;/li&gt;&lt;/ol&gt;&lt;/div&gt;&lt;div class=&#34;card_try&#34; style=&#34;font-family: SimSun; line-height: 18px; color: rgb(51, 51, 51); background-color: rgb(249, 249, 249);&#34;&gt;&lt;p style=&#34;border: 0px; border-collapse: collapse; border-spacing: 0px; list-style: none; margin: 0px; padding: 0px 0px 5px; font-family: &#39;Microsoft YaHei&#39;; font-size: 14px; line-height: 19px;&#34;&gt;想了解某位学术大咖在哪些领域有所建树？动动手指搜一搜：&lt;/p&gt;&lt;ol class=&#34;card_try_list2&#34; style=&#34;border: 0px; border-collapse: collapse; border-spacing: 0px; list-style: none; margin: 0px; padding: 0px; font-family: Arial;&#34;&gt;&lt;li style=&#34;border: 0px; border-collapse: collapse; border-spacing: 0px; list-style: none; margin: 0px; padding: 0px; display: inline;&#34;&gt;&lt;a href=&#34;http://cn.bing.com/academic/search?q=Michael+Stonebraker&amp;amp;form=QBRE&#34; style=&#34;color: rgb(96, 0, 144); text-decoration: none;&#34;&gt;Michael Stonebraker&lt;/a&gt;&lt;/li&gt;&lt;i&gt;&amp;nbsp;·&amp;nbsp;&lt;/i&gt;&lt;li style=&#34;border: 0px; border-collapse: collapse; border-spacing: 0px; list-style: none; margin: 0px; padding: 0px; display: inline;&#34;&gt;&lt;a href=&#34;http://cn.bing.com/academic/search?q=Hsiaowuen+Hon&amp;amp;qs=n&amp;amp;form=QBRE&#34; style=&#34;color: rgb(96, 0, 144); text-decoration: none;&#34;&gt;Hsiaowuen Hon&lt;/a&gt;&lt;/li&gt;&lt;/ol&gt;&lt;/div&gt;', '7', '2016-01-02 17:48:55');

-- ----------------------------
-- Table structure for `block_item`
-- ----------------------------
DROP TABLE IF EXISTS `block_item`;
CREATE TABLE `block_item` (
  `id` int(11) NOT NULL DEFAULT '0',
  `name` text,
  `url` text,
  `owner` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of block_item
-- ----------------------------

-- ----------------------------
-- Table structure for `catalog`
-- ----------------------------
DROP TABLE IF EXISTS `catalog`;
CREATE TABLE `catalog` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` text NOT NULL,
  `creater` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of catalog
-- ----------------------------
INSERT INTO `catalog` VALUES ('1', '测试分类', '7');
INSERT INTO `catalog` VALUES ('2', '测试子分类', '7');

-- ----------------------------
-- Table structure for `group`
-- ----------------------------
DROP TABLE IF EXISTS `group`;
CREATE TABLE `group` (
  `id` tinyint(11) NOT NULL AUTO_INCREMENT,
  `name` text NOT NULL,
  `catalog` tinyint(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of group
-- ----------------------------
INSERT INTO `group` VALUES ('8', '管理员组', '0');
INSERT INTO `group` VALUES ('11', 'CMS组', '1');

-- ----------------------------
-- Table structure for `image`
-- ----------------------------
DROP TABLE IF EXISTS `image`;
CREATE TABLE `image` (
  `id` tinyint(4) NOT NULL AUTO_INCREMENT,
  `name` text NOT NULL,
  `url` text NOT NULL,
  `description` text NOT NULL,
  `creater` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of image
-- ----------------------------
INSERT INTO `image` VALUES ('1', 'secondaaarytile', '/upload/20160104220832_dURohWplZNHeGiGJ_secondarytile.png', 'a反对反对法', '7');

-- ----------------------------
-- Table structure for `link`
-- ----------------------------
DROP TABLE IF EXISTS `link`;
CREATE TABLE `link` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` text NOT NULL,
  `url` text NOT NULL,
  `logo` text NOT NULL,
  `style` tinyint(4) NOT NULL,
  `creater` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of link
-- ----------------------------

-- ----------------------------
-- Table structure for `module`
-- ----------------------------
DROP TABLE IF EXISTS `module`;
CREATE TABLE `module` (
  `id` varchar(36) NOT NULL,
  `name` text NOT NULL,
  `description` text NOT NULL,
  `enableflag` tinyint(4) NOT NULL,
  `defaultflag` tinyint(4) NOT NULL,
  `styleflag` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of module
-- ----------------------------
INSERT INTO `module` VALUES ('f17133ec-63e9-4b46-8757-e6ca1af6fe3e', '博客', '博客模块', '1', '1', '0');
INSERT INTO `module` VALUES ('f17133ec-63e9-4b46-8757-e6ca1af6fe4e', 'Blog', 'Blog2', '0', '0', '0');

-- ----------------------------
-- Table structure for `module_block`
-- ----------------------------
DROP TABLE IF EXISTS `module_block`;
CREATE TABLE `module_block` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` text NOT NULL,
  `owner` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of module_block
-- ----------------------------
INSERT INTO `module_block` VALUES ('3', '导航栏', 'f17133ec-63e9-4b46-8757-e6ca1af6fe3e');
INSERT INTO `module_block` VALUES ('4', ' 文章分类', 'f17133ec-63e9-4b46-8757-e6ca1af6fe3e');
INSERT INTO `module_block` VALUES ('5', '标签云', 'f17133ec-63e9-4b46-8757-e6ca1af6fe3e');
INSERT INTO `module_block` VALUES ('10', '123', 'f17133ec-63e9-4b46-8757-e6ca1af6fe3e');

-- ----------------------------
-- Table structure for `page_block`
-- ----------------------------
DROP TABLE IF EXISTS `page_block`;
CREATE TABLE `page_block` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `url` text,
  `block` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=45 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of page_block
-- ----------------------------
INSERT INTO `page_block` VALUES ('26', '/guestbook/', '5');
INSERT INTO `page_block` VALUES ('27', '/guestbook/', '7');
INSERT INTO `page_block` VALUES ('33', '/aboutsite/', '5');
INSERT INTO `page_block` VALUES ('36', '/', '3');
INSERT INTO `page_block` VALUES ('37', '/', '4');
INSERT INTO `page_block` VALUES ('38', '/', '5');
INSERT INTO `page_block` VALUES ('39', '/view/', '3');
INSERT INTO `page_block` VALUES ('40', '/view/', '4');
INSERT INTO `page_block` VALUES ('41', '/view/', '5');
INSERT INTO `page_block` VALUES ('42', '/aboutme/', '3');
INSERT INTO `page_block` VALUES ('43', '/aboutme/', '4');
INSERT INTO `page_block` VALUES ('44', '/aboutme/', '5');

-- ----------------------------
-- Table structure for `resource`
-- ----------------------------
DROP TABLE IF EXISTS `resource`;
CREATE TABLE `resource` (
  `oid` tinyint(4) NOT NULL AUTO_INCREMENT,
  `name` text NOT NULL,
  `type` tinyint(4) NOT NULL,
  `id` tinyint(4) NOT NULL,
  PRIMARY KEY (`oid`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of resource
-- ----------------------------
INSERT INTO `resource` VALUES ('1', '测试分类', '1', '1');
INSERT INTO `resource` VALUES ('2', '测试子分类', '1', '2');
INSERT INTO `resource` VALUES ('3', '   测试标题', '0', '1');
INSERT INTO `resource` VALUES ('4', 'secondaaarytile', '2', '1');

-- ----------------------------
-- Table structure for `resource_relative`
-- ----------------------------
DROP TABLE IF EXISTS `resource_relative`;
CREATE TABLE `resource_relative` (
  `id` tinyint(4) NOT NULL AUTO_INCREMENT,
  `src` tinyint(4) NOT NULL,
  `srcType` tinyint(4) NOT NULL,
  `dst` tinyint(4) NOT NULL,
  `dstType` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of resource_relative
-- ----------------------------
INSERT INTO `resource_relative` VALUES ('1', '1', '0', '1', '1');
INSERT INTO `resource_relative` VALUES ('2', '1', '2', '1', '1');

-- ----------------------------
-- Table structure for `system_config`
-- ----------------------------
DROP TABLE IF EXISTS `system_config`;
CREATE TABLE `system_config` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `key` text NOT NULL,
  `value` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of system_config
-- ----------------------------
INSERT INTO `system_config` VALUES ('1', '@systemName', 'WebCenter');
INSERT INTO `system_config` VALUES ('2', '@systemDomain', '127.0.0.1:8888');
INSERT INTO `system_config` VALUES ('3', '@systemEMailServer', 'smtp.126.com:25');
INSERT INTO `system_config` VALUES ('4', '@systemEMailAccount', 'rangh@126.com');
INSERT INTO `system_config` VALUES ('5', '@systemEMailPassword', 'hRangh@13924');

-- ----------------------------
-- Table structure for `user`
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` tinyint(4) NOT NULL AUTO_INCREMENT,
  `account` varchar(45) NOT NULL,
  `password` varchar(45) NOT NULL,
  `nickname` varchar(45) NOT NULL,
  `email` varchar(45) NOT NULL,
  `group` text NOT NULL,
  `status` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`,`account`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('7', 'rangh@126.com', 'rangh', 'rangh', 'rangh@126.com', '8', '0');
INSERT INTO `user` VALUES ('36', '123', '123', '123', 'rangh@126.com', '8,11', '0');
INSERT INTO `user` VALUES ('37', 'rangh', 'rangh', 'rangh', 'rangh@126.com', '11', '0');
