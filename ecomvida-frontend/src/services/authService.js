/* global process */
import axios from "axios";

// URL dinámica desde variable de entorno
const API = process.env.REACT_APP_AUTH_URL;

export const login = async (datos) => {
  try {
    const res = await axios.post(`${API}/login`, datos);
    return res.data;
  } catch (error) {
    throw error.response?.data?.mensaje || "Error al iniciar sesión";
  }
};

export const register = async (datos) => {
  try {
    const res = await axios.post(`${API}/register`, datos);
    return res.data;
  } catch (error) {
    throw error.response?.data?.mensaje || "Error al registrarse";
  }
};
