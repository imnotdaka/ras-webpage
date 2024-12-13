import { createContext, ReactNode, useContext, useEffect, useState } from "react";
import axios from "./axios";
import { FieldValues } from "react-hook-form";
import Cookies from "js-cookie"



interface AuthContextType {
    signup: (user: FieldValues) => Promise<void>;
    signin: (user: FieldValues) => Promise<void>;
    loading: boolean
    isAuthenticated: boolean
    isModalOpen: boolean
    setIsModalOpen: React.Dispatch<React.SetStateAction<boolean>>
}

export const AuthContext = createContext<AuthContextType | null>(null)

export const useAuth = () => {
    const context = useContext(AuthContext)
    if (!context) {
        throw new Error("userAuth must be used within an AuthProvider")
    }
    return context
}

async function registerRequest(user: FieldValues) {
    return axios.post(`/user`, user)
}
async function signinRequest(user: FieldValues) {
    return axios.post(`/auth/user`, user)
}
async function jwtRequest(token: FieldValues) {
    return axios.post(`/auth/jwt`, token)
}

export const AuthProvider = ({ children }: { children: ReactNode }) => {
    const [isAuthenticated, setIsAuthenticated] = useState(false)
    const [loading, setLoading] = useState(true)
    const [isModalOpen, setIsModalOpen] = useState(false)

    useEffect(() => {
        const jwtReq = async () => {
            const cookies = Cookies.get()
            if (!cookies["x-jwt-token"]) {
                setIsAuthenticated(false)
                setLoading(false)
                return
            }
            try {
                console.log("cookie:", cookies["x-jwt-token"])
                const res = await jwtRequest({ "x-jwt-token": cookies["x-jwt-token"] })

                if (!res.data) {
                    console.log("error res", res)
                    setIsAuthenticated(false)
                    setLoading(false)
                    return
                }

                setIsAuthenticated(true)
                console.log("res.data:", res.data)

            } catch (error) {

                setIsAuthenticated(false)
                setLoading(false)
                console.log(isAuthenticated, loading)
                console.log("catch err:", error)
            }

        }
        jwtReq()


    }, [])
    // useEffect(() => {
    //     if (user) {

    //         setLoading(false)
    //         console.log("user:", user)

    //     }
    // }, [user])

    async function signup(user: FieldValues) {
        try {
            await registerRequest(user)
            setIsAuthenticated(true)
        }
        catch (error) {
            console.log(error)
        }
    }

    const signin = async (user: FieldValues) => {
        try {
            await signinRequest(user)
            setIsAuthenticated(true)
        } catch (error) {
            console.log(error)
        }
    }


    return <AuthContext.Provider
        value={{
            signup,
            signin,
            loading,
            isAuthenticated,
            isModalOpen,
            setIsModalOpen
        }}
    >
        {children}
    </AuthContext.Provider >
}