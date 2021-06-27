DROP DATABASE IF EXISTS budgieDB;
CREATE DATABASE budgieDB;
USE budgieDB;

CREATE TABLE tblUsers (
	fldID 		 BIGINT AUTO_INCREMENT PRIMARY KEY,
	fldCreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
    fldUpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP(),
	fldFirstName VARCHAR(50) NOT NULL,
	fldLastName  VARCHAR(50) NOT NULL,
	fldEmail     VARCHAR(50) NOT NULL UNIQUE,
	fldPassword  VARCHAR(255) NOT NULL
);

CREATE TABLE tblExpenseCat (
	fldID BIGINT AUTO_INCREMENT PRIMARY KEY,
    fldCategoryName VARCHAR(25) NOT NULL
);

CREATE TABLE tblExpense (
	fldID BIGINT AUTO_INCREMENT PRIMARY KEY,
    fldCreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
    fldUpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP(),
    fldTitle VARCHAR(50),
    fldDate DATE NOT NULL,
	fldAmount DOUBLE NOT NULL,
    fldNote TEXT,
    fldFKCategoryID BIGINT NOT NULL,
	fldFKUserID BIGINT NOT NULL,
	FOREIGN KEY(fldFKUserID) REFERENCES tblUsers(fldID) ON DELETE CASCADE,
    FOREIGN KEY(fldFKCategoryID) REFERENCES tblExpenseCat(fldID)
);

/* Insert expense categories */
INSERT INTO tblExpenseCat(fldCategoryName) VALUES ("Uncategorized");
INSERT INTO tblExpenseCat(fldCategoryName) VALUES ("Food");
INSERT INTO tblExpenseCat(fldCategoryName) VALUES ("Utility");
INSERT INTO tblExpenseCat(fldCategoryName) VALUES ("Rent/Mortgage");
INSERT INTO tblExpenseCat(fldCategoryName) VALUES ("Hobbies");
INSERT INTO tblExpenseCat(fldCategoryName) VALUES ("Other");