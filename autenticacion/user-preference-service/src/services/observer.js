class Observer {
  actualizar(data) {
    console.log(`[Observer] Preferencias actualizadas:`, data);
    // Aquí podrías notificar por Webhook, Kafka, etc.
  }
}

class Sujeto {
  constructor() {
    this.observadores = [];
  }

  agregar(observador) {
    this.observadores.push(observador);
  }

  notificar(data) {
    this.observadores.forEach(obs => obs.actualizar(data));
  }
}

module.exports = {
  Observer,
  Sujeto
};
