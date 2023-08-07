import { BrowserRouter, Routes, Route } from "react-router-dom";
import { Appointments } from './Components/Appointments';
import { Department } from './Components/Department';
import { Footer } from './Components/Footer';
import { Home } from "./Components/Home";
import { Header } from "./Components/Header";
import Protected from "./Components/Protected";
import { About } from "./Components/About";
import { useEffect, useState } from "react";
import { Dashboard } from "./Components/Dashboard";
import AuthGuard  from "./auth/AuthGuard";
import AuthForDashboard from "./auth/AuthForDashboard";

function App() {
  //var key  = true;  
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  const handleLogin = () => {
    localStorage.setItem('isAuthenticated', true);
    setIsAuthenticated(true);
  };

  const handleLogout = () => {
    localStorage.removeItem('isAuthenticated');
    setIsAuthenticated(false);
  }


  return (
    <div className="App">
      <BrowserRouter>
        <Header handleLogin={handleLogin} handleLogout={handleLogout}/>
        <Routes>
          <Route path='/'>
            <Route index element={<AuthForDashboard Component={Home}/>} />
            <Route path="/dashboard" element={<AuthGuard role={["patient","doctor"]} Component={Dashboard} />} />
            <Route path='/appointments' element={<AuthGuard role={["patient"]} Component={Appointments} />} />
            <Route path='/departments' element={<Department />} />
            <Route path="/about" element={<About />} /> 
            <Route path="*" element={<AuthForDashboard Component={Home}/>} />
          </Route>
        </Routes>
        <Footer />
      </BrowserRouter>
    </div>
  );
}

export default App;