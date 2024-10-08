CREATE TABLE IF NOT EXISTS entries (
`id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
userId INT UNSIGNED NOT NULL,
`title` VARCHAR(255) NOT NULL,
`content` TEXT NOT NULL,
`topic` VARCHAR(255),
`createdAt` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

PRIMARY KEY (`id`),
FOREIGN KEY (`userId`) REFERENCES users(`id`)
)
