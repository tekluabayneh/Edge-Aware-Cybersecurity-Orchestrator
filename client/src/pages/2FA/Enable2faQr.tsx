import { useState, useEffect } from "react";
import { useLocation } from "react-router-dom";
import api from "../../config/axios";
import { AxiosError } from "axios";
import toast from "react-hot-toast";
import { useNavigate } from "react-router-dom";
const Enable2faQr = () => {
  //@ts-ignore
  const [qrUrl, setQrUrl] = useState<"" | undefined>("");
  const [secret, setSecret] = useState<null | "">("");
  const [code, setCode] = useState("");
  const [loading, setLoading] = useState(true);
  const [verifying, setVerifying] = useState(false);
  const [message, setMessage] = useState(null);
  const [qrLoaded, setQrLoaded] = useState(false);
  const [copied, setCopied] = useState(false);
  const [mounted, setMounted] = useState(false);
  const location = useLocation()
  const navigate = useNavigate()
  const { otpUrl, qr } = location.state

  useEffect(() => {
    setTimeout(() => setMounted(true), 50);
    const t = setTimeout(() => {
      setQrUrl(qr);
      setSecret(otpUrl);
      setLoading(false);
    }, 900);
    return () => clearTimeout(t);
  }, []);

  const handleVerify = async () => {
    if (code.length < 6) {
      setMessage({ type: "error", text: "" });
      toast("Enter a valid 6-digit code.");
      return;
    }

    try {
      setVerifying(true);
      setCode("")
      const res = await api.post("api/check/2fa", { "code": code })
      if (res.data.message) {
        toast.success(res.data.message);
      }

      if (res.status == 200) {
        navigate("/Dashboard")
      }

    } catch (error) {
      if (error instanceof AxiosError) {
        toast.error(error.response?.data.message);
      } else {
        toast.error("operation failed");
      }
    }

    setMessage(null);
    setVerifying(false);
  };

  const handleCopy = () => {
    navigator.clipboard.writeText(secret);
    setCopied(true);
    setTimeout(() => setCopied(false), 2200);
  };

  const qrUrlImg = `data:image/png;base64,${qr}`

  return (
    <div className="bg-[#030303] min-h-screen flex items-center justify-center px-4 relative overflow-hidden">

      {/* Ambient orbs */}
      <div className="pointer-events-none absolute w-[480px] h-[480px] rounded-full top-[-80px] left-[-120px] bg-[radial-gradient(circle,rgba(99,60,180,0.12)_0%,transparent_70%)]" />
      <div className="pointer-events-none absolute w-[400px] h-[400px] rounded-full bottom-[-60px] right-[-100px] bg-[radial-gradient(circle,rgba(30,100,200,0.08)_0%,transparent_70%)]" />

      {/* Card */}
      <div
        className={`w-full max-w-[400px] bg-[rgba(8,8,8,0.95)] rounded-[20px] pt-9 pb-8 px-8 shadow-[0_32px_80px_rgba(0,0,0,0.8),0_0_0_1px_rgba(255,255,255,0.06)] transition-opacity ${mounted ? "opacity-100" : "opacity-0"}`}
      >

        {/* Header */}
        <div className="flex flex-col items-center mb-7">

          <div className="w-[44px] h-[44px] mb-4 flex items-center justify-center bg-white/5 border border-white/10 rounded-xl">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="rgba(255,255,255,0.75)" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round">
              <rect x="3" y="11" width="18" height="11" rx="2" />
              <path d="M7 11V7a5 5 0 0 1 10 0v4" />
            </svg>
          </div>

          <h1 className="text-white text-[16px] font-semibold tracking-[-0.02em] mb-1">
            Enable Two-Factor Authentication
          </h1>

          <p className="text-white/35 text-[12px] text-center leading-[1.6]">
            Scan the QR code with your authenticator app,<br />
            or copy the secret key manually.
          </p>

        </div>

        {/* QR */}
        <div className="flex justify-center mb-6">
          {loading ? (
            <div className="w-[180px] h-[180px] bg-white/5 border border-white/10 rounded-[14px] flex items-center justify-center">
              <svg className="animate-spin" width="20" height="20" viewBox="0 0 24 24">
                <circle cx="12" cy="12" r="10" stroke="rgba(255,255,255,0.08)" strokeWidth="2.5" fill="none" />
                <path d="M12 2a10 10 0 0 1 10 10" stroke="rgba(255,255,255,0.45)" strokeWidth="2.5" strokeLinecap="round" />
              </svg>
            </div>
          ) : (
            <div className={`${qrLoaded ? "opacity-100" : "opacity-0"} transition-opacity`}>
              <div className="p-3 bg-white rounded-[14px] shadow-[0_0_0_1px_rgba(0,0,0,0.1),0_8px_32px_rgba(0,0,0,0.5)]">
                <img
                  src={qrUrlImg}
                  alt="2FA QR"
                  onLoad={() => setQrLoaded(true)}
                  className="w-[156px] h-[156px] block rounded"
                />
              </div>
            </div>
          )}
        </div>

        {/* Secret */}
        {secret && (
          <div className="mb-6">

            <p className="text-white/25 text-[11px] text-center mb-2 tracking-[0.06em] uppercase">
              Manual secret key
            </p>

            <button
              onClick={handleCopy}
              className="w-full flex items-center justify-between gap-3 bg-white/5 border border-white/10 rounded-[10px] py-[10px] px-[14px] cursor-pointer"
            >

              <code className="text-white/70 hide-scrollbar border-none overflow-scroll text-[12px] font-mono text-left tracking-[0.18em] font-sm">
                {secret}
              </code>

              <span className={`flex-shrink-0 transition-opacity ${copied ? "opacity-100" : "opacity-40"}`}>
                {copied ? (
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="rgb(74,222,128)" strokeWidth="2.5">
                    <polyline points="20 6 9 17 4 12" />
                  </svg>
                ) : (
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="rgba(255,255,255,0.6)" strokeWidth="1.8">
                    <rect x="9" y="9" width="13" height="13" rx="2" />
                    <path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1" />
                  </svg>
                )}
              </span>

            </button>

            {copied && (
              <p className="text-green-400 text-[11px] text-center mt-1">
                Copied to clipboard
              </p>
            )}

          </div>
        )}

        {/* Divider */}
        <div className="h-px bg-white/10 mb-5" />

        {/* Input */}
        <div className="mb-3">

          <label className="text-white/35 text-[11px] uppercase tracking-[0.06em] block mb-2">
            Confirmation code
          </label>

          <input
            type="text"
            inputMode="numeric"
            maxLength={6}
            value={code}
            onChange={(e) => {
              setCode(e.target.value.replace(/\D/g, "").slice(0, 6));
              setMessage(null);
            }}
            placeholder="• • • • • •"
            className="w-full outline-none bg-white/5 border border-white/10 rounded-[10px] py-[13px] px-4 text-white text-[18px] font-mono tracking-[0.3em] text-center caret-white/60"
          />

        </div>

        {/* Message */}
        {message && (
          <div className={`mb-4 text-center text-sm ${message.type === "success" ? "text-green-400" : "text-red-400"}`}>
            {message.text}
          </div>
        )}

        {/* Button */}
        <button
          onClick={handleVerify}
          disabled={verifying || code.length < 6}
          className="w-full py-[13px] px-4 rounded-[10px] text-[13px] font-semibold tracking-[-0.01em] bg-purple-600 text-white hover:bg-purple-500 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          {verifying ? "Verifying…" : "Confirm & Enable"}
        </button>

      </div>

      {/* Footer */}
      <p className="absolute bottom-7 left-0 right-0 text-center text-white/20 text-[11px]">
        Having trouble? <span className="text-white/40 cursor-pointer">Get help</span>
      </p>

    </div>
  );
};

export default Enable2faQr;
