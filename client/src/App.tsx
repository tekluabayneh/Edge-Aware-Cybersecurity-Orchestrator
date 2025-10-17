import "./App.css";
import Routers from "./router";

import { Link, useLocation } from "react-router-dom";
import {
  LayoutDashboard,
  AlertTriangle,
  User,
  Shield,
  Activity,
} from "lucide-react";

const navigationItems = [
  {
    title: "Dashboard",
    url: "/Dashboard",
    icon: LayoutDashboard,
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

const App = () => {
  const location = useLocation();
  const isLandingPage =
    location.pathname === "/" || location.pathname == "/auth";
  return (
    <div className="w-full min-h-screen bg-slate-950 text-gray-100 flex">
      {/* Sidebar */}
      {!isLandingPage && (
        <aside className="fixed left-0 top-0 h-full w-56 bg-slate-900 border-r border-slate-800 z-50">
          <div className="p-6 border-b border-slate-800">
            <div className="flex items-center gap-3">
              <div className="w-10 h-10 bg-gradient-to-br from-indigo-500 to-violet-500 rounded-lg flex items-center justify-center">
                <Shield className="w-6 h-6 text-white" />
              </div>
              <div>
                <h1 className="text-xl font-bold bg-gradient-to-r from-indigo-400 to-violet-400 bg-clip-text text-transparent">
                  EdgeGuard
                </h1>
                <p className="text-xs text-gray-500">Security Platform</p>
              </div>
            </div>
          </div>

          <nav className="p-4 space-y-2">
            {navigationItems.map((item) => {
              const isActive = location.pathname === item.url;
              return (
                <Link
                  key={item.title}
                  to={item.url}
                  className={`flex items-center gap-3 px-4 py-3 rounded-lg transition-all duration-200 ${
                    isActive
                      ? "bg-gradient-to-r from-indigo-600/20 to-violet-600/20 text-violet-400 border border-violet-500/30"
                      : "text-gray-400 hover:text-gray-200 hover:bg-slate-800"
                  }`}
                >
                  <item.icon className="w-5 h-5" />
                  <span className="font-medium">{item.title}</span>
                </Link>
              );
            })}
          </nav>

          <div className="absolute bottom-0 left-0 right-0 p-4 border-t border-slate-800">
            <div className="flex items-center gap-2 text-xs">
              <Activity className="w-4 h-4 text-green-500 animate-pulse" />
              <span className="text-gray-500">System Active</span>
            </div>
          </div>
        </aside>
      )}

      {/* Main Content */}
      <main
        className={`${
          isLandingPage ? "ml-0" : "ml-60"
        }  flex-1 pr-3 relative z-10`}
      >
        <Routers />
      </main>

      {/* Background Effects */}
      <div className="fixed inset-0 pointer-events-none overflow-hidden z-0">
        <div className="absolute top-0 right-0 w-96 h-96 bg-indigo-500/5 rounded-full blur-3xl"></div>
        <div className="absolute bottom-0 left-64 w-96 h-96 bg-violet-500/5 rounded-full blur-3xl"></div>
      </div>
    </div>
  );
};
export default App;
