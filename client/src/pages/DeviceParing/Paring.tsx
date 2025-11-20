import  { useState, useEffect } from 'react';
import { useQuery } from '@tanstack/react-query';
import { Copy, RefreshCw, Check } from 'lucide-react';
import DeviceList from '../../components/DeviceList/DeviceList';
import data from "../../data/Device.json"
const style = document.createElement('style');
style.textContent = `
  @keyframes fadeIn {
    from { opacity: 0; }
    to { opacity: 1; }
  }
  @keyframes slideUp {
    from { opacity: 0; transform: translateY(20px); }
    to { opacity: 1; transform: translateY(0); }
  }
`;
document.head.appendChild(style);

export default function PairDevice() {
  const [pairingToken, setPairingToken] = useState('');
  const [copied, setCopied] = useState(false);

  const { data: devices = [], isLoading } = useQuery({
    queryKey: ['devices'],
    queryFn: () => data
  });

  const generateToken = () => {
    const token = Math.random().toString(36).substring(2, 10).toUpperCase();
    setPairingToken(token);
  };

  useEffect(() => {
    generateToken();
  }, []);

  const handleCopy = async () => {
    await navigator.clipboard.writeText(pairingToken);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  const handleRefresh = () => {
    generateToken();
    setCopied(false);
  };

  return (
    <div className="min-h-screen bg-black">
      <div className="max-w-5xl mx-auto px-6 py-12">
        {/* Header */}
        <div className="mb-12 animate-[fadeIn_0.6s_ease-out]">
          <h1 className="text-3xl font-light text-white mb-2">
            Pair Device
          </h1>
          <p className="text-gray-400">
            Connect your devices securely using a pairing token
          </p>
        </div>

        {/* Pairing Token Card */}
        <div className="bg-gradient-to-br from-gray-900 to-gray-800 rounded-xl border border-gray-700 p-8 mb-8 shadow-2xl animate-[slideUp_0.7s_ease-out]">
          <div className="text-center">
            <h2 className="text-sm font-medium text-gray-400 uppercase tracking-wide mb-4">
              Pairing Token
            </h2>

            <div className="bg-black/50 rounded-lg px-8 py-6 mb-6 inline-block backdrop-blur-sm border border-gray-700 transition-all duration-300 hover:scale-105">
              <div className="font-mono text-4xl font-semibold text-white tracking-wider">
                {pairingToken}
              </div>
            </div>

            <div className="flex items-center justify-center gap-3 mb-6">
              <button
                onClick={handleCopy}
                className="inline-flex cursor-pointer items-center gap-2 px-6 py-3 bg-white text-black rounded-lg hover:bg-gray-200 transition-all duration-200 shadow-lg hover:shadow-xl hover:scale-105 active:scale-95"
              >
                {copied ? (
                  <>
                    <Check className="w-4 h-4 animate-[bounce_0.5s_ease-in-out]" />
                    Copied
                  </>
                ) : (
                  <>
                    <Copy className="w-4 h-4" />
                    Copy Token
                  </>
                )}
              </button>

              <button
                onClick={handleRefresh}
                className="inline-flex cursor-pointer items-center gap-2 px-6 py-3 bg-gray-800 border border-gray-700 text-gray-200 rounded-lg hover:bg-gray-700 transition-all duration-200 hover:scale-105 active:scale-95"
              >
                <RefreshCw className="w-4 h-4" />
                Refresh
              </button>
            </div>

            <p className="text-sm text-gray-400 max-w-md mx-auto">
              Open your agent on your device and enter this token to pair.
            </p>
          </div>
        </div>

        {/* Paired Devices */}
        <div className="mb-4 animate-[fadeIn_0.9s_ease-out]">
          <h2 className="text-xl font-light text-white mb-4">
            Paired Devices
          </h2>
        </div>

        {isLoading ? (
          <div className="bg-gray-900 rounded-lg border border-gray-800 p-12 text-center">
            <div className="inline-block w-8 h-8 border-3 border-gray-700 border-t-white rounded-full animate-spin" />
          </div>
        ) : (
          <DeviceList devices={devices} />
        )}
      </div>
    </div>
  );
}