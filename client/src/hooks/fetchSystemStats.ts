
import toast from "react-hot-toast"
import api from "../config/axios"
import { AxiosError } from "axios"

export const FechUserAgentSystem = async () => {
    try {
        const res = await api.get("/api/data/telemetry")
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

