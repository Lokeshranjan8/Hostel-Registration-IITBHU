CREATE TYPE no_dues_type AS ENUM ( 'PENDING','CLEARED','NOT_CLEARED');

CREATE TABLE no_dues (
    id              SERIAL PRIMARY KEY,
    student_id      INTEGER REFERENCES students(id) ON DELETE CASCADE,
    session_id      INTEGER REFERENCES hostel_sessions(id) ON DELETE CASCADE,
    status          no_dues_type DEFAULT 'PENDING',
    reason          TEXT,
    verified_by     INTEGER REFERENCES admins(id) ON DELETE SET NULL,
    verified_at     TIMESTAMP,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)