
import { Suspense, lazy } from "react";
import { Routes, Route } from "react-router-dom";
import Loading from "./components/loading";

import Footer from "./components/footer/footer";
const Dashboard = lazy(() => import("./pages/Dashboard/dashbaord"));
const Alerts = lazy(() => import("./pages/Alert/Alert"));
const Profiles = lazy(() => import("./pages/Profile/Profile"));
const Landing = lazy(() => import("./pages/Landing/landing"));
const Auth = lazy(() => import("./pages/Auth/Auth"));
const Download = lazy(() => import("./pages/Download/download"));
const Device = lazy(() => import("./pages/DeviceParing/Paring"));
const NotFound = lazy(() => import("./pages/NotFound/Notfound"))
const Routers = () => {
  return (
    <>
      <Suspense fallback={<Loading />}>
        <Routes>
          <Route path="/" element={<Landing />} />
          <Route path="/Download" element={<Download/>} />
          <Route path="/dashboard" element={<Dashboard />} />
          <Route path="/paring" element={<Device/>} />
          <Route path="/auth" element={<Auth />} />
          <Route path="/Profile" element={<Profiles />} />
          <Route path="/Alerts" element={<Alerts />} />
          <Route path="*" element={<NotFound />} />
        </Routes>
      </Suspense>
      {/* Footer */}
      <Footer />
    </>
  );
};
export default Routers;
