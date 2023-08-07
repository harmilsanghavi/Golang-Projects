-- MySQL dump 10.13  Distrib 8.0.32, for Linux (x86_64)
--
-- Host: localhost    Database: hospital_system
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
-- Table structure for table `appoinment_books`
--

DROP TABLE IF EXISTS `appoinment_books`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `appoinment_books` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `first_name` longtext,
  `last_name` longtext,
  `address` longtext,
  `number` bigint DEFAULT NULL,
  `gender` longtext,
  `age` bigint DEFAULT NULL,
  `day` longtext,
  `apooinment_time` longtext,
  `user_id` bigint unsigned DEFAULT NULL,
  `role_name` longtext,
  `doctor_id` bigint unsigned DEFAULT NULL,
  `status` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_appoinment_books_deleted_at` (`deleted_at`),
  KEY `fk_appoinment_books_users` (`user_id`),
  CONSTRAINT `fk_appoinment_books_users` FOREIGN KEY (`user_id`) REFERENCES `signups` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `appoinment_books`
--

LOCK TABLES `appoinment_books` WRITE;
/*!40000 ALTER TABLE `appoinment_books` DISABLE KEYS */;
INSERT INTO `appoinment_books` VALUES (14,'2023-06-27 14:31:53.974','2023-07-13 07:37:12.090',NULL,'harmil','sanghavi','weuiryewruieyhruiw',7894561230,'male',20,'2023-06-28','1',1,'patient',8,'complete'),(15,'2023-06-27 14:32:31.057','2023-06-27 14:32:31.057','2023-06-27 14:32:47.375','vatsal','sanghavi','eryuweuirewyr',4567891230,'male',22,'2023-06-28','3',3,'patient',8,'complete'),(16,'2023-06-27 15:39:30.633','2023-06-27 15:39:30.633',NULL,'harmil','sanghavi','weryuiweurwy',7894561230,'male',21,'2023-06-26','1',1,'patient',8,'complete'),(17,'2023-06-27 15:40:58.257','2023-07-13 06:13:22.671',NULL,'vatsal','sanghavi','uiuou88oui',7894560123,'male',25,'2023-06-26','4',3,'patient',8,'complete'),(18,'2023-07-03 14:17:49.859','2023-07-03 14:17:49.859','2023-07-03 14:32:04.827','riya','shah','sdeuiyeieyhfriu hv',4567891230,'female',20,'2023-07-03','8',17,'patient',19,'complete'),(19,'2023-07-03 14:24:14.372','2023-07-03 14:24:14.372',NULL,'riya','shah','eijhurirhteruit',4561237890,'female',20,'2023-07-04','2',17,'patient',9,'complete'),(20,'2023-07-03 14:31:05.215','2023-07-03 14:31:05.215',NULL,'riya','shah','deutyhrutyh',4561327890,'female',20,'2023-07-06','3',17,'patient',10,'complete'),(21,'2023-07-03 14:39:03.692','2023-07-03 14:39:03.692',NULL,'riya','shah','eryuweireyrui',4561237890,'female',20,'2023-07-04','4',17,'patient',19,'complete'),(26,'2023-07-07 08:02:54.246','2023-07-07 08:02:54.246',NULL,'harmil','sanghavi','efrere',1234567890,'male',25,'2023-07-08','1',1,'patient',19,'complete'),(27,'2023-07-07 09:14:29.363','2023-07-07 09:14:29.363',NULL,'harmil','sanghavi','efrere',1234567890,'male',25,'2023-07-06','5',1,'patient',19,'complete'),(28,'2023-07-10 13:06:59.680','2023-07-10 13:06:59.680',NULL,'harmil','sanghavi','weretrett',7894561230,'male',20,'2023-07-11','1',1,'patient',9,'complete'),(29,'2023-07-11 04:20:16.863','2023-07-11 04:20:16.863',NULL,'harmil','sanghavi','gyturuierteoiruw',7894561230,'male',23,'2023-07-12','4',1,'patient',9,'complete'),(30,'2023-07-11 04:22:31.381','2023-07-11 04:22:31.381',NULL,'harmil','sanghavi','vbdfvbhfjub',4561237890,'male',24,'2023-07-12','5',1,'patient',28,'complete'),(32,'2023-07-11 05:14:33.985','2023-07-11 05:14:33.985',NULL,'harmil','sanghavi','gjfdgfhadsf',1234567890,'male',30,'2023-07-13','3',1,'patient',13,'complete'),(33,'2023-07-11 05:15:22.831','2023-07-11 05:15:22.831',NULL,'harmil','sanghavi','yhkyukhfghgthh',7845123690,'male',35,'2023-07-12','5',1,'patient',26,'complete'),(34,'2023-07-11 05:16:41.903','2023-07-11 05:16:41.903',NULL,'harmil','sanghavi','dghrthyhrythyth',4567891230,'male',40,'2023-07-29','3',1,'patient',10,'pending'),(35,'2023-07-11 05:16:56.672','2023-07-11 05:16:56.672',NULL,'harmil','sanghavi','dghrthyhrythyth',4567891230,'male',40,'2023-07-27','4',1,'patient',28,'pending'),(36,'2023-07-11 12:47:43.036','2023-07-13 07:41:37.987',NULL,'harmil','sanghavi','dfghthjytjyujky',4567891230,'male',15,'2023-07-13','1',1,'patient',9,'complete'),(37,'2023-07-12 06:18:46.288','2023-07-13 07:49:06.154',NULL,'vishwa','bhatt','grtgthyt',4561237890,'female',22,'2023-07-13','6',31,'patient',9,'complete'),(38,'2023-07-12 11:48:32.123','2023-07-13 13:06:55.393',NULL,'harmil','sanghavi','eryhtjuy7juyg',4567891230,'male',23,'2023-07-14','2',1,'patient',8,'complete'),(39,'2023-07-12 12:17:55.413','2023-07-13 06:13:30.792',NULL,'harmil','sanghavi','tuyjuyuyghjfgh',4567891230,'male',23,'2023-07-15','1',1,'patient',8,'reject'),(40,'2023-07-13 13:10:14.525','2023-07-13 13:11:35.578',NULL,'harmil','sanghavi','ertrytutu',7845123690,'male',20,'2023-07-14','3',1,'patient',8,'complete'),(41,'2023-07-14 03:10:01.060','2023-07-14 03:10:01.060',NULL,'vatsal','sanghavi','dfgthythythth',78974561230,'male',22,'2023-07-15','3',3,'patient',8,'reject'),(43,'2023-07-14 03:16:16.999','2023-07-14 03:16:16.999',NULL,'vatsal','sanghavi','rey6tuuyu',4567891230,'male',24,'2023-07-16','7',3,'patient',19,'pending');
/*!40000 ALTER TABLE `appoinment_books` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `appoinment_times`
--

DROP TABLE IF EXISTS `appoinment_times`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `appoinment_times` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_appoinment_times_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `appoinment_times`
--

LOCK TABLES `appoinment_times` WRITE;
/*!40000 ALTER TABLE `appoinment_times` DISABLE KEYS */;
INSERT INTO `appoinment_times` VALUES (1,NULL,NULL,NULL,'2023-06-26 10:00:00.000'),(2,NULL,NULL,NULL,'2023-06-26 11:00:00.000'),(3,NULL,NULL,NULL,'2023-06-26 12:00:00.000'),(4,NULL,NULL,NULL,'2023-06-26 13:00:00.000'),(5,NULL,NULL,NULL,'2023-06-26 15:00:00.000'),(6,NULL,NULL,NULL,'2023-06-26 16:00:00.000'),(7,NULL,NULL,NULL,'2023-06-26 17:00:00.000'),(8,NULL,NULL,NULL,'2023-06-26 18:30:00.000'),(9,NULL,NULL,NULL,'2023-06-26 19:00:00.000'),(10,NULL,NULL,NULL,'2023-06-26 19:30:00.000');
/*!40000 ALTER TABLE `appoinment_times` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `available_roles`
--

DROP TABLE IF EXISTS `available_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `available_roles` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `role_name` longtext,
  `role_key` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_available_roles_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `available_roles`
--

LOCK TABLES `available_roles` WRITE;
/*!40000 ALTER TABLE `available_roles` DISABLE KEYS */;
INSERT INTO `available_roles` VALUES (1,'2023-07-14 10:32:57.673','2023-07-14 10:32:57.673',NULL,'admin','admin'),(2,'2023-07-14 10:32:57.673','2023-07-14 10:32:57.673',NULL,'doctor','doctor'),(3,'2023-07-14 10:32:57.673','2023-07-14 10:32:57.673',NULL,'patient','patient');
/*!40000 ALTER TABLE `available_roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `doctor_details`
--

DROP TABLE IF EXISTS `doctor_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `doctor_details` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `first_name` longtext,
  `last_name` longtext,
  `number` bigint DEFAULT NULL,
  `gender` longtext,
  `age` bigint DEFAULT NULL,
  `qualification` longtext,
  `department` longtext,
  `user_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_doctor_details_deleted_at` (`deleted_at`),
  KEY `fk_doctor_details_users` (`user_id`),
  CONSTRAINT `fk_doctor_details_users` FOREIGN KEY (`user_id`) REFERENCES `signups` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `doctor_details`
--

LOCK TABLES `doctor_details` WRITE;
/*!40000 ALTER TABLE `doctor_details` DISABLE KEYS */;
INSERT INTO `doctor_details` VALUES (1,'2023-06-26 15:14:57.429','2023-06-26 15:14:57.429',NULL,'peter','parker',1234567890,'male',30,'M.B.B.S','nurology',8),(2,'2023-06-26 15:16:08.156','2023-06-26 15:16:08.156',NULL,'lali','lali',1234567890,'female',35,'M.B.B.S','psycjiatrists',9),(4,'2023-06-28 09:47:41.485','2023-06-28 09:47:41.485',NULL,'vivek','valand',7894560321,'male',23,'M.B.B.S','nurology',10),(8,'2023-06-28 17:26:15.939','2023-06-28 17:26:15.939',NULL,'raj','sanghavi',7894561230,'male',30,'M.B.B.S','nurology',13),(10,'2023-07-03 14:16:30.937','2023-07-03 14:16:30.937',NULL,'diya','shah',4561237890,'female',40,'M.B.B.S','hard wed',19),(12,'2023-07-05 08:30:50.409','2023-07-05 08:30:50.409',NULL,'jay','patel',1234567890,'male',30,'M.B.B.S','nurology',26),(13,'2023-07-06 08:59:51.822','2023-07-06 08:59:51.822',NULL,'rajesh','parmar',4567891230,'male',35,'M.B.B.S','hard wed',28);
/*!40000 ALTER TABLE `doctor_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `signups`
--

DROP TABLE IF EXISTS `signups`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `signups` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `email` longtext,
  `password` longtext,
  `role_name` longtext,
  `number` bigint DEFAULT NULL,
  `role_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_signups_deleted_at` (`deleted_at`),
  KEY `fk_signups_roles` (`role_id`),
  CONSTRAINT `fk_signups_roles` FOREIGN KEY (`role_id`) REFERENCES `available_roles` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `signups`
--

LOCK TABLES `signups` WRITE;
/*!40000 ALTER TABLE `signups` DISABLE KEYS */;
INSERT INTO `signups` VALUES (1,'2023-06-26 10:11:54.422','2023-06-26 10:11:54.422',NULL,'harmil@gmail.com','Harmil@123','patient',NULL,NULL),(3,'2023-06-26 14:27:25.901','2023-06-26 14:27:25.901',NULL,'vatsal@gmail.com','Vatsal@123','patient',NULL,NULL),(5,'2023-06-26 14:53:05.792','2023-06-26 14:53:05.792',NULL,'jhon@gmail.com','Jhon@123','patient',NULL,NULL),(8,'2023-06-26 15:14:57.428','2023-06-26 15:14:57.428',NULL,'peter@gmail.com','Peter@123','doctor',NULL,NULL),(9,'2023-06-26 15:16:08.155','2023-06-26 15:16:08.155',NULL,'lali@gmail.com','Lali@123','doctor',NULL,NULL),(10,'2023-06-28 09:47:41.483','2023-06-28 09:47:41.483',NULL,'vivek@gmail.com','Vivek@123','doctor',NULL,NULL),(13,'2023-06-28 17:26:15.930','2023-06-28 17:26:15.930',NULL,'raj@gmail.com','Raj@123','doctor',NULL,NULL),(14,'2023-06-28 17:28:08.042','2023-06-28 17:28:08.042',NULL,'viren@gmail.com','Viren@123','patient',NULL,NULL),(15,'2023-07-03 14:04:01.234','2023-07-03 14:04:01.234',NULL,'sarthank@gmail.com','Sarthank@123','',NULL,NULL),(17,'2023-07-03 14:12:41.029','2023-07-03 14:12:41.029',NULL,'riya@gmail.com','Riya@123','patient',NULL,NULL),(19,'2023-07-03 14:16:30.935','2023-07-03 14:16:30.935',NULL,'diya@gmail.com','Diya@123','doctor',NULL,NULL),(20,'2023-07-03 11:57:17.512','2023-07-03 11:57:17.512',NULL,'josef@gmail.com','Josef@123','patient',NULL,NULL),(21,'2023-07-03 12:16:28.596','2023-07-03 12:16:28.596',NULL,'darshil@gmail.com','Darshil@123','patient',NULL,NULL),(26,'2023-07-05 08:30:50.399','2023-07-05 08:30:50.399',NULL,'Jay@gmail.com','Jay@123','doctor',NULL,NULL),(27,'2023-07-05 08:32:31.294','2023-07-05 08:32:31.294',NULL,'viren@gmail.com','Viren@123','patient',NULL,NULL),(28,'2023-07-06 08:59:51.815','2023-07-06 08:59:51.815',NULL,'rajesh@gmail.com','Rajesh@123','doctor',NULL,NULL),(29,'2023-07-06 10:52:31.673','2023-07-06 10:52:31.673',NULL,'archar@gmail.com','Archar@123','patient',NULL,NULL),(30,'2023-07-06 11:30:01.849','2023-07-06 11:30:01.849',NULL,'miniyan@gmail.com','Miniyan@123','patient',NULL,NULL),(31,'2023-07-12 05:41:11.108','2023-07-12 05:41:11.108',NULL,'vishwa@gmail.com','Vishwa@123','patient',NULL,NULL);
/*!40000 ALTER TABLE `signups` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `temps`
--

DROP TABLE IF EXISTS `temps`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `temps` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_temps_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `temps`
--

LOCK TABLES `temps` WRITE;
/*!40000 ALTER TABLE `temps` DISABLE KEYS */;
INSERT INTO `temps` VALUES (1,'2023-07-13 12:20:05.351','2023-07-13 12:20:05.351',NULL,'2023-07-05 10:00:00.000'),(2,'2023-07-13 12:20:05.351','2023-07-13 12:20:05.351',NULL,'2023-07-05 11:00:00.000'),(3,'2023-07-13 12:20:05.351','2023-07-13 12:20:05.351',NULL,'2023-07-05 12:00:00.000'),(4,'2023-07-13 12:20:05.351','2023-07-13 12:20:05.351',NULL,'2023-07-05 13:00:00.000'),(5,'2023-07-13 12:20:05.351','2023-07-13 12:20:05.351',NULL,'2023-07-05 15:00:00.000'),(6,'2023-07-13 12:20:05.351','2023-07-13 12:20:05.351',NULL,'2023-07-05 16:00:00.000'),(7,'2023-07-13 12:20:05.351','2023-07-13 12:20:05.351',NULL,'2023-07-05 17:00:00.000'),(8,'2023-07-13 12:20:05.351','2023-07-13 12:20:05.351',NULL,'2023-07-05 18:30:00.000'),(9,'2023-07-13 12:20:05.351','2023-07-13 12:20:05.351',NULL,'2023-07-05 19:00:00.000'),(10,'2023-07-13 12:20:05.351','2023-07-13 12:20:05.351',NULL,'2023-07-05 19:30:00.000');
/*!40000 ALTER TABLE `temps` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-08-07 14:48:58
