import { AxiosError } from "axios";
import api from "../config/axios";
import toast from "react-hot-toast";

const GetAllAlerts = async () => {
    try {
        const res = await api.get("/api/device/alert")
        // console.log(data.message)
        // toast.success(data.message);
        console.log("data", res.data)
        return res.data
    } catch (error) {
        if (error instanceof AxiosError) {
            toast.error(error.response?.data.message);
        } else {
            toast.error("operation failed");
        }
    }
}

export default GetAllAlerts
