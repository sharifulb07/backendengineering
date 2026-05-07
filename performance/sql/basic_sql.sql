


-- SELECT product_name
-- FROM products
-- WHERE EXISTS (
--     SELECT 1
--     FROM users
--     WHERE users.id = products.user_id
-- );

-- Select all rows from 'cars'
-- SELECT  year FROM cars GROUP BY year;

-- Select all rows from 'cars UNION SELECT * FROM products'
-- SELECT * FROM person UNION SELECT * FROM products;



-- SELECT users.name, orders.product_name, amount FROM users INNER JOIN orders ON users.id=orders.user_id;
-- SELECT users.name, orders.product_name, amount FROM users LEFT JOIN orders ON users.id=orders.user_id;
-- SELECT users.name, orders.product_name, amount FROM users RIGHT JOIN orders ON users.id=orders.user_id;
-- CROSS JOIN HAS NO CONDITION 
-- SELECT users.name, orders.product_name, amount FROM users CROSS JOIN orders;



-- INSERT INTO orders(user_id, product_name, amount)
-- VALUES
-- (1, 'Laptop', 1200.00),
-- (1, 'Mouse', 20.00),
-- (2, 'Phone', 900.00),
-- (3, 'Keyboard', 50.00);


-- INSERT INTO users (name)
-- VALUES
-- ('Shariful'),
-- ('Rahim'),
-- ('Karim'),
-- ('Sakib');

-- CREATE TABLE users(
--     id SERIAL PRIMARY KEY, 
--     name VARCHAR(100) NOT NULL
-- );

-- CREATE TABLE orders(
--     id SERIAL PRIMARY KEY,
--     user_id INT, 
--     product_name VARCHAR(100),
--     amount NUMERIC (10,2),
--     FOREIGN KEY (user_id) REFERENCES users(id)
-- );




-- SELECT * FROM products;
-- Select all rows from 'person'
-- SELECT * FROM person;


-- -- Insert data into 'products'
-- INSERT INTO products (name, category, price, stock, brand)
-- VALUES 
--  ('iPhone 15', 'Mobile', 1200.00, 10, 'Apple'),
--     ('Galaxy S23', 'Mobile', 950.00, 15, 'Samsung'),
--     ('MacBook Air M2', 'Laptop', 1400.00, 5, 'Apple'),
--     ('Dell XPS 13', 'Laptop', 1300.00, 8, 'Dell'),
--     ('Sony WH-1000XM5', 'Headphones', 350.00, 20, 'Sony'),
--     ('Logitech MX Master 3', 'Mouse', 99.99, 50, 'Logitech'),
--     ('Apple Watch Series 9', 'Wearable', 499.00, 12, 'Apple');


-- Select all rows from 'cars'


-- Select all rows from 'cars'
-- SELECT * FROM cars;

-- SELECT * FROM cars WHERE brand IN ('BMW', 'Toyota', 'Honda', 'MERCIDIS');
-- SELECT * FROM cars WHERE brand NOT IN ('BMW', 'Toyota', 'Honda', 'MERCIDIS');

-- Select all rows from 'cars'
-- SELECT * FROM cars WHERE model LIKE '%o%';
-- SELECT * FROM cars WHERE model LIKE 'C%c';



-- SELECT SUM(price) FROM cars;
-- SELECT ROUND(AVG(price), 0) FROM cars;

-- SELECT MAX(price) AS HIGHTEST_PRICE FROM cars;

-- Select all rows from 'cars LIMIT 2'
-- SELECT * FROM cars LIMIT 1 OFFSET 3;

-- Select all rows from 'cars ORDER BY year'
-- SELECT * FROM cars WHERE year IS NOT NULL ORDER BY year ASC;


-- Insert data into 'brand'
-- INSERT INTO cars (brand, model, year, price)
-- VALUES ('Tesla', 'Motion-Ds', 2024, 5022145);

-- Select all rows from 'cars'
-- SELECT * FROM cars;


-- SELECT brand, model FROM cars WHERE year>2019;
-- SELECT brand, model FROM cars WHERE year=2019;




-- Select all rows from 'cars'



-- SELECT * FROM cars;




-- Select all rows from 'cars'
-- SELECT COUNT( technology) FROM teacher;