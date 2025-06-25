const factory = require('../services/factory');

const enviarTokenRecuperacion = async (req, res) => {
  try {
    const resultado = await factory.crearServicio('forgot').ejecutar(req.body);
    res.status(200).json(resultado);
  } catch (error) {
    res.status(400).json({ error: error.message });
  }
};

module.exports = enviarTokenRecuperacion;
