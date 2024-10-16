export default function Footer() {
    return (
        <footer className="bg-white w-full shadow dark:bg-gray-900">
            <div className="w-full max-w-screen-xl mx-auto p-3 md:py-8">
                <div className="sm:flex sm:items-center sm:justify-between">
                    <a href="./#" className="flex items-center mb-4 sm:mb-0 space-x-3 rtl:space-x-reverse">
                        <img src="./LOGORASHD.png" className="h-10 mr-3" alt="RAS logo" />
                        <span className="self-center text-2xl font-semibold whitespace-nowrap dark:text-white">Rosario Autos Sport</span>
                    </a>
                    <ul className="flex flex-wrap items-center mb-6 text-sm font-medium text-gray-500 sm:mb-0 dark:text-gray-400">
                        <li>
                            <a href="#" className="hover:underline me-4 md:me-6">Acerca de</a>
                        </li>

                        <li>
                            <a href="#" className="hover:underline me-4 md:me-6">Contacto</a>
                        </li>
                        <li>
                            <a href="#" className="hover:underline me-4 md:me-6">Políticas de privacidad</a>
                        </li>
                        <li>
                            <a href="#" className="hover:underline">Licensing</a>
                        </li>
                    </ul>
                </div>
                <hr className="my-6 border-gray-200 sm:mx-auto dark:border-gray-700 lg:my-8" />
                <span className="block text-sm text-gray-500 sm:text-center dark:text-gray-400">© 2024 <a href="/" className="hover:underline">Rosario Autos Sport</a>. All Rights Reserved.</span>
            </div>
        </footer>
    )
}
