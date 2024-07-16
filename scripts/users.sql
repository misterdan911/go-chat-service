DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `display_name` varchar(50) DEFAULT NULL,
  `phone_number` varchar(15) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `phone_number` (`phone_number`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

INSERT INTO `users` (`id`, `username`, `email`, `password`, `display_name`, `phone_number`, `created_at`, `updated_at`) VALUES
(1, 'danu', 'dciptadi@gmail.com', '$2a$14$wQP/JBKBWn0sVYKR3/3Fb.C7wL.cQZda3Pls2xE4GlJjP40VUprxC', 'Danu Ciptadi', NULL, NULL, NULL);
