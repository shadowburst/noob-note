CREATE TABLE IF NOT EXISTS students (
    id INTEGER PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    classroom_id INTEGER NOT NULL,
    FOREIGN KEY (classroom_id) REFERENCES classrooms (id) ON DELETE CASCADE
);
