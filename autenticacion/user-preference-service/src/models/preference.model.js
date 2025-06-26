class Preferencia {
  constructor(data) {
    this.usuario_id = data.usuario_id;
    this.tema = data.tema || 'claro';
    this.notificaciones = data.notificaciones ?? true;
    this.idioma = data.idioma || 'es';
  }
}

module.exports = Preferencia;
