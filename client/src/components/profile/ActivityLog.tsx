import React from "react";
import { Activity, Clock } from "lucide-react";
import { format } from "date-fns";

const recentActivities = [
  { action: "Security scan completed", timestamp: new Date(Date.now() - 3600000) },
  { action: "Alert dismissed: Low priority phishing attempt", timestamp: new Date(Date.now() - 7200000) },
  { action: "System settings updated", timestamp: new Date(Date.now() - 14400000) },
  { action: "2FA enabled", timestamp: new Date(Date.now() - 86400000) },
  { action: "Password changed", timestamp: new Date(Date.now() - 172800000) }
];

export default function ActivityLog() {
  return (
    <div className="bg-gray-900 border border-gray-800 rounded-xl p-6">
      <h3 className="text-lg font-semibold text-white mb-6 flex items-center gap-2">
        <Activity className="w-5 h-5 text-cyan-400" />
        Recent Activity
      </h3>
      <div className="space-y-4">
        {recentActivities.map((activity, index) => (
          <div key={index} className="flex items-start gap-3 p-3 bg-gray-800/50 rounded-lg hover:bg-gray-800 transition-colors">
            <div className="w-2 h-2 bg-cyan-400 rounded-full mt-2"></div>
            <div className="flex-1">
              <p className="text-sm text-white">{activity.action}</p>
              <p className="text-xs text-gray-500 flex items-center gap-1 mt-1">
                <Clock className="w-3 h-3" />
                {format(activity.timestamp, "MMM d, yyyy 'at' HH:mm")}
              </p>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
