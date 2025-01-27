var jwt = require('jsonwebtoken');
var fs = require('fs');
var generator = require('generate-password');
const crypto = require("crypto");

// function untuk mengenerate password 6 karakter
const generatePassword = async() => {
    return generator.generate({
        length: 6,
        numbers: true
    });
}

const createUser = async(req, res, callback) => {
    try {
        await model.sequelize.transaction(async (t) => {
            const { nik, role } = req.body;
            const password = await generatePassword();
            const user = await model.User.create({
                id: UUID(),
                nik,
                role,
                password: Crypt.hash_password(password),
            }, {
                transaction: t,
            })
    
            return {
                'nik': nik,
                'role': role,
                'password': password
            }
        }).then(async(user)=>{
            callback(user, '')
        })
      } catch (error) {
        callback('', error.message)
      }
}

module.exports = {
    generatePassword,
    createUser,
};