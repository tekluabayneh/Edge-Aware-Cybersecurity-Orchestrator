import { useState, useEffect } from "react";
import { motion } from "framer-motion";
import { Home, ArrowLeft } from "lucide-react";
import { Link } from "react-router-dom";

const FloatingOrb = ({ delay, size, x, y, color }) => (
  <motion.div
    className="absolute rounded-full blur-3xl opacity-20"
    style={{
      width: size,
      height: size,
      left: x,
      top: y,
      background: color,
    }}
    animate={{
      y: [0, -30, 0, 20, 0],
      x: [0, 15, -10, 5, 0],
      scale: [1, 1.1, 0.95, 1.05, 1],
    }}
    transition={{
      duration: 8,
      delay,
      repeat: Infinity,
      ease: "easeInOut",
    }}
  />
);

export default function PageNotFound() {
  const [mousePos, setMousePos] = useState({ x: 0, y: 0 });

  useEffect(() => {
    const handleMouseMove = (e) => {
      setMousePos({
        x: (e.clientX / window.innerWidth - 0.5) * 20,
        y: (e.clientY / window.innerHeight - 0.5) * 20,
      });
    };
    window.addEventListener("mousemove", handleMouseMove);
    return () => window.removeEventListener("mousemove", handleMouseMove);
  }, []);

  return (
    <div className="min-h-screen bg-[#fafafa] relative overflow-hidden flex items-center justify-center px-6">
      {/* Floating background orbs */}
      <FloatingOrb delay={0} size={400} x="10%" y="20%" color="linear-gradient(135deg, #e0e7ff, #c7d2fe)" />
      <FloatingOrb delay={1.5} size={300} x="65%" y="10%" color="linear-gradient(135deg, #fce7f3, #fbcfe8)" />
      <FloatingOrb delay={3} size={350} x="50%" y="60%" color="linear-gradient(135deg, #dbeafe, #bfdbfe)" />
      <FloatingOrb delay={2} size={250} x="80%" y="70%" color="linear-gradient(135deg, #f3e8ff, #e9d5ff)" />

      {/* Subtle grid */}
      <div
        className="absolute inset-0 opacity-[0.03]"
        style={{
          backgroundImage: `radial-gradient(circle, #000 1px, transparent 1px)`,
          backgroundSize: "40px 40px",
        }}
      />

      <motion.div
        className="relative z-10 text-center max-w-lg"
        style={{
          x: mousePos.x * 0.3,
          y: mousePos.y * 0.3,
        }}
      >
        {/* Big 404 number */}
        <motion.div
          initial={{ opacity: 0, y: 40 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.8, ease: [0.22, 1, 0.36, 1] }}
        >
          <h1
            className="text-[10rem] sm:text-[13rem] font-black leading-none tracking-tighter select-none"
            style={{
              background: "linear-gradient(135deg, #1e1b4b 0%, #6366f1 40%, #a78bfa 70%, #c4b5fd 100%)",
              WebkitBackgroundClip: "text",
              WebkitTextFillColor: "transparent",
              filter: "drop-shadow(0 4px 30px rgba(99, 102, 241, 0.15))",
            }}
          >
            404
          </h1>
        </motion.div>

        {/* Divider line */}
        <motion.div
          className="mx-auto h-[2px] bg-gradient-to-r from-transparent via-indigo-300 to-transparent"
          initial={{ width: 0, opacity: 0 }}
          animate={{ width: "60%", opacity: 1 }}
          transition={{ duration: 1, delay: 0.3, ease: [0.22, 1, 0.36, 1] }}
        />

        {/* Text */}
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.7, delay: 0.5, ease: [0.22, 1, 0.36, 1] }}
          className="mt-8 space-y-3"
        >
          <h2 className="text-2xl sm:text-3xl font-semibold text-slate-800 tracking-tight">
            Page not found
          </h2>
          <p className="text-slate-500 text-base sm:text-lg leading-relaxed max-w-md mx-auto">
            The page you're looking for doesn't exist or has been moved.
            Let's get you back on track.
          </p>
        </motion.div>

        {/* Buttons */}
        <motion.div
          className="mt-10 flex flex-col sm:flex-row items-center justify-center gap-3"
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.7, delay: 0.7, ease: [0.22, 1, 0.36, 1] }}
        >
          <button
            className="bg-slate-900 hover:bg-slate-800 flex align-center pt-3 justify-center text-white rounded-full px-8 h-12 text-sm font-medium shadow-lg shadow-slate-900/10 transition-all duration-300 hover:shadow-xl hover:shadow-slate-900/20 hover:-translate-y-0.5"
          >
              <Home className="mr-2 h-4 w-4" />
            <Link to={"/dashboard"}>
              Back to Home
            </Link>
          </button>
          <button
            className="text-slate-600 hover:text-slate-900 pt-3 cursor-pointer flex align-center rounded-full px-8 h-12 text-sm font-medium transition-all duration-300 hover:-translate-y-0.5"
            onClick={() => window.history.back()}
          >
            <ArrowLeft className="mr-2 h-4 w-4 pt-1" />
            Go Back
          </button>
        </motion.div>

        {/* Decorative sparkle dots */}
        <motion.div
          className="mt-16 flex items-center justify-center gap-2"
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          transition={{ delay: 1, duration: 1 }}
        >
          {[0, 1, 2].map((i) => (
            <motion.div
              key={i}
              className="h-1.5 w-1.5 rounded-full bg-indigo-300"
              animate={{ opacity: [0.3, 1, 0.3], scale: [0.8, 1.2, 0.8] }}
              transition={{ duration: 2, delay: i * 0.3, repeat: Infinity, ease: "easeInOut" }}
            />
          ))}
        </motion.div>
      </motion.div>
    </div>
  );
}