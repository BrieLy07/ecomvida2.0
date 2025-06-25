const factory = require('../services/factory');

const loginUsuario = async (req, res) => {
  try {
    const resultado = await factory.crearServicio('login').ejecutar(req.body);
    res.status(200).json(resultado);
  } catch (error) {
    res.status(401).json({ error: error.message });
  }
};

module.exports = loginUsuario;
