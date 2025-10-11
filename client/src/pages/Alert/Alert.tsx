import React, { useState, useEffect } from "react";
import { Alert } from "@/entities/Alert";
import { AlertTriangle, Search } from "lucide-react";
import { Input } from "@/components/ui/input";
import AlertFilters from "../components/alerts/AlertFilters";
import AlertCard from "../components/alerts/AlertCard";

export default function Alerts() {
  const [alerts, setAlerts] = useState([]);
  const [filteredAlerts, setFilteredAlerts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [activeFilter, setActiveFilter] = useState("all");
  const [searchTerm, setSearchTerm] = useState("");

  useEffect(() => {
    loadAlerts();
  }, []);

  useEffect(() => {
    filterAlerts();
  }, [alerts, activeFilter, searchTerm]);

  const loadAlerts = async () => {
    setLoading(true);
    const data = await Alert.list("-detected_at");
    setAlerts(data);
    setLoading(false);
  };

  const filterAlerts = () => {
    let filtered = alerts;

    if (activeFilter !== "all") {
      filtered = filtered.filter(alert => alert.severity === activeFilter);
    }

    if (searchTerm) {
      filtered = filtered.filter(alert =>
        alert.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
        alert.description.toLowerCase().includes(searchTerm.toLowerCase())
      );
    }

    setFilteredAlerts(filtered);
  };

  return (
    <div className="space-y-8">
      {/* Header */}
      <div className="flex flex-col gap-2">
        <h1 className="text-3xl font-bold text-white flex items-center gap-3">
          <AlertTriangle className="w-8 h-8 text-cyan-400" />
          Security Alerts
        </h1>
        <p className="text-gray-400">Monitor and manage security threats in real-time</p>
      </div>

      {/* Stats Banner */}
      <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
        <div className="bg-gradient-to-br from-red-500/10 to-rose-500/10 border border-red-500/30 rounded-lg p-4">
          <p className="text-2xl font-bold text-red-400">
            {alerts.filter(a => a.severity === "critical").length}
          </p>
          <p className="text-sm text-gray-400">Critical</p>
        </div>
        <div className="bg-gradient-to-br from-orange-500/10 to-amber-500/10 border border-orange-500/30 rounded-lg p-4">
          <p className="text-2xl font-bold text-orange-400">
            {alerts.filter(a => a.severity === "high").length}
          </p>
          <p className="text-sm text-gray-400">High</p>
        </div>
        <div className="bg-gradient-to-br from-yellow-500/10 to-yellow-600/10 border border-yellow-500/30 rounded-lg p-4">
          <p className="text-2xl font-bold text-yellow-400">
            {alerts.filter(a => a.severity === "medium").length}
          </p>
          <p className="text-sm text-gray-400">Medium</p>
        </div>
        <div className="bg-gradient-to-br from-blue-500/10 to-cyan-500/10 border border-blue-500/30 rounded-lg p-4">
          <p className="text-2xl font-bold text-blue-400">
            {alerts.filter(a => a.severity === "low").length}
          </p>
          <p className="text-sm text-gray-400">Low</p>
        </div>
      </div>

      {/* Search and Filters */}
      <div className="flex flex-col md:flex-row gap-4">
        <div className="relative flex-1">
          <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 w-4 h-4 text-gray-500" />
          <Input
            placeholder="Search alerts..."
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
            className="pl-10 bg-gray-900 border-gray-800 text-white placeholder-gray-500"
          />
        </div>
      </div>

      <AlertFilters activeFilter={activeFilter} setActiveFilter={setActiveFilter} />

      {/* Alerts List */}
      <div className="space-y-4">
        {loading ? (
          <div className="text-center py-12 text-gray-400">Loading alerts...</div>
        ) : filteredAlerts.length === 0 ? (
          <div className="text-center py-12 text-gray-400">No alerts found</div>
        ) : (
          filteredAlerts.map((alert) => (
            <AlertCard key={alert.id} alert={alert} />
          ))
        )}
      </div>
    </div>
  );
}
