// pages/submissions.js
"use client";

import SubmissionsList from '@/components/SubmissionsList';

export default function SubmissionsPage() {
  return (
    <div className="container mx-auto p-8">
      <h1 className="text-2xl font-bold mb-6">View Submissions</h1>
      <SubmissionsList />
    </div>
  );
}