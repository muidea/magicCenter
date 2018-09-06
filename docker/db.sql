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
  `reserve` tinyint(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `account_group`
--

LOCK TABLES `account_group` WRITE;
/*!40000 ALTER TABLE `account_group` DISABLE KEYS */;
INSERT INTO `account_group` VALUES (1,'系统管理组',0,'系统管理组描述信息',1),(2,'普通用户组',0,'普通用户组描述信息',1);
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
  `reserve` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`,`account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `account_user`
--

LOCK TABLES `account_user` WRITE;
/*!40000 ALTER TABLE `account_user` DISABLE KEYS */;
INSERT INTO `account_user` VALUES (1,'admin','123','admin@muidea.com','1',0,'2018-03-20 00:00:00',1);
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
) ENGINE=InnoDB AUTO_INCREMENT=78 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `authority_acl`
--

LOCK TABLES `authority_acl` WRITE;
/*!40000 ALTER TABLE `authority_acl` DISABLE KEYS */;
INSERT INTO `authority_acl` VALUES (1,'/fileregistry/file/','POST','b467c59d-10a5-4875-b617-66662f8824fa',0,1),(2,'/fileregistry/file/','GET','b467c59d-10a5-4875-b617-66662f8824fa',0,0),(3,'/fileregistry/file/:id','DELETE','b467c59d-10a5-4875-b617-66662f8824fa',0,1),(4,'/static/**','GET','e9a778e8-1098-4d48-80fc-811782fe2798',0,0),(5,'/module/','GET','a86ebf5a-9666-4b0d-a12c-acb0c91a03f5',0,2),(6,'/module/:id','GET','a86ebf5a-9666-4b0d-a12c-acb0c91a03f5',0,2),(7,'/cache/item/:id','GET','0424492f-420a-42fb-9106-3882c07bf99e',0,1),(8,'/cache/item/','POST','0424492f-420a-42fb-9106-3882c07bf99e',0,1),(9,'/cache/item/:id','DELETE','0424492f-420a-42fb-9106-3882c07bf99e',0,1),(10,'/account/user/:id','GET','b9e35167-b2a3-43ae-8c57-9b4379475e47',0,1),(11,'/account/user/','GET','b9e35167-b2a3-43ae-8c57-9b4379475e47',0,1),(12,'/account/user/','POST','b9e35167-b2a3-43ae-8c57-9b4379475e47',0,0),(13,'/account/user/:id','PUT','b9e35167-b2a3-43ae-8c57-9b4379475e47',0,1),(14,'/account/user/:id','DELETE','b9e35167-b2a3-43ae-8c57-9b4379475e47',0,2),(15,'/account/group/:id','GET','b9e35167-b2a3-43ae-8c57-9b4379475e47',0,1),(16,'/account/group/','GET','b9e35167-b2a3-43ae-8c57-9b4379475e47',0,1),(17,'/account/group/','POST','b9e35167-b2a3-43ae-8c57-9b4379475e47',0,2),(18,'/account/group/:id','PUT','b9e35167-b2a3-43ae-8c57-9b4379475e47',0,2),(19,'/account/group/:id','DELETE','b9e35167-b2a3-43ae-8c57-9b4379475e47',0,2),(20,'/endpoint/registry/','GET','fa404076-ebf1-4ad6-bedf-2fd6d114ab05',0,2),(21,'/endpoint/registry/:id','GET','fa404076-ebf1-4ad6-bedf-2fd6d114ab05',0,2),(22,'/endpoint/registry/','POST','fa404076-ebf1-4ad6-bedf-2fd6d114ab05',0,2),(23,'/endpoint/registry/:id','DELETE','fa404076-ebf1-4ad6-bedf-2fd6d114ab05',0,2),(24,'/endpoint/registry/:id','PUT','fa404076-ebf1-4ad6-bedf-2fd6d114ab05',0,2),(25,'/cas/user/','POST','759a2ee4-147a-4169-ba89-15c0c692bc16',0,0),(26,'/cas/user/','DELETE','759a2ee4-147a-4169-ba89-15c0c692bc16',0,1),(27,'/cas/user/','GET','759a2ee4-147a-4169-ba89-15c0c692bc16',0,1),(28,'/cas/endpoint/','POST','759a2ee4-147a-4169-ba89-15c0c692bc16',0,0),(29,'/cas/endpoint/','DELETE','759a2ee4-147a-4169-ba89-15c0c692bc16',0,1),(30,'/cas/endpoint/','GET','759a2ee4-147a-4169-ba89-15c0c692bc16',0,1),(31,'/system/config/','GET','5b9965b6-b2be-4072-87e2-25b4f96aee54',0,2),(32,'/system/config/','PUT','5b9965b6-b2be-4072-87e2-25b4f96aee54',0,2),(33,'/system/menu/','GET','5b9965b6-b2be-4072-87e2-25b4f96aee54',0,2),(34,'/system/dashboard/','GET','5b9965b6-b2be-4072-87e2-25b4f96aee54',0,2),(35,'/content/summary/','GET','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,1),(36,'/content/summary/:id','GET','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,1),(37,'/content/summarys/','GET','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,1),(38,'/content/article/:id','GET','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,0),(39,'/content/articles/','GET','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,0),(40,'/content/article/','POST','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,1),(41,'/content/article/:id','PUT','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,1),(42,'/content/article/:id','DELETE','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,2),(43,'/content/catalog/:id','GET','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,0),(44,'/content/catalog/','GET','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,0),(45,'/content/catalogs/','GET','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,0),(46,'/content/catalog/','POST','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,1),(47,'/content/catalog/:id','PUT','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,1),(48,'/content/catalog/:id','DELETE','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,2),(49,'/content/link/:id','GET','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,0),(50,'/content/links/','GET','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,0),(51,'/content/link/','POST','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,1),(52,'/content/link/:id','PUT','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,1),(53,'/content/link/:id','DELETE','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,2),(54,'/content/media/:id','GET','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,0),(55,'/content/medias/','GET','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,0),(56,'/content/media/','POST','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,1),(57,'/content/media/batch/','POST','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,1),(58,'/content/media/:id','PUT','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,1),(59,'/content/media/:id','DELETE','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,2),(60,'/content/comments/','GET','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,0),(61,'/content/comment/','POST','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,1),(62,'/content/comment/:id','PUT','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,1),(63,'/content/comment/:id','DELETE','3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',0,2),(64,'/authority/acl/','GET','158e11b7-adee-4b0d-afc9-0b47145195bd',0,2),(65,'/authority/acl/:id','GET','158e11b7-adee-4b0d-afc9-0b47145195bd',0,2),(66,'/authority/acl/','POST','158e11b7-adee-4b0d-afc9-0b47145195bd',0,2),(67,'/authority/acl/:id','DELETE','158e11b7-adee-4b0d-afc9-0b47145195bd',0,2),(68,'/authority/acl/:id','PUT','158e11b7-adee-4b0d-afc9-0b47145195bd',0,2),(69,'/authority/acls/','PUT','158e11b7-adee-4b0d-afc9-0b47145195bd',0,2),(70,'/authority/acl/authgroup/:id','GET','158e11b7-adee-4b0d-afc9-0b47145195bd',0,2),(71,'/authority/acl/authgroup/:id','PUT','158e11b7-adee-4b0d-afc9-0b47145195bd',0,2),(72,'/authority/module/','GET','158e11b7-adee-4b0d-afc9-0b47145195bd',0,2),(73,'/authority/module/:id','GET','158e11b7-adee-4b0d-afc9-0b47145195bd',0,2),(74,'/authority/module/:id','PUT','158e11b7-adee-4b0d-afc9-0b47145195bd',0,2),(75,'/authority/user/','GET','158e11b7-adee-4b0d-afc9-0b47145195bd',0,1),(76,'/authority/user/:id','GET','158e11b7-adee-4b0d-afc9-0b47145195bd',0,2),(77,'/authority/user/:id','PUT','158e11b7-adee-4b0d-afc9-0b47145195bd',0,1);
/*!40000 ALTER TABLE `authority_acl` ENABLE KEYS */;
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
INSERT INTO `authority_module` VALUES (39,0,'158e11b7-adee-4b0d-afc9-0b47145195bd',2),(63,0,'5b9965b6-b2be-4072-87e2-25b4f96aee54',2),(70,0,'e9a778e8-1098-4d48-80fc-811782fe2798',2),(78,0,'0424492f-420a-42fb-9106-3882c07bf99e',2),(80,0,'a86ebf5a-9666-4b0d-a12c-acb0c91a03f5',2),(86,0,'b9e35167-b2a3-43ae-8c57-9b4379475e47',2),(90,0,'b467c59d-10a5-4875-b617-66662f8824fa',2),(109,1,'0424492f-420a-42fb-9106-3882c07bf99e',1),(110,0,'759a2ee4-147a-4169-ba89-15c0c692bc16',2),(111,1,'759a2ee4-147a-4169-ba89-15c0c692bc16',1),(112,0,'3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',2),(113,53,'3a7123ec-63f0-5e46-1234-e6ca1af6fe4e',1);
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
) ENGINE=InnoDB AUTO_INCREMENT=221 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `common_fileregistry`
--

LOCK TABLES `common_fileregistry` WRITE;
/*!40000 ALTER TABLE `common_fileregistry` DISABLE KEYS */;
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
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `common_option`
--

LOCK TABLES `common_option` WRITE;
/*!40000 ALTER TABLE `common_option` DISABLE KEYS */;
INSERT INTO `common_option` VALUES (3,'@system_mailServer','smtp.126.com:25','SystemInternalConfig'),(4,'@system_mailAccount','rangh@126.com','SystemInternalConfig'),(5,'@system_mailPassword','hRangh@13924','SystemInternalConfig'),(6,'@application_logo','http://localhost:8888/api/system/','SystemInternalConfig'),(13,'@application_name','magicCenter','SystemInternalConfig'),(14,'@application_description','rangh\'s magicCenter','SystemInternalConfig'),(15,'@application_domain','muidea.com','SystemInternalConfig'),(16,'@system_uploadPath','upload','SystemInternalConfig'),(17,'@system_staticPath','./static/','SystemInternalConfig'),(47,'@application_startupData','startup_TimeStamp:2018-08-10 13:15:59','SystemInternalConfig');
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
INSERT INTO `common_resource` VALUES (1,1,'magicDRP','auto setup catalog description','catalog','2018-09-06 19:27:16',0),(2,2,'order','订单信息','catalog','2018-09-06 19:27:16',0),(3,3,'history','历史订单','catalog','2018-09-06 19:27:16',0),(4,4,'report','订单报表','catalog','2018-09-06 19:27:16',0),(5,5,'bill','账单信息','catalog','2018-09-06 19:27:16',0),(6,6,'partner','代理商信息','catalog','2018-09-06 19:27:16',0),(7,7,'product','产品信息','catalog','2018-09-06 19:27:16',0),(8,8,'offline','下线产品','catalog','2018-09-06 19:27:16',0),(9,9,'realtime','实时订单','catalog','2018-09-06 19:27:16',0),(10,10,'district','区域信息','catalog','2018-09-06 19:27:16',0),(11,11,'realtime','实时账单','catalog','2018-09-06 19:29:17',1),(12,12,'history','历史账单','catalog','2018-09-06 19:29:33',1),(13,13,'report','账单报表','catalog','2018-09-06 19:29:52',1),(14,14,'setting','系统设置','catalog','2018-09-06 19:30:24',1);
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
) ENGINE=InnoDB AUTO_INCREMENT=23862 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `common_resource_relative`
--

LOCK TABLES `common_resource_relative` WRITE;
/*!40000 ALTER TABLE `common_resource_relative` DISABLE KEYS */;
INSERT INTO `common_resource_relative` VALUES (23849,2,1),(23850,3,2),(23851,4,2),(23852,5,1),(23853,6,1),(23854,7,1),(23855,8,7),(23856,9,2),(23857,10,1),(23858,11,5),(23859,12,5),(23860,13,5),(23861,14,1);
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
INSERT INTO `content_catalog` VALUES (1,'magicDRP','auto setup catalog description','2018-09-06 19:27:16',0),(2,'order','订单信息','2018-09-06 19:27:16',0),(3,'history','历史订单','2018-09-06 19:27:16',0),(4,'report','订单报表','2018-09-06 19:27:16',0),(5,'bill','账单信息','2018-09-06 19:27:16',0),(6,'partner','代理商信息','2018-09-06 19:27:16',0),(7,'product','产品信息','2018-09-06 19:27:16',0),(8,'offline','下线产品','2018-09-06 19:27:16',0),(9,'realtime','实时订单','2018-09-06 19:27:16',0),(10,'district','区域信息','2018-09-06 19:27:16',0),(11,'realtime','实时账单','2018-09-06 19:29:17',1),(12,'history','历史账单','2018-09-06 19:29:33',1),(13,'report','账单报表','2018-09-06 19:29:52',1),(14,'setting','系统设置','2018-09-06 19:30:24',1);
/*!40000 ALTER TABLE `content_catalog` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `content_comment`
--

DROP TABLE IF EXISTS `content_comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `content_comment` (
  `id` int(11) unsigned NOT NULL,
  `subject` text NOT NULL,
  `content` text NOT NULL,
  `createdate` datetime NOT NULL,
  `creater` int(11) NOT NULL,
  `flag` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `content_comment`
--

LOCK TABLES `content_comment` WRITE;
/*!40000 ALTER TABLE `content_comment` DISABLE KEYS */;
/*!40000 ALTER TABLE `content_comment` ENABLE KEYS */;
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
/*!40000 ALTER TABLE `content_media` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `endpoint_registry`
--

DROP TABLE IF EXISTS `endpoint_registry`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `endpoint_registry` (
  `eid` int(11) NOT NULL AUTO_INCREMENT,
  `id` text NOT NULL,
  `name` text NOT NULL,
  `description` text NOT NULL,
  `user` text NOT NULL,
  `status` int(11) NOT NULL,
  `authToken` text,
  PRIMARY KEY (`eid`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `endpoint_registry`
--

LOCK TABLES `endpoint_registry` WRITE;
/*!40000 ALTER TABLE `endpoint_registry` DISABLE KEYS */;
INSERT INTO `endpoint_registry` VALUES (12,'0a1ef1fe-7ff2-46bc-a5cd-32b8b35bbe2c','magicDRP','','',0,'7e8oGovuWOUnzqiCXOspqtnTTYTFZC65');
/*!40000 ALTER TABLE `endpoint_registry` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-09-06 19:31:36
