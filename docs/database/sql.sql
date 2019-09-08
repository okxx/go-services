-- MySQL dump 10.13  Distrib 8.0.15, for Win64 (x86_64)
--
-- Host: localhost    Database: pets
-- ------------------------------------------------------
-- Server version	8.0.15

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `tbl_category`
--

DROP TABLE IF EXISTS `tbl_category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `tbl_category` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `introduction` varchar(5000) DEFAULT NULL COMMENT '介绍',
  `feature` text COMMENT '特征',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `step` tinyint(1) DEFAULT '0' COMMENT '进度：0=待抓取 1=已抓取',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`) /*!80000 INVISIBLE */
) ENGINE=InnoDB AUTO_INCREMENT=52 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='类别';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tbl_category`
--

LOCK TABLES `tbl_category` WRITE;
/*!40000 ALTER TABLE `tbl_category` DISABLE KEYS */;
INSERT INTO `tbl_category` VALUES (1,'牛头梗','','','2019-04-28 07:27:04','2019-04-28 07:44:42',0),(3,'斯塔福斗牛梗','','','2019-04-28 07:27:57','2019-04-28 07:44:42',0),(6,'阿根廷杜高犬','','','2019-04-28 07:32:00','2019-04-28 07:44:42',0),(7,'古梗犬','','','2019-04-28 07:32:00','2019-04-28 07:44:42',0),(8,'平毛猎狐梗','','','2019-04-28 07:32:00','2019-04-28 07:44:42',0),(9,'杰克罗素梗','','','2019-04-28 07:32:00','2019-04-28 07:46:19',0),(16,'贵宾犬/贵妇犬','','','2019-04-28 07:48:35','2019-04-28 07:48:35',0),(17,'台湾犬','','','2019-04-28 07:48:35','2019-04-28 07:48:35',0),(18,'迷你贵宾犬','','','2019-04-28 07:48:35','2019-04-28 07:48:35',0),(19,'卷毛比熊犬','','','2019-04-28 07:48:35','2019-04-28 07:48:35',0),(20,'杂交狮子狗','','','2019-04-28 07:48:35','2019-04-28 07:48:35',0),(21,'可卡','','','2019-04-28 07:48:35','2019-04-28 07:48:35',0),(22,'纪州犬','','','2019-04-28 07:48:56','2019-04-28 07:48:56',0),(23,'下司犬','','','2019-04-28 07:48:56','2019-04-28 07:48:56',0),(24,'广东潮州犬','','','2019-04-28 07:48:56','2019-04-28 07:48:56',0),(25,'斗牛梗','','','2019-04-28 07:48:56','2019-04-28 07:48:56',0),(26,'丰山犬','','','2019-04-28 07:48:56','2019-04-28 07:48:56',0),(27,'中华田园犬','','','2019-04-28 07:48:56','2019-04-28 07:48:56',0),(28,'边境牧羊犬','','','2019-04-28 07:49:13','2019-04-28 07:49:13',0),(29,'串串狗','','','2019-04-28 07:49:13','2019-04-28 07:49:13',0),(30,'史宾格犬','','','2019-04-28 07:49:13','2019-04-28 07:49:13',0),(31,'蝴蝶犬','','','2019-04-28 07:49:13','2019-04-28 07:49:13',0),(32,'日本狆','','','2019-04-28 07:49:13','2019-04-28 07:49:13',0),(34,'布偶猫','','','2019-04-28 07:49:32','2019-04-28 07:49:32',0),(35,'狮子猫','','','2019-04-28 07:49:32','2019-04-28 07:49:32',0),(36,'挪威森林猫','','','2019-04-28 07:49:32','2019-04-28 07:49:32',0),(37,'金吉拉猫','','','2019-04-28 07:49:32','2019-04-28 07:49:32',0),(38,'褴褛猫','','','2019-04-28 07:49:32','2019-04-28 07:49:32',0),(39,'伯曼猫','','','2019-04-28 07:49:32','2019-04-28 07:49:32',0),(40,'丝毛狗','','','2019-04-28 07:50:02','2019-04-28 07:50:02',0),(41,'博美','','','2019-04-28 07:50:02','2019-04-28 07:50:02',0),(42,'德国狐狸犬','','','2019-04-28 07:50:02','2019-04-28 07:50:02',0),(43,'荷兰毛狮犬','','','2019-04-28 07:50:02','2019-04-28 07:50:02',0),(45,'爱斯基摩犬','','','2019-04-28 07:50:02','2019-04-28 07:50:02',0);
/*!40000 ALTER TABLE `tbl_category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tbl_identification`
--

DROP TABLE IF EXISTS `tbl_identification`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `tbl_identification` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL,
  `image` varchar(500) NOT NULL,
  `last` varchar(1000) NOT NULL DEFAULT '',
  `result_list` varchar(2000) NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `image` (`image`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='识别结果';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tbl_identification`
--

LOCK TABLES `tbl_identification` WRITE;
/*!40000 ALTER TABLE `tbl_identification` DISABLE KEYS */;
INSERT INTO `tbl_identification` VALUES (10,1,'a746bed0eb6f35b388a6fe3b4e8cfdfb','','','2019-04-19 03:09:52','2019-04-19 03:09:52'),(11,1,'75286fab27c6b5b57030514b837a6049','','','2019-04-19 03:24:05','2019-04-19 03:24:05'),(12,1,'67179e2fd20aa58cd1024067bbfc24ef','','','2019-04-19 03:24:06','2019-04-19 03:24:06'),(13,1,'ab9b29756e14d008bed50d5260463aa6','','','2019-04-19 03:24:08','2019-04-19 03:24:08'),(14,1,'c8640365937e5e233ca61cbe38146b57','','','2019-04-19 03:24:13','2019-04-19 03:24:13'),(15,1,'0904b89aa88c93849c1f7dba09531873','','','2019-04-19 03:24:16','2019-04-19 03:24:16'),(16,1,'2f8236adb42687e89709ef0cd50d7c44','','','2019-04-19 07:08:23','2019-04-19 07:08:23'),(17,1,'91cfac51baaeccddc381657ac759f4a1','','','2019-04-19 07:11:04','2019-04-19 07:11:04'),(18,1,'/ai/animal/734288ae1599e3f56b10198771782b97','','','2019-04-19 08:20:07','2019-04-19 08:20:07'),(19,1,'/ai/animal/5822766191d3d49036c94ca0d08f3e32','','','2019-04-19 09:29:24','2019-04-19 09:29:24'),(20,1,'e69df8b9f9164a9fe7e921d668a19606','{\"score\":\"0.145458\",\"name\":\"意大利灵缇犬\",\"baike_info\":{\"baike_url\":\"http://baike.baidu.com/item/%E6%A0%BC%E5%8A%9B%E7%8A%AC/10863882\",\"image_url\":\"http://imgsrc.baidu.com/baike/pic/item/8601a18b87d6277ff5fca52121381f30e924fc39.jpg\",\"description\":\"格力犬(greadog)又名灵缇(Greyhound)，原产于中东地区。成年犬身高68-76公分，体重27-32公斤，是极难得的纯种犬。近代世界公认澳洲灵缇最为优秀，各国赛狗均选自澳洲。\"}}','[{\"score\":\"0.145458\",\"name\":\"意大利灵缇犬\",\"baike_info\":{\"baike_url\":\"http://baike.baidu.com/item/%E6%A0%BC%E5%8A%9B%E7%8A%AC/10863882\",\"image_url\":\"http://imgsrc.baidu.com/baike/pic/item/8601a18b87d6277ff5fca52121381f30e924fc39.jpg\",\"description\":\"格力犬(greadog)又名灵缇(Greyhound)，原产于中东地区。成年犬身高68-76公分，体重27-32公斤，是极难得的纯种犬。近代世界公认澳洲灵缇最为优秀，各国赛狗均选自澳洲。\"}},{\"score\":\"0.0931281\",\"name\":\"比特狗\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.0771934\",\"name\":\"丹麦布罗荷马獒\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.0705703\",\"name\":\"阿根廷杜高犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.0705401\",\"name\":\"惠比特犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.0451985\",\"name\":\"牛头梗\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}}]','2019-04-19 09:31:41','2019-04-19 09:31:41'),(21,1,'4c3511018fed9308d922b8f01f4f8951','','','2019-04-22 09:26:10','2019-04-22 09:26:10'),(22,1,'be52aa2d48d5d928092b20735a82ee2e','','','2019-04-22 09:39:50','2019-04-22 09:39:50'),(23,3,'36cf3a105bac557c4bbce57c530ca064','','','2019-04-28 05:59:16','2019-04-28 05:59:16'),(24,3,'c0792d17763102850b2d49196f745f46fa92d438c1c63ee21a4ad5a52b2e1209','{\"score\":\"0.999984\",\"name\":\"非动物\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}}','[{\"score\":\"0.999984\",\"name\":\"非动物\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}}]','2019-04-28 06:03:54','2019-04-28 07:01:37'),(25,3,'9f942e0b38df8a7050191f23c1ce77085146175715039ee89a7d25f49747c5b3','{\"score\":\"0.99755\",\"name\":\"牛头梗\",\"baike_info\":{\"baike_url\":\"http://baike.baidu.com/item/%E7%89%9B%E5%A4%B4%E6%A2%97/469635\",\"image_url\":\"http://imgsrc.baidu.com/baike/pic/item/4e4a20a4462309f7bf6a9e31750e0cf3d7cad68f.jpg\",\"description\":\"牛头梗(Bull Terrier)这一品种有两个品系——白色牛头梗和有色牛头梗。该品种的起源可追溯到1835年。大家一致相信该品种是由现已绝种的英国白梗繁育产生的。几年后，为达到其外观上的标准，该犬与西班牙指示犬进行杂交，至今仍可在该犬的身上偶尔发现指示犬的遗传特征。牛头梗是为绅士们饲养的，为此这些绅士强烈要求公平竞赛，大家不要与弄虚作假的人进行比赛。牛头梗被教以勇敢地进行自卫和保护它的主人，但决不允许它去尝试挑战——白色的品种被称作“白色的骑士”，这是它保持至今的荣誉。不了解该犬的人并不知道牛头梗其实是一种很友善的犬，它就是靠其性情才繁衍兴旺，有时也争斗和嬉戏。多数人更喜欢一种相对平衡的动物，而不是在某些方面古怪的动物。但总的来讲，牛头梗拥有作为一种斗犬的完善组合——活泼和机敏。牛头梗天生性情活跃兴奋度极高，个体犬只或经过人为的训练后具有争斗性，在犬类中从不让步，甚至伤害其它犬类。但牛头梗相对人来说性格还是比较温顺、聪明听话，对主人忠心而且服从性强，对儿童特别和善友好，亲切耐心，牛头梗是具有贵宾的性格跟斗犬身材的结合体如照顾得 当，可成为忠实的家庭守卫犬。作为一种伴侣犬或者说是犬类中的斗士，牛头梗必须强壮、敏捷而且要勇敢。\"}}','[{\"score\":\"0.99755\",\"name\":\"牛头梗\",\"baike_info\":{\"baike_url\":\"http://baike.baidu.com/item/%E7%89%9B%E5%A4%B4%E6%A2%97/469635\",\"image_url\":\"http://imgsrc.baidu.com/baike/pic/item/4e4a20a4462309f7bf6a9e31750e0cf3d7cad68f.jpg\",\"description\":\"牛头梗(Bull Terrier)这一品种有两个品系——白色牛头梗和有色牛头梗。该品种的起源可追溯到1835年。大家一致相信该品种是由现已绝种的英国白梗繁育产生的。几年后，为达到其外观上的标准，该犬与西班牙指示犬进行杂交，至今仍可在该犬的身上偶尔发现指示犬的遗传特征。牛头梗是为绅士们饲养的，为此这些绅士强烈要求公平竞赛，大家不要与弄虚作假的人进行比赛。牛头梗被教以勇敢地进行自卫和保护它的主人，但决不允许它去尝试挑战——白色的品种被称作“白色的骑士”，这是它保持至今的荣誉。不了解该犬的人并不知道牛头梗其实是一种很友善的犬，它就是靠其性情才繁衍兴旺，有时也争斗和嬉戏。多数人更喜欢一种相对平衡的动物，而不是在某些方面古怪的动物。但总的来讲，牛头梗拥有作为一种斗犬的完善组合——活泼和机敏。牛头梗天生性情活跃兴奋度极高，个体犬只或经过人为的训练后具有争斗性，在犬类中从不让步，甚至伤害其它犬类。但牛头梗相对人来说性格还是比较温顺、聪明听话，对主人忠心而且服从性强，对儿童特别和善友好，亲切耐心，牛头梗是具有贵宾的性格跟斗犬身材的结合体如照顾得 当，可成为忠实的家庭守卫犬。作为一种伴侣犬或者说是犬类中的斗士，牛头梗必须强壮、敏捷而且要勇敢。\"}},{\"score\":\"0.000685228\",\"name\":\"斯塔福斗牛梗\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.000154363\",\"name\":\"阿根廷杜高犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.000117852\",\"name\":\"古梗犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.000104242\",\"name\":\"平毛猎狐梗\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"9.68728e-05\",\"name\":\"杰克罗素梗\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}}]','2019-04-28 06:37:30','2019-04-28 07:44:42'),(26,3,'50d9f5e4bd2b82f0056dd4e7c915ad7a67d1fd24596d33a3ea4a9bee303ef3d5','{\"score\":\"0.189952\",\"name\":\"纪州犬\",\"baike_info\":{\"baike_url\":\"http://baike.baidu.com/item/%E7%BA%AA%E5%B7%9E%E7%8A%AC/164424\",\"image_url\":\"http://imgsrc.baidu.com/baike/pic/item/b21bb051f81986184e684b3149ed2e738bd4e65c.jpg\",\"description\":\"纪州犬很早以前就在日本的和歌山县出现。三重县为中心的山岳地带 ，被作为猎猪、猎鹿的犬种。力气强大，外貌洋溢着野性味道，尾巴卷曲，后肢具有狼爪。脸部及额头稍微宽广。脸颊很显著，眼睛也斜上方微微吊起，肌肉结实精悍。毛色以白色居多，也有红色、芝麻色和虎斑色等等。此犬对陌生人戒备性强，对主人衷心。勇敢性比较突出，有警戒心的原因，让人看了不敢接近。是看家护院的优秀犬种。产地为日本纪伊半岛。\"}}','[{\"score\":\"0.189952\",\"name\":\"纪州犬\",\"baike_info\":{\"baike_url\":\"http://baike.baidu.com/item/%E7%BA%AA%E5%B7%9E%E7%8A%AC/164424\",\"image_url\":\"http://imgsrc.baidu.com/baike/pic/item/b21bb051f81986184e684b3149ed2e738bd4e65c.jpg\",\"description\":\"纪州犬很早以前就在日本的和歌山县出现。三重县为中心的山岳地带 ，被作为猎猪、猎鹿的犬种。力气强大，外貌洋溢着野性味道，尾巴卷曲，后肢具有狼爪。脸部及额头稍微宽广。脸颊很显著，眼睛也斜上方微微吊起，肌肉结实精悍。毛色以白色居多，也有红色、芝麻色和虎斑色等等。此犬对陌生人戒备性强，对主人衷心。勇敢性比较突出，有警戒心的原因，让人看了不敢接近。是看家护院的优秀犬种。产地为日本纪伊半岛。\"}},{\"score\":\"0.13561\",\"name\":\"下司犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.134946\",\"name\":\"广东潮州犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.133856\",\"name\":\"斗牛梗\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.0710248\",\"name\":\"丰山犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.0653756\",\"name\":\"中华田园犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}}]','2019-04-28 07:45:33','2019-04-28 07:45:37'),(27,3,'f927a6856b81347c767b1b2641c2fe7fc4e3d09f445c1a603aecf59e3c78991b','{\"score\":\"0.683441\",\"name\":\"边境牧羊犬\",\"baike_info\":{\"baike_url\":\"http://baike.baidu.com/item/%E8%BE%B9%E5%A2%83%E7%89%A7%E7%BE%8A%E7%8A%AC/406513\",\"image_url\":\"http://imgsrc.baidu.com/baike/pic/item/d0c8a786c9177f3e2b44aa517dcf3bc79f3d567d.jpg\",\"description\":\"边境牧羊犬(Border Collie)，原产于苏格兰边境，为柯利牧羊犬的一种，具有强烈的牧羊本能，天性聪颖、善于察言观色，能准确明白主人的指示，可借由眼神的注视而驱动羊群移动或旋转，被当成牧羊犬已有多年的历史，在世界犬种智商排行第一名。特点是聪明、学习能力强、理解力高、容易训练、善于和主人沟通、温和、忠诚、服从性好，其忠心程度可以用如影随形来形容。由于温和忠诚的性格不乱叫，一度成为最受城市人口欢迎的宠物。而且边境牧羊犬还是飞盘狗最具竞争力的犬种，是历年飞盘狗世界杯大赛的主角。\"}}','[{\"score\":\"0.683441\",\"name\":\"边境牧羊犬\",\"baike_info\":{\"baike_url\":\"http://baike.baidu.com/item/%E8%BE%B9%E5%A2%83%E7%89%A7%E7%BE%8A%E7%8A%AC/406513\",\"image_url\":\"http://imgsrc.baidu.com/baike/pic/item/d0c8a786c9177f3e2b44aa517dcf3bc79f3d567d.jpg\",\"description\":\"边境牧羊犬(Border Collie)，原产于苏格兰边境，为柯利牧羊犬的一种，具有强烈的牧羊本能，天性聪颖、善于察言观色，能准确明白主人的指示，可借由眼神的注视而驱动羊群移动或旋转，被当成牧羊犬已有多年的历史，在世界犬种智商排行第一名。特点是聪明、学习能力强、理解力高、容易训练、善于和主人沟通、温和、忠诚、服从性好，其忠心程度可以用如影随形来形容。由于温和忠诚的性格不乱叫，一度成为最受城市人口欢迎的宠物。而且边境牧羊犬还是飞盘狗最具竞争力的犬种，是历年飞盘狗世界杯大赛的主角。\"}},{\"score\":\"0.0804629\",\"name\":\"串串狗\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.0801911\",\"name\":\"史宾格犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.0423184\",\"name\":\"蝴蝶犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.0247415\",\"name\":\"日本狆\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.00628662\",\"name\":\"中华田园犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}}]','2019-04-28 07:45:45','2019-04-28 07:45:49'),(28,3,'3894cf0a7808a532c22ca0a431e2690437de87baf04e3a0129adf9840a04516a','','','2019-04-28 07:46:03','2019-04-28 07:46:03'),(29,3,'128f7d1b18fe81e500dafa659d9864582aba9913eb4fc2f95f4299ac8a1cf77c','{\"score\":\"0.767868\",\"name\":\"串串狗\",\"baike_info\":{\"baike_url\":\"http://baike.baidu.com/item/%E4%B8%B2%E4%B8%B2%E7%8B%97/2768810\",\"image_url\":\"http://imgsrc.baidu.com/baike/pic/item/d50735fae6cd7b89bc99b910032442a7d8330eaf.jpg\",\"description\":\"串串狗是所有杂交品种狗或品种特性不明显的狗的总称，不是某个具体的犬种，基因不纯\"}}','[{\"score\":\"0.767868\",\"name\":\"串串狗\",\"baike_info\":{\"baike_url\":\"http://baike.baidu.com/item/%E4%B8%B2%E4%B8%B2%E7%8B%97/2768810\",\"image_url\":\"http://imgsrc.baidu.com/baike/pic/item/d50735fae6cd7b89bc99b910032442a7d8330eaf.jpg\",\"description\":\"串串狗是所有杂交品种狗或品种特性不明显的狗的总称，不是某个具体的犬种，基因不纯\"}},{\"score\":\"0.0637682\",\"name\":\"中华田园犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.0449601\",\"name\":\"库达犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.0336948\",\"name\":\"蝴蝶犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.011418\",\"name\":\"美系秋田犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.010579\",\"name\":\"杰克罗素梗\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}}]','2019-04-28 07:46:13','2019-04-28 07:46:19'),(30,3,'0e24df2b5bbee46084ac9fa32451c0777b2bac205e85416c9da56f9c1273d2de','{\"score\":\"0.978097\",\"name\":\"布偶猫\",\"baike_info\":{\"baike_url\":\"http://baike.baidu.com/item/%E5%B8%83%E5%81%B6%E7%8C%AB/642121\",\"image_url\":\"http://imgsrc.baidu.com/baike/pic/item/4034970a304e251fc3ec88c8af86c9177f3e53e2.jpg\",\"description\":\"布偶猫，又称“布拉多尔(Ragdoll)”，发源于美国，是一种杂交品种宠物猫。是现存体型最大、体重最重的猫之一。头呈楔形，眼大而圆，被毛丰厚，四肢较长且富有肉感，尾长，身体柔软，毛色有重点色、手套色或双色等等。布偶猫较为温顺好静，对人友善。其美丽优雅又非常类似于狗的性格(Puppy cat)而又被称为\\\"仙女猫\\\"，\\\"小狗猫\\\"。特殊的外貌和温和的性格是布偶猫最大的特点之一。\"}}','[{\"score\":\"0.978097\",\"name\":\"布偶猫\",\"baike_info\":{\"baike_url\":\"http://baike.baidu.com/item/%E5%B8%83%E5%81%B6%E7%8C%AB/642121\",\"image_url\":\"http://imgsrc.baidu.com/baike/pic/item/4034970a304e251fc3ec88c8af86c9177f3e53e2.jpg\",\"description\":\"布偶猫，又称“布拉多尔(Ragdoll)”，发源于美国，是一种杂交品种宠物猫。是现存体型最大、体重最重的猫之一。头呈楔形，眼大而圆，被毛丰厚，四肢较长且富有肉感，尾长，身体柔软，毛色有重点色、手套色或双色等等。布偶猫较为温顺好静，对人友善。其美丽优雅又非常类似于狗的性格(Puppy cat)而又被称为\\\"仙女猫\\\"，\\\"小狗猫\\\"。特殊的外貌和温和的性格是布偶猫最大的特点之一。\"}},{\"score\":\"0.00522007\",\"name\":\"狮子猫\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.0049363\",\"name\":\"挪威森林猫\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.00211385\",\"name\":\"金吉拉猫\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.00170313\",\"name\":\"褴褛猫\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.000621644\",\"name\":\"伯曼猫\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}}]','2019-04-28 07:46:34','2019-04-28 07:46:38'),(31,3,'b3a81ecf15be3437551d5d147cb032bd5af024781c7532add4068b1e32f3c00f','{\"score\":\"0.76537\",\"name\":\"丝毛狗\",\"baike_info\":{\"baike_url\":\"http://baike.baidu.com/item/%E4%B8%9D%E6%AF%9B%E7%8B%97/901957\",\"image_url\":\"http://imgsrc.baidu.com/baike/pic/item/b219ebc4b74543a957b0e62c1e178a82b80114a2.jpg\",\"description\":\"丝毛犬身材虽然矮小一些，但食量却多于其他种犬的，每天至少需供给肉一百八十克，不少的食量对于我们来说会是负担吗，这个就要因人而异了。还有就是除了肉食之外，还需给等量的素食，少糖的饼干也是可以替代的。\"}}','[{\"score\":\"0.76537\",\"name\":\"丝毛狗\",\"baike_info\":{\"baike_url\":\"http://baike.baidu.com/item/%E4%B8%9D%E6%AF%9B%E7%8B%97/901957\",\"image_url\":\"http://imgsrc.baidu.com/baike/pic/item/b219ebc4b74543a957b0e62c1e178a82b80114a2.jpg\",\"description\":\"丝毛犬身材虽然矮小一些，但食量却多于其他种犬的，每天至少需供给肉一百八十克，不少的食量对于我们来说会是负担吗，这个就要因人而异了。还有就是除了肉食之外，还需给等量的素食，少糖的饼干也是可以替代的。\"}},{\"score\":\"0.0701848\",\"name\":\"博美\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.0130148\",\"name\":\"德国狐狸犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.00934069\",\"name\":\"荷兰毛狮犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.00475413\",\"name\":\"蝴蝶犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.00423595\",\"name\":\"爱斯基摩犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}}]','2019-04-28 07:46:50','2019-04-28 07:46:54'),(32,3,'4287382588e9c5251004f13453b750838849ca3a21b86ffbda7f4be82df53e00','{\"score\":\"0.746315\",\"name\":\"贵宾犬/贵妇犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}}','[{\"score\":\"0.746315\",\"name\":\"贵宾犬/贵妇犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.208121\",\"name\":\"台湾犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.0334318\",\"name\":\"迷你贵宾犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.0030288\",\"name\":\"卷毛比熊犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.0010864\",\"name\":\"杂交狮子狗\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.00098021\",\"name\":\"可卡\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}}]','2019-04-28 07:47:03','2019-04-28 07:48:35'),(33,3,'270ba863cbcc53491f0eaaf9a2af3a4784a39d308d25bf3d91dba83b4aa28df7','{\"score\":\"0.189952\",\"name\":\"纪州犬\",\"baike_info\":{\"baike_url\":\"http://baike.baidu.com/item/%E7%BA%AA%E5%B7%9E%E7%8A%AC/164424\",\"image_url\":\"http://imgsrc.baidu.com/baike/pic/item/b21bb051f81986184e684b3149ed2e738bd4e65c.jpg\",\"description\":\"纪州犬很早以前就在日本的和歌山县出现。三重县为中心的山岳地带 ，被作为猎猪、猎鹿的犬种。力气强大，外貌洋溢着野性味道，尾巴卷曲，后肢具有狼爪。脸部及额头稍微宽广。脸颊很显著，眼睛也斜上方微微吊起，肌肉结实精悍。毛色以白色居多，也有红色、芝麻色和虎斑色等等。此犬对陌生人戒备性强，对主人衷心。勇敢性比较突出，有警戒心的原因，让人看了不敢接近。是看家护院的优秀犬种。产地为日本纪伊半岛。\"}}','[{\"score\":\"0.189952\",\"name\":\"纪州犬\",\"baike_info\":{\"baike_url\":\"http://baike.baidu.com/item/%E7%BA%AA%E5%B7%9E%E7%8A%AC/164424\",\"image_url\":\"http://imgsrc.baidu.com/baike/pic/item/b21bb051f81986184e684b3149ed2e738bd4e65c.jpg\",\"description\":\"纪州犬很早以前就在日本的和歌山县出现。三重县为中心的山岳地带 ，被作为猎猪、猎鹿的犬种。力气强大，外貌洋溢着野性味道，尾巴卷曲，后肢具有狼爪。脸部及额头稍微宽广。脸颊很显著，眼睛也斜上方微微吊起，肌肉结实精悍。毛色以白色居多，也有红色、芝麻色和虎斑色等等。此犬对陌生人戒备性强，对主人衷心。勇敢性比较突出，有警戒心的原因，让人看了不敢接近。是看家护院的优秀犬种。产地为日本纪伊半岛。\"}},{\"score\":\"0.13561\",\"name\":\"下司犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.134946\",\"name\":\"广东潮州犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.133856\",\"name\":\"斗牛梗\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.0710248\",\"name\":\"丰山犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.0653756\",\"name\":\"中华田园犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}}]','2019-04-28 07:48:49','2019-04-28 07:48:56'),(34,3,'f081cc6c5889a4cc8ca32d42943d9c70ca32adcecfcef49e8fb3a86ff59a07c7','{\"score\":\"0.683441\",\"name\":\"边境牧羊犬\",\"baike_info\":{\"baike_url\":\"http://baike.baidu.com/item/%E8%BE%B9%E5%A2%83%E7%89%A7%E7%BE%8A%E7%8A%AC/406513\",\"image_url\":\"http://imgsrc.baidu.com/baike/pic/item/d0c8a786c9177f3e2b44aa517dcf3bc79f3d567d.jpg\",\"description\":\"边境牧羊犬(Border Collie)，原产于苏格兰边境，为柯利牧羊犬的一种，具有强烈的牧羊本能，天性聪颖、善于察言观色，能准确明白主人的指示，可借由眼神的注视而驱动羊群移动或旋转，被当成牧羊犬已有多年的历史，在世界犬种智商排行第一名。特点是聪明、学习能力强、理解力高、容易训练、善于和主人沟通、温和、忠诚、服从性好，其忠心程度可以用如影随形来形容。由于温和忠诚的性格不乱叫，一度成为最受城市人口欢迎的宠物。而且边境牧羊犬还是飞盘狗最具竞争力的犬种，是历年飞盘狗世界杯大赛的主角。\"}}','[{\"score\":\"0.683441\",\"name\":\"边境牧羊犬\",\"baike_info\":{\"baike_url\":\"http://baike.baidu.com/item/%E8%BE%B9%E5%A2%83%E7%89%A7%E7%BE%8A%E7%8A%AC/406513\",\"image_url\":\"http://imgsrc.baidu.com/baike/pic/item/d0c8a786c9177f3e2b44aa517dcf3bc79f3d567d.jpg\",\"description\":\"边境牧羊犬(Border Collie)，原产于苏格兰边境，为柯利牧羊犬的一种，具有强烈的牧羊本能，天性聪颖、善于察言观色，能准确明白主人的指示，可借由眼神的注视而驱动羊群移动或旋转，被当成牧羊犬已有多年的历史，在世界犬种智商排行第一名。特点是聪明、学习能力强、理解力高、容易训练、善于和主人沟通、温和、忠诚、服从性好，其忠心程度可以用如影随形来形容。由于温和忠诚的性格不乱叫，一度成为最受城市人口欢迎的宠物。而且边境牧羊犬还是飞盘狗最具竞争力的犬种，是历年飞盘狗世界杯大赛的主角。\"}},{\"score\":\"0.0804629\",\"name\":\"串串狗\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.0801911\",\"name\":\"史宾格犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.0423184\",\"name\":\"蝴蝶犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.0247415\",\"name\":\"日本狆\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.00628662\",\"name\":\"中华田园犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}}]','2019-04-28 07:49:08','2019-04-28 07:49:13'),(35,3,'79c33e924468292e2ca523a3bdd924cadcc388125d5f130f84529b721ee0a5cf','{\"score\":\"0.978097\",\"name\":\"布偶猫\",\"baike_info\":{\"baike_url\":\"http://baike.baidu.com/item/%E5%B8%83%E5%81%B6%E7%8C%AB/642121\",\"image_url\":\"http://imgsrc.baidu.com/baike/pic/item/4034970a304e251fc3ec88c8af86c9177f3e53e2.jpg\",\"description\":\"布偶猫，又称“布拉多尔(Ragdoll)”，发源于美国，是一种杂交品种宠物猫。是现存体型最大、体重最重的猫之一。头呈楔形，眼大而圆，被毛丰厚，四肢较长且富有肉感，尾长，身体柔软，毛色有重点色、手套色或双色等等。布偶猫较为温顺好静，对人友善。其美丽优雅又非常类似于狗的性格(Puppy cat)而又被称为\\\"仙女猫\\\"，\\\"小狗猫\\\"。特殊的外貌和温和的性格是布偶猫最大的特点之一。\"}}','[{\"score\":\"0.978097\",\"name\":\"布偶猫\",\"baike_info\":{\"baike_url\":\"http://baike.baidu.com/item/%E5%B8%83%E5%81%B6%E7%8C%AB/642121\",\"image_url\":\"http://imgsrc.baidu.com/baike/pic/item/4034970a304e251fc3ec88c8af86c9177f3e53e2.jpg\",\"description\":\"布偶猫，又称“布拉多尔(Ragdoll)”，发源于美国，是一种杂交品种宠物猫。是现存体型最大、体重最重的猫之一。头呈楔形，眼大而圆，被毛丰厚，四肢较长且富有肉感，尾长，身体柔软，毛色有重点色、手套色或双色等等。布偶猫较为温顺好静，对人友善。其美丽优雅又非常类似于狗的性格(Puppy cat)而又被称为\\\"仙女猫\\\"，\\\"小狗猫\\\"。特殊的外貌和温和的性格是布偶猫最大的特点之一。\"}},{\"score\":\"0.00522007\",\"name\":\"狮子猫\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.0049363\",\"name\":\"挪威森林猫\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.00211385\",\"name\":\"金吉拉猫\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.00170313\",\"name\":\"褴褛猫\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.000621644\",\"name\":\"伯曼猫\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}}]','2019-04-28 07:49:28','2019-04-28 07:49:32'),(36,3,'09342dcadbacafb22aa0e848df16da1be0895307f1469ccc9d969e3cc8efc290','','','2019-04-28 07:49:54','2019-04-28 07:49:54'),(37,3,'0cd3fcb87c2279b0797c766d99c45a10338858568c83f9019b459592b881289b','','','2019-04-28 07:49:54','2019-04-28 07:49:54'),(38,3,'ede1e79cb28cca4561dee5f9a8d6bf9ee03bc8c76573be67c21432e8db53100d','{\"score\":\"0.76537\",\"name\":\"丝毛狗\",\"baike_info\":{\"baike_url\":\"http://baike.baidu.com/item/%E4%B8%9D%E6%AF%9B%E7%8B%97/901957\",\"image_url\":\"http://imgsrc.baidu.com/baike/pic/item/b219ebc4b74543a957b0e62c1e178a82b80114a2.jpg\",\"description\":\"丝毛犬身材虽然矮小一些，但食量却多于其他种犬的，每天至少需供给肉一百八十克，不少的食量对于我们来说会是负担吗，这个就要因人而异了。还有就是除了肉食之外，还需给等量的素食，少糖的饼干也是可以替代的。\"}}','[{\"score\":\"0.76537\",\"name\":\"丝毛狗\",\"baike_info\":{\"baike_url\":\"http://baike.baidu.com/item/%E4%B8%9D%E6%AF%9B%E7%8B%97/901957\",\"image_url\":\"http://imgsrc.baidu.com/baike/pic/item/b219ebc4b74543a957b0e62c1e178a82b80114a2.jpg\",\"description\":\"丝毛犬身材虽然矮小一些，但食量却多于其他种犬的，每天至少需供给肉一百八十克，不少的食量对于我们来说会是负担吗，这个就要因人而异了。还有就是除了肉食之外，还需给等量的素食，少糖的饼干也是可以替代的。\"}},{\"score\":\"0.0701848\",\"name\":\"博美\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.0130148\",\"name\":\"德国狐狸犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.00934069\",\"name\":\"荷兰毛狮犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.00475413\",\"name\":\"蝴蝶犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.00423595\",\"name\":\"爱斯基摩犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}}]','2019-04-28 07:49:57','2019-04-28 07:50:02'),(39,3,'e5c66225477781412fada38aab54356e3573964885c8cd54545e98eda9fc2a87','{\"score\":\"0.746315\",\"name\":\"贵宾犬/贵妇犬\",\"baike_info\":{\"baike_url\":\"http://baike.baidu.com/item/%E4%B8%9D%E6%AF%9B%E7%8B%97/901957\",\"image_url\":\"http://imgsrc.baidu.com/baike/pic/item/b219ebc4b74543a957b0e62c1e178a82b80114a2.jpg\",\"description\":\"丝毛犬身材虽然矮小一些，但食量却多于其他种犬的，每天至少需供给肉一百八十克，不少的食量对于我们来说会是负担吗，这个就要因人而异了。还有就是除了肉食之外，还需给等量的素食，少糖的饼干也是可以替代的。\"}}','[{\"score\":\"0.746315\",\"name\":\"贵宾犬/贵妇犬\",\"baike_info\":{\"baike_url\":\"http://baike.baidu.com/item/%E4%B8%9D%E6%AF%9B%E7%8B%97/901957\",\"image_url\":\"http://imgsrc.baidu.com/baike/pic/item/b219ebc4b74543a957b0e62c1e178a82b80114a2.jpg\",\"description\":\"丝毛犬身材虽然矮小一些，但食量却多于其他种犬的，每天至少需供给肉一百八十克，不少的食量对于我们来说会是负担吗，这个就要因人而异了。还有就是除了肉食之外，还需给等量的素食，少糖的饼干也是可以替代的。\"}},{\"score\":\"0.208121\",\"name\":\"台湾犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.0334318\",\"name\":\"迷你贵宾犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.0030288\",\"name\":\"卷毛比熊犬\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.0010864\",\"name\":\"杂交狮子狗\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}},{\"score\":\"0.00098021\",\"name\":\"可卡\",\"baike_info\":{\"baike_url\":\"\",\"image_url\":\"\",\"description\":\"\"}}]','2019-04-28 07:50:43','2019-04-28 07:51:44');
/*!40000 ALTER TABLE `tbl_identification` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tbl_pet`
--

DROP TABLE IF EXISTS `tbl_pet`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `tbl_pet` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cid` int(11) NOT NULL DEFAULT '0' COMMENT 'category id',
  `user_id` int(11) DEFAULT '0',
  `nickname` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '昵称',
  `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '头像',
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  PRIMARY KEY (`id`),
  KEY `cid` (`cid`),
  KEY `userid` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='宠物';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tbl_pet`
--

LOCK TABLES `tbl_pet` WRITE;
/*!40000 ALTER TABLE `tbl_pet` DISABLE KEYS */;
INSERT INTO `tbl_pet` VALUES (23,28,3,'芝麻','f927a6856b81347c767b1b2641c2fe7fc4e3d09f445c1a603aecf59e3c78991b','2019-04-28 09:26:19','2019-04-28 09:26:19');
/*!40000 ALTER TABLE `tbl_pet` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tbl_sms`
--

DROP TABLE IF EXISTS `tbl_sms`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `tbl_sms` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `account` varchar(20) DEFAULT NULL,
  `sms` varchar(255) DEFAULT NULL,
  `mode` varchar(45) DEFAULT NULL,
  `effective` int(10) DEFAULT NULL COMMENT '有效时间-单位/秒',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `status` (`status`),
  KEY `account` (`account`) /*!80000 INVISIBLE */,
  KEY `mode` (`mode`) /*!80000 INVISIBLE */
) ENGINE=InnoDB AUTO_INCREMENT=84 DEFAULT CHARSET=utf8 COMMENT='验证码';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tbl_sms`
--

LOCK TABLES `tbl_sms` WRITE;
/*!40000 ALTER TABLE `tbl_sms` DISABLE KEYS */;
INSERT INTO `tbl_sms` VALUES (76,'18511566625','633911','REGISTER',1800,'2019-04-26 08:46:53','2019-04-28 03:23:05',1),(77,'18511566625','122492','REGISTER',1800,'2019-04-26 08:49:40','2019-04-28 03:23:05',1),(78,'18511566625','873452','REGISTER',1800,'2019-04-26 08:50:41','2019-04-28 03:23:05',1),(79,'18511566625','544366','REGISTER',1800,'2019-04-26 08:57:57','2019-04-28 03:23:05',1),(80,'18511566625','858661','REGISTER',1800,'2019-04-26 09:29:19','2019-04-28 03:23:05',1),(81,'18511566625','722243','REGISTER',1800,'2019-04-26 10:01:10','2019-04-28 03:23:05',1),(82,'18511566625','461525','REGISTER',1800,'2019-04-28 03:07:18','2019-04-28 03:23:05',1),(83,'18511566625','573419','REGISTER',1800,'2019-04-28 03:23:05','2019-04-28 03:23:05',0);
/*!40000 ALTER TABLE `tbl_sms` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tbl_user`
--

DROP TABLE IF EXISTS `tbl_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `tbl_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `account` varchar(20) NOT NULL COMMENT '账号',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `salting` varchar(255) NOT NULL COMMENT '盐',
  `ticket` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '票据',
  `nickname` varchar(255) NOT NULL DEFAULT '' COMMENT '昵称',
  `status` tinyint(1) NOT NULL COMMENT '状态：0=正常 1=无效',
  `created_at` timestamp NOT NULL COMMENT '创建时间',
  `updated_at` timestamp NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='user';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tbl_user`
--

LOCK TABLES `tbl_user` WRITE;
/*!40000 ALTER TABLE `tbl_user` DISABLE KEYS */;
INSERT INTO `tbl_user` VALUES (3,'18511566625','2f18b55739a3325964b544f6959c3879','KJaOhVdFnVpUbyjaHSHfkULBp','0e76870dcc87a605e5f3eedacac597dc','KJaOhVdFnVpUbyjaHSHfkULBp',0,'2019-04-28 03:23:53','2019-04-28 10:17:39');
/*!40000 ALTER TABLE `tbl_user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-04-28 18:39:13
