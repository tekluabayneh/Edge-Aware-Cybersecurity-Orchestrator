import { useState } from "react";
import { motion, AnimatePresence } from "framer-motion";
import {
    Download as DownloadIcon,
    Terminal,
    Copy,
    Check,
    AlertCircle,
    CheckCircle2,
    Bell,
    Globe,
    Package,
    Shield,
    ExternalLink,
    type LucideIcon
} from "lucide-react";

// Types
interface InstallMethod {
    id: string;
    name: string;
    icon: LucideIcon;
    description: string;
    commands: string[];
    comingSoon?: boolean;
}

interface InstallMethods {
    linux: InstallMethod[];
    macos: InstallMethod[];
    windows: InstallMethod[];
}

interface OSConfig {
    icon: LucideIcon;
    label: string;
    color: string;
}

interface Requirement {
    icon: LucideIcon;
    title: string;
    desc: string;
}

// Installation methods data
const installMethods: InstallMethods = {
    linux: [
        {
            id: "deb",
            name: "APT Package (.deb)",
            icon: Package,
            description: "Recommended for Ubuntu/Debian",
            commands: [
                "wget https://github.com/tekluabayneh/Edge-Aware-Cybersecurity-Orchestrator/releases/download/v2.0.29/agent_2.0.29_amd64.deb",
                "sudo apt install ./agent_2.0.29_amd64.deb",
                "sudo agent"
            ]
        },
        {
            id: "curl",
            name: "Install Script",
            icon: Terminal,
            description: "Quick one-liner for any distro",
            commands: [
                "curl -sSL https://raw.githubusercontent.com/tekluabayneh/Edge-Aware-Cybersecurity-Orchestrator/v2.0.29/scripts/install.sh | sudo bash",
                "sudo agent"
            ]
        },
        {
            id: "tarball",
            name: "Manual (.tar.gz)",
            icon: DownloadIcon,
            description: "Universal Linux method",
            commands: [
                "curl -sSL https://github.com/tekluabayneh/Edge-Aware-Cybersecurity-Orchestrator/releases/download/v2.0.29/agent_2.0.29_linux_amd64.tar.gz -o agent.tar.gz",
                "tar -xzf agent.tar.gz",
                "sudo mv agent /usr/local/bin/agent",
                "sudo agent"
            ]
        }
    ],
    macos: [
        {
            id: "brew",
            name: "Homebrew",
            icon: Package,
            description: "Recommended for macOS",
            commands: [
                "brew tap tekluabayneh/edge-aware",
                "brew install agent",
                "sudo agent"
            ]
        },
        {
            id: "curl-mac",
            name: "Install Script",
            icon: Terminal,
            description: "Quick installation",
            commands: [
                "curl -sSL https://raw.githubusercontent.com/tekluabayneh/Edge-Aware-Cybersecurity-Orchestrator/v2.0.29/scripts/install-macos.sh | sudo bash",
                "sudo agent"
            ]
        },
        {
            id: "tarball-mac",
            name: "Manual (.tar.gz)",
            icon: DownloadIcon,
            description: "Direct download",
            commands: [
                "curl -sSL https://github.com/tekluabayneh/Edge-Aware-Cybersecurity-Orchestrator/releases/download/v2.0.29/agent_2.0.29_darwin_amd64.tar.gz -o agent.tar.gz",
                "tar -xzf agent.tar.gz",
                "sudo mv agent /usr/local/bin/agent",
                "sudo agent"
            ]
        }
    ],
    windows: [
        {
            id: "exe",
            name: "Installer (.exe)",
            icon: DownloadIcon,
            description: "Recommended - UAC elevated",
            commands: [
                "# Download from Releases page:",
                "https://github.com/tekluabayneh/Edge-Aware-Cybersecurity-Orchestrator/releases/download/v2.0.29/agent_installer_2.0.29_amd64.exe",
                "",
                "# Then run in PowerShell (Admin):",
                "Start-Process -FilePath \".\\agent_installer_2.0.29_amd64.exe\" -Verb RunAs"
            ]
        },
        {
            id: "zip",
            name: "Portable (.zip)",
            icon: Package,
            description: "No installation required",
            commands: [
                "# Download ZIP from Releases",
                "# Extract and run in PowerShell:",
                "cd C:\\path\\to\\extracted",
                ".\\agent.exe"
            ]
        },
        {
            id: "winget",
            name: "Winget (Coming Soon)",
            icon: Terminal,
            description: "Microsoft Package Manager",
            commands: [
                "winget install tekluabayneh.Agent",
                "agent"
            ],
            comingSoon: true
        }
    ]
};

// OS Configuration
const osConfig: Record<string, OSConfig> = {
    linux: { icon: Globe, label: "Linux", color: "from-blue-500 to-cyan-500" },
    macos: { icon: Terminal, label: "macOS", color: "from-purple-500 to-pink-500" },
    windows: { icon: Shield, label: "Windows", color: "from-green-500 to-emerald-500" }
};

// Requirements data
const requirements: Requirement[] = [
    { icon: Shield, title: "64-bit System", desc: "x86_64 or ARM64 architecture" },
    { icon: Bell, title: "Admin Access", desc: "Sudo/root for system installation" },
    { icon: CheckCircle2, title: "Network", desc: "Internet for updates & alerts" }
];

// Copyable command block component
interface CommandBlockProps {
    commands: string[];
    language?: string;
}

function CommandBlock({ commands, language = "bash" }: CommandBlockProps) {
    const [copied, setCopied] = useState(false);

    const copyToClipboard = () => {
        navigator.clipboard.writeText(commands.join("\n"));
        setCopied(true);
        setTimeout(() => setCopied(false), 2000);
    };

    return (
        <div className="relative group">
            <div className="absolute -inset-0.5 bg-gradient-to-r from-cyan-500 to-purple-500 rounded-xl opacity-30 group-hover:opacity-50 transition-opacity duration-300 blur" />
            <div className="relative bg-gray-950 rounded-xl border border-gray-800 overflow-hidden">
                {/* Header */}
                <div className="flex items-center justify-between px-4 py-3 bg-gray-900/50 border-b border-gray-800">
                    <div className="flex items-center gap-2">
                        <Terminal className="w-4 h-4 text-cyan-400" />
                        <span className="text-xs font-mono text-gray-400">{language}</span>
                    </div>
                    <button
                        onClick={copyToClipboard}
                        className="flex items-center gap-1.5 px-3 py-1.5 text-xs font-medium text-gray-400 hover:text-white bg-gray-800 hover:bg-gray-700 rounded-lg transition-all duration-200"
                    >
                        {copied ? (
                            <>
                                <Check className="w-3.5 h-3.5 text-green-400" />
                                <span className="text-green-400">Copied!</span>
                            </>
                        ) : (
                            <>
                                <Copy className="w-3.5 h-3.5" />
                                <span className="cursor-pointer">Copy</span>
                            </>
                        )}
                    </button>
                </div>

                {/* Commands */}
                <div className="p-4 overflow-x-auto">
                    <pre className="text-sm font-mono leading-relaxed">
                        {commands.map((cmd, idx) => (
                            <div key={idx} className={cmd.startsWith("#") || cmd.startsWith("https") ? "text-gray-500" : "text-cyan-300"}>
                                {cmd.startsWith("#") || cmd.startsWith("https") || cmd.startsWith("cd") || cmd.startsWith(".") ? cmd : `$ ${cmd}`}
                            </div>
                        ))}
                    </pre>
                </div>
            </div>
        </div>
    );
}

// Method card component
interface MethodCardProps {
    method: InstallMethod;
    isActive: boolean;
    onClick: () => void;
}

function MethodCard({ method, isActive, onClick }: MethodCardProps) {
    const Icon = method.icon;

    return (
        <button
            onClick={onClick}
            className={`w-full text-left p-4 rounded-xl border transition-all duration-300 ${isActive
                ? "bg-gradient-to-br from-cyan-500/20 to-purple-500/20 border-cyan-500/50 shadow-lg shadow-cyan-500/10"
                : "bg-gray-900/50 border-gray-800 hover:border-gray-700 hover:bg-gray-800/50"
                }`}
        >
            <div className="flex items-start gap-3">
                <div className={`p-2 rounded-lg ${isActive ? "bg-cyan-500/20" : "bg-gray-800"}`}>
                    <Icon className={`w-5 h-5 ${isActive ? "text-cyan-400" : "text-gray-400"}`} />
                </div>
                <div className="flex-1">
                    <div className="flex items-center gap-2">
                        <h4 className={`font-semibold ${isActive ? "text-white" : "text-gray-300"}`}>
                            {method.name}
                        </h4>
                        {method.comingSoon && (
                            <span className="px-2 py-0.5 text-xs font-medium text-yellow-400 bg-yellow-400/10 rounded-full">
                                Soon
                            </span>
                        )}
                    </div>
                    <p className="text-sm text-gray-500 mt-1">{method.description}</p>
                </div>
            </div>
        </button>
    );
}

// OS Tab component
interface OSTabProps {
    os: string;
    icon: LucideIcon;
    label: string;
    isActive: boolean;
    onClick: () => void;
}

function OSTab({ icon: Icon, label, isActive, onClick }: OSTabProps) {
    return (
        <button
            onClick={onClick}
            className={`flex items-center gap-3 px-6 py-4 cursor-pointer rounded-xl font-semibold transition-all duration-300 ${isActive
                ? "bg-gradient-to-r from-cyan-500 to-purple-500 text-white shadow-lg shadow-cyan-500/25"
                : "bg-gray-900/50 text-gray-400 hover:text-white hover:bg-gray-800/50"
                }`}
        >
            <Icon className="w-5 h-5" />
            <span className="text-lg">{label}</span>
        </button>
    );
}

// Main component
export default function Download() {
    const [activeOS, setActiveOS] = useState<"linux" | "macos" | "windows">("linux");
    const [activeMethod, setActiveMethod] = useState(0);

    const currentMethods = installMethods[activeOS];
    const currentMethod = currentMethods[activeMethod];
    const OSIcon = osConfig[activeOS].icon;

    return (
        <div className="min-h-screen bg-black text-white overflow-hidden relative">
            {/* Animated background */}
            <div className="absolute inset-0 bg-gradient-to-br from-blue-950/20 via-black to-purple-950/20" />
            <div className="absolute inset-0">
                <div className="absolute top-1/4 left-1/4 w-96 h-96 bg-blue-500/10 rounded-full blur-3xl animate-pulse" />
                <div className="absolute bottom-1/4 right-1/4 w-96 h-96 bg-purple-500/10 rounded-full blur-3xl animate-pulse" style={{ animationDelay: "1s" }} />
            </div>

            <div className="relative z-10 container mx-auto px-6 py-16 md:py-24 max-w-7xl">
                {/* Hero */}
                <motion.div
                    initial={{ opacity: 0, y: 30 }}
                    animate={{ opacity: 1, y: 0 }}
                    transition={{ duration: 0.8 }}
                    className="text-center mb-16"
                >
                    <h1 className="text-5xl md:text-7xl font-bold mb-6 bg-gradient-to-r from-blue-400 via-cyan-300 to-purple-400 bg-clip-text text-transparent">
                        Install Agent
                    </h1>
                    <p className="text-xl md:text-2xl text-gray-400 max-w-3xl mx-auto">
                        Choose your platform and installation method. Terminal commands ready to copy.
                    </p>
                </motion.div>

                {/* Prerequisites */}
                <motion.div
                    initial={{ opacity: 0, y: 30 }}
                    animate={{ opacity: 1, y: 0 }}
                    transition={{ duration: 0.8, delay: 0.2 }}
                    className="mb-16"
                >
                    <div className="bg-gradient-to-br from-gray-900/50 to-gray-800/30 border border-gray-800/50 backdrop-blur-xl rounded-2xl p-8">
                        <div className="flex items-start gap-4 mb-6">
                            <div className="p-3 rounded-xl bg-gradient-to-br from-cyan-500/20 to-blue-500/20 border border-cyan-500/30">
                                <AlertCircle className="w-6 h-6 text-cyan-400" />
                            </div>
                            <div>
                                <h3 className="text-2xl font-semibold text-white">Requirements</h3>
                                <p className="text-gray-400">Ensure your system meets these specifications</p>
                            </div>
                        </div>

                        <div className="grid md:grid-cols-3 gap-4">
                            {requirements.map((req, idx) => (
                                <div key={idx} className="flex items-start gap-3 p-4 rounded-xl bg-black/40 border border-gray-800/50">
                                    <req.icon className="w-5 h-5 text-cyan-400 mt-1 flex-shrink-0" />
                                    <div>
                                        <h4 className="font-semibold text-white">{req.title}</h4>
                                        <p className="text-sm text-gray-500">{req.desc}</p>
                                    </div>
                                </div>
                            ))}
                        </div>
                    </div>
                </motion.div>

                {/* OS Tabs */}
                <motion.div
                    initial={{ opacity: 0, y: 30 }}
                    animate={{ opacity: 1, y: 0 }}
                    transition={{ duration: 0.8, delay: 0.4 }}
                    className="mb-12"
                >
                    <div className="flex flex-wrap justify-center gap-4">
                        {Object.entries(osConfig).map(([key, config]) => (
                            <OSTab
                                key={key}
                                os={key}
                                icon={config.icon}
                                label={config.label}
                                isActive={activeOS === key}
                                onClick={() => {
                                    setActiveOS(key as "linux" | "macos" | "windows");
                                    setActiveMethod(0);
                                }}
                            />
                        ))}
                    </div>
                </motion.div>

                {/* Installation Methods */}
                <motion.div
                    initial={{ opacity: 0, y: 30 }}
                    animate={{ opacity: 1, y: 0 }}
                    transition={{ duration: 0.8, delay: 0.6 }}
                    className="grid lg:grid-cols-3 gap-8"
                >
                    {/* Method Selection */}
                    <div className="lg:col-span-1 space-y-3">
                        <h3 className="text-xl font-semibold text-white mb-4">Installation Method</h3>
                        {currentMethods.map((method, idx) => (
                            <MethodCard
                                key={method.id}
                                method={method}
                                isActive={activeMethod === idx}
                                onClick={() => setActiveMethod(idx)}
                            />
                        ))}
                    </div>

                    {/* Command Display */}
                    <div className="lg:col-span-2">
                        <div className="bg-gradient-to-br from-gray-900/50 to-gray-800/30 border border-gray-800/50 backdrop-blur-xl rounded-2xl p-8">
                            <div className="flex items-center gap-4 mb-6">
                                <div className={`p-3 rounded-xl bg-gradient-to-br ${osConfig[activeOS].color} bg-opacity-20`}>
                                    <OSIcon className="w-6 h-6 text-white" />
                                </div>
                                <div>
                                    <h3 className="text-2xl font-semibold text-white">{currentMethod.name}</h3>
                                    <p className="text-gray-400">{currentMethod.description}</p>
                                </div>
                            </div>

                            <AnimatePresence mode="wait">
                                <motion.div
                                    key={activeMethod}
                                    initial={{ opacity: 0, x: -20 }}
                                    animate={{ opacity: 1, x: 0 }}
                                    exit={{ opacity: 0, x: 20 }}
                                    transition={{ duration: 0.3 }}
                                >
                                    <CommandBlock commands={currentMethod.commands} />
                                </motion.div>
                            </AnimatePresence>

                            {/* Additional Info */}
                            <div className="mt-6 p-4 rounded-xl bg-blue-500/10 border border-blue-500/20">
                                <div className="flex items-start gap-3">
                                    <ExternalLink className="w-5 h-5 text-blue-400 mt-0.5 flex-shrink-0" />
                                    <div>
                                        <h4 className="font-semibold text-blue-300 mb-1">Need the latest version?</h4>
                                        <p className="text-sm text-gray-400">
                                            Check{" "}
                                            <a
                                                href="https://github.com/tekluabayneh/Edge-Aware-Cybersecurity-Orchestrator/releases"
                                                target="_blank"
                                                rel="noopener noreferrer"
                                                className="text-cyan-400 hover:text-cyan-300 underline"
                                            >
                                                GitHub Releases
                                            </a>{" "}
                                            for all available versions and changelogs.
                                        </p>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </motion.div>

                {/* Footer */}
                <motion.div
                    initial={{ opacity: 0 }}
                    animate={{ opacity: 1 }}
                    transition={{ duration: 0.8, delay: 0.8 }}
                    className="text-center mt-16"
                >
                    <p className="text-gray-500">
                        Having issues?{" "}
                        <a href="/support" className="text-cyan-400 hover:text-cyan-300 underline">
                            Contact Support
                        </a>
                    </p>
                </motion.div>
            </div>
        </div>
    );
}
