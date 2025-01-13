"use client";
import { useState, useEffect } from 'react';
import { ChevronLeft, ChevronRight, PlusCircle } from 'lucide-react';
import Link from 'next/link';

export default function Sidebar() {
  const [isExpanded, setIsExpanded] = useState(undefined);

  useEffect(() => {
    setIsExpanded(true);
  }, []);

  if (isExpanded === undefined) {
    return null;
  }

  return (
    <div className={`
      fixed inset-y-0 left-0 transform transition-all duration-300
      bg-white shadow-lg z-30
      ${isExpanded ? 'w-[360px]' : 'w-16'}
      ${isExpanded ? 'translate-x-0' : '-translate-x-[calc(100%-64px)]'}
    `}>
      <div className="p-6">
        <div className="flex justify-between items-center mb-6">
          <h1 className={`text-xl font-medium ${!isExpanded && 'hidden'}`}>
            Find Resources
          </h1>
          {isExpanded && (
            <Link
              href="/resources/add"
              className="flex items-center gap-2 px-3 py-2 text-sm text-blue-600 hover:bg-blue-50 rounded-lg transition-colors"
            >
              <PlusCircle size={16} />
              <span>Add Resource</span>
            </Link>
          )}
        </div>

        <div className={`transition-opacity duration-300 ${isExpanded ? 'opacity-100' : 'opacity-0'}`}>
          {isExpanded && (
            <div className="space-y-6">
              {/* Sort Options */}
              <div>
                <label className="text-sm text-gray-600">Sort by</label>
                <select className="mt-1 w-full p-2 border rounded bg-white">
                  <option>Most Active</option>
                  <option>Distance</option>
                </select>
              </div>

              {/* Map Settings */}
              <div>
                <label className="text-sm text-gray-600">Map Settings</label>
                <select className="mt-1 w-full p-2 border rounded bg-white">
                  <option>Road Map</option>
                  <option>Satellite</option>
                </select>
              </div>

              {/* Resource Types */}
              <div>
                <label className="text-sm text-gray-600">Resource Types</label>
                <div className="mt-2 space-y-2">
                  {['ADHD', 'Autism', 'Anxiety', 'Depression'].map(type => (
                    <label key={type} className="flex items-center">
                      <input type="checkbox" className="mr-2" />
                      <span className="text-sm">{type}</span>
                    </label>
                  ))}
                </div>
              </div>
            </div>
          )}
        </div>
      </div>

      <button
        onClick={() => setIsExpanded(!isExpanded)}
        className="absolute -right-3 top-6 bg-white rounded-full shadow-lg p-1.5 z-10"
      >
        {isExpanded ? <ChevronLeft size={16} /> : <ChevronRight size={16} />}
      </button>
    </div>
  );
}