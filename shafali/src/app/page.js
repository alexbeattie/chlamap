'use client';

import ResourceMap from '@/components/ResourceMap';
import SearchResults from '@/components/SearchResults';

export default function Home() {
  return (
    <div className="h-screen w-screen relative flex">
      {/* Search Results Sidebar */}
      <div className="w-[400px] h-full overflow-y-auto bg-white shadow-lg z-10">
        <SearchResults />
      </div>

      {/* Full-screen Map */}
      <div className="flex-1 h-full">
        <ResourceMap />
      </div>
    </div>
  );
}