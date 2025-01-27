const { check, validationResult } = require('express-validator');
const  passwordValidator = require('password-validator');

const reqBodyClient = [
    check('nik').isLength({ min: 16, max:16 }).withMessage('NIK harus memiliki tepat 16 karakter'),
    check('password').isLength({ min: 1 }).withMessage('Password tidak boleh kosong')
];

const reqBodyRegister = [
    check('nik').isLength({ min:16, max: 16 }).withMessage('NIK harus memiliki tepat 16 karakter')
    .custom(async (nik,{req}) => { 
        const existingUser =  await model.User.findOne({
        where: { nik: nik },
        }) 
        if (existingUser) { 
            throw new Error('NIK tersebut sudah terdaftar') 
        } 
    }),
    check('role').isLength({ min: 1 }).withMessage('Role tidak boleh kosong')
                .isIn(['superadmin','admin','user'])
                .withMessage("Tidak sesuai pilihan, harap masukan : superadmin, admin atau user"),
];


module.exports = {
  reqBodyClient,
  reqBodyRegister,
}