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
INSERT INTO `account_group` VALUES (1,'系统管理组',0,'系统管理组描述信息',1),(2,'普通用户组',0,'普通用户组描述信息',1),(3,'magicBlog用户组',2,'MagicBlog用户组描述信息',0),(4,'magicShare用户组',2,'magicShare用户分组描述',0);
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
INSERT INTO `account_user` VALUES (1,'admin','123','admin@muidea.com','1',0,'2018-03-20 00:00:00',1),(2,'test','123','test@126.com','4',0,'2018-07-08 21:13:12',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=137 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `common_fileregistry`
--

LOCK TABLES `common_fileregistry` WRITE;
/*!40000 ALTER TABLE `common_fileregistry` DISABLE KEYS */;
INSERT INTO `common_fileregistry` VALUES (74,'arrifoxwj3arcnkotqubjzkpczsjnsxo','7谢瑶瑶.pdf','static/upload/arrifoxwj3arcnkotqubjzkpczsjnsxo/7谢瑶瑶.pdf','2018-08-11 13:17:01',0),(75,'rueis9lp09d7zrxim6ppa6vxkvohzyun','logoko (1).png','static/upload/rueis9lp09d7zrxim6ppa6vxkvohzyun/logoko (1).png','2018-08-11 13:50:28',0),(76,'csa7rqmrfda14h5djknhsfr9aglv1vzp','dKbkpPXKfvZzWCM.png','static/upload/csa7rqmrfda14h5djknhsfr9aglv1vzp/dKbkpPXKfvZzWCM.png','2018-08-11 14:47:14',0),(77,'ev3kbjjrijqodxugpksxhuikhhujql32','dKbkpPXKfvZzWCM.png','static/upload/ev3kbjjrijqodxugpksxhuikhhujql32/dKbkpPXKfvZzWCM.png','2018-08-11 14:49:56',0),(78,'wkkdqkm5nawxznffbq6be4zptjujiqzs','dKbkpPXKfvZzWCM.png','static/upload/wkkdqkm5nawxznffbq6be4zptjujiqzs/dKbkpPXKfvZzWCM.png','2018-08-11 14:50:14',0),(79,'r01azpkqddcmtbeccri0oownvbgjszvp','dKbkpPXKfvZzWCM.png','static/upload/r01azpkqddcmtbeccri0oownvbgjszvp/dKbkpPXKfvZzWCM.png','2018-08-11 14:51:08',0),(80,'azzqb9pn6vhohquc4ozmrtyuqw83th3n','dKbkpPXKfvZzWCM.png','static/upload/azzqb9pn6vhohquc4ozmrtyuqw83th3n/dKbkpPXKfvZzWCM.png','2018-08-11 14:52:04',0),(81,'8d6m185lfyztyzvqc18m6kzmbniismqq','dKbkpPXKfvZzWCM.png','static/upload/8d6m185lfyztyzvqc18m6kzmbniismqq/dKbkpPXKfvZzWCM.png','2018-08-11 18:25:40',0),(82,'6oimqyxsu6o52bwpwp0qngnldy5z1alz','dKbkpPXKfvZzWCM.png','static/upload/6oimqyxsu6o52bwpwp0qngnldy5z1alz/dKbkpPXKfvZzWCM.png','2018-08-11 19:09:25',0),(83,'s8ppvszgocubp5whwtxoui890autecep','dKbkpPXKfvZzWCM.png','static/upload/s8ppvszgocubp5whwtxoui890autecep/dKbkpPXKfvZzWCM.png','2018-08-11 19:12:05',0),(84,'sjrghru7sgqjk6lfamc91jld2716fuvi','dKbkpPXKfvZzWCM.png','static/upload/sjrghru7sgqjk6lfamc91jld2716fuvi/dKbkpPXKfvZzWCM.png','2018-08-11 22:49:37',0),(85,'izaef5nk2wvm5yhfw6mbkuimdqs9fcai','备案幕布.png','static/upload/izaef5nk2wvm5yhfw6mbkuimdqs9fcai/备案幕布.png','2018-08-11 22:50:21',0),(86,'ltslkfqq0h09nayx5umwqnklv3lfgxim','域名证书.jpg','static/upload/ltslkfqq0h09nayx5umwqnklv3lfgxim/域名证书.jpg','2018-08-11 22:50:40',0),(90,'s6517swv0q1gethxaphahycp3ow9kp4e','th.jpeg','static/upload/s6517swv0q1gethxaphahycp3ow9kp4e/th.jpeg','2018-08-13 15:59:21',0),(91,'zhkxtplmydqhvutivpjxecb6oh82pcw7','zhegexiaritouqingliangjingmeibizhi_488270_3.jpg','static/upload/zhkxtplmydqhvutivpjxecb6oh82pcw7/zhegexiaritouqingliangjingmeibizhi_488270_3.jpg','2018-08-13 16:08:13',0),(128,'ltqmmcum9eyyabf8ptsklac6nlchonra','休闲夹克衫.jpg','static/upload/ltqmmcum9eyyabf8ptsklac6nlchonra/休闲夹克衫.jpg','2018-08-18 15:53:14',0),(129,'eb16r6tuyluim5gkyolbqxk8dymbdyw8','修身衬衣.jpeg','static/upload/eb16r6tuyluim5gkyolbqxk8dymbdyw8/修身衬衣.jpeg','2018-08-18 15:53:28',0),(130,'yctavmpnc6obsyk6kn6unbvtgwfd8l4h','商务小西装.jpeg','static/upload/yctavmpnc6obsyk6kn6unbvtgwfd8l4h/商务小西装.jpeg','2018-08-18 15:53:44',0),(131,'eicdfe7bsvm67ugmgadkuad6612kusiz','毛呢大衣.jpg','static/upload/eicdfe7bsvm67ugmgadkuad6612kusiz/毛呢大衣.jpg','2018-08-18 15:53:59',0),(132,'epyfwz8laiqo7msyzjupzxmmb9xyn0fq','get-pip.py','static/upload/epyfwz8laiqo7msyzjupzxmmb9xyn0fq/get-pip.py','2018-08-21 16:25:49',0),(133,'syulsbghmctuflmbm2snualgjmmxkkm5','get-pip.py','static/upload/syulsbghmctuflmbm2snualgjmmxkkm5/get-pip.py','2018-08-21 16:27:36',0),(134,'fjhrtg21ajulfo54sb37ppbfkhcqldhk','get-pip.py','static/upload/fjhrtg21ajulfo54sb37ppbfkhcqldhk/get-pip.py','2018-08-21 16:29:06',0),(135,'yp3f9a8j877vib4msvzahbue0on02fiv','get-pip.py','static/upload/yp3f9a8j877vib4msvzahbue0on02fiv/get-pip.py','2018-08-21 16:30:45',0),(136,'lrcujyqrcp7cxnm6rkjxfzvvfxagojws','17k基于kubernetes容器云平台实践v1.1.pptx','static/upload/lrcujyqrcp7cxnm6rkjxfzvvfxagojws/17k基于kubernetes容器云平台实践v1.1.pptx','2018-08-21 16:31:31',0);
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
INSERT INTO `common_resource` VALUES (7,2,'magicBlog','magicBlog 分组','catalog','2018-04-23 23:10:50',1),(9,1,'测试文章','这里一些测试内容','article','2018-05-01 22:24:29',1),(15,3,'About','个人介绍','article','2018-05-05 18:14:32',1),(16,4,'Contact','交流信息','article','2018-05-05 22:21:36',1),(17,8,'Catalog','分类信息','catalog','2018-05-05 22:26:21',1),(18,9,'Index','主页信息','catalog','2018-05-05 22:30:08',1),(19,5,'404','404页面','article','2018-05-05 22:31:11',1),(20,10,'技术文章','这是技术文章的描述','catalog','2018-05-06 13:41:35',1),(42,10,'测试内容,这是一篇测试文章','rules: [ { required: true }, ],','article','2018-06-09 09:42:13',1),(75,26,'magicShare','auto update catalog description','catalog','2018-07-09 19:31:07',1),(87,32,'shareCatalog','共享文件','catalog','2018-07-15 18:02:57',1),(88,33,'privacyCatalog','私有文件','catalog','2018-07-16 16:09:11',1),(122,54,'magicDRP','magicDRP','catalog','2018-08-10 21:46:23',1),(123,55,'district','区域定义','catalog','2018-08-10 21:46:59',1),(124,56,'partner','partner代理商','catalog','2018-08-10 21:47:22',1),(125,57,'product','产品信息','catalog','2018-08-10 21:47:47',1),(131,58,'安徽省','340000','catalog','2018-08-11 14:11:07',1),(132,59,'六安市','341500','catalog','2018-08-11 14:11:38',1),(133,60,'霍邱县','341522','catalog','2018-08-11 14:12:04',1),(134,61,'霍山县','341525','catalog','2018-08-11 14:12:32',1),(135,62,'金安区','341502','catalog','2018-08-11 14:12:52',1),(136,63,'澳门特别行政区','820000','catalog','2018-08-11 14:13:42',1),(137,64,'澳门半岛','820100','catalog','2018-08-11 14:14:02',1),(138,65,'离岛','820200','catalog','2018-08-11 14:14:21',1),(139,66,'浙江省','330000','catalog','2018-08-11 14:15:05',1),(140,67,'杭州市','330100','catalog','2018-08-11 14:15:26',1),(141,68,'滨江区','330108','catalog','2018-08-11 14:15:46',1),(142,69,'淳安县','330127','catalog','2018-08-11 14:16:08',1),(143,70,'富阳区','330183','catalog','2018-08-11 14:16:33',1),(144,71,'拱墅区','330105','catalog','2018-08-11 14:16:57',1),(262,41,'休闲夹克衫.jpg','{\"id\":0,\"name\":\"休闲夹克衫.jpg\",\"description\":\"休闲夹克衫\",\"model\":\"pro_0001\",\"image\":\"ltqmmcum9eyyabf8ptsklac6nlchonra\",\"district\":[]}','media','2018-08-18 15:53:24',0),(263,42,'修身衬衣.jpeg','{\"id\":0,\"name\":\"修身衬衣.jpeg\",\"description\":\"修身衬衣\",\"model\":\"pro_0002\",\"image\":\"eb16r6tuyluim5gkyolbqxk8dymbdyw8\",\"district\":[]}','media','2018-08-18 15:53:40',0),(264,43,'商务小西装.jpeg','{\"id\":0,\"name\":\"商务小西装.jpeg\",\"description\":\"商务小西装\",\"model\":\"pro_0003\",\"image\":\"yctavmpnc6obsyk6kn6unbvtgwfd8l4h\",\"district\":[]}','media','2018-08-18 15:53:55',0),(265,44,'毛呢大衣.jpg','{\"id\":0,\"name\":\"毛呢大衣.jpg\",\"description\":\"毛呢大衣\",\"model\":\"pro_0004\",\"image\":\"eicdfe7bsvm67ugmgadkuad6612kusiz\",\"district\":[]}','media','2018-08-18 15:54:06',0),(266,11,'11111111111','{\"id\":11,\"name\":\"测试\",\"telephone\":\"11111111111\",\"wechat\":\"1111\",\"referee\":\"\",\"district\":{\"id\":58,\"label\":\"安徽省\",\"value\":\"340000\",\"children\":[{\"id\":59,\"label\":\"六安市\",\"value\":\"341500\",\"children\":[{\"id\":62,\"label\":\"金安区\",\"value\":\"341502\",\"children\":[]}]}]},\"product\":[{\"id\":41,\"name\":\"休闲夹克衫.jpg\",\"description\":\"休闲夹克衫\",\"model\":\"pro_0001\",\"image\":\"ltqmmcum9eyyabf8ptsklac6nlchonra\"}]}','article','2018-08-21 16:58:09',0),(267,1,'userCode','oG-Py1Lq4QaQ4fxI6AqMzfabTT_g','comment','2018-08-21 16:58:09',0),(268,12,'11111111111','{\"id\":0,\"name\":\"测试\",\"telephone\":\"11111111111\",\"wechat\":\"12121\",\"referee\":\"\",\"district\":{\"id\":58,\"label\":\"安徽省\",\"value\":\"340000\",\"children\":[{\"id\":59,\"label\":\"六安市\",\"value\":\"341500\",\"children\":[{\"id\":62,\"label\":\"金安区\",\"value\":\"341502\",\"children\":[]}]}]},\"product\":[]}','article','2018-08-22 09:24:29',0),(269,2,'userCode','aabbad','comment','2018-08-22 09:24:29',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=884 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `common_resource_relative`
--

LOCK TABLES `common_resource_relative` WRITE;
/*!40000 ALTER TABLE `common_resource_relative` DISABLE KEYS */;
INSERT INTO `common_resource_relative` VALUES (52,17,7),(55,18,7),(63,20,17),(64,20,18),(80,9,17),(81,9,18),(154,15,7),(155,16,7),(156,19,7),(248,88,75),(261,87,75),(268,42,17),(269,42,18),(304,123,122),(305,124,122),(306,125,122),(312,131,123),(313,132,131),(314,133,132),(315,134,132),(316,135,132),(317,136,123),(318,137,136),(319,138,136),(320,139,123),(321,140,139),(322,141,140),(323,142,140),(324,143,140),(325,144,140),(747,262,125),(748,263,125),(749,264,125),(750,265,125),(873,267,266),(874,266,124),(875,266,131),(876,266,132),(877,266,135),(878,266,262),(879,268,124),(880,268,131),(881,268,132),(882,268,135),(883,269,268);
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
INSERT INTO `content_article` VALUES (1,'测试文章','\n不少朋友都知道我在“[++极客时间++](https://time.geekbang.org/)”上开了一个收费专栏，这个专栏会开设大约一年的时间，一共会发布104篇文章。现在，我在上面以每周两篇文章的频率已发布了27篇文章了，也就是差不多两个半月的时间。新的一年开始了，写专栏这个事对我来说是第一次，在这个过程中有一些感想，所以，我想在这里说一下这些感受和一些相关的故事，算是一个记录，也算是对我专栏的正式介绍，还希望能得到、大家的喜欢和指点。（当然，CoolShell这边还是会持续更新的）\n\n测试内容\n\n​\n\n​\n\n看看效果\n\n- **A\n- **B\n','2018-05-08 23:33:43',1),(3,'About','## 个人介绍\n\n我叫黄冬朋，互联网技术爱好者，Gopher!\n\n2016年以前专注于C++跨平台服务器后台应用系统开发，擅长通讯服务器，数据管理软件架构开发。\n\n2014年开始接触Golang，曾经也是Python的爱好者和推广者，自从接触到Golang后，就被它的设计哲理所吸引。 开始各种场合推荐Go，并逐步开始使用Go进行系统开发。\n\n欢迎跟大家交流Kubernetes，Docker，Cloud。\n\n希望能跟大家多多交流，微信：21883911\n','2018-06-10 21:24:13',1),(4,'Contact','## 站点介绍\n\n记录建设云基础平台过程中经历的心路历程，欢迎与大家一起相互交流。\n\n本站使用的技术栈：Docker + Golang + React + MySQL\n\n前后端都是由我个人纯手工打造，引用了部分开源项目(后面单独说明)。本人C++后台开发出身，部分内容可能会姿势不对，欢迎大家拍砖！\n\n也欢迎与大家相互交流，分享心得，也诚邀美工和前端的朋友一起合作，欢迎联系！\n\n交换站点链接，请加我微信并说明！\n\n1、为什么要建本站？\n\n为了实现多年夙愿，也为了对基础平台进行功能验证。\n\n2、站点代码开源么？\n\n开源，GitHub地址: [magicBlog](https://github.com/muidea/magicBlog)\n\n3、本站引用到的项目\n\nAnt Design\n','2018-06-10 21:30:28',1),(5,'404','# **找不到内容了！**\n\n**如果你喜欢本站，欢迎交流！**\n','2018-06-10 21:32:20',1),(10,'测试内容,这是一篇测试文章','rules: [ { required: true }, ],\n\n这里只是看看效果，\n![](http://www.linuxdaxue.com/wp-content/uploads/2017/07/image-1.png)\n\n看看效果怎样啊\n','2018-07-27 19:25:20',1),(11,'11111111111','{\"id\":11,\"name\":\"测试\",\"telephone\":\"11111111111\",\"wechat\":\"1111\",\"referee\":\"\",\"district\":{\"id\":58,\"label\":\"安徽省\",\"value\":\"340000\",\"children\":[{\"id\":59,\"label\":\"六安市\",\"value\":\"341500\",\"children\":[{\"id\":62,\"label\":\"金安区\",\"value\":\"341502\",\"children\":[]}]}]},\"product\":[{\"id\":41,\"name\":\"休闲夹克衫.jpg\",\"description\":\"休闲夹克衫\",\"model\":\"pro_0001\",\"image\":\"ltqmmcum9eyyabf8ptsklac6nlchonra\"}]}','2018-08-21 16:58:17',0),(12,'11111111111','{\"id\":0,\"name\":\"测试\",\"telephone\":\"11111111111\",\"wechat\":\"12121\",\"referee\":\"\",\"district\":{\"id\":58,\"label\":\"安徽省\",\"value\":\"340000\",\"children\":[{\"id\":59,\"label\":\"六安市\",\"value\":\"341500\",\"children\":[{\"id\":62,\"label\":\"金安区\",\"value\":\"341502\",\"children\":[]}]}]},\"product\":[]}','2018-08-22 09:24:29',0);
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
INSERT INTO `content_catalog` VALUES (2,'magicBlog','MagicBlog catalog description','2018-06-10 21:33:15',1),(8,'Catalog','Catalog分类','2018-05-05 22:26:21',1),(9,'Index','Index','2018-05-05 22:30:08',1),(10,'技术文章','技术文章','2018-05-06 13:41:35',1),(26,'magicShare','auto update catalog description','2018-07-19 10:19:42',1),(32,'shareCatalog','共享文件','2018-07-19 10:19:42',1),(33,'privacyCatalog','私有文件','2018-07-17 11:17:27',1),(54,'magicDRP','magicDRP','2018-08-10 21:46:23',1),(55,'district','区域定义','2018-08-10 21:46:59',1),(56,'partner','partner代理商','2018-08-10 21:47:22',1),(57,'product','产品信息','2018-08-10 21:47:47',1),(58,'安徽省','340000','2018-08-11 14:11:07',1),(59,'六安市','341500','2018-08-11 14:11:38',1),(60,'霍邱县','341522','2018-08-11 14:12:04',1),(61,'霍山县','341525','2018-08-11 14:12:32',1),(62,'金安区','341502','2018-08-11 14:12:52',1),(63,'澳门特别行政区','820000','2018-08-11 14:13:42',1),(64,'澳门半岛','820100','2018-08-11 14:14:02',1),(65,'离岛','820200','2018-08-11 14:14:21',1),(66,'浙江省','330000','2018-08-11 14:15:05',1),(67,'杭州市','330100','2018-08-11 14:15:26',1),(68,'滨江区','330108','2018-08-11 14:15:46',1),(69,'淳安县','330127','2018-08-11 14:16:08',1),(70,'富阳区','330183','2018-08-11 14:16:33',1),(71,'拱墅区','330105','2018-08-11 14:16:57',1);
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
INSERT INTO `content_comment` VALUES (1,'userCode','oG-Py1Lq4QaQ4fxI6AqMzfabTT_g','2018-08-21 16:58:09',0,0),(2,'userCode','aabbad','2018-08-22 09:24:29',0,0);
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
INSERT INTO `content_media` VALUES (41,'休闲夹克衫.jpg','{\"id\":0,\"name\":\"休闲夹克衫.jpg\",\"description\":\"休闲夹克衫\",\"model\":\"pro_0001\",\"image\":\"ltqmmcum9eyyabf8ptsklac6nlchonra\",\"district\":[]}','ltqmmcum9eyyabf8ptsklac6nlchonra','2018-08-18 15:53:24',0,3650),(42,'修身衬衣.jpeg','{\"id\":0,\"name\":\"修身衬衣.jpeg\",\"description\":\"修身衬衣\",\"model\":\"pro_0002\",\"image\":\"eb16r6tuyluim5gkyolbqxk8dymbdyw8\",\"district\":[]}','eb16r6tuyluim5gkyolbqxk8dymbdyw8','2018-08-18 15:53:40',0,3650),(43,'商务小西装.jpeg','{\"id\":0,\"name\":\"商务小西装.jpeg\",\"description\":\"商务小西装\",\"model\":\"pro_0003\",\"image\":\"yctavmpnc6obsyk6kn6unbvtgwfd8l4h\",\"district\":[]}','yctavmpnc6obsyk6kn6unbvtgwfd8l4h','2018-08-18 15:53:55',0,3650),(44,'毛呢大衣.jpg','{\"id\":0,\"name\":\"毛呢大衣.jpg\",\"description\":\"毛呢大衣\",\"model\":\"pro_0004\",\"image\":\"eicdfe7bsvm67ugmgadkuad6612kusiz\",\"district\":[]}','eicdfe7bsvm67ugmgadkuad6612kusiz','2018-08-18 15:54:06',0,3650);
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
INSERT INTO `endpoint_registry` VALUES (5,'f0e078a8-6de8-4273-88a4-dccef60ff88f','magicBlog','magicBlog是一个博客应用','1',0,'yTtWiuuoGifPVfcK5Mf4mdu8mGl78E3y'),(6,'b92c3028-cadb-43e8-8fcd-576d8ffcfcc5','magicShare','magicShare','2',0,'ADYiib9Ss3roQ5lNhhv601wmx87f5Hsq'),(10,'8a30313b-50ed-4a3a-b52b-e0d3e000d400','magicProtal','门户','',0,'ZA7gLdcpeabhFFQXWH7N3HSXHxK2LnAU'),(11,'c88d1ab9-ec65-4c88-9a8b-9218e0826744','magicMall','','',0,'xQJQ4UwGy3IHFponoLMSa5HGw5J3R6vK'),(12,'0a1ef1fe-7ff2-46bc-a5cd-32b8b35bbe2c','magicDRP','','',0,'7e8oGovuWOUnzqiCXOspqtnTTYTFZC65');
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

-- Dump completed on 2018-08-22  9:35:06
