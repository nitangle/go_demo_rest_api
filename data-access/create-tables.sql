use book_manager;
DROP TABLE IF EXISTS users,books;
create table users(
    id INT AUTO_INCREMENT NOT NULL,
    name  VARCHAR(200) NOT NULL ,
    PRIMARY KEY(`id`)
);

CREATE TABLE books (
  id         INT AUTO_INCREMENT NOT NULL,
  title      VARCHAR(128) NOT NULL,
  author     VARCHAR(255) NOT NULL,
  category     VARCHAR(255) NOT NULL,
  lentTo      INT,
  lentOn      TIMESTAMP, 
  FOREIGN KEY(lentTo) REFERENCES users(id),
  PRIMARY KEY (`id`)
);

INSERT INTO books 
  (title, author, category) 
VALUES 
  ('Blue Train', 'John Coltrane', 'cs'),
  ('Giant Steps', 'John Coltrane', 'fiction'),
  ('Jeru', 'Gerry Mulligan', 'technical'),
  ('Sarah Vaughan', 'Sarah Vaughan', 'drama');