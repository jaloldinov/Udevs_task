#### guidline

![Test Image](/project.png)

## Open every service seperately otherwise you might  get error :)
<br />

#### FIRST OF ALL YOU NEED TO SET UP AN ENV ACCORDING TO .ENV_EXAMPLE FILE (but I left .gitignore :) )
<br />

## service folder : 
- create database

- then run with 

       make migration-up:

       make run

## api_gateway folder : 

        make run

### After running all service you can see the result in the following link:

 http://localhost:8080/swagger/index.html#/ 


<br />
<br />
<br />

# Task: Create catalog of books.

1) For creating, receiving, modifying and deleting books must have API
2) The API must have swagger documentation which you can test methods.
3) When creating an API, you can use any framework (gin, fiber, bego, echo).
4) Books should be stored in postgresql.
5) There are must be at least two tables in database, for example (books, book_category) and they must be interconnected (reference).
6) Separate API methods from postgres methods to create a microservice and use grpc.