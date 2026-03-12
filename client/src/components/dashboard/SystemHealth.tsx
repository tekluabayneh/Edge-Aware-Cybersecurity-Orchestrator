import { useQuery } from '@tanstack/react-query'
// import { healthMetrics } from '../../data/filters'
import { FechUserAgentSystem } from '../../hooks/fetchSystemStats'
import { useEffect, useState } from 'react';
import { Server, Cpu, HardDrive, Network, type LucideIcon } from "lucide-react";
type SysInfoType = {
    icon: LucideIcon,
    color: string,
    value: string | number,
    name: string,
}

export default function SystemHealth() {
    const [systemInfo, setSystemInfo] = useState<SysInfoType[] | []>([])

    const { data, isLoading, error } = useQuery({
        queryKey: ["systemAlert"],
        queryFn: FechUserAgentSystem
    });
    useEffect(() => {
        const iconMap: any = {
            cpu: { icon: Cpu, bar: "from-cyan-500 to-cyan-600", iconColor: "text-cyan-400" },
            ram: { icon: Server, bar: "from-blue-500 to-blue-600", iconColor: "text-blue-400" },
            disk: { icon: HardDrive, bar: "from-green-500 to-green-600", iconColor: "text-green-400" },
            network: { icon: Network, bar: "from-purple-500 to-purple-600", iconColor: "text-purple-400" },
        };

        const SysInfo = Object.entries(data?.data[0]?.system?.payload || {}).filter(([key,]) => iconMap[key])
            .map(
                ([key, value]) => ({
                    name: key,
                    value,
                    icon: iconMap[key]?.icon ?? Network,
                    bar: iconMap[key].bar,
                    color: iconMap[key]?.iconColor ?? "green",
                })
            );

        setSystemInfo(SysInfo);
    }, [data, setSystemInfo]);

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
            <h3 className="text-lg font-semibold text-white mb-6">System Health</h3>
            <div className="grid grid-cols-2 gap-4">
                {systemInfo.map((metric) => (
                    <div key={metric.name} className="space-y-2">
                        <div className="flex items-center justify-between">
                            <span className="text-sm text-gray-400 flex items-center gap-2">
                                <metric.icon className="w-4 h-4 text-cyan-400" />
                                {metric.name}
                            </span>
                            <span className="text-sm font-semibold text-white">
                                {Number(metric.value) ? Math.round(metric.value) : metric.value}%
                            </span>
                        </div>
                        <div className="h-2 bg-gray-800 rounded-full overflow-hidden">
                            <div
                                className={`h-full bg-gradient-to-r ${metric.bar} transition-all duration-1000 ease-out`}
                                style={{ width: `${metric.value}%` }}
                            ></div>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    )
}
