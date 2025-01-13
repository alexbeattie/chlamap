// src/app/api/submissions/[id]/route.js
import { executeQuery } from '@/lib/db';

export async function DELETE(request, { params }) {
  try {
    const { id } = params;
    const query = 'DELETE FROM form_submissions WHERE id = $1';
    await executeQuery(query, [id]);

    return new Response(JSON.stringify({ message: 'Submission deleted' }), {
      status: 200,
      headers: { 'Content-Type': 'application/json' },
    });
  } catch (error) {
    console.error('Error deleting submission:', error);
    return new Response(JSON.stringify({ message: 'Error deleting submission' }), {
      status: 500,
      headers: { 'Content-Type': 'application/json' },
    });
  }
}