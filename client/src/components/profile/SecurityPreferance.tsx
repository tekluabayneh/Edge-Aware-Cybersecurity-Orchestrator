import { useEffect, useState } from 'react'
import { Lock, Bell, Mail } from 'lucide-react'
import { UpdateUserProfile, type userInfoType } from '../../hooks/fetchUserProfile'
import { useMutation } from '@tanstack/react-query'
import api from '../../config/axios'
import toast from 'react-hot-toast'
import { AxiosError } from 'axios'
import { useNavigate } from 'react-router-dom'
export type ProfileInfo = {
    two_fa: boolean
    notification: boolean
    alert_notification: boolean
}

export default function SecurityPreferences({ user }: { user: userInfoType | null }) {
    const navigator = useNavigate()
    const [profileInfo, setProfileInfo] = useState<ProfileInfo>({
        two_fa: false,
        notification: false,
        alert_notification: true
    })

    useEffect(() => {
        if (!user) return

        setProfileInfo({
            // @ts-expect-error index is type number
            two_fa: user.two_fa ?? false,
            // @ts-expect-error index is type number
            notification: user.notification ?? false,
            // @ts-expect-error index is type number
            alert_notification: user.alert_notification ?? false
        })
    }, [user])

    const togglePreference = (key: keyof ProfileInfo) => {
        setProfileInfo(prev => ({
            ...prev,
            [key]: !prev[key]
        }))
    }

    const mutation = useMutation({
        mutationFn: UpdateUserProfile,
    })


    async function UPdate2FA(flag: boolean) {
        try {
            const res = await api.post("/api/send/2fa", flag ? { Enable: "enable" } : { Enable: "disable" })
            if (res.data?.data?.otpUrl != undefined || res.data?.data?.rq != undefined) {
                navigator("/EnableQr", {
                    state: {
                        otpUrl: res.data?.data?.otpUrl,
                        qr: res.data?.data?.qr
                    }
                })
            }
            if (res.data.message) {
                toast.success(res.data.message);
            }

        } catch (error) {
            if (error instanceof AxiosError) {
                toast.error(error.response?.data.message);
            } else {
                toast.error("operation failed");
            }
        }

    }


    const SubmitUpdateProfile = () => {
        UPdate2FA(profileInfo.two_fa)
        mutation.mutate(profileInfo)
    }

    if (!user) return null

    return (
        <div className="bg-gray-900 border border-gray-800 rounded-xl p-6">
            <h3 className="text-lg font-semibold text-white mb-6">
                Security Preferences
            </h3>

            <div className="space-y-4">

                {/* TWO FA */}
                <div className="flex items-center justify-between p-4 bg-gray-800/50 rounded-lg">
                    <div className="flex items-center gap-3">
                        <Lock className="w-5 h-5 text-cyan-400" />
                        <div>
                            <p className="text-sm font-medium text-white">
                                Two-Factor Authentication
                            </p>
                            <p className="text-xs text-gray-400">
                                Add extra security to your account
                            </p>
                        </div>
                    </div>

                    <button
                        onClick={() => togglePreference("two_fa")}
                        className={`relative w-12 h-6 rounded-full transition-colors ${profileInfo.two_fa ? 'bg-cyan-500' : 'bg-gray-700'
                            }`}
                    >
                        <span
                            className={`absolute top-1 left-1 w-4 h-4 bg-white rounded-full transition-transform ${profileInfo.two_fa ? 'translate-x-6' : ''
                                }`}
                        />
                    </button>
                </div>

                {/* EMAIL */}
                <div className="flex items-center justify-between p-4 bg-gray-800/50 rounded-lg">
                    <div className="flex items-center gap-3">
                        <Mail className="w-5 h-5 text-cyan-400" />
                        <div>
                            <p className="text-sm font-medium text-white">
                                Email Notifications
                            </p>
                            <p className="text-xs text-gray-400">
                                Receive security updates via email
                            </p>
                        </div>
                    </div>

                    <button
                        onClick={() => togglePreference("notification")}
                        className={`relative w-12 h-6 rounded-full transition-colors ${profileInfo.notification ? 'bg-cyan-500' : 'bg-gray-700'
                            }`}
                    >
                        <span
                            className={`absolute top-1 left-1 w-4 h-4 bg-white rounded-full transition-transform ${profileInfo.notification ? 'translate-x-6' : ''
                                }`}
                        />
                    </button>
                </div>

                {/* ALERT */}
                <div className="flex items-center justify-between p-4 bg-gray-800/50 rounded-lg">
                    <div className="flex items-center gap-3">
                        <Bell className="w-5 h-5 text-cyan-400" />
                        <div>
                            <p className="text-sm font-medium text-white">
                                Alert Notifications
                            </p>
                            <p className="text-xs text-gray-400">
                                Get notified of critical threats
                            </p>
                        </div>
                    </div>

                    <button
                        onClick={() => togglePreference("alert_notification")}
                        className={`relative w-12 h-6 rounded-full transition-colors ${profileInfo.alert_notification ? 'bg-cyan-500' : 'bg-gray-700'
                            }`}
                    >
                        <span
                            className={`absolute top-1 left-1 w-4 h-4 bg-white rounded-full transition-transform ${profileInfo.alert_notification ? 'translate-x-6' : ''
                                }`}
                        />
                    </button>
                </div>

            </div>

            <button onClick={() => SubmitUpdateProfile()} className="w-full rounded cursor-pointer p-2 mt-6 bg-gradient-to-r from-blue-600 to-cyan-600 hover:from-blue-700 hover:to-cyan-700">

                Save Preferences
            </button>
        </div>
    )
}
