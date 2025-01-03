import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Separator } from "@/components/ui/separator"
import { useAuth } from "@/context/AuthContext"
import { useSubscription } from "@/context/SubscriptionContext"
import { Button } from "./ui/button"




export default function Profile() {

  const {
    subscription,
    isSubLoading,
    parsedDateCreated,
    parsedNextPayment,
    cancelSubscription
  } = useSubscription()

  const { user } = useAuth()

  return (
    <div className="min-h-screen bg-gray-100 py-3 px-4">
      <div className="container mx-auto max-w-4xl">
        <h1 className="text-3xl font-bold mb-2 ml-2">Perfil</h1>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          {/* Información Personal */}
          <Card className="border-zinc-400">
            <CardHeader>
              <CardTitle className="text-xl">Información Personal</CardTitle>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                <div>
                  <h3 className="font-semibold text-sm text-gray-500">Nombre Completo</h3>
                  <p className="text-lg">{user.first_name} {user.last_name}</p>
                </div>
                <div>
                  <h3 className="font-semibold text-sm text-gray-500">Email</h3>
                  <p className="text-lg">{user.email}</p>
                </div>
              </div>
            </CardContent>
          </Card>

          {/* Información de Suscripción */}
          {user.subscribed && !isSubLoading && subscription && (
            <Card className="border-zinc-400">
              <CardHeader>
                <CardTitle className="text-xl">Información de Suscripción</CardTitle>
              </CardHeader>
              <CardContent>
                <div className="space-y-4">
                  <div>
                    <h3 className="font-semibold text-sm text-gray-500">Plan</h3>
                    <p className="text-lg">{subscription.reason}</p>
                  </div>
                  <div>
                    <h3 className="font-semibold text-sm text-gray-500">Estado</h3>
                    <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-sm font-medium bg-green-100 text-green-800">
                      {subscription.status}
                    </span>
                  </div>
                  <Separator />
                  <div>
                    <h3 className="font-semibold text-sm text-gray-500">Fecha de Inicio</h3>
                    <p className="text-lg">{parsedDateCreated.format("DD/MM/YYYY")}</p>
                  </div>
                  <div>
                    <h3 className="font-semibold text-sm text-gray-500">Próximo Pago</h3>
                    <p className="text-lg">{parsedNextPayment.format("DD/MM/YYYY")}</p>
                  </div>
                  <div className="pt-4">
                    <Button
                      variant="destructive"
                      className="w-full"
                      onClick={() => {
                        if (confirm('¿Estás seguro de que deseas cancelar tu suscripción?')) {
                          cancelSubscription()
                        }
                      }}
                    >
                      Cancelar Suscripción
                    </Button>
                  </div>
                </div>
              </CardContent>
            </Card>)}
        </div>
      </div>
    </div>
  )
}