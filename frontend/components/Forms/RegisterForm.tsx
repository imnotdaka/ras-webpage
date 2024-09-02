export default function Register() {
    return (
        <div className="max-w-lg mx-auto  bg-white dark:bg-gray-800 shadow-md px-8 py-4 flex flex-col items-center">
            <form action="#" className="w-full flex flex-col gap-2">
                <div className="flex items-start flex-col justify-start">
                    <label htmlFor="firstName" className="text-sm text-gray-700 dark:text-gray-200 mr-2">Nombre</label>
                    <input type="text" id="firstName" name="firstName" className="w-full px-3 dark:text-gray-200 dark:bg-gray-900 py-2 rounded-md border border-gray-300 dark:border-gray-700 focus:outline-none focus:ring-1 focus:ring-blue-500" />
                </div>

                <div className="flex items-start flex-col justify-start">
                    <label htmlFor="lastName" className="text-sm text-gray-700 dark:text-gray-200 mr-2">Apellido:</label>
                    <input type="text" id="lastName" name="lastName" className="w-full px-3 dark:text-gray-200 dark:bg-gray-900 py-2 rounded-md border border-gray-300 dark:border-gray-700 focus:outline-none focus:ring-1 focus:ring-blue-500" />
                </div>

                <div className="flex items-start flex-col justify-start">
                    <label htmlFor="username" className="text-sm text-gray-700 dark:text-gray-200 mr-2">DNI:</label>
                    <input type="text" id="username" name="username" className="w-full px-3 dark:text-gray-200 dark:bg-gray-900 py-2 rounded-md border border-gray-300 dark:border-gray-700 focus:outline-none focus:ring-1 focus:ring-blue-500" />
                </div>

                <div className="flex items-start flex-col justify-start">
                    <label htmlFor="email" className="text-sm text-gray-700 dark:text-gray-200 mr-2">Correo electrónico:</label>
                    <input type="email" id="email" name="email" className="w-full px-3 dark:text-gray-200 dark:bg-gray-900 py-2 rounded-md border border-gray-300 dark:border-gray-700 focus:outline-none focus:ring-1 focus:ring-blue-500" />
                </div>

                <div className="flex items-start flex-col justify-start">
                    <label htmlFor="password" className="text-sm text-gray-700 dark:text-gray-200 mr-2">Contraseña:</label>
                    <input type="password" id="password" name="password" className="w-full px-3 dark:text-gray-200 dark:bg-gray-900 py-2 rounded-md border border-gray-300 dark:border-gray-700 focus:outline-none focus:ring-1 focus:ring-blue-500" />
                </div>

                <div className="flex items-start flex-col justify-start">
                    <label htmlFor="confirmPassword" className="text-sm text-gray-700 dark:text-gray-200 mr-2">Confirmar contraseña:</label>
                    <input type="password" id="confirmPassword" name="confirmPassword" className="w-full px-3 dark:text-gray-200 dark:bg-gray-900 py-2 rounded-md border border-gray-300 dark:border-gray-700 focus:outline-none focus:ring-1 focus:ring-blue-500" />
                </div>

                <button type="submit" className="bg-blue-500 hover:bg-blue-600 text-white font-medium py-2 px-4 rounded-md shadow-sm">Registrar</button>
            </form>

            <div className="mt-4 text-center">
                <span className="text-sm text-gray-500 dark:text-gray-300">Ya tenes una cuenta? </span>
                <a href="./#" className="text-blue-500 hover:text-blue-600 text-sm">Iniciar sesión</a>
            </div>
        </div >
    )
}
