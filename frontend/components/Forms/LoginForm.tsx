import { useForm } from "react-hook-form"
import { useAuth } from "../../context/AuthContext"
import { useNavigate } from "react-router-dom"
import { useEffect } from "react"

export default function Login() {
    const {
        register,
        handleSubmit,
        formState: { errors } 
    } = useForm()
    const { signin, isAuthenticated } = useAuth()

    const onSubmit = handleSubmit(async data => {
        console.log(data)
        await signin(data)
    })
    const navigate = useNavigate()
    useEffect(() => {
        if (isAuthenticated) navigate("/membership")

    }, [isAuthenticated])

    return (
        <div className="flex items-center justify-center mt-12 w-full dark:bg-gray-950 h-full">
            <div className="bg-white dark:bg-gray-900 px-8 py-5 max-w-md">
                <h1 className="text-2xl font-bold text-center mb-4 dark:text-gray-200">Iniciar sesión</h1>
                <form onSubmit={onSubmit}>
                    <div className="mb-4">
                        <label htmlFor="email" className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">Correo electrónico</label>
                        <input type="email" id="email" className="w-full px-3 dark:text-gray-200 dark:bg-gray-900 py-2 rounded-md border border-gray-300 dark:border-gray-700 focus:outline-none focus:ring-1 focus:ring-blue-500" placeholder="tu@email.com" required {...register("email", { required: true })} />
                        {errors.email && <p className="text-red-500">Correo requerido</p>}
                    </div>
                    <div className="mb-4">
                        <label htmlFor="password" className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">Contraseña</label>
                        <input type="password" id="password" className="w-full px-3 dark:text-gray-200 dark:bg-gray-900 py-2 rounded-md border border-gray-300 dark:border-gray-700 focus:outline-none focus:ring-1 focus:ring-blue-500" placeholder="Contraseña" required {...register("password", { required: true })} />
                        {errors.password && <p className="text-red-500">Contraseña requerida</p>}
                        {/* <a href="#"
                            className="text-xs text-gray-600 hover:text-indigo-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">¿Olvidaste tu contraseña?</a> */}
                    </div>
                    <div className="flex items-center justify-between mb-4">
                        {/* <div className="flex items-center">
                            <input type="checkbox" id="remember" className="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500 focus:outline-none" checked />
                            <label htmlFor="remember" className="ml-2 block text-sm text-gray-700 dark:text-gray-300">Iniciar sesión automáticamente</label>
                        </div> */}
                        <a href="./register"
                            className="text-xs text-indigo-500 hover:text-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">Crear cuenta</a>
                    </div>
                    <button type="submit" className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">Iniciar sesión</button>
                </form>
            </div>
        </div>
    )
}
