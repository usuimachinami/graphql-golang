-- phpMyAdmin SQL Dump
-- version 4.7.1
-- https://www.phpmyadmin.net/
--
-- Host: mysql
-- Generation Time: Jun 28, 2018 at 01:53 PM
-- Server version: 5.7.18
-- PHP Version: 7.0.16

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `video`
--

-- --------------------------------------------------------

--
-- Table structure for table `favorites`
--

CREATE TABLE `favorites` (
  `id` int(11) NOT NULL,
  `full_title_id` varchar(6) NOT NULL,
  `user_id` int(11) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `favorites`
--

INSERT INTO `favorites` (`id`, `full_title_id`, `user_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, '000001', 12, '2017-12-21 01:25:20', '2017-12-21 01:25:20', NULL),
(2, '000001', 11, '2017-12-21 20:12:03', '2017-12-21 20:12:03', NULL),
(3, '000001', 1, '2017-12-21 20:56:24', '2017-12-21 20:56:24', NULL),
(4, '000002', 1, '2017-12-21 21:07:11', '2017-12-21 21:07:11', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `stories`
--

CREATE TABLE `stories` (
  `full_title_id` varchar(6) NOT NULL,
  `story_id` varchar(3) NOT NULL,
  `full_story_id` varchar(9) NOT NULL,
  `story_name` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `stories`
--

INSERT INTO `stories` (`full_title_id`, `story_id`, `full_story_id`, `story_name`, `created_at`, `updated_at`, `deleted_at`) VALUES
('000001', '001', '000001001', '第一話', '2017-12-15 18:24:27', '2017-12-15 18:24:27', NULL),
('000001', '002', '000001002', '第二話', '2017-12-15 18:24:27', '2017-12-15 18:24:27', NULL),
('000002', '001', '000002001', '第一話', '2017-12-15 18:24:27', '2017-12-15 18:24:27', NULL),
('000002', '002', '000002002', '第二話', '2017-12-15 18:24:27', '2017-12-15 18:24:27', NULL),
('001001', '001', '001001001', '第一話', '2017-12-15 18:24:27', '2017-12-15 18:24:27', NULL),
('001001', '002', '001001002', '第二話', '2017-12-15 18:24:27', '2017-12-15 18:24:27', NULL),
('001002', '001', '001002001', '第一話', '2017-12-15 18:24:27', '2017-12-15 18:24:27', NULL),
('001002', '002', '001002002', '第二話', '2017-12-15 18:24:27', '2017-12-15 18:24:27', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `titles`
--

CREATE TABLE `titles` (
  `publisher_id` varchar(3) NOT NULL,
  `title_id` varchar(3) NOT NULL,
  `full_title_id` varchar(6) NOT NULL,
  `title_name` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `titles`
--

INSERT INTO `titles` (`publisher_id`, `title_id`, `full_title_id`, `title_name`, `created_at`, `updated_at`, `deleted_at`) VALUES
('000', '001', '000001', '番組①第一シーズン', '2017-12-15 16:36:30', '2018-06-28 13:52:40', NULL),
('000', '002', '000002', '番組①第二シーズン', '2017-12-15 16:36:46', '2018-06-28 13:53:10', NULL),
('001', '001', '001001', '番組②第一シーズン', '2017-12-15 16:36:30', '2018-06-28 13:53:00', NULL),
('001', '002', '001002', '番組②第二シーズン', '2017-12-15 16:36:46', '2018-06-28 13:53:06', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `password`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'testuser', 'testpassword', '2017-12-16 00:00:00', '2018-06-28 13:45:39', NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `favorites`
--
ALTER TABLE `favorites`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `stories`
--
ALTER TABLE `stories`
  ADD PRIMARY KEY (`full_story_id`),
  ADD UNIQUE KEY `full_title_id` (`full_story_id`);

--
-- Indexes for table `titles`
--
ALTER TABLE `titles`
  ADD PRIMARY KEY (`full_title_id`),
  ADD UNIQUE KEY `full_title_id` (`full_title_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `favorites`
--
ALTER TABLE `favorites`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
