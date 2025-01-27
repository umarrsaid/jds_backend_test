'use strict';
const { Model } = require('sequelize');
const dataTypes = require('sequelize-mock/src/data-types');
const sequelizePaginate = require('sequelize-paginate');

module.exports = (sequelize, DataTypes) => {
  class Product extends Model {
    static associate(models) {
      //
    }
  };

  Product.init({
    id: {
      type: DataTypes.UUID,
      primaryKey: true,
    },
    product: DataTypes.STRING,
    department: DataTypes.STRING,
    price: DataTypes.DOUBLE,
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
    modelName: 'Product',
    tableName: 'products'
  });

  sequelizePaginate.paginate(Product);

  return Product;
};