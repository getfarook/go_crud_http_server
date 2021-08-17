About this Project:
-----------------------------
This is a simple crud project to demonstrate the http-server working in go.
Postgres in the local system is used in the backend
REST APIs were tested using postman application.
APIs are saved in postman under collection CRUD-HTTP-SERVER-GO





Postgres Details:
-----------------------------
postgres installed in my macbook locally was used

Database Name: go_httpserver_crud
User ID: postgres
Password: admin123
Table Name: partners

Sql Queries Used to Create Database and the Tables:

CREATE DATABASE go_httpserver_crud;

CREATE TABLE partners (
    id SERIAL PRIMARY KEY, 
    name TEXT NOT NULL, 
    age INT, 
    dob DATE, 
    balance NUMERIC, 
    access BOOLEAN);



INSERT INTO partners ( name, age, dob, balance, access)
VALUES ('Mohammed Farook', 38, '1982-07-27', 12500.50, true);

INSERT INTO partners ( name, age, dob, balance, access)
VALUES ('Thomas George', 35, '1990-07-29', 12500.50, 'true');
