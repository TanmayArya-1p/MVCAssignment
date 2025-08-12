import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom' // Add this import
import './index.css'
import LoginScreen from './screens/login'
import LandingScreen from './screens/landing'
import RegisterScreen from './screens/register'
import HomeScreen from './screens/home/home'
import NotFoundScreen from './screens/not-found'
import OrderScreen from './screens/order'
import ItemScreen from './screens/items'
import Modal from 'react-modal';
import UserScreen from './screens/users';

Modal.setAppElement('#root');

createRoot(document.getElementById('root')).render(
  <Router>
    <Routes>
      <Route path="/" element={<LandingScreen />} />
      <Route path="/login" element={<LoginScreen />} />
      <Route path="/register" element={<RegisterScreen />} />
      <Route path="/home" element={<HomeScreen />} />
      <Route path="/order/:orderid" element={<OrderScreen />} />
      <Route path="/items" element={<ItemScreen />} />
      <Route path="/users" element={<UserScreen />} />
      <Route path="/*" element={<NotFoundScreen />} />

    </Routes>
  </Router>
)