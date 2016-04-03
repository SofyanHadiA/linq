SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

--
-- Table structure for table `admin`
--

CREATE TABLE `admin` (
  `admin_id` int(11) NOT NULL AUTO_INCREMENT,
  `admin_username` varchar(200) CHARACTER SET latin1 NOT NULL,
  `admin_password` varchar(300) CHARACTER SET latin1 NOT NULL,
  `admin_email` varchar(300) CHARACTER SET latin1 NOT NULL,
  PRIMARY KEY (`admin_id`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 ROW_FORMAT=FIXED AUTO_INCREMENT=4 ;

--
-- Dumping data for table `admin`
--

INSERT INTO `admin` (`admin_id`, `admin_username`, `admin_password`, `admin_email`) VALUES
(1, 'admin', '25f9e794323b453885f5181f1b624d0b', 'admin@hotmail.com');

-- --------------------------------------------------------

--
-- Table structure for table `blocked_users`
--

CREATE TABLE `blocked_users` (
  `b_id` int(11) NOT NULL AUTO_INCREMENT,
  `blocker_id` int(11) DEFAULT NULL,
  `blocked_id` int(11) DEFAULT NULL,
  `block_status` enum('0','1') NOT NULL,
  `created` int(13) NOT NULL DEFAULT '1411570461',
  PRIMARY KEY (`b_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=latin1 AUTO_INCREMENT=3 ;

-- --------------------------------------------------------

--
-- Table structure for table `comments`
--

CREATE TABLE `comments` (
  `com_id` int(11) NOT NULL AUTO_INCREMENT,
  `comment` text CHARACTER SET utf8,
  `hashtag` varchar(255) DEFAULT NULL,
  `p_mention` varchar(255) DEFAULT NULL,
  `msg_id_fk` int(11) DEFAULT NULL,
  `uid_fk` int(11) DEFAULT NULL,
  `ip` varchar(30) DEFAULT NULL,
  `created` int(11) NOT NULL DEFAULT '1269249260',
  `like_count` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`com_id`)
) ENGINE=MyISAM  DEFAULT CHARSET=latin1 AUTO_INCREMENT=153 ;

-- --------------------------------------------------------

--
-- Table structure for table `comment_like`
--

CREATE TABLE `comment_like` (
  `clike_id` int(11) NOT NULL AUTO_INCREMENT,
  `com_id_fk` int(11) NOT NULL,
  `uid_fk` int(11) NOT NULL,
  `ouid_fk` int(11) DEFAULT NULL,
  `created` int(11) DEFAULT NULL,
  PRIMARY KEY (`clike_id`)
) ENGINE=MyISAM  DEFAULT CHARSET=latin1 AUTO_INCREMENT=22 ;

-- --------------------------------------------------------

--
-- Table structure for table `configuration`
--

CREATE TABLE `configuration` (
  `con_id` int(11) NOT NULL AUTO_INCREMENT,
  `notificationPerPage` int(3) DEFAULT NULL,
  `uploadImage` int(11) DEFAULT NULL,
  `profileWidth` int(11) DEFAULT NULL,
  `friendsWidgetPerPage` int(4) DEFAULT NULL,
  `adminPerPage` int(3) DEFAULT NULL,
  `annoncement` varchar(500) CHARACTER SET utf8 DEFAULT NULL,
  `scriptName` varchar(100) CHARACTER SET utf8 DEFAULT NULL,
  `scriptTitle` varchar(500) CHARACTER SET utf8 DEFAULT NULL,
  `scriptLogo` varchar(300) DEFAULT NULL,
  `forgot` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`con_id`)
) ENGINE=MyISAM  DEFAULT CHARSET=latin1 AUTO_INCREMENT=4 ;

--
-- Dumping data for table `configuration`
--

INSERT INTO `configuration` (`con_id`, `notificationPerPage`, `uploadImage`, `profileWidth`, `friendsWidgetPerPage`, `adminPerPage`, `annoncement`, `scriptName`, `scriptTitle`, `scriptLogo`, `forgot`) VALUES
(1, 10, 50120, 580, 20, 25, 'Hi, visitor. Well come to SocialMatv1.2 Social Networking Platform. We are working hardly for you. You can buy SocialMatv1.2 just 26$. We are giving very good support faster and final results. No need to wait. Just click http://www.codecanyon.net/item/socialmat-social-networking-platform/11734904 link and easily used.', 'SocialMatv1.2', 'The SocialMat is a very Different social networking platform. If you like it please signup or signin', 'logo_1441181158.png', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `conversation`
--

CREATE TABLE `conversation` (
  `c_id` int(11) NOT NULL AUTO_INCREMENT,
  `user_one` int(11) NOT NULL,
  `user_two` int(11) NOT NULL,
  `ip` varchar(30) DEFAULT NULL,
  `time` int(11) DEFAULT NULL,
  PRIMARY KEY (`c_id`)
) ENGINE=MyISAM  DEFAULT CHARSET=latin1 AUTO_INCREMENT=49 ;

-- --------------------------------------------------------

--
-- Table structure for table `conversation_images`
--

CREATE TABLE `conversation_images` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `image_path` varchar(30) DEFAULT NULL,
  `uid_fk` int(11) NOT NULL,
  `group_id_fk` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM  DEFAULT CHARSET=latin1 AUTO_INCREMENT=33 ;

-- --------------------------------------------------------

--
-- Table structure for table `conversation_reply`
--

CREATE TABLE `conversation_reply` (
  `cr_id` int(11) NOT NULL AUTO_INCREMENT,
  `reply` text CHARACTER SET utf8,
  `upload` varchar(30) DEFAULT NULL,
  `user_id_fk` int(11) NOT NULL,
  `ip` varchar(30) NOT NULL,
  `time` int(11) NOT NULL,
  `c_id_fk` int(11) NOT NULL,
  `read_status` int(11) NOT NULL DEFAULT '1',
  PRIMARY KEY (`cr_id`)
) ENGINE=MyISAM  DEFAULT CHARSET=latin1 AUTO_INCREMENT=807 ;

-- --------------------------------------------------------

--
-- Table structure for table `friends`
--

CREATE TABLE `friends` (
  `friend_id` int(11) NOT NULL AUTO_INCREMENT,
  `friend_one` int(11) DEFAULT NULL,
  `friend_two` int(11) DEFAULT NULL,
  `role` varchar(5) DEFAULT NULL,
  `block` enum('0','1') NOT NULL,
  `created` int(13) NOT NULL DEFAULT '1411570461',
  PRIMARY KEY (`friend_id`)
) ENGINE=MyISAM  DEFAULT CHARSET=latin1 AUTO_INCREMENT=281 ;

-- --------------------------------------------------------

--
-- Table structure for table `messages`
--

CREATE TABLE `messages` (
  `msg_id` int(11) NOT NULL AUTO_INCREMENT,
  `message` text CHARACTER SET utf8,
  `uid_fk` int(11) DEFAULT NULL,
  `ip` varchar(30) DEFAULT NULL,
  `created` int(11) NOT NULL DEFAULT '1269249260',
  `uploads` varchar(30) NOT NULL DEFAULT '',
  `like_count` int(11) NOT NULL DEFAULT '0',
  `comment_count` int(11) NOT NULL DEFAULT '0',
  `group_id_fk` int(11) NOT NULL DEFAULT '0',
  `secret_option` enum('0','1','2') DEFAULT '0',
  `share_count` int(11) DEFAULT '0',
  `video` varchar(30) DEFAULT NULL,
  `hashtag` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `p_mention` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`msg_id`)
) ENGINE=MyISAM  DEFAULT CHARSET=latin1 AUTO_INCREMENT=561 ;

-- --------------------------------------------------------

--
-- Table structure for table `message_like`
--

CREATE TABLE `message_like` (
  `like_id` int(11) NOT NULL AUTO_INCREMENT,
  `msg_id_fk` int(11) NOT NULL,
  `uid_fk` int(11) NOT NULL,
  `ouid_fk` int(11) DEFAULT NULL,
  `created` int(11) DEFAULT NULL,
  PRIMARY KEY (`like_id`)
) ENGINE=MyISAM  DEFAULT CHARSET=latin1 AUTO_INCREMENT=99 ;

-- --------------------------------------------------------

--
-- Table structure for table `message_share`
--

CREATE TABLE `message_share` (
  `share_id` int(11) NOT NULL AUTO_INCREMENT,
  `msg_id_fk` int(11) NOT NULL,
  `uid_fk` int(11) NOT NULL,
  `ouid_fk` int(11) DEFAULT NULL,
  `created` int(11) DEFAULT NULL,
  PRIMARY KEY (`share_id`)
) ENGINE=MyISAM  DEFAULT CHARSET=latin1 AUTO_INCREMENT=106 ;

-- --------------------------------------------------------

--
-- Table structure for table `notification_helper`
--

CREATE TABLE `notification_helper` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) DEFAULT NULL,
  `help_status` int(11) NOT NULL DEFAULT '1',
  `group_id_fk` int(11) NOT NULL DEFAULT '0',
  `created` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `remember_tokens`
--

CREATE TABLE `remember_tokens` (
  `uid` int(11) NOT NULL,
  `token` varchar(128) NOT NULL,
  `expires` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `reported_post`
--

CREATE TABLE `reported_post` (
  `r_id` int(11) NOT NULL AUTO_INCREMENT,
  `report_id` int(11) DEFAULT NULL,
  `report_type` varchar(1) DEFAULT NULL,
  `uid_fk` int(11) DEFAULT NULL,
  `ip` varchar(30) DEFAULT NULL,
  `created` int(11) NOT NULL DEFAULT '1269249260',
  PRIMARY KEY (`r_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `s_advertisements`
--

CREATE TABLE `s_advertisements` (
  `ads_id` int(11) NOT NULL AUTO_INCREMENT,
  `ads_title` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `ads_description` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `ads_img` varchar(255) DEFAULT NULL,
  `ads_url` varchar(900) DEFAULT NULL,
  `ads_durl` varchar(200) CHARACTER SET utf8 DEFAULT NULL,
  `status` int(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`ads_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=latin1 AUTO_INCREMENT=6 ;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `uid` int(20) NOT NULL AUTO_INCREMENT,
  `email` varchar(350) NOT NULL,
  `username` varchar(250) NOT NULL,
  `name` varchar(200) CHARACTER SET utf8 DEFAULT NULL,
  `password` varchar(350) NOT NULL,
  `forgot_code` text NOT NULL,
  `notification_created` int(11) DEFAULT NULL,
  `last_login` int(13) DEFAULT NULL,
  `join_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `photos_count` int(11) NOT NULL DEFAULT '0',
  `profile_pic` varchar(200) DEFAULT NULL,
  `mini_profile_pic` varchar(200) DEFAULT NULL,
  `mini_profile_pic_status` int(11) NOT NULL DEFAULT '1',
  `profile_bg_position` varchar(20) NOT NULL DEFAULT '0',
  `profile_pic_status` int(11) NOT NULL DEFAULT '1',
  `conversation_count` int(11) NOT NULL DEFAULT '0',
  `updates_count` int(11) NOT NULL DEFAULT '0',
  `friend_count` int(11) NOT NULL DEFAULT '0',
  `status` int(1) NOT NULL DEFAULT '1',
  `user_case` varchar(500) CHARACTER SET utf8 DEFAULT NULL,
  `user_city` varchar(80) DEFAULT NULL,
  `hide_following` enum('0','1','2') DEFAULT '0',
  `hide_followers` enum('0','1','2') DEFAULT '0',
  `hide_photos` enum('0','1','2') DEFAULT '0',
  `hide_maynow_people` enum('0','1') NOT NULL DEFAULT '0',
  `verified` enum('0','1') NOT NULL DEFAULT '0',
  `provider` varchar(10) DEFAULT NULL,
  `provider_id` int(30) DEFAULT NULL,
  `announcement_readed` enum('0','1') NOT NULL DEFAULT '0',
  `u_lang` enum('1','2','3','4') NOT NULL DEFAULT '1',
  `s_fb` varchar(80) DEFAULT NULL,
  `s_tw` varchar(80) DEFAULT NULL,
  `s_tumb` varchar(80) DEFAULT NULL,
  `s_yout` varchar(80) DEFAULT NULL,
  `s_inst` varchar(80) DEFAULT NULL,
  `s_gp` varchar(255) DEFAULT NULL,
  `gender` enum('0','1','2','3') NOT NULL DEFAULT '0',
  `b_day` varchar(2) DEFAULT NULL,
  `u_month` varchar(30) CHARACTER SET utf8 DEFAULT NULL,
  `u_year` varchar(5) DEFAULT NULL,
  PRIMARY KEY (`uid`)
) ENGINE=MyISAM  DEFAULT CHARSET=latin1 AUTO_INCREMENT=56 ;

-- --------------------------------------------------------

--
-- Table structure for table `user_profile_images`
--

CREATE TABLE `user_profile_images` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `profile_image_path` varchar(30) DEFAULT NULL,
  `uid_fk` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM  DEFAULT CHARSET=latin1 AUTO_INCREMENT=84 ;

-- --------------------------------------------------------

--
-- Table structure for table `user_uploads`
--

CREATE TABLE `user_uploads` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `image_path` varchar(30) DEFAULT NULL,
  `uid_fk` int(11) DEFAULT NULL,
  `group_id_fk` int(11) NOT NULL DEFAULT '0',
  `video_name` varchar(30) DEFAULT NULL,
  `video_ext` varchar(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM  DEFAULT CHARSET=latin1 AUTO_INCREMENT=150 ;
