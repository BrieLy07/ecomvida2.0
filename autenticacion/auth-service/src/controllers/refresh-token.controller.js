const factory = require('../services/factory');

const refrescarToken = async (req, res) => {
  try {
    const { refreshToken } = req.body;
    const resultado = await factory.crearServicio('refresh').ejecutar({ refreshToken });
    res.status(200).json(resultado);
  } catch (error) {
    res.status(403).json({ error: error.message });
  }
};

module.exports = refrescarToken;
