import toast from "react-hot-toast"
import api from "../config/axios"
import { AxiosError } from "axios"

export const FetchDeviceList = async () => {
    try {
        const res = await api.get("/api/list/agent")
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

export const GenerateParingToken = async () => {
    try {
        const res = await api.post("/api/paringToken/token")
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


