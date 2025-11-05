import { motion } from "framer-motion"
import { Globe, Shield, Terminal, Zap } from "lucide-react"
const Preview = () => { 

    return  ( 
        <>
            {/* Right Content - Animated Dashboard Preview */}
            <motion.div
              initial={{ opacity: 0, scale: 0.9 }}
              animate={{ opacity: 1, scale: 1 }}
              transition={{ duration: 0.8, delay: 0.4 }}
              className="relative"
            >
              {/* Glow Effect */}
              <div className="absolute inset-0 bg-gradient-to-r from-cyan-500/30 to-purple-600/30 blur-3xl rounded-full"></div>

              <div className="relative bg-gradient-to-br from-gray-900/90 to-black/90 border border-gray-700/50 rounded-3xl p-8 backdrop-blur-2xl shadow-2xl">
                <div className="space-y-4">
                  {/* Status Card */}
                  <motion.div
                    whileHover={{ scale: 1.02 }}
                    className="flex items-center justify-between p-4 bg-gradient-to-r from-gray-800/80 to-gray-900/80 rounded-xl border border-cyan-500/30 backdrop-blur-xl"
                  >
                    <div className="flex items-center gap-3">
                      <motion.div
                        animate={{ rotate: 360 }}
                        transition={{
                          duration: 3,
                          repeat: Infinity,
                          ease: "linear",
                        }}
                        className="w-12 h-12 bg-gradient-to-br from-green-500/30 to-emerald-600/30 rounded-xl flex items-center justify-center"
                      >
                        <Shield className="w-6 h-6 text-green-400" />
                      </motion.div>
                      <div>
                        <p className="font-semibold text-white">
                          All Systems Protected
                        </p>
                        <p className="text-sm text-gray-400">
                          Real-time monitoring active
                        </p>
                      </div>
                    </div>
                    <motion.div
                      animate={{ scale: [1, 1.2, 1] }}
                      transition={{ duration: 2, repeat: Infinity }}
                      className="w-3 h-3 bg-green-400 rounded-full"
                    ></motion.div>
                  </motion.div>

                  {/* Stats Grid */}
                  <div className="grid grid-cols-2 gap-4">
                    {[
                      {
                        icon: Globe,
                        value: "99.9%",
                        label: "Uptime",
                        color: "from-blue-500 to-cyan-500",
                      },
                      {
                        icon: Zap,
                        value: "<3ms",
                        label: "Response",
                        color: "from-yellow-500 to-orange-500",
                      },
                    ].map((stat, i) => (
                      <motion.div
                        key={i}
                        whileHover={{ scale: 1.05, rotate: 2 }}
                        className="p-4 bg-gradient-to-br from-gray-800/50 to-gray-900/50 rounded-xl border border-gray-700/50 backdrop-blur-xl"
                      >
                        <stat.icon
                          className={`w-8 h-8 mb-3 bg-gradient-to-r ${stat.color} bg-clip-text text-transparent`}
                        />
                        <p className="text-3xl font-bold text-white">
                          {stat.value}
                        </p>
                        <p className="text-sm text-gray-400 mt-1">
                          {stat.label}
                        </p>
                      </motion.div>
                    ))}
                  </div>

                  {/* Terminal Preview */}
                  <motion.div
                    whileHover={{ scale: 1.02 }}
                    className="p-4 bg-gradient-to-br from-cyan-500/10 to-blue-600/10 rounded-xl border border-cyan-500/30 backdrop-blur-xl"
                  >
                    <div className="flex items-center gap-2 mb-3">
                      <Terminal className="w-4 h-4 text-cyan-400" />
                      <span className="text-sm text-cyan-300 font-mono">
                        edgeguard@terminal
                      </span>
                    </div>
                    <p className="font-mono text-sm text-gray-200 mb-1">
                      $ edgeguard status
                    </p>
                    <motion.p
                      initial={{ opacity: 0 }}
                      animate={{ opacity: 1 }}
                      transition={{ delay: 1.5 }}
                      className="font-mono text-sm text-green-400"
                    >
                      ✓ Protection active • 3 devices secured
                    </motion.p>
                  </motion.div>

                  {/* Live Activity Indicator */}
                  <div className="flex items-center justify-center gap-2 pt-2">
                    <motion.div
                      animate={{ scale: [1, 1.5, 1], opacity: [1, 0.5, 1] }}
                      transition={{ duration: 2, repeat: Infinity }}
                      className="w-2 h-2 bg-cyan-400 rounded-full"
                    ></motion.div>
                    <span className="text-xs text-gray-400">
                      Live monitoring in progress
                    </span>
                  </div>
                </div>
              </div>
            </motion.div>
            
        </>
    )
}
export default Preview