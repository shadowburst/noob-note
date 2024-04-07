CREATE TABLE IF NOT EXISTS grades (
    id INTEGER PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    cycle_id INTEGER NOT NULL,
    FOREIGN KEY (cycle_id) REFERENCES cycles (id) ON DELETE CASCADE
);

INSERT INTO grades
    (id, name, cycle_id)
VALUES 
    (1, 'TPS', 1),
    (2, 'PS', 1),
    (3, 'MS', 1),
    (4, 'GS', 1),
    (5, 'CP', 2),
    (6, 'CE1', 2),
    (7, 'CE2', 2),
    (8, 'CM1', 3),
    (9, 'CM2', 3);
