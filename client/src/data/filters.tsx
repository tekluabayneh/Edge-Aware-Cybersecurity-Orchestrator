import { Server, Cpu, HardDrive, Network } from "lucide-react";

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