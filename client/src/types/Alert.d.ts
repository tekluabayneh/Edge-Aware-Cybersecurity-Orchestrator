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
