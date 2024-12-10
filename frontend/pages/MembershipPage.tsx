import { PaymentButton } from "@/components/Payment/PaymentButton";
import Footer from "../components/Footer";
import NavBar from "../components/NavBar";
import axios from "../context/axios";
import { useEffect, useState } from "react"
import Checkout from "@/components/Checkout/Checkout";


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

    // const submit = async () => {
    //     axios.post('create_suscription',
    //         {
    //             headers: {
    //                 "Content-Type": "application/json",
    //             },
    //             body: JSON.stringify({
    //                 "token": "7e718c632edc536955cdf1a0d9a99365",
    //                 "issuer_id": "310",
    //                 "payment_method_id": "visa",
    //                 "transaction_amount": 20,
    //                 "installments": 1,
    //                 "payer": {
    //                     "email": "uwu@uwu.com",
    //                     "identification": {
    //                         "type": "DNI",
    //                         "number": "12345678"
    //                     }
    //                 }
    //             })
    //         })
    //         .then((response) => {
    //             console.log(response)
    //         })
    //         .catch((error) => {
    //             console.log(error)
    //         })

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
                        {/* <button className="bg-neutral-800 p-10" /> */}
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
                    {/* <div>{plans.map(plan => (




                    <a href={"https://www.mercadopago.com.ar/subscriptions/checkout?preapproval_plan_id=" + plan.id} key={plan.id} className="block max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700">

                        <h5 className="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">{plan.reason}</h5>
                        <p className="font-normal text-gray-700 dark:text-gray-400">{plan.auto_recurring.frequency} {plan.auto_recurring.frequency_type} - ${plan.auto_recurring.transaction_amount}</p>
                    </a>
                ))}
                </div> */}
                    <Footer />
                </div >
            )
        }
    }