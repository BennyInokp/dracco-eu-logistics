CREATE DATABASE draco_logistics_eu;

USE draco_logistics_eu;

CREATE TABLE customers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    customer_id INT NOT NULL,
    customer_name VARCHAR(255) NOT NULL,
    customer_email VARCHAR(255) NOT NULL,
    customer_phone VARCHAR(255) NOT NULL,
    tax_id VARCHAR(255) NOT NULL,
    verified BOOLEAN NOT NULL,
    password VARCHAR(255) NOT NULL,
    is_transporter BOOLEAN NOT NULL,
    is_warehouse BOOLEAN NOT NULL
);



CREATE TABLE transporters (
  id INT PRIMARY KEY,
  transporter_name VARCHAR(50),
  transporter_email VARCHAR(100),
  transporter_phone VARCHAR(20),
  vehicle_brand VARCHAR(50),
  vehicle_model VARCHAR(50),
  vehicle_license_plate VARCHAR(20),
  vehicle_max_capacity INT,
  is_verified BOOLEAN,
  tax_id VARCHAR(20)
);

CREATE TABLE documents (
  id INT PRIMARY KEY,
  transporter_id INT,
  document_type VARCHAR(50),
  document_description VARCHAR(100),
  document_url VARCHAR(200),
  FOREIGN KEY (transporter_id) REFERENCES transporters(id)
);

CREATE TABLE parcels (
  id VARCHAR(50) PRIMARY KEY,
  company_name VARCHAR(100),
  company_id INT,
  company_phone VARCHAR(20),
  tax_id VARCHAR(20),
  description VARCHAR(200),
  weight FLOAT,
  pickup_location_address VARCHAR(200),
  pickup_location_city VARCHAR(100),
  pickup_location_zip_code VARCHAR(20),
  pickup_location_longitude VARCHAR(20),
  pickup_location_latitude VARCHAR(20),
  pick_up_date DATETIME,
  drop_off_date DATETIME,
  drop_off_destination_address VARCHAR(200),
  drop_off_destination_city VARCHAR(100),
  drop_off_destination_zip_code VARCHAR(20),
  drop_off_destination_longitude VARCHAR(20),
  drop_off_destination_latitude VARCHAR(20),
  is_eu_palet BOOLEAN,
  offer_price FLOAT,
  country_name VARCHAR(100),
  country_code VARCHAR(10),
  transporter_id INT,
  is_accepted BOOLEAN,
  is_completed BOOLEAN,
  FOREIGN KEY (transporter_id) REFERENCES transporters(id)
);

CREATE TABLE locations (
  id INT PRIMARY KEY,
  address VARCHAR(200),
  city VARCHAR(100),
  zip_code VARCHAR(20),
  longitude VARCHAR(20),
  latitude VARCHAR(20)
);

CREATE TABLE countries (
  code VARCHAR(10) PRIMARY KEY,
  name VARCHAR(100)
);

CREATE TABLE dimensions (
  id INT PRIMARY KEY,
  length FLOAT,
  width FLOAT,
  height FLOAT
);
