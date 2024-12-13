import { initMercadoPago, CardPayment } from '@mercadopago/sdk-react';
import axios from '../../context/axios';
import { Plan } from '@/pages/MembershipPage'
import { useRef, useState } from 'react';
import { SubscriptionProps, SubscriptionStatus } from './StatusCard'
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

export interface SuscriptionID {
    suscription_id: string | undefined
}

initMercadoPago('APP_USR-f922c249-0a13-4b0e-9aef-274fc6493fd9');

function Checkout({ plan }: CheckPrep) {
    const [data, setData] = useState<SubscriptionProps>()
    const [isPaymentSuccessful, setIsPaymentSuccessful] = useState(false)

    const formRef = useRef<HTMLFormElement>(null);

    const err = ({ type, cause, message }: ErrorInterface) => {
        console.log("onerror:", type, cause, message)
    }

    return (
        <form ref={formRef}>
            <CardPayment
                customization={{ paymentMethods: { maxInstallments: 1, minInstallments: 1 } }}
                initialization={{ amount: 20 }}
                locale='es-AR'
                onSubmit={async (cardFormData) => {
                    console.log(cardFormData)

                    const reqData: Request = {
                        card_token_id: cardFormData.token,
                        preapproval_plan_id: plan.id,
                        payer_email: cardFormData.payer.email,
                        reason: plan.reason,
                    }
                    console.log(reqData)
                    axios.post<SubscriptionProps>('create_suscription',
                        { ...reqData },

                        {
                            headers: {
                                "Content-Type": "application/json",
                            },
                        })
                        .then((res) => {
                            window.cardPaymentBrickController.unmount()
                            // setIsModalOpen(false)
                            setData(res.data)
                            setIsPaymentSuccessful(true)
                            console.log("res.data sub:", res.data)
                        })
                        .catch((error) => {
                            console.log(error)
                        })
                }}
                onError={err}
            />
            {isPaymentSuccessful &&
                <SubscriptionStatus props={data} />
            }

        </form>
    )
}

export default Checkout
