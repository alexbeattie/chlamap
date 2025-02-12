'use client';

import ResourceMap from '@/components/ResourceMap';
import ResourceSearch from '@/components/ResourceSearch';
import SearchPanel from '@/components/SearchPanel';
import SearchResults from '@/components/SearchResults';
import Sidebar from '@/components/Sidebar';

export default function Home() {
  return (
    <div className="h-screen w-screen relative flex">
      {/* Search Results Sidebar */}
      <div className="w-[400px] h-full overflow-y-auto z-10">
        {/* <Sidebar /> */}
        <SearchPanel />
        {/* <SearchPanel /> */}
        {/* <ResourceSearch /> */}
        {/* <SearchResults /> */}
      </div>

      {/* Full-screen Map */}
      <div className="flex-1 h-full">
        <ResourceMap />
      </div>
    </div>
  );
}