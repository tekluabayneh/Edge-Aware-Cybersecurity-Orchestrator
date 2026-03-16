import { Server, BluetoothConnected, Cpu, HardDrive, Network } from "lucide-react";
import {
    Monitor as WindowsIcon,
    Apple, Smartphone,
    LayoutDashboard,
    AlertTriangle,
    User,
    Download,
} from "lucide-react";

export const filters = [
    { value: "all", label: "All Alerts" },
    { value: "critical", label: "Critical" },
    { value: "high", label: "High" },
    { value: "medium", label: "Medium" },
    { value: "low", label: "Low" },
];
export const healthMetrics = [
    { name: "CPU Usage", value: 45, icon: Cpu, color: "cyan" },
    { name: "Memory", value: 67, icon: Server, color: "blue" },
    { name: "Storage", value: 38, icon: HardDrive, color: "green" },
    { name: "Network", value: 82, icon: Network, color: "purple" },
];
export const recentActivities = [
    {
        action: "Security scan completed",
        timestamp: new Date(Date.now() - 3600000),
    },
    {
        action: "Alert dismissed: Low priority phishing attempt",
        timestamp: new Date(Date.now() - 7200000),
    },
    {
        action: "System settings updated",
        timestamp: new Date(Date.now() - 14400000),
    },
    { action: "2FA enabled", timestamp: new Date(Date.now() - 86400000) },
    { action: "Password changed", timestamp: new Date(Date.now() - 172800000) },
];



export const downloadButtons = [
    {
        os: "Windows",
        icon: WindowsIcon,
        color: "from-blue-500 to-cyan-500",
        hoverColor: "hover:shadow-cyan-500/50",
        downloadUrl: "/downloads/agent-windows.exe",
        githubRelease: "https://github.com/yourrepo/releases/tag/v1.0.0",
    },
    {
        os: "macOS",
        icon: Apple,
        color: "from-purple-500 to-pink-500",
        hoverColor: "hover:shadow-purple-500/50",
        downloadUrl: "/downloads/agent-macos.dmg",
        githubRelease: "https://github.com/yourrepo/releases/tag/v1.0.0",
    },
    {
        os: "Linux",
        icon: Smartphone,
        color: "from-orange-500 to-red-500",
        hoverColor: "hover:shadow-orange-500/50",
        downloadUrl: "/downloads/agent-linux.tar.gz",
        githubRelease: "https://github.com/yourrepo/releases/tag/v1.0.0",
    },
];


export const navigationItems = [
    {
        title: "Download",
        url: "/Download",
        icon: Download,
    },
    {
        title: "Dashboard",
        url: "/Dashboard",
        icon: LayoutDashboard,
    },
    {
        title: "Pair Device",
        url: "/paring",
        icon: BluetoothConnected,
    },
    {
        title: "Alerts",
        url: "/Alerts",
        icon: AlertTriangle,
    },
    {
        title: "Profile",
        url: "/Profile",
        icon: User,
    },
];

export const flipVariants = {
    initial: (isSignUp: boolean) => ({
        rotateY: isSignUp ? -180 : 180,
        opacity: 0,
        scale: 0.9,
    }),
    animate: {
        rotateY: 0,
        opacity: 1,
        scale: 1,
        transition: {
            duration: 0.6,
            ease: [0.45, 0, 0.55, 1],
        },
    },
    exit: (isSignUp: boolean) => ({
        rotateY: isSignUp ? 180 : -180,
        opacity: 0,
        scale: 0.9,
        transition: {
            duration: 0.5,
            ease: [0.45, 0, 0.55, 1],
        },
    }),
};

































    
