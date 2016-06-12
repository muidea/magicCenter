/*
Navicat MySQL Data Transfer

Source Server         : magicid
Source Server Version : 50624
Source Host           : localhost:3306
Source Database       : magiccenter_db

Target Server Type    : MYSQL
Target Server Version : 50624
File Encoding         : 65001

Date: 2016-06-12 23:13:04
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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of article
-- ----------------------------
INSERT INTO `article` VALUES ('2', '测试文章', '这是一些测试内容，占坑用的信息。。。。。&lt;div&gt;&lt;div class=&#34;cLeft&#34; style=&#34;margin: 0px 27px 20px 0px; padding: 0px 0px 0px 10px; color: rgb(51, 51, 51); font-family: simsun; font-size: 12px; display: inline-block; float: left; width: 637px; line-height: 20px; background-color: rgb(255, 255, 255);&#34;&gt;&lt;p style=&#34;margin: 0px; padding: 0px; text-align: justify; clear: both;&#34;&gt;岗位职责 1、 负责PC端/手机端UI设计； 2、 根据交互设计和产品规划，完成产品相关的用户界面视觉设计； 3、 根据视觉设计的发展趋势及用户研究不断优化产品UI。 岗位要求： 1、 大专或以上学历，具有艺术设计或计算机关专业优先； 2、 具备1年以上网站或APP设计工作经验； 3、 熟练掌握P...&lt;/p&gt;&lt;/div&gt;&lt;div class=&#34;cRight&#34; style=&#34;margin: 0px 0px 20px; padding: 0px; color: rgb(51, 51, 51); font-family: simsun; font-size: 12px; display: inline-block; float: left; line-height: normal; background-color: rgb(255, 255, 255);&#34;&gt;&lt;a class=&#34;applyJob&#34; style=&#34;color: white; outline: none; display: block; width: 95px; height: 30px; text-align: center; line-height: 30px; font-size: 14px; font-weight: bold; margin-top: 10px; border: 1px solid rgb(255, 140, 7); border-radius: 3px; background-color: rgb(255, 160, 21);&#34;&gt;申请职位&lt;/a&gt;&lt;a class=&#34;collectJob&#34; style=&#34;outline: none; display: block; width: 80px; height: 30px; text-align: center; line-height: 30px; font-size: 14px; font-weight: bold; margin-top: 10px; padding-left: 18px; background: url(&amp;quot;http://img01.zpin.net.cn/2014/seo/images/collect.png&amp;quot;) 5px 5px no-repeat;&#34;&gt;收藏职位&lt;/a&gt;&lt;div&gt;&lt;br&gt;&lt;/div&gt;&lt;/div&gt;&lt;/div&gt;', '1', '2016-05-25 21:46:43');
INSERT INTO `article` VALUES ('3', '关于本站', '本站基本信息介绍', '1', '2016-05-20 22:19:51');
INSERT INTO `article` VALUES ('4', '关于本人', '本人基本信息介绍', '1', '2016-05-20 22:20:09');

-- ----------------------------
-- Table structure for `block`
-- ----------------------------
DROP TABLE IF EXISTS `block`;
CREATE TABLE `block` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` text NOT NULL,
  `tag` text NOT NULL,
  `style` tinyint(4) unsigned zerofill NOT NULL,
  `owner` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of block
-- ----------------------------
INSERT INTO `block` VALUES ('1', '导航栏', 'nav', '0000', 'f17133ec-63e9-4b46-8757-e6ca1af6fe3e');
INSERT INTO `block` VALUES ('2', '链接区', 'link', '0000', 'f17133ec-63e9-4b46-8757-e6ca1af6fe3e');
INSERT INTO `block` VALUES ('3', '标签云', 'tags', '0000', 'f17133ec-63e9-4b46-8757-e6ca1af6fe3e');
INSERT INTO `block` VALUES ('4', '文章列表', 'list', '0000', 'f17133ec-63e9-4b46-8757-e6ca1af6fe3e');
INSERT INTO `block` VALUES ('12', '文章视图', 'post', '0001', 'f17133ec-63e9-4b46-8757-e6ca1af6fe3e');

-- ----------------------------
-- Table structure for `catalog`
-- ----------------------------
DROP TABLE IF EXISTS `catalog`;
CREATE TABLE `catalog` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` text NOT NULL,
  `creater` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of catalog
-- ----------------------------
INSERT INTO `catalog` VALUES ('4', '测试分类', '1');
INSERT INTO `catalog` VALUES ('5', 'Linux相关', '1');
INSERT INTO `catalog` VALUES ('6', 'Go', '1');
INSERT INTO `catalog` VALUES ('7', '系统部署', '1');

-- ----------------------------
-- Table structure for `entity`
-- ----------------------------
DROP TABLE IF EXISTS `entity`;
CREATE TABLE `entity` (
  `id` varchar(36) NOT NULL,
  `name` text NOT NULL,
  `description` text NOT NULL,
  `enableflag` tinyint(4) NOT NULL,
  `defaultflag` tinyint(4) NOT NULL,
  `module` varchar(36) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of entity
-- ----------------------------

-- ----------------------------
-- Table structure for `group`
-- ----------------------------
DROP TABLE IF EXISTS `group`;
CREATE TABLE `group` (
  `id` tinyint(11) NOT NULL AUTO_INCREMENT,
  `name` text NOT NULL,
  `creater` tinyint(4) NOT NULL,
  `catalog` tinyint(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of group
-- ----------------------------
INSERT INTO `group` VALUES ('8', '管理员组', '0', '0');
INSERT INTO `group` VALUES ('16', 'CMS组', '0', '1');

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of image
-- ----------------------------

-- ----------------------------
-- Table structure for `item`
-- ----------------------------
DROP TABLE IF EXISTS `item`;
CREATE TABLE `item` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `rid` int(11) NOT NULL,
  `rtype` tinyint(4) NOT NULL,
  `owner` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=105 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of item
-- ----------------------------
INSERT INTO `item` VALUES ('54', '4', '1', '3');
INSERT INTO `item` VALUES ('55', '5', '1', '3');
INSERT INTO `item` VALUES ('56', '5', '1', '9');
INSERT INTO `item` VALUES ('67', '2', '0', '4');
INSERT INTO `item` VALUES ('68', '4', '1', '4');
INSERT INTO `item` VALUES ('69', '5', '1', '4');
INSERT INTO `item` VALUES ('70', '6', '1', '4');
INSERT INTO `item` VALUES ('71', '7', '1', '4');
INSERT INTO `item` VALUES ('85', '4', '1', '2');
INSERT INTO `item` VALUES ('86', '5', '1', '2');
INSERT INTO `item` VALUES ('87', '6', '1', '2');
INSERT INTO `item` VALUES ('88', '1', '3', '2');
INSERT INTO `item` VALUES ('93', '2', '0', '12');
INSERT INTO `item` VALUES ('94', '3', '0', '12');
INSERT INTO `item` VALUES ('95', '4', '0', '12');
INSERT INTO `item` VALUES ('101', '3', '0', '1');
INSERT INTO `item` VALUES ('102', '4', '0', '1');
INSERT INTO `item` VALUES ('103', '4', '1', '1');
INSERT INTO `item` VALUES ('104', '2', '3', '1');

-- ----------------------------
-- Table structure for `link`
-- ----------------------------
DROP TABLE IF EXISTS `link`;
CREATE TABLE `link` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` text NOT NULL,
  `url` text NOT NULL,
  `logo` text NOT NULL,
  `creater` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of link
-- ----------------------------
INSERT INTO `link` VALUES ('1', '123', 'http://cn.bing.com/', '43423423', '1');
INSERT INTO `link` VALUES ('2', '首页', 'http://127.0.0.11:8888/', 'aaa', '1');

-- ----------------------------
-- Table structure for `module`
-- ----------------------------
DROP TABLE IF EXISTS `module`;
CREATE TABLE `module` (
  `id` varchar(36) NOT NULL,
  `name` text NOT NULL,
  `uri` text NOT NULL,
  `description` text NOT NULL,
  `enableflag` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of module
-- ----------------------------
INSERT INTO `module` VALUES ('f17133ec-63e9-4b46-8757-e6ca1af6fe3e', '博客', '/blog', '博客模块', '1');
INSERT INTO `module` VALUES ('f17133ec-63e9-4b46-8757-e6ca1af6fe4e', 'Blog', '/blog2', 'Blog2', '0');

-- ----------------------------
-- Table structure for `option`
-- ----------------------------
DROP TABLE IF EXISTS `option`;
CREATE TABLE `option` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `key` text NOT NULL,
  `value` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of option
-- ----------------------------
INSERT INTO `option` VALUES ('1', '@application_name', 'MagicCenter');
INSERT INTO `option` VALUES ('2', '@application_domain', '127.0.0.1:8888');
INSERT INTO `option` VALUES ('3', '@system_mailServer', 'smtp.126.com:25');
INSERT INTO `option` VALUES ('4', '@system_mailAccount', 'rangh@126.com');
INSERT INTO `option` VALUES ('5', '@system_mailPassword', 'hRangh@13924');
INSERT INTO `option` VALUES ('6', '@application_logo', '223');
INSERT INTO `option` VALUES ('7', '@system_defaultModule', 'f17133ec-63e9-4b46-8757-e6ca1af6fe3e');

-- ----------------------------
-- Table structure for `page`
-- ----------------------------
DROP TABLE IF EXISTS `page`;
CREATE TABLE `page` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `owner` text NOT NULL,
  `url` text NOT NULL,
  `block` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of page
-- ----------------------------
INSERT INTO `page` VALUES ('22', 'f17133ec-63e9-4b46-8757-e6ca1af6fe4e', '/view/', '9');
INSERT INTO `page` VALUES ('28', 'f17133ec-63e9-4b46-8757-e6ca1af6fe3e', '/view/', '1');
INSERT INTO `page` VALUES ('29', 'f17133ec-63e9-4b46-8757-e6ca1af6fe3e', '/view/', '2');
INSERT INTO `page` VALUES ('30', 'f17133ec-63e9-4b46-8757-e6ca1af6fe3e', '/view/', '3');
INSERT INTO `page` VALUES ('40', 'f17133ec-63e9-4b46-8757-e6ca1af6fe3e', '/catalog/', '1');
INSERT INTO `page` VALUES ('41', 'f17133ec-63e9-4b46-8757-e6ca1af6fe3e', '/catalog/', '2');
INSERT INTO `page` VALUES ('42', 'f17133ec-63e9-4b46-8757-e6ca1af6fe3e', '/catalog/', '4');
INSERT INTO `page` VALUES ('43', 'f17133ec-63e9-4b46-8757-e6ca1af6fe3e', '/', '1');
INSERT INTO `page` VALUES ('44', 'f17133ec-63e9-4b46-8757-e6ca1af6fe3e', '/', '2');
INSERT INTO `page` VALUES ('45', 'f17133ec-63e9-4b46-8757-e6ca1af6fe3e', '/', '3');
INSERT INTO `page` VALUES ('46', 'f17133ec-63e9-4b46-8757-e6ca1af6fe3e', '/', '4');
INSERT INTO `page` VALUES ('47', 'f17133ec-63e9-4b46-8757-e6ca1af6fe3e', '/', '12');

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
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of resource
-- ----------------------------
INSERT INTO `resource` VALUES ('1', '测试分类', '1', '1');
INSERT INTO `resource` VALUES ('2', 'Linux相关', '1', '2');
INSERT INTO `resource` VALUES ('7', 'Go', '1', '3');
INSERT INTO `resource` VALUES ('9', '测试分类', '1', '4');
INSERT INTO `resource` VALUES ('10', 'Linux相关', '1', '5');
INSERT INTO `resource` VALUES ('11', 'Go', '1', '6');
INSERT INTO `resource` VALUES ('12', '测试文章', '0', '2');
INSERT INTO `resource` VALUES ('13', '系统部署', '1', '7');
INSERT INTO `resource` VALUES ('14', '关于本站', '0', '3');
INSERT INTO `resource` VALUES ('15', '关于本人', '0', '4');
INSERT INTO `resource` VALUES ('16', '123', '3', '1');
INSERT INTO `resource` VALUES ('17', '首页', '3', '2');

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
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of resource_relative
-- ----------------------------
INSERT INTO `resource_relative` VALUES ('17', '7', '1', '4', '1');
INSERT INTO `resource_relative` VALUES ('18', '7', '1', '5', '1');
INSERT INTO `resource_relative` VALUES ('19', '7', '1', '6', '1');
INSERT INTO `resource_relative` VALUES ('22', '3', '0', '7', '1');
INSERT INTO `resource_relative` VALUES ('23', '4', '0', '7', '1');
INSERT INTO `resource_relative` VALUES ('25', '2', '3', '7', '1');
INSERT INTO `resource_relative` VALUES ('26', '1', '3', '7', '1');
INSERT INTO `resource_relative` VALUES ('27', '2', '0', '4', '1');
INSERT INTO `resource_relative` VALUES ('28', '2', '0', '7', '1');

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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', 'rangh', '123', 'rangh', 'rangh@126.com', '8', '0');
INSERT INTO `user` VALUES ('2', 'test', '123', 'test', 'rangh@foxmail.com', '8,16', '2');

-- ----------------------------
-- Table structure for `user_function`
-- ----------------------------
DROP TABLE IF EXISTS `user_function`;
CREATE TABLE `user_function` (
  `id` tinyint(11) NOT NULL AUTO_INCREMENT,
  `uid` tinyint(11) NOT NULL,
  `entity` char(36) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user_function
-- ----------------------------
