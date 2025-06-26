const dynamo = require('../config/dynamodb.config');
const Preferencia = require('../models/preference.model');
const { Sujeto, Observer } = require('./observer');

const tabla = process.env.DYNAMO_TABLE;
const sujeto = new Sujeto();
sujeto.agregar(new Observer());

class PreferencesService {
  async obtener(usuarioId) {
    const params = {
      TableName: tabla,
      Key: { usuario_id: usuarioId }
    };

    const resultado = await dynamo.get(params).promise();
    return resultado.Item;
  }

  async actualizar(usuarioId, data) {
    const actualizada = new Preferencia({ usuario_id: usuarioId, ...data });

    const params = {
      TableName: tabla,
      Item: actualizada
    };

    await dynamo.put(params).promise();
    sujeto.notificar(actualizada);
    return actualizada;
  }
}

module.exports = new PreferencesService();
