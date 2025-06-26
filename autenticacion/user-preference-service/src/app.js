const express = require('express');
const dotenv = require('dotenv');
const preferencesRoutes = require('./routes/preferences.routes');

dotenv.config();

const app = express();
app.use(express.json());

app.use('/api', preferencesRoutes);

const PORT = process.env.PORT || 3004;
app.listen(PORT, () => {
  console.log(`✅ user-preferences-service corriendo en el puerto ${PORT}`);
});
