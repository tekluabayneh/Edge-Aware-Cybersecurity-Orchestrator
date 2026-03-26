import toast from "react-hot-toast"
import api from "../config/axios"
import { AxiosError } from "axios"

export type NotificationType = {
    ID: number,
    UserID: number,
    Title: string,
    Message: string,
    IsRead: boolean,
    CreatedAt: string,
}

export const FetchAllNotification = async () => {
    try {
        const res = await api.get("/api/getAll/get")
        return res.data
    } catch (error) {
        if (error instanceof AxiosError) {
            toast.error(error.response?.data.message);
            return []
        } else {
            toast.error("operation failed");
            return []
        }
    }

}

export const MarkRadSingleNotification = async (id: number) => {
    console.log(id)
    try {
        const res = await api.post("/api/updateSingle/updateSingle")
        toast.success(res?.data.message);
        return res.data
    } catch (error) {
        if (error instanceof AxiosError) {
            toast.error(error.response?.data.message);
            return []
        } else {
            toast.error("operation failed");
            return []
        }
    }

}

export const MarkRadAlllNotification = async () => {
    try {
        const res = await api.post("/api/updatAll/update")
        toast.success(res?.data.message);
        return res.data
    } catch (error) {
        if (error instanceof AxiosError) {
            toast.error(error.response?.data.message);
            return []
        } else {
            toast.error("operation failed");
            return []
        }
    }

}


