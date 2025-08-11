import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom' // Add this import
import './index.css'
import LoginScreen from './screens/login'
import LandingScreen from './screens/landing'
import RegisterScreen from './screens/register'
import HomeScreen from './screens/home/home'
import NotFoundScreen from './screens/not-found'

createRoot(document.getElementById('root')).render(
  <Router>
    <Routes>
      <Route path="/" element={<LandingScreen />} />
      <Route path="/login" element={<LoginScreen />} />
      <Route path="/register" element={<RegisterScreen />} />
      <Route path="/home" element={<HomeScreen />} />
      <Route path="/*" element={<NotFoundScreen />} />

    </Routes>
  </Router>
)