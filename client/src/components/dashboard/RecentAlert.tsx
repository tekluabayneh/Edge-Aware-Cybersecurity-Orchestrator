import { AlertTriangle, Clock } from 'lucide-react'
import { Link } from 'react-router-dom'
import { severityColorsForAlert } from '../../constants/colors'
import type { AlertType } from '../../types/Alert'
import GetAllAlerts from '../../hooks/fetchAlerts'
import { useEffect, useState } from 'react'
import { useQuery } from '@tanstack/react-query'

export default function RecentAlerts() {
    const [alerts, setAlerts] = useState<AlertType[]>([])

    const { data, isLoading, error } = useQuery({
        queryKey: ["userAlert"],
        queryFn: GetAllAlerts
    });

    useEffect(() => {
        filterAlerts()
    }, [data, setAlerts])

    const filterAlerts = () => {
        if (!data?.alert) {
            setAlerts([])
            return
        }
        setAlerts(data?.alert)
    }

    if (isLoading) {
        return (
            <div className="flex items-center justify-center h-screen">
                <div className="text-gray-400">Loading alerts...</div>
            </div>
        )
    }

    if (!data) {
        return (
            <div className="flex items-center justify-center h-screen">
                <div className="text-gray-400">Opps! you don't have alerts yet</div>
            </div>
        )
    }

    if (error) return <p className="text-red-500">Something went wrong</p>;


    return (
        <div className="bg-gray-900 border border-gray-800 rounded-xl p-6">
            <div className="flex justify-between items-center mb-6">
                <h3 className="text-lg font-semibold text-white flex items-center gap-2">
                    <AlertTriangle className="w-5 h-5 text-cyan-400" />
                    Recent Alerts
                </h3>
                <Link
                    to={'/Alerts'}
                    className="text-sm text-cyan-400 hover:text-cyan-300 transition-colors"
                >
                    View All →
                </Link>
            </div>
            <div className="space-y-3">
                {alerts?.slice(0, 5).map((alert, idx) => (
                    <div
                        key={alert._id || idx}
                        className="p-4 bg-gray-800/50 border border-gray-700 rounded-lg hover:border-cyan-500/30 transition-all duration-200"
                    >
                        <div className="flex items-start justify-between mb-2">
                            <div className="flex items-center gap-2">
                                <div
                                    className={`w-2 h-2 rounded-full bg-gradient-to-r ${severityColorsForAlert[alert.risk_level]} animate-pulse`}
                                ></div>
                                <span className="text-sm font-medium text-white">
                                    {alert.summery}
                                </span>
                            </div>
                            <span className={`px-2 rounded-sm text-sm bg-gradient-to-r ${severityColorsForAlert[alert.risk_level]} `}>
                                {alert.risk_level}
                            </span>
                        </div>
                        <p className="text-xs text-gray-400 mb-2">{alert.message}</p>
                        <div className="flex items-center gap-4 text-xs text-gray-500">
                            <span className="flex items-center gap-1">
                                <Clock className="w-3 h-3" />
                                {alert.CreatedAt}
                            </span>
                            <span className="px-2 py-1 bg-gray-700/50 rounded">
                                {alert.alertType}
                            </span>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    )
}
