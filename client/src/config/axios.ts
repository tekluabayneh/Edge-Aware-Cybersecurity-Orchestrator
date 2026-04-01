import axios from "axios"

const BASE_URL = "__VITE_REPLACE_ME__";

if (!BASE_URL) {
  console.log("baseURL is not found")
}

const api = axios.create({ baseURL: BASE_URL, timeout: 20000, })

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
