import { motion } from "framer-motion"
import {Shield,  Terminal,Download} from "lucide-react";
const HowItWork = () =>{ 
    return ( 
        <> 
         
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
      </>
    )
}

export default HowItWork