-- MySQL dump 10.13  Distrib 8.0.32, for Linux (x86_64)
--
-- Host: localhost    Database: goblog
-- ------------------------------------------------------
-- Server version	8.0.32

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `demos`
--

DROP TABLE IF EXISTS `demos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `demos` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `fname` longtext,
  `lname` longtext,
  `email` longtext,
  `password` longtext,
  `number` bigint DEFAULT NULL,
  `profile` longtext,
  `is_delete` varchar(45) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_demos_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `demos`
--

LOCK TABLES `demos` WRITE;
/*!40000 ALTER TABLE `demos` DISABLE KEYS */;
INSERT INTO `demos` VALUES (26,'2023-05-13 15:58:56.049','2023-05-24 08:58:40.646',NULL,'ram','siya','ram@gmail.com','$2a$14$/5U/gbq8dER8YhegjnXTZOfP2opM.oLqEfKuyCo03y.iz/XYpAFl.',1234567890,'/images/profile-img.png','0'),(27,'2023-05-13 15:59:12.355','2023-05-13 15:59:12.355',NULL,'sita','ram','sita@gmail.com','$2a$14$pfSrrne.e0z714HoRwqYUO0rAn3/wtmI8xxssAjK40wmH0SxuJZCW',1234567890,'/images/about-removebg-preview.png',NULL),(28,'2023-05-13 15:59:44.982','2023-05-13 15:59:44.982',NULL,'lakhan','lakhan','lakhan@gmail.com','$2a$14$4uc4yrWozZk2WL0TytBfdudtJZzTOBkHT68Gyz28Yjba545Gan8ae',1234567890,'/images/banner.jpg',NULL),(29,'2023-05-13 16:00:10.103','2023-05-13 16:00:10.103',NULL,'Harmil','sanghavi','harmil@gmail.com','$2a$14$wp4pouusyHUMnfxO6PcUVeaEQ6PgeMKJ1unodpkgbkNRd9jpkL1S2',1234567890,'/images/dark_sepia_nature_background_210839.jpg',NULL),(30,'2023-05-13 16:00:29.133','2023-05-13 16:00:29.133',NULL,'vivek','valand','vivek@gmail.com','$2a$14$3sEOtFNBRmT68vQGnop9leYAe/H2OtWsuI2gqgCPmPi11dPJXkRZC',1234567890,'/images/images.jpeg',NULL),(31,'2023-05-15 11:22:32.683','2023-05-15 11:22:32.683',NULL,'harsh','patel','harsh@gmail.com','$2a$14$ACaN1Z/1mNj1pizV0.6PEOvKmkihE341.aozqhl4cMkFZGW8oV92C',1234567890,'/images/nature-3125912_1280.jpg','0'),(32,'2023-05-15 11:23:03.398','2023-05-15 15:19:12.194',NULL,'viren','rathod','viren@gmail.com','$2a$14$Hf5haWQyw4m6ZV4GHYWHDeHhz7yGDYAm7F0AzapWfFeJmwy.U.hca',1234567890,'/images/flowers-276014_1280.jpg','0'),(33,'2023-05-15 11:23:21.461','2023-05-16 10:05:45.182',NULL,'mit','mit','mit@gmail.com','$2a$14$3qUAStUs2C76iyOV4xb.h.Kx2S.0eNho/w3fF7m/TO8qWBSSC3ODm',1234567890,'/images/Polygon.png','0'),(34,'2023-05-15 14:11:47.730','2023-05-15 14:11:47.730',NULL,'vikram','vikram','vikram@gmail.com','$2a$14$T.wTqTDq0275tbzLK5InVOM/TUaL88dbACmT/Cyx2fwXoIXnuI6Aa',1234567890,'/images/about-removebg-preview.png','0'),(35,'2023-05-15 14:12:07.548','2023-05-15 14:12:07.548',NULL,'isha','desai','isha@gmail.com','$2a$14$8/w0qUViOXRpGmRC9X82Pua6BdbhT06YH2nkNyGLef5ipqlrQfltC',1234567890,'/images/dark_sepia_nature_background_210839.jpg','0'),(36,'2023-05-15 14:12:27.073','2023-05-15 14:12:27.073',NULL,'riya','riya','riya@gmail.com','$2a$14$RnIfJHRlLLWXCXI3s7TXeeZB.jrxPj2R1nN80xKjQV4Ex7YovVCK.',1234567890,'/images/images.jpeg','0'),(37,'2023-05-15 14:12:55.223','2023-05-15 14:12:55.223',NULL,'peter','parkar','peter@gmail.com','$2a$14$5y6nyFinsXGQuu7CNjTDb.k9vVKzMoKdmAJU5g5o4VwHpbR8LcQpK',1234567890,'/images/nature-3125912_1280.jpg','0'),(38,'2023-05-15 14:13:32.633','2023-05-16 09:52:27.206',NULL,'john','john','john@gmail.com','$2a$14$aoUDhMw6NA.4/2H4k.Q8K.x4U65fDKl2MS7ZDK1RMpic10rLPeV32',1234567890,'/images/Icon_Like.png','0'),(49,'2023-05-16 09:57:13.585','2023-05-24 10:30:49.066',NULL,' jay ','patel',' jay@gmail.com ','$2a$14$itaMNDsKLMjCtRlsxS1oeOiBuAYJ/.xK3qMO64ePD2FLV2R4Sxa.u',1234567890,'/images/Polygon.png','0'),(50,'2023-05-16 10:06:34.473','2023-05-16 10:06:34.473',NULL,' raj','raj','raj@gmail.com','$2a$14$L1CbCoR8CVXJwnOSZUhS7uMgbGwh3AfH.eRGPE5/IDjiyaa28tS2e',1234567890,'/images/Rectangle.png','0'),(51,'2023-05-24 10:38:22.051','2023-05-24 10:41:07.373',NULL,'domanic ','torento','domanic@gmail.com','$2a$14$.UBdcE5eeNutYDPG66BhCuMxdMVpLiFjtCwBUN35YmzcJZjN7HMW2',0,'/images/Path(2).png','1'),(52,'2023-05-24 10:44:29.146','2023-05-24 10:45:52.474',NULL,' domanic','torento',' domanic@gmail.com','$2a$14$FPEis3m1WTA8LJ1Y3.gvRedzGHXlLZoP28/rIW0mYgc5kixKZOWvi',1234567890,'/images/Oval.png','1'),(53,'2023-05-27 16:38:23.352','2023-05-27 16:38:50.248',NULL,' demo','demo','demo@gmail.com','$2a$14$zQS2lg9hSdtVXJYtHPjKwe5AEholcyBLM4asanOSyQg1B6OwqPDzS',1234567890,'/images/Blog.png','1');
/*!40000 ALTER TABLE `demos` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-08-07 14:46:23
