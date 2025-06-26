const express = require('express');
const obtenerPreferencias = require('../controllers/get-preferences.controller');
const actualizarPreferencias = require('../controllers/update-preferences.controller');

const router = express.Router();

router.get('/users/:id/preferences', obtenerPreferencias);
router.put('/users/:id/preferences', actualizarPreferencias);

module.exports = router;
