import { useState } from "react";
import { useAuth } from "../context/AuthContext";
import { login } from "../services/authService";
import { useNavigate } from "react-router-dom";

export default function Login() {
  const { iniciarSesion } = useAuth();
  const [usuario, setUsuario] = useState("");
  const [clave, setClave] = useState("");
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const manejarLogin = async (e) => {
    e.preventDefault();
    try {
      const data = await login({ usuario, clave });
      iniciarSesion(data);
      navigate("/perfil");
    } catch (err) {
      setError(err);
    }
  };

  return (
    <div className="max-w-md mx-auto mt-10 bg-white p-6 shadow">
      <h2 className="text-2xl font-bold mb-4 text-center">Iniciar Sesión</h2>
      {error && <div className="text-red-600 mb-2 text-center">{error}</div>}
      <form onSubmit={manejarLogin} className="space-y-4">
        <input
          type="text"
          placeholder="Usuario"
          value={usuario}
          onChange={(e) => setUsuario(e.target.value)}
          className="w-full border p-2"
        />
        <input
          type="password"
          placeholder="Contraseña"
          value={clave}
          onChange={(e) => setClave(e.target.value)}
          className="w-full border p-2"
        />
        <button type="submit" className="w-full bg-green-600 text-white p-2">
          Entrar
        </button>
      </form>
    </div>
  );
}
