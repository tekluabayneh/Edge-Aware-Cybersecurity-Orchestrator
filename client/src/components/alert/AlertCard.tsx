import React from "react";
import { Clock, MapPin, Server } from "lucide-react";
import { format } from "date-fns";

const severityColors = {
  critical: "bg-red-500/20 text-red-400 border-red-500/50",
  high: "bg-orange-500/20 text-orange-400 border-orange-500/50",
  medium: "bg-yellow-500/20 text-yellow-400 border-yellow-500/50",
  low: "bg-blue-500/20 text-blue-400 border-blue-500/50",
};

const statusColors = {
  active: "bg-red-500/20 text-red-400 border-red-500/50",
  investigating: "bg-yellow-500/20 text-yellow-400 border-yellow-500/50",
  resolved: "bg-green-500/20 text-green-400 border-green-500/50",
  dismissed: "bg-gray-500/20 text-gray-400 border-gray-500/50",
};

const categoryIcons = {
  malware: "🦠",
  phishing: "🎣",
  unauthorized_access: "🔓",
  data_breach: "💾",
  ddos: "⚡",
  vulnerability: "🔍",
};

export default function AlertCard({ alert }) {
  return (
    <div className="bg-gray-900 border border-gray-800 rounded-xl p-6 hover:border-cyan-500/30 transition-all duration-300">
      <div className="flex flex-col md:flex-row md:items-start justify-between gap-4 mb-4">
        <div className="flex items-start gap-3">
          <div className="text-2xl">{categoryIcons[alert.category]}</div>
          <div className="flex-1">
            <h3 className="text-lg font-semibold text-white mb-1">
              {alert.title}
            </h3>
            <p className="text-sm text-gray-400">{alert.description}</p>
          </div>
        </div>
        <div className="flex gap-2">
          <div
            className={`${
              severityColors[alert.severity]
            } border  rounded-lg px-3 text-xs font-semibold`}
          >
            {alert.severity}
          </div>
          <div
            className={`${
              statusColors[alert.status]
            } border  uppercase text-xs rounded-lg px-3 font-semibold`}
          >
            {alert.status}
          </div>
        </div>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-3 gap-4 pt-4 border-t border-gray-800">
        <div className="flex items-center gap-2 text-sm text-gray-400">
          <Clock className="w-4 h-4 text-cyan-400" />
          <span>
            {format(
              new Date(alert.detected_at || alert.created_date),
              "MMM d, yyyy HH:mm"
            )}
          </span>
        </div>
        {alert.source_ip && (
          <div className="flex items-center gap-2 text-sm text-gray-400">
            <MapPin className="w-4 h-4 text-cyan-400" />
            <span>{alert.source_ip}</span>
          </div>
        )}
        {alert.target && (
          <div className="flex items-center gap-2 text-sm text-gray-400">
            <Server className="w-4 h-4 text-cyan-400" />
            <span>{alert.target}</span>
          </div>
        )}
      </div>
    </div>
  );
}
