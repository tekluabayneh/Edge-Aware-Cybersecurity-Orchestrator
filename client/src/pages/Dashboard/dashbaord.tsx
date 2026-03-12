import { useState, useEffect } from 'react'
import type { AlertType } from '../../types/Alert'


import { Shield, Activity, Eye, AlertCircle, AlertTriangleIcon } from 'lucide-react'
import MetricCard from '../../components/dashboard/MatricCard'
import SecurityScore from '../../components/dashboard/Security'
import RecentAlerts from '../../components/dashboard/RecentAlert'
import SystemHealth from '../../components/dashboard/SystemHealth'
import GetAllAlerts from '../../hooks/fetchAlerts'
import { useQuery } from '@tanstack/react-query'


export default function Dashboard() {
    const [alerts, setAlerts] = useState<AlertType[]>([])
    const params = new URLSearchParams(window.location.search);
    const token = params.get("token");
    const isAuthLogin = Boolean(token)

    if (isAuthLogin && token != null) {
        cookieStore.set("token", token)
        window.history.replaceState({}, document.title, window.location.pathname);
    }

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


    const activeAlerts = alerts.filter((a) => a.status === 'active').length
    const criticalAlerts = alerts.filter((a) => a.risk_level === 'critical').length
    const analyzedEvenCount = alerts.length
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
            <div className="flex flex-row justify-between p-2 align-center  gap-2">
                <div className="flex flex-col gap-2">

                    <h1 className="text-3xl font-bold text-white">Security Dashboard</h1>
                    <p className="text-gray-400">
                        Real-time overview of your security posture
                    </p>
                </div>
                <p className='cursor-pointer'>{AlertTriangleIcon}</p>
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
                    value={analyzedEvenCount}
                    change="+8%"
                    trend="up"
                    icon={Eye}
                />
            </div>

            {/* Main Content Grid */}
            <div className="grid lg:grid-cols-3 gap-6">
                <div className="lg:col-span-2 space-y-6">
                    <RecentAlerts />
                    <SystemHealth />
                </div>
                <div>
                    <SecurityScore />
                </div>
            </div>
        </div>
    )
}
