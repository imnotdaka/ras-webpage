import { initMercadoPago, CardPayment,  } from '@mercadopago/sdk-react';
import axios from '../../context/axios';
import { Plan } from '@/pages/MembershipPage'
// import axios from "../../context/axios";

interface CheckPrep {
    plan: Plan
}

interface Request {
    card_token_id: string,
    preapproval_plan_id: string,
    payer_email: string | any,
    reason: string
}

interface ErrorInterface {
    type: 'non_critical' | 'critical';
    cause: string;
    message: string;
}

initMercadoPago('APP_USR-f922c249-0a13-4b0e-9aef-274fc6493fd9');

function Checkout({ plan }: CheckPrep) {

    const err = ({type, cause, message}: ErrorInterface) => {
        console.log("onerror:", type, cause, message)
    }

    return (
            <CardPayment
                initialization={{ amount: 20 }}
                onSubmit={async (cardFormData) => {
                    console.log(cardFormData)

                    const reqData: Request = {
                        card_token_id: cardFormData.token,
                        preapproval_plan_id: plan.id,
                        payer_email: cardFormData.payer.email,
                        reason: plan.reason,
                    }
                    console.log(reqData)
                    axios.post('create_suscription',
                        {...reqData},

                        {
                            headers: {
                                "Content-Type": "application/json",
                            },
                        })
                        .then((response) => {
                            console.log(response)
                        })
                        .catch((error) => {
                            console.log(error)
                        })
                }}
                onError={err}
            />
    )
}

export default Checkout
