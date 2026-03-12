import { format } from 'date-fns';
import { useQuery } from '@tanstack/react-query';
import { FetchDeviceList } from '../../hooks/ParingDevice';

export default function DeviceList() {

    const formatLastSeen = (dateString: string) => {
        try {
            return format(new Date(dateString), 'MMM d, yyyy • h:mm a');
        } catch {
            return 'N/A';
        }
    };


    const { data, isLoading, error } = useQuery({
        queryKey: ["DeviceLIst"],
        queryFn: FetchDeviceList
    });


    if (isLoading) {
        return (
            <div className="flex items-center justify-center h-screen">
                <div className="text-gray-400">Loading profile...</div>
            </div>
        )
    }
    if (error) return <p>Something went wrong</p>;

    return (
        <div className="bg-gray-900 rounded-lg border border-gray-800 overflow-hidden shadow-xl">
            <div className="overflow-x-auto">
                <table className="w-full">
                    <thead>
                        <tr className="border-b border-gray-800">
                            <th className="px-6 py-4 text-left text-xs font-medium text-gray-400 uppercase tracking-wider">
                                Device Name
                            </th>
                            <th className="px-6 py-4 text-left text-xs font-medium text-gray-400 uppercase tracking-wider">
                                OS
                            </th>
                            <th className="px-6 py-4 text-left text-xs font-medium text-gray-400 uppercase tracking-wider">
                                Status
                            </th>
                            <th className="px-6 py-4 text-left text-xs font-medium text-gray-400 uppercase tracking-wider">
                                Last Seen
                            </th>
                        </tr>
                    </thead>
                    <tbody className="divide-y divide-gray-800">
                        {data.length === 0 ? (
                            <tr>
                                <td colSpan="4" className="px-6 py-12 text-center text-sm text-gray-500">
                                    No devices paired yet
                                </td>
                            </tr>
                        ) : (
                            (Array.isArray(data?.agents) ? data.agents : data?.agents ? [data.agents] : []).map((device, index) => (
                                <tr
                                    key={device.MachineID}
                                    className="hover:bg-gray-800 transition-all duration-200"
                                    style={{
                                        animation: `slideUp 0.5s ease-out ${index * 0.1}s both`
                                                                                                                                                                                                                                                                        }}
                                        >
                                        <td className="px-6 py-4 text-sm font-medium text-white">
                                            {device.Os}
                                    </td>
                                            <td className="px-6 py-4 text-sm text-gray-400">
                                                {device.Os}
                                    </td>
                                                <td className="px-6 py-4">
                                                    <span
                                                        className={`inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full text-xs font-medium transition-all duration-300 ${device.status === 'online'
                                                            ? 'bg-emerald-500/20 text-emerald-400 border border-emerald-500/30'
                                                            : 'bg-gray-700/50 text-gray-400 border border-gray-600/30'
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            }`}
                                        >
                                                    <span
                                                        className={`w-1.5 h-1.5 rounded-full ${device.Status === 'online' ? 'bg-emerald-400 animate-pulse' : 'bg-gray-500'
                                                            }`}
                                                    />
                                                    {device.Status === 'online' ? 'Online' : 'Offline'}
                        </span>
                                    </td>
                                                    <td className="px-6 py-4 text-sm text-gray-400">
                                                        {formatLastSeen(device.CreatedAt)}
                    </td>
                                </tr>
                            ))
                        )}
                                                    </tbody>
                                                </table>
                                            </div >
                                        </div >
                        );
}
