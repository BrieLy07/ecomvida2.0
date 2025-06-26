const preferencesService = require('../services/preferences.service');

const obtenerPreferencias = async (req, res) => {
  try {
    const { id } = req.params;
    const resultado = await preferencesService.obtener(id);

    if (!resultado) {
      return res.status(404).json({ mensaje: 'Preferencias no encontradas' });
    }

    res.status(200).json(resultado);
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
};

module.exports = obtenerPreferencias;
