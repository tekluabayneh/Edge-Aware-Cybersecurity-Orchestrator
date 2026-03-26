import toast from "react-hot-toast"
import api from "../config/axios"
import { AxiosError } from "axios"
import { type ProfileInfo } from "../components/profile/SecurityPreferance"
export type userInfoType = {
    alert_notification: string
    created_at: string
    email: string
    id: number
    name: string
    notification: string
    phone: string
    photo: string
    two_fa: string
}

type UserResponse = {
    id: number
    name: string
    il: string
    photo: string
    ne: string
    a: string
    notification: string
    alertnotification: string
    createdat: string
}



export const FechUserProfile = async (): Promise<UserResponse | []> => {
    try {
        const res = await api.get("/api/get/profile")
        return res.data
    } catch (error) {
        if (error instanceof AxiosError) {
            toast.error(error.response?.data.message);
            if (error.status == 401 || error.status == 400) {
                window.location.href = "/Auth"
            }
            return []
        } else {
            toast.error("operation failed");
            return []
        }
    }

}

export const UpdateUserProfile = async (usreinfo: ProfileInfo): Promise<userInfoType | []> => {
    try {
        const res = await api.post("/api/user/profile", usreinfo)
        if (res.status) {
            toast.success(res.data.message);
        }
        return res.data
    } catch (error) {
        if (error instanceof AxiosError) {
            toast.error(error.response?.data.message);
            if (error.status == 401 || error.status == 400) {
                window.location.href = "/Auth"
            }
            return []
        } else {
            toast.error("operation failed");
            return []
        }
    }
}
