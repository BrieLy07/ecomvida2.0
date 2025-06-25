const factory = require('../services/factory');

const cerrarSesion = async (req, res) => {
  try {
    const token = req.headers.authorization?.split(' ')[1];
    if (!token) throw new Error('Token no enviado');

    await factory.crearServicio('logout').ejecutar({ token });
    res.status(200).json({ mensaje: 'Sesión cerrada correctamente' });
  } catch (error) {
    res.status(400).json({ error: error.message });
  }
};

module.exports = cerrarSesion;
