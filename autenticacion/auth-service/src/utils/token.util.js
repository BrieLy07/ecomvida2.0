const jwt = require('jsonwebtoken');

const blacklist = new Set();

function agregarATokenBlacklist(token) {
  blacklist.add(token);
}

function estaEnBlacklist(token) {
  return blacklist.has(token);
}

function verificarToken(token, secreto) {
  if (estaEnBlacklist(token)) {
    throw new Error('Token inválido');
  }

  return jwt.verify(token, secreto);
}

module.exports = {
  agregarATokenBlacklist,
  estaEnBlacklist,
  verificarToken
};
