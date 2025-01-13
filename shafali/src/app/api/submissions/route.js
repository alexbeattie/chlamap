// src/app/api/submissions/route.js
import { executeQuery } from '@/lib/db';

export async function GET(request) {
  try {
    const { searchParams } = new URL(request.url);
    const page = parseInt(searchParams.get('page')) || 1;
    const limit = parseInt(searchParams.get('limit')) || 10;
    const search = searchParams.get('search') || '';
    const sortBy = searchParams.get('sortBy') || 'created_at';
    const sortOrder = searchParams.get('sortOrder') || 'desc';
    const dateFilter = searchParams.get('dateFilter') || '';

    const offset = (page - 1) * limit;

    let dateCondition = '';
    if (dateFilter === 'today') {
      dateCondition = 'AND created_at >= CURRENT_DATE';
    } else if (dateFilter === 'week') {
      dateCondition = 'AND created_at >= CURRENT_DATE - INTERVAL \'7 days\'';
    } else if (dateFilter === 'month') {
      dateCondition = 'AND created_at >= CURRENT_DATE - INTERVAL \'1 month\'';
    }

    const query = `
      SELECT * FROM form_submissions 
      WHERE (name ILIKE $1 OR email ILIKE $1 OR message ILIKE $1)
      ${dateCondition}
      ORDER BY ${sortBy} ${sortOrder}
      LIMIT $2 OFFSET $3
    `;

    const result = await executeQuery(query, [`%${search}%`, limit, offset]);

    const countQuery = `
      SELECT COUNT(*) FROM form_submissions 
      WHERE (name ILIKE $1 OR email ILIKE $1 OR message ILIKE $1)
      ${dateCondition}
    `;
    const countResult = await executeQuery(countQuery, [`%${search}%`]);

    return new Response(JSON.stringify({
      data: result.rows,
      pagination: {
        total: parseInt(countResult.rows[0].count),
        page,
        limit
      }
    }), {
      status: 200,
      headers: { 'Content-Type': 'application/json' },
    });
  } catch (error) {
    console.error('Error fetching submissions:', error);
    return new Response(JSON.stringify({ message: 'Error fetching data' }), {
      status: 500,
      headers: { 'Content-Type': 'application/json' },
    });
  }
}