import React, { useState, useEffect } from "react";
import { User } from "@/entities/User";
import UserInfoCard from "../components/profile/UserInfoCard";
import ActivityLog from "../components/profile/ActivityLog";
import SecurityPreferences from "../components/profile/SecurityPreferences";

export default function Profile() {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    loadUserData();
  }, []);

  const loadUserData = async () => {
    try {
      const userData = await User.me();
      setUser(userData);
    } catch (error) {
      console.error("Error loading user data:", error);
    }
    setLoading(false);
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center h-screen">
        <div className="text-gray-400">Loading profile...</div>
      </div>
    );
  }

  return (
    <div className="space-y-8">
      {/* Header */}
      <div className="flex flex-col gap-2">
        <h1 className="text-3xl font-bold text-white">User Profile</h1>
        <p className="text-gray-400">Manage your account settings and preferences</p>
      </div>

      {/* Profile Grid */}
      <div className="grid lg:grid-cols-3 gap-6">
        <div className="lg:col-span-1">
          <UserInfoCard user={user} />
        </div>
        <div className="lg:col-span-2 space-y-6">
          <SecurityPreferences />
          <ActivityLog />
        </div>
      </div>
    </div>
  );
}
