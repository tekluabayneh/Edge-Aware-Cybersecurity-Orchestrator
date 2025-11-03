import { motion } from "framer-motion"
import { ChevronRight, Shield, Sparkles } from "lucide-react"
const CTA = () => { 

    return ( 
<>
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

      </>
    )
}
export default CTA