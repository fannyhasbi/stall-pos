CREATE TABLE employee (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(50) NOT NULL,
  username VARCHAR(50) NOT NULL,
  password VARCHAR(255) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE customer (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(50),
  type INT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE product (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(100) NOT NULL,
  description VARCHAR(500) NULL,
  price FLOAT NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE orders (
  id INT NOT NULL AUTO_INCREMENT,
  status SMALLINT NOT NULL DEFAULT 1,
  total_price FLOAT NOT NULL,
  ordered_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  done_at TIMESTAMP NULL,
  cust_type SMALLINT NOT NULL DEFAULT 1,
  emp_id INT NULL,
  cust_id INT NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (emp_id) REFERENCES employee(id),
  FOREIGN KEY (cust_id) REFERENCES customer(id)
);

CREATE  TABLE order_details (
  id INT NOT NULL AUTO_INCREMENT,
  product_id INT NOT NULL,
  order_id INT NOT NULL,
  FOREIGN KEY (product_id) REFERENCES product(id),
  FOREIGN KEY (order_id) REFERENCES orders(id),
  PRIMARY KEY (id)
);