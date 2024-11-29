import { initMercadoPago, CardPayment } from '@mercadopago/sdk-react';
// import axios from "../../context/axios";

initMercadoPago('APP_USR-f922c249-0a13-4b0e-9aef-274fc6493fd9');

function Checkout() {
    return (
        <div>
            <CardPayment
                initialization={{ amount: 20 }}
                onSubmit={async (cardFormData) => {
                    console.log(cardFormData)
                    // axios.post('create_suscription',
                    //     {
                    //         headers: {
                    //             "Content-Type": "application/json",
                    //         },
                    //         body: JSON.stringify(cardFormData)
                    //     })
                    //     .then((response) => {
                    //         console.log(response)
                    //     })
                    //     .catch((error) => {
                    //         console.log(error)
                    //     })
                }}
            />
        </div>
    )
}

export default Checkout
