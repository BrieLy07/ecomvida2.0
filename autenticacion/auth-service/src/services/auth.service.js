const jwt = require('jsonwebtoken');
const bcrypt = require('bcrypt');
const User = require('../models/user.model');
const { agregarATokenBlacklist } = require('../utils/token.util');

class ServicioRegistro {
  async ejecutar(data) {
    const { nombre, apellido, usuario, correo, numero, contrasena } = data;

    // Validación simple
    if (!nombre || !apellido || !usuario || !correo || !numero || !contrasena) {
      throw new Error('Todos los campos son obligatorios');
    }

    // Verificar existencia
    const existente = await User.findOne({ where: { correo } });
    if (existente) {
      throw new Error('Ya existe un usuario con ese correo');
    }

    // Hash de la contraseña
    const hash = await bcrypt.hash(contrasena, 10);

    // Crear usuario
    const nuevoUsuario = await User.create({
      nombre,
      apellido,
      usuario,
      correo,
      numero,
      contrasena: hash,
    });

    // No retornar contraseña
    const { contrasena: _, ...usuarioSinClave } = nuevoUsuario.toJSON();
    return usuarioSinClave;
  }
}

class ServicioLogin {
  async ejecutar(data) {
    const { correo, contrasena } = data;

    if (!correo || !contrasena) {
      throw new Error('Correo y contraseña son requeridos');
    }

    const usuario = await User.findOne({ where: { correo } });
    if (!usuario) {
      throw new Error('Credenciales inválidas');
    }

    const esValida = await bcrypt.compare(contrasena, usuario.contrasena);
    if (!esValida) {
      throw new Error('Credenciales inválidas');
    }

    // Generar token JWT
    const payload = { id: usuario.id, correo: usuario.correo };
    const token = jwt.sign(payload, process.env.JWT_SECRET, { expiresIn: '1h' });

    return { token, usuario: { id: usuario.id, correo: usuario.correo } };
  }
}

class ServicioRefresh {
  async ejecutar(data) {
    const { refreshToken } = data;

    if (!refreshToken) {
      throw new Error('Token de actualización requerido');
    }

    try {
      const decoded = jwt.verify(refreshToken, process.env.JWT_REFRESH_SECRET);
      const nuevoAccessToken = jwt.sign(
        { id: decoded.id, correo: decoded.correo },
        process.env.JWT_SECRET,
        { expiresIn: '1h' }
      );

      return { accessToken: nuevoAccessToken };
    } catch (err) {
      throw new Error('Token inválido o expirado');
    }
  }
}

class ServicioLogout {
  async ejecutar({ token }) {
    agregarATokenBlacklist(token);
    return true;
  }
}

class ServicioForgotPassword {
  async ejecutar(data) {
    const { correo } = data;

    if (!correo) throw new Error('Correo requerido');

    const usuario = await User.findOne({ where: { correo } });
    if (!usuario) throw new Error('Correo no registrado');

    const payload = { id: usuario.id, correo: usuario.correo };
    const token = jwt.sign(payload, process.env.JWT_SECRET, { expiresIn: '15m' });

    // Simulamos envío por correo (solo lo devolvemos en la respuesta)
    return { tokenRecuperacion: token };
  }
}

class ServicioResetPassword {
  async ejecutar(data) {
    const { token, nuevaContrasena } = data;

    if (!token || !nuevaContrasena) {
      throw new Error('Token y nueva contraseña son requeridos');
    }

    let decoded;
    try {
      decoded = jwt.verify(token, process.env.JWT_SECRET);
    } catch (err) {
      throw new Error('Token inválido o expirado');
    }

    const usuario = await User.findOne({ where: { id: decoded.id } });
    if (!usuario) throw new Error('Usuario no encontrado');

    const hash = await bcrypt.hash(nuevaContrasena, 10);
    usuario.contrasena = hash;
    await usuario.save();

    return true;
  }
}

module.exports = { ServicioRegistro, ServicioLogin, ServicioRefresh, ServicioLogout, ServicioForgotPassword, ServicioResetPassword };
