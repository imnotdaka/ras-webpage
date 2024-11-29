import Footer from "../components/Footer";
import NavBar from "../components/NavBar";
import axios from "../context/axios";
import { useEffect, useState } from "react"


interface AutoRecurring {
    frequency: number;
    frequency_type: string;
    transaction_amount: number;
}

interface Plan {
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
                <div>{plans.map(plan => (




                    <a href={"https://www.mercadopago.com.ar/subscriptions/checkout?preapproval_plan_id=" + plan.id} key={plan.id} className="block max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700">

                        <h5 className="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">{plan.reason}</h5>
                        <p className="font-normal text-gray-700 dark:text-gray-400">{plan.auto_recurring.frequency} {plan.auto_recurring.frequency_type} - ${plan.auto_recurring.transaction_amount}</p>
                    </a>
                ))}
                </div>
                <Footer />
            </div >
        )
    }
}

