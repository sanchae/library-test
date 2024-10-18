CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL
);
INSERT INTO users (first_name, last_name) VALUES
('John', 'Doe'),
('Jane', 'Smith'),
('Michael', 'Johnson'),
('Emily', 'Brown'),
('David', 'Wilson'),
('Sarah', 'Taylor'),
('Christopher', 'Anderson'),
('Jessica', 'Martinez'),
('Matthew', 'Thomas'),
('Lauren', 'Garcia');