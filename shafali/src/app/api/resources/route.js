// app/api/resources/route.js
import { Pool } from 'pg';
import { NextResponse } from 'next/server';

const pool = new Pool({
  user: process.env.DB_USER,
  host: process.env.DB_HOST,
  database: process.env.DB_NAME,
  password: process.env.DB_PASSWORD,
  port: parseInt(process.env.DB_PORT)
});

export async function POST(req) {
  try {
    const body = await req.json();

    const {
      name,
      description,
      latitude,
      longitude,
      diagnoses,
      address,
      contact_info
    } = body;

    const result = await pool.query(
      `INSERT INTO resources (
        id,
        name,
        description,
        latitude,
        longitude,
        diagnoses,
        address,
        contact_info
      ) VALUES (
        gen_random_uuid(),
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7
      )
      RETURNING id`,
      [
        name,
        description,
        latitude,
        longitude,
        diagnoses,
        address,
        JSON.stringify(contact_info)
      ]
    );

    return NextResponse.json({
      message: 'Resource created successfully',
      id: result.rows[0].id
    });

  } catch (error) {
    console.error('Error creating resource:', error);
    return NextResponse.json(
      { error: 'Error creating resource: ' + error.message },
      { status: 500 }
    );
  }
}