import React from "react";
import { Shield } from "lucide-react";

export default function SecurityScore({ score }) {
  const percentage = (score / 100) * 100;
  const getScoreColor = () => {
    if (score >= 80) return "from-green-500 to-emerald-500";
    if (score >= 60) return "from-yellow-500 to-orange-500";
    return "from-red-500 to-rose-500";
  };

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
                <stop offset="0%" className={`text-green-500`} stopColor="currentColor" />
                <stop offset="100%" className={`text-cyan-500`} stopColor="currentColor" />
              </linearGradient>
            </defs>
          </svg>
          <div className="absolute inset-0 flex flex-col items-center justify-center">
            <Shield className="w-8 h-8 text-cyan-400 mb-2" />
            <p className="text-4xl font-bold text-white">{score}</p>
            <p className="text-sm text-gray-400">/ 100</p>
          </div>
        </div>
        <p className="text-gray-400 text-sm">
          Your security posture is <span className="text-cyan-400 font-semibold">Excellent</span>
        </p>
      </div>
    </div>
  );
}
