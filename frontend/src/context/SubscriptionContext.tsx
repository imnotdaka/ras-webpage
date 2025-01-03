import { createContext, ReactNode, useContext, useEffect, useState } from "react"
import { useAuth } from "@/context/AuthContext"
import dayjs, { Dayjs } from "dayjs";
import axios from "./axios";

interface SubscriptionContextType {
    subscription: Subscription | undefined
    isSubLoading: boolean
    parsedDateCreated: Dayjs
    parsedNextPayment: Dayjs
    cancelSubscription: () => Promise<void>;
}

type Subscription = {
    reason: string
    status: string
    date_created: string
    next_payment_date: string
}

export const SubscriptionContext = createContext<SubscriptionContextType | null>(null)

export const useSubscription = () => {
    const context = useContext(SubscriptionContext)
    if (!context) {
        throw new Error("useSubscription must be used within an SubscriptionProvider")
    }
    return context
}

const cancelSubscriptionRequest = async () => {
    return axios.put(`/cancel_subscription`)
}

export const SubscriptionProvider = ({ children }: { children: ReactNode }) => {
    const [subscription, setSubscription] = useState<Subscription>()
    const [isSubLoading, setIsSubLoading] = useState(true)

    const { setToken, loading, GetSubscription } = useAuth()

    const cancelSubscription = async () => {
        await cancelSubscriptionRequest()
        setSubscription(undefined)
        setToken("")
    }

    useEffect(() => {
        const run = async () => {
            if (!loading) {
                const res = await GetSubscription()
                setSubscription(res)
                setIsSubLoading(false)
            }
        }
        run()
    }, [loading])

    const parsedDateCreated = dayjs(subscription?.date_created);
    const parsedNextPayment = dayjs(subscription?.next_payment_date);

    return (
        <SubscriptionContext.Provider
            value={{
                subscription,
                isSubLoading,
                parsedDateCreated,
                parsedNextPayment,
                cancelSubscription
            }}
        >
            {children}
        </SubscriptionContext.Provider>
    )
}