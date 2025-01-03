import './index.css'
import { BrowserRouter, Routes, Route } from 'react-router-dom'
import Register from './pages/RegisterPage'
import Login from './pages/Login'
import { AuthProvider } from './context/AuthContext'
import HomePage from './pages/HomePage'
import ProtectRoute from './ProtectRoute'
import MembershipPage from './pages/MembershipPage'
import { QueryClient, QueryClientProvider } from 'react-query'
import ProfilePage from './pages/ProfilePage'
import { SubscriptionProvider } from './context/SubscriptionContext'
import SubscriptionRedirect from './components/SubscriptionRedirect'



function App() {
  const queryClient = new QueryClient()
  return (
    <AuthProvider>
      <BrowserRouter>
        <QueryClientProvider client={queryClient}>
          <SubscriptionProvider>
            <Routes>
              <Route path="/" element={<HomePage />} />
              <Route path="/register" element={<Register />} />
              <Route path="/login" element={<Login />} />
              <Route element={<ProtectRoute />} >
                <Route path="/profile" element={<ProfilePage />} />
                <Route path="/membership" element={
                  <SubscriptionRedirect>
                    <MembershipPage />
                  </SubscriptionRedirect>} />
              </Route>
            </Routes>
          </SubscriptionProvider>
        </QueryClientProvider>
      </BrowserRouter>
    </AuthProvider>
  )
}

export default App
