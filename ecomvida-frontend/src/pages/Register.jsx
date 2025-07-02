import { useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";

export default function Register() {
  const [datos, setDatos] = useState({
    nombre: "",
    apellido: "",
    usuario: "",
    correo: "",
    numero: "",
    clave: ""
  });
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const manejarCambio = (e) => {
    setDatos({ ...datos, [e.target.name]: e.target.value });
  };

  const manejarSubmit = async (e) => {
    e.preventDefault();
    try {
      await axios.post("http://<IP-EC2>:3001/register", datos); // reemplaza <IP-EC2>
      navigate("/login");
    } catch (err) {
      setError(err.response?.data?.mensaje || "Error al registrar");
    }
  };

  return (
    <div className="max-w-md mx-auto mt-10 bg-white p-6 shadow">
      <h2 className="text-2xl font-bold mb-4 text-center">Registro</h2>
      {error && <div className="text-red-600 mb-2 text-center">{error}</div>}
      <form onSubmit={manejarSubmit} className="space-y-3">
        <input name="nombre" placeholder="Nombre" onChange={manejarCambio} className="w-full border p-2" />
        <input name="apellido" placeholder="Apellido" onChange={manejarCambio} className="w-full border p-2" />
        <input name="usuario" placeholder="Usuario" onChange={manejarCambio} className="w-full border p-2" />
        <input name="correo" type="email" placeholder="Correo" onChange={manejarCambio} className="w-full border p-2" />
        <input name="numero" placeholder="Número" onChange={manejarCambio} className="w-full border p-2" />
        <input name="clave" type="password" placeholder="Contraseña" onChange={manejarCambio} className="w-full border p-2" />
        <button type="submit" className="w-full bg-green-600 text-white p-2">
          Registrarse
        </button>
      </form>
    </div>
  );
}
