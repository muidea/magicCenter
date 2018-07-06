CREATE DATABASE  IF NOT EXISTS `magiccenter_db` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `magiccenter_db`;
-- MySQL dump 10.13  Distrib 5.7.17, for macos10.12 (x86_64)
--
-- Host: 127.0.0.1    Database: magiccenter_db
-- ------------------------------------------------------
-- Server version	5.7.15

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `account_group`
--

DROP TABLE IF EXISTS `account_group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `account_group` (
  `id` int(11) NOT NULL,
  `name` text NOT NULL,
  `catalog` int(11) NOT NULL,
  `description` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `account_group`
--

LOCK TABLES `account_group` WRITE;
/*!40000 ALTER TABLE `account_group` DISABLE KEYS */;
INSERT INTO `account_group` VALUES (0,'系统管理组',0,'系统管理组描述信息'),(1,'MagicBlog用户组',0,'magicBlog用户组');
/*!40000 ALTER TABLE `account_group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `account_user`
--

DROP TABLE IF EXISTS `account_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `account_user` (
  `id` int(11) NOT NULL,
  `account` varchar(45) NOT NULL,
  `password` varchar(45) NOT NULL,
  `email` varchar(45) NOT NULL,
  `groups` text NOT NULL,
  `status` tinyint(4) NOT NULL,
  `registertime` datetime NOT NULL,
  PRIMARY KEY (`id`,`account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `account_user`
--

LOCK TABLES `account_user` WRITE;
/*!40000 ALTER TABLE `account_user` DISABLE KEYS */;
INSERT INTO `account_user` VALUES (0,'admin@muidea.com','123','admin@muidea.com','0',0,'2018-03-20 00:00:00');
/*!40000 ALTER TABLE `account_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `authority_acl`
--

DROP TABLE IF EXISTS `authority_acl`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `authority_acl` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `url` text NOT NULL,
  `method` text NOT NULL,
  `module` text NOT NULL,
  `status` int(11) NOT NULL DEFAULT '0',
  `authgroup` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=70 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `authority_acl`
--

LOCK TABLES `authority_acl` WRITE;
/*!40000 ALTER TABLE `authority_acl` DISABLE KEYS */;
INSERT INTO `authority_acl` VALUES (1,'/system/config/','GET','5b9965b6-b2be-4072-87e2-25b4f96aee54',0,2),(2,'/system/config/','PUT','5b9965b6-b2be-4072-87e2-25b4f96aee54',0,2),(3,'/system/menu/','GET','5b9965b6-b2be-4072-87e2-25b4f96aee54',0,2),(4,'/system/dashboard/','GET','5b9965b6-b2be-4072-87e2-25b4f96aee54',0,2),(5,'/account/user/:id','GET','b9e35167-b2a3-43ae-8c57-9b4379475e47',0,1),(6,'/account/user/','GET','b9e35167-b2a3-43ae-8c57-9b4379475e47',0,1),(7,'/account/user/','POST','b9e35167-b2a3-43ae-8c57-9b4379475e47',0,0),(8,'/account/user/:id','PUT','b9e35167-b2a3-43ae-8c57-9b4379475e47',0,1),(9,'/account/user/:id','DELETE','b9e35167-b2a3-43ae-8c57-9b4379475e47',0,2),(10,'/account/group/:id','GET','b9e35167-b2a3-43ae-8c57-9b4379475e47',0,1),(11,'/account/group/','GET','b9e35167-b2a3-43ae-8c57-9b4379475e47',0,1),(12,'/account/group/','POST','b9e35167-b2a3-43ae-8c57-9b4379475e47',0,2),(13,'/account/group/:id','PUT','b9e35167-b2a3-43ae-8c57-9b4379475e47',0,2),(14,'/account/group/:id','DELETE','b9e35167-b2a3-43ae-8c57-9b4379475e47',0,2),(15,'/fileregistry/file/','POST','b467c59d-10a5-4875-b617-66662f8824fa',0,1),(16,'/fileregistry/file/','GET','b467c59d-10a5-4875-b617-66662f8824fa',0,0),(17,'/fileregistry/file/:id','DELETE','b467c59d-10a5-4875-b617-66662f8824fa',0,1),(18,'/static/**','GET','e9a778e8-1098-4d48-80fc-811782fe2798',0,0),(19,'/authority/acl/','GET','158e11b7-adee-4b0d-afc9-0b47145195bd',0,2),(20,'/authority/acl/:id','GET','158e11b7-adee-4b0d-afc9-0b47145195bd',0,2),(21,'/authority/acl/','POST','158e11b7-adee-4b0d-afc9-0b47145195bd',0,2),(22,'/authority/acl/:id','DELETE','158e11b7-adee-4b0d-afc9-0b47145195bd',0,2),(23,'/authority/acl/:id','PUT','158e11b7-adee-4b0d-afc9-0b47145195bd',0,2),(24,'/authority/acls/','PUT','158e11b7-adee-4b0d-afc9-0b47145195bd',0,2),(25,'/authority/acl/authgroup/:id','GET','158e11b7-adee-4b0d-afc9-0b47145195bd',0,2),(26,'/authority/acl/authgroup/:id','PUT','158e11b7-adee-4b0d-afc9-0b47145195bd',0,2),(27,'/authority/module/','GET','158e11b7-adee-4b0d-afc9-0b47145195bd',0,2),(28,'/authority/module/:id','GET','158e11b7-adee-4b0d-afc9-0b47145195bd',0,2),(29,'/authority/module/:id','PUT','158e11b7-adee-4b0d-afc9-0b47145195bd',0,2),(30,'/authority/user/','GET','158e11b7-adee-4b0d-afc9-0b47145195bd',0,1),(31,'/authority/user/:id','GET','158e11b7-adee-4b0d-afc9-0b47145195bd',0,2),(32,'/authority/user/:id','PUT','158e11b7-adee-4b0d-afc9-0b47145195bd',0,1),(33,'/authority/endpoint/','GET','158e11b7-adee-4b0d-afc9-0b47145195bd',0,2),(34,'/authority/endpoint/:id','GET','158e11b7-adee-4b0d-afc9-0b47145195bd',0,2),(35,'/authority/endpoint/','POST','158e11b7-adee-4b0d-afc9-0b47145195bd',0,2),(36,'/authority/endpoint/:id','DELETE','158e11b7-adee-4b0d-afc9-0b47145195bd',0,2),(37,'/authority/endpoint/:id','PUT','158e11b7-adee-4b0d-afc9-0b47145195bd',0,2),(38,'/authority/endpoint/verify/:id','GET','158e11b7-adee-4b0d-afc9-0b47145195bd',0,0),(39,'/module/','GET','a86ebf5a-9666-4b0d-a12c-acb0c91a03f5',0,2),(40,'/module/:id','GET','a86ebf5a-9666-4b0d-a12c-acb0c91a03f5',0,2),(41,'/cache/item/:id','GET','0424492f-420a-42fb-9106-3882c07bf99e',0,1),(42,'/cache/item/','POST','0424492f-420a-42fb-9106-3882c07bf99e',0,1),(43,'/cache/item/:id','DELETE','0424492f-420a-42fb-9106-3882c07bf99e',0,1),(44,'/content/summary/:id','GET','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,1),(45,'/content/article/:id','GET','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,0),(46,'/content/articles/','GET','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,0),(47,'/content/article/','POST','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,1),(48,'/content/article/:id','PUT','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,1),(49,'/content/article/:id','DELETE','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,2),(50,'/content/catalog/:id','GET','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,0),(51,'/content/catalog/','GET','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,0),(52,'/content/catalogs/','GET','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,0),(53,'/content/catalog/','POST','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,1),(54,'/content/catalog/:id','PUT','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,1),(55,'/content/catalog/:id','DELETE','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,2),(56,'/content/link/:id','GET','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,0),(57,'/content/links/','GET','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,0),(58,'/content/link/','POST','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,1),(59,'/content/link/:id','PUT','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,1),(60,'/content/link/:id','DELETE','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,2),(61,'/content/media/:id','GET','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,0),(62,'/content/medias/','GET','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,0),(63,'/content/media/','POST','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,1),(64,'/content/media/batch/','POST','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,1),(65,'/content/media/:id','PUT','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,1),(66,'/content/media/:id','DELETE','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,2),(67,'/cas/user/','POST','759a2ee4-147a-4169-ba89-15c0c692bc16',0,0),(68,'/cas/user/','DELETE','759a2ee4-147a-4169-ba89-15c0c692bc16',0,1),(69,'/cas/user/','GET','759a2ee4-147a-4169-ba89-15c0c692bc16',0,1);
/*!40000 ALTER TABLE `authority_acl` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `authority_endpoint`
--

DROP TABLE IF EXISTS `authority_endpoint`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `authority_endpoint` (
  `eid` int(11) NOT NULL AUTO_INCREMENT,
  `id` text NOT NULL,
  `name` text NOT NULL,
  `description` text NOT NULL,
  `user` text NOT NULL,
  `status` int(11) NOT NULL,
  `authToken` text,
  PRIMARY KEY (`eid`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `authority_endpoint`
--

LOCK TABLES `authority_endpoint` WRITE;
/*!40000 ALTER TABLE `authority_endpoint` DISABLE KEYS */;
INSERT INTO `authority_endpoint` VALUES (5,'f0e078a8-6de8-4273-88a4-dccef60ff88f','magicBlog','magicBlog是一个博客应用','0',0,'yTtWiuuoGifPVfcK5Mf4mdu8mGl78E3y');
/*!40000 ALTER TABLE `authority_endpoint` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `authority_module`
--

DROP TABLE IF EXISTS `authority_module`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `authority_module` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user` int(11) NOT NULL,
  `module` text NOT NULL,
  `authgroup` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=114 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `authority_module`
--

LOCK TABLES `authority_module` WRITE;
/*!40000 ALTER TABLE `authority_module` DISABLE KEYS */;
INSERT INTO `authority_module` VALUES (39,0,'158e11b7-adee-4b0d-afc9-0b47145195bd',2),(63,0,'5b9965b6-b2be-4072-87e2-25b4f96aee54',2),(66,0,'7fe4a6fa-b73a-401f-bd37-71e76670d18c',2),(68,0,'12675100-da3d-42f3-a3fb-68aadc189730',2),(70,0,'e9a778e8-1098-4d48-80fc-811782fe2798',2),(78,0,'0424492f-420a-42fb-9106-3882c07bf99e',2),(80,0,'a86ebf5a-9666-4b0d-a12c-acb0c91a03f5',2),(86,0,'b9e35167-b2a3-43ae-8c57-9b4379475e47',2),(90,0,'b467c59d-10a5-4875-b617-66662f8824fa',2),(109,1,'0424492f-420a-42fb-9106-3882c07bf99e',1),(110,0,'759a2ee4-147a-4169-ba89-15c0c692bc16',2),(111,1,'759a2ee4-147a-4169-ba89-15c0c692bc16',1),(112,0,'3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',2),(113,53,'3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',1);
/*!40000 ALTER TABLE `authority_module` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `common_fileregistry`
--

DROP TABLE IF EXISTS `common_fileregistry`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `common_fileregistry` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `filetoken` text NOT NULL,
  `filename` text NOT NULL,
  `filepath` text NOT NULL,
  `uploaddate` datetime NOT NULL,
  `reserveflag` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `common_fileregistry`
--

LOCK TABLES `common_fileregistry` WRITE;
/*!40000 ALTER TABLE `common_fileregistry` DISABLE KEYS */;
INSERT INTO `common_fileregistry` VALUES (31,'d3fxpiq6igoo5scvukeiafsqgxcr3f2h','kubernetes指南.pdf','static/upload/d3fxpiq6igoo5scvukeiafsqgxcr3f2h/kubernetes指南.pdf','2018-07-06 16:39:30',0),(32,'v4ggbcusyxknglegfm9lepwrmxxxskil','当mongodb遇见iot.pdf','static/upload/v4ggbcusyxknglegfm9lepwrmxxxskil/当mongodb遇见iot.pdf','2018-07-06 17:08:41',0),(33,'1udryvloas0inuslpgtxxaagipolcadk','4江骏.pdf','static/upload/1udryvloas0inuslpgtxxaagipolcadk/4江骏.pdf','2018-07-06 17:08:47',0);
/*!40000 ALTER TABLE `common_fileregistry` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `common_option`
--

DROP TABLE IF EXISTS `common_option`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `common_option` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `key` text NOT NULL,
  `value` text NOT NULL,
  `owner` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `common_option`
--

LOCK TABLES `common_option` WRITE;
/*!40000 ALTER TABLE `common_option` DISABLE KEYS */;
INSERT INTO `common_option` VALUES (3,'@system_mailServer','smtp.126.com:25','SystemInternalConfig'),(4,'@system_mailAccount','rangh@126.com','SystemInternalConfig'),(5,'@system_mailPassword','hRangh@13924','SystemInternalConfig'),(6,'@application_logo','http://localhost:8888/api/system/','SystemInternalConfig'),(13,'@application_name','magicCenter','SystemInternalConfig'),(14,'@application_description','rangh\'s magicCenter','SystemInternalConfig'),(15,'@application_domain','muidea.com','SystemInternalConfig'),(16,'@system_uploadPath','upload','SystemInternalConfig'),(17,'@system_staticPath','./static/','SystemInternalConfig'),(37,'@application_startupData','startup_TimeStamp:2018-07-04 14:58:43','SystemInternalConfig');
/*!40000 ALTER TABLE `common_option` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `common_resource`
--

DROP TABLE IF EXISTS `common_resource`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `common_resource` (
  `oid` int(11) NOT NULL,
  `id` int(11) NOT NULL,
  `name` text NOT NULL,
  `description` text NOT NULL,
  `type` text NOT NULL,
  `createtime` datetime NOT NULL,
  `owner` int(11) NOT NULL,
  PRIMARY KEY (`oid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `common_resource`
--

LOCK TABLES `common_resource` WRITE;
/*!40000 ALTER TABLE `common_resource` DISABLE KEYS */;
INSERT INTO `common_resource` VALUES (0,0,'默认Content分组','默认分组的描述','catalog','2018-03-01 00:00:00',0),(7,2,'magicBlog','magicBlog 分组','catalog','2018-04-23 23:10:50',1),(9,1,'测试文章','这里一些测试内容','article','2018-05-01 22:24:29',0),(15,3,'About','个人介绍','article','2018-05-05 18:14:32',0),(16,4,'Contact','交流信息','article','2018-05-05 22:21:36',0),(17,8,'Catalog','分类信息','catalog','2018-05-05 22:26:21',0),(18,9,'Index','主页信息','catalog','2018-05-05 22:30:08',0),(19,5,'404','404页面','article','2018-05-05 22:31:11',0),(20,10,'技术文章','这是技术文章的描述','catalog','2018-05-06 13:41:35',0),(42,10,'测试内容,这是一篇测试文章','rules: [ { required: true }, ],','article','2018-06-09 09:42:13',0),(44,11,'ts','','catalog','2018-07-04 16:06:35',0),(45,12,'aaa','','catalog','2018-07-04 16:08:31',0),(46,13,'aa','','catalog','2018-07-04 16:11:17',0),(47,14,'12','','catalog','2018-07-04 16:13:38',0),(50,15,'ca','','catalog','2018-07-04 16:36:34',0),(52,16,'bb','','catalog','2018-07-04 16:37:35',0),(54,17,'gg','','catalog','2018-07-04 16:38:00',0),(57,18,'测试','','catalog','2018-07-04 16:55:48',0),(58,19,'a','','catalog','2018-07-06 11:09:07',0),(61,20,'啊','','catalog','2018-07-06 13:36:29',0),(62,21,'111','','catalog','2018-07-06 14:37:33',0),(65,22,'测试内容','','catalog','2018-07-06 16:37:08',0),(68,5,'kubernetes指南.pdf','大放送','media','2018-07-06 16:39:41',0),(69,23,'分类','','catalog','2018-07-06 17:09:09',0),(70,6,'当mongodb遇见iot.pdf','看看分类效果','media','2018-07-06 17:09:09',0),(71,7,'4江骏.pdf','看看分类效果','media','2018-07-06 17:09:09',0);
/*!40000 ALTER TABLE `common_resource` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `common_resource_relative`
--

DROP TABLE IF EXISTS `common_resource_relative`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `common_resource_relative` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `src` int(11) NOT NULL,
  `dst` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=198 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `common_resource_relative`
--

LOCK TABLES `common_resource_relative` WRITE;
/*!40000 ALTER TABLE `common_resource_relative` DISABLE KEYS */;
INSERT INTO `common_resource_relative` VALUES (52,17,7),(55,18,7),(63,20,17),(64,20,18),(80,9,17),(81,9,18),(154,15,7),(155,16,7),(156,19,7),(159,42,17),(160,42,18),(161,44,0),(162,45,0),(163,46,0),(164,47,0),(167,50,0),(169,52,0),(171,54,0),(174,57,0),(177,58,0),(182,61,0),(186,62,0),(189,65,0),(192,68,57),(193,69,0),(194,70,57),(195,70,69),(196,71,57),(197,71,69);
/*!40000 ALTER TABLE `common_resource_relative` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `content_article`
--

DROP TABLE IF EXISTS `content_article`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `content_article` (
  `id` int(11) NOT NULL,
  `title` tinytext NOT NULL,
  `content` text NOT NULL,
  `createdate` datetime NOT NULL,
  `creater` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `content_article`
--

LOCK TABLES `content_article` WRITE;
/*!40000 ALTER TABLE `content_article` DISABLE KEYS */;
INSERT INTO `content_article` VALUES (1,'测试文章','\n不少朋友都知道我在“[++极客时间++](https://time.geekbang.org/)”上开了一个收费专栏，这个专栏会开设大约一年的时间，一共会发布104篇文章。现在，我在上面以每周两篇文章的频率已发布了27篇文章了，也就是差不多两个半月的时间。新的一年开始了，写专栏这个事对我来说是第一次，在这个过程中有一些感想，所以，我想在这里说一下这些感受和一些相关的故事，算是一个记录，也算是对我专栏的正式介绍，还希望能得到、大家的喜欢和指点。（当然，CoolShell这边还是会持续更新的）\n\n测试内容\n\n​\n\n​\n\n看看效果\n\n- **A\n- **B\n','2018-05-08 23:33:43',0),(3,'About','## 个人介绍\n\n我叫黄冬朋，互联网技术爱好者，Gopher!\n\n2016年以前专注于C++跨平台服务器后台应用系统开发，擅长通讯服务器，数据管理软件架构开发。\n\n2014年开始接触Golang，曾经也是Python的爱好者和推广者，自从接触到Golang后，就被它的设计哲理所吸引。 开始各种场合推荐Go，并逐步开始使用Go进行系统开发。\n\n欢迎跟大家交流Kubernetes，Docker，Cloud。\n\n希望能跟大家多多交流，微信：21883911\n','2018-06-10 21:24:13',0),(4,'Contact','## 站点介绍\n\n记录建设云基础平台过程中经历的心路历程，欢迎与大家一起相互交流。\n\n本站使用的技术栈：Docker + Golang + React + MySQL\n\n前后端都是由我个人纯手工打造，引用了部分开源项目(后面单独说明)。本人C++后台开发出身，部分内容可能会姿势不对，欢迎大家拍砖！\n\n也欢迎与大家相互交流，分享心得，也诚邀美工和前端的朋友一起合作，欢迎联系！\n\n交换站点链接，请加我微信并说明！\n\n1、为什么要建本站？\n\n为了实现多年夙愿，也为了对基础平台进行功能验证。\n\n2、站点代码开源么？\n\n开源，GitHub地址: [magicBlog](https://github.com/muidea/magicBlog)\n\n3、本站引用到的项目\n\nAnt Design\n','2018-06-10 21:30:28',0),(5,'404','# **找不到内容了！**\n\n**如果你喜欢本站，欢迎交流！**\n','2018-06-10 21:32:20',0),(10,'测试内容,这是一篇测试文章','rules: [ { required: true }, ],\n\n这里只是看看效果，\n\n看看效果怎样啊\n','2018-06-13 22:29:34',0);
/*!40000 ALTER TABLE `content_article` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `content_catalog`
--

DROP TABLE IF EXISTS `content_catalog`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `content_catalog` (
  `id` int(11) NOT NULL,
  `name` text NOT NULL,
  `description` text,
  `createdate` datetime NOT NULL,
  `creater` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `content_catalog`
--

LOCK TABLES `content_catalog` WRITE;
/*!40000 ALTER TABLE `content_catalog` DISABLE KEYS */;
INSERT INTO `content_catalog` VALUES (0,'默认Content分组','系统默认的Content分组','2018-03-01 00:00:00',0),(2,'magicBlog','MagicBlog catalog description','2018-06-10 21:33:15',0),(8,'Catalog','Catalog分类','2018-05-05 22:26:21',0),(9,'Index','Index','2018-05-05 22:30:08',0),(10,'技术文章','技术文章','2018-05-06 13:41:35',0),(11,'ts','','2018-07-04 16:06:35',0),(12,'aaa','','2018-07-04 16:08:31',0),(13,'aa','','2018-07-04 16:11:17',0),(14,'12','','2018-07-04 16:13:38',0),(15,'ca','','2018-07-04 16:36:34',0),(16,'bb','','2018-07-04 16:37:35',0),(17,'gg','','2018-07-04 16:38:00',0),(18,'测试','','2018-07-04 16:55:48',0),(19,'a','','2018-07-06 11:09:07',0),(20,'啊','','2018-07-06 13:36:29',0),(21,'111','','2018-07-06 14:37:33',0),(22,'测试内容','','2018-07-06 16:37:08',0),(23,'分类','','2018-07-06 17:09:09',0);
/*!40000 ALTER TABLE `content_catalog` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `content_link`
--

DROP TABLE IF EXISTS `content_link`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `content_link` (
  `id` int(10) unsigned NOT NULL,
  `name` text NOT NULL,
  `description` text NOT NULL,
  `url` text NOT NULL,
  `logo` text NOT NULL,
  `createdate` datetime NOT NULL,
  `creater` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `content_link`
--

LOCK TABLES `content_link` WRITE;
/*!40000 ALTER TABLE `content_link` DISABLE KEYS */;
/*!40000 ALTER TABLE `content_link` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `content_media`
--

DROP TABLE IF EXISTS `content_media`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `content_media` (
  `id` int(11) NOT NULL,
  `name` text NOT NULL,
  `description` text NOT NULL,
  `fileToken` text NOT NULL,
  `createdate` datetime NOT NULL,
  `creater` int(11) NOT NULL,
  `expiration` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `content_media`
--

LOCK TABLES `content_media` WRITE;
/*!40000 ALTER TABLE `content_media` DISABLE KEYS */;
INSERT INTO `content_media` VALUES (5,'kubernetes指南.pdf','大放送','d3fxpiq6igoo5scvukeiafsqgxcr3f2h','2018-07-06 16:39:41',0,10),(6,'当mongodb遇见iot.pdf','看看分类效果','v4ggbcusyxknglegfm9lepwrmxxxskil','2018-07-06 17:09:09',0,10),(7,'4江骏.pdf','看看分类效果','1udryvloas0inuslpgtxxaagipolcadk','2018-07-06 17:09:09',0,10);
/*!40000 ALTER TABLE `content_media` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-07-06 23:37:03
