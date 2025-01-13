// src/app/api/resources/[id]/route.js
import { executeQuery } from '@/lib/db';
import { NextResponse } from 'next/server';

export async function GET(request, { params }) {
  try {
    const { id } = params;
    const result = await executeQuery(
      'SELECT * FROM resources WHERE id = $1',
      [id]
    );

    if (result.rows.length === 0) {
      return NextResponse.json(
        { error: 'Resource not found' },
        { status: 404 }
      );
    }

    return NextResponse.json(result.rows[0]);
  } catch (error) {
    return NextResponse.json(
      { error: 'Failed to fetch resource' },
      { status: 500 }
    );
  }
}

export async function PUT(request, { params }) {
  try {
    const { id } = params;
    const { name, description, diagnosis, latitude, longitude, address, contact_info } = await request.json();

    const result = await executeQuery(
      `UPDATE resources 
       SET name=$1, description=$2, diagnosis=$3, latitude=$4, longitude=$5, 
           address=$6, contact_info=$7
       WHERE id=$8
       RETURNING *`,
      [name, description, diagnosis, latitude, longitude, address, contact_info, id]
    );

    if (result.rows.length === 0) {
      return NextResponse.json(
        { error: 'Resource not found' },
        { status: 404 }
      );
    }

    return NextResponse.json(result.rows[0]);
  } catch (error) {
    return NextResponse.json(
      { error: 'Failed to update resource' },
      { status: 500 }
    );
  }
}
