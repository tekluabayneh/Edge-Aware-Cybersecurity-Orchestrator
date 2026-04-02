import axios from "axios"


let BASE_URL = "__VITE_REPLACE_ME__";

const api = axios.create({ baseURL: BASE_URL, timeout: 20000, withCredentials: true })

async function FechBaseURL(): Promise<void> {
  try {
    const res = await axios.get("https://gist.githubusercontent.com/tekluabayneh/f959fffc8449bcaf9901007cb25d830a/raw/backend_base_url")
    console.log("url", res.data.split("=")[1].trim())

    let freshUrl = res.data.split("=")[1].trim()

    api.defaults.baseURL = freshUrl
    BASE_URL = freshUrl
  } catch (err) {
    console.log(err)
  }
}

FechBaseURL()


api.interceptors.request.use(async (config) => {
  const publicRoutes = ["/api/auth/l/login", "/api/auth/l/register", "/api/oauth/google", "/api/oauth/github"];


  if (publicRoutes.some(route => config.url?.includes(route))) {
    return config;
  }

  const token = localStorage.getItem("token");
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }


  return config;
}, (error) => {
  return Promise.reject(error);
});

export default api
