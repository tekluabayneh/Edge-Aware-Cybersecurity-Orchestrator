import { useState, useEffect, useRef } from "react";
import { useQuery } from "@tanstack/react-query";
import { FetchAllNotification, MarkRadAlllNotification, MarkRadSingleNotification, type NotificationType } from "../../hooks/Notification";
import { Bell, Check, Inbox } from "lucide-react";


export default function NotificationBell() {
    const [open, setOpen] = useState(false);
    const [notifications, setNotifications] = useState<NotificationType[] | []>([]);
    const ref = useRef<HTMLDivElement>(null);

    const { data, isLoading } = useQuery({
        queryKey: ["allNotification"],
        queryFn: FetchAllNotification
    });

    useEffect(() => {
        if (!data) {
            setNotifications([])
            return
        }
        setNotifications(data?.notification)
    }, [data, setNotifications])


    // Close on outside click
    useEffect(() => {
        const handler = (e: MouseEvent) => {
            if (ref.current && e.target instanceof Node && !ref.current.contains(e.target)) {
                setOpen(false)
            }
        }
        document.addEventListener("mousedown", handler);
        return () => document.removeEventListener("mousedown", handler);
    }, []);


    const markRead = async (id: number, user_id: number) => {
        console.log(user_id)
        MarkRadSingleNotification(id)
    };

    const markAllRead = async () => {
        MarkRadAlllNotification()
    };

    if (isLoading) {
        return (
            <div className="flex items-center justify-center h-screen">
                <div className="text-gray-400">Loading alerts...</div>
            </div>
        )
    }

    if (!data) {
        return (
            <div className="flex items-center justify-center h-screen">
                <div className="text-gray-400">Opps! you don't have alerts yet</div>
            </div>
        )
    }

    return (
        <div className="relative" ref={ref}>
            {/* Bell button */}
            <button
                onClick={() => setOpen((v) => !v)}
                className="relative p-2 rounded-xl hover:bg-slate-100 transition-colors focus:outline-none"
            >
                <Bell className="h-5 w-5 cursor-pointer text-slate-600" />
                {notifications?.length > 0 && (
                    <span className="absolute -top-0.5 -right-0.5 flex h-[18px] min-w-[18px] items-center justify-center rounded-full bg-red-500 px-1 text-[10px] font-bold text-white ring-2 ring-white">
                        {notifications?.length > 9 ? "9+" : notifications?.length}
                    </span>
                )}
            </button>

            {/* Popover */}
            {open && (
                <div className="absolute right-0 mt-2 w-80 bg-white rounded-2xl shadow-xl border border-slate-200 overflow-hidden z-50">
                    {/* Header */}
                    <div className="flex items-center justify-between px-4 py-3 border-b border-slate-100">
                        <span className="text-sm font-semibold text-slate-900">Notifications</span>
                        {notifications?.length > 0 && (
                            <button
                                onClick={markAllRead}
                                className="text-xs text-slate-400 hover:text-slate-700 transition-colors"
                            >
                                Mark all read
                            </button>
                        )}
                    </div>

                    {/* List */}
                    <div className="max-h-96 overflow-y-auto">
                        {notifications?.length === 0 || !notifications ? (
                            <div className="flex flex-col items-center justify-center py-10 gap-2">
                                <Inbox className="h-6 w-6 text-slate-300" />
                                <p className="text-sm text-slate-400">You're all caught up!</p>
                            </div>
                        ) : (
                            notifications?.map((n) => {
                                return (
                                    <div
                                        key={n.ID}
                                        className={`flex gap-3 px-4 py-3 border-b border-slate-100 last:border-0 ${!n.IsRead ? "bg-slate-50" : "bg-white"}`}
                                    >
                                        {/* <div className={`shrink-0 h-8 w-8 rounded-full flex items-center justify-center`}> */}
                                        {/* <Icon className={`h-4 w-4`} /> */}
                                        {/* </div> */}
                                        <div className="flex-1 min-w-0">
                                            <div className="flex items-start justify-between gap-2">
                                                <p className={`text-sm leading-snug ${!n.IsRead ? "font-semibold text-slate-900" : "font-medium text-slate-500"}`}>
                                                    {n.Title}
                                                </p>
                                                {!n.IsRead && (
                                                    <button onClick={() => markRead(n.ID, n.UserID)} className="shrink-0 cursor-pointer p-1 rounded hover:bg-slate-200 transition-colors">
                                                        <Check className="h-3 w-3 text-slate-400" />
                                                    </button>
                                                )}
                                            </div>
                                            <p className="text-xs text-slate-500 mt-0.5 line-clamp-2">{n.Message}</p>
                                            <p className="text-[11px] text-slate-400 mt-1">
                                                {new Date(n.CreatedAt).toLocaleString("en-US", {
                                                    year: "numeric",
                                                    month: "short",
                                                    day: "numeric",
                                                    hour: "2-digit",
                                                    minute: "2-digit",
                                                })}
                                            </p>
                                        </div>
                                    </div>
                                );
                            })
                        )}
                    </div>
                </div>
            )}
        </div>
    );
}
