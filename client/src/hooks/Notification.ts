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

export const FetchAllNotification = async (): NotificationType => {
    try {
        const res = await api.get("/api/getAll/get")
        return res.data
    } catch (error) {
        if (error instanceof AxiosError) {
            toast.error(error.response?.data.message);
        } else {
            toast.error("operation failed");
        }
    }

}

export const MarkRadSingleNotification = async () => {
    try {
        const res = await api.post("/api/updateSingle/updateSingle")
        toast.success(res?.data.message);
        return res.data
    } catch (error) {
        if (error instanceof AxiosError) {
            toast.error(error.response?.data.message);
        } else {
            toast.error("operation failed");
        }
    }

}

export const MarkRadAlllNotification = async (id: number) => {
    try {
        const res = await api.post("/api/updatAll/update", { id: id })
        toast.success(res?.data.message);
        return res.data
    } catch (error) {
        if (error instanceof AxiosError) {
            toast.error(error.response?.data.message);
        } else {
            toast.error("operation failed");
        }
    }

}


