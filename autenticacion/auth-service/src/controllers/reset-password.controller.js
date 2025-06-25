const factory = require('../services/factory');

const cambiarContrasena = async (req, res) => {
  try {
    const resultado = await factory.crearServicio('reset').ejecutar(req.body);
    res.status(200).json({ mensaje: 'Contraseña actualizada correctamente' });
  } catch (error) {
    res.status(400).json({ error: error.message });
  }
};

module.exports = cambiarContrasena;
