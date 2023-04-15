-- Create custom ENUM type for user roles
CREATE TYPE user_role AS ENUM ('librarian', 'customer');

-- Create authors table
CREATE TABLE authors (
                         id SERIAL PRIMARY KEY,
                         name VARCHAR(255) NOT NULL,
                         biography TEXT,
                         birth_date DATE
);

-- Create categories table
CREATE TABLE categories (
                            id SERIAL PRIMARY KEY,
                            name VARCHAR(255) NOT NULL,
                            description TEXT
);

-- Create books table
CREATE TABLE books (
                       id SERIAL PRIMARY KEY,
                       title VARCHAR(255) NOT NULL,
                       author_id INTEGER REFERENCES authors(id),
                       category_id INTEGER REFERENCES categories(id),
                       publication_date DATE NOT NULL,
                       isbn VARCHAR(13) NOT NULL UNIQUE,
                       summary TEXT,
                       cover_image_url VARCHAR(255)
);

-- Create users table
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       first_name VARCHAR(255) NOT NULL,
                       last_name VARCHAR(255) NOT NULL,
                       email VARCHAR(255) NOT NULL UNIQUE,
                       password VARCHAR(255) NOT NULL,
                       role user_role NOT NULL
);

-- Create book_loans table
CREATE TABLE book_loans (
                            id SERIAL PRIMARY KEY,
                            book_id INTEGER REFERENCES books(id),
                            user_id INTEGER REFERENCES users(id),
                            due_date DATE NOT NULL,
                            returned_date DATE
);
