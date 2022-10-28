CREATE TABLE `Employees` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(255) NOT NULL,
	`status` VARCHAR(255) NOT NULL,
	`password` VARCHAR(255) NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE `Transactions` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`idSender` INT NOT NULL,
	`idReceiver` INT NOT NULL,
	`value` INT NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE `Items` (
	`name` VARCHAR(255) NOT NULL,
	`price` INT NOT NULL,
	`ownerId` INT NOT NULL
);

CREATE TABLE `Sessions` (
	`userId` INT NOT NULL,
	`status` VARCHAR(255) NOT NULL,
	`refreshToken` VARCHAR(255) NOT NULL,
	`timeClose` INT NOT NULL
);

ALTER TABLE `Transactions` ADD CONSTRAINT `Transactions_fk0` FOREIGN KEY (`idSender`) REFERENCES `Employees`(`id`);

ALTER TABLE `Transactions` ADD CONSTRAINT `Transactions_fk1` FOREIGN KEY (`idReceiver`) REFERENCES `Employees`(`id`);

ALTER TABLE `Items` ADD CONSTRAINT `Items_fk0` FOREIGN KEY (`ownerId`) REFERENCES `Employees`(`id`);

INSERT INTO Employees VALUES(1, "Bob", "admin", "qwerty123");
INSERT INTO Employees VALUES(2, "Jack", "boss", "bobqewsax");
INSERT INTO Employees VALUES(3, "Nick", "employee", "asdcxvsa");

INSERT INTO Transactions VALUES(1, 1, 2, 100);

INSERT INTO Items VALUES("Account CS:GO", 50, 1);







