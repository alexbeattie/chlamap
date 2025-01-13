-- Up migration
-- First, we create a function that helps find nearby resources based on location and diagnoses
-- This function uses the spatial index we already created in the first migration
CREATE OR REPLACE FUNCTION find_nearby_resources(
    search_lat DECIMAL,
    search_lng DECIMAL,
    radius_miles DECIMAL,
    diagnosis_filter TEXT[] DEFAULT NULL
)
RETURNS TABLE (
    -- Return all relevant resource information plus calculated distance
    id UUID,
    name VARCHAR,
    description TEXT,
    address TEXT,
    distance_miles DECIMAL,
    diagnoses TEXT[],
    contact_info JSONB
) AS $$
BEGIN
    RETURN QUERY
    SELECT 
        r.id,
        r.name,
        r.description,
        r.address,
        -- Convert meters to miles and round for readable distances
        ROUND(
            (ST_Distance(
                ST_SetSRID(ST_MakePoint(r.longitude, r.latitude), 4326)::geography,
                ST_SetSRID(ST_MakePoint(search_lng, search_lat), 4326)::geography
            ) / 1609.344)::numeric,
            2
        ) as distance_miles,
        r.diagnoses,
        r.contact_info
    FROM 
        resources r
    WHERE 
        -- Use PostGIS ST_DWithin for efficient radius search
        -- This takes advantage of our spatial index
        ST_DWithin(
            ST_SetSRID(ST_MakePoint(r.longitude, r.latitude), 4326)::geography,
            ST_SetSRID(ST_MakePoint(search_lng, search_lat), 4326)::geography,
            radius_miles * 1609.344  -- Convert miles to meters for PostGIS
        )
        -- Use the GIN index we created for efficient diagnosis filtering
        AND (diagnosis_filter IS NULL OR r.diagnoses && diagnosis_filter)
    ORDER BY 
        -- Sort results by distance from search point
        ST_Distance(
            ST_SetSRID(ST_MakePoint(r.longitude, r.latitude), 4326)::geography,
            ST_SetSRID(ST_MakePoint(search_lng, search_lat), 4326)::geography
        );
END;
$$ LANGUAGE plpgsql;

-- Down migration
-- Simply remove the function if we need to roll back
DROP FUNCTION IF EXISTS find_nearby_resources;