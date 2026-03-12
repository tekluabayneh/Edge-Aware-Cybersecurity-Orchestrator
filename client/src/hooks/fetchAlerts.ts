import { AxiosError } from "axios";
import api from "../config/axios";
import toast from "react-hot-toast";

export const GetAllAlerts = async () => {
    try {
        const res = await api.get("/api/device/alert")
        return res.data
    } catch (error) {
        if (error instanceof AxiosError) {
            if (error.status == 401 || error.status == 400) {
                window.location.href = "/Auth"
            }
        } else {
            toast.error("operation failed");
        }
    }
}

export default GetAllAlerts
