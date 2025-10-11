import React, { useState } from "react";
import { Lock, Bell, Mail, Shield } from "lucide-react";
import { Button } from "@/components/ui/button";

export default function SecurityPreferences() {
  const [preferences, setPreferences] = useState({
    twoFactor: true,
    emailNotifications: true,
    alertNotifications: true,
    autoLock: false
  });

  const togglePreference = (key) => {
    setPreferences(prev => ({ ...prev, [key]: !prev[key] }));
  };

  return (
    <div className="bg-gray-900 border border-gray-800 rounded-xl p-6">
      <h3 className="text-lg font-semibold text-white mb-6">Security Preferences</h3>
      <div className="space-y-4">
        <div className="flex items-center justify-between p-4 bg-gray-800/50 rounded-lg">
          <div className="flex items-center gap-3">
            <Lock className="w-5 h-5 text-cyan-400" />
            <div>
              <p className="text-sm font-medium text-white">Two-Factor Authentication</p>
              <p className="text-xs text-gray-400">Add extra security to your account</p>
            </div>
          </div>
          <button
            onClick={() => togglePreference("twoFactor")}
            className={`relative w-12 h-6 rounded-full transition-colors ${
              preferences.twoFactor ? "bg-cyan-500" : "bg-gray-700"
            }`}
          >
            <span
              className={`absolute top-1 left-1 w-4 h-4 bg-white rounded-full transition-transform ${
                preferences.twoFactor ? "translate-x-6" : ""
              }`}
            />
          </button>
        </div>

        <div className="flex items-center justify-between p-4 bg-gray-800/50 rounded-lg">
          <div className="flex items-center gap-3">
            <Mail className="w-5 h-5 text-cyan-400" />
            <div>
              <p className="text-sm font-medium text-white">Email Notifications</p>
              <p className="text-xs text-gray-400">Receive security updates via email</p>
            </div>
          </div>
          <button
            onClick={() => togglePreference("emailNotifications")}
            className={`relative w-12 h-6 rounded-full transition-colors ${
              preferences.emailNotifications ? "bg-cyan-500" : "bg-gray-700"
            }`}
          >
            <span
              className={`absolute top-1 left-1 w-4 h-4 bg-white rounded-full transition-transform ${
                preferences.emailNotifications ? "translate-x-6" : ""
              }`}
            />
          </button>
        </div>

        <div className="flex items-center justify-between p-4 bg-gray-800/50 rounded-lg">
          <div className="flex items-center gap-3">
            <Bell className="w-5 h-5 text-cyan-400" />
            <div>
              <p className="text-sm font-medium text-white">Alert Notifications</p>
              <p className="text-xs text-gray-400">Get notified of critical threats</p>
            </div>
          </div>
          <button
            onClick={() => togglePreference("alertNotifications")}
            className={`relative w-12 h-6 rounded-full transition-colors ${
              preferences.alertNotifications ? "bg-cyan-500" : "bg-gray-700"
            }`}
          >
            <span
              className={`absolute top-1 left-1 w-4 h-4 bg-white rounded-full transition-transform ${
                preferences.alertNotifications ? "translate-x-6" : ""
              }`}
            />
          </button>
        </div>

        <div className="flex items-center justify-between p-4 bg-gray-800/50 rounded-lg">
          <div className="flex items-center gap-3">
            <Shield className="w-5 h-5 text-cyan-400" />
            <div>
              <p className="text-sm font-medium text-white">Auto-Lock Session</p>
              <p className="text-xs text-gray-400">Lock after 15 minutes of inactivity</p>
            </div>
          </div>
          <button
            onClick={() => togglePreference("autoLock")}
            className={`relative w-12 h-6 rounded-full transition-colors ${
              preferences.autoLock ? "bg-cyan-500" : "bg-gray-700"
            }`}
          >
            <span
              className={`absolute top-1 left-1 w-4 h-4 bg-white rounded-full transition-transform ${
                preferences.autoLock ? "translate-x-6" : ""
              }`}
            />
          </button>
        </div>
      </div>

      <Button className="w-full mt-6 bg-gradient-to-r from-blue-600 to-cyan-600 hover:from-blue-700 hover:to-cyan-700">
        Save Preferences
      </Button>
    </div>
  );
}
