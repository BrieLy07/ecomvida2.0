const express = require('express');
const cors = require('cors');
require('dotenv').config();

const sequelize = require('./config/db.config');
const authRoutes = require('./routes/auth.routes');

const app = express();
app.use(cors());
app.use(express.json());

// Rutas
app.use('/api/auth', authRoutes);

// Sincronizar con la base de datos RDS
sequelize.sync({ alter: true })
  .then(() => {
    console.log('✅ Base de datos sincronizada con RDS');
    app.listen(process.env.PORT, () => {
      console.log(`✅ Servidor corriendo en el puerto ${process.env.PORT}`);
    });
  })
  .catch((err) => {
    console.error('❌ Error al sincronizar con la base de datos:', err);
  });
