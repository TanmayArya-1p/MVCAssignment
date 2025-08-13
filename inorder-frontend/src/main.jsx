import { createRoot } from 'react-dom/client'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import '@/index.css'
import LoginScreen from '@/screens/login'
import LandingScreen from '@/screens/landing'
import RegisterScreen from '@/screens/register'
import HomeScreen from '@/screens/home/home'
import NotFoundScreen from '@/screens/notFound'
import OrderScreen from '@/screens/order'
import ItemScreen from '@/screens/items'
import Modal from 'react-modal';
import UserScreen from '@/screens/users';
import UpdateItemScreen from '@/screens/updateItem'
import AdminProtectedRoute from '@/middleware/adminRoute'
import SignedIn from './middleware/signedIn'

Modal.setAppElement('#root');

createRoot(document.getElementById('root')).render(
  <Router>
    <Routes>
      <Route path="/" element={<LandingScreen />} />
      <Route path="/login" element={<LoginScreen />} />
      <Route path="/register" element={<RegisterScreen />} />
    
      <Route
        path="/home"
        element={
          <SignedIn>
              <HomeScreen />
          </SignedIn>
        }
      />  

      <Route
        path="/order/:orderid"
        element={
          <SignedIn>
              <OrderScreen />
          </SignedIn>
        }
      />     

      <Route
        path="/items"
        element={
          <SignedIn>
            <AdminProtectedRoute>
              <ItemScreen />
            </AdminProtectedRoute>
          </SignedIn>
        }
      />

      <Route
        path="/users"
        element={
          <SignedIn>
            <AdminProtectedRoute>
              <UserScreen />
            </AdminProtectedRoute>
          </SignedIn>
        }
      />

      <Route
        path="/items/update"
        element={
          <SignedIn>
            <AdminProtectedRoute>
              <UpdateItemScreen />
            </AdminProtectedRoute>
          </SignedIn>
        }
      />

      <Route path="/*" element={<NotFoundScreen />} />

    </Routes>
  </Router>
)