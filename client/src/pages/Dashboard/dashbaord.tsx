import React, { useState, useEffect } from "react";

//@ts-expect-error
import Alert from "../../Entities/Alert";

import { Shield, Activity, Eye, AlertCircle } from "lucide-react";
import MetricCard from "../../components/dashboard/MatricCard";
import SecurityScore from "../../components/dashboard/Security";
import RecentAlerts from "../../components/dashboard/RecentAlert";
import SystemHealth from "../../components/dashboard/SystemHealth";

export default function Dashboard() {
  const [alerts, setAlerts] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    loadData();
  }, []);

  const loadData = async () => {
    setLoading(true);
    const alertsData = await Alert.list("-detected_at");
    setAlerts(alertsData);
    setLoading(false);
  };

  const activeAlerts = alerts.filter((a) => a.status === "active").length;
  const criticalAlerts = alerts.filter((a) => a.severity === "critical").length;

  return (
    <div className="space-y-8">
      {/* Header */}
      <div className="flex flex-col gap-2">
        <h1 className="text-3xl font-bold text-white">Security Dashboard</h1>
        <p className="text-gray-400">
          Real-time overview of your security posture
        </p>
      </div>

      {/* Metrics Grid */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <MetricCard
          title="Active Threats"
          value={activeAlerts}
          change="-12%"
          trend="down"
          icon={AlertCircle}
        />
        <MetricCard
          title="Critical Alerts"
          value={criticalAlerts}
          change="+5%"
          trend="up"
          icon={Shield}
        />
        <MetricCard
          title="Systems Monitored"
          value="24"
          change="+2"
          trend="up"
          icon={Activity}
        />
        <MetricCard
          title="Events Analyzed"
          value="1.2M"
          change="+8%"
          trend="up"
          icon={Eye}
        />
      </div>

      {/* Main Content Grid */}
      <div className="grid lg:grid-cols-3 gap-6">
        <div className="lg:col-span-2 space-y-6">
          <RecentAlerts alerts={alerts} />
          <SystemHealth />
        </div>
        <div>
          <SecurityScore score={60} />
        </div>
      </div>
    </div>
  );
}
