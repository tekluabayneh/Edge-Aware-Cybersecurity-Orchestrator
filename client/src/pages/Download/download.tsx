import React, { useState } from "react";
import { motion } from "framer-motion";
import {  Download as DownloadIcon, AlertCircle, CheckCircle2, Bell } from "lucide-react";
import { downloadButtons } from "../../data/filters";

export default function Download() {
  const [hoveredOS, setHoveredOS] = useState(null);
  return (
    <div className="min-h-screen bg-black text-white overflow-hidden relative">
      {/* Animated background gradient */}
      <div className="absolute inset-0 bg-gradient-to-br from-blue-950/20 via-black to-purple-950/20" />
      <div className="absolute inset-0">
        <div className="absolute top-1/4 left-1/4 w-96 h-96 bg-blue-500/10 rounded-full blur-3xl animate-pulse" />
        <div className="absolute bottom-1/4 right-1/4 w-96 h-96 bg-purple-500/10 rounded-full blur-3xl animate-pulse" style={{ animationDelay: "1s" }} />
      </div>

      <div className="relative z-10 container mx-auto px-6 py-16 md:py-24 max-w-6xl">
        {/* Hero Section */}
        <motion.div
          initial={{ opacity: 0, y: 30 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.8 }}
          className="text-center mb-16 md:mb-24"
        >
          <h1 className="text-5xl md:text-7xl font-bold mb-6 bg-gradient-to-r from-blue-400 via-cyan-300 to-purple-400 bg-clip-text text-transparent leading-tight">
            Download Agent
          </h1>
          <p className="text-xl md:text-2xl text-gray-400 max-w-3xl mx-auto leading-relaxed">
            Your intelligent monitoring companion. Stay connected, get instant alerts, and never miss what matters.
          </p>
        </motion.div>

        {/* Prerequisites Card */}
        <motion.div
          initial={{ opacity: 0, y: 30 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.8, delay: 0.2 }}
          className="mb-16"
        >
          <div className="bg-gradient-to-br from-gray-900/50 to-gray-800/30 border-gray-800/50 backdrop-blur-xl">
            <div className="p-8 md:p-10">
              <div className="flex items-start gap-4 mb-6">
                <div className="p-3 rounded-xl bg-gradient-to-br from-cyan-500/20 to-blue-500/20 border border-cyan-500/30">
                  <AlertCircle className="w-6 h-6 text-cyan-400" />
                </div>
                <div>
                  <h3 className="text-2xl font-semibold mb-2 text-white">Before You Start</h3>
                  <p className="text-gray-400 text-lg">Please review these requirements</p>
                </div>
              </div>

              <div className="grid md:grid-cols-2 gap-6 mt-8">
                <div className="flex items-start gap-4 p-5 rounded-xl bg-black/40 border border-gray-800/50">
                  <CheckCircle2 className="w-5 h-5 text-green-400 mt-1 flex-shrink-0" />
                  <div>
                    <h4 className="font-semibold text-white mb-2 text-lg">64-bit System Required</h4>
                    <p className="text-gray-400 leading-relaxed">
                      Ensure your operating system is running on a 64-bit architecture for optimal performance
                    </p>
                  </div>
                </div>

                <div className="flex items-start gap-4 p-5 rounded-xl bg-black/40 border border-gray-800/50">
                  <Bell className="w-5 h-5 text-blue-400 mt-1 flex-shrink-0" />
                  <div>
                    <h4 className="font-semibold text-white mb-2 text-lg">Sign Up for Alerts</h4>
                    <p className="text-gray-400 leading-relaxed">
                      Register to receive real-time notifications via email and other methods you configure
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </motion.div>

        {/* Download Buttons Section */}
        <motion.div
          initial={{ opacity: 0, y: 30 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.8, delay: 0.4 }}
        >
          <h2 className="text-3xl md:text-4xl font-bold text-center mb-12 text-white">
            Choose Your Platform
          </h2>

          <div className="grid md:grid-cols-3 gap-6 md:gap-8">
            {downloadButtons.map((btn, index) => (
              <motion.a
                key={btn.os}
                href={btn.downloadUrl}
                initial={{ opacity: 0, y: 30 }}
                animate={{ opacity: 1, y: 0 }}
                transition={{ duration: 0.6, delay: 0.5 + index * 0.1 }}
                onMouseEnter={() => setHoveredOS(btn.os)}
                onMouseLeave={() => setHoveredOS(null)}
                className="block group"
              >
                <div className={`
                  relative overflow-hidden bg-gradient-to-br from-gray-900/80 to-gray-800/50 
                  border-gray-700/50 backdrop-blur-xl transition-all duration-500 
                  hover:scale-105 hover:shadow-2xl ${btn.hoverColor}
                  cursor-pointer
                `}>
                  <div className="p-8 md:p-10 flex flex-col items-center text-center relative z-10">
                    {/* Gradient overlay on hover */}
                    <div className={`
                      absolute inset-0 bg-gradient-to-br ${btn.color} opacity-0 
                      group-hover:opacity-10 transition-opacity duration-500
                    `} />

                    {/* Icon */}
                    <div className={`
                      relative w-20 h-20 md:w-24 md:h-24 mb-6 rounded-2xl 
                      bg-gradient-to-br ${btn.color} p-0.5
                      group-hover:scale-110 transition-transform duration-500
                    `}>
                      <div className="w-full h-full bg-black rounded-2xl flex items-center justify-center">
                        <btn.icon className="w-10 h-10 md:w-12 md:h-12 text-white" />
                      </div>
                    </div>

                    {/* OS Name */}
                    <h3 className="text-2xl md:text-3xl font-bold mb-4 text-white">
                      {btn.os}
                    </h3>

                    {/* Download Button */}
                    <button
                      className={`
                        w-full bg-gradient-to-r ${btn.color} text-white font-semibold 
                        py-5 text-lg rounded-xl border-0 shadow-lg
                        group-hover:shadow-2xl transition-all duration-500
                        flex items-center justify-center gap-2
                      `}
                    >
                      <DownloadIcon className={`
                        w-5 h-5 transition-transform duration-500
                        ${hoveredOS === btn.os ? "translate-y-1" : ""}
                      `} />
                      Download for {btn.os}
                    </button>

                    {/* File info */}
                    <p className="text-gray-500 text-sm mt-4">64-bit installer</p>
                  </div>
                </div>
              </motion.a>
            ))}
          </div>
        </motion.div>

        {/* Footer Note */}
        <motion.div
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          transition={{ duration: 0.8, delay: 0.8 }}
          className="text-center mt-16 md:mt-20"
        >
          <p className="text-gray-500 text-lg">
            Need help? Contact our support team anytime.
          </p>
        </motion.div>
      </div>
    </div>
  );
}