const express = require('express');
const router = express.Router();
const jwt = require('jsonwebtoken');
const fs = require('fs');
const { check, validationResult } = require('express-validator');
const Validator = require(appDir + '/app/validation/AuthValidations');
const AuthRepo = require(appDir + '/app/repository/Auth');
const { Op } = require("sequelize");

// function
var verifyuser = async(nik, password) => {
    const user = await model.User.findOne({ 
                    raw: true,
                    where: {nik: nik}});
    if(user === null){
        return { message: 'User tidak terdaftar', status: false}
    }else{
        if (user.password == Crypt.hash_password(password)){
            return { message: 'Berhasil Login', status: true, id: user.id, nik: user.nik, role: user.role }
        }else{
            return { message: 'Password salah', status: false }
        }
    }
}   

router.post('/login', Validator.reqBodyClient, async(req, res, next) => {
    try{
        const errors = validationResult(req);
        if (!errors.isEmpty()) {
            return responser.error(res, 200, "Login Error", (errors['errors'][0].param+" "+errors['errors'][0].msg))
        }
        
        var {nik, password} = req.body
        var user = await verifyuser(nik, password)        
        if(user.status == true ){
            data = {
                id: user.id,
                nik: user.nik,
                role: user.role
            }
            var cert = process.env.JWT_KEY_SECRET;;
            var token = jwt.sign({ exp: Math.floor(Date.now() / 1000) + (1440 * 60), data: data }, cert, { algorithm: 'HS256' });
            
            responser.success(res, 200, "Login Succesfully", {...data,token:token})
        }else{
            responser.error(res, 400, "Login Error", user.message)
        }
    } catch (error) {
        responser.error(res, 400, "Login Error", error.message)
    }
});

router.post('/register', Validator.reqBodyRegister, async(req, res, next) => {
    console.log('req ',req.body)
    const errors = validationResult(req);
    if (!errors.isEmpty()) {
        console.log(errors)
        return responser.error(res, 400, 'Register Error', errors['errors']);
    }

    AuthRepo.createUser(req, res, (data, error) =>{
        if(!error){
            responser.success(res, 200, "Register Succesfully", data)
        }else{
            responser.error(res, 400, "Register Error", error)
        }
    })
});

router.get('/private-claim', auth(), (req, res) => {
  // req.user berisi data user dari payload JWT, yang di extract di middleware
  const userData = req.user;

  const privateClaims = {
    id: userData.id,
    nik: userData.nik,
    role: userData.role
  };

  responser.success(res, 200, "Get Private claim data Successfully", privateClaims)
});


module.exports = router;