CREATE TABLE IF NOT EXISTS cycles (
    id INTEGER PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

INSERT INTO cycles
    (id, name)
VALUES
    (1, 'Cycle 1'),
    (2, 'Cycle 2'),
    (3, 'Cycle 3');
