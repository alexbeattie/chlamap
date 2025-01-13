// src/components/ResourceSearch.jsx
'use client';
import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { Search } from 'lucide-react';

export default function ResourceSearch() {  // Changed to default export
  const router = useRouter();
  const [address, setAddress] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');
  const [searchRadius, setSearchRadius] = useState(10);
  const [selectedTypes, setSelectedTypes] = useState([]);

  const resourceTypes = ['ADHD', 'Autism', 'Anxiety', 'Depression'];

  const handleSearch = async () => {
    setLoading(true);
    setError('');

    try {
      // First, geocode the address
      const geocodeResponse = await fetch(`/api/geocode?address=${encodeURIComponent(address)}`);
      const geocodeData = await geocodeResponse.json();

      if (!geocodeData.results?.[0]?.geometry?.location) {
        throw new Error('Invalid address');
      }

      const { lat, lng } = geocodeData.results[0].geometry.location;

      // Now search for resources
      const searchParams = new URLSearchParams({
        lat: lat.toString(),
        lon: lng.toString(),
        radius: (searchRadius * 1609.34).toString(), // Convert miles to meters
        ...(selectedTypes.length && { q: selectedTypes.join(',') })
      });

      const searchResponse = await fetch(`/api/resources/search?${searchParams}`);
      const resources = await searchResponse.json();

      // Update URL with search params
      router.push(`/results?${searchParams.toString()}`);

    } catch (err) {
      setError('Failed to search resources. Please try again.');
      console.error('Search error:', err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="max-w-2xl mx-auto p-6 space-y-6">
      <div className="relative">
        <input
          type="text"
          placeholder="Enter your address or zip code"
          value={address}
          onChange={(e) => setAddress(e.target.value)}
          className="w-full px-4 py-2 border rounded-lg pr-10 focus:ring-2 focus:ring-blue-500 outline-none"
        />
        <Search className="absolute right-3 top-2.5 text-gray-400" size={20} />
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-2">
          Search Radius (miles)
        </label>
        <input
          type="range"
          min="1"
          max="50"
          value={searchRadius}
          onChange={(e) => setSearchRadius(Number(e.target.value))}
          className="w-full"
        />
        <div className="text-sm text-gray-600 text-center">{searchRadius} miles</div>
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-2">
          Resource Types
        </label>
        <div className="flex flex-wrap gap-2">
          {resourceTypes.map(type => (
            <button
              key={type}
              onClick={() => {
                setSelectedTypes(prev =>
                  prev.includes(type)
                    ? prev.filter(t => t !== type)
                    : [...prev, type]
                );
              }}
              className={`px-3 py-1 rounded-full text-sm transition-colors
                ${selectedTypes.includes(type)
                  ? 'bg-blue-600 text-white'
                  : 'bg-gray-100 text-gray-700 hover:bg-gray-200'}`}
            >
              {type}
            </button>
          ))}
        </div>
      </div>

      {error && (
        <div className="text-red-600 text-sm">
          {error}
        </div>
      )}

      <button
        onClick={handleSearch}
        disabled={loading || !address}
        className={`w-full py-2 px-4 rounded-lg text-white font-medium
          ${loading || !address
            ? 'bg-blue-400 cursor-not-allowed'
            : 'bg-blue-600 hover:bg-blue-700'}`}
      >
        {loading ? 'Searching...' : 'Search Resources'}
      </button>
    </div>
  );
}