import { AxiosError } from "axios";
import React, { useState, useRef, useEffect } from "react";
import toast from "react-hot-toast";
import { useNavigate } from "react-router-dom";
import api from "../../config/axios";

const Verify2FA = () => {
    const [digits, setDigits] = useState(["", "", "", "", "", ""]);
    const [verifying, setVerifying] = useState(false);
    const [error, setError] = useState("");
    const [shake, setShake] = useState(false);
    const [success, setSuccess] = useState(false);
    const [mounted, setMounted] = useState(false);
    const [bounceIdx, setBounceIdx] = useState(null);
    const inputsRef = useRef([]);
    const navigator = useNavigate()

    useEffect(() => {
        setTimeout(() => setMounted(true), 50);
        setTimeout(() => inputsRef.current[0]?.focus(), 200);
    }, []);

    const handleChange = (index, value) => {
        if (!/^\d?$/.test(value)) return;

        const newDigits = [...digits];
        newDigits[index] = value;
        setDigits(newDigits);
        setError("");

        if (value) {
            setBounceIdx(index);
            setTimeout(() => setBounceIdx(null), 260);
            if (index < 5) inputsRef.current[index + 1]?.focus();
        }

        if (value && index === 5 && newDigits.every((d) => d !== "")) {
            handleVerify(newDigits.join(""));
        }
    };

    const handleKeyDown = (index, e) => {
        if (e.key === "Backspace" && !digits[index] && index > 0) {
            inputsRef.current[index - 1]?.focus();
        }
    };

    const handlePaste = (e) => {
        e.preventDefault();
        const pasted = e.clipboardData.getData("text").replace(/\D/g, "").slice(0, 6);

        if (!pasted.length) return;

        const newDigits = Array.from({ length: 6 }, (_, i) => pasted[i] || "");
        setDigits(newDigits);

        const next = newDigits.findIndex((d) => !d);
        inputsRef.current[next === -1 ? 5 : next]?.focus();

        if (newDigits.every((d) => d)) handleVerify(newDigits.join(""));
    };

    const handleVerify = async (codeStr) => {
        const toVerify = codeStr || digits.join("");

        if (toVerify.length < 6) {
            setError("Please enter all 6 digits.");
            return;
        }

        setVerifying(true);
        setError("");
        try {
            const res = await api.post("/api/verify2fa/verify", { code: toVerify })

            if (res.status == 200) {
                setSuccess(true);
                navigator("/Dashboard")
            }
            if (res.data.message) {
                toast.success(res.data.message);
            }

        } catch (error) {
            setShake(true);
            setTimeout(() => setShake(false), 500);
            setError("Invalid code. Please try again.");
            setDigits(["", "", "", "", "", ""]);
            setTimeout(() => inputsRef.current[0]?.focus(), 420);
            if (error instanceof AxiosError) {
                toast.error(error.response?.data.message);
            } else {
                toast.error("operation failed");
            }
        }

        setVerifying(false);
    };

    const code = digits.join("");

    return (
        <div className="bg-[#030303] min-w-full min-h-screen flex items-center justify-center px-4 relative overflow-hidden">

            {/* Ambient orbs */}
            <div className="pointer-events-none absolute w-[500px] h-[500px] rounded-full top-[-100px] right-[-80px] bg-[radial-gradient(circle,rgba(80,50,180,0.1)_0%,transparent_70%)]" />
            <div className="pointer-events-none absolute w-[420px] h-[420px] rounded-full bottom-[-80px] left-[-80px] bg-[radial-gradient(circle,rgba(20,80,200,0.07)_0%,transparent_70%)]" />

            {/* Card */}
            <div
                className={`w-full max-w-[400px] bg-[rgba(8,8,8,0.95)] rounded-[20px] pt-9 pb-8 px-8 shadow-[0_32px_80px_rgba(0,0,0,0.8),0_0_0_1px_rgba(255,255,255,0.06)] transition-opacity ${mounted ? "opacity-100" : "opacity-0"}`}
            >

                {/* Header */}
                <div className="flex flex-col items-center mb-8">

                    <div className="w-[44px] h-[44px] mb-4 flex items-center justify-center bg-white/5 border border-white/10 rounded-xl">
                        <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="rgba(255,255,255,0.75)" strokeWidth="1.8">
                            <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z" />
                        </svg>
                    </div>

                    <h1 className="text-white text-[16px] font-semibold tracking-[-0.02em] mb-1">
                        Two-Factor Authentication
                    </h1>

                    <p className="text-white/35 text-[12px] leading-[1.6] text-center">
                        Enter the 6-digit code from your<br />
                        authenticator app.
                    </p>

                </div>

                {/* Digit Inputs */}
                <div className="mb-5">

                    <div className={`flex items-center justify-center gap-2 ${shake ? "animate-pulse" : ""}`}>

                        {digits.map((digit, i) => (
                            <React.Fragment key={i}>

                                {i === 3 && (
                                    <div className="w-3 flex items-center justify-center">
                                        <div className="w-1 h-1 bg-white/20 rounded-full" />
                                    </div>
                                )}

                                {/* @ts-expect-error handleVerify shold be empty in her */}
                                <input ref={(el) => (inputsRef.current[i] = el)}
                                    type="text"
                                    inputMode="numeric"
                                    maxLength={1}
                                    value={digit}
                                    onChange={(e) => handleChange(i, e.target.value)}
                                    onKeyDown={(e) => handleKeyDown(i, e)}
                                    onPaste={i === 0 ? handlePaste : undefined}
                                    disabled={verifying || success}
                                    className={`w-[42px] h-[52px] text-center bg-white/5 border border-white/10 rounded-[10px] text-white text-[20px] font-mono focus:outline-none transition
                  ${digit ? "border-white/30" : ""}
                  ${error ? "border-red-400" : ""}
                  ${success ? "border-green-400" : ""}
                  ${bounceIdx === i ? "animate-bounce" : ""}`}
                                />

                            </React.Fragment>
                        ))}

                    </div>

                </div>

                {/* Error */}
                {error && (
                    <div className="text-red-400 text-[12px] text-center mb-3">
                        {error}
                    </div>
                )}

                {/* Success */}
                {success && (
                    <div className="text-green-400 text-[12px] text-center mb-3 flex items-center justify-center gap-1">
                        <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2.5">
                            <polyline points="20 6 9 17 4 12" />
                        </svg>
                        Verified — redirecting…
                    </div>
                )}

                {/* Button */}
                {/* @ts-expect-error handleVerify shold be empty in her */}
                <button onClick={() => handleVerify()}
                    disabled={verifying || code.length < 6 || success}
                    className="w-full py-[13px] rounded-[10px] text-[13px] font-semibold tracking-[-0.01em] bg-purple-600 text-white hover:bg-purple-500 disabled:opacity-50 disabled:cursor-not-allowed"
                >
                    {verifying ? "Verifying…" : success ? "Verified" : "Verify & Login"}
                </button>

                {/* Divider */}
                <div className="h-px bg-white/10 my-5" />

                {/* Help Links */}
                <div className="flex justify-between">

                    <button className="text-white/25 text-[11px] hover:text-white/50">
                        Use backup code
                    </button>

                    <button className="text-white/25 text-[11px] hover:text-white/50">
                        Need help?
                    </button>

                </div>

            </div>

            {/* Footer */}
            <p className="absolute bottom-7 left-0 right-0 text-center text-white/20 text-[11px]">
                Protected by two-factor authentication
            </p>

        </div>
    );
};

export default Verify2FA;
