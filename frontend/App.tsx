import './index.css'
import { BrowserRouter, Routes, Route } from 'react-router-dom'
import Register from './pages/Register'
import Login from './pages/Login'
import { AuthProvider } from './context/AuthContext'
import HomePage from './pages/HomePage'
import ProtectRoute from './ProtectRoute'
import MembershipPage from './pages/MembershipPage'
import AfterPayment from './components/AfterPayment'
import { QueryClient, QueryClientProvider } from 'react-query'



function App() {
  const queryClient = new QueryClient()
  return (
    <AuthProvider>
      <BrowserRouter>
        <QueryClientProvider client={queryClient}>
          <Routes>
            <Route path="/" element={<HomePage />} />
            <Route path="/register" element={<Register />} />
            <Route path="/login" element={<Login />} />
            <Route element={<ProtectRoute />} >
              <Route path="/membership" element={<MembershipPage />} />
              <Route path="/afterpayment" element={<AfterPayment />} />
            </Route>

          </Routes>
        </QueryClientProvider>
      </BrowserRouter>
    </AuthProvider >
  )
}

export default App
