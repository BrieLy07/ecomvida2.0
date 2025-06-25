const { ServicioRegistro, ServicioLogin, ServicioRefresh, ServicioLogout, ServicioForgotPassword, ServicioResetPassword } = require('./auth.service');

class FabricaDeServicios {
  crearServicio(tipo) {
    switch (tipo) {
      case 'registro': return new ServicioRegistro();
      case 'login': return new ServicioLogin();
      case 'refresh': return new ServicioRefresh();
      case 'logout': return new ServicioLogout();
      case 'forgot': return new ServicioForgotPassword();
      case 'reset': return new ServicioResetPassword();
      default: throw new Error('Tipo de servicio no reconocido');
    }
  }
}

module.exports = new FabricaDeServicios();
