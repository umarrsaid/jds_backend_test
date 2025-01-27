'use strict';
/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up(queryInterface, Sequelize) {
    await queryInterface.createTable('products', {
      id: {
        allowNull: false,
        primaryKey: true,
        type: Sequelize.CHAR(36)
      },
      product: {
        type: Sequelize.STRING(100),
        defaultValue: null        
      },
      department: {
        type: Sequelize.STRING(100),
        defaultValue: null        
      },
      price: {
        type: Sequelize.DOUBLE,
        defaultValue: 0        
      },
      created_at: {
        allowNull: false,
        type: Sequelize.DATE,
        defaultValue: Sequelize.fn('NOW')
      },
      updated_at: {
        allowNull: false,
        type: Sequelize.DATE,
        defaultValue: Sequelize.fn('NOW')
      }
    });
  },
  async down(queryInterface, Sequelize) {
    await queryInterface.dropTable('products');
  }
};