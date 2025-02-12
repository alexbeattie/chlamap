// components/HospitalSidebar.tsx
import React from "react";

interface Hospital {
  id: string;
  name: string;
  address?: string;
}

interface SidebarProps {
  hospitals: Hospital[];
  selectedHospital: Hospital | null;
  onHospitalSelect: (hospital: Hospital) => void;
}

const HospitalSidebar: React.FC<SidebarProps> = ({
  hospitals,
  selectedHospital,
  onHospitalSelect,
}) => {
  const getItemClassName = (hospital: Hospital) => {
    const baseClass = "p-2 cursor-pointer";
    return `${baseClass} ${
      selectedHospital?.id === hospital.id ? "bg-gray-300" : "hover:bg-gray-100"
    }`;
  };

  return (
    <div className="p-4">
      <h2 className="text-lg font-semibold">Hospitals</h2>
      <ul className="mt-4">
        {hospitals.map((hospital) => (
          <li
            key={hospital.id}
            onClick={() => onHospitalSelect(hospital)}
            className={getItemClassName(hospital)}
          >
            {hospital.name}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default HospitalSidebar;
