import { motion } from "framer-motion";
import { ArrowRight, Shield } from "lucide-react";
const Navigation = () => {
  return (
    <>
      {/* Navigation */}
      <motion.nav
        initial={{ y: -100 }}
        animate={{ y: 0 }}
        className="fixed top-0 w-full z-50 bg-black/50 backdrop-blur-2xl border-b border-gray-800/50"
      >
        <div className="max-w-7xl mx-auto px-4 sm:px-6 py-4">
          <div className="flex justify-between items-center">
            <motion.div
              whileHover={{ scale: 1.05 }}
              className="flex items-center gap-2 cursor-pointer"
            >
              <div className="w-8 h-8 bg-gradient-to-br from-cyan-400 to-blue-600 rounded-lg flex items-center justify-center">
                <Shield className="w-5 h-5 text-white" />
              </div>
              <span className="text-xl sm:text-2xl font-bold bg-gradient-to-r from-cyan-400 via-blue-500 to-purple-600 bg-clip-text text-transparent">
                EdgeGuard
              </span>
            </motion.div>

            <div className="flex items-center gap-6">
              <a
                href="#features"
                className="hidden sm:inline text-gray-300 hover:text-cyan-400 transition-colors"
              >
                Features
              </a>
              <a
                href="#how-it-works"
                className="hidden sm:inline text-gray-300 hover:text-cyan-400 transition-colors"
              >
                How It Works
              </a>
              <motion.div
                whileHover={{ scale: 1.05 }}
                whileTap={{ scale: 0.95 }}
              >
                <button className="bg-gradient-to-r flex items-center cursor-pointer p-2 rounded-md from-cyan-500 via-blue-600 to-purple-600 hover:from-cyan-600 hover:via-blue-700 hover:to-purple-700 border-0 shadow-lg shadow-cyan-500/50">
                    <a href="/auth">
                  Get Started
                  </a>
                  <ArrowRight className="ml-2 w-4 h-4" />
                </button>
              </motion.div>
            </div>
          </div>
        </div>
      </motion.nav>
    </>
  );
};
export default Navigation;
