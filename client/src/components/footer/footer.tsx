import { Shield } from 'lucide-react'
const Footer = () => {
  return (
    <>
      <footer className="relative py-8 px-4 sm:px-6 border-t border-gray-800/50">
        <div className="max-w-7xl mx-auto flex flex-col sm:flex-row justify-between items-center gap-4">
          <div className="flex items-center gap-2">
            <Shield className="w-6 h-6 text-cyan-400" />
            <span className="text-xl font-bold bg-gradient-to-r from-cyan-400 to-blue-500 bg-clip-text text-transparent">
              EdgeGuard
            </span>
          </div>
          <p className="text-gray-500 text-sm text-center">
            © 2024 EdgeGuard. The future of cybersecurity.
          </p>
        </div>
      </footer>
    </>
  )
}
export default Footer
