-- MySQL dump 10.13  Distrib 8.0.19, for Win64 (x86_64)
--
-- Host: localhost    Database: meono
-- ------------------------------------------------------
-- Server version	11.5.2-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `game_tb`
--

DROP TABLE IF EXISTS `game_tb`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `game_tb` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `status` varchar(50) DEFAULT NULL,
  `playuser` varchar(50) DEFAULT NULL,
  `bai` varchar(50) DEFAULT NULL,
  `bobai` varchar(1000) DEFAULT NULL,
  `rote` int(11) DEFAULT NULL,
  `gs` varchar(100) DEFAULT NULL,
  `gd` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `game_tb`
--

LOCK TABLES `game_tb` WRITE;
/*!40000 ALTER TABLE `game_tb` DISABLE KEYS */;
INSERT INTO `game_tb` VALUES (1,'p','6','6','6,6,3,4,6,2,7,7,3,1,6,4,3,4,7,7,7,3,5',1,'4','');
/*!40000 ALTER TABLE `game_tb` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `log_tb`
--

DROP TABLE IF EXISTS `log_tb`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `log_tb` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `mm` varchar(1000) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=385 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `log_tb`
--

LOCK TABLES `log_tb` WRITE;
/*!40000 ALTER TABLE `log_tb` DISABLE KEYS */;
INSERT INTO `log_tb` VALUES (374,'bat dau game'),(375,'an di dau tien'),(376,'an: rut bai'),(377,'cong: rut bai'),(378,'an: rut bai'),(379,'cong: rut bai'),(380,'an: rut bai'),(381,'cong: da xao bai'),(382,'cong: danh bai reverse'),(383,'an: danh bai reverse'),(384,'cong: da xao bai');
/*!40000 ALTER TABLE `log_tb` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_tb`
--

DROP TABLE IF EXISTS `user_tb`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_tb` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT NULL,
  `status` varchar(50) DEFAULT NULL,
  `arr` varchar(300) DEFAULT NULL,
  `bom` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_tb`
--

LOCK TABLES `user_tb` WRITE;
/*!40000 ALTER TABLE `user_tb` DISABLE KEYS */;
INSERT INTO `user_tb` VALUES (6,'cong','p','2,4,4,5','0'),(7,'an','p','2,3,3,4,5,5,7','0'),(8,'','','',''),(9,'','','',''),(10,'','','',''),(11,'','','',''),(12,'','','',''),(13,'','','',''),(14,'','','',''),(15,'','','',''),(16,'','','',''),(17,'','','',''),(18,'','','',''),(19,'','','',''),(20,'','','',''),(21,'','','',''),(22,'','','',''),(23,'','','',''),(24,'','','','');
/*!40000 ALTER TABLE `user_tb` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'meono'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-09-23 16:56:57
