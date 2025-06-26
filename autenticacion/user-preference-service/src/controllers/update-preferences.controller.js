const preferencesService = require('../services/preferences.service');

const actualizarPreferencias = async (req, res) => {
  try {
    const { id } = req.params;
    const data = req.body;
    const resultado = await preferencesService.actualizar(id, data);

    res.status(200).json(resultado);
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
};

module.exports = actualizarPreferencias;
