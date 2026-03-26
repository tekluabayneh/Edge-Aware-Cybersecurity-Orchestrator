import { Suspense, lazy } from "react";
import { Routes, Route } from "react-router-dom";
import Loading from "./components/loading";

import Footer from "./components/footer/footer";
import Enable2faQr from "./pages/2FA/Enable2faQr";
import Verify2FA from "./pages/2FA/Verify2FA";
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
                    <Route path="/Download" element={<Download />} />
                    <Route path="/Alerts" element={<Alerts />} />
                    <Route path="/Dashboard" element={<Dashboard />} />
                    <Route path="/Paring" element={<Device />} />
                    <Route path="/Auth" element={<Auth />} />
                    <Route path="/Profile" element={<Profiles />} />
                    <Route path="/EnableQr" element={<Enable2faQr />} />
                    <Route path="/Verify2Fa" element={<Verify2FA />} />
                    <Route path="*" element={<NotFound />} />
                </Routes>
            </Suspense>
            <Footer />
        </>
    );
};
export default Routers;
