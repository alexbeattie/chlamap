"use client";

import { useState, useEffect } from 'react';
import { AlertCircle } from 'lucide-react';
import { Alert, AlertDescription } from '@/components/ui/alert';

export default function SubmissionsList() {
  const [submissions, setSubmissions] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [page, setPage] = useState(1);
  const [search, setSearch] = useState('');
  const [sortBy, setSortBy] = useState('created_at');
  const [sortOrder, setSortOrder] = useState('desc');
  const [dateFilter, setDateFilter] = useState('');
  const [totalPages, setTotalPages] = useState(1);

  const fetchSubmissions = async () => {
    try {
      setLoading(true);
      setError(null);

      // Construct query parameters
      const params = new URLSearchParams({
        page: page.toString(),
        limit: '10',
        sort_by: sortBy,
        sort_order: sortOrder
      });

      if (search) {
        params.append('search', search);
      }

      if (dateFilter) {
        params.append('date_filter', dateFilter);
      }

      const response = await fetch(`/api/submissions?${params.toString()}`);

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const data = await response.json();

      // Update state with the response data
      setSubmissions(data.submissions || []);
      setTotalPages(Math.ceil((data.total || 0) / 10));
    } catch (error) {
      console.error('Error fetching submissions:', error);
      setError('Failed to load submissions. Please try again.');
    } finally {
      setLoading(false);
    }
  };

  const handleDelete = async (id) => {
    if (!window.confirm('Are you sure you want to delete this submission?')) {
      return;
    }

    try {
      const response = await fetch(`/api/submissions/${id}`, {
        method: 'DELETE',
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      await fetchSubmissions();
    } catch (error) {
      console.error('Error deleting submission:', error);
      setError('Failed to delete submission. Please try again.');
    }
  };

  const handleSort = (column) => {
    if (sortBy === column) {
      setSortOrder(sortOrder === 'asc' ? 'desc' : 'asc');
    } else {
      setSortBy(column);
      setSortOrder('asc');
    }
  };

  useEffect(() => {
    fetchSubmissions();
  }, [page, search, sortBy, sortOrder, dateFilter]);

  if (loading && submissions.length === 0) {
    return (
      <div className="flex justify-center items-center h-64">
        <div className="text-gray-600">Loading submissions...</div>
      </div>
    );
  }

  return (
    <div className="bg-white rounded-lg shadow p-6">
      {error && (
        <Alert variant="destructive" className="mb-4">
          <AlertCircle className="h-4 w-4" />
          <AlertDescription>{error}</AlertDescription>
        </Alert>
      )}

      <div className="flex gap-4 mb-6">
        <input
          type="text"
          value={search}
          onChange={(e) => setSearch(e.target.value)}
          placeholder="Search submissions..."
          className="flex-1 p-2 border rounded focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
        />
        <select
          value={dateFilter}
          onChange={(e) => setDateFilter(e.target.value)}
          className="p-2 border rounded focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
        >
          <option value="">All time</option>
          <option value="today">Today</option>
          <option value="week">This week</option>
          <option value="month">This month</option>
        </select>
      </div>

      <div className="overflow-x-auto">
        <table className="w-full text-gray-900">
          <thead>
            <tr className="bg-gray-50">
              <th
                onClick={() => handleSort('name')}
                className="cursor-pointer text-left p-3 hover:bg-gray-100"
              >
                Name {sortBy === 'name' && (sortOrder === 'asc' ? '↑' : '↓')}
              </th>
              <th
                onClick={() => handleSort('email')}
                className="cursor-pointer text-left p-3 hover:bg-gray-100"
              >
                Email {sortBy === 'email' && (sortOrder === 'asc' ? '↑' : '↓')}
              </th>
              <th className="text-left p-3">Message</th>
              <th
                onClick={() => handleSort('created_at')}
                className="cursor-pointer text-left p-3 hover:bg-gray-100"
              >
                Date {sortBy === 'created_at' && (sortOrder === 'asc' ? '↑' : '↓')}
              </th>
              <th className="text-left p-3">Actions</th>
            </tr>
          </thead>
          <tbody>
            {submissions.map((submission) => (
              <tr key={submission.id} className="border-t hover:bg-gray-50">
                <td className="p-3">{submission.name}</td>
                <td className="p-3">{submission.email}</td>
                <td className="p-3">
                  <div className="max-w-xs truncate">{submission.message}</div>
                </td>
                <td className="p-3">
                  {new Date(submission.created_at).toLocaleDateString()}
                </td>
                <td className="p-3">
                  <button
                    onClick={() => handleDelete(submission.id)}
                    className="text-red-500 hover:text-red-700 focus:outline-none"
                  >
                    Delete
                  </button>
                </td>
              </tr>
            ))}
            {submissions.length === 0 && !loading && (
              <tr>
                <td colSpan="5" className="p-3 text-center text-gray-500">
                  No submissions found
                </td>
              </tr>
            )}
          </tbody>
        </table>
      </div>

      <div className="mt-4 flex justify-between items-center">
        <button
          onClick={() => setPage(prev => Math.max(prev - 1, 1))}
          disabled={page === 1}
          className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 disabled:bg-gray-300 disabled:cursor-not-allowed"
        >
          Previous
        </button>
        <span className="text-sm text-gray-600">
          Page {page} of {totalPages}
        </span>
        <button
          onClick={() => setPage(prev => prev + 1)}
          disabled={page >= totalPages}
          className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 disabled:bg-gray-300 disabled:cursor-not-allowed"
        >
          Next
        </button>
      </div>
    </div>
  );
}