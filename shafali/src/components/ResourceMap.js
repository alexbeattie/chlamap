import { useState, useEffect } from 'react';
import { GoogleMap, LoadScript, Marker, InfoWindow } from '@react-google-maps/api';
import { useSearchParams } from 'next/navigation';

const mapContainerStyle = {
  width: '100%',
  height: '100vh'
};

const defaultCenter = { lat: 40.7128, lng: -74.0060 }; // New York City as default
const libraries = ['places', 'geometry'];
const availableDiagnoses = ["ADHD", "Autism", "Anxiety", "Depression"];

const mapOptions = {
  disableDefaultUI: false,
  zoomControl: true,
  mapTypeControl: true,
  scaleControl: true,
  streetViewControl: true,
  rotateControl: true,
  fullscreenControl: true
};

export default function ResourceMap() {
  const searchParams = useSearchParams();
  const [map, setMap] = useState(null);
  const [center, setCenter] = useState(defaultCenter);
  const [resources, setResources] = useState([]);
  const [selectedResource, setSelectedResource] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [loadError, setLoadError] = useState(null);
  const [selectedDiagnoses, setSelectedDiagnoses] = useState([]);
  const [searchInput, setSearchInput] = useState('');
  const [searchRadius, setSearchRadius] = useState(16093.4); // 10 miles in meters

  const fetchNearbyResources = async (lat, lng, diagnoses = []) => {
    try {
      setLoading(true);
      // Using the Gin backend endpoint
      let url = `/api/resources/nearby?lat=${lat}&lng=${lng}`;

      if (diagnoses.length > 0) {
        url += `&q=${encodeURIComponent(diagnoses.join(','))}`;
      }

      const response = await fetch(url);
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const data = await response.json();
      setResources(data);
      setError(null);
    } catch (err) {
      setError('Failed to load resources. Please try again later.');
      console.error('Fetch error:', err);
    } finally {
      setLoading(false);
    }
  };

  const handleSearch = () => {
    if (searchInput && map) {
      const geocoder = new window.google.maps.Geocoder();
      geocoder.geocode({ address: searchInput }, (results, status) => {
        if (status === 'OK' && results[0]) {
          const location = results[0].geometry.location;
          const newCenter = {
            lat: location.lat(),
            lng: location.lng()
          };
          setCenter(newCenter);
          map.panTo(newCenter);

          // Fetch resources near the searched location
          fetchNearbyResources(newCenter.lat, newCenter.lng, selectedDiagnoses);
        }
      });
    }
  };

  const handleDiagnosisToggle = (diagnosis) => {
    const newDiagnoses = selectedDiagnoses.includes(diagnosis)
      ? selectedDiagnoses.filter(d => d !== diagnosis)
      : [...selectedDiagnoses, diagnosis];

    setSelectedDiagnoses(newDiagnoses);

    // Refetch resources with updated diagnoses
    if (map) {
      fetchNearbyResources(center.lat, center.lng, newDiagnoses);
    }
  };

  const handleLoadError = (error) => {
    console.error('Error loading Google Maps:', error);
    setLoadError('Failed to load Google Maps. Please check your internet connection and try again.');
  };

  const handleMapLoad = (map) => {
    console.log('Map loaded successfully');
    setMap(map);
  };

  useEffect(() => {
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(
        (position) => {
          const newCenter = {
            lat: position.coords.latitude,
            lng: position.coords.longitude
          };
          setCenter(newCenter);
          if (map) {
            map.panTo(newCenter);
            // Fetch resources near user's location
            fetchNearbyResources(newCenter.lat, newCenter.lng, selectedDiagnoses);
          }
        },
        (error) => {
          console.log('Geolocation error:', error);
          // Fetch resources for default location if geolocation fails
          fetchNearbyResources(defaultCenter.lat, defaultCenter.lng, selectedDiagnoses);
        }
      );
    }
  }, [map]);

  if (loading && !map) {
    return (
      <div className="flex-1 flex items-center justify-center">
        <div className="text-gray-600">Loading map...</div>
      </div>
    );
  }

  if (loadError) {
    return (
      <div className="flex-1 flex items-center justify-center">
        <div className="text-red-600 max-w-md text-center p-4">
          {loadError}
        </div>
      </div>
    );
  }

  return (
    <div className="h-full relative">
      <div className="absolute top-4 left-1/2 transform -translate-x-1/2 z-20 w-full max-w-2xl px-4">
        <div className="bg-white/90 backdrop-blur-sm rounded-lg shadow-lg p-4">
          <div className="flex gap-2">
            <input
              type="text"
              placeholder="Search locations..."
              value={searchInput}
              onChange={(e) => setSearchInput(e.target.value)}
              onKeyPress={(e) => e.key === 'Enter' && handleSearch()}
              className="flex-1 px-4 py-2 border rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
            />
            <button
              onClick={handleSearch}
              className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
            >
              Search
            </button>
          </div>
          <div className="mt-2 flex flex-wrap gap-2">
            {availableDiagnoses.map((diagnosis) => (
              <button
                key={diagnosis}
                onClick={() => handleDiagnosisToggle(diagnosis)}
                className={`px-3 py-1 rounded-full text-sm transition-colors ${selectedDiagnoses.includes(diagnosis)
                    ? 'bg-blue-600 text-white'
                    : 'bg-gray-100 text-gray-700 hover:bg-gray-200'
                  }`}
              >
                {diagnosis}
              </button>
            ))}
          </div>
        </div>
      </div>

      <LoadScript
        googleMapsApiKey={process.env.NEXT_PUBLIC_GOOGLE_MAPS_API_KEY}
        libraries={libraries}
        onError={handleLoadError}
      >
        <GoogleMap
          mapContainerStyle={mapContainerStyle}
          center={center}
          zoom={13}
          onLoad={handleMapLoad}
          options={mapOptions}
        >
          {resources.map((resource) => (
            <Marker
              key={resource.id}
              position={{
                lat: parseFloat(resource.latitude),
                lng: parseFloat(resource.longitude)
              }}
              onClick={() => setSelectedResource(resource)}
            />
          ))}

          {selectedResource && (
            <InfoWindow
              position={{
                lat: parseFloat(selectedResource.latitude),
                lng: parseFloat(selectedResource.longitude)
              }}
              onCloseClick={() => setSelectedResource(null)}
            >
              <div className="p-2">
                <h3 className="font-semibold mb-1">{selectedResource.name}</h3>
                <p className="text-sm mb-2">{selectedResource.address}</p>
                {selectedResource.phone && (
                  <p className="text-sm">{selectedResource.phone}</p>
                )}
                {selectedResource.website && (
                  <a
                    href={selectedResource.website}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="text-sm text-blue-600 hover:text-blue-800"
                  >
                    Visit Website
                  </a>
                )}
              </div>
            </InfoWindow>
          )}
        </GoogleMap>
      </LoadScript>
    </div>
  );
}