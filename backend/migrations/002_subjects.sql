CREATE TABLE IF NOT EXISTS subjects (
    id INTEGER PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

INSERT INTO subjects
    (id, name)
VALUES 
    (1, 'Maths'),
    (2, 'Anglais'),
    (3, 'Français'),
    (4, 'Histoire'),
    (5, 'Géographie'),
    (6, 'Sport'),
    (7, 'Science');
