CREATE TYPE gender_type AS ENUM ('MALE', 'FEMALE');
CREATE TYPE program_type AS ENUM ('BTECH', 'IDD', 'MTECH', 'PHD');
CREATE TYPE enrollment_status_type AS ENUM ('ACTIVE', 'INACTIVE');

CREATE TABLE students (
    id                  SERIAL PRIMARY KEY,
    student_id          VARCHAR(20) UNIQUE NOT NULL,
    full_name           VARCHAR(100) NOT NULL,
    date_of_birth       DATE NOT NULL,
    gender              gender_type NOT NULL,

    institute_email     VARCHAR(100) UNIQUE NOT NULL,
    phone_number        VARCHAR(15) NOT NULL,

    password_hash       VARCHAR(255) NOT NULL,

    program             program_type NOT NULL,
    branch              VARCHAR(100) NOT NULL,

    current_year        INTEGER NOT NULL CHECK (current_year BETWEEN 1 AND 5),
    current_semester    INTEGER NOT NULL CHECK (current_semester BETWEEN 1 AND 10),

    is_verified         BOOLEAN DEFAULT FALSE,
    verified_at         TIMESTAMP,
    verified_by         INTEGER REFERENCES admins(id) ON DELETE SET NULL,

    enrollment_status   enrollment_status_type DEFAULT 'ACTIVE',

    created_at          TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);