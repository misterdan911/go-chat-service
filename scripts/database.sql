DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `phone_number` varchar(15) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `phone_number` (`phone_number`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;


DROP TABLE IF EXISTS `private_chat`;
CREATE TABLE `private_chat` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `sender_id` int(11) DEFAULT NULL,
  `receiver_id` int(11) DEFAULT NULL,
  `message_content` text DEFAULT NULL,
  `media_url` varchar(255) DEFAULT NULL,
  `sent_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `is_sent` tinyint(1) DEFAULT 0,
  `is_delivered` tinyint(1) DEFAULT 0,
  `is_seen` tinyint(1) DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

DROP TABLE IF EXISTS `chat_group`;
CREATE TABLE `chat_group` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `group_name` varchar(50) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

DROP TABLE IF EXISTS `group_member`;
CREATE TABLE group_member (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    chat_group_id INT,
    user_id INT,
    joined_at datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `group_chat`;
CREATE TABLE group_chat (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    chat_group INT,
    sender_id INT,
    message_content TEXT,
    media_url VARCHAR(255),  -- if the message includes media
    sent_at  datetime DEFAULT NULL,
    is_sent TINYINT(1) DEFAULT 0,  -- flag for sent message
    is_delivered TINYINT(1) DEFAULT 0,  -- flag for delivered message
    is_seen TINYINT(1) DEFAULT 0,  -- flag for seen message
    PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `group_message_status`;
CREATE TABLE group_message_status (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    message_id INT,
    user_id INT,
    is_delivered TINYINT(1) DEFAULT 0,  -- flag for delivered message
    is_seen TINYINT(1) DEFAULT 0,  -- flag for seen message
    PRIMARY KEY (id)
);
