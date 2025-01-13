'use client';
import SearchResults from '@/components/SearchResults';

export default function ResultsPage() {
  return (
    <div className="min-h-screen bg-gray-50">
      <main className="container mx-auto py-12">
        <SearchResults />
      </main>
    </div>
  );
}