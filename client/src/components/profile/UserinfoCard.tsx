import { User, Mail, Shield, Calendar } from 'lucide-react'
import { format } from 'date-fns'
import type { userType } from '../../types/Alert'
export default function UserInfoCard({ user }: { user: userType | null }) {
  return (
    <div className="bg-gray-900 border border-gray-800 rounded-xl p-6">
      <h3 className="text-lg font-semibold text-white mb-6">
        Account Information
      </h3>
      <div className="flex flex-col items-center mb-6">
        <div className="w-24 h-24 bg-gradient-to-br from-blue-500 to-cyan-500 rounded-full flex items-center justify-center mb-4">
          <User className="w-12 h-12 text-white" />
        </div>
        <h2 className="text-2xl font-bold text-white">
          {user?.full_name || 'User'}
        </h2>
        <p className="text-gray-400 text-sm">{user?.role || 'user'}</p>
      </div>

      <div className="space-y-4">
        <div className="flex items-center gap-3 p-3 bg-gray-800/50 rounded-lg">
          <Mail className="w-5 h-5 text-cyan-400" />
          <div>
            <p className="text-xs text-gray-400">Email Address</p>
            <p className="text-sm text-white">
              {user?.email || 'user@edge-aware.com'}
            </p>
          </div>
        </div>

        <div className="flex items-center gap-3 p-3 bg-gray-800/50 rounded-lg">
          <Shield className="w-5 h-5 text-cyan-400" />
          <div>
            <p className="text-xs text-gray-400">Account Type</p>
            <p className="text-sm text-white capitalize">
              {user?.role || 'Standard User'}
            </p>
          </div>
        </div>

        <div className="flex items-center gap-3 p-3 bg-gray-800/50 rounded-lg">
          <Calendar className="w-5 h-5 text-cyan-400" />
          <div>
            <p className="text-xs text-gray-400">Member Since</p>
            <p className="text-sm text-white">
              {user?.created_date
                ? format(new Date(user.created_date), 'MMMM d, yyyy')
                : 'N/A'}
            </p>
          </div>
        </div>
      </div>
    </div>
  )
}
