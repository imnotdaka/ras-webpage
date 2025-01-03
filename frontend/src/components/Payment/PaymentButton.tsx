import { PaymentModal } from './PaymentModal'
import { Plan } from '@/pages/MembershipPage'
import { useAuth } from '@/context/AuthContext'

interface PaymentButtonProps {
  plan: Plan
}

export function PaymentButton({ plan }: PaymentButtonProps) {
  const { isModalOpen, setIsModalOpen } = useAuth()


  return (
    <>
      <button
        onClick={() => setIsModalOpen(true)}
        className="w-full bg-blue-600 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded transition duration-300"
      >
        Suscribirse
      </button>
      {isModalOpen && (
        <PaymentModal 
          isOpen={isModalOpen} 
          onClose={() => setIsModalOpen(false)} 
          plan={plan}
        />
      )}
    </>
  )
}

