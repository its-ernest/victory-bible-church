-- Schema setup
CREATE SCHEMA IF NOT EXISTS church;
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- status lookup table
CREATE TABLE IF NOT EXISTS church.member_status (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL
);

-- members table
CREATE TABLE IF NOT EXISTS church.members (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    phone VARCHAR(20) UNIQUE NOT NULL,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    email VARCHAR(255) UNIQUE,
    status_id INTEGER REFERENCES church.member_status(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- initial seeds
INSERT INTO church.member_status (name) 
VALUES ('Member'), ('Visitor'), ('Pastor'), ('Admin')
ON CONFLICT (name) DO NOTHING;

select * from church.member_status;
select * from church.members;