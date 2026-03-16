import { useState } from "react";
import { motion, AnimatePresence } from "framer-motion";
import { ArrowRight, Lock, Mail, User } from "lucide-react";
import axios, { AxiosError } from "axios";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import toast from "react-hot-toast";
import type { RegisterRequest } from "../../types/Alert";
import { validateFormData } from "../../utils/ValidateForm";
const Register = () => {
    const { register, handleSubmit, formState: { errors } } = useForm()
    const [isLoading, setisLoading] = useState(false)
    const BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL;
    const navigator = useNavigate()

    if (!BASE_URL) {
        console.log("base url is empty")
        return
    }

    const SubmitREgisterForm = async (data: RegisterRequest) => {
        console.log(data)
        try {
            setisLoading(true)
            const response = await axios.post(`${BASE_URL}/api/auth/r/register`, data)
            toast.success(response.data.message);

            const loginRes = await axios.post(`${BASE_URL}/api/auth/l/login`, {
                email: data.email,
                password: data.password
            })

            await cookieStore.set("token", loginRes.data.token)

            navigator("/paring")
            setisLoading(false)
        } catch (err) {
            setisLoading(false)
            if (err instanceof AxiosError) {
                toast.error(err.response?.data.message);
            } else {
                toast.error("operation failed");
            }
        }
    }

    const onSubmit = (data: RegisterRequest): void => {
        if (!validateFormData(data)) {
            return
        }
        SubmitREgisterForm(data)
    }

    return (
        <form className="relative w-full" style={{ perspective: 1000, }} onSubmit={handleSubmit(onSubmit)} >

            <AnimatePresence mode="wait">
                <motion.div
                    key="signin"
                    initial="initial"
                    animate="animate"
                    exit="exit"
                    className="space-y-5 backface-hidden"
                    style={{
                        transformStyle: "preserve-3d",
                    }}
                >
                    <>
                        <div className="space-y-2">
                            <label className="text-sm font-medium text-gray-300 flex items-center gap-2">
                                <Mail className="w-4 h-4" />
                                User Name
                            </label>
                            <input
                                type="text"
                                placeholder="username"
                                {...register("name", { required: "username is required" })}
                                className="bg-gray-800 w-full placeholder:pl-3 rounded-md border-gray-700 text-white placeholder:text-gray-500 h-12 focus:border-cyan-500 focus:ring-cyan-500"
                            />
                            {errors.name && <p className="text-red-500">{errors.name.message}</p>}

                        </div>

                        <div className="space-y-2">
                            <label className="text-sm font-medium text-gray-300 flex items-center gap-2">
                                <Mail className="w-4 h-4" />
                                Email Address
                            </label>
                            <input
                                type="email"
                                placeholder="Email"
                                {...register("email", { required: "Email is required" })}
                                className="bg-gray-800 w-full placeholder:pl-3 rounded-md border-gray-700 text-white placeholder:text-gray-500 h-12 focus:border-cyan-500 focus:ring-cyan-500"
                            />
                            {errors.email && <p className="text-red-500">{errors.email.message}</p>}

                        </div>

                        {/* Password */}
                        <div className="space-y-2">
                            <label className="text-sm font-medium text-gray-300 flex items-center gap-2">
                                <Lock className="w-4 h-4" />
                                Password
                            </label>
                            <input
                                type="password"
                                placeholder="Password"
                                {...register("password", { required: "Password is required" })}
                                className="bg-gray-800 w-full placeholder:pl-3 rounded-md border-gray-700 text-white placeholder:text-gray-500 h-12 focus:border-cyan-500 focus:ring-cyan-500"
                            />
                            {errors.password && <p className="text-red-500">{errors.password.message}</p>}
                        </div>

                        <div className="flex items-center justify-between text-sm">
                            <label className="flex items-center gap-2 text-gray-400 cursor-pointer">
                                <input
                                    type="checkbox"
                                    className="rounded border-gray-600 bg-gray-800"
                                />
                                Remember me
                            </label>
                            <a
                                href="#"
                                className="text-cyan-400 hover:text-cyan-300 transition-colors"
                            >
                                Forgot password?
                            </a>
                        </div>

                        <button
                            type="submit"
                            disabled={isLoading}
                            className="w-full h-12 flex items-center justify-center rounded-md cursor-pointer bg-gradient-to-r from-cyan-500 to-blue-600 hover:from-cyan-600 hover:to-blue-700 text-white font-semibold text-base group"
                        >
                            {isLoading ? (
                                <div className="flex items-center gap-2">
                                    <div className="w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin"></div>
                                    Signing up...
                                </div>
                            ) : (
                                <>
                                    Sign Up
                                    <ArrowRight className="ml-2 w-5 h-5 group-hover:translate-x-1 transition-transform" />
                                </>
                            )}
                        </button>

                    </>
                </motion.div>
            </AnimatePresence>
        </form>
    );
}

export default Register;
