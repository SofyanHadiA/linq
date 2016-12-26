CREATE TABLE `products` (
 `uid` varchar(36) NOT NULL,
 `title` varchar(200) NOT NULL,
 `category` varchar(36) NOT NULL,
 `buy_price` decimal(10,0) NOT NULL,
 `sell_price` decimal(10,0) NOT NULL,
 `image` varchar(250) NOT NULL,
 `stock` int(11) NOT NULL,
 `code` varchar(200) NOT NULL,
 `deleted` tinyint(1) NOT NULL,
 `created` datetime NOT NULL,
 `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 PRIMARY KEY (`uid`),
 UNIQUE KEY `titile` (`title`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

	
CREATE TABLE `product_categories` (
 `uid` varchar(36) NOT NULL,
 `title` varchar(200) NOT NULL,
 `description` text NOT NULL,
 `slug` varchar(200) NOT NULL,
 `deleted` tinyint(1) NOT NULL DEFAULT '0',
 `created` datetime NOT NULL,
 `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;