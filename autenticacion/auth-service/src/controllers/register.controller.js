const factory = require('../services/factory');

const registrarUsuario = async (req, res) => {
  try {
    const resultado = await factory.crearServicio('registro').ejecutar(req.body);
    res.status(201).json({ mensaje: 'Usuario registrado con éxito', usuario: resultado });
  } catch (error) {
    res.status(400).json({ error: error.message });
  }
};

module.exports = registrarUsuario;
