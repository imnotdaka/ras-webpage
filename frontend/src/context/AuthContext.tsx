import { createContext, ReactNode, useContext, useEffect, useLayoutEffect, useState } from "react";
import axios from "./axios";
import { FieldValues } from "react-hook-form";

type User = {
    first_name: string
    last_name: string
    email: string
    subscribed: boolean
}

type Subscription = {
    reason: string
    status: string
    date_created: string
    next_payment_date: string
}

interface AuthContextType {
    token: string
    setToken: React.Dispatch<React.SetStateAction<string>>;
    signup: (user: FieldValues) => Promise<void>;
    signin: (user: FieldValues) => Promise<void>;
    logout: () => Promise<void>;
    loading: boolean
    isAuthenticated: boolean
    isModalOpen: boolean
    setIsModalOpen: React.Dispatch<React.SetStateAction<boolean>>
    user: User
    GetSubscription: () => Promise<Subscription>
}

export const AuthContext = createContext<AuthContextType | null>(null)

export const useAuth = () => {
    const context = useContext(AuthContext)
    if (!context) {
        throw new Error("useAuth must be used within an AuthProvider")
    }
    return context
}

async function registerRequest(user: FieldValues) {
    return axios.post(`/user`, user)
}
async function signinRequest(user: FieldValues) {
    return axios.post(`/auth/user`, user)
}
async function subscriptionRequest() {
    return axios.get(`/subscription`)
}
async function logoutRequest() {
    return axios.post("/auth/logout")
}

export const AuthProvider = ({ children }: { children: ReactNode }) => {
    const [isAuthenticated, setIsAuthenticated] = useState(false)
    const [loading, setLoading] = useState(true)
    const [isModalOpen, setIsModalOpen] = useState(false)
    const [token, setToken] = useState("")
    const [user, setUser] = useState<User>({ first_name: "", last_name: "", email: "", subscribed: false })

    useEffect(() => {

        const fetchMe = async () => {
            try {
                console.log("authme")
                const res = await axios.get("/auth/me")
                console.log("res:", res.data)
                setToken(res.data.token)
                setUser(res.data.user)
                setIsAuthenticated(true)
                setLoading(false)
            }
            catch (err) {
                console.log("error in authme", err)
                setToken("")
            }
        }
        fetchMe()
    }, [token])

    useLayoutEffect(() => {
        const authInterceptor = axios.interceptors.request.use((config) => {
            config.headers.Authorization =
                token
                    ? `${token}`
                    : config.headers.Authorization
            return config
        })

        return () => {
            axios.interceptors.request.eject(authInterceptor)
        }
    }, [token])

    // useLayoutEffect(() => {
    //     const refreshInterceptor = axios.interceptors.response.use((res) => {

    //         }
    //     )
    // })

    async function GetSubscription() {
        try {
            const res = await subscriptionRequest()
            console.log(res.data)
            return res.data
        } catch (error) {
            console.log(error)
        }
    }

    async function signup(user: FieldValues) {
        try {
            const res = await registerRequest(user)
            setToken(res.data)
            setIsAuthenticated(true)
            setLoading(false)
        }
        catch (error) {
            console.log(error)
        }
    }

    const signin = async (user: FieldValues) => {
        try {
            const res = await signinRequest(user)
            setToken(res.data)
            setIsAuthenticated(true)

        } catch (error) {
            console.log(error)
        }
        setLoading(false)
    }

    const logout = async () => {
        try {
            await logoutRequest()
            setToken("")
            setIsAuthenticated(false)
            setLoading(false)
        } catch (error) {
            console.log(error)
        }
    }

    return <AuthContext.Provider
        value={{
            signup,
            signin,
            logout,
            loading,
            isAuthenticated,
            isModalOpen,
            setIsModalOpen,
            token,
            setToken,
            user,
            GetSubscription
        }}
    >
        {children}
    </AuthContext.Provider >
}