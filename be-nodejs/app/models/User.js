'use strict';
const { Model } = require('sequelize');
const sequelizePaginate = require('sequelize-paginate');

module.exports = (sequelize, DataTypes) => {
  class User extends Model {
    static associate(models) {
      // definisikan associate table disini
    }
  }
  User.init({
    nik: {
      type: DataTypes.CHAR,
      unique:  true
    },
    password: {
      type: DataTypes.STRING
    },
    role: DataTypes.ENUM('superadmin', 'admin','user'),
    createdAt: {
      type: DataTypes.DATE,
      field: 'created_at'
    },
    updatedAt: {
      type: DataTypes.DATE,
      field: 'updated_at'
    }
  }, {
    sequelize,
    modelName: 'User',
    tableName: 'users',
    hooks: {
      afterDestroy: function (instance, options) {
        //hooks (cascade)
      }
    }
  });

  sequelizePaginate.paginate(User);

  return User;
};