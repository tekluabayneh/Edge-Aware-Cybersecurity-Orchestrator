import { useQuery } from '@tanstack/react-query';
import { Shield } from 'lucide-react'
import { useEffect, useState } from 'react';
import { FechUserAgentSystem } from '../../hooks/fetchSystemStats';

export default function SecurityScore() {
    const [score, setscore] = useState<number>(0)
    const percentage = (score / 100) * 100

    const { data, isLoading, error } = useQuery({
        queryKey: ["systemAlert"],
        queryFn: FechUserAgentSystem
    });

    useEffect(() => {
        const antivirus = data?.data[0]?.security?.payload.antivirus?.running
        const firewall = data?.data[0]?.security?.payload.firewall?.enabled
        const network = data?.data[0]?.network?.is_up_and_running


        const checks = [antivirus, firewall, network];
        const trueCount = checks.filter(Boolean).length;
        let overAllScore = Math.round((trueCount / checks.length) * 100)
        setscore(overAllScore);
    }, [data, setscore]);

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
                <div className="text-gray-400">Opps! you don't have score yet</div>
            </div>
        )
    }

    if (error) return <p className="text-red-500">Something went wrong</p>;


    return (
        <div className="bg-gray-900 border border-gray-800 rounded-xl p-6">
            <h3 className="text-lg font-semibold text-white mb-6">Security Score</h3>
            <div className="flex flex-col items-center">
                <div className="relative w-48 h-48 mb-6">
                    <svg className="transform -rotate-90 w-48 h-48">
                        <circle
                            cx="96"
                            cy="96"
                            r="88"
                            stroke="currentColor"
                            strokeWidth="12"
                            fill="none"
                            className="text-gray-800"
                        />
                        <circle
                            cx="96"
                            cy="96"
                            r="88"
                            stroke="url(#gradient)"
                            strokeWidth="12"
                            fill="none"
                            strokeDasharray={`${percentage * 5.53} 553`}
                            strokeLinecap="round"
                            className="transition-all duration-1000 ease-out"
                        />
                        <defs>
                            <linearGradient id="gradient" x1="0%" y1="0%" x2="100%" y2="100%">
                                <stop
                                    offset="0%"
                                    stopColor={
                                        score >= 90
                                            ? '#22c55e'
                                            : score >= 75
                                                ? '#84cc16'
                                                : score >= 60
                                                    ? '#facc15'
                                                    : score >= 40
                                                        ? '#f97316'
                                                        : '#ef4444'
                                    }
                                />
                                <stop
                                    offset="100%"
                                    stopColor={
                                        score >= 90
                                            ? '#10b981'
                                            : score >= 75
                                                ? '#a3e635'
                                                : score >= 60
                                                    ? '#f59e0b'
                                                    : score >= 40
                                                        ? '#f43f5e'
                                                        : '#dc2626'
                                    }
                                />                            </linearGradient>
                        </defs>
                    </svg>
                    <div className="absolute inset-0 flex flex-col items-center justify-center">
                        <Shield className="w-8 h-8 text-cyan-400 mb-2" />
                        <p className="text-4xl font-bold text-white">{score}</p>
                        <p className="text-sm text-gray-400">/ 100</p>
                    </div>
                </div>
                <p className="text-gray-400 text-sm">
                    Your security posture is{' '}
                    <span className="text-cyan-400 font-semibold">{score >= 80 ? "Excellent" : score > 50 ? "Not Bad" : "Very Bad"}</span>
                </p>
            </div>
        </div>
    )
}
