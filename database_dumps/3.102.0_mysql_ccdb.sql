-- MySQL dump 10.13  Distrib 8.0.23, for Linux (x86_64)
--
-- Host: localhost    Database: ccdb
-- ------------------------------------------------------
-- Server version	8.0.25

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
-- Table structure for table `app_annotations`
--

DROP TABLE IF EXISTS `app_annotations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `app_annotations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key` varchar(1000) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(5000) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `app_annotations_guid_index` (`guid`),
  KEY `app_annotations_created_at_index` (`created_at`),
  KEY `app_annotations_updated_at_index` (`updated_at`),
  KEY `fk_app_annotations_resource_guid_index` (`resource_guid`),
  CONSTRAINT `fk_app_annotations_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `apps` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `app_annotations`
--

LOCK TABLES `app_annotations` WRITE;
/*!40000 ALTER TABLE `app_annotations` DISABLE KEYS */;
/*!40000 ALTER TABLE `app_annotations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `app_events`
--

DROP TABLE IF EXISTS `app_events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `app_events` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `app_id` int NOT NULL,
  `instance_guid` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `instance_index` int NOT NULL,
  `exit_status` int NOT NULL,
  `timestamp` datetime NOT NULL,
  `exit_description` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `app_events_guid_index` (`guid`),
  KEY `app_events_app_id_index` (`app_id`),
  KEY `app_events_created_at_index` (`created_at`),
  KEY `app_events_updated_at_index` (`updated_at`),
  CONSTRAINT `fk_app_events_app_id` FOREIGN KEY (`app_id`) REFERENCES `processes` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `app_events`
--

LOCK TABLES `app_events` WRITE;
/*!40000 ALTER TABLE `app_events` DISABLE KEYS */;
/*!40000 ALTER TABLE `app_events` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `app_labels`
--

DROP TABLE IF EXISTS `app_labels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `app_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key_name` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `app_labels_guid_index` (`guid`),
  KEY `app_labels_created_at_index` (`created_at`),
  KEY `app_labels_updated_at_index` (`updated_at`),
  KEY `fk_app_labels_resource_guid_index` (`resource_guid`),
  KEY `app_labels_compound_index` (`key_prefix`,`key_name`,`value`),
  CONSTRAINT `fk_app_labels_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `apps` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `app_labels`
--

LOCK TABLES `app_labels` WRITE;
/*!40000 ALTER TABLE `app_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `app_labels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `app_usage_events`
--

DROP TABLE IF EXISTS `app_usage_events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `app_usage_events` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `instance_count` int NOT NULL,
  `memory_in_mb_per_instance` int NOT NULL,
  `state` varchar(255) COLLATE utf8_bin NOT NULL,
  `app_guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `app_name` varchar(255) COLLATE utf8_bin NOT NULL,
  `space_guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `space_name` varchar(255) COLLATE utf8_bin NOT NULL,
  `org_guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `buildpack_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `buildpack_name` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `package_state` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `parent_app_name` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `parent_app_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `process_type` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `task_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `task_name` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `package_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `previous_state` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `previous_package_state` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `previous_memory_in_mb_per_instance` int DEFAULT NULL,
  `previous_instance_count` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `app_usage_events_guid_index` (`guid`),
  KEY `usage_events_created_at_index` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `app_usage_events`
--

LOCK TABLES `app_usage_events` WRITE;
/*!40000 ALTER TABLE `app_usage_events` DISABLE KEYS */;
/*!40000 ALTER TABLE `app_usage_events` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `apps`
--

DROP TABLE IF EXISTS `apps`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `apps` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `space_guid` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `droplet_guid` varchar(255) DEFAULT NULL,
  `desired_state` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT 'STOPPED',
  `encrypted_environment_variables` text CHARACTER SET utf8 COLLATE utf8_bin,
  `salt` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `max_task_sequence_id` int DEFAULT '1',
  `buildpack_cache_sha256_checksum` varchar(255) DEFAULT NULL,
  `enable_ssh` tinyint(1) DEFAULT NULL,
  `encryption_key_label` varchar(255) DEFAULT NULL,
  `encryption_iterations` int NOT NULL DEFAULT '2048',
  `revisions_enabled` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `apps_v3_guid_index` (`guid`),
  UNIQUE KEY `apps_v3_space_guid_name_index` (`space_guid`,`name`),
  KEY `apps_v3_created_at_index` (`created_at`),
  KEY `apps_v3_updated_at_index` (`updated_at`),
  KEY `apps_v3_space_guid_index` (`space_guid`),
  KEY `apps_v3_name_index` (`name`),
  KEY `apps_desired_droplet_guid` (`droplet_guid`),
  CONSTRAINT `fk_apps_space_guid` FOREIGN KEY (`space_guid`) REFERENCES `spaces` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `apps`
--

LOCK TABLES `apps` WRITE;
/*!40000 ALTER TABLE `apps` DISABLE KEYS */;
/*!40000 ALTER TABLE `apps` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `build_annotations`
--

DROP TABLE IF EXISTS `build_annotations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `build_annotations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key` varchar(1000) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(5000) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `build_annotations_guid_index` (`guid`),
  KEY `build_annotations_created_at_index` (`created_at`),
  KEY `build_annotations_updated_at_index` (`updated_at`),
  KEY `fk_build_annotations_resource_guid_index` (`resource_guid`),
  CONSTRAINT `fk_build_annotations_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `builds` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `build_annotations`
--

LOCK TABLES `build_annotations` WRITE;
/*!40000 ALTER TABLE `build_annotations` DISABLE KEYS */;
/*!40000 ALTER TABLE `build_annotations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `build_labels`
--

DROP TABLE IF EXISTS `build_labels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `build_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key_name` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `build_labels_guid_index` (`guid`),
  KEY `build_labels_created_at_index` (`created_at`),
  KEY `build_labels_updated_at_index` (`updated_at`),
  KEY `fk_build_labels_resource_guid_index` (`resource_guid`),
  KEY `build_labels_compound_index` (`key_prefix`,`key_name`,`value`),
  CONSTRAINT `fk_build_labels_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `builds` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `build_labels`
--

LOCK TABLES `build_labels` WRITE;
/*!40000 ALTER TABLE `build_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `build_labels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `buildpack_annotations`
--

DROP TABLE IF EXISTS `buildpack_annotations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `buildpack_annotations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key` varchar(1000) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(5000) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `buildpack_annotations_guid_index` (`guid`),
  KEY `buildpack_annotations_created_at_index` (`created_at`),
  KEY `buildpack_annotations_updated_at_index` (`updated_at`),
  KEY `fk_buildpack_annotations_resource_guid_index` (`resource_guid`),
  CONSTRAINT `fk_buildpack_annotations_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `buildpacks` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `buildpack_annotations`
--

LOCK TABLES `buildpack_annotations` WRITE;
/*!40000 ALTER TABLE `buildpack_annotations` DISABLE KEYS */;
/*!40000 ALTER TABLE `buildpack_annotations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `buildpack_labels`
--

DROP TABLE IF EXISTS `buildpack_labels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `buildpack_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key_name` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `buildpack_labels_guid_index` (`guid`),
  KEY `buildpack_labels_created_at_index` (`created_at`),
  KEY `buildpack_labels_updated_at_index` (`updated_at`),
  KEY `fk_buildpack_labels_resource_guid_index` (`resource_guid`),
  KEY `buildpack_labels_compound_index` (`key_prefix`,`key_name`,`value`),
  CONSTRAINT `fk_buildpack_labels_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `buildpacks` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `buildpack_labels`
--

LOCK TABLES `buildpack_labels` WRITE;
/*!40000 ALTER TABLE `buildpack_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `buildpack_labels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `buildpack_lifecycle_buildpacks`
--

DROP TABLE IF EXISTS `buildpack_lifecycle_buildpacks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `buildpack_lifecycle_buildpacks` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `admin_buildpack_name` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `encrypted_buildpack_url` varchar(16000) COLLATE utf8_bin DEFAULT NULL,
  `encrypted_buildpack_url_salt` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `buildpack_lifecycle_data_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `encryption_key_label` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `version` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `buildpack_name` varchar(2047) COLLATE utf8_bin DEFAULT NULL,
  `encryption_iterations` int NOT NULL DEFAULT '2048',
  PRIMARY KEY (`id`),
  UNIQUE KEY `buildpack_lifecycle_buildpacks_guid_index` (`guid`),
  KEY `buildpack_lifecycle_buildpacks_created_at_index` (`created_at`),
  KEY `buildpack_lifecycle_buildpacks_updated_at_index` (`updated_at`),
  KEY `bl_buildpack_bldata_guid_index` (`buildpack_lifecycle_data_guid`),
  CONSTRAINT `fk_blbuildpack_bldata_guid` FOREIGN KEY (`buildpack_lifecycle_data_guid`) REFERENCES `buildpack_lifecycle_data` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `buildpack_lifecycle_buildpacks`
--

LOCK TABLES `buildpack_lifecycle_buildpacks` WRITE;
/*!40000 ALTER TABLE `buildpack_lifecycle_buildpacks` DISABLE KEYS */;
/*!40000 ALTER TABLE `buildpack_lifecycle_buildpacks` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `buildpack_lifecycle_data`
--

DROP TABLE IF EXISTS `buildpack_lifecycle_data`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `buildpack_lifecycle_data` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `app_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `droplet_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `stack` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `encrypted_buildpack_url` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `encrypted_buildpack_url_salt` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `admin_buildpack_name` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `build_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `encryption_key_label` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `encryption_iterations` int NOT NULL DEFAULT '2048',
  PRIMARY KEY (`id`),
  UNIQUE KEY `buildpack_lifecycle_data_guid_index` (`guid`),
  KEY `buildpack_lifecycle_data_created_at_index` (`created_at`),
  KEY `buildpack_lifecycle_data_updated_at_index` (`updated_at`),
  KEY `buildpack_lifecycle_data_app_guid` (`app_guid`),
  KEY `bp_lifecycle_data_droplet_guid` (`droplet_guid`),
  KEY `buildpack_lifecycle_data_admin_buildpack_name_index` (`admin_buildpack_name`),
  KEY `buildpack_lifecycle_data_build_guid_index` (`build_guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `buildpack_lifecycle_data`
--

LOCK TABLES `buildpack_lifecycle_data` WRITE;
/*!40000 ALTER TABLE `buildpack_lifecycle_data` DISABLE KEYS */;
/*!40000 ALTER TABLE `buildpack_lifecycle_data` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `buildpacks`
--

DROP TABLE IF EXISTS `buildpacks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `buildpacks` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) COLLATE utf8_bin NOT NULL,
  `key` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `position` int NOT NULL,
  `enabled` tinyint(1) DEFAULT '1',
  `locked` tinyint(1) DEFAULT '0',
  `filename` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `sha256_checksum` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `stack` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `buildpacks_guid_index` (`guid`),
  UNIQUE KEY `unique_name_and_stack` (`name`,`stack`),
  KEY `buildpacks_created_at_index` (`created_at`),
  KEY `buildpacks_updated_at_index` (`updated_at`),
  KEY `buildpacks_key_index` (`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `buildpacks`
--

LOCK TABLES `buildpacks` WRITE;
/*!40000 ALTER TABLE `buildpacks` DISABLE KEYS */;
/*!40000 ALTER TABLE `buildpacks` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `builds`
--

DROP TABLE IF EXISTS `builds`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `builds` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `state` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `package_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `error_description` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `app_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `error_id` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `created_by_user_guid` text COLLATE utf8_bin,
  `created_by_user_name` text COLLATE utf8_bin,
  `created_by_user_email` text COLLATE utf8_bin,
  `staging_memory_in_mb` int DEFAULT NULL,
  `staging_disk_in_mb` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `builds_guid_index` (`guid`),
  KEY `builds_created_at_index` (`created_at`),
  KEY `builds_updated_at_index` (`updated_at`),
  KEY `builds_app_guid_index` (`app_guid`),
  KEY `builds_state_index` (`state`),
  CONSTRAINT `fk_builds_app_guid` FOREIGN KEY (`app_guid`) REFERENCES `apps` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `builds`
--

LOCK TABLES `builds` WRITE;
/*!40000 ALTER TABLE `builds` DISABLE KEYS */;
/*!40000 ALTER TABLE `builds` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `clock_jobs`
--

DROP TABLE IF EXISTS `clock_jobs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `clock_jobs` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8_bin NOT NULL,
  `last_started_at` datetime DEFAULT NULL,
  `last_completed_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `clock_jobs_name_unique` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `clock_jobs`
--

LOCK TABLES `clock_jobs` WRITE;
/*!40000 ALTER TABLE `clock_jobs` DISABLE KEYS */;
/*!40000 ALTER TABLE `clock_jobs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `delayed_jobs`
--

DROP TABLE IF EXISTS `delayed_jobs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `delayed_jobs` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `priority` int DEFAULT '0',
  `attempts` int DEFAULT '0',
  `handler` longtext COLLATE utf8_bin,
  `last_error` text COLLATE utf8_bin,
  `run_at` datetime DEFAULT NULL,
  `locked_at` datetime DEFAULT NULL,
  `failed_at` datetime DEFAULT NULL,
  `locked_by` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `queue` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `cf_api_error` text COLLATE utf8_bin,
  PRIMARY KEY (`id`),
  UNIQUE KEY `dj_guid_index` (`guid`),
  KEY `dj_created_at_index` (`created_at`),
  KEY `dj_updated_at_index` (`updated_at`),
  KEY `delayed_jobs_reserve` (`queue`,`locked_at`,`locked_by`,`failed_at`,`run_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `delayed_jobs`
--

LOCK TABLES `delayed_jobs` WRITE;
/*!40000 ALTER TABLE `delayed_jobs` DISABLE KEYS */;
/*!40000 ALTER TABLE `delayed_jobs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `deployment_annotations`
--

DROP TABLE IF EXISTS `deployment_annotations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `deployment_annotations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key` varchar(1000) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(5000) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `deployment_annotations_guid_index` (`guid`),
  KEY `deployment_annotations_created_at_index` (`created_at`),
  KEY `deployment_annotations_updated_at_index` (`updated_at`),
  KEY `fk_deployment_annotations_resource_guid_index` (`resource_guid`),
  CONSTRAINT `fk_deployment_annotations_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `deployments` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `deployment_annotations`
--

LOCK TABLES `deployment_annotations` WRITE;
/*!40000 ALTER TABLE `deployment_annotations` DISABLE KEYS */;
/*!40000 ALTER TABLE `deployment_annotations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `deployment_labels`
--

DROP TABLE IF EXISTS `deployment_labels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `deployment_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key_name` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `deployment_labels_guid_index` (`guid`),
  KEY `deployment_labels_created_at_index` (`created_at`),
  KEY `deployment_labels_updated_at_index` (`updated_at`),
  KEY `fk_deployment_labels_resource_guid_index` (`resource_guid`),
  KEY `deployment_labels_compound_index` (`key_prefix`,`key_name`,`value`),
  CONSTRAINT `fk_deployment_labels_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `deployments` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `deployment_labels`
--

LOCK TABLES `deployment_labels` WRITE;
/*!40000 ALTER TABLE `deployment_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `deployment_labels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `deployment_processes`
--

DROP TABLE IF EXISTS `deployment_processes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `deployment_processes` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `process_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `process_type` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `deployment_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `deployment_processes_guid_index` (`guid`),
  KEY `deployment_processes_created_at_index` (`created_at`),
  KEY `deployment_processes_updated_at_index` (`updated_at`),
  KEY `deployment_processes_deployment_guid_index` (`deployment_guid`),
  CONSTRAINT `fk_deployment_processes_deployment_guid` FOREIGN KEY (`deployment_guid`) REFERENCES `deployments` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `deployment_processes`
--

LOCK TABLES `deployment_processes` WRITE;
/*!40000 ALTER TABLE `deployment_processes` DISABLE KEYS */;
/*!40000 ALTER TABLE `deployment_processes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `deployments`
--

DROP TABLE IF EXISTS `deployments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `deployments` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `state` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `app_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `droplet_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `deploying_web_process_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `previous_droplet_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `original_web_process_instance_count` int NOT NULL,
  `revision_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `revision_version` int DEFAULT NULL,
  `last_healthy_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status_value` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `status_reason` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `strategy` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT 'rolling',
  PRIMARY KEY (`id`),
  UNIQUE KEY `deployments_guid_index` (`guid`),
  KEY `deployments_created_at_index` (`created_at`),
  KEY `deployments_updated_at_index` (`updated_at`),
  KEY `deployments_app_guid_index` (`app_guid`),
  KEY `deployments_state_index` (`state`),
  KEY `deployments_status_value_index` (`status_value`),
  KEY `deployments_status_reason_index` (`status_reason`),
  CONSTRAINT `deployments_app_guid_fkey` FOREIGN KEY (`app_guid`) REFERENCES `apps` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `deployments`
--

LOCK TABLES `deployments` WRITE;
/*!40000 ALTER TABLE `deployments` DISABLE KEYS */;
/*!40000 ALTER TABLE `deployments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `domain_annotations`
--

DROP TABLE IF EXISTS `domain_annotations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `domain_annotations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key` varchar(1000) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(5000) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `domain_annotations_guid_index` (`guid`),
  KEY `domain_annotations_created_at_index` (`created_at`),
  KEY `domain_annotations_updated_at_index` (`updated_at`),
  KEY `fk_domain_annotations_resource_guid_index` (`resource_guid`),
  CONSTRAINT `fk_domain_annotations_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `domains` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `domain_annotations`
--

LOCK TABLES `domain_annotations` WRITE;
/*!40000 ALTER TABLE `domain_annotations` DISABLE KEYS */;
/*!40000 ALTER TABLE `domain_annotations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `domain_labels`
--

DROP TABLE IF EXISTS `domain_labels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `domain_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key_name` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `domain_labels_guid_index` (`guid`),
  KEY `domain_labels_created_at_index` (`created_at`),
  KEY `domain_labels_updated_at_index` (`updated_at`),
  KEY `fk_domain_labels_resource_guid_index` (`resource_guid`),
  KEY `domain_labels_compound_index` (`key_prefix`,`key_name`,`value`),
  CONSTRAINT `fk_domain_labels_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `domains` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `domain_labels`
--

LOCK TABLES `domain_labels` WRITE;
/*!40000 ALTER TABLE `domain_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `domain_labels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `domains`
--

DROP TABLE IF EXISTS `domains`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `domains` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `wildcard` tinyint(1) NOT NULL DEFAULT '1',
  `owning_organization_id` int DEFAULT NULL,
  `router_group_guid` varchar(255) DEFAULT NULL,
  `internal` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `domains_guid_index` (`guid`),
  UNIQUE KEY `domains_name_index` (`name`),
  KEY `domains_created_at_index` (`created_at`),
  KEY `domains_updated_at_index` (`updated_at`),
  KEY `fk_domains_owning_org_id` (`owning_organization_id`),
  CONSTRAINT `fk_domains_owning_org_id` FOREIGN KEY (`owning_organization_id`) REFERENCES `organizations` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `domains`
--

LOCK TABLES `domains` WRITE;
/*!40000 ALTER TABLE `domains` DISABLE KEYS */;
/*!40000 ALTER TABLE `domains` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `droplet_annotations`
--

DROP TABLE IF EXISTS `droplet_annotations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `droplet_annotations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key` varchar(1000) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(5000) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `droplet_annotations_guid_index` (`guid`),
  KEY `droplet_annotations_created_at_index` (`created_at`),
  KEY `droplet_annotations_updated_at_index` (`updated_at`),
  KEY `fk_droplet_annotations_resource_guid_index` (`resource_guid`),
  CONSTRAINT `fk_droplet_annotations_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `droplets` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `droplet_annotations`
--

LOCK TABLES `droplet_annotations` WRITE;
/*!40000 ALTER TABLE `droplet_annotations` DISABLE KEYS */;
/*!40000 ALTER TABLE `droplet_annotations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `droplet_labels`
--

DROP TABLE IF EXISTS `droplet_labels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `droplet_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key_name` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `droplet_labels_guid_index` (`guid`),
  KEY `droplet_labels_created_at_index` (`created_at`),
  KEY `droplet_labels_updated_at_index` (`updated_at`),
  KEY `fk_droplet_labels_resource_guid_index` (`resource_guid`),
  KEY `droplet_labels_compound_index` (`key_prefix`,`key_name`,`value`),
  CONSTRAINT `fk_droplet_labels_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `droplets` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `droplet_labels`
--

LOCK TABLES `droplet_labels` WRITE;
/*!40000 ALTER TABLE `droplet_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `droplet_labels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `droplets`
--

DROP TABLE IF EXISTS `droplets`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `droplets` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `droplet_hash` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `execution_metadata` text COLLATE utf8_bin,
  `state` varchar(255) COLLATE utf8_bin NOT NULL,
  `process_types` text COLLATE utf8_bin,
  `error_id` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `error_description` text COLLATE utf8_bin,
  `encrypted_environment_variables` text COLLATE utf8_bin,
  `salt` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `staging_memory_in_mb` int DEFAULT NULL,
  `staging_disk_in_mb` int DEFAULT NULL,
  `buildpack_receipt_buildpack` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `buildpack_receipt_buildpack_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `buildpack_receipt_detect_output` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `docker_receipt_image` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `package_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `app_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `sha256_checksum` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `build_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `docker_receipt_username` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `docker_receipt_password_salt` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `encrypted_docker_receipt_password` varchar(16000) COLLATE utf8_bin DEFAULT NULL,
  `encryption_key_label` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `encryption_iterations` int NOT NULL DEFAULT '2048',
  `sidecars` text COLLATE utf8_bin,
  PRIMARY KEY (`id`),
  UNIQUE KEY `droplets_guid_index` (`guid`),
  KEY `droplets_created_at_index` (`created_at`),
  KEY `droplets_updated_at_index` (`updated_at`),
  KEY `droplets_droplet_hash_index` (`droplet_hash`),
  KEY `droplets_state_index` (`state`),
  KEY `package_guid_index` (`package_guid`),
  KEY `droplets_sha256_checksum_index` (`sha256_checksum`),
  KEY `droplet_app_guid_index` (`app_guid`),
  KEY `droplet_build_guid_index` (`build_guid`),
  KEY `droplets_guid_droplet_hash_index` (`guid`,`droplet_hash`),
  CONSTRAINT `fk_droplets_app_guid` FOREIGN KEY (`app_guid`) REFERENCES `apps` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `droplets`
--

LOCK TABLES `droplets` WRITE;
/*!40000 ALTER TABLE `droplets` DISABLE KEYS */;
/*!40000 ALTER TABLE `droplets` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `encryption_key_sentinels`
--

DROP TABLE IF EXISTS `encryption_key_sentinels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `encryption_key_sentinels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `expected_value` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `encrypted_value` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `encryption_key_label` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `salt` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `encryption_iterations` int NOT NULL DEFAULT '2048',
  PRIMARY KEY (`id`),
  UNIQUE KEY `encryption_key_sentinels_guid_index` (`guid`),
  UNIQUE KEY `encryption_key_sentinels_label_index` (`encryption_key_label`),
  KEY `encryption_key_sentinels_created_at_index` (`created_at`),
  KEY `encryption_key_sentinels_updated_at_index` (`updated_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `encryption_key_sentinels`
--

LOCK TABLES `encryption_key_sentinels` WRITE;
/*!40000 ALTER TABLE `encryption_key_sentinels` DISABLE KEYS */;
/*!40000 ALTER TABLE `encryption_key_sentinels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `env_groups`
--

DROP TABLE IF EXISTS `env_groups`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `env_groups` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) COLLATE utf8_bin NOT NULL,
  `environment_json` text COLLATE utf8_bin,
  `salt` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `encryption_key_label` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `encryption_iterations` int NOT NULL DEFAULT '2048',
  PRIMARY KEY (`id`),
  UNIQUE KEY `evg_guid_index` (`guid`),
  UNIQUE KEY `env_groups_name_index` (`name`),
  KEY `evg_created_at_index` (`created_at`),
  KEY `evg_updated_at_index` (`updated_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `env_groups`
--

LOCK TABLES `env_groups` WRITE;
/*!40000 ALTER TABLE `env_groups` DISABLE KEYS */;
/*!40000 ALTER TABLE `env_groups` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `events`
--

DROP TABLE IF EXISTS `events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `events` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `timestamp` datetime NOT NULL,
  `type` varchar(255) COLLATE utf8_bin NOT NULL,
  `actor` varchar(255) COLLATE utf8_bin NOT NULL,
  `actor_type` varchar(255) COLLATE utf8_bin NOT NULL,
  `actee` varchar(255) COLLATE utf8_bin NOT NULL,
  `actee_type` varchar(255) COLLATE utf8_bin NOT NULL,
  `metadata` text COLLATE utf8_bin,
  `organization_guid` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `space_guid` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `actor_name` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `actee_name` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `actor_username` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `events_guid_index` (`guid`),
  KEY `events_created_at_index` (`created_at`),
  KEY `events_updated_at_index` (`updated_at`),
  KEY `events_type_index` (`type`),
  KEY `events_actee_index` (`actee`),
  KEY `events_space_guid_index` (`space_guid`),
  KEY `events_organization_guid_index` (`organization_guid`),
  KEY `events_actee_type_index` (`actee_type`),
  KEY `events_timestamp_id_index` (`timestamp`,`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `events`
--

LOCK TABLES `events` WRITE;
/*!40000 ALTER TABLE `events` DISABLE KEYS */;
/*!40000 ALTER TABLE `events` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `feature_flags`
--

DROP TABLE IF EXISTS `feature_flags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `feature_flags` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) COLLATE utf8_bin NOT NULL,
  `enabled` tinyint(1) NOT NULL,
  `error_message` text COLLATE utf8_bin,
  PRIMARY KEY (`id`),
  UNIQUE KEY `feature_flag_guid_index` (`guid`),
  UNIQUE KEY `feature_flags_name_index` (`name`),
  KEY `feature_flag_created_at_index` (`created_at`),
  KEY `feature_flag_updated_at_index` (`updated_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `feature_flags`
--

LOCK TABLES `feature_flags` WRITE;
/*!40000 ALTER TABLE `feature_flags` DISABLE KEYS */;
/*!40000 ALTER TABLE `feature_flags` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `isolation_segment_annotations`
--

DROP TABLE IF EXISTS `isolation_segment_annotations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `isolation_segment_annotations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key` varchar(1000) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(5000) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `isolation_segment_annotations_guid_index` (`guid`),
  KEY `isolation_segment_annotations_created_at_index` (`created_at`),
  KEY `isolation_segment_annotations_updated_at_index` (`updated_at`),
  KEY `fk_isolation_segment_annotations_resource_guid_index` (`resource_guid`),
  CONSTRAINT `fk_isolation_segment_annotations_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `isolation_segments` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `isolation_segment_annotations`
--

LOCK TABLES `isolation_segment_annotations` WRITE;
/*!40000 ALTER TABLE `isolation_segment_annotations` DISABLE KEYS */;
/*!40000 ALTER TABLE `isolation_segment_annotations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `isolation_segment_labels`
--

DROP TABLE IF EXISTS `isolation_segment_labels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `isolation_segment_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key_name` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `isolation_segment_labels_guid_index` (`guid`),
  KEY `isolation_segment_labels_created_at_index` (`created_at`),
  KEY `isolation_segment_labels_updated_at_index` (`updated_at`),
  KEY `fk_isolation_segment_labels_resource_guid_index` (`resource_guid`),
  KEY `isolation_segment_labels_compound_index` (`key_prefix`,`key_name`,`value`),
  CONSTRAINT `fk_isolation_segment_labels_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `isolation_segments` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `isolation_segment_labels`
--

LOCK TABLES `isolation_segment_labels` WRITE;
/*!40000 ALTER TABLE `isolation_segment_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `isolation_segment_labels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `isolation_segments`
--

DROP TABLE IF EXISTS `isolation_segments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `isolation_segments` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `isolation_segments_guid_index` (`guid`),
  UNIQUE KEY `isolation_segment_name_unique_constraint` (`name`),
  KEY `isolation_segments_created_at_index` (`created_at`),
  KEY `isolation_segments_updated_at_index` (`updated_at`),
  KEY `isolation_segments_name_index` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `isolation_segments`
--

LOCK TABLES `isolation_segments` WRITE;
/*!40000 ALTER TABLE `isolation_segments` DISABLE KEYS */;
/*!40000 ALTER TABLE `isolation_segments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `job_warnings`
--

DROP TABLE IF EXISTS `job_warnings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `job_warnings` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `detail` varchar(16000) COLLATE utf8_bin NOT NULL,
  `job_id` int NOT NULL,
  `fk_jobs_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `job_warnings_guid_index` (`guid`),
  KEY `fk_jobs_id` (`fk_jobs_id`),
  KEY `job_warnings_created_at_index` (`created_at`),
  KEY `job_warnings_updated_at_index` (`updated_at`),
  KEY `fk_jobs_id_index` (`job_id`),
  CONSTRAINT `job_warnings_ibfk_1` FOREIGN KEY (`fk_jobs_id`) REFERENCES `jobs` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `job_warnings`
--

LOCK TABLES `job_warnings` WRITE;
/*!40000 ALTER TABLE `job_warnings` DISABLE KEYS */;
/*!40000 ALTER TABLE `job_warnings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `jobs`
--

DROP TABLE IF EXISTS `jobs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `jobs` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `state` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `operation` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `resource_type` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `delayed_job_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `cf_api_error` varchar(16000) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `jobs_guid_index` (`guid`),
  KEY `jobs_created_at_index` (`created_at`),
  KEY `jobs_updated_at_index` (`updated_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `jobs`
--

LOCK TABLES `jobs` WRITE;
/*!40000 ALTER TABLE `jobs` DISABLE KEYS */;
/*!40000 ALTER TABLE `jobs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `kpack_lifecycle_data`
--

DROP TABLE IF EXISTS `kpack_lifecycle_data`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `kpack_lifecycle_data` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `build_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `droplet_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `app_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `buildpacks` varchar(5000) COLLATE utf8_bin DEFAULT '[]',
  PRIMARY KEY (`id`),
  UNIQUE KEY `kpack_lifecycle_data_guid_index` (`guid`),
  KEY `kpack_lifecycle_data_created_at_index` (`created_at`),
  KEY `kpack_lifecycle_data_updated_at_index` (`updated_at`),
  KEY `fk_kpack_lifecycle_build_guid_index` (`build_guid`),
  KEY `fk_kpack_lifecycle_app_guid` (`app_guid`),
  KEY `kpack_lifecycle_droplet_guid_index` (`droplet_guid`),
  CONSTRAINT `fk_kpack_lifecycle_app_guid` FOREIGN KEY (`app_guid`) REFERENCES `apps` (`guid`),
  CONSTRAINT `fk_kpack_lifecycle_build_guid` FOREIGN KEY (`build_guid`) REFERENCES `builds` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kpack_lifecycle_data`
--

LOCK TABLES `kpack_lifecycle_data` WRITE;
/*!40000 ALTER TABLE `kpack_lifecycle_data` DISABLE KEYS */;
/*!40000 ALTER TABLE `kpack_lifecycle_data` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lockings`
--

DROP TABLE IF EXISTS `lockings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `lockings` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `lockings_name_index` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lockings`
--

LOCK TABLES `lockings` WRITE;
/*!40000 ALTER TABLE `lockings` DISABLE KEYS */;
/*!40000 ALTER TABLE `lockings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `organization_annotations`
--

DROP TABLE IF EXISTS `organization_annotations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `organization_annotations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key` varchar(1000) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(5000) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `organization_annotations_guid_index` (`guid`),
  KEY `organization_annotations_created_at_index` (`created_at`),
  KEY `organization_annotations_updated_at_index` (`updated_at`),
  KEY `fk_organization_annotations_resource_guid_index` (`resource_guid`),
  CONSTRAINT `fk_organization_annotations_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `organizations` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `organization_annotations`
--

LOCK TABLES `organization_annotations` WRITE;
/*!40000 ALTER TABLE `organization_annotations` DISABLE KEYS */;
/*!40000 ALTER TABLE `organization_annotations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `organization_labels`
--

DROP TABLE IF EXISTS `organization_labels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `organization_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key_name` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `organization_labels_guid_index` (`guid`),
  KEY `organization_labels_created_at_index` (`created_at`),
  KEY `organization_labels_updated_at_index` (`updated_at`),
  KEY `fk_organization_labels_resource_guid_index` (`resource_guid`),
  KEY `organization_labels_compound_index` (`key_prefix`,`key_name`,`value`),
  CONSTRAINT `fk_organization_labels_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `organizations` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `organization_labels`
--

LOCK TABLES `organization_labels` WRITE;
/*!40000 ALTER TABLE `organization_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `organization_labels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `organizations`
--

DROP TABLE IF EXISTS `organizations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `organizations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `billing_enabled` tinyint(1) NOT NULL DEFAULT '0',
  `quota_definition_id` int NOT NULL,
  `status` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT 'active',
  `default_isolation_segment_guid` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `organizations_guid_index` (`guid`),
  UNIQUE KEY `organizations_name_index` (`name`),
  KEY `organizations_created_at_index` (`created_at`),
  KEY `organizations_updated_at_index` (`updated_at`),
  KEY `fk_org_quota_definition_id` (`quota_definition_id`),
  KEY `organizations_isolation_segments_pk` (`guid`,`default_isolation_segment_guid`),
  CONSTRAINT `fk_org_quota_definition_id` FOREIGN KEY (`quota_definition_id`) REFERENCES `quota_definitions` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `organizations`
--

LOCK TABLES `organizations` WRITE;
/*!40000 ALTER TABLE `organizations` DISABLE KEYS */;
/*!40000 ALTER TABLE `organizations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `organizations_auditors`
--

DROP TABLE IF EXISTS `organizations_auditors`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `organizations_auditors` (
  `organization_id` int NOT NULL,
  `user_id` int NOT NULL,
  `organizations_auditors_pk` int NOT NULL AUTO_INCREMENT,
  `role_guid` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`organizations_auditors_pk`),
  UNIQUE KEY `org_auditors_idx` (`organization_id`,`user_id`),
  KEY `org_auditors_user_fk` (`user_id`),
  KEY `organizations_auditors_role_guid_index` (`role_guid`),
  KEY `organizations_auditors_created_at_index` (`created_at`),
  KEY `organizations_auditors_updated_at_index` (`updated_at`),
  CONSTRAINT `org_auditors_org_fk` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`),
  CONSTRAINT `org_auditors_user_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `organizations_auditors`
--

LOCK TABLES `organizations_auditors` WRITE;
/*!40000 ALTER TABLE `organizations_auditors` DISABLE KEYS */;
/*!40000 ALTER TABLE `organizations_auditors` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `organizations_billing_managers`
--

DROP TABLE IF EXISTS `organizations_billing_managers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `organizations_billing_managers` (
  `organization_id` int NOT NULL,
  `user_id` int NOT NULL,
  `organizations_billing_managers_pk` int NOT NULL AUTO_INCREMENT,
  `role_guid` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`organizations_billing_managers_pk`),
  UNIQUE KEY `org_billing_managers_idx` (`organization_id`,`user_id`),
  KEY `org_billing_managers_user_fk` (`user_id`),
  KEY `organizations_billing_managers_role_guid_index` (`role_guid`),
  KEY `organizations_billing_managers_created_at_index` (`created_at`),
  KEY `organizations_billing_managers_updated_at_index` (`updated_at`),
  CONSTRAINT `org_billing_managers_org_fk` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`),
  CONSTRAINT `org_billing_managers_user_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `organizations_billing_managers`
--

LOCK TABLES `organizations_billing_managers` WRITE;
/*!40000 ALTER TABLE `organizations_billing_managers` DISABLE KEYS */;
/*!40000 ALTER TABLE `organizations_billing_managers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `organizations_isolation_segments`
--

DROP TABLE IF EXISTS `organizations_isolation_segments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `organizations_isolation_segments` (
  `organization_guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `isolation_segment_guid` varchar(255) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`organization_guid`,`isolation_segment_guid`),
  KEY `fk_isolation_segments_guid` (`isolation_segment_guid`),
  CONSTRAINT `fk_isolation_segments_guid` FOREIGN KEY (`isolation_segment_guid`) REFERENCES `isolation_segments` (`guid`),
  CONSTRAINT `fk_organization_guid` FOREIGN KEY (`organization_guid`) REFERENCES `organizations` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `organizations_isolation_segments`
--

LOCK TABLES `organizations_isolation_segments` WRITE;
/*!40000 ALTER TABLE `organizations_isolation_segments` DISABLE KEYS */;
/*!40000 ALTER TABLE `organizations_isolation_segments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `organizations_managers`
--

DROP TABLE IF EXISTS `organizations_managers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `organizations_managers` (
  `organization_id` int NOT NULL,
  `user_id` int NOT NULL,
  `organizations_managers_pk` int NOT NULL AUTO_INCREMENT,
  `role_guid` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`organizations_managers_pk`),
  UNIQUE KEY `org_managers_idx` (`organization_id`,`user_id`),
  KEY `org_managers_user_fk` (`user_id`),
  KEY `organizations_managers_role_guid_index` (`role_guid`),
  KEY `organizations_managers_created_at_index` (`created_at`),
  KEY `organizations_managers_updated_at_index` (`updated_at`),
  CONSTRAINT `org_managers_org_fk` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`),
  CONSTRAINT `org_managers_user_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `organizations_managers`
--

LOCK TABLES `organizations_managers` WRITE;
/*!40000 ALTER TABLE `organizations_managers` DISABLE KEYS */;
/*!40000 ALTER TABLE `organizations_managers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `organizations_private_domains`
--

DROP TABLE IF EXISTS `organizations_private_domains`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `organizations_private_domains` (
  `organization_id` int NOT NULL,
  `private_domain_id` int NOT NULL,
  `organizations_private_domains_pk` int NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`organizations_private_domains_pk`),
  UNIQUE KEY `orgs_pd_ids` (`organization_id`,`private_domain_id`),
  KEY `fk_private_domain_id` (`private_domain_id`),
  CONSTRAINT `fk_organization_id` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`),
  CONSTRAINT `fk_private_domain_id` FOREIGN KEY (`private_domain_id`) REFERENCES `domains` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `organizations_private_domains`
--

LOCK TABLES `organizations_private_domains` WRITE;
/*!40000 ALTER TABLE `organizations_private_domains` DISABLE KEYS */;
/*!40000 ALTER TABLE `organizations_private_domains` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `organizations_users`
--

DROP TABLE IF EXISTS `organizations_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `organizations_users` (
  `organization_id` int NOT NULL,
  `user_id` int NOT NULL,
  `organizations_users_pk` int NOT NULL AUTO_INCREMENT,
  `role_guid` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`organizations_users_pk`),
  UNIQUE KEY `org_users_idx` (`organization_id`,`user_id`),
  KEY `org_users_user_fk` (`user_id`),
  KEY `organizations_users_role_guid_index` (`role_guid`),
  KEY `organizations_users_created_at_index` (`created_at`),
  KEY `organizations_users_updated_at_index` (`updated_at`),
  CONSTRAINT `org_users_org_fk` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`),
  CONSTRAINT `org_users_user_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `organizations_users`
--

LOCK TABLES `organizations_users` WRITE;
/*!40000 ALTER TABLE `organizations_users` DISABLE KEYS */;
/*!40000 ALTER TABLE `organizations_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `orphaned_blobs`
--

DROP TABLE IF EXISTS `orphaned_blobs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `orphaned_blobs` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `blob_key` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `dirty_count` int DEFAULT NULL,
  `blobstore_type` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `orphaned_blobs_guid_index` (`guid`),
  UNIQUE KEY `orphaned_blobs_unique_blob_index` (`blob_key`,`blobstore_type`),
  KEY `orphaned_blobs_created_at_index` (`created_at`),
  KEY `orphaned_blobs_updated_at_index` (`updated_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orphaned_blobs`
--

LOCK TABLES `orphaned_blobs` WRITE;
/*!40000 ALTER TABLE `orphaned_blobs` DISABLE KEYS */;
/*!40000 ALTER TABLE `orphaned_blobs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `package_annotations`
--

DROP TABLE IF EXISTS `package_annotations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `package_annotations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key` varchar(1000) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(5000) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `package_annotations_guid_index` (`guid`),
  KEY `package_annotations_created_at_index` (`created_at`),
  KEY `package_annotations_updated_at_index` (`updated_at`),
  KEY `fk_package_annotations_resource_guid_index` (`resource_guid`),
  CONSTRAINT `fk_package_annotations_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `packages` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `package_annotations`
--

LOCK TABLES `package_annotations` WRITE;
/*!40000 ALTER TABLE `package_annotations` DISABLE KEYS */;
/*!40000 ALTER TABLE `package_annotations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `package_labels`
--

DROP TABLE IF EXISTS `package_labels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `package_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key_name` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `package_labels_guid_index` (`guid`),
  KEY `package_labels_created_at_index` (`created_at`),
  KEY `package_labels_updated_at_index` (`updated_at`),
  KEY `fk_package_labels_resource_guid_index` (`resource_guid`),
  KEY `package_labels_compound_index` (`key_prefix`,`key_name`,`value`),
  CONSTRAINT `fk_package_labels_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `packages` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `package_labels`
--

LOCK TABLES `package_labels` WRITE;
/*!40000 ALTER TABLE `package_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `package_labels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `packages`
--

DROP TABLE IF EXISTS `packages`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `packages` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `type` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `package_hash` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `state` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `error` text CHARACTER SET utf8 COLLATE utf8_bin,
  `app_guid` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `docker_image` text,
  `sha256_checksum` varchar(255) DEFAULT NULL,
  `docker_username` varchar(255) DEFAULT NULL,
  `docker_password_salt` varchar(255) DEFAULT NULL,
  `encrypted_docker_password` varchar(16000) DEFAULT NULL,
  `encryption_key_label` varchar(255) DEFAULT NULL,
  `encryption_iterations` int NOT NULL DEFAULT '2048',
  PRIMARY KEY (`id`),
  UNIQUE KEY `packages_guid_index` (`guid`),
  KEY `packages_created_at_index` (`created_at`),
  KEY `packages_updated_at_index` (`updated_at`),
  KEY `packages_type_index` (`type`),
  KEY `package_app_guid_index` (`app_guid`),
  KEY `packages_state_index` (`state`),
  CONSTRAINT `fk_packages_app_guid` FOREIGN KEY (`app_guid`) REFERENCES `apps` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `packages`
--

LOCK TABLES `packages` WRITE;
/*!40000 ALTER TABLE `packages` DISABLE KEYS */;
/*!40000 ALTER TABLE `packages` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `process_annotations`
--

DROP TABLE IF EXISTS `process_annotations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `process_annotations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key` varchar(1000) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(5000) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `process_annotations_guid_index` (`guid`),
  KEY `process_annotations_created_at_index` (`created_at`),
  KEY `process_annotations_updated_at_index` (`updated_at`),
  KEY `fk_process_annotations_resource_guid_index` (`resource_guid`),
  CONSTRAINT `fk_process_annotations_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `processes` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `process_annotations`
--

LOCK TABLES `process_annotations` WRITE;
/*!40000 ALTER TABLE `process_annotations` DISABLE KEYS */;
/*!40000 ALTER TABLE `process_annotations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `process_labels`
--

DROP TABLE IF EXISTS `process_labels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `process_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key_name` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `process_labels_guid_index` (`guid`),
  KEY `process_labels_created_at_index` (`created_at`),
  KEY `process_labels_updated_at_index` (`updated_at`),
  KEY `fk_process_labels_resource_guid_index` (`resource_guid`),
  KEY `process_labels_compound_index` (`key_prefix`,`key_name`,`value`),
  CONSTRAINT `fk_process_labels_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `processes` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `process_labels`
--

LOCK TABLES `process_labels` WRITE;
/*!40000 ALTER TABLE `process_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `process_labels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `processes`
--

DROP TABLE IF EXISTS `processes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `processes` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `production` tinyint(1) DEFAULT '0',
  `memory` int DEFAULT NULL,
  `instances` int DEFAULT '1',
  `file_descriptors` int DEFAULT '16384',
  `disk_quota` int DEFAULT '2048',
  `state` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT 'STOPPED',
  `version` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `metadata` varchar(4096) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '{}',
  `detected_buildpack` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `not_deleted` tinyint(1) DEFAULT '1',
  `health_check_timeout` int DEFAULT NULL,
  `diego` tinyint(1) DEFAULT '0',
  `package_updated_at` datetime DEFAULT NULL,
  `app_guid` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `type` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT 'web',
  `health_check_type` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT 'port',
  `command` varchar(4096) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `enable_ssh` tinyint(1) DEFAULT '0',
  `encrypted_docker_credentials_json` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `docker_salt` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `ports` text,
  `health_check_http_endpoint` text,
  `health_check_invocation_timeout` int DEFAULT NULL,
  `revision_guid` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `apps_guid_index` (`guid`),
  KEY `apps_created_at_index` (`created_at`),
  KEY `apps_updated_at_index` (`updated_at`),
  KEY `apps_diego_index` (`diego`),
  KEY `processes_app_guid_index` (`app_guid`),
  CONSTRAINT `fk_processes_app_guid` FOREIGN KEY (`app_guid`) REFERENCES `apps` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `processes`
--

LOCK TABLES `processes` WRITE;
/*!40000 ALTER TABLE `processes` DISABLE KEYS */;
/*!40000 ALTER TABLE `processes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `quota_definitions`
--

DROP TABLE IF EXISTS `quota_definitions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `quota_definitions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `non_basic_services_allowed` tinyint(1) NOT NULL,
  `total_services` int NOT NULL,
  `memory_limit` int NOT NULL,
  `total_routes` int NOT NULL,
  `instance_memory_limit` int NOT NULL DEFAULT '-1',
  `total_private_domains` int NOT NULL DEFAULT '-1',
  `app_instance_limit` int DEFAULT '-1',
  `app_task_limit` int DEFAULT '-1',
  `total_service_keys` int DEFAULT '-1',
  `total_reserved_route_ports` int DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `qd_guid_index` (`guid`),
  UNIQUE KEY `qd_name_index` (`name`),
  KEY `qd_created_at_index` (`created_at`),
  KEY `qd_updated_at_index` (`updated_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `quota_definitions`
--

LOCK TABLES `quota_definitions` WRITE;
/*!40000 ALTER TABLE `quota_definitions` DISABLE KEYS */;
/*!40000 ALTER TABLE `quota_definitions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `request_counts`
--

DROP TABLE IF EXISTS `request_counts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `request_counts` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `count` int DEFAULT '0',
  `valid_until` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `request_counts_user_guid_index` (`user_guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `request_counts`
--

LOCK TABLES `request_counts` WRITE;
/*!40000 ALTER TABLE `request_counts` DISABLE KEYS */;
/*!40000 ALTER TABLE `request_counts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `revision_annotations`
--

DROP TABLE IF EXISTS `revision_annotations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `revision_annotations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key` varchar(1000) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(5000) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `revision_annotations_guid_index` (`guid`),
  KEY `revision_annotations_created_at_index` (`created_at`),
  KEY `revision_annotations_updated_at_index` (`updated_at`),
  KEY `fk_revision_annotations_resource_guid_index` (`resource_guid`),
  CONSTRAINT `fk_revision_annotations_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `revisions` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `revision_annotations`
--

LOCK TABLES `revision_annotations` WRITE;
/*!40000 ALTER TABLE `revision_annotations` DISABLE KEYS */;
/*!40000 ALTER TABLE `revision_annotations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `revision_labels`
--

DROP TABLE IF EXISTS `revision_labels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `revision_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key_name` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `revision_labels_guid_index` (`guid`),
  KEY `revision_labels_created_at_index` (`created_at`),
  KEY `revision_labels_updated_at_index` (`updated_at`),
  KEY `fk_revision_labels_resource_guid_index` (`resource_guid`),
  KEY `revision_labels_compound_index` (`key_prefix`,`key_name`,`value`),
  CONSTRAINT `fk_revision_labels_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `revisions` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `revision_labels`
--

LOCK TABLES `revision_labels` WRITE;
/*!40000 ALTER TABLE `revision_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `revision_labels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `revision_process_commands`
--

DROP TABLE IF EXISTS `revision_process_commands`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `revision_process_commands` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `revision_guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `process_type` varchar(255) COLLATE utf8_bin NOT NULL,
  `process_command` varchar(4096) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `revision_process_commands_guid_index` (`guid`),
  KEY `revision_process_commands_created_at_index` (`created_at`),
  KEY `revision_process_commands_updated_at_index` (`updated_at`),
  KEY `rev_commands_revision_guid_index` (`revision_guid`),
  CONSTRAINT `rev_commands_revision_guid_fkey` FOREIGN KEY (`revision_guid`) REFERENCES `revisions` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `revision_process_commands`
--

LOCK TABLES `revision_process_commands` WRITE;
/*!40000 ALTER TABLE `revision_process_commands` DISABLE KEYS */;
/*!40000 ALTER TABLE `revision_process_commands` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `revision_sidecar_process_types`
--

DROP TABLE IF EXISTS `revision_sidecar_process_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `revision_sidecar_process_types` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `type` varchar(255) COLLATE utf8_bin NOT NULL,
  `revision_sidecar_guid` varchar(255) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `revision_sidecar_process_types_guid_index` (`guid`),
  KEY `revision_sidecar_process_types_created_at_index` (`created_at`),
  KEY `revision_sidecar_process_types_updated_at_index` (`updated_at`),
  KEY `fk_revision_sidecar_proc_type_sidecar_guid_index` (`revision_sidecar_guid`),
  CONSTRAINT `fk_revision_sidecar_proc_type_sidecar_guid` FOREIGN KEY (`revision_sidecar_guid`) REFERENCES `revision_sidecars` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `revision_sidecar_process_types`
--

LOCK TABLES `revision_sidecar_process_types` WRITE;
/*!40000 ALTER TABLE `revision_sidecar_process_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `revision_sidecar_process_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `revision_sidecars`
--

DROP TABLE IF EXISTS `revision_sidecars`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `revision_sidecars` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) COLLATE utf8_bin NOT NULL,
  `command` varchar(4096) COLLATE utf8_bin NOT NULL,
  `revision_guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `memory` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `revision_sidecars_guid_index` (`guid`),
  KEY `revision_sidecars_created_at_index` (`created_at`),
  KEY `revision_sidecars_updated_at_index` (`updated_at`),
  KEY `fk_sidecar_revision_guid_index` (`revision_guid`),
  CONSTRAINT `fk_sidecar_revision_guid` FOREIGN KEY (`revision_guid`) REFERENCES `revisions` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `revision_sidecars`
--

LOCK TABLES `revision_sidecars` WRITE;
/*!40000 ALTER TABLE `revision_sidecars` DISABLE KEYS */;
/*!40000 ALTER TABLE `revision_sidecars` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `revisions`
--

DROP TABLE IF EXISTS `revisions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `revisions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `app_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `version` int DEFAULT '1',
  `droplet_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `encrypted_environment_variables` text COLLATE utf8_bin,
  `salt` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `encryption_key_label` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `encryption_iterations` int NOT NULL DEFAULT '2048',
  `description` text COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `revisions_guid_index` (`guid`),
  KEY `revisions_created_at_index` (`created_at`),
  KEY `revisions_updated_at_index` (`updated_at`),
  KEY `fk_revision_app_guid_index` (`app_guid`),
  CONSTRAINT `fk_revision_app_guid` FOREIGN KEY (`app_guid`) REFERENCES `apps` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `revisions`
--

LOCK TABLES `revisions` WRITE;
/*!40000 ALTER TABLE `revisions` DISABLE KEYS */;
/*!40000 ALTER TABLE `revisions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `route_annotations`
--

DROP TABLE IF EXISTS `route_annotations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `route_annotations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key` varchar(1000) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(5000) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `route_annotations_guid_index` (`guid`),
  KEY `route_annotations_created_at_index` (`created_at`),
  KEY `route_annotations_updated_at_index` (`updated_at`),
  KEY `fk_route_annotations_resource_guid_index` (`resource_guid`),
  CONSTRAINT `fk_route_annotations_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `routes` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `route_annotations`
--

LOCK TABLES `route_annotations` WRITE;
/*!40000 ALTER TABLE `route_annotations` DISABLE KEYS */;
/*!40000 ALTER TABLE `route_annotations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `route_binding_annotations`
--

DROP TABLE IF EXISTS `route_binding_annotations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `route_binding_annotations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key` varchar(1000) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(5000) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `route_binding_annotations_guid_index` (`guid`),
  KEY `route_binding_annotations_created_at_index` (`created_at`),
  KEY `route_binding_annotations_updated_at_index` (`updated_at`),
  KEY `fk_route_binding_annotations_resource_guid_index` (`resource_guid`),
  CONSTRAINT `fk_route_binding_annotations_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `route_bindings` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `route_binding_annotations`
--

LOCK TABLES `route_binding_annotations` WRITE;
/*!40000 ALTER TABLE `route_binding_annotations` DISABLE KEYS */;
/*!40000 ALTER TABLE `route_binding_annotations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `route_binding_labels`
--

DROP TABLE IF EXISTS `route_binding_labels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `route_binding_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key_name` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `route_binding_labels_guid_index` (`guid`),
  KEY `route_binding_labels_created_at_index` (`created_at`),
  KEY `route_binding_labels_updated_at_index` (`updated_at`),
  KEY `fk_route_binding_labels_resource_guid_index` (`resource_guid`),
  KEY `route_binding_labels_compound_index` (`key_prefix`,`key_name`,`value`),
  CONSTRAINT `fk_route_binding_labels_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `route_bindings` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `route_binding_labels`
--

LOCK TABLES `route_binding_labels` WRITE;
/*!40000 ALTER TABLE `route_binding_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `route_binding_labels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `route_binding_operations`
--

DROP TABLE IF EXISTS `route_binding_operations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `route_binding_operations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `route_binding_id` int DEFAULT NULL,
  `state` varchar(255) COLLATE utf8_bin NOT NULL,
  `type` varchar(255) COLLATE utf8_bin NOT NULL,
  `description` varchar(10000) COLLATE utf8_bin DEFAULT NULL,
  `broker_provided_operation` varchar(10000) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `route_binding_id_index` (`route_binding_id`),
  KEY `route_binding_operations_created_at_index` (`created_at`),
  KEY `route_binding_operations_updated_at_index` (`updated_at`),
  CONSTRAINT `fk_route_binding_op_route_binding_id` FOREIGN KEY (`route_binding_id`) REFERENCES `route_bindings` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `route_binding_operations`
--

LOCK TABLES `route_binding_operations` WRITE;
/*!40000 ALTER TABLE `route_binding_operations` DISABLE KEYS */;
/*!40000 ALTER TABLE `route_binding_operations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `route_bindings`
--

DROP TABLE IF EXISTS `route_bindings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `route_bindings` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `route_id` int DEFAULT NULL,
  `service_instance_id` int DEFAULT NULL,
  `route_service_url` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `route_bindings_guid_index` (`guid`),
  KEY `route_id` (`route_id`),
  KEY `service_instance_id` (`service_instance_id`),
  KEY `route_bindings_created_at_index` (`created_at`),
  KEY `route_bindings_updated_at_index` (`updated_at`),
  CONSTRAINT `route_bindings_ibfk_1` FOREIGN KEY (`route_id`) REFERENCES `routes` (`id`),
  CONSTRAINT `route_bindings_ibfk_2` FOREIGN KEY (`service_instance_id`) REFERENCES `service_instances` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `route_bindings`
--

LOCK TABLES `route_bindings` WRITE;
/*!40000 ALTER TABLE `route_bindings` DISABLE KEYS */;
/*!40000 ALTER TABLE `route_bindings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `route_labels`
--

DROP TABLE IF EXISTS `route_labels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `route_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key_name` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `route_labels_guid_index` (`guid`),
  KEY `route_labels_created_at_index` (`created_at`),
  KEY `route_labels_updated_at_index` (`updated_at`),
  KEY `fk_route_labels_resource_guid_index` (`resource_guid`),
  KEY `route_labels_compound_index` (`key_prefix`,`key_name`,`value`),
  CONSTRAINT `fk_route_labels_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `routes` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `route_labels`
--

LOCK TABLES `route_labels` WRITE;
/*!40000 ALTER TABLE `route_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `route_labels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `route_mappings`
--

DROP TABLE IF EXISTS `route_mappings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `route_mappings` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `app_port` int DEFAULT '-1',
  `guid` varchar(255) NOT NULL,
  `app_guid` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `route_guid` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `process_type` varchar(255) DEFAULT NULL,
  `weight` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `apps_routes_guid_index` (`guid`),
  UNIQUE KEY `route_mappings_app_guid_route_guid_process_type_app_port_key` (`app_guid`,`route_guid`,`process_type`,`app_port`),
  KEY `apps_routes_created_at_index` (`created_at`),
  KEY `apps_routes_updated_at_index` (`updated_at`),
  KEY `route_mappings_process_type_index` (`process_type`),
  KEY `fk_route_mappings_route_guid` (`route_guid`),
  KEY `route_mappings_app_guid_index` (`app_guid`),
  CONSTRAINT `fk_route_mappings_app_guid` FOREIGN KEY (`app_guid`) REFERENCES `apps` (`guid`),
  CONSTRAINT `fk_route_mappings_route_guid` FOREIGN KEY (`route_guid`) REFERENCES `routes` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `route_mappings`
--

LOCK TABLES `route_mappings` WRITE;
/*!40000 ALTER TABLE `route_mappings` DISABLE KEYS */;
/*!40000 ALTER TABLE `route_mappings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `routes`
--

DROP TABLE IF EXISTS `routes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `routes` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `host` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `domain_id` int NOT NULL,
  `space_id` int NOT NULL,
  `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `port` int NOT NULL DEFAULT '0',
  `vip_offset` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `routes_guid_index` (`guid`),
  UNIQUE KEY `routes_host_domain_id_path_port_index` (`host`,`domain_id`,`path`,`port`),
  UNIQUE KEY `routes_vip_offset_index` (`vip_offset`),
  KEY `fk_routes_domain_id` (`domain_id`),
  KEY `fk_routes_space_id` (`space_id`),
  KEY `routes_created_at_index` (`created_at`),
  KEY `routes_updated_at_index` (`updated_at`),
  CONSTRAINT `fk_routes_domain_id` FOREIGN KEY (`domain_id`) REFERENCES `domains` (`id`),
  CONSTRAINT `fk_routes_space_id` FOREIGN KEY (`space_id`) REFERENCES `spaces` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `routes`
--

LOCK TABLES `routes` WRITE;
/*!40000 ALTER TABLE `routes` DISABLE KEYS */;
/*!40000 ALTER TABLE `routes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `schema_migrations`
--

DROP TABLE IF EXISTS `schema_migrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `schema_migrations` (
  `filename` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`filename`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `schema_migrations`
--

LOCK TABLES `schema_migrations` WRITE;
/*!40000 ALTER TABLE `schema_migrations` DISABLE KEYS */;
INSERT INTO `schema_migrations` VALUES ('20130131184954_new_initial_schema.rb'),('20130219194917_create_stacks_table.rb'),('20130313210735_add_extra_to_service.rb'),('20130314221003_add_extra_to_service_plans.rb'),('20130318175647_convert_mysql_text_types_for_services.rb'),('20130323005501_remove_framework_and_runtime.rb'),('20130328225143_add_unique_id_to_services.rb'),('20130329165619_add_unique_id_to_service_plans.rb'),('20130405004709_add_detected_buildpack_to_app.rb'),('20130408215903_add_staging_task_id.rb'),('20130425230435_add_crash_events.rb'),('20130426001853_add_public_to_service_plans.rb'),('20130426175420_add_can_access_non_public_plans_to_orgs.rb'),('20130429225754_add_salt_to_service_bindings.rb'),('20130430184538_add_salt_to_service_auth_tokens.rb'),('20130430205654_add_salt_to_service_instances.rb'),('20130501002952_use_text_fields_for_service_credentials.rb'),('20130501224802_rename_crash_events_to_app_events.rb'),('20130503211948_add_rds_column_to_quota.rb'),('20130508214259_rename_free_rds_to_trial_db_allowed_in_quota_definition.rb'),('20130508220251_add_dashboard_url_to_service_instances.rb'),('20130509212940_convert_mysql_type_for_environment_json.rb'),('20130515213335_add_kill_after_multiple_restarts_to_apps.rb'),('20130522233916_add_deleted_column_to_apps.rb'),('20130529183520_add_not_deleted_column.rb'),('20130530173342_fix_permission_fk_problems.rb'),('20130530210927_add_status_to_organization.rb'),('20130603223751_convert_latin1_to_utf8_for_mysql.rb'),('20130611175239_create_service_plan_visibilities_table.rb'),('20130611235653_drop_organizations_can_access_non_public_plans.rb'),('20130621222741_allow_null_service_plan_and_add_kind_on_service_instances.rb'),('20130628184803_delete_kind_from_service_instances.rb'),('20130628233316_add_is_gateway_service_to_service_instances.rb'),('20130712215250_limit_service_instance_name_to_50_chars.rb'),('20130718003848_add_index_to_service_instances_on_gateway_name.rb'),('20130718210849_add_tasks.rb'),('20130718234910_add_bindable_to_services.rb'),('20130719002215_remove_not_null_from_unique_id.rb'),('20130719214629_add_salt_and_encrypted_env_json_to_app_table.rb'),('20130725213922_create_events_table.rb'),('20130730000217_create_service_brokers_table.rb'),('20130730000318_add_tags_to_service.rb'),('20130730214030_lengthen_encrypted_environment_variables.rb'),('20130806175100_support_30char_identifiers_for_oracle.rb'),('20130814211636_add_secure_token_to_tasks.rb'),('20130815225217_add_documentation_url_to_services.rb'),('20130821333222_create_buildpacks_table.rb'),('20130823211228_add_indices_to_events.rb'),('20130824002024_add_broker_id_to_services_and_nullable_provider_version_url.rb'),('20130826210728_add_salt_to_service_brokers.rb'),('20130829221542_add_unique_index_on_service_plans_unique_id.rb'),('20130910221313_add_syslog_drain_url_to_bindings.rb'),('20130911111938_encrypt_app_env_json.rb'),('20130911220131_add_long_description_to_services.rb'),('20130913165030_add_requires_to_services.rb'),('20130919164043_denormalize_space_and_add_org_guid_to_events.rb'),('20130919211247_add_sequel_delayed_job.rb'),('20130927215822_add_admin_buildpack_id_to_apps.rb'),('20131002170606_create_droplets.rb'),('20131003173318_service_brokers_have_usernames_and_passwords.rb'),('20131016215721_rename_priority_to_position.rb'),('20131018002858_add_syslog_to_service_instances.rb'),('20131018171757_add_active_to_service_plans.rb'),('20131022223542_service_instance_credentials_are_optional.rb'),('20131024203451_use_text_for_metadata_in_events.rb'),('20131104223608_fix_events_timestamps_and_truncate.rb'),('20131107223211_add_enabled_to_buildpack.rb'),('20131111155640_add_health_check_timeout_to_apps.rb'),('20131114185320_add_total_routes_to_quota_definitions.rb'),('20131119225844_move_organization_managers_into_organization.rb'),('20131212222630_drop_domains_spaces_table.rb'),('20131213232039_remove_app_memory_default.rb'),('20131217175551_create_app_usage_events.rb'),('20140114175047_add_locked_to_buildpack.rb'),('20140115094157_add_filename_to_buildpacks.rb'),('20140130215438_change_services_provider_default.rb'),('20140204221153_add_purging_to_services.rb'),('20140206190102_add_dashboard_client_id_to_services.rb'),('20140208001107_reset_service_provider_default_value_for_mysql.rb'),('20140211223904_drop_domains_organizations.rb'),('20140218201113_add_cf_api_error_to_delayed_jobs.rb'),('20140220184651_remove_trial_db_allowed.rb'),('20140227191931_add_actor_actee_names_to_events.rb'),('20140305230940_create_service_dashboard_client.rb'),('20140306000008_remove_sso_client_id_from_services.rb'),('20140307201639_replace_service_id_on_broker_with_service_broker_id.rb'),('20140319191826_add_buildpack_info_to_app_usage_events.rb'),('20140324172951_add_detected_buildpack_guid_to_app.rb'),('20140325171355_add_service_broker_id_index_to_service_dashboard_clients.rb'),('20140402183459_change_app_instances_default.rb'),('20140407230239_add_detected_buildpack_name_to_apps.rb'),('20140512175050_add_service_usage_event.rb'),('20140514210916_add_detected_command_to_droplet.rb'),('20140515155207_add_staging_failed_reason_to_apps.rb'),('20140528174243_add_app_security_groups.rb'),('20140603184619_add_app_security_group_spaces_join_table.rb'),('20140609180716_add_staging_default_to_app_security_groups.rb'),('20140609234412_add_running_default_to_app_security_groups.rb'),('20140623205358_make_name_unique_for_app_security_groups.rb'),('20140624232412_change_app_security_groups_to_security_groups.rb'),('20140707180618_truncate_securitygroups.rb'),('20140708223526_drop_tasks.rb'),('20140716213753_add_actee_index_to_events.rb'),('20140721225153_add_diego_flag_to_apps.rb'),('20140723172206_add_instance_memory_limit_to_quota_definitions.rb'),('20140723212942_add_space_quota_definitions.rb'),('20140724185938_add_space_quota_definition_to_space.rb'),('20140724215343_add_feature_flags.rb'),('20140805223232_add_docker_image.rb'),('20140811174704_add_error_message_to_feature_flags.rb'),('20140819094032_add_environment_variable_groups.rb'),('20140826170851_add_salt_and_encrypt_env_group_vars.rb'),('20140827202612_add_package_updated_at_to_apps.rb'),('20140908220352_add_execution_data_to_droplets.rb'),('20140929212559_change_droplet_cols_to_text.rb'),('20141008205150_create_lockings.rb'),('20141022211551_add_updateable_column_to_services.rb'),('20141029182220_create_apps_v3.rb'),('20141030213445_add_app_guid_to_processes.rb'),('20141105173044_add_package_pending_since_to_apps.rb'),('20141120182308_reprocess_diego_apps.rb'),('20141120193405_add_type_to_apps.rb'),('20141204175821_add_package_state_to_app_usage_events.rb'),('20141204180212_add_package_state_to_app_usage_events.rb'),('20141210194308_add_name_col_v3_apps.rb'),('20141212001210_update_delayed_jobs_indexes.rb'),('20141215184824_add_locked_by_to_delayed_jobs_index.rb'),('20141216183550_add_health_check_type_to_apps.rb'),('20141226222846_create_packages.rb'),('20150113000312_add_state_and_state_description_to_service_instances.rb'),('20150113201824_fix_created_at_column.rb'),('20150122183634_create_v3_droplets.rb'),('20150122222844_add_app_guid_to_package.rb'),('20150124164257_add_organizations_private_domains.rb'),('20150127013821_add_events_index.rb'),('20150127194942_create_service_instance_operations.rb'),('20150130181135_grow_metadata_column.rb'),('20150202222406_revert_accidental_migration_on_service_instances.rb'),('20150202222559_make_metadata_column_long_varchar_on_apps.rb'),('20150202222600_add_command_to_processes.rb'),('20150204012903_drop_state_and_state_description.rb'),('20150204222823_add_proposed_changes_to_service_instance_operations.rb'),('20150211002332_add_desired_droplet_guid_to_v3_apps.rb'),('20150211010713_add_app_guid_to_v3_droplets.rb'),('20150211182540_add_failure_reason_to_droplet.rb'),('20150211190122_add_detected_start_command_to_droplet.rb'),('20150213001251_add_index_to_diego_app_flag.rb'),('20150218192154_remove_space_guid_from_v3_packages.rb'),('20150220010146_remove_kill_after_multiple_restarts_from_apps.rb'),('20150306233007_increase_size_of_delayed_job_handler.rb'),('20150311204445_add_desired_state_to_v3_apps.rb'),('20150313233039_create_apps_v3_routes.rb'),('20150316184259_create_service_key_table.rb'),('20150318185941_add_encrypted_environment_variables_to_apps_v3.rb'),('20150319150641_add_encrypted_environment_variables_to_v3_droplets.rb'),('20150323170053_change_service_instance_description_to_text.rb'),('20150323234355_recreate_apps_v3_routes.rb'),('20150324232809_add_fk_v3_apps_packages_droplets_processes.rb'),('20150325224808_add_v3_attrs_to_app_usage_events.rb'),('20150327080540_add_cached_docker_image_to_droplets.rb'),('20150403175058_add_index_to_droplets_droplet_hash.rb'),('20150403190653_add_procfile_to_droplets.rb'),('20150407213536_add_index_to_stack_id.rb'),('20150414113235_remove_space_id.rb'),('20150421190248_add_allow_ssh_to_app.rb'),('20150422000255_route_path_field.rb'),('20150430214950_add_allow_ssh_to_spaces.rb'),('20150501181106_rename_apps_allow_ssh_to_enable_ssh.rb'),('20150511134747_add_docker_image_credentials.rb'),('20150514190458_fix_mysql_collations.rb'),('20150515230939_add_case_insensitive_to_route_path.rb'),('20150521205906_add_tags_to_service_instance.rb'),('20150528183140_increase_service_instance_tags_column_length.rb'),('20150618190133_add_staging_failed_description_to_app.rb'),('20150623175237_add_total_private_domains_to_quota_definitions.rb'),('20150625171651_change_stack_description_to_allow_null.rb'),('20150625213234_create_service_instance_dashboard_clients.rb'),('20150626000000_add_buildpack_to_v3_app.rb'),('20150626231641_v3_droplets_generalize_buildpack_field.rb'),('20150702004354_add_service_broker_id_to_space.rb'),('20150707202136_rename_desired_droplet_guid_to_droplet_guid.rb'),('20150709165719_truncate_billing_events_table.rb'),('20150710171553_add_app_instance_limit_to_quota_definition.rb'),('20150713133551_enlarge_service_keys_credentials.rb'),('20150720182530_v3_droplets_error_field.rb'),('20150817175317_add_route_service_url_to_service_instance.rb'),('20150819230845_add_service_instance_id_to_route.rb'),('20150821153140_remove_service_instance_dashboard.rb'),('20150827235417_remove_route_service_url_from_service_instances.rb'),('20150904215710_create_route_binding_table.rb'),('20150910221617_drop_routes_service_instance_id_column.rb'),('20150910221699_add_events_index.rb'),('20150916213812_add_route_service_url_to_route_binding.rb'),('20150928221442_add_router_group_guid_to_shared_domains.rb'),('20150928231352_add_app_instance_limit_to_space_quota_definition.rb'),('20151006170705_add_port_column_to_routes_table.rb'),('20151006224020_add_route_service_url_to_user_provided_service_instances.rb'),('20151008211628_add_lifecycle_info_to_droplet.rb'),('20151013004418_add_lifecycle_info_to_v3_app.rb'),('20151016173720_drop_app_lifecycle_json.rb'),('20151016173728_drop_droplet_lifecycle_json.rb'),('20151016173751_create_buildpack_lifecycle_data.rb'),('20151026231841_add_ports_to_v2_app.rb'),('20151028210802_drop_buildpack_from_v3_app_model.rb'),('20151102175640_create_package_docker_data.rb'),('20151110215644_rename_buildpack_and_stack_in_droplet_model.rb'),('20151119192340_add_docker_receipt_to_droplet.rb'),('20151123232837_use_text_fields_for_service_tags.rb'),('20151124013630_increase_text_size_for_service_instances_tags.rb'),('20151217235335_remove_unused_package_cols.rb'),('20151222182812_create_v3_service_bindings.rb'),('20151231224207_increase_security_group_rules_length.rb'),('20160113223027_alter_apps_routes_table.rb'),('20160114182148_create_tasks.rb'),('20160127192643_add_result_to_tasks.rb'),('20160127195206_add_environment_variables_to_tasks.rb'),('20160129195310_add_memory_in_mb_to_tasks.rb'),('20160203011824_create_route_mappings_table.rb'),('20160203013305_add_salt_to_tasks.rb'),('20160203223853_add_task_info_to_app_usage_event.rb'),('20160210184012_add_app_task_limit_to_quota_definitions.rb'),('20160210191133_add_app_task_limit_to_space_quota_definitions.rb'),('20160221152904_add_total_service_keys_to_quota_definitions.rb'),('20160221153023_add_total_service_keys_to_space_quota_definitions.rb'),('20160301215655_add_unique_constraint_on_apps_routes_table.rb'),('20160303002319_add_indices_for_filtered_task_columns.rb'),('20160303231742_add_package_guid_to_usage_events.rb'),('20160322184040_update_app_port.rb'),('20160328074419_add_previous_values_in_app_usage_events.rb'),('20160411172037_add_total_reserved_route_ports_to_quota_definitions.rb'),('20160416005940_add_total_reserved_route_ports_to_space_quota_definitions.rb'),('20160502214345_add_volume_mounts_to_service_binding.rb'),('20160502225745_add_volume_mounts_salt_to_service_binding.rb'),('20160504182904_increase_task_command_size.rb'),('20160504214134_increase_volume_mounts_size.rb'),('20160504232502_add_volume_mounts_to_v3_service_bindings.rb'),('20160510172035_increase_task_environment_variables_size.rb'),('20160517190429_rename_memory_limit_to_staging_memory_in_mb_for_v3_droplets.rb'),('20160523172247_rename_disk_limit_to_staging_disk_in_mb_for_v3_droplets.rb'),('20160601165902_add_broker_provided_operation_to_instance_last_operation.rb'),('20160601173727_add_port_to_route_mappings.rb'),('20160621182906_remove_space_id_from_events.rb'),('20160802210551_add_salt_and_encrypt_buildpack.rb'),('20160808182741_isolation_segment.rb'),('20160817214144_space_isolation_segment_association.rb'),('20160914165525_migrate_v2_app_data_to_v3.rb'),('20160919172325_remove_droplet_fk_constraint_for_tasks.rb'),('20160920170627_add_task_sequence_id.rb'),('20160920175805_add_app_model_max_task_sequence_id.rb'),('20160922164519_default_task_sequence_id_to_one.rb'),('20160922213611_organization_isolation_segments.rb'),('20160926221309_change_service_instance_dashboard_url_to_have_max_length.rb'),('20161005205815_fix_droplets_process_types_json.rb'),('20161006184718_create_request_counts.rb'),('20161006221433_add_valid_until_to_request_counts.rb'),('20161011141422_service_binding_volume_mounts_convert_to_2_10_format.rb'),('20161024221405_add_disk_in_mb_to_tasks.rb'),('20161028210215_add_default_updated_at.rb'),('20161104184720_create_staging_security_groups_spaces.rb'),('20161114113512_add_bindable_to_service_plans.rb'),('20161206005057_add_health_check_http_endpoint_to_processes.rb'),('20161215190126_add_sha56_checksum_to_droplet_model.rb'),('20170109174921_builpack_cache_sha256_checksum.rb'),('20170110195809_builpack_sha256_checksum.rb'),('20170111005234_packages_sha256_checksum.rb'),('20170201224823_add_username_to_events.rb'),('20170214000310_create_clock_jobs.rb'),('20170221232946_add_last_completed_at_to_clock_jobs.rb'),('20170303011654_add_app_guid_index_to_droplets.rb'),('20170303012525_add_app_guid_index_to_packages.rb'),('20170321205040_create_build_model.rb'),('20170411172945_add_error_description_to_builds.rb'),('20170412173354_add_docker_receipt_image_to_builds.rb'),('20170413213539_remove_service_instance_id_from_routes.rb'),('20170418185436_add_app_guid_to_builds.rb'),('20170420184502_add_error_id_to_build.rb'),('20170425173340_add_docker_credentials_to_packages.rb'),('20170502171127_remove_orphaned_apps.rb'),('20170502181209_add_fk_apps_space_guid.rb'),('20170504210922_add_docker_creds_to_droplet.rb'),('20170505163434_remove_unused_columns_from_builds.rb'),('20170505205924_add_app_guid_foreign_key_to_builds.rb'),('20170522233826_remove_duplicate_route_mapping_entries.rb'),('20170524214613_add_created_by_to_builds.rb'),('20170524232621_create_orphaned_blobs.rb'),('20170526225825_add_directory_key_to_orphaned_blobs.rb'),('20170601205215_add_tasks_app_guid_index.rb'),('20170602113045_add_instance_create_schema_to_service_plans.rb'),('20170605182822_create_jobs.rb'),('20170609172743_add_delayed_job_guid_to_job.rb'),('20170609220018_orphaned_blob_index_on_key_and_type.rb'),('20170620171034_add_index_for_buildpack_key.rb'),('20170620171244_add_compound_index_on_droplet_guid_and_droplet_hash.rb'),('20170630224921_create_buildpack_lifecycle_buildpacks.rb'),('20170712153045_add_instance_update_schema_to_service_plans.rb'),('20170717211406_add_cf_api_error_to_jobs.rb'),('20170719182326_add_app_guid_index_to_builds.rb'),('20170719182821_add_indices_to_service_bindings.rb'),('20170719183350_add_indices_to_service_usage_events.rb'),('20170720233439_change_encrypted_docker_password_on_packages.rb'),('20170721205940_add_missing_task_stopped_usage_events.rb'),('20170724090255_add_binding_create_schema_to_service_plans.rb'),('20170724170303_grow_service_instances_syslog_drain_url.rb'),('20170724173748_grow_service_bindings_syslog_drain_url.rb'),('20170802230125_add_missing_task_stopped_usage_events_second_attempt.rb'),('20170814212845_delete_unused_tables.rb'),('20170815190541_add_enable_ssh_to_apps.rb'),('20170815233431_migrate_enable_ssh_from_processes_to_apps.rb'),('20170920143711_create_service_instance_shares.rb'),('20171013223336_remove_deprecated_v1_service_fields.rb'),('20171103163351_add_name_to_service_bindings.rb'),('20171106202032_change_processes_with_health_check_timeout_0_to_nil.rb'),('20171120214253_remove_buildpack_receipt_stack_name_from_droplets.rb'),('20171123191651_add_index_to_service_binding_on_name.rb'),('20171220183100_add_encryption_key_label_column_to_tables_with_encrypted_columns.rb'),('20180115151922_add_instances_and_bindings_retrievable_to_services.rb'),('20180125181819_add_internal_to_domains.rb'),('20180220224558_add_missing_primary_keys.rb'),('20180315195737_add_index_to_builds_state.rb'),('20180319143620_create_service_binding_operations.rb'),('20180419180706_add_process_healthcheck_invocation_timeout.rb'),('20180420185709_add_version_bpname_to_buildpack_lifecycle_buildpacks.rb'),('20180424202908_create_deployments_table.rb'),('20180501171507_add_droplet_guid_to_deployments.rb'),('20180502232322_add_stack_to_buildpack.rb'),('20180502234947_add_index_to_packages_state.rb'),('20180515220732_create_primary_key_for_organizations_auditors.rb'),('20180515221609_create_primary_key_for_organizations_billing_managers.rb'),('20180515221623_create_primary_key_for_organizations_managers.rb'),('20180515221638_create_primary_key_for_organizations_private_domains.rb'),('20180515221652_create_primary_key_for_organizations_users.rb'),('20180515221706_create_primary_key_for_security_groups_spaces.rb'),('20180515221720_create_primary_key_for_spaces_auditors.rb'),('20180515221734_create_primary_key_for_spaces_developers.rb'),('20180515221748_create_primary_key_for_spaces_managers.rb'),('20180515221803_create_primary_key_for_staging_security_groups_spaces.rb'),('20180522211345_add_webish_process_to_deployments.rb'),('20180523205142_backfill_web_processes_for_v3_apps.rb'),('20180628223056_rename_webish_process_guid_on_deployments.rb'),('20180703233121_set_missing_fields_for_backfilled_web_processes.rb'),('20180710115626_change_broker_catalog_descriptions_to_type_text.rb'),('20180726120275_clear_process_command_for_buildpacks.rb'),('20180813181554_add_route_weight_to_route_mappings.rb'),('20180813221823_clear_process_command_and_metadata_command.rb'),('20180814184641_change_encrypted_docker_password_on_droplets.rb'),('20180828172307_add_previous_droplet_to_deployments.rb'),('20180904174127_add_original_web_process_instance_count_to_deployments.rb'),('20180904210247_add_index_to_deployments_on_state.rb'),('20180917222717_create_encryption_key_sentinels_table.rb'),('20180921102908_remove_uniqueness_on_index_broker_url_from_service_brokers.rb'),('20180924142348_remove_uniqueness_on_index_unique_id_from_services.rb'),('20180924142407_remove_uniqueness_on_index_unique_id_from_service_plans.rb'),('20180925150440_remove_services_label_provider_index.rb'),('20180927105539_add_broker_name_and_guid_to_service_usage_event.rb'),('20181002165615_create_deployment_processes_table.rb'),('20181015223531_create_app_labels.rb'),('20181030175334_create_org_labels.rb'),('20181031232313_create_space_labels.rb'),('20181101211039_drop_and_recreate_app_space_org_labels.rb'),('20181107191728_add_revisions.rb'),('20181109191715_add_revision_guid_to_deployments.rb'),('20181112143236_add_plan_updateable_to_service_plans.rb'),('20181112220156_create_app_annotations.rb'),('20181112222754_add_version_to_revisions.rb'),('20181120230039_increase_app_annotations_resource_guid_column_length.rb'),('20181120231620_create_organization_annotations_table.rb'),('20181127224109_add_vip_to_routes.rb'),('20181129180059_populate_vips.rb'),('20181204184506_create_space_annotations_table.rb'),('20181207184247_create_droplet_labels_table.rb'),('20181211230616_create_droplet_annotations_table.rb'),('20181219235732_create_package_labels_table.rb'),('20181219235744_create_package_annotations_table.rb'),('20181220234322_add_revisions_enabled_to_apps.rb'),('20190102192213_add_droplet_guid_to_revision.rb'),('20190109223722_create_stack_labels_table.rb'),('20190109223733_create_stack_annotations_table.rb'),('20190110224601_create_isolation_segment_metadata_tables.rb'),('20190110225015_create_task_annotations_table.rb'),('20190110225031_create_task_labels_table.rb'),('20190116193201_create_process_metadata_tables.rb'),('20190116220853_create_revision_labels_table.rb'),('20190116220909_create_revision_annotations_table.rb'),('20190116225537_create_deployment_metadata_tables.rb'),('20190123184851_add_environment_variables_to_revisions.rb'),('20190128183942_create_build_labels_table.rb'),('20190128183957_create_build_annotations_table.rb'),('20190128233032_create_buildpack_metadata_tables.rb'),('20190206191247_add_custom_commands_to_revisions.rb'),('20190219161111_add_maximum_polling_duration_to_service_plans.rb'),('20190220001829_add_encryption_iteration_count_to_apps.rb'),('20190220003558_add_encryption_iteration_count_to_buildpack_lifecycle_buildpacks.rb'),('20190220003646_add_encryption_iteration_count_to_buildpack_lifecycle_data.rb'),('20190220003740_add_encryption_iteration_count_to_droplets.rb'),('20190220003808_add_encryption_iteration_count_to_env_groups.rb'),('20190220003847_add_encryption_iteration_count_to_packages.rb'),('20190220003903_add_encryption_iteration_count_to_revisions.rb'),('20190220003916_add_encryption_iteration_count_to_tasks.rb'),('20190220004407_add_encryption_iteration_count_to_service_bindings.rb'),('20190220004428_add_encryption_iteration_count_to_service_brokers.rb'),('20190220004441_add_encryption_iteration_count_to_service_instances.rb'),('20190220004454_add_encryption_iteration_count_to_service_keys.rb'),('20190221221532_create_revision_process_commands_table.rb'),('20190222004908_remove_revisions_encrypted_commands_by_process_type.rb'),('20190222190051_add_encryption_iteration_count_to_encryption_key_sentinels.rb'),('20190223002950_remove_encryption_iteration_defaults.rb'),('20190227175600_add_allow_context_updates_to_services.rb'),('20190302003850_add_revision_description_column.rb'),('20190319220026_create_sidecars.rb'),('20190320234126_create_service_instance_labels_table.rb'),('20190320234146_create_service_instance_annotations_table.rb'),('20190321214408_change_sidecar_process_type_name_to_type.rb'),('20190326212736_add_app_guid_to_sidecar_process_types.rb'),('20190409110217_add_maintenance_info_to_service_plan.rb'),('20190417232206_add_is_client_to_users_table.rb'),('20190430171333_add_last_healthy_at_to_deployments.rb'),('20190502172509_create_domain_metadata_tables.rb'),('20190503090026_add_maintenance_info_to_service_instance.rb'),('20190531214252_create_route_metadata_tables.rb'),('20190617170734_add_key_prefix_to_package_annotations.rb'),('20190617175452_add_key_prefix_to_buildpack_annotations.rb'),('20190617181003_add_key_prefix_to_droplet_annotations.rb'),('20190617181111_add_key_prefix_to_app_annotations.rb'),('20190617181159_add_key_prefix_to_revision_annotations.rb'),('20190617181245_add_key_prefix_to_task_annotations.rb'),('20190617181325_add_key_prefix_to_process_annotations.rb'),('20190617181421_add_key_prefix_to_isolation_segment_annotations.rb'),('20190617181509_add_key_prefix_to_stack_annotations.rb'),('20190617181730_add_key_prefix_to_service_instance_annotations.rb'),('20190617181810_add_key_prefix_to_deployment_annotations.rb'),('20190617181849_add_key_prefix_to_domain_annotations.rb'),('20190617181924_add_key_prefix_to_space_annotations.rb'),('20190617182113_add_key_prefix_to_organization_annotations.rb'),('20190617182135_add_key_prefix_to_build_annotations.rb'),('20190617182203_add_key_prefix_to_route_annotations.rb'),('20190628183006_change_default_route_mapping_weight_to_null.rb'),('20190701233808_add_memory_column_to_sidecars.rb'),('20190711205517_add_value_and_reason_status_to_deployments.rb'),('20190712210940_backfill_status_for_deployments.rb'),('20190716215540_add_status_indices_to_deployments.rb'),('20190905131313_create_service_broker_states.rb'),('20190913184731_create_user_metadata_tables.rb'),('20190916184059_create_revision_sidecars_table.rb'),('20190923131251_create_job_warnings_table.rb'),('20190926205451_add_origin_to_sidecar.rb'),('20190930225152_add_sidecars_to_droplets.rb'),('20191014212939_add_guid_and_timestamps_to_roles_join_tables.rb'),('20191014233211_fill_guid_and_timestamps_for_roles_join_tables.rb'),('20191023211109_add_strategy_to_deployment.rb'),('20191031135500_create_service_broker_update_requests.rb'),('20191112121711_add_state_to_service_brokers_table.rb'),('20191112154745_drop_service_broker_states_table.rb'),('20191125192326_create_kpack_lifecycle_data.rb'),('20191203153701_create_service_broker_metadata_tables.rb'),('20191204144101_create_service_broker_update_request_metadata_tables.rb'),('20191218000612_add_guid_and_timestamps_to_spaces_auditors.rb'),('20191218000616_add_guid_and_timestamps_to_spaces_managers.rb'),('20191218000620_add_guid_and_timestamps_to_spaces_developers.rb'),('20191218000625_add_guid_and_timestamps_to_organizations_auditors.rb'),('20191218000629_add_guid_and_timestamps_to_organizations_billing_managers.rb'),('20191218000633_add_guid_and_timestamps_to_organizations_managers.rb'),('20191218000638_add_guid_and_timestamps_to_organizations_users.rb'),('20191218001006_fill_guid_and_timestamps_for_spaces_auditors.rb'),('20191218001010_fill_guid_and_timestamps_for_spaces_managers.rb'),('20191218001015_fill_guid_and_timestamps_for_spaces_developers.rb'),('20191218001019_fill_guid_and_timestamps_for_organizations_auditors.rb'),('20191218001024_fill_guid_and_timestamps_for_organizations_billing_managers.rb'),('20191218001028_fill_guid_and_timestamps_for_organizations_managers.rb'),('20191218001034_fill_guid_and_timestamps_for_organizations_users.rb'),('20200110000101_add_app_guid_index_to_route_mappings.rb'),('20200114202134_change_app_revisions_enabled_to_default_true.rb'),('20200121104723_create_service_offering_metadata.rb'),('20200210103523_create_service_plan_metadata.rb'),('20200211194856_change_revision_env_vars_to_text.rb'),('20200218223146_add_app_kpack_lifecycle_data.rb'),('20200507224709_change_crash_events_exit_description_to_text.rb'),('20200727211409_add_kpack_encrypted_buildpack_infos_column.rb'),('20200826090000_create_route_binding_operations.rb'),('20201013153721_add_service_instance_constraint_to_service_bindings.rb'),('20201020204056_index_kpack_lifecycle_data_droplet_guid.rb'),('20201030105232_create_service_route_bindings_metadata_table.rb'),('20201124153320_create_service_key_operations.rb'),('20210105151532_create_service_bindings_metadata_tables.rb'),('20210106120645_create_service_keys_metadata_tables.rb'),('20210427221541_create_spaces_application_supporters_table.rb'),('20210607213940_add_staging_memory_and_disk_to_builds.rb');
/*!40000 ALTER TABLE `schema_migrations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `security_groups`
--

DROP TABLE IF EXISTS `security_groups`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `security_groups` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) COLLATE utf8_bin NOT NULL,
  `rules` mediumtext COLLATE utf8_bin,
  `staging_default` tinyint(1) DEFAULT '0',
  `running_default` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `sg_guid_index` (`guid`),
  KEY `sg_created_at_index` (`created_at`),
  KEY `sg_updated_at_index` (`updated_at`),
  KEY `sgs_staging_default_index` (`staging_default`),
  KEY `sgs_running_default_index` (`running_default`),
  KEY `sg_name_index` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `security_groups`
--

LOCK TABLES `security_groups` WRITE;
/*!40000 ALTER TABLE `security_groups` DISABLE KEYS */;
/*!40000 ALTER TABLE `security_groups` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `security_groups_spaces`
--

DROP TABLE IF EXISTS `security_groups_spaces`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `security_groups_spaces` (
  `security_group_id` int NOT NULL,
  `space_id` int NOT NULL,
  `security_groups_spaces_pk` int NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`security_groups_spaces_pk`),
  KEY `fk_space_id` (`space_id`),
  KEY `sgs_spaces_ids` (`security_group_id`,`space_id`),
  CONSTRAINT `fk_security_group_id` FOREIGN KEY (`security_group_id`) REFERENCES `security_groups` (`id`),
  CONSTRAINT `fk_space_id` FOREIGN KEY (`space_id`) REFERENCES `spaces` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `security_groups_spaces`
--

LOCK TABLES `security_groups_spaces` WRITE;
/*!40000 ALTER TABLE `security_groups_spaces` DISABLE KEYS */;
/*!40000 ALTER TABLE `security_groups_spaces` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_binding_annotations`
--

DROP TABLE IF EXISTS `service_binding_annotations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_binding_annotations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key` varchar(1000) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(5000) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `service_binding_annotations_guid_index` (`guid`),
  KEY `service_binding_annotations_created_at_index` (`created_at`),
  KEY `service_binding_annotations_updated_at_index` (`updated_at`),
  KEY `fk_service_binding_annotations_resource_guid_index` (`resource_guid`),
  CONSTRAINT `fk_service_binding_annotations_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `service_bindings` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_binding_annotations`
--

LOCK TABLES `service_binding_annotations` WRITE;
/*!40000 ALTER TABLE `service_binding_annotations` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_binding_annotations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_binding_labels`
--

DROP TABLE IF EXISTS `service_binding_labels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_binding_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key_name` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `service_binding_labels_guid_index` (`guid`),
  KEY `service_binding_labels_created_at_index` (`created_at`),
  KEY `service_binding_labels_updated_at_index` (`updated_at`),
  KEY `fk_service_binding_labels_resource_guid_index` (`resource_guid`),
  KEY `service_binding_labels_compound_index` (`key_prefix`,`key_name`,`value`),
  CONSTRAINT `fk_service_binding_labels_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `service_bindings` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_binding_labels`
--

LOCK TABLES `service_binding_labels` WRITE;
/*!40000 ALTER TABLE `service_binding_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_binding_labels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_binding_operations`
--

DROP TABLE IF EXISTS `service_binding_operations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_binding_operations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `service_binding_id` int DEFAULT NULL,
  `state` varchar(255) COLLATE utf8_bin NOT NULL,
  `type` varchar(255) COLLATE utf8_bin NOT NULL,
  `description` varchar(10000) COLLATE utf8_bin DEFAULT NULL,
  `broker_provided_operation` varchar(10000) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `svc_binding_id_index` (`service_binding_id`),
  KEY `service_binding_operations_created_at_index` (`created_at`),
  KEY `service_binding_operations_updated_at_index` (`updated_at`),
  CONSTRAINT `fk_svc_binding_op_svc_binding_id` FOREIGN KEY (`service_binding_id`) REFERENCES `service_bindings` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_binding_operations`
--

LOCK TABLES `service_binding_operations` WRITE;
/*!40000 ALTER TABLE `service_binding_operations` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_binding_operations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_bindings`
--

DROP TABLE IF EXISTS `service_bindings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_bindings` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `credentials` text CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `salt` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `syslog_drain_url` text,
  `volume_mounts` text,
  `volume_mounts_salt` varchar(255) DEFAULT NULL,
  `app_guid` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `service_instance_guid` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `type` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `encryption_key_label` varchar(255) DEFAULT NULL,
  `encryption_iterations` int NOT NULL DEFAULT '2048',
  PRIMARY KEY (`id`),
  UNIQUE KEY `sb_guid_index` (`guid`),
  UNIQUE KEY `unique_service_binding_service_instance_guid_app_guid` (`service_instance_guid`,`app_guid`),
  UNIQUE KEY `unique_service_binding_app_guid_name` (`app_guid`,`name`),
  KEY `sb_created_at_index` (`created_at`),
  KEY `sb_updated_at_index` (`updated_at`),
  KEY `service_bindings_app_guid_index` (`app_guid`),
  KEY `service_bindings_service_instance_guid_index` (`service_instance_guid`),
  KEY `service_bindings_name_index` (`name`),
  CONSTRAINT `fk_service_bindings_app_guid` FOREIGN KEY (`app_guid`) REFERENCES `apps` (`guid`),
  CONSTRAINT `fk_service_bindings_service_instance_guid` FOREIGN KEY (`service_instance_guid`) REFERENCES `service_instances` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_bindings`
--

LOCK TABLES `service_bindings` WRITE;
/*!40000 ALTER TABLE `service_bindings` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_bindings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_broker_annotations`
--

DROP TABLE IF EXISTS `service_broker_annotations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_broker_annotations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key` varchar(1000) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(5000) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `service_broker_annotations_guid_index` (`guid`),
  KEY `service_broker_annotations_created_at_index` (`created_at`),
  KEY `service_broker_annotations_updated_at_index` (`updated_at`),
  KEY `fk_service_broker_annotations_resource_guid_index` (`resource_guid`),
  CONSTRAINT `fk_service_broker_annotations_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `service_brokers` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_broker_annotations`
--

LOCK TABLES `service_broker_annotations` WRITE;
/*!40000 ALTER TABLE `service_broker_annotations` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_broker_annotations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_broker_labels`
--

DROP TABLE IF EXISTS `service_broker_labels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_broker_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key_name` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `service_broker_labels_guid_index` (`guid`),
  KEY `service_broker_labels_created_at_index` (`created_at`),
  KEY `service_broker_labels_updated_at_index` (`updated_at`),
  KEY `fk_service_broker_labels_resource_guid_index` (`resource_guid`),
  KEY `service_broker_labels_compound_index` (`key_prefix`,`key_name`,`value`),
  CONSTRAINT `fk_service_broker_labels_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `service_brokers` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_broker_labels`
--

LOCK TABLES `service_broker_labels` WRITE;
/*!40000 ALTER TABLE `service_broker_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_broker_labels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_broker_update_request_annotations`
--

DROP TABLE IF EXISTS `service_broker_update_request_annotations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_broker_update_request_annotations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key` varchar(1000) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(5000) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `service_broker_update_request_annotations_guid_index` (`guid`),
  KEY `service_broker_update_request_annotations_created_at_index` (`created_at`),
  KEY `service_broker_update_request_annotations_updated_at_index` (`updated_at`),
  KEY `fk_sb_update_request_annotations_resource_guid_index` (`resource_guid`),
  CONSTRAINT `fk_sb_update_request_annotations_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `service_broker_update_requests` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_broker_update_request_annotations`
--

LOCK TABLES `service_broker_update_request_annotations` WRITE;
/*!40000 ALTER TABLE `service_broker_update_request_annotations` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_broker_update_request_annotations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_broker_update_request_labels`
--

DROP TABLE IF EXISTS `service_broker_update_request_labels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_broker_update_request_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key_name` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `service_broker_update_request_labels_guid_index` (`guid`),
  KEY `service_broker_update_request_labels_created_at_index` (`created_at`),
  KEY `service_broker_update_request_labels_updated_at_index` (`updated_at`),
  KEY `fk_sb_update_request_labels_resource_guid_index` (`resource_guid`),
  KEY `sb_update_request_labels_compound_index` (`key_prefix`,`key_name`,`value`),
  CONSTRAINT `fk_sb_update_request_labels_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `service_broker_update_requests` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_broker_update_request_labels`
--

LOCK TABLES `service_broker_update_request_labels` WRITE;
/*!40000 ALTER TABLE `service_broker_update_request_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_broker_update_request_labels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_broker_update_requests`
--

DROP TABLE IF EXISTS `service_broker_update_requests`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_broker_update_requests` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `broker_url` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `authentication` varchar(16000) COLLATE utf8_bin DEFAULT NULL,
  `salt` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `encryption_key_label` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `encryption_iterations` int NOT NULL DEFAULT '2048',
  `service_broker_id` int NOT NULL,
  `fk_service_brokers_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `service_broker_update_requests_guid_index` (`guid`),
  KEY `fk_service_brokers_id` (`fk_service_brokers_id`),
  KEY `service_broker_update_requests_created_at_index` (`created_at`),
  KEY `service_broker_update_requests_updated_at_index` (`updated_at`),
  CONSTRAINT `service_broker_update_requests_ibfk_1` FOREIGN KEY (`fk_service_brokers_id`) REFERENCES `service_brokers` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_broker_update_requests`
--

LOCK TABLES `service_broker_update_requests` WRITE;
/*!40000 ALTER TABLE `service_broker_update_requests` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_broker_update_requests` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_brokers`
--

DROP TABLE IF EXISTS `service_brokers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_brokers` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) COLLATE utf8_bin NOT NULL,
  `broker_url` varchar(255) COLLATE utf8_bin NOT NULL,
  `auth_password` varchar(255) COLLATE utf8_bin NOT NULL,
  `salt` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `auth_username` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `space_id` int DEFAULT NULL,
  `encryption_key_label` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `encryption_iterations` int NOT NULL DEFAULT '2048',
  `state` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `sbrokers_guid_index` (`guid`),
  UNIQUE KEY `service_brokers_name_index` (`name`),
  KEY `sbrokers_created_at_index` (`created_at`),
  KEY `sbrokers_updated_at_index` (`updated_at`),
  KEY `space_id` (`space_id`),
  KEY `sb_broker_url_index` (`broker_url`),
  CONSTRAINT `service_brokers_ibfk_1` FOREIGN KEY (`space_id`) REFERENCES `spaces` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_brokers`
--

LOCK TABLES `service_brokers` WRITE;
/*!40000 ALTER TABLE `service_brokers` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_brokers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_dashboard_clients`
--

DROP TABLE IF EXISTS `service_dashboard_clients`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_dashboard_clients` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `uaa_id` varchar(255) COLLATE utf8_bin NOT NULL,
  `service_broker_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `s_d_clients_uaa_id_unique` (`uaa_id`),
  KEY `s_d_clients_created_at_index` (`created_at`),
  KEY `s_d_clients_updated_at_index` (`updated_at`),
  KEY `svc_dash_cli_svc_brkr_id_idx` (`service_broker_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_dashboard_clients`
--

LOCK TABLES `service_dashboard_clients` WRITE;
/*!40000 ALTER TABLE `service_dashboard_clients` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_dashboard_clients` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_instance_annotations`
--

DROP TABLE IF EXISTS `service_instance_annotations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_instance_annotations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key` varchar(1000) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(5000) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `service_instance_annotations_guid_index` (`guid`),
  KEY `service_instance_annotations_created_at_index` (`created_at`),
  KEY `service_instance_annotations_updated_at_index` (`updated_at`),
  KEY `fk_service_instance_annotations_resource_guid_index` (`resource_guid`),
  CONSTRAINT `fk_service_instance_annotations_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `service_instances` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_instance_annotations`
--

LOCK TABLES `service_instance_annotations` WRITE;
/*!40000 ALTER TABLE `service_instance_annotations` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_instance_annotations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_instance_labels`
--

DROP TABLE IF EXISTS `service_instance_labels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_instance_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key_name` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `service_instance_labels_guid_index` (`guid`),
  KEY `service_instance_labels_created_at_index` (`created_at`),
  KEY `service_instance_labels_updated_at_index` (`updated_at`),
  KEY `fk_service_instance_labels_resource_guid_index` (`resource_guid`),
  KEY `service_instance_labels_compound_index` (`key_prefix`,`key_name`,`value`),
  CONSTRAINT `fk_service_instance_labels_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `service_instances` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_instance_labels`
--

LOCK TABLES `service_instance_labels` WRITE;
/*!40000 ALTER TABLE `service_instance_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_instance_labels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_instance_operations`
--

DROP TABLE IF EXISTS `service_instance_operations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_instance_operations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `service_instance_id` int DEFAULT NULL,
  `type` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `state` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `description` text COLLATE utf8_bin,
  `proposed_changes` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '{}',
  `broker_provided_operation` text COLLATE utf8_bin,
  PRIMARY KEY (`id`),
  UNIQUE KEY `svc_inst_op_guid_index` (`guid`),
  KEY `svc_inst_op_created_at_index` (`created_at`),
  KEY `svc_inst_op_updated_at_index` (`updated_at`),
  KEY `svc_instance_id_index` (`service_instance_id`),
  CONSTRAINT `fk_svc_inst_op_svc_instance_id` FOREIGN KEY (`service_instance_id`) REFERENCES `service_instances` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_instance_operations`
--

LOCK TABLES `service_instance_operations` WRITE;
/*!40000 ALTER TABLE `service_instance_operations` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_instance_operations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_instance_shares`
--

DROP TABLE IF EXISTS `service_instance_shares`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_instance_shares` (
  `service_instance_guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `target_space_guid` varchar(255) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`service_instance_guid`,`target_space_guid`),
  KEY `fk_target_space_guid` (`target_space_guid`),
  CONSTRAINT `fk_service_instance_guid` FOREIGN KEY (`service_instance_guid`) REFERENCES `service_instances` (`guid`) ON DELETE CASCADE,
  CONSTRAINT `fk_target_space_guid` FOREIGN KEY (`target_space_guid`) REFERENCES `spaces` (`guid`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_instance_shares`
--

LOCK TABLES `service_instance_shares` WRITE;
/*!40000 ALTER TABLE `service_instance_shares` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_instance_shares` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_instances`
--

DROP TABLE IF EXISTS `service_instances`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_instances` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `credentials` text CHARACTER SET utf8 COLLATE utf8_bin,
  `gateway_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `gateway_data` varchar(2048) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `space_id` int NOT NULL,
  `service_plan_id` int DEFAULT NULL,
  `salt` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `dashboard_url` varchar(16000) DEFAULT NULL,
  `is_gateway_service` tinyint(1) NOT NULL DEFAULT '1',
  `syslog_drain_url` text,
  `tags` text,
  `route_service_url` varchar(255) DEFAULT NULL,
  `encryption_key_label` varchar(255) DEFAULT NULL,
  `encryption_iterations` int NOT NULL DEFAULT '2048',
  `maintenance_info` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `si_guid_index` (`guid`),
  UNIQUE KEY `si_space_id_name_index` (`space_id`,`name`),
  KEY `si_created_at_index` (`created_at`),
  KEY `si_updated_at_index` (`updated_at`),
  KEY `service_instances_name_index` (`name`),
  KEY `svc_instances_service_plan_id` (`service_plan_id`),
  KEY `si_gateway_name_index` (`gateway_name`),
  CONSTRAINT `service_instances_space_id` FOREIGN KEY (`space_id`) REFERENCES `spaces` (`id`),
  CONSTRAINT `svc_instances_service_plan_id` FOREIGN KEY (`service_plan_id`) REFERENCES `service_plans` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_instances`
--

LOCK TABLES `service_instances` WRITE;
/*!40000 ALTER TABLE `service_instances` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_instances` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_key_annotations`
--

DROP TABLE IF EXISTS `service_key_annotations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_key_annotations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key` varchar(1000) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(5000) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `service_key_annotations_guid_index` (`guid`),
  KEY `service_key_annotations_created_at_index` (`created_at`),
  KEY `service_key_annotations_updated_at_index` (`updated_at`),
  KEY `fk_service_key_annotations_resource_guid_index` (`resource_guid`),
  CONSTRAINT `fk_service_key_annotations_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `service_keys` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_key_annotations`
--

LOCK TABLES `service_key_annotations` WRITE;
/*!40000 ALTER TABLE `service_key_annotations` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_key_annotations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_key_labels`
--

DROP TABLE IF EXISTS `service_key_labels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_key_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key_name` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `service_key_labels_guid_index` (`guid`),
  KEY `service_key_labels_created_at_index` (`created_at`),
  KEY `service_key_labels_updated_at_index` (`updated_at`),
  KEY `fk_service_key_labels_resource_guid_index` (`resource_guid`),
  KEY `service_key_labels_compound_index` (`key_prefix`,`key_name`,`value`),
  CONSTRAINT `fk_service_key_labels_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `service_keys` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_key_labels`
--

LOCK TABLES `service_key_labels` WRITE;
/*!40000 ALTER TABLE `service_key_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_key_labels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_key_operations`
--

DROP TABLE IF EXISTS `service_key_operations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_key_operations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `service_key_id` int DEFAULT NULL,
  `state` varchar(255) COLLATE utf8_bin NOT NULL,
  `type` varchar(255) COLLATE utf8_bin NOT NULL,
  `description` varchar(10000) COLLATE utf8_bin DEFAULT NULL,
  `broker_provided_operation` varchar(10000) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `svc_key_id_index` (`service_key_id`),
  KEY `service_key_operations_created_at_index` (`created_at`),
  KEY `service_key_operations_updated_at_index` (`updated_at`),
  CONSTRAINT `fk_svc_key_op_svc_key_id` FOREIGN KEY (`service_key_id`) REFERENCES `service_keys` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_key_operations`
--

LOCK TABLES `service_key_operations` WRITE;
/*!40000 ALTER TABLE `service_key_operations` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_key_operations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_keys`
--

DROP TABLE IF EXISTS `service_keys`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_keys` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) COLLATE utf8_bin NOT NULL,
  `salt` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `credentials` text COLLATE utf8_bin NOT NULL,
  `service_instance_id` int NOT NULL,
  `encryption_key_label` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `encryption_iterations` int NOT NULL DEFAULT '2048',
  PRIMARY KEY (`id`),
  UNIQUE KEY `sk_guid_index` (`guid`),
  UNIQUE KEY `svc_key_name_instance_id_index` (`name`,`service_instance_id`),
  KEY `fk_svc_key_svc_instance_id` (`service_instance_id`),
  KEY `sk_created_at_index` (`created_at`),
  KEY `sk_updated_at_index` (`updated_at`),
  CONSTRAINT `fk_svc_key_svc_instance_id` FOREIGN KEY (`service_instance_id`) REFERENCES `service_instances` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_keys`
--

LOCK TABLES `service_keys` WRITE;
/*!40000 ALTER TABLE `service_keys` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_keys` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_offering_annotations`
--

DROP TABLE IF EXISTS `service_offering_annotations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_offering_annotations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key` varchar(1000) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(5000) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `service_offering_annotations_guid_index` (`guid`),
  KEY `service_offering_annotations_created_at_index` (`created_at`),
  KEY `service_offering_annotations_updated_at_index` (`updated_at`),
  KEY `fk_service_offering_annotations_resource_guid_index` (`resource_guid`),
  CONSTRAINT `fk_service_offering_annotations_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `services` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_offering_annotations`
--

LOCK TABLES `service_offering_annotations` WRITE;
/*!40000 ALTER TABLE `service_offering_annotations` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_offering_annotations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_offering_labels`
--

DROP TABLE IF EXISTS `service_offering_labels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_offering_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key_name` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `service_offering_labels_guid_index` (`guid`),
  KEY `service_offering_labels_created_at_index` (`created_at`),
  KEY `service_offering_labels_updated_at_index` (`updated_at`),
  KEY `fk_service_offering_labels_resource_guid_index` (`resource_guid`),
  KEY `service_offering_labels_compound_index` (`key_prefix`,`key_name`,`value`),
  CONSTRAINT `fk_service_offering_labels_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `services` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_offering_labels`
--

LOCK TABLES `service_offering_labels` WRITE;
/*!40000 ALTER TABLE `service_offering_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_offering_labels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_plan_annotations`
--

DROP TABLE IF EXISTS `service_plan_annotations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_plan_annotations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key` varchar(1000) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(5000) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `service_plan_annotations_guid_index` (`guid`),
  KEY `service_plan_annotations_created_at_index` (`created_at`),
  KEY `service_plan_annotations_updated_at_index` (`updated_at`),
  KEY `fk_service_plan_annotations_resource_guid_index` (`resource_guid`),
  CONSTRAINT `fk_service_plan_annotations_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `service_plans` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_plan_annotations`
--

LOCK TABLES `service_plan_annotations` WRITE;
/*!40000 ALTER TABLE `service_plan_annotations` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_plan_annotations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_plan_labels`
--

DROP TABLE IF EXISTS `service_plan_labels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_plan_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key_name` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `service_plan_labels_guid_index` (`guid`),
  KEY `service_plan_labels_created_at_index` (`created_at`),
  KEY `service_plan_labels_updated_at_index` (`updated_at`),
  KEY `fk_service_plan_labels_resource_guid_index` (`resource_guid`),
  KEY `service_plan_labels_compound_index` (`key_prefix`,`key_name`,`value`),
  CONSTRAINT `fk_service_plan_labels_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `service_plans` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_plan_labels`
--

LOCK TABLES `service_plan_labels` WRITE;
/*!40000 ALTER TABLE `service_plan_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_plan_labels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_plan_visibilities`
--

DROP TABLE IF EXISTS `service_plan_visibilities`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_plan_visibilities` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `service_plan_id` int NOT NULL,
  `organization_id` int NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `spv_guid_index` (`guid`),
  UNIQUE KEY `spv_org_id_sp_id_index` (`organization_id`,`service_plan_id`),
  KEY `spv_created_at_index` (`created_at`),
  KEY `spv_updated_at_index` (`updated_at`),
  KEY `fk_spv_service_plan_id` (`service_plan_id`),
  CONSTRAINT `fk_spv_organization_id` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`),
  CONSTRAINT `fk_spv_service_plan_id` FOREIGN KEY (`service_plan_id`) REFERENCES `service_plans` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_plan_visibilities`
--

LOCK TABLES `service_plan_visibilities` WRITE;
/*!40000 ALTER TABLE `service_plan_visibilities` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_plan_visibilities` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_plans`
--

DROP TABLE IF EXISTS `service_plans`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_plans` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `description` text NOT NULL,
  `free` tinyint(1) NOT NULL,
  `service_id` int NOT NULL,
  `extra` mediumtext CHARACTER SET utf8 COLLATE utf8_bin,
  `unique_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `public` tinyint(1) DEFAULT '1',
  `active` tinyint(1) DEFAULT '1',
  `bindable` tinyint(1) DEFAULT NULL,
  `create_instance_schema` text,
  `update_instance_schema` text,
  `create_binding_schema` text,
  `plan_updateable` tinyint(1) DEFAULT NULL,
  `maximum_polling_duration` int DEFAULT NULL,
  `maintenance_info` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `service_plans_guid_index` (`guid`),
  UNIQUE KEY `svc_plan_svc_id_name_index` (`service_id`,`name`),
  KEY `service_plans_created_at_index` (`created_at`),
  KEY `service_plans_updated_at_index` (`updated_at`),
  KEY `service_plans_unique_id_index` (`unique_id`),
  CONSTRAINT `fk_service_plans_service_id` FOREIGN KEY (`service_id`) REFERENCES `services` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_plans`
--

LOCK TABLES `service_plans` WRITE;
/*!40000 ALTER TABLE `service_plans` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_plans` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_usage_events`
--

DROP TABLE IF EXISTS `service_usage_events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_usage_events` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `state` varchar(255) COLLATE utf8_bin NOT NULL,
  `org_guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `space_guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `space_name` varchar(255) COLLATE utf8_bin NOT NULL,
  `service_instance_guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `service_instance_name` varchar(255) COLLATE utf8_bin NOT NULL,
  `service_instance_type` varchar(255) COLLATE utf8_bin NOT NULL,
  `service_plan_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `service_plan_name` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `service_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `service_label` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `service_broker_name` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `service_broker_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `usage_events_guid_index` (`guid`),
  KEY `created_at_index` (`created_at`),
  KEY `service_usage_events_service_guid_index` (`service_guid`),
  KEY `service_usage_events_service_instance_type_index` (`service_instance_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_usage_events`
--

LOCK TABLES `service_usage_events` WRITE;
/*!40000 ALTER TABLE `service_usage_events` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_usage_events` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `services`
--

DROP TABLE IF EXISTS `services`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `services` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `label` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `description` text NOT NULL,
  `info_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `acls` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `timeout` int DEFAULT NULL,
  `active` tinyint(1) DEFAULT '0',
  `extra` mediumtext CHARACTER SET utf8 COLLATE utf8_bin,
  `unique_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `bindable` tinyint(1) NOT NULL,
  `tags` text,
  `documentation_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `service_broker_id` int DEFAULT NULL,
  `long_description` text CHARACTER SET utf8 COLLATE utf8_bin,
  `requires` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `purging` tinyint(1) NOT NULL DEFAULT '0',
  `plan_updateable` tinyint(1) DEFAULT '0',
  `bindings_retrievable` tinyint(1) NOT NULL DEFAULT '0',
  `instances_retrievable` tinyint(1) NOT NULL DEFAULT '0',
  `allow_context_updates` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `services_guid_index` (`guid`),
  KEY `services_created_at_index` (`created_at`),
  KEY `services_updated_at_index` (`updated_at`),
  KEY `services_label_index` (`label`),
  KEY `fk_services_service_broker_id` (`service_broker_id`),
  KEY `services_unique_id_index` (`unique_id`),
  CONSTRAINT `fk_services_service_broker_id` FOREIGN KEY (`service_broker_id`) REFERENCES `service_brokers` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `services`
--

LOCK TABLES `services` WRITE;
/*!40000 ALTER TABLE `services` DISABLE KEYS */;
/*!40000 ALTER TABLE `services` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sidecar_process_types`
--

DROP TABLE IF EXISTS `sidecar_process_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sidecar_process_types` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `type` varchar(255) COLLATE utf8_bin NOT NULL,
  `sidecar_guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `app_guid` varchar(255) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `sidecar_process_types_guid_index` (`guid`),
  KEY `sidecar_process_types_created_at_index` (`created_at`),
  KEY `sidecar_process_types_updated_at_index` (`updated_at`),
  KEY `fk_sidecar_proc_type_sidecar_guid_index` (`sidecar_guid`),
  CONSTRAINT `fk_sidecar_proc_type_sidecar_guid` FOREIGN KEY (`sidecar_guid`) REFERENCES `sidecars` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sidecar_process_types`
--

LOCK TABLES `sidecar_process_types` WRITE;
/*!40000 ALTER TABLE `sidecar_process_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `sidecar_process_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sidecars`
--

DROP TABLE IF EXISTS `sidecars`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sidecars` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) COLLATE utf8_bin NOT NULL,
  `command` varchar(4096) COLLATE utf8_bin NOT NULL,
  `app_guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `memory` int DEFAULT NULL,
  `origin` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT 'user',
  PRIMARY KEY (`id`),
  UNIQUE KEY `sidecars_guid_index` (`guid`),
  KEY `sidecars_created_at_index` (`created_at`),
  KEY `sidecars_updated_at_index` (`updated_at`),
  KEY `fk_sidecar_app_guid_index` (`app_guid`),
  CONSTRAINT `fk_sidecar_app_guid` FOREIGN KEY (`app_guid`) REFERENCES `apps` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sidecars`
--

LOCK TABLES `sidecars` WRITE;
/*!40000 ALTER TABLE `sidecars` DISABLE KEYS */;
/*!40000 ALTER TABLE `sidecars` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `space_annotations`
--

DROP TABLE IF EXISTS `space_annotations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `space_annotations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key` varchar(1000) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(5000) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `space_annotations_guid_index` (`guid`),
  KEY `space_annotations_created_at_index` (`created_at`),
  KEY `space_annotations_updated_at_index` (`updated_at`),
  KEY `fk_space_annotations_resource_guid_index` (`resource_guid`),
  CONSTRAINT `fk_space_annotations_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `spaces` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `space_annotations`
--

LOCK TABLES `space_annotations` WRITE;
/*!40000 ALTER TABLE `space_annotations` DISABLE KEYS */;
/*!40000 ALTER TABLE `space_annotations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `space_labels`
--

DROP TABLE IF EXISTS `space_labels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `space_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key_name` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `space_labels_guid_index` (`guid`),
  KEY `space_labels_created_at_index` (`created_at`),
  KEY `space_labels_updated_at_index` (`updated_at`),
  KEY `fk_space_labels_resource_guid_index` (`resource_guid`),
  KEY `space_labels_compound_index` (`key_prefix`,`key_name`,`value`),
  CONSTRAINT `fk_space_labels_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `spaces` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `space_labels`
--

LOCK TABLES `space_labels` WRITE;
/*!40000 ALTER TABLE `space_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `space_labels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `space_quota_definitions`
--

DROP TABLE IF EXISTS `space_quota_definitions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `space_quota_definitions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) COLLATE utf8_bin NOT NULL,
  `non_basic_services_allowed` tinyint(1) NOT NULL,
  `total_services` int NOT NULL,
  `memory_limit` int NOT NULL,
  `total_routes` int NOT NULL,
  `instance_memory_limit` int NOT NULL DEFAULT '-1',
  `organization_id` int NOT NULL,
  `app_instance_limit` int DEFAULT '-1',
  `app_task_limit` int DEFAULT '5',
  `total_service_keys` int NOT NULL DEFAULT '-1',
  `total_reserved_route_ports` int DEFAULT '-1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `sqd_guid_index` (`guid`),
  UNIQUE KEY `sqd_org_id_index` (`organization_id`,`name`),
  KEY `sqd_created_at_index` (`created_at`),
  KEY `sqd_updated_at_index` (`updated_at`),
  CONSTRAINT `fk_sqd_organization_id` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `space_quota_definitions`
--

LOCK TABLES `space_quota_definitions` WRITE;
/*!40000 ALTER TABLE `space_quota_definitions` DISABLE KEYS */;
/*!40000 ALTER TABLE `space_quota_definitions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `spaces`
--

DROP TABLE IF EXISTS `spaces`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `spaces` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `organization_id` int NOT NULL,
  `space_quota_definition_id` int DEFAULT NULL,
  `allow_ssh` tinyint(1) DEFAULT '1',
  `isolation_segment_guid` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `spaces_guid_index` (`guid`),
  UNIQUE KEY `spaces_org_id_name_index` (`organization_id`,`name`),
  KEY `spaces_created_at_index` (`created_at`),
  KEY `spaces_updated_at_index` (`updated_at`),
  KEY `fk_space_sqd_id` (`space_quota_definition_id`),
  KEY `fk_spaces_isolation_segment_guid` (`isolation_segment_guid`),
  CONSTRAINT `fk_space_sqd_id` FOREIGN KEY (`space_quota_definition_id`) REFERENCES `space_quota_definitions` (`id`),
  CONSTRAINT `fk_spaces_isolation_segment_guid` FOREIGN KEY (`isolation_segment_guid`) REFERENCES `isolation_segments` (`guid`),
  CONSTRAINT `fk_spaces_organization_id` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `spaces`
--

LOCK TABLES `spaces` WRITE;
/*!40000 ALTER TABLE `spaces` DISABLE KEYS */;
/*!40000 ALTER TABLE `spaces` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `spaces_application_supporters`
--

DROP TABLE IF EXISTS `spaces_application_supporters`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `spaces_application_supporters` (
  `spaces_application_supporters_pk` int NOT NULL AUTO_INCREMENT,
  `role_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `space_id` int NOT NULL,
  `user_id` int NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`spaces_application_supporters_pk`),
  UNIQUE KEY `spaces_application_supporters_user_space_index` (`space_id`,`user_id`),
  KEY `spaces_application_supporters_user_fk` (`user_id`),
  KEY `spaces_application_supporters_role_guid_index` (`role_guid`),
  KEY `spaces_application_supporters_created_at_index` (`created_at`),
  KEY `spaces_application_supporters_updated_at_index` (`updated_at`),
  CONSTRAINT `spaces_application_supporters_space_fk` FOREIGN KEY (`space_id`) REFERENCES `spaces` (`id`),
  CONSTRAINT `spaces_application_supporters_user_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `spaces_application_supporters`
--

LOCK TABLES `spaces_application_supporters` WRITE;
/*!40000 ALTER TABLE `spaces_application_supporters` DISABLE KEYS */;
/*!40000 ALTER TABLE `spaces_application_supporters` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `spaces_auditors`
--

DROP TABLE IF EXISTS `spaces_auditors`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `spaces_auditors` (
  `space_id` int NOT NULL,
  `user_id` int NOT NULL,
  `spaces_auditors_pk` int NOT NULL AUTO_INCREMENT,
  `role_guid` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`spaces_auditors_pk`),
  UNIQUE KEY `space_auditors_idx` (`space_id`,`user_id`),
  KEY `space_auditors_user_fk` (`user_id`),
  KEY `spaces_auditors_role_guid_index` (`role_guid`),
  KEY `spaces_auditors_created_at_index` (`created_at`),
  KEY `spaces_auditors_updated_at_index` (`updated_at`),
  CONSTRAINT `space_auditors_space_fk` FOREIGN KEY (`space_id`) REFERENCES `spaces` (`id`),
  CONSTRAINT `space_auditors_user_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `spaces_auditors`
--

LOCK TABLES `spaces_auditors` WRITE;
/*!40000 ALTER TABLE `spaces_auditors` DISABLE KEYS */;
/*!40000 ALTER TABLE `spaces_auditors` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `spaces_developers`
--

DROP TABLE IF EXISTS `spaces_developers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `spaces_developers` (
  `space_id` int NOT NULL,
  `user_id` int NOT NULL,
  `spaces_developers_pk` int NOT NULL AUTO_INCREMENT,
  `role_guid` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`spaces_developers_pk`),
  UNIQUE KEY `space_developers_idx` (`space_id`,`user_id`),
  KEY `space_developers_user_fk` (`user_id`),
  KEY `spaces_developers_role_guid_index` (`role_guid`),
  KEY `spaces_developers_created_at_index` (`created_at`),
  KEY `spaces_developers_updated_at_index` (`updated_at`),
  CONSTRAINT `space_developers_space_fk` FOREIGN KEY (`space_id`) REFERENCES `spaces` (`id`),
  CONSTRAINT `space_developers_user_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `spaces_developers`
--

LOCK TABLES `spaces_developers` WRITE;
/*!40000 ALTER TABLE `spaces_developers` DISABLE KEYS */;
/*!40000 ALTER TABLE `spaces_developers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `spaces_managers`
--

DROP TABLE IF EXISTS `spaces_managers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `spaces_managers` (
  `space_id` int NOT NULL,
  `user_id` int NOT NULL,
  `spaces_managers_pk` int NOT NULL AUTO_INCREMENT,
  `role_guid` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`spaces_managers_pk`),
  UNIQUE KEY `space_managers_idx` (`space_id`,`user_id`),
  KEY `space_managers_user_fk` (`user_id`),
  KEY `spaces_managers_role_guid_index` (`role_guid`),
  KEY `spaces_managers_created_at_index` (`created_at`),
  KEY `spaces_managers_updated_at_index` (`updated_at`),
  CONSTRAINT `space_managers_space_fk` FOREIGN KEY (`space_id`) REFERENCES `spaces` (`id`),
  CONSTRAINT `space_managers_user_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `spaces_managers`
--

LOCK TABLES `spaces_managers` WRITE;
/*!40000 ALTER TABLE `spaces_managers` DISABLE KEYS */;
/*!40000 ALTER TABLE `spaces_managers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `stack_annotations`
--

DROP TABLE IF EXISTS `stack_annotations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `stack_annotations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key` varchar(1000) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(5000) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `stack_annotations_guid_index` (`guid`),
  KEY `stack_annotations_created_at_index` (`created_at`),
  KEY `stack_annotations_updated_at_index` (`updated_at`),
  KEY `fk_stack_annotations_resource_guid_index` (`resource_guid`),
  CONSTRAINT `fk_stack_annotations_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `stacks` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `stack_annotations`
--

LOCK TABLES `stack_annotations` WRITE;
/*!40000 ALTER TABLE `stack_annotations` DISABLE KEYS */;
/*!40000 ALTER TABLE `stack_annotations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `stack_labels`
--

DROP TABLE IF EXISTS `stack_labels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `stack_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key_name` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `stack_labels_guid_index` (`guid`),
  KEY `stack_labels_created_at_index` (`created_at`),
  KEY `stack_labels_updated_at_index` (`updated_at`),
  KEY `fk_stack_labels_resource_guid_index` (`resource_guid`),
  KEY `stack_labels_compound_index` (`key_prefix`,`key_name`,`value`),
  CONSTRAINT `fk_stack_labels_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `stacks` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `stack_labels`
--

LOCK TABLES `stack_labels` WRITE;
/*!40000 ALTER TABLE `stack_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `stack_labels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `stacks`
--

DROP TABLE IF EXISTS `stacks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `stacks` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `stacks_guid_index` (`guid`),
  UNIQUE KEY `stacks_name_index` (`name`),
  KEY `stacks_created_at_index` (`created_at`),
  KEY `stacks_updated_at_index` (`updated_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `stacks`
--

LOCK TABLES `stacks` WRITE;
/*!40000 ALTER TABLE `stacks` DISABLE KEYS */;
/*!40000 ALTER TABLE `stacks` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `staging_security_groups_spaces`
--

DROP TABLE IF EXISTS `staging_security_groups_spaces`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `staging_security_groups_spaces` (
  `staging_security_group_id` int NOT NULL,
  `staging_space_id` int NOT NULL,
  `staging_security_groups_spaces_pk` int NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`staging_security_groups_spaces_pk`),
  UNIQUE KEY `staging_security_groups_spaces_ids` (`staging_security_group_id`,`staging_space_id`),
  KEY `fk_staging_space_id` (`staging_space_id`),
  CONSTRAINT `fk_staging_security_group_id` FOREIGN KEY (`staging_security_group_id`) REFERENCES `security_groups` (`id`),
  CONSTRAINT `fk_staging_space_id` FOREIGN KEY (`staging_space_id`) REFERENCES `spaces` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `staging_security_groups_spaces`
--

LOCK TABLES `staging_security_groups_spaces` WRITE;
/*!40000 ALTER TABLE `staging_security_groups_spaces` DISABLE KEYS */;
/*!40000 ALTER TABLE `staging_security_groups_spaces` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `task_annotations`
--

DROP TABLE IF EXISTS `task_annotations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `task_annotations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key` varchar(1000) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(5000) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `task_annotations_guid_index` (`guid`),
  KEY `task_annotations_created_at_index` (`created_at`),
  KEY `task_annotations_updated_at_index` (`updated_at`),
  KEY `fk_task_annotations_resource_guid_index` (`resource_guid`),
  CONSTRAINT `fk_task_annotations_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `tasks` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `task_annotations`
--

LOCK TABLES `task_annotations` WRITE;
/*!40000 ALTER TABLE `task_annotations` DISABLE KEYS */;
/*!40000 ALTER TABLE `task_annotations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `task_labels`
--

DROP TABLE IF EXISTS `task_labels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `task_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key_name` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `task_labels_guid_index` (`guid`),
  KEY `task_labels_created_at_index` (`created_at`),
  KEY `task_labels_updated_at_index` (`updated_at`),
  KEY `fk_task_labels_resource_guid_index` (`resource_guid`),
  KEY `task_labels_compound_index` (`key_prefix`,`key_name`,`value`),
  CONSTRAINT `fk_task_labels_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `tasks` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `task_labels`
--

LOCK TABLES `task_labels` WRITE;
/*!40000 ALTER TABLE `task_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `task_labels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tasks`
--

DROP TABLE IF EXISTS `tasks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tasks` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `command` text COLLATE utf8_bin NOT NULL,
  `state` varchar(255) COLLATE utf8_bin NOT NULL,
  `memory_in_mb` int DEFAULT NULL,
  `encrypted_environment_variables` text COLLATE utf8_bin,
  `salt` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `failure_reason` varchar(4096) COLLATE utf8_bin DEFAULT NULL,
  `app_guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `droplet_guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `sequence_id` int DEFAULT NULL,
  `disk_in_mb` int DEFAULT NULL,
  `encryption_key_label` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `encryption_iterations` int NOT NULL DEFAULT '2048',
  PRIMARY KEY (`id`),
  UNIQUE KEY `tasks_guid_index` (`guid`),
  UNIQUE KEY `unique_task_app_guid_sequence_id` (`app_guid`,`sequence_id`),
  KEY `fk_tasks_droplet_guid` (`droplet_guid`),
  KEY `tasks_created_at_index` (`created_at`),
  KEY `tasks_updated_at_index` (`updated_at`),
  KEY `tasks_name_index` (`name`),
  KEY `tasks_state_index` (`state`),
  KEY `tasks_app_guid_index` (`app_guid`),
  CONSTRAINT `fk_tasks_app_guid` FOREIGN KEY (`app_guid`) REFERENCES `apps` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tasks`
--

LOCK TABLES `tasks` WRITE;
/*!40000 ALTER TABLE `tasks` DISABLE KEYS */;
/*!40000 ALTER TABLE `tasks` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_annotations`
--

DROP TABLE IF EXISTS `user_annotations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_annotations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key` varchar(1000) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(5000) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_annotations_guid_index` (`guid`),
  KEY `user_annotations_created_at_index` (`created_at`),
  KEY `user_annotations_updated_at_index` (`updated_at`),
  KEY `fk_user_annotations_resource_guid_index` (`resource_guid`),
  CONSTRAINT `fk_user_annotations_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `users` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_annotations`
--

LOCK TABLES `user_annotations` WRITE;
/*!40000 ALTER TABLE `user_annotations` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_annotations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_labels`
--

DROP TABLE IF EXISTS `user_labels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `resource_guid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `key_prefix` varchar(253) COLLATE utf8_bin DEFAULT NULL,
  `key_name` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(63) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_labels_guid_index` (`guid`),
  KEY `user_labels_created_at_index` (`created_at`),
  KEY `user_labels_updated_at_index` (`updated_at`),
  KEY `fk_user_labels_resource_guid_index` (`resource_guid`),
  KEY `user_labels_compound_index` (`key_prefix`,`key_name`,`value`),
  CONSTRAINT `fk_user_labels_resource_guid` FOREIGN KEY (`resource_guid`) REFERENCES `users` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_labels`
--

LOCK TABLES `user_labels` WRITE;
/*!40000 ALTER TABLE `user_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_labels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `guid` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `default_space_id` int DEFAULT NULL,
  `admin` tinyint(1) DEFAULT '0',
  `active` tinyint(1) DEFAULT '0',
  `is_oauth_client` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `users_guid_index` (`guid`),
  KEY `fk_users_default_space_id` (`default_space_id`),
  KEY `users_created_at_index` (`created_at`),
  KEY `users_updated_at_index` (`updated_at`),
  CONSTRAINT `fk_users_default_space_id` FOREIGN KEY (`default_space_id`) REFERENCES `spaces` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
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

-- Dump completed on 2021-07-27 15:11:15
