'use strict';

const fs = require('fs');
const path = require('path');
const Sequelize = require('sequelize');
const basename = path.basename(__filename);
const env = process.env.NODE_ENV || 'development';
const config = require(appDir+ '/config/config.js')[env];
const db = {};

let sequelize;
if (config.use_env_variable) {
  sequelize = new Sequelize(process.env[config.use_env_variable], config);
} else {
  sequelize = new Sequelize(config.database, config.username, config.password, {
      host: config.host,
      dialect: config.dialect,
      port: config.port
  });
}

sequelize.authenticate().then(() => {
  console.log('Connection has been established successfully.');
}).catch((error) => {
  console.error('Unable to connect to the database: ', error);
});

const files = []
const sortDir = (maniDir) => {
  const folders = []
  const CheckFile = (filePath) => fs.statSync(filePath).isFile()
  const sortPath = (dir) => {
    fs.readdirSync(dir)
      .filter((file) => file.indexOf(".") !== 0 && file !== "index.js")
      .forEach((res) => {
        const filePath = path.join(dir, res)
        if (CheckFile(filePath)) {
          files.push(filePath)
        } else {
          folders.push(filePath)
        }
      })
  }
  folders.push(maniDir)
  let i = 0
  do {
    sortPath(folders[i])
    i += 1
  } while (i < folders.length)
}
sortDir(__dirname)

files.forEach((file) => {
  const model = require(file)(sequelize, Sequelize.DataTypes)
  db[model.name] = model
})

Object.keys(db).forEach((modelName) => {
  if (db[modelName].associate) {
    db[modelName].associate(db)
  }
})

db.sequelize = sequelize;
db.sequelize = sequelize;

module.exports = db;