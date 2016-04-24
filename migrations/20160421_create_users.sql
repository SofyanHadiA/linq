-- up!

CREATE TABLE users
(
  `uid` varchar(16),
  `email` varchar(100),
  `username` varchar(25),
  `first_name` varchar(25),
  `last_name` varchar(25),
  `password` varchar(100),
  `lastLogin` int(11),
  `forgot_code` text NOT NULL,
  `notification_created` int(11) DEFAULT NULL,
  `last_login` int(13) DEFAULT NULL,
  `address` varchar(100),
  
  PRIMARY KEY (`uid`)
);
ALTER TABLE users ADD city varchar(100);
ALTER TABLE users ADD state varchar(100);
ALTER TABLE users ADD country varchar(100);
ALTER TABLE users ADD zip varchar(10);
ALTER TABLE users ADD phone_number varchar(20);
ALTER TABLE users modify uid varchar(36) not null;
ALTER TABLE users ADD photo varchar(200) ;