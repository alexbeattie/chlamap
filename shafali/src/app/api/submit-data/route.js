// src/app/api/submit-data/route.js
import { executeQuery } from '@/lib/db';

export async function POST(req) {
  try {
    const body = await req.json();
    const { name, email, message } = body;

    const query = `
      INSERT INTO form_submissions (name, email, message)
      VALUES ($1, $2, $3)
      RETURNING id
    `;

    await executeQuery(query, [name, email, message]);
    return new Response(JSON.stringify({ message: 'Success' }), {
      status: 200,
      headers: { 'Content-Type': 'application/json' },
    });
  } catch (error) {
    console.error(error);
    return new Response(JSON.stringify({ message: 'Error submitting data' }), {
      status: 500,
      headers: { 'Content-Type': 'application/json' },
    });
  }
}