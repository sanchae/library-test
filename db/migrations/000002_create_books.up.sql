CREATE TABLE IF NOT EXISTS books (
    book_id SERIAL PRIMARY KEY,
    book_title VARCHAR(255) NOT NULL,
    available_copies INT NOT NULL
);