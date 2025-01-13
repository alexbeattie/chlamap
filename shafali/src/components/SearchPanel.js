"use client";
import { useState } from 'react';
import { Search } from 'lucide-react';
import { useRouter, useSearchParams } from 'next/navigation';

export default function SearchPanel() {
  const [isCollapsed, setIsCollapsed] = useState(false);
  const [searchInput, setSearchInput] = useState('');
  const [selectedDiagnoses, setSelectedDiagnoses] = useState([]);
  const availableDiagnoses = ['ADHD', 'Autism', 'Anxiety', 'Depression'];

  return (
    <div className={`fixed left-0 top-0 h-full bg-black/90 text-white transition-all duration-300 
      ${isCollapsed ? 'w-12' : 'w-[380px]'}`}>

      {!isCollapsed ? (
        <div className="p-6 space-y-6">
          <h1 className="text-2xl">Find Resources</h1>

          <div className="relative">
            <Search className="absolute left-3 top-2.5 h-5 w-5 text-gray-500" />
            <input
              type="text"
              placeholder="Search by location..."
              className="w-full pl-10 p-2 rounded bg-white text-black"
            />
          </div>

          <div>
            <h2 className="mb-3">Resource Types</h2>
            <div className="space-y-2">
              {availableDiagnoses.map(diagnosis => (
                <label key={diagnosis} className="flex items-center">
                  <input type="checkbox" className="mr-3" />
                  {diagnosis}
                </label>
              ))}
            </div>
          </div>

          <div>
            <h2 className="mb-3">Distance</h2>
            <input
              type="text"
              placeholder="Within 5 miles"
              className="w-full p-2 rounded bg-white text-black"
            />
          </div>

          <button className="w-full bg-white text-black p-2 rounded">
            Apply Filters
          </button>
        </div>
      ) : null}

      <button
        onClick={() => setIsCollapsed(!isCollapsed)}
        className="absolute top-4 -right-4 bg-black text-white p-2 rounded-full"
      >
        {isCollapsed ? '→' : '←'}
      </button>
    </div>
  );
}