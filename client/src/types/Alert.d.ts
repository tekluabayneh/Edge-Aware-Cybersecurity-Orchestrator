import type { number } from "framer-motion"

export type AlertType = {
    CreatedAt: string
    agent_id: string
    agent_token: string
    alertType: string
    message: string
    netetwork: object
    performance: object
    risk_level: string
    security: string
    status: string
    summery: string
}

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
AgentID:string,
AgentToken:string
AgentVersion :string 
CreatedAt:string
DeviceName:string 
ID:number 
LastSeen:string 
MachineID:string
Os:string
Status:string
UserID: string
}





