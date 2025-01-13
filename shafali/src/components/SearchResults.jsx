// src/components/SearchResults.jsx
'use client';
import { useState, useEffect } from 'react';
import { useSearchParams } from 'next/navigation';
import { MapPin, Phone, Clock, Mail, Calendar } from 'lucide-react';

export default function SearchResults() {
  const searchParams = useSearchParams();
  const [resources, setResources] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    const fetchResults = async () => {
      setLoading(true);
      try {
        // Get all the current search parameters
        const params = new URLSearchParams();
        searchParams.forEach((value, key) => {
          params.append(key, value);
        });

        const response = await fetch(`/api/resources/search?${params.toString()}`);
        if (!response.ok) throw new Error('Failed to fetch resources');
        const data = await response.json();
        setResources(data);
      } catch (err) {
        setError('Failed to load resources. Please try again.');
        console.error('Results error:', err);
      } finally {
        setLoading(false);
      }
    };

    fetchResults();
  }, [searchParams]);

  if (loading) {
    return (
      <div className="flex justify-center items-center min-h-[50vh]">
        <div className="text-gray-600">Loading resources...</div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="max-w-2xl mx-auto p-6">
        <div className="bg-red-50 text-red-600 p-4 rounded-lg">
          {error}
        </div>
      </div>
    );
  }

  if (resources.length === 0) {
    return (
      <div className="max-w-2xl mx-auto p-6">
        <div className="text-center text-gray-600">
          No resources found in this area. Try expanding your search radius.
        </div>
      </div>
    );
  }

  return (
    <div className="p-4">
      <h2 className="text-2xl font-bold mb-6 text-gray-400">
        Found {resources.length} Resources
      </h2>

      <div className="space-y-4">
        {resources.map((resource) => (
          <div
            key={resource.id || Math.random()}
            className="bg-white rounded-lg shadow-md p-4 hover:shadow-lg transition-shadow"
          >
            <h3 className="text-lg font-semibold mb-2 text-gray-800">
              {resource.name}
            </h3>

            <div className="space-y-2 text-sm text-gray-600">
              {resource.address && (
                <div className="flex items-center gap-2">
                  <MapPin size={16} />
                  <span>{resource.address}</span>
                </div>
              )}

              {resource.contact_info && (
                <>
                  {resource.contact_info.phone && (
                    <div className="flex items-center gap-2">
                      <Phone size={16} />
                      <span>{resource.contact_info.phone}</span>
                    </div>
                  )}
                  {resource.contact_info.email && (
                    <div className="flex items-center gap-2">
                      <Mail size={16} />
                      <span>{resource.contact_info.email}</span>
                    </div>
                  )}
                  {resource.contact_info.hours && (
                    <div className="flex items-center gap-2">
                      <Calendar size={16} />
                      <span>{resource.contact_info.hours}</span>
                    </div>
                  )}
                </>
              )}

              {resource.distance !== undefined && (
                <div className="flex items-center gap-2">
                  <Clock size={16} />
                  <span>{(resource.distance / 1609.34).toFixed(1)} miles away</span>
                </div>
              )}
            </div>

            {resource.diagnoses && (
              <div className="mt-3 flex flex-wrap gap-2">
                {resource.diagnoses.map((type) => (
                  <span
                    key={type}
                    className="px-2 py-1 bg-blue-100 text-blue-800 text-xs rounded-full"
                  >
                    {type}
                  </span>
                ))}
              </div>
            )}

            {resource.description && (
              <p className="mt-3 text-sm text-gray-600">
                {resource.description}
              </p>
            )}
          </div>
        ))}
      </div>
    </div>
  );
}