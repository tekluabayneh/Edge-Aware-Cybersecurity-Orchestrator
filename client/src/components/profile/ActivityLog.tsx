import { Activity, Clock } from 'lucide-react'
import GetAllAlerts from '../../hooks/fetchAlerts';
import { useQuery } from '@tanstack/react-query';


export default function ActivityLog() {
    const { data, isLoading, error } = useQuery({
        queryKey: ["userAlert"],
        queryFn: GetAllAlerts
    });

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
            <h3 className="text-lg font-semibold text-white mb-6 flex items-center gap-2">
                <Activity className="w-5 h-5 text-cyan-400" />
                Recent Activity
            </h3>
            <div className="space-y-4">
                {data["alert"]?.slice(0, 5).map((activity, index) => (
                    <div
                        key={index}
                        className="flex items-start gap-3 p-3 bg-gray-800/50 rounded-lg hover:bg-gray-800 transition-colors"
                    >
                        <div className="w-2 h-2 bg-cyan-400 rounded-full mt-2"></div>
                        <div className="flex-1">
                            <p className="text-sm text-white">{activity.message}</p>
                            <p className="text-xs text-gray-500 flex items-center gap-1 mt-1">
                                <Clock className="w-3 h-3" />
                                {activity.CreatedAt ? activity.CreatedAt : activity.created_at}
                            </p>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    )
}
