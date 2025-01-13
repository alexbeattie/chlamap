// src/app/api/resources/search/route.js
import { executeQuery } from '@/lib/db';
import { NextResponse } from 'next/server';

export async function GET(request) {
  const { searchParams } = new URL(request.url);
  const q = searchParams.get('q')?.split(',') || null;
  const lat = parseFloat(searchParams.get('lat'));
  const lon = parseFloat(searchParams.get('lon'));
  const radius = parseFloat(searchParams.get('radius')); // in meters

  try {
    // Use PostGIS for spatial search
    let query = `
      SELECT 
        id,
        name,
        description,
        latitude,
        longitude,
        diagnoses,
        address,
        contact_info::text as contact_info,
        ST_Distance(
          ST_SetSRID(ST_MakePoint(longitude, latitude), 4326)::geography,
          ST_SetSRID(ST_MakePoint($1, $2), 4326)::geography
        ) as distance
      FROM resources
      WHERE ST_DWithin(
        ST_SetSRID(ST_MakePoint(longitude, latitude), 4326)::geography,
        ST_SetSRID(ST_MakePoint($1, $2), 4326)::geography,
        $3
      )
    `;

    const params = [lon, lat, radius];
    let paramCount = 3;

    // Add diagnosis filter if provided
    if (q) {
      paramCount++;
      query += ` AND diagnoses && $${paramCount}`;
      params.push(Array.isArray(q) ? q : [q]);
    }

    query += ` ORDER BY distance`;

    const result = await executeQuery(query, params);

    // Parse the contact_info JSON string back to an object
    const resources = result.rows.map(resource => ({
      ...resource,
      contact_info: JSON.parse(resource.contact_info)
    }));

    return NextResponse.json(resources);
  } catch (error) {
    console.error('Search error:', error);
    return NextResponse.json(
      { error: 'Search failed' },
      { status: 500 }
    );
  }
}