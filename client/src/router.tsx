<<<<<<< HEAD
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

const Routers = () => {
  return (
    <>
      <Suspense fallback={<Loading />}>
        <Routes>
          <Route path="/" element={<Landing />} />
          <Route path="/Download" element={<Download/>} />
          <Route path="/dashboard" element={<Dashboard />} />
          <Route path="/auth" element={<Auth />} />
          <Route path="/Profile" element={<Profiles />} />
          <Route path="/Alerts" element={<Alerts />} />
        </Routes>
      </Suspense>
      {/* Footer */}
      <Footer />
    </>
  );
};
export default Routers;
=======
import React, { Suspense, lazy } from 'react'
import { Routes, Route } from 'react-router-dom'
import Loading from './components/loading'

import Footer from './components/footer/footer'
const Dashboard = lazy(() => import('./pages/Dashboard/dashbaord'))
const Alerts = lazy(() => import('./pages/Alert/Alert'))
const Profiles = lazy(() => import('./pages/Profile/Profile'))
const Landing = lazy(() => import('./pages/Landing/landing'))
const Auth = lazy(() => import('./pages/Auth/Auth'))

const Routers = () => {
  return (
    <>
      <Suspense fallback={<Loading />}>
        <Routes>
          <Route path="/" element={<Landing />} />
          <Route path="/dashboard" element={<Dashboard />} />
          <Route path="/auth" element={<Auth />} />
          <Route path="/Profile" element={<Profiles />} />
          <Route path="/Alerts" element={<Alerts />} />
        </Routes>
      </Suspense>
      {/* Footer */}
      <Footer />
    </>
  )
}
export default Routers
>>>>>>> 51a603614efb67c2b802b1a2f50a6c7eba24100d
