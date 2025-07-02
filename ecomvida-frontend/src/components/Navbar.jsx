import { Link } from "react-router-dom";

export default function Navbar() {
  return (
    <nav className="bg-white shadow p-4 flex justify-between">
      <div className="text-xl font-bold text-green-600">ECOMVIDA</div>
      <div className="space-x-4">
        <Link to="/" className="text-gray-700 hover:text-green-600">Inicio</Link>
        <Link to="/login" className="text-gray-700 hover:text-green-600">Login</Link>
        <Link to="/register" className="text-gray-700 hover:text-green-600">Registro</Link>
        <Link to="/perfil" className="text-gray-700 hover:text-green-600">Perfil</Link>
      </div>
    </nav>
  );
}
