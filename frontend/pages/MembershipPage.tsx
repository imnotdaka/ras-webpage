import { PaymentButton } from "@/components/Payment/PaymentButton";
import Footer from "../components/Footer";
import NavBar from "../components/NavBar";
import axios from "../context/axios";
import { useEffect, useState } from "react"


interface AutoRecurring {
    frequency: number;
    frequency_type: string;
    transaction_amount: number;
}

export interface Plan {
    id: string;
    reason: string;
    auto_recurring: AutoRecurring;
}

export default function MembershipPage() {
    const [plans, setPlans] = useState<Plan[]>([])

        async function getPlans() {
            try {
                const res = await axios.get('/get_plans');
                console.log("getplans: ", res.data)
                setPlans(res.data);
            } catch (error) {
                console.error("Error in getplans:", error);
            }
        }
        useEffect(() => {
            getPlans()
        }, [])

        if (!plans) {
            return (
                <div>
                    <NavBar />
                    <h1>No hay planes de suscripcion por el momento.</h1>
                    <Footer />
                </div>
            )
        }
        else {
            return (
                <div>
                    <NavBar />

                    <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
                        {plans.map((plan) => (
                            <div key={plan.id} className="bg-white shadow-lg rounded-lg overflow-hidden flex flex-col">
                                <div className="p-6">
                                    <h2 className="text-2xl font-bold mb-2">{plan.reason}</h2>
                                    <p className="text-gray-600 mb-4">${plan.auto_recurring.transaction_amount} / {plan.auto_recurring.frequency} mes</p>
                                </div>
                                <div className="mt-auto p-6 bg-gray-50">
                                    <PaymentButton plan={plan} />
                                </div>
                            </div>
                        ))}
                    </div>
                    <Footer />
                </div >
            )
        }
    }