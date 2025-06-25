const express = require('express');
const registrarUsuario = require('../controllers/register.controller');
const loginUsuario = require('../controllers/login.controller');
const refrescarToken = require('../controllers/refresh-token.controller');
const cerrarSesion = require('../controllers/logout.controller');
const enviarTokenRecuperacion = require('../controllers/forgot-password.controller');
const cambiarContrasena = require('../controllers/reset-password.controller');

const router = express.Router();

router.post('/register', registrarUsuario);
router.post('/login', loginUsuario);
router.post('/refresh-token', refrescarToken);
router.post('/logout', cerrarSesion);
router.post('/forgot-password', enviarTokenRecuperacion);
router.post('/reset-password', cambiarContrasena);

module.exports = router;
