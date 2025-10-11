import { AlertTriangle, Clock } from "lucide-react";
import { format } from "date-fns";
import { Link } from "react-router-dom";

const severityColors = {
  critical: "from-red-500 to-rose-500",
  high: "from-orange-500 to-amber-500",
  medium: "from-yellow-500 to-yellow-600",
  low: "from-blue-500 to-cyan-500"
};

export default function RecentAlerts({ alerts }) {
  return (
    <div className="bg-gray-900 border border-gray-800 rounded-xl p-6">
      <div className="flex justify-between items-center mb-6">
        <h3 className="text-lg font-semibold text-white flex items-center gap-2">
          <AlertTriangle className="w-5 h-5 text-cyan-400" />
          Recent Alerts
        </h3>
        <Link to={"/Alerts"} className="text-sm text-cyan-400 hover:text-cyan-300 transition-colors">
          View All →
        </Link>
      </div>
      <div className="space-y-3">
        {alerts.slice(0, 5).map((alert) => (
          <div
            key={alert.id}
            className="p-4 bg-gray-800/50 border border-gray-700 rounded-lg hover:border-cyan-500/30 transition-all duration-200"
          >
            <div className="flex items-start justify-between mb-2">
              <div className="flex items-center gap-2">
                <div className={`w-2 h-2 rounded-full bg-gradient-to-r ${severityColors[alert.severity]} animate-pulse`}></div>
                <span className="text-sm font-medium text-white">{alert.title}</span>
              </div>
              <span className="text-xs text-gray-500 uppercase font-semibold">{alert.severity}</span>
            </div>
            <p className="text-xs text-gray-400 mb-2">{alert.description}</p>
            <div className="flex items-center gap-4 text-xs text-gray-500">
              <span className="flex items-center gap-1">
                <Clock className="w-3 h-3" />
                {format(new Date(alert.detected_at || alert.created_date), "MMM d, HH:mm")}
              </span>
              <span className="px-2 py-1 bg-gray-700/50 rounded">{alert.category}</span>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
