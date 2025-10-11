import React from "react";
import { Filter } from "lucide-react";
import { Button } from "@/components/ui/button";

export default function AlertFilters({ activeFilter, setActiveFilter }) {
  const filters = [
    { value: "all", label: "All Alerts" },
    { value: "critical", label: "Critical" },
    { value: "high", label: "High" },
    { value: "medium", label: "Medium" },
    { value: "low", label: "Low" }
  ];

  return (
    <div className="flex flex-wrap gap-3">
      <div className="flex items-center gap-2 text-gray-400">
        <Filter className="w-4 h-4" />
        <span className="text-sm font-medium">Filter by Severity:</span>
      </div>
      {filters.map((filter) => (
        <Button
          key={filter.value}
          onClick={() => setActiveFilter(filter.value)}
          variant="outline"
          className={`transition-all duration-200 ${
            activeFilter === filter.value
              ? "bg-cyan-500/20 border-cyan-500 text-cyan-400"
              : "bg-gray-800 border-gray-700 text-gray-400 hover:bg-gray-700 hover:text-gray-200"
          }`}
        >
          {filter.label}
        </Button>
      ))}
    </div>
  );
}
