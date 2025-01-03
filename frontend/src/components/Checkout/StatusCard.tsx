import { CheckCircle, XCircle, Clock, DollarSign, Calendar, FileText, X } from 'lucide-react'
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { motion, AnimatePresence } from 'framer-motion'
import { useEffect, useState } from 'react'
import dayjs from 'dayjs'

type SubscriptionStatus = 'authorized' | 'cancelled' | 'pending'

interface SubProps {
    props: SubscriptionProps | undefined
}

export type SubscriptionProps = {
    status: SubscriptionStatus
    reason: string
    date_created: string
    transaction_amount: number
}

export function SubscriptionStatus({ props }: SubProps) {
    const [showStatus, setShowStatus] = useState(false)

    const parsedDateCreated = dayjs(props?.date_created);

    useEffect(() => {
        if (props != undefined) {
            setShowStatus(true)
            console.log("Props not undefined", props)
        }
    }, [props])

    const handleCloseStatus = () => {
        setShowStatus(false)
    }

    return (
        <div className="min-h-screen bg-gray-100 flex flex-col items-center justify-center p-4">
            <AnimatePresence>
                {showStatus && (
                    <motion.div
                        className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4"
                        initial={{ opacity: 0 }}
                        animate={{ opacity: 1 }}
                        exit={{ opacity: 0 }}
                    >
                        <motion.div
                            className="bg-white rounded-lg shadow-xl max-w-2xl w-full"
                            initial={{ scale: 0.9, opacity: 0 }}
                            animate={{ scale: 1, opacity: 1 }}
                            exit={{ scale: 0.9, opacity: 0 }}
                            transition={{ type: 'spring', damping: 25, stiffness: 300 }}
                        >
                            {props ? (
                                <StatusCard
                                    status={props.status}
                                    reason={props.reason}
                                    amount={props.transaction_amount.toString()}
                                    date={parsedDateCreated.format('DD/MM/YYYY')}
                                    onClose={handleCloseStatus}
                                />
                            ) : null}
                        </motion.div>
                    </motion.div>
                )}
            </AnimatePresence>
        </div>
    )
}

function StatusCard({
    status,
    reason,
    amount,
    date,
    onClose
}: {
    status: SubscriptionStatus
    reason: string
    amount: string
    date: string
    onClose: () => void
}) {
    const statusConfig = {
        authorized: {
            title: "Suscripción Activa",
            description: "Tu suscripción está activa y al día. ¡Gracias por tu preferencia!",
            icon: <CheckCircle className="text-green-500" size={36} />,
            color: "border-t-8 border-green-500"
        },
        cancelled: {
            title: "Suscripción Cancelada",
            description: "Tu suscripción ha sido cancelada. Esperamos que vuelvas pronto.",
            icon: <XCircle className="text-red-500" size={36} />,
            color: "border-t-8 border-red-500"
        },
        pending: {
            title: "Suscripción Pendiente",
            description: "Tu suscripción está en proceso de activación. Te notificaremos cuando esté lista.",
            icon: <Clock className="text-yellow-500" size={36} />,
            color: "border-t-8 border-yellow-500"
        }
    }

    const config = statusConfig[status]

    return (
        <Card className={`w-full ${config.color}`}>
            <CardHeader>
                <div className="flex justify-between items-start">
                    <div className="flex items-center">
                        {config.icon}
                        <CardTitle className="text-3xl ml-3">{config.title}</CardTitle>
                    </div>
                    <Button variant="ghost" size="icon" onClick={onClose} aria-label="Cerrar">
                        <X size={24} />
                    </Button>
                </div>
            </CardHeader>
            <CardContent>
                <CardDescription className="text-xl mb-8">{config.description}</CardDescription>
                <div className="space-y-4 mb-8">
                    <InfoRow icon={<FileText size={24} />} label="Tipo de Suscripción" value={reason} />
                    <InfoRow icon={<DollarSign size={24} />} label="Monto" value={amount} />
                    <InfoRow icon={<Calendar size={24} />} label="Fecha" value={date} />
                </div>
            </CardContent>
            <CardFooter>
                <Button className="w-full" onClick={onClose}>
                    Entendido
                </Button>
            </CardFooter>
        </Card>
    )
}

function InfoRow({ icon, label, value }: { icon: React.ReactNode; label: string; value: string }) {
    return (
        <div className="flex items-center text-lg">
            <span className="text-gray-400 mr-3">{icon}</span>
            <span className="font-medium mr-2">{label}:</span>
            <span className="font-bold">{value}</span>
        </div>
    )
}

