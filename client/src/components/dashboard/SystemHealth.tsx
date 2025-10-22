import { healthMetrics } from '../../data/filters'

export default function SystemHealth() {
  return (
    <div className="bg-gray-900 border border-gray-800 rounded-xl p-6">
      <h3 className="text-lg font-semibold text-white mb-6">System Health</h3>
      <div className="grid grid-cols-2 gap-4">
        {healthMetrics.map((metric) => (
          <div key={metric.name} className="space-y-2">
            <div className="flex items-center justify-between">
              <span className="text-sm text-gray-400 flex items-center gap-2">
                <metric.icon className="w-4 h-4 text-cyan-400" />
                {metric.name}
              </span>
              <span className="text-sm font-semibold text-white">
                {metric.value}%
              </span>
            </div>
            <div className="h-2 bg-gray-800 rounded-full overflow-hidden">
              <div
                className={`h-full bg-gradient-to-r from-${metric.color}-500 to-${metric.color}-600 transition-all duration-1000 ease-out`}
                style={{ width: `${metric.value}%` }}
              ></div>
            </div>
          </div>
        ))}
      </div>
    </div>
  )
}
