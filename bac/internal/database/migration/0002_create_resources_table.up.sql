-- Up migration
CREATE TABLE resources (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    latitude DECIMAL(10, 8) NOT NULL,
    longitude DECIMAL(11, 8) NOT NULL,
    diagnoses TEXT[] NOT NULL,
    address TEXT,
    contact_info JSONB,
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