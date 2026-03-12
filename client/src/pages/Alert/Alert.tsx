import { useState, useEffect } from 'react'
import { AlertTriangle, Search } from 'lucide-react'
import AlertFilters from '../../components/alert/AlertFilter'
import AlertCard from '../../components/alert/AlertCard'
import type { AlertType } from '../../types/Alert'
import GetAllAlerts from '../../hooks/fetchAlerts'
import { useQuery } from '@tanstack/react-query'
import NotificationBell from '../../components/Notification/Notification'

export default function Alerts() {
    const [filteredAlerts, setFilteredAlerts] = useState<AlertType[]>([])
    const [activeFilter, setActiveFilter] = useState<string>('all')
    const [searchTerm, setSearchTerm] = useState('')

    const { data, isLoading, error } = useQuery({
        queryKey: ["userAlert"],
        queryFn: GetAllAlerts
    });

    useEffect(() => {
        filterAlerts()
    }, [data, activeFilter, searchTerm])

    const filterAlerts = () => {
        if (!data?.alert) {
            setFilteredAlerts([])
            return
        }

        let filtered = [...data.alert]

        if (activeFilter !== 'all') {
            filtered = filtered.filter((alert) => alert.risk_level === activeFilter)
        }

        if (searchTerm) {
            const search = searchTerm.toLowerCase()
            filtered = filtered.filter(
                (alert) =>
                    alert.title?.toLowerCase().includes(search) ||
                    alert.description?.toLowerCase().includes(search)
            )
        }

        setFilteredAlerts(filtered)
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
        <div className="space-y-8">
            {/* Header */}
            <div className="flex flex-row justify-between align-center p-2 gap-2">
                <div className="flex flex-col gap-2">
                    <h1 className="text-3xl font-bold text-white flex items-center gap-3">
                        <AlertTriangle className="w-8 h-8 text-cyan-400" />
                        Security Alerts
                    </h1>
                    <p className="text-gray-400">
                        Monitor and manage security threats in real-time
                    </p>

                </div>
                {<NotificationBell />}
            </div>

            {/* Stats Banner */}
            <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
                <div className="bg-gradient-to-br from-red-500/10 to-rose-500/10 border border-red-500/30 rounded-lg p-4">
                    <p className="text-2xl font-bold text-red-400">
                        {data?.alert?.filter((a: any) => a.risk_level === 'critical').length || 0}
                    </p>
                    <p className="text-sm text-gray-400">Critical</p>
                </div>
                <div className="bg-gradient-to-br from-orange-500/10 to-amber-500/10 border border-orange-500/30 rounded-lg p-4">
                    <p className="text-2xl font-bold text-orange-400">
                        {data?.alert?.filter((a: any) => a.risk_level === 'high').length || 0}
                    </p>
                    <p className="text-sm text-gray-400">High</p>
                </div>
                <div className="bg-gradient-to-br from-yellow-500/10 to-yellow-600/10 border border-yellow-500/30 rounded-lg p-4">
                    <p className="text-2xl font-bold text-yellow-400">
                        {data?.alert?.filter((a: any) => a.risk_level === 'medium').length || 0}
                    </p>
                    <p className="text-sm text-gray-400">Medium</p>
                </div>
                <div className="bg-gradient-to-br from-blue-500/10 to-cyan-500/10 border border-blue-500/30 rounded-lg p-4">
                    <p className="text-2xl font-bold text-blue-400">
                        {data?.alert?.filter((a: any) => a.risk_level === 'low').length || 0}
                    </p>
                    <p className="text-sm text-gray-400">Low</p>
                </div>
            </div>

            {/* Search and Filters */}
            <div className="flex flex-col md:flex-row gap-4">
                <div className="relative flex-1 w-full">
                    <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 w-4 h-4 text-gray-500" />
                    <input
                        placeholder="Search alerts..."
                        value={searchTerm}
                        onChange={(e) => setSearchTerm(e.target.value)}
                        className="pl-10 bg-gray-900 border-gray-800 text-white placeholder-gray-500 rounded-lg w-full p-2"
                    />
                </div>
            </div>

            <AlertFilters setActiveFilter={setActiveFilter} />

            {/* Alerts List */}
            <div className="space-y-4">
                {filteredAlerts.length === 0 ? (
                    <div className="text-center py-12 text-gray-400">No alerts found</div>
                ) : (
                    filteredAlerts.map((alert, idx) => (
                        <AlertCard key={alert._id || alert.id || idx} alert={alert} />
                    ))
                )}
            </div>
        </div>
    )
}

