import axios from "axios"


const api = axios.create({
    baseURL: import.meta.env.VITE_BACKEND_BASE_URL,
    timeout: 1000,
})

// add token to each request
api.interceptors.request.use(async (config) => {
    const token = await cookieStore.get("token")
    if (!token) {
        return
    }
    if (token?.value) {
        config.headers.Authorization = `Bearer ${token.value}`;
    }
    return config
})

export default api
