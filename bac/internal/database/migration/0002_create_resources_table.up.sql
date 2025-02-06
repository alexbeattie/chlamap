-- Up migration
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS resources (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    address VARCHAR(255),
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION,
    diagnoses TEXT,  -- Store as comma-separated string
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_resources_location ON resources USING gist (
    ST_SetSRID(ST_MakePoint(longitude, latitude), 4326)
);
CREATE INDEX idx_resources_diagnoses ON resources USING gin (diagnoses);

-- Down migration
DROP TABLE IF EXISTS resources;
DROP INDEX IF EXISTS idx_resources_location;
DROP INDEX IF EXISTS idx_resources_diagnoses;