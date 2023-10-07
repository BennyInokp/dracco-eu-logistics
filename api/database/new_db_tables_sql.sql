-- Create the schema "dbo" if it doesn't already exist
CREATE SCHEMA IF NOT EXISTS dbo;

-- Create the "users" table
CREATE TABLE dbo.users (
  user_id INT PRIMARY KEY,
  name VARCHAR(255),
  email VARCHAR(255),
  phone VARCHAR(20),
  address VARCHAR(255),
  type VARCHAR(5),
  vat_number VARCHAR(30),
  is_verified BOOLEAN,
  modification_date DATE,
  password VARCHAR(255)
);

-- Create the "distributors" table
CREATE TABLE dbo.distributors (
  distributor_id INT PRIMARY KEY,
  user_id INT,
  FOREIGN KEY (user_id) REFERENCES dbo.users(user_id)
);

-- Create the "packages" table
CREATE TABLE dbo.packages (
  package_id INT PRIMARY KEY,
  distributor_id INT,
  origin VARCHAR(255),
  destination VARCHAR(255),
  type VARCHAR(100),
  est_value FLOAT,
  offer_price FLOAT,
  publication_date DATE,
  modification_date DATE,
  deadline_date DATE,
  dimensiones VARCHAR(50),
  eu_palet BOOLEAN,
  accepted BOOLEAN,
  completed BOOLEAN,
  FOREIGN KEY (distributor_id) REFERENCES dbo.distributors(distributor_id)
);

-- Create the "transporters" table
CREATE TABLE dbo.transporters (
  transporter_id INT PRIMARY KEY,
  user_id INT,
  FOREIGN KEY (user_id) REFERENCES dbo.users(user_id)
);

-- Create the "vehicles" table
CREATE TABLE dbo.vehicles (
  vehicle_id INT PRIMARY KEY,
  transporter_id INT,
  plate VARCHAR(15),
  brand VARCHAR(50),
  model VARCHAR(50),
  capacity_kg DECIMAL(10,2),
  FOREIGN KEY (transporter_id) REFERENCES dbo.transporters(transporter_id)
);

-- Create the "orders" table
CREATE TABLE dbo.orders (
  order_id BIGINT PRIMARY KEY,
  package_id INT,
  transporter_id INT,
  distributor_id INT,
  order_price FLOAT,
  modification_date TIMESTAMP,
  est_arrival_date TIMESTAMP,
  order_date TIMESTAMP,
  order_status VARCHAR(50),
  FOREIGN KEY (package_id) REFERENCES dbo.packages(package_id),
  FOREIGN KEY (transporter_id) REFERENCES dbo.transporters(transporter_id),
  FOREIGN KEY (distributor_id) REFERENCES dbo.distributors(distributor_id)
);

-- Create the "tracking" table
CREATE TABLE dbo.tracking (
  tracking_id INT PRIMARY KEY,
  order_id BIGINT,
  latitude DECIMAL(10,6),
  longitude DECIMAL(10,6),
  modification_date TIMESTAMP,
  FOREIGN KEY (order_id) REFERENCES dbo.orders(order_id)
);

-- Create the "documents" table
CREATE TABLE dbo.documents (
  document_id INT PRIMARY KEY,
  user_id INT,
  order_id BIGINT,
  type VARCHAR(255),
  description VARCHAR(255),
  name VARCHAR(255),
  FOREIGN KEY (user_id) REFERENCES dbo.users(user_id),
  FOREIGN KEY (order_id) REFERENCES dbo.orders(order_id)
);

-- Create the "audit_log" table
CREATE TABLE dbo.audit_log (
  log_id INT PRIMARY KEY AUTO_INCREMENT,
  table_name VARCHAR(255),
  action VARCHAR(10),
  change_date TIMESTAMP,
  user_id INT,
  executed_sql TEXT
);
