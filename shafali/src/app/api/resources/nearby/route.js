// src/app/api/resources/nearby/route.js
import { executeQuery } from '@/lib/db';
import { NextResponse } from 'next/server';

export async function GET(request) {
  // Handle nearby resources logic here
  // Example: ?latitude=34.0522&longitude=-118.2437
  const { searchParams } = new URL(request.url);
  const lat = searchParams.get('latitude');
  const lng = searchParams.get('longitude');

  try {
    const result = await executeQuery(
      `SELECT *, 
       point($1, $2) <@> point(longitude, latitude)::point as distance
       FROM resources
       ORDER BY point($1, $2) <@> point(longitude, latitude)
       LIMIT 10`,
      [lng, lat]
    );

    return NextResponse.json(result.rows);
  } catch (error) {
    return NextResponse.json(
      { error: 'Failed to fetch nearby resources' },
      { status: 500 }
    );
  }
}