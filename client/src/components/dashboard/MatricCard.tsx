import { TrendingUp, TrendingDown } from "lucide-react";
import type { MetricCardProps } from "../../types/Alert";
export default function MetricCard({
  title,
  value,
  change,
  icon: Icon,
  trend,
}: MetricCardProps) {
  const isPositive = trend === "up";

  return (
    <div className="bg-gray-900 border border-gray-800 rounded-xl p-6 hover:border-cyan-500/30 transition-all duration-300 group">
      <div className="flex justify-between items-start mb-4">
        <div className="w-12 h-12 bg-gradient-to-br from-blue-500/20 to-cyan-500/20 rounded-lg flex items-center justify-center group-hover:scale-110 transition-transform duration-300">
          <Icon className="w-6 h-6 text-cyan-400" />
        </div>
        {change && (
          <div
            className={`flex items-center gap-1 text-xs font-medium ${
              isPositive ? "text-green-400" : "text-red-400"
            }`}
          >
            {isPositive ? (
              <TrendingUp className="w-3 h-3" />
            ) : (
              <TrendingDown className="w-3 h-3" />
            )}
            {change}
          </div>
        )}
      </div>
      <h3 className="text-gray-400 text-sm font-medium mb-1">{title}</h3>
      <p className="text-3xl font-bold text-white">{value}</p>
    </div>
  );
}
