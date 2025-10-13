import React, { Suspense, lazy } from "react";
import { Routes, Route } from "react-router-dom";
import Loading from "./components/loading";
const Dashboard = lazy(() => import("./pages/Dashboard/dashbaord"));
const Alerts = lazy(() => import("./pages/Alert/Alert"))
const Profiles = lazy(() => import("./pages/Profile/Profile"))

const Routers = () => {
  return (
    <Suspense fallback={<Loading />}>
      <Routes>
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="/Profile" element={<Profiles/>} />
        <Route path="/Alerts" element={<Alerts/>} />
      </Routes>
    </Suspense>
  );
};
export default Routers;
