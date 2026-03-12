import type { number } from "framer-motion"

export type AlertType = {
  id: number
  title: string
  description: string
  severity: string
  status: string
  detected_at: string
  source_ip: string
  target: string
}

export type userType = {
  email: string
  full_name: string
  created_date: string
}

type MetricCardProps = {
  title: string
  value: string | number
  change: string
  trend: 'up' | 'down'
  icon: React.ComponentType
}

export type DeviceType = {
  user_id: number,
  machine_id: number,
  agent_version: string,
  os:string,
  status: string,
  last_seen:string,
  created_at:string,
}

type FormDataType ={
    password:string,
    email: string
}
type RegisterRequest = FormDataType & {
  username: string;
};