import { ChevronDown, Menu } from 'lucide-react'
import { Button } from "@/components/ui/button"
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import {
    Sheet,
    SheetContent,
    SheetTrigger,
} from "@/components/ui/sheet"
import { useAuth } from "@/context/AuthContext"
import { useNavigate } from 'react-router-dom'

export default function Navbar() {

    const navigate = useNavigate()

    const { user, isAuthenticated, logout } = useAuth()

    const NavItems = () => (
        <>
            {/* <a href="/eventos" className="hover:text-gray-300">Eventos</a>
            <a href="/galeria" className="hover:text-gray-300">Galería</a> */}

            {isAuthenticated ? (
                <DropdownMenu>
                    <DropdownMenuTrigger asChild>
                        <Button variant="ghost" className="text-white flex items-center">
                            <img
                                src="../../usuario.png"
                                alt="User Avatar"
                                className="w-8 h-8 rounded-full mr-2"
                            />
                            <ChevronDown className="h-4 w-4" />
                        </Button>
                    </DropdownMenuTrigger>
                    <DropdownMenuContent>
                        <DropdownMenuItem>
                            <a href="/profile" className="w-full">Ver Perfil</a>
                        </DropdownMenuItem>
                        <DropdownMenuItem>
                            <button onClick={() => {
                                logout()
                                navigate("/")
                            }} className="w-full text-left">
                                Cerrar Sesión
                            </button>
                        </DropdownMenuItem>
                    </DropdownMenuContent>
                </DropdownMenu>
            ) : (
                <>
                    <Button variant="ghost" className="text-white bg-black  hover:bg-gray-200 hover:text-black">
                        <a href="/login">
                            Iniciar Sesión
                        </a>
                    </Button>
                    <Button variant="outline" className="text-white bg-black border-hidden hover:bg-gray-200 hover:text-black">
                        <a href="/register">
                            Registrarse
                        </a>
                    </Button>
                </>
            )}{!user.subscribed && (
                <a href="/membership">
                    <Button variant="default" className="bg-black text-white hover:bg-gray-200 hover:text-black">
                        Hacerse Miembro
                    </Button>
                </a>
            )}
        </>
    )

    return (
        <nav className="bg-black text-white p-2">
            <div className="container mx-auto flex justify-between items-center">
                {/* Logo y nombre del grupo */}
                <div className="flex items-center space-x-2">
                    <img
                        src="../../LOGORASHD.png"
                        alt="Logo Rosario Autos Sport"
                        className="w-8 h-8 object-contain"
                    />
                    <a href="/" className="text-xl font-bold">
                        Rosario Autos Sport
                    </a>
                </div>

                {/* Menú de escritorio */}
                <div className="hidden md:flex items-center space-x-4">
                    <NavItems />
                </div>

                {/* Menú móvil */}
                <div className="md:hidden">
                    <Sheet>
                        <SheetTrigger asChild>
                            <Button variant="ghost" size="icon" className="text-white">
                                <Menu className="h-6 w-6" />
                                <span className="sr-only">Abrir menú</span>
                            </Button>
                        </SheetTrigger>
                        <SheetContent side="right" className="w-[300px] bg-black text-white">
                            {isAuthenticated ? (
                                <div className="flex flex-col items-center py-4 border-b border-gray-800">
                                    <img
                                        src="../../usuario.png"
                                        alt="User Avatar"
                                        className="w-16 h-16 rounded-full mb-2"
                                    />
                                    {!user.subscribed && (
                                        <a href="/membership" className="text-white hover:text-gray-300 py-2 mt-5">
                                            Hacerse Miembro
                                        </a>
                                    )}
                                    <a href="/profile" className="text-white hover:text-gray-300 py-2 mt-5">
                                        Ver Perfil
                                    </a>
                                    <button
                                        onClick={() => {
                                            logout()
                                            navigate("/")
                                        }} className="text-white hover:text-gray-300 py-2">
                                        Cerrar Sesión</button>
                                </div>
                            ) :
                                (<div className="flex flex-col items-center py-4 border-b border-gray-800">
                                    <a href="/login" className="text-white hover:text-gray-300 py-2 mt-5">
                                        Iniciar Sesión
                                    </a>
                                    <a href="/register" className="text-white hover:text-gray-300 py-2 mt-5">
                                        Registrarse
                                    </a>
                                </div>)}


                        </SheetContent>
                    </Sheet>
                </div>
            </div>
        </nav >
    )
}

