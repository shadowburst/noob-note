CREATE TABLE IF NOT EXISTS test_results (
    student_id INTEGER NOT NULL,
    test_id INTEGER NOT NULL,
    blue_count INTEGER DEFAULT 0,
    green_count INTEGER DEFAULT 0,
    yellow_count INTEGER DEFAULT 0,
    red_count INTEGER DEFAULT 0,
    FOREIGN KEY (student_id) REFERENCES students (id) ON DELETE CASCADE,
    FOREIGN KEY (test_id) REFERENCES tests (id) ON DELETE CASCADE,
    PRIMARY KEY(student_id, test_id)
);
