'use strict';
/** @type {import('sequelize-cli').Migration} */
const Crypt = require('../app/lib/Crypt');
const uuid = require('uuid')

module.exports = {
  async up (queryInterface, Sequelize) {
    await queryInterface.bulkInsert('products', [
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "218.00",
      department: "Outdoors",
      product: "Salad"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "784.00",
      department: "Games",
      product: "Chips"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "652.00",
      department: "Automotive",
      product: "Bike"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "544.00",
      department: "Tools",
      product: "Tuna"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "643.00",
      department: "Toys",
      product: "Fish"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "100.00",
      department: "Baby",
      product: "Bacon"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "584.00",
      department: "Clothing",
      product: "Pants"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "265.00",
      department: "Kids",
      product: "Bacon"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "493.00",
      department: "Health",
      product: "Computer"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "877.00",
      department: "Games",
      product: "Soap"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "293.00",
      department: "Beauty",
      product: "Table"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "546.00",
      department: "Books",
      product: "Chair"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "906.00",
      department: "Automotive",
      product: "Salad"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "870.00",
      department: "Computers",
      product: "Car"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "443.00",
      department: "Movies",
      product: "Pants"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "810.00",
      department: "Toys",
      product: "Towels"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "33.00",
      department: "Health",
      product: "Car"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "803.00",
      department: "Outdoors",
      product: "Table"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "219.00",
      department: "Computers",
      product: "Fish"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "878.00",
      department: "Garden",
      product: "Cheese"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "556.00",
      department: "Games",
      product: "Keyboard"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "38.00",
      department: "Clothing",
      product: "Hat"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "381.00",
      department: "Music",
      product: "Keyboard"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "881.00",
      department: "Outdoors",
      product: "Shoes"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "327.00",
      department: "Movies",
      product: "Chicken"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "538.00",
      department: "Sports",
      product: "Salad"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "621.00",
      department: "Shoes",
      product: "Bike"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "849.00",
      department: "Beauty",
      product: "Chips"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "869.00",
      department: "Games",
      product: "Ball"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "840.00",
      department: "Automotive",
      product: "Towels"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "353.00",
      department: "Garden",
      product: "Hat"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "645.00",
      department: "Sports",
      product: "Sausages"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "304.00",
      department: "Kids",
      product: "Shoes"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "665.00",
      department: "Home",
      product: "Computer"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "179.00",
      department: "Jewelery",
      product: "Table"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "752.00",
      department: "Books",
      product: "Pizza"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "653.00",
      department: "Kids",
      product: "Chicken"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "307.00",
      department: "Games",
      product: "Sausages"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "740.00",
      department: "Garden",
      product: "Bike"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "111.00",
      department: "Games",
      product: "Bike"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "894.00",
      department: "Industrial",
      product: "Cheese"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "421.00",
      department: "Clothing",
      product: "Ball"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "819.00",
      department: "Tools",
      product: "Soap"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "929.00",
      department: "Home",
      product: "Bike"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "231.00",
      department: "Shoes",
      product: "Mouse"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "991.00",
      department: "Tools",
      product: "Gloves"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "211.00",
      department: "Music",
      product: "Car"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "408.00",
      department: "Shoes",
      product: "Cheese"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "137.00",
      department: "Grocery",
      product: "Table"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "505.00",
      department: "Shoes",
      product: "Salad"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "921.00",
      department: "Automotive",
      product: "Mouse"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "146.00",
      department: "Beauty",
      product: "Gloves"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "559.00",
      department: "Automotive",
      product: "Hat"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "96.00",
      department: "Beauty",
      product: "Tuna"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "82.00",
      department: "Health",
      product: "Bacon"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "285.00",
      department: "Garden",
      product: "Chips"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "977.00",
      department: "Home",
      product: "Keyboard"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "973.00",
      department: "Kids",
      product: "Car"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "950.00",
      department: "Kids",
      product: "Sausages"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "180.00",
      department: "Industrial",
      product: "Table"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "276.00",
      department: "Baby",
      product: "Gloves"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "72.00",
      department: "Books",
      product: "Bike"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "140.00",
      department: "Movies",
      product: "Tuna"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "916.00",
      department: "Beauty",
      product: "Salad"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "416.00",
      department: "Garden",
      product: "Hat"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "875.00",
      department: "Clothing",
      product: "Bike"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "51.00",
      department: "Health",
      product: "Pizza"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "779.00",
      department: "Books",
      product: "Cheese"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "773.00",
      department: "Automotive",
      product: "Chair"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "4.00",
      department: "Games",
      product: "Computer"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "355.00",
      department: "Movies",
      product: "Chicken"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "918.00",
      department: "Baby",
      product: "Car"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "13.00",
      department: "Automotive",
      product: "Car"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "5.00",
      department: "Computers",
      product: "Fish"
   },
   {
      id: uuid.v4(),
      created_at: new Date(),
      updated_at: new Date(),
      price: "576.00",
      department: "Health",
      product: "Ball"
   },
   {
      created_at: new Date(),
      updated_at: new Date(),
      price: "102.00",
      department: "Garden",
      product: "Salad",
      id: uuid.v4()
   },
   {
      created_at: new Date(),
      updated_at: new Date(),
      price: "169.00",
      department: "Movies",
      product: "Towels",
      id: uuid.v4()
   },
   {
      created_at: new Date(),
      updated_at: new Date(),
      price: "907.00",
      department: "Outdoors",
      product: "Car",
      id: uuid.v4()
   },
   {
      created_at: new Date(),
      updated_at: new Date(),
      price: "653.00",
      department: "Automotive",
      product: "Tuna",
      id: uuid.v4()
   },
   {
      created_at: new Date(),
      updated_at: new Date(),
      price: "485.00",
      department: "Toys",
      product: "Pants",
      id: uuid.v4()
   },
   {
      created_at: new Date(),
      updated_at: new Date(),
      price: "596.00",
      department: "Tools",
      product: "Chicken",
      id: uuid.v4()
   },
   {
      created_at: new Date(),
      updated_at: new Date(),
      price: "587.00",
      department: "Toys",
      product: "Pizza",
      id: uuid.v4()
   },
   {
      created_at: new Date(),
      updated_at: new Date(),
      price: "470.00",
      department: "Clothing",
      product: "Mouse",
      id: uuid.v4()
   },
   {
      created_at: new Date(),
      updated_at: new Date(),
      price: "764.00",
      department: "Grocery",
      product: "Car",
      id: uuid.v4()
   },
   {
      created_at: new Date(),
      updated_at: new Date(),
      price: "810.00",
      department: "Clothing",
      product: "Sausages",
      id: uuid.v4()
   },
   {
      created_at: new Date(),
      updated_at: new Date(),
      price: "782.00",
      department: "Grocery",
      product: "Towels",
      id: uuid.v4()
   },
   {
      created_at: new Date(),
      updated_at: new Date(),
      price: "395.00",
      department: "Home",
      product: "Pants",
      id: uuid.v4()
   },
   {
      created_at: new Date(),
      updated_at: new Date(),
      price: "758.00",
      department: "Sports",
      product: "Keyboard",
      id: uuid.v4()
   },
   {
      created_at: new Date(),
      updated_at: new Date(),
      price: "837.00",
      department: "Games",
      product: "Pizza",
      id: uuid.v4()
   },
   {
      created_at: new Date(),
      updated_at: new Date(),
      price: "267.00",
      department: "Jewelery",
      product: "Soap",
      id: uuid.v4()
   },
   {
      created_at: new Date(),
      updated_at: new Date(),
      price: "908.00",
      department: "Kids",
      product: "Keyboard",
      id: uuid.v4()
   },
   {
      created_at: new Date(),
      updated_at: new Date(),
      price: "924.00",
      department: "Computers",
      product: "Mouse",
      id: uuid.v4()
   },
   {
      created_at: new Date(),
      updated_at: new Date(),
      price: "644.00",
      department: "Industrial",
      product: "Chips",
      id: uuid.v4()
   },
   {
      created_at: new Date(),
      updated_at: new Date(),
      price: "866.00",
      department: "Tools",
      product: "Fish",
      id: uuid.v4()
   },
   {
      created_at: new Date(),
      updated_at: new Date(),
      price: "775.00",
      department: "Jewelery",
      product: "Chips",
      id: uuid.v4()
   },
   {
      created_at: new Date(),
      updated_at: new Date(),
      price: "584.00",
      department: "Outdoors",
      product: "Keyboard",
      id: uuid.v4()
   },
   {
      created_at: new Date(),
      updated_at: new Date(),
      price: "940.00",
      department: "Games",
      product: "Soap",
      id: uuid.v4()
   },
   {
      created_at: new Date(),
      updated_at: new Date(),
      price: "908.00",
      department: "Home",
      product: "Gloves",
      id: uuid.v4()
   },
   {
      created_at: new Date(),
      updated_at: new Date(),
      price: "226.00",
      department: "Kids",
      product: "Tuna",
      id: uuid.v4()
   },
   {
      created_at: new Date(),
      updated_at: new Date(),
      price: "979.00",
      department: "Clothing",
      product: "Bike",
      id: uuid.v4()
   }
   ], {});
  },

  async down (queryInterface, Sequelize) {
    await queryInterface.bulkDelete('products', null, {});
  }
};
