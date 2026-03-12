import UserInfoCard from '../../components/profile/UserinfoCard'
import ActivityLog from '../../components/profile/ActivityLog'
import SecurityPreferences from '../../components/profile/SecurityPreferance'
import { FechUserProfile } from '../../hooks/fetchUserProfile'
import { useQuery } from '@tanstack/react-query'

export default function Profile() {
    const { data, isLoading, error } = useQuery({
        queryKey: ["userProfile"],
        queryFn: FechUserProfile
    });

    if (isLoading) {
        return (
            <div className="flex items-center justify-center h-screen">
                <div className="text-gray-400">Loading profile...</div>
            </div>
        )
    }
    if (error) return <p>Something went wrong</p>;

    return (
        <div className="space-y-8">
            {/* Header */}
            <div className="flex flex-col gap-2">
                <h1 className="text-3xl font-bold text-white">User Profile</h1>
                <p className="text-gray-400">
                    Manage your account settings and preferences
                </p>
            </div>

            {/* Profile Grid */}
            <div className="grid lg:grid-cols-3 gap-6">
                <div className="lg:col-span-1">
                    <UserInfoCard user={data["UserInfo"]} />
                </div>
                <div className="lg:col-span-2 space-y-6">
                    <SecurityPreferences user={data["UserInfo"]} />
                    <ActivityLog />
                </div>
            </div>
        </div>
    )
}
