import React, { useEffect, useState } from "react";
import {
  Shield,
  Zap,
  Lock,
  Globe,
  ChevronRight,
  Check,
  Terminal,
  Star,
  Sparkles,
  Download,
} from "lucide-react";
import { motion } from "framer-motion";
import Navigation from "../../components/navigation/navigation";

export default function Home() {
  const [scrollY, setScrollY] = useState(0);
  const [mousePosition, setMousePosition] = useState({ x: 0, y: 0 });

  useEffect(() => {
    const handleScroll = () => setScrollY(window.scrollY);
    const handleMouseMove = (e) => {
      setMousePosition({ x: e.clientX, y: e.clientY });
    };

    window.addEventListener("scroll", handleScroll);
    window.addEventListener("mousemove", handleMouseMove);

    return () => {
      window.removeEventListener("scroll", handleScroll);
      window.removeEventListener("mousemove", handleMouseMove);
    };
  }, []);

  return (
    <div className="min-h-screen bg-black text-white overflow-x-hidden relative">
      {/* Cursor Glow Effect */}
      <div
        className="pointer-events-none fixed inset-0 z-30 transition duration-300"
        style={{
          background: `radial-gradient(600px at ${mousePosition.x}px ${mousePosition.y}px, rgba(0, 212, 255, 0.15), transparent 80%)`,
        }}
      />
      <Navigation />
      {/* Hero Section - GOD MODE */}
      <section className="relative min-h-screen flex items-center justify-center pt-20 pb-12 px-4 sm:px-6">
        {/* Animated Background Layers */}
        <div className="absolute inset-0 overflow-hidden">
          {/* Main Gradient */}
          <div className="absolute inset-0 bg-gradient-to-br from-black via-gray-900 to-black"></div>

          {/* Animated Orbs */}
          <motion.div
            animate={{
              scale: [1, 1.2, 1],
              rotate: [0, 180, 360],
            }}
            transition={{ duration: 20, repeat: Infinity, ease: "linear" }}
            className="absolute top-1/4 left-1/4 w-[500px] h-[500px] bg-gradient-to-r from-cyan-500/30 to-blue-600/30 rounded-full blur-[150px]"
          />
          <motion.div
            animate={{
              scale: [1.2, 1, 1.2],
              rotate: [360, 180, 0],
            }}
            transition={{ duration: 25, repeat: Infinity, ease: "linear" }}
            className="absolute bottom-1/4 right-1/4 w-[600px] h-[600px] bg-gradient-to-l from-purple-600/30 to-pink-500/30 rounded-full blur-[150px]"
          />

          {/* Grid Pattern with Animation */}
          <div
            className="absolute inset-0 bg-[linear-gradient(to_right,#1f1f1f_1px,transparent_1px),linear-gradient(to_bottom,#1f1f1f_1px,transparent_1px)] bg-[size:4rem_4rem] opacity-20"
            style={{ transform: `translateY(${scrollY * 0.1}px)` }}
          />

          {/* Floating Particles */}
          <div className="absolute inset-0">
            {[...Array(30)].map((_, i) => (
              <motion.div
                key={i}
                className="absolute rounded-full"
                style={{
                  width: Math.random() * 4 + 1 + "px",
                  height: Math.random() * 4 + 1 + "px",
                  left: `${Math.random() * 100}%`,
                  top: `${Math.random() * 100}%`,
                  background: `rgba(0, 212, 255, ${Math.random() * 0.5 + 0.2})`,
                }}
                animate={{
                  y: [0, -100, 0],
                  opacity: [0, 1, 0],
                }}
                transition={{
                  duration: Math.random() * 3 + 2,
                  repeat: Infinity,
                  delay: Math.random() * 2,
                }}
              />
            ))}
          </div>
        </div>

        <div className="relative z-10 max-w-7xl mx-auto w-full">
          <div className="grid lg:grid-cols-2 gap-12 items-center">
            {/* Left Content */}
            <div className="space-y-8 text-center lg:text-left">
              <motion.div
                initial={{ opacity: 0, y: 30 }}
                animate={{ opacity: 1, y: 0 }}
                transition={{ duration: 0.8 }}
              >
                <motion.div
                  whileHover={{ scale: 1.05 }}
                  className="inline-flex items-center gap-2 px-4 py-2 bg-gradient-to-r from-cyan-500/20 to-blue-600/20 border border-cyan-500/50 rounded-full mb-8 backdrop-blur-xl"
                >
                  <Sparkles className="w-4 h-4 text-cyan-400 animate-pulse" />
                  <span className="text-cyan-300 text-sm font-medium">
                    Next-Gen Edge Security
                  </span>
                  <div className="w-2 h-2 bg-green-400 rounded-full animate-pulse"></div>
                </motion.div>

                <h1 className="text-5xl sm:text-6xl lg:text-7xl xl:text-8xl font-bold leading-tight mb-6">
                  <motion.span
                    initial={{ opacity: 0, x: -20 }}
                    animate={{ opacity: 1, x: 0 }}
                    transition={{ delay: 0.2 }}
                    className="block text-white"
                  >
                    Secure Your
                  </motion.span>
                  <motion.span
                    initial={{ opacity: 0, x: -20 }}
                    animate={{ opacity: 1, x: 0 }}
                    transition={{ delay: 0.4 }}
                    className="block bg-gradient-to-r from-cyan-400 via-blue-500 to-purple-600 bg-clip-text text-transparent mt-2"
                  >
                    Digital Empire
                  </motion.span>
                </h1>

                <motion.p
                  initial={{ opacity: 0 }}
                  animate={{ opacity: 1 }}
                  transition={{ delay: 0.6 }}
                  className="text-xl text-gray-300 leading-relaxed max-w-2xl"
                >
                  Revolutionary edge-aware protection powered. Install once,
                  protect forever. Experience cybersecurity that adapts to
                  threats before they happen.
                </motion.p>
              </motion.div>

              <motion.div
                initial={{ opacity: 0, y: 20 }}
                animate={{ opacity: 1, y: 0 }}
                transition={{ delay: 0.8 }}
                className="flex flex-col sm:flex-row gap-4 justify-center lg:justify-start"
              >
                <motion.div
                  whileHover={{ scale: 1.05 }}
                  whileTap={{ scale: 0.95 }}
                >
                  <button className="bg-gradient-to-r flex items-center cursor-pointer from-cyan-500 via-blue-600 to-purple-600 hover:from-cyan-600 hover:via-blue-700 hover:to-purple-700 text-lg px-10 py-7 rounded-xl shadow-2xl shadow-cyan-500/50 border-0 group">
                    <Zap className="w-5 h-5 mr-2 group-hover:animate-pulse" />
                    <a href="/auth">
                    Start Protection
                    </a>
                    <ChevronRight className="w-5 h-5 ml-2 group-hover:translate-x-1 transition-transform" />
                  </button>
                </motion.div>
              </motion.div>

              {/* Trust Indicators */}
              <motion.div
                initial={{ opacity: 0 }}
                animate={{ opacity: 1 }}
                transition={{ delay: 1 }}
                className="flex flex-wrap items-center justify-center lg:justify-start gap-6 pt-6"
              >
                {[
                  { icon: Check, text: "Zero Configuration" },
                  { icon: Zap, text: "Real-time Protection" },
                  { icon: Star, text: "5-Star Rated" },
                ].map((item, i) => (
                  <motion.div
                    key={i}
                    whileHover={{ scale: 1.1 }}
                    className="flex items-center gap-2"
                  >
                    <div className="w-8 h-8 bg-gradient-to-br from-cyan-500/30 to-blue-600/30 rounded-lg flex items-center justify-center">
                      <item.icon className="w-4 h-4 text-cyan-400" />
                    </div>
                    <span className="text-gray-300 text-sm">{item.text}</span>
                  </motion.div>
                ))}
              </motion.div>
            </div>

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
          </div>

          {/* Scroll Indicator */}
          <motion.div
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            transition={{ delay: 1.5 }}
            className="absolute bottom-8 left-1/2 transform -translate-x-1/2"
          >
            <motion.div
              animate={{ y: [0, 10, 0] }}
              transition={{ duration: 1.5, repeat: Infinity }}
              className="flex flex-col items-center gap-2"
            >
              <span className="text-xs text-gray-500 uppercase tracking-wider">
                Scroll to explore
              </span>
              <div className="w-6 h-10 border-2 border-gray-600 rounded-full flex justify-center">
                <motion.div
                  animate={{ y: [0, 12, 0] }}
                  transition={{ duration: 1.5, repeat: Infinity }}
                  className="w-1 h-3 bg-cyan-400 rounded-full mt-2"
                />
              </div>
            </motion.div>
          </motion.div>
        </div>
      </section>
      {/* Features Section */}
      <section id="features" className="relative py-20 px-4 sm:px-6">
        <div className="absolute inset-0 bg-gradient-to-b from-black via-gray-950 to-black"></div>

        <div className="relative z-10 max-w-7xl mx-auto">
          <motion.div
            initial={{ opacity: 0, y: 30 }}
            whileInView={{ opacity: 1, y: 0 }}
            viewport={{ once: true }}
            className="text-center mb-16"
          >
            <div className="inline-flex items-center gap-2 px-4 py-2 bg-cyan-500/10 border border-cyan-500/30 rounded-full mb-6">
              <Star className="w-4 h-4 text-cyan-400" />
              <span className="text-cyan-300 text-sm">Premium Features</span>
            </div>
            <h2 className="text-4xl sm:text-5xl font-bold mb-4 text-white">
              Military-Grade Protection
            </h2>
            <p className="text-xl text-gray-400 max-w-2xl mx-auto">
              Enterprise security for everyone
            </p>
          </motion.div>

          <div className="grid sm:grid-cols-2 lg:grid-cols-3 gap-6">
            {[
              {
                icon: Shield,
                title: "Edge Detection",
                desc: "AI-powered threat detection at the edge. Stop attacks before they reach your network.",
                color: "from-cyan-500 to-blue-600",
                delay: 0.1,
              },
              {
                icon: Zap,
                title: "Lightning Fast",
                desc: "Optimized Go-based agent with minimal latency. Protection that doesn't slow you down.",
                color: "from-yellow-500 to-orange-600",
                delay: 0.2,
              },
              {
                icon: Lock,
                title: "Zero Trust",
                desc: "Military-grade encryption. Every connection verified. Trust nothing, verify everything.",
                color: "from-purple-500 to-pink-600",
                delay: 0.3,
              },
            ].map((feature, i) => (
              <motion.div
                key={i}
                initial={{ opacity: 0, y: 30 }}
                whileInView={{ opacity: 1, y: 0 }}
                whileHover={{ scale: 1.05, rotate: 1 }}
                viewport={{ once: true }}
                transition={{ delay: feature.delay }}
              >
                <div className="bg-gradient-to-br from-gray-900/80 to-black/80 border-gray-700/50 p-8 h-full backdrop-blur-xl hover:border-cyan-500/50 transition-all group">
                  <motion.div
                    whileHover={{ rotate: 360 }}
                    transition={{ duration: 0.6 }}
                    className={`w-14 h-14 bg-gradient-to-br ${feature.color} rounded-xl flex items-center justify-center mb-6 opacity-80 group-hover:opacity-100`}
                  >
                    <feature.icon className="w-7 h-7 text-white" />
                  </motion.div>
                  <h3 className="text-2xl font-bold mb-4 text-white">
                    {feature.title}
                  </h3>
                  <p className="text-gray-400 leading-relaxed">
                    {feature.desc}
                  </p>
                </div>
              </motion.div>
            ))}
          </div>
        </div>
      </section>
      {/* How It Works - Simplified */}
      <section id="how-it-works" className="relative py-20 px-4 sm:px-6">
        <div className="relative z-10 max-w-7xl mx-auto">
          <motion.div
            initial={{ opacity: 0, y: 30 }}
            whileInView={{ opacity: 1, y: 0 }}
            viewport={{ once: true }}
            className="text-center mb-16"
          >
            <h2 className="text-4xl sm:text-5xl font-bold mb-4 text-white">
              Get Started in 3 Steps
            </h2>
            <p className="text-xl text-gray-400">
              Deploy in minutes, protect forever
            </p>
          </motion.div>

          <div className="grid md:grid-cols-3 gap-8">
            {[
              {
                num: "01",
                title: "Download",
                desc: "Get our lightweight agent",
                icon: Download,
              },
              {
                num: "02",
                title: "Install",
                desc: "One command installation",
                icon: Terminal,
              },
              {
                num: "03",
                title: "Protected",
                desc: "Real-time monitoring active",
                icon: Shield,
              },
            ].map((step, i) => (
              <motion.div
                key={i}
                initial={{ opacity: 0, y: 30 }}
                whileInView={{ opacity: 1, y: 0 }}
                whileHover={{ scale: 1.05 }}
                viewport={{ once: true }}
                transition={{ delay: i * 0.2 }}
                className="relative"
              >
                <div className="bg-gray-900/80 border border-gray-700/50 rounded-2xl p-8 backdrop-blur-xl">
                  <div className="text-7xl font-bold text-transparent bg-gradient-to-br from-cyan-500/20 to-blue-600/20 bg-clip-text mb-6">
                    {step.num}
                  </div>
                  <div className="w-16 h-16 bg-gradient-to-br from-cyan-500/30 to-blue-600/30 rounded-xl flex items-center justify-center mb-6 mx-auto">
                    <step.icon className="w-8 h-8 text-cyan-400" />
                  </div>
                  <h3 className="text-2xl font-bold mb-3 text-white text-center">
                    {step.title}
                  </h3>
                  <p className="text-gray-400 text-center">{step.desc}</p>
                </div>
              </motion.div>
            ))}
          </div>
        </div>
      </section>
      {/* CTA Section */}
      <section className="relative py-20 px-4 sm:px-6">
        <div className="relative z-10 max-w-4xl mx-auto">
          <motion.div
            initial={{ opacity: 0, scale: 0.9 }}
            whileInView={{ opacity: 1, scale: 1 }}
            viewport={{ once: true }}
            className="relative overflow-hidden bg-gradient-to-br from-cyan-500/10 via-blue-600/10 to-purple-600/10 border border-cyan-500/30 rounded-3xl p-12 text-center backdrop-blur-2xl"
          >
            <div className="absolute inset-0 bg-gradient-to-r from-cyan-500/5 to-purple-600/5"></div>
            <div className="relative z-10">
              <motion.div
                animate={{ rotate: 360 }}
                transition={{ duration: 20, repeat: Infinity, ease: "linear" }}
                className="w-20 h-20 mx-auto mb-6 bg-gradient-to-br from-cyan-500 to-purple-600 rounded-2xl flex items-center justify-center"
              >
                <Shield className="w-10 h-10 text-white" />
              </motion.div>
              <h2 className="text-4xl sm:text-5xl font-bold mb-6 text-white">
                Ready for God-Mode Security?
              </h2>
              <p className="text-xl text-gray-300 mb-8 max-w-2xl mx-auto">
                Join the elite. Protect your digital empire with next-generation
                cybersecurity.
              </p>
              <motion.div
                whileHover={{ scale: 1.05 }}
                whileTap={{ scale: 0.95 }}
              >
                <button className="bg-gradient-to-r flex items-center cursor-pointer mx-auto from-cyan-500 via-blue-600 to-purple-600 hover:from-cyan-600 hover:via-blue-700 hover:to-purple-700 text-lg px-12 py-7 rounded-xl shadow-2xl shadow-cyan-500/50">
                  <Sparkles className="w-5 h-5 mr-2" />
                    <a href="/auth">
                  Activate Protection
                  </a>
                  <ChevronRight className="w-5 h-5 ml-2" />
                </button>
              </motion.div>
            </div>
          </motion.div>
        </div>
      </section>
    </div>
  );
}
