-- Create users table
CREATE TABLE user_cars (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    registration_number TEXT NOT NULL,
    model TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
