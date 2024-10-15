import axios from "axios";
import { useEffect, useState } from "react"

interface AutoRecurring {
    frequency: number;
    frequency_type: string;
}

interface Plan {
    id: string;
    reason: string;
    auto_recurring: AutoRecurring;
    transaction_amount: number;
}

export default function MembershipPage() {
    const [plans, setPlans] = useState<Plan[]>([])

    async function getPlans() {
        try {
            const res = await axios.get('/api/preapproval_plan/search', {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer APP_USR-96f4d433-0142-4d57-914e-0fa08414d273' //Public key doesnt work
                }
            });
            console.log(res.data.results)
            setPlans(res.data.results);
        } catch (error) {
            console.error("Error in getplans:", error);
        }
    }
    useEffect(() => {
      getPlans()
    }, [])
    



    return (
        <div>{plans.map(plan => (
            <div key={plan.id}>
                <h3>{plan.reason}</h3>
                <p>{plan.transaction_amount}</p>
                <p>{plan.auto_recurring.frequency}</p>
                <p>{plan.auto_recurring.frequency_type}</p>
            </div>
        ))}</div>
    )
}