import { Navigate, Outlet } from "react-router-dom"
import { useAuth } from "./context/AuthContext"
import LoadingPage from "./pages/LoadingPage"

export default function ProtectRoute() {
    const { loading, isAuthenticated } = useAuth()

    if (loading) return <LoadingPage />
    if (!loading && !isAuthenticated) return <Navigate to="/login" replace />

    return (
        <Outlet />
    )
}
