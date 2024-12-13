import { StatusScreen } from "@mercadopago/sdk-react";

interface Init {
    id: string
}

interface ErrorInterface {
    type: 'non_critical' | 'critical';
    cause: string;
    message: string;
}

function AfterPay({id}: Init) {

    const err = ({ type, cause, message }: ErrorInterface) => {
        console.log("onerror:", type, cause, message)
    }

    const initialization = {
        paymentId: id
    }

    return (
        <StatusScreen
            initialization={initialization} 
            onError={err}/>
    )
}

export default AfterPay
