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

-- ministries table
CREATE TABLE IF NOT EXISTS church.ministries (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL UNIQUE, -- 'Choir', 'Media', 'Ushers'
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- the join table (member <-> ministry)
CREATE TABLE IF NOT EXISTS church.member_ministries (
    member_id UUID REFERENCES church.members(id) ON DELETE CASCADE,
    ministry_id UUID REFERENCES church.ministries(id) ON DELETE CASCADE,
    joined_at TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY (member_id, ministry_id)
);

--- sermons table
CREATE TABLE IF NOT EXISTS church.sermons (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    speaker VARCHAR(100) NOT NULL,
    series VARCHAR(100),
    video_url TEXT, -- YouTube link
    audio_url TEXT, -- audio link
    summary TEXT,
    tags TEXT[],
    preached_at DATE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- initial seeds
INSERT INTO church.member_status (name) 
VALUES ('Member'), ('Visitor'), ('Pastor'), ('Admin')
ON CONFLICT (name) DO NOTHING;

INSERT INTO church.ministries (name, description) 
VALUES ('Choir', 'Worship team'), ('Media', 'Tech team');

select * from church.member_status;
select * from church.members;