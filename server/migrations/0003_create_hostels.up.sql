-- CREATE TYPE gender_type AS ENUM ('MALE','FEMALE');

CREATE TABLE hostels (
    id              SERIAL PRIMARY KEY,
    hostel_name     VARCHAR(100) UNIQUE NOT NULL,     
    gender_type     gender_type NOT NULL,         
    total_capacity  INTEGER NOT NULL,             
    warden_id       INTEGER REFERENCES admins(id) ON DELETE SET NULL,
    is_active       BOOLEAN DEFAULT TRUE,
    created_by      INTEGER REFERENCES admins(id) ON DELETE SET NULL,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);