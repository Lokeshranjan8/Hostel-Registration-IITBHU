CREATE TYPE session_status_type AS ENUM ('UPCOMING', 'ACTIVE', 'CLOSED');
CREATE TYPE session_term_type AS ENUM ('ODD', 'EVEN');

CREATE TABLE hostel_sessions (
    id              SERIAL PRIMARY KEY,
    session_name    VARCHAR(50) UNIQUE NOT NULL,  -- "2024 ODD", "2024 EVEN"
    term            session_term_type NOT NULL,    -- ODD or EVEN
    academic_year   VARCHAR(10) NOT NULL,
    start_date      DATE NOT NULL,
    end_date        DATE NOT NULL,
    status          session_status_type DEFAULT 'UPCOMING',
    created_by      INTEGER REFERENCES admins(id) ON DELETE SET NULL,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);