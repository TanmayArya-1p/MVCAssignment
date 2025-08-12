-- MySQL dump 10.13  Distrib 8.4.5, for Linux (x86_64)
--
-- Host: localhost    Database: mvc
-- ------------------------------------------------------
-- Server version	8.4.5-0ubuntu0.2

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
-- Table structure for table `items`
--

DROP TABLE IF EXISTS `items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `price` float NOT NULL,
  `image` varchar(255) NOT NULL DEFAULT '/images/placeholder.png',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `items`
--

LOCK TABLES `items` WRITE;
/*!40000 ALTER TABLE `items` DISABLE KEYS */;
INSERT INTO `items` VALUES (1,'burger','just burger',20,'/images/placeholder.png'),(2,'fries','just fries',30,'/images/placeholder.png'),(3,'smoothie','just smoothie',100,'/images/placeholder.png'),(4,'sushi','just sushi',100,'/images/placeholder.png'),(5,'pasta','just pasta',100,'/images/placeholder.png'),(6,'pizza','just pizza',120,'/images/placeholder.png'),(7,'maggi','just maggi',120,'/images/placeholder.png'),(8,'ramen','just ramen',20,'/images/placeholder.png');
/*!40000 ALTER TABLE `items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_items`
--

DROP TABLE IF EXISTS `order_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `order_items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `order_id` int NOT NULL,
  `item_id` int NOT NULL,
  `instructions` varchar(255) DEFAULT NULL,
  `quantity` int NOT NULL,
  `price` float NOT NULL,
  `issued_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status` enum('pending','preparing','served') NOT NULL DEFAULT 'pending',
  PRIMARY KEY (`id`),
  KEY `item_id` (`item_id`),
  KEY `idx_order_item_order` (`order_id`),
  CONSTRAINT `order_items_ibfk_1` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`),
  CONSTRAINT `order_items_ibfk_2` FOREIGN KEY (`item_id`) REFERENCES `items` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_items`
--

LOCK TABLES `order_items` WRITE;
/*!40000 ALTER TABLE `order_items` DISABLE KEYS */;
INSERT INTO `order_items` VALUES (1,1,3,'',3,100,'2025-08-12 21:17:59','preparing'),(2,1,5,'',1,100,'2025-08-12 21:17:59','preparing'),(3,1,2,'',3,30,'2025-08-12 21:17:59','preparing'),(4,1,6,'',2,120,'2025-08-12 21:17:59','preparing'),(5,2,1,'',4,20,'2025-08-12 21:18:09','served'),(6,3,3,'',1,100,'2025-08-12 21:18:19','served'),(7,3,4,'',1,100,'2025-08-12 21:18:19','served'),(8,4,2,'',1,30,'2025-08-12 21:18:30','served'),(9,4,3,'',2,100,'2025-08-12 21:18:30','served'),(10,5,2,'',2,30,'2025-08-12 21:18:42','pending'),(11,5,3,'',2,100,'2025-08-12 21:18:42','pending'),(12,6,3,'extra sweet',2,100,'2025-08-12 21:19:00','pending'),(13,6,1,'',1,20,'2025-08-12 21:19:00','pending'),(14,6,2,'',1,30,'2025-08-12 21:19:00','pending');
/*!40000 ALTER TABLE `order_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `orders`
--

DROP TABLE IF EXISTS `orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `orders` (
  `id` int NOT NULL AUTO_INCREMENT,
  `issued_by` int NOT NULL,
  `issued_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status` enum('pending','preparing','served','billed','paid') NOT NULL DEFAULT 'pending',
  `billable_amount` float DEFAULT NULL,
  `table_no` int DEFAULT NULL,
  `waiter` int DEFAULT NULL,
  `paid_at` timestamp NULL DEFAULT NULL,
  `tip` float DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `issued_by` (`issued_by`),
  KEY `waiter` (`waiter`),
  CONSTRAINT `orders_ibfk_1` FOREIGN KEY (`issued_by`) REFERENCES `users` (`id`),
  CONSTRAINT `orders_ibfk_2` FOREIGN KEY (`waiter`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

LOCK TABLES `orders` WRITE;
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
INSERT INTO `orders` VALUES (1,3,'2025-08-12 21:17:51','preparing',730,1,NULL,NULL,NULL),(2,3,'2025-08-12 21:18:05','served',80,2,NULL,NULL,NULL),(3,3,'2025-08-12 21:18:15','billed',200,6,NULL,NULL,NULL),(4,3,'2025-08-12 21:18:25','paid',230,74,2,'2025-08-12 21:20:01',270),(5,3,'2025-08-12 21:18:37','pending',260,11,NULL,NULL,NULL),(6,3,'2025-08-12 21:18:51','pending',250,65,NULL,NULL,NULL);
/*!40000 ALTER TABLE `orders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `refresh_jti`
--

DROP TABLE IF EXISTS `refresh_jti`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `refresh_jti` (
  `jti` varchar(36) NOT NULL,
  `issued_by` int NOT NULL,
  `expires_at` bigint NOT NULL,
  PRIMARY KEY (`jti`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `refresh_jti`
--

LOCK TABLES `refresh_jti` WRITE;
/*!40000 ALTER TABLE `refresh_jti` DISABLE KEYS */;
INSERT INTO `refresh_jti` VALUES ('532c28ed-8e19-4eb2-bebe-42afc08eb8e0',1,1755035247),('54504af9-3a02-4231-afcb-041a7eef8ca2',2,1755035410),('c1b6617b-191f-4f3d-9286-3eed95468a58',1,1755035226),('e0cfe8e5-38bc-4821-b428-9bff9dad649d',3,1755035353),('e99288bb-e118-448a-af55-fc3d69650018',1,1755035049);
/*!40000 ALTER TABLE `refresh_jti` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `schema_migrations`
--

DROP TABLE IF EXISTS `schema_migrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `schema_migrations` (
  `version` bigint NOT NULL,
  `dirty` tinyint(1) NOT NULL,
  PRIMARY KEY (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `schema_migrations`
--

LOCK TABLES `schema_migrations` WRITE;
/*!40000 ALTER TABLE `schema_migrations` DISABLE KEYS */;
INSERT INTO `schema_migrations` VALUES (5,0);
/*!40000 ALTER TABLE `schema_migrations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tag_rel`
--

DROP TABLE IF EXISTS `tag_rel`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tag_rel` (
  `id` int NOT NULL AUTO_INCREMENT,
  `item_id` int NOT NULL,
  `tag_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `tag_rel_item_id_idx` (`item_id`),
  KEY `tag_rel_tag_id_idx` (`tag_id`),
  CONSTRAINT `tag_rel_ibfk_1` FOREIGN KEY (`item_id`) REFERENCES `items` (`id`),
  CONSTRAINT `tag_rel_ibfk_2` FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tag_rel`
--

LOCK TABLES `tag_rel` WRITE;
/*!40000 ALTER TABLE `tag_rel` DISABLE KEYS */;
INSERT INTO `tag_rel` VALUES (1,1,1),(2,1,2),(3,2,3),(4,2,4),(5,3,5),(6,3,6),(7,4,7),(8,4,8),(9,5,9),(10,6,9),(11,7,4),(12,7,10),(13,8,7);
/*!40000 ALTER TABLE `tag_rel` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tags`
--

DROP TABLE IF EXISTS `tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tags` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tags`
--

LOCK TABLES `tags` WRITE;
/*!40000 ALTER TABLE `tags` DISABLE KEYS */;
INSERT INTO `tags` VALUES (1,'hot'),(2,'spicy'),(3,'fried'),(4,'snack'),(5,'drinks'),(6,'sweet'),(7,'japanese'),(8,'fish'),(9,'italian'),(10,'indian');
/*!40000 ALTER TABLE `tags` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` enum('admin','chef','customer') NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'admin','$2b$10$bk4Kk88nnC2IaWLpSy57LuaX1U6CrYmwSn.xYbF3wVqb69y1mh0wC','admin','2025-08-12 21:13:09'),(2,'chef','$2a$10$EZTo3z0HI59Oo06Wbd8zcO6fnzfzIsDIthVYLswFv8tv8WpVa5twO','chef','2025-08-12 21:13:09'),(3,'test','$2a$10$wBDfegu7gNqvfagOk/Z2.eo7RiKjwgwcwXSoad0x3VYiRYK.qjR/e','customer','2025-08-12 21:17:38');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-08-13  2:50:30
