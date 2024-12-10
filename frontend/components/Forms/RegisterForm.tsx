import { useForm } from "react-hook-form"
import { useAuth } from "../../context/AuthContext"
import { useEffect } from "react"
import { useNavigate } from "react-router-dom"


export default function Register() {
    const {
        register,
        handleSubmit,
        formState: { errors }
    } = useForm()
    const { signup, isAuthenticated } = useAuth()
    const navigate = useNavigate()
    useEffect(() => {
        if (isAuthenticated) navigate("/membership")

    }, [isAuthenticated])


    return (
        <div className="max-w-lg mx-auto  bg-white  dark:bg-gray-900 shadow-md px-8 py-1 flex flex-col items-center">
            <div className="text-2xl text-gray-700 dark:text-gray-200 mr-2">Registro</div>
            <form className="w-full flex flex-col gap-2" onSubmit={handleSubmit(async (values) => {
                console.log(values)
                signup(values)
            })}>
                <div className="flex items-start flex-col justify-start">
                    <label htmlFor="firstName" className="text-sm text-gray-700 dark:text-gray-200 mr-2">Nombre</label>
                    <input type="text" id="firstName" className="w-full px-3 dark:text-gray-200 dark:bg-gray-900 py-2 rounded-md border border-gray-300 dark:border-gray-700 focus:outline-none focus:ring-1 focus:ring-blue-500" {...register("first_name", { required: true })} />
                    {errors.first_name && (<p className="text-red-500">Nombre requerido</p>)}
                </div>


                <div className="flex items-start flex-col justify-start">
                    <label htmlFor="lastName" className="text-sm text-gray-700 dark:text-gray-200 mr-2">Apellido:</label>
                    <input type="text" id="lastName" className="w-full px-3 dark:text-gray-200 dark:bg-gray-900 py-2 rounded-md border border-gray-300 dark:border-gray-700 focus:outline-none focus:ring-1 focus:ring-blue-500" required {...register("last_name", { required: true })} />
                    {errors.last_name && (<p className="text-red-500">Apellido requerido</p>)}
                </div>

                <div className="flex items-start flex-col justify-start">
                    <label htmlFor="email" className="text-sm text-gray-700 dark:text-gray-200 mr-2">Correo electrónico:</label>
                    <input type="email" id="email" className="w-full px-3 dark:text-gray-200 dark:bg-gray-900 py-2 rounded-md border border-gray-300 dark:border-gray-700 focus:outline-none focus:ring-1 focus:ring-blue-500" required {...register("email", { required: true })} />
                    {errors.email && (<p className="text-red-500">Correo requerido</p>)}
                </div>

                <div className="flex items-start flex-col justify-start">
                    <label htmlFor="password" className="text-sm text-gray-700 dark:text-gray-200 mr-2">Contraseña:</label>
                    <input type="password" id="password" className="w-full px-3 dark:text-gray-200 dark:bg-gray-900 py-2 rounded-md border border-gray-300 dark:border-gray-700 focus:outline-none focus:ring-1 focus:ring-blue-500" required {...register("password", { required: true })} />
                    {errors.password && (<p className="text-red-500">Contraseña requerida</p>)}
                </div>

                {/* <div className="flex items-start flex-col justify-start">
                    <label htmlFor="confirmPassword" className="text-sm text-gray-700 dark:text-gray-200 mr-2">Confirmar contraseña:</label>
                    <input type="password" id="confirm_password" className="w-full px-3 dark:text-gray-200 dark:bg-gray-900 py-2 rounded-md border border-gray-300 dark:border-gray-700 focus:outline-none focus:ring-1 focus:ring-blue-500" required/>
                    {errors.confirm_password && (<p className="text-red-500">Contraseña requerida</p>)}
                </div> */}

                <button type="submit" className="bg-blue-500 hover:bg-blue-600 text-white font-medium py-2 px-4 rounded-md shadow-sm">Registrar</button>
            </form>

            <div className="mt-4 text-center">
                <span className="text-sm text-gray-500 dark:text-gray-300">Ya tenes una cuenta? </span>
                <a href="./login" className="text-blue-500 hover:text-blue-600 text-sm">Iniciar sesión</a>
            </div>
        </div >
    )
}
