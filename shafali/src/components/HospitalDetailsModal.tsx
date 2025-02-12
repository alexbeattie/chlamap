// components/HospitalDetailsModal.tsx
import React from "react";

interface Hospital {
  id: string;
  name: string;
  address?: string;
}

interface ModalProps {
  hospital: Hospital;
  onClose: () => void;
  onGetDirections: () => void;
}

const HospitalDetailsModal: React.FC<ModalProps> = ({
  hospital,
  onClose,
  onGetDirections,
}) => {
  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
      <div className="bg-white p-6 rounded-lg w-96">
        <h2 className="text-xl font-semibold">{hospital.name}</h2>
        <p className="mt-2">{hospital.address}</p>
        <div className="flex gap-2 mt-4">
          <button
            onClick={onGetDirections}
            className="bg-gray-500 text-white px-4 py-2 rounded hover:bg-gray-600"
          >
            Get Directions
          </button>
          <button
            onClick={onClose}
            className="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600"
          >
            Close
          </button>
        </div>
      </div>
    </div>
  );
};

export default HospitalDetailsModal;
