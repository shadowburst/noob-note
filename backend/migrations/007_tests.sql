CREATE TABLE IF NOT EXISTS tests (
    id INTEGER PRIMARY KEY,
    date VARCHAR(255) NOT NULL,
    term_id INTEGER NOT NULL,
    classroom_id INTEGER NOT NULL,
    subject_id INTEGER NOT NULL,
    FOREIGN KEY (term_id) REFERENCES terms (id) ON DELETE CASCADE,
    FOREIGN KEY (classroom_id) REFERENCES classrooms (id) ON DELETE CASCADE,
    FOREIGN KEY (subject_id) REFERENCES subjects (id) ON DELETE CASCADE
);

