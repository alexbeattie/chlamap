"use client";

import { useState, useEffect } from "react";
import {
  GoogleMap,
  LoadScript,
  Marker,
  InfoWindow,
} from "@react-google-maps/api";

const mapContainerStyle = {
  width: "100%",
  height: "100vh",
};

const defaultCenter = { lat: 34.0522, lng: -118.2437 }; // Los Angeles as default
const libraries = ["places", "geometry"];

const mapOptions = {
  disableDefaultUI: false,
  zoomControl: true,
  mapTypeControl: true,
  streetViewControl: true,
  fullscreenControl: true,
};

export default function ResourceMap() {
  const [map, setMap] = useState(null);
  const [center, setCenter] = useState(defaultCenter);
  const [resources, setResources] = useState([]); // Regional Centers
  const [selectedResource, setSelectedResource] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  // 🏥 **Fetch Regional Centers from Backend**
  const fetchRegionalCenters = async () => {
    try {
      setLoading(true);
      const response = await fetch(
        "http://localhost:8080/api/regional-centers"
      ); // Ensure correct URL
      if (!response.ok)
        throw new Error(`HTTP error! status: ${response.status}`);
      const data = await response.json();
      console.log("Fetched regional centers:", data); // Debugging
      setResources(data);
      setError(null);
    } catch (err) {
      setError("Failed to load regional centers.");
      console.error("Fetch error:", err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchRegionalCenters();
  }, []);

  useEffect(() => {
    console.log("Updated resources:", resources);
  }, [resources]); // Log whenever resources updates
  
  return (
    <div className="flex h-screen">
      {/* 🏥 Sidebar */}
      <div className="w-80 flex-shrink-0 border-r border-gray-200 bg-white overflow-y-auto p-4">
        <h2 className="text-lg font-semibold mb-3">Regional Centers</h2>
        {loading ? (
          <p className="text-gray-600">Loading...</p>
        ) : error ? (
          <p className="text-red-600">{error}</p>
        ) : (
          <ul>
            {resources.map((resource) => (
              <li
                key={resource.id}
                className={`p-2 cursor-pointer rounded ${
                  selectedResource?.id === resource.id
                    ? "bg-blue-200"
                    : "hover:bg-gray-100"
                }`}
                onClick={() => setSelectedResource(resource)}
              >
                <strong>{resource.name}</strong>
                <p className="text-sm text-gray-600">{resource.address}</p>
              </li>
            ))}
          </ul>
        )}
      </div>

      {/* 🗺️ Map */}
      <div className="flex-1 relative">
        <LoadScript
          googleMapsApiKey={process.env.NEXT_PUBLIC_GOOGLE_MAPS_API_KEY}
          libraries={libraries}
        >
          <GoogleMap
            mapContainerStyle={mapContainerStyle}
            center={center}
            zoom={10}
            onLoad={setMap}
            options={mapOptions}
          >
            {resources
              .filter((resource) => resource.latitude && resource.longitude) // Filter out invalid locations
              .map((resource) => (
                <Marker
                  key={resource.id}
                  position={{
                    lat: parseFloat(resource.latitude),
                    lng: parseFloat(resource.longitude),
                  }}
                  onClick={() => setSelectedResource(resource)}
                />
              ))}

            {/* 📍 Info Window */}
            {selectedResource &&
              selectedResource.latitude &&
              selectedResource.longitude && (
                <InfoWindow
                  position={{
                    lat: parseFloat(selectedResource.latitude),
                    lng: parseFloat(selectedResource.longitude),
                  }}
                  onCloseClick={() => setSelectedResource(null)}
                >
                  <div className="p-2">
                    <h3 className="font-semibold">{selectedResource.name}</h3>
                    <p className="text-sm">{selectedResource.address}</p>
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
    </div>
  );
}
