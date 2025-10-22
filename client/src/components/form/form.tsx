import { useState, type FormEvent } from 'react'
import { motion, AnimatePresence } from 'framer-motion'
import { ArrowRight, Lock, Mail, User } from 'lucide-react'

interface FormProps {
  isSignUp: boolean
}

const Form = ({ isSignUp }: FormProps) => {
  const [email, setEmail] = useState('')
  const [isLoading, setIsLoading] = useState(false)

  const handleAuth = async (e: FormEvent) => {
    e.preventDefault()
    setIsLoading(true)
  }

  const flipVariants = {
    initial: (isSignUp: boolean) => ({
      rotateY: isSignUp ? -180 : 180,
      opacity: 0,
      scale: 0.9,
    }),
    animate: {
      rotateY: 0,
      opacity: 1,
      scale: 1,
      transition: {
        duration: 0.6,
        ease: [0.45, 0, 0.55, 1],
      },
    },
    exit: (isSignUp: boolean) => ({
      rotateY: isSignUp ? 180 : -180,
      opacity: 0,
      scale: 0.9,
      transition: {
        duration: 0.5,
        ease: [0.45, 0, 0.55, 1],
      },
    }),
  }

  return (
    <div
      className="relative w-full"
      style={{
        perspective: 1000,
      }}
    >
      <AnimatePresence mode="wait" custom={isSignUp}>
        <motion.form
          key={isSignUp ? 'signup' : 'signin'}
          onSubmit={handleAuth}
          variants={flipVariants}
          initial="initial"
          animate="animate"
          exit="exit"
          custom={isSignUp}
          className="space-y-5 backface-hidden"
          style={{
            transformStyle: 'preserve-3d',
          }}
        >
          {isSignUp ? (
            <>
              {/* Username */}
              <div className="space-y-2">
                <label className="text-sm font-medium text-gray-300 flex items-center gap-2">
                  <User className="w-4 h-4" />
                  Username
                </label>
                <input
                  type="text"
                  placeholder="Enter your name"
                  className="bg-gray-800 w-full placeholder:pl-3 rounded-md border-gray-700 text-white placeholder:text-gray-500 h-12 focus:border-purple-500 focus:ring-purple-500"
                  required
                />
              </div>

              {/* Email */}
              <div className="space-y-2">
                <label className="text-sm font-medium text-gray-300 flex items-center gap-2">
                  <Mail className="w-4 h-4" />
                  Email Address
                </label>
                <input
                  type="email"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                  placeholder="you@company.com"
                  className="bg-gray-800 w-full placeholder:pl-3 rounded-md border-gray-700 text-white placeholder:text-gray-500 h-12 focus:border-purple-500 focus:ring-purple-500"
                  required
                />
              </div>

              {/* Password */}
              <div className="space-y-2">
                <label className="text-sm font-medium text-gray-300 flex items-center gap-2">
                  <Lock className="w-4 h-4" />
                  Password
                </label>
                <input
                  type="password"
                  placeholder="Create a password"
                  className="bg-gray-800 w-full placeholder:pl-3 rounded-md border-gray-700 text-white placeholder:text-gray-500 h-12 focus:border-purple-500 focus:ring-purple-500"
                  required
                />
              </div>

              <button
                type="submit"
                disabled={isLoading}
                className="w-full h-12 flex items-center justify-center rounded-md cursor-pointer bg-gradient-to-r from-purple-500 to-pink-600 hover:from-purple-600 hover:to-pink-700 text-white font-semibold text-base group"
              >
                {isLoading ? (
                  <div className="flex items-center gap-2">
                    <div className="w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin"></div>
                    Creating...
                  </div>
                ) : (
                  <>
                    Sign Up
                    <ArrowRight className="ml-2 w-5 h-5 group-hover:translate-x-1 transition-transform" />
                  </>
                )}
              </button>
            </>
          ) : (
            <>
              {/* Email */}
              <div className="space-y-2">
                <label className="text-sm font-medium text-gray-300 flex items-center gap-2">
                  <Mail className="w-4 h-4" />
                  Email Address
                </label>
                <input
                  type="email"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                  placeholder="you@company.com"
                  className="bg-gray-800 w-full placeholder:pl-3 rounded-md border-gray-700 text-white placeholder:text-gray-500 h-12 focus:border-cyan-500 focus:ring-cyan-500"
                  required
                />
              </div>

              {/* Password */}
              <div className="space-y-2">
                <label className="text-sm font-medium text-gray-300 flex items-center gap-2">
                  <Lock className="w-4 h-4" />
                  Password
                </label>
                <input
                  type="password"
                  placeholder="Enter your password"
                  className="bg-gray-800 w-full placeholder:pl-3 rounded-md border-gray-700 text-white placeholder:text-gray-500 h-12 focus:border-cyan-500 focus:ring-cyan-500"
                  required
                />
              </div>

              <div className="flex items-center justify-between text-sm">
                <label className="flex items-center gap-2 text-gray-400 cursor-pointer">
                  <input
                    type="checkbox"
                    className="rounded border-gray-600 bg-gray-800"
                  />
                  Remember me
                </label>
                <a
                  href="#"
                  className="text-cyan-400 hover:text-cyan-300 transition-colors"
                >
                  Forgot password?
                </a>
              </div>

              <button
                type="submit"
                disabled={isLoading}
                className="w-full h-12 flex items-center justify-center rounded-md cursor-pointer bg-gradient-to-r from-cyan-500 to-blue-600 hover:from-cyan-600 hover:to-blue-700 text-white font-semibold text-base group"
              >
                {isLoading ? (
                  <div className="flex items-center gap-2">
                    <div className="w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin"></div>
                    Signing In...
                  </div>
                ) : (
                  <>
                    Sign In
                    <ArrowRight className="ml-2 w-5 h-5 group-hover:translate-x-1 transition-transform" />
                  </>
                )}
              </button>
            </>
          )}
        </motion.form>
      </AnimatePresence>
    </div>
  )
}

export default Form
