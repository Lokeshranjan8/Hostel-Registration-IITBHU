CREATE TABLE Students (
    id SERIAL PRIMARY KEY,
    student_id VARCHAR(20) UNIQUE NOT NULL,
    full_name VARCHAR(100) NOT NULL,
    institute_email VARCHAR(100)  UNIQUE NOT NULL,
    phone_number VARCHAR(15), 


    gender  VARCHAR(10) NOT NULL,
    branch VARCHAR(100) NOT NULL,
    program VARCHAR(20) NOT NULL,
    current_year  INTEGER NOT NULL,
    current_semester INTEGER  NOT NULL,

    is_verified BOOLEAN DEFAULT FALSE,
    -- use for no-dues status

    data_of_birth DATE NOT NULL,
    enrollement_status VARCHAR(20) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP

);
