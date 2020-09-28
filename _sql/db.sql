SET NAMES utf8;
SET time_zone = '+03:00';
DROP TABLE IF EXISTS links;

CREATE TABLE `links` (
    `id` VARCHAR(255) NOT NULL,
    `link` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8;

