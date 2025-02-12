"use client";
import { useState } from "react";
import { Search } from "lucide-react";

export default function SearchPanel() {
  const [isCollapsed, setIsCollapsed] = useState(false);
  const [searchInput, setSearchInput] = useState("");
  const [selectedDiagnoses, setSelectedDiagnoses] = useState([]);
  const availableDiagnoses = ["ADHD", "Autism", "Anxiety", "Depression"];

  return (
    <div
      className={`fixed left-0 top-0 h-full bg-white shadow-lg transition-all duration-300 border-r border-gray-200
        ${isCollapsed ? "w-14" : "w-[360px]"}
      `}
    >
      {/* Sidebar Content */}
      {!isCollapsed && (
        <div className="p-6 space-y-6">
          {/* Title */}
          <h1 className="text-2xl font-semibold text-black">Find Resources</h1>

          {/* Search Bar */}
          <div className="relative">
            <Search className="absolute left-3 top-2.5 h-5 w-5 text-gray-500" />
            <input
              type="text"
              placeholder="Search by location..."
              className="w-full pl-10 p-2 rounded-md border border-gray-300 bg-white text-black shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>

          {/* Resource Types */}
          <div>
            <h2 className="mb-2 text-lg font-medium text-gray-800">Resource Types</h2>
            <div className="space-y-2">
              {availableDiagnoses.map((diagnosis) => (
                <label key={diagnosis} className="flex items-center text-gray-700">
                  <input
                    type="checkbox"
                    className="mr-3 accent-blue-600"
                  />
                  {diagnosis}
                </label>
              ))}
            </div>
          </div>

          {/* Distance Input */}
          <div>
            <h2 className="mb-2 text-lg font-medium text-gray-800">Distance</h2>
            <input
              type="text"
              placeholder="Within 5 miles"
              className="w-full p-2 rounded-md border border-gray-300 bg-white text-black shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>

          {/* Apply Filters Button */}
          <button className="w-full bg-blue-600 text-white p-2 rounded-md hover:bg-blue-700 transition-all">
            Apply Filters
          </button>
        </div>
      )}

      {/* Toggle Button */}
      <button
        onClick={() => setIsCollapsed(!isCollapsed)}
        className="absolute top-6 -right-5 bg-white border border-gray-300 text-gray-600 shadow-md p-2 rounded-full hover:bg-gray-100 transition-all"
      >
        {isCollapsed ? "→" : "←"}
      </button>
    </div>
  );
}
