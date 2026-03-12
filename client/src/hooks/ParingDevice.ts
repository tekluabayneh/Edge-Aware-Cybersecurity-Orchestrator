import toast from "react-hot-toast"
import api from "../config/axios"
import { AxiosError } from "axios"
import { useNavigate } from "react-router-dom"


export const FetchDeviceList = async (): DeviceType => {
    try {
        const res = await api.get("/api/list/agent")
        return res.data
    } catch (error) {
        if (error instanceof AxiosError) {
            toast.error(error.response?.data.message);
            if (error.status == 401 || error.status == 400) {
                window.location.href = "/Auth"
            }
        } else {
            toast.error("operation failed");
        }
    }

}

export const GenerateParingToken = async (): DeviceType => {
    try {
        const res = await api.post("/api/paringToken/token")
        return res.data
    } catch (error) {
        if (error instanceof AxiosError) {
            toast.error(error.response?.data.message);
            if (error.status == 401 || error.status == 400) {
                window.location.href = "/Auth"
            }
        } else {
            toast.error("operation failed");
        }
    }

}


