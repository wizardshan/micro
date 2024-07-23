CREATE TABLE `users` (
     `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
     `hash_id` varchar(20) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
     `mobile` varchar(11) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
     `password` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
     `level` int(11) NOT NULL DEFAULT '0',
     `nickname` varchar(20) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
     `avatar` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
     `bio` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
     `amount` int(11) NOT NULL DEFAULT '0',
     `status` varchar(10) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
     `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
     `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
     PRIMARY KEY (`id`),
     UNIQUE KEY `mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin