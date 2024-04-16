# Essential data to run the demo.

//////////////////////////////////////////////////////////

# PACKAGES AND COMMANDS

First make sure you have the gqlgen package installed, to install it I ran the following command:
# NOTE: This command only needs to be executed once.
# go get -d github.com/99designs/gqlgen

Every time the 'schema.graphql' file is updated, the command must be executed:
# go run github.com/99designs/gqlgen generate

When you want to run the project, use the command
# go run server.go

////////////////////////////////////////////////////////////

# NOTE: To run the demo, make sure you have the essential packages installed and the database configured.

To configure the database, you can do it in the following demo path:

# config\config.go

////////////////////////////////////////////////////////////

# DATABASE

We will use a local database, for this database you will have to use the following Script

CREATE DATABASE inventory;

USE inventory;

CREATE TABLE refrigerators (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    brand VARCHAR(255) NOT NULL,
    color VARCHAR(255) NOT NULL,
    category VARCHAR(255) NOT NULL
);

CREATE TABLE televisions (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    brand VARCHAR(255) NOT NULL,
    color VARCHAR(255) NOT NULL,
    category VARCHAR(255) NOT NULL
);

CREATE TABLE cellphones (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    brand VARCHAR(255) NOT NULL,
    color VARCHAR(255) NOT NULL,
    category VARCHAR(255) NOT NULL
);
