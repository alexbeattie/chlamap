// pages/submit.js

"use client";

import { useEffect, useRef } from 'react';
import DataEntryForm from '@/components/DataEntryForm';

export default function SubmitPage() {
  return (
    <div className="container mx-auto py-8">
      <h1 className="text-2xl font-bold mb-6">Submit New Resource</h1>
      <DataEntryForm />
    </div>
  );
}