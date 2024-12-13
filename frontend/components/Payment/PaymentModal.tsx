import { Dialog, DialogContent } from '@/components/ui/dialog'
// import { Input } from '@/components/ui/input'
// import { Label } from '@/components/ui/label'
import { Plan } from '@/pages/MembershipPage'
import Checkout from '../Checkout/Checkout'

interface PaymentModalProps {
  isOpen: boolean
  onClose: () => void
  plan: Plan
}

export function PaymentModal({ isOpen, onClose, plan }: PaymentModalProps) {

  return (
    <Dialog open={isOpen} onOpenChange={onClose}>
      <DialogContent>
        <Checkout plan={plan}/>
      </DialogContent>
    </Dialog>
  )
}

