type AlertNotification = {
    id: string,
    type: 'success' | 'error',
    body: string,
    footer?: string
  }
type NotificationProps = Omit<AlertNotification, 'id'>

export type { 
    AlertNotification,
    NotificationProps
}