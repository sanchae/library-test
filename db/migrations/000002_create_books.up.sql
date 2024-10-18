CREATE TABLE IF NOT EXISTS books (
    book_id SERIAL PRIMARY KEY,
    book_title VARCHAR(255) NOT NULL,
    available_copies INT NOT NULL
);
INSERT INTO books (book_title, available_copies) VALUES
('To Kill a Mockingbird', 5),
('1984', 3),
('Pride and Prejudice', 4),
('The Great Gatsby', 2),
('One Hundred Years of Solitude', 3),
('Brave New World', 4),
('The Catcher in the Rye', 3),
('Lord of the Flies', 5),
('The Hobbit', 2),
('Fahrenheit 451', 4);