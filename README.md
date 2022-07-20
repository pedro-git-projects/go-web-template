# Go Web Server Template

This repository contains a simple architectural template for Go web applications.

## Project Structure

- **bin**: Compiled application binaries for deployment to the production server

- **cmd/api**: Application-specific code. This includes code for running the server, reading and writing HTTP requests and managing authentication 

- **internal**: Ancillary packages used by the API, as code for database interaction, data validation, sending emails and so on. 

- **migarations**: SQL migration files for the database 

- **remote**: Configuration files and setup scripts for the production server 

- **go.mod**: Project dependencies, versions and module path 

- **Makefile**: Recipes for automating administrative tasks 

