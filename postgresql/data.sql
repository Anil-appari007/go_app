-- CREATE DATABASE inventory;
-- \l

-- \c inventory

CREATE TABLE inventory (
  id SERIAL UNIQUE PRIMARY KEY, 
  name VARCHAR (50) UNIQUE NOT NULL, 
  price NUMERIC(5, 2),
  sales INT,
  stock INT
);

-- \dt

INSERT INTO inventory (name, price, sales, stock)
VALUES ('apple', 11.11, 22, 44);

INSERT INTO inventory (name, price, sales, stock)
VALUES 
('banana', 22.5, 53, 65),
('mango', 65.9, 44, 44),
('orange', 3.54, 12, 98);

INSERT INTO inventory (name, price, sales, stock)
VALUES ('papayya', 33, 45, 90);

-- select * from inventory;

-- select * from inventory where name = 'apple';

-- update inventory
-- SET price = 12, sales = 20, stock = 49
-- WHERE name = 'apple';


-- DELETE FROM inventory
-- WHERE name = 'apple';

