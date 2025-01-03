import { Button } from "@/components/ui/button"

export default function Home() {
    return (
        <div className="min-h-screen bg-white font-roboto">
            <main>
                {/* Hero Section */}
                <section className="relative h-[70vh] bg-cover bg-center" style={{ backgroundImage: "url('../../LOGORASHD.png')" }}>
                    <div className="absolute inset-0 bg-black bg-opacity-70 flex items-center justify-center">
                        <div className="text-center text-white">
                            <h1 className="text-5xl font-bold mb-4">Rosario Autos Sport</h1>
                            <p className="text-xl mb-8">Donde la pasión por los autos clásicos se encuentra con la modernidad</p>
                            {/* <Button size="lg" className="bg-white text-black hover:bg-gray-200">Descubre Más</Button> */}
                        </div>
                    </div>
                </section>

                {/* About Section */}
                <section className="py-16 bg-gray-100">
                    <div className="container mx-auto px-4">
                        <h2 className="text-3xl font-bold mb-8 text-center text-black">Sobre Nosotros</h2>
                        <p className="text-lg text-center max-w-3xl mx-auto mb-12 text-gray-700">
                            Rosario Autos Sport es un grupo apasionado de entusiastas de los autos que celebra la rica historia del automovilismo mientras abraza las innovaciones modernas. Organizamos eventos, exhibiciones y recorridos que unen a amantes de autos de todas las épocas.
                        </p>
                    </div>
                </section>

                {/* Gallery Section */}
                {/* <section className="py-16 bg-white">
                    <div className="container mx-auto px-4">
                        <h2 className="text-3xl font-bold mb-8 text-center text-black">Galería de Eventos</h2>
                        <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
                            {[1, 2, 3, 4, 5, 6, 7, 8].map((i) => (
                                <div key={i} className="relative h-64">
                                    <img
                                        src={`/placeholder.jpg`}
                                        alt={`Evento ${i}`}
                                        className="w-full h-full object-cover rounded-lg"
                                    />
                                </div>
                            ))}
                        </div>
                    </div>
                </section> */}
            </main>
        </div>
    )
}
