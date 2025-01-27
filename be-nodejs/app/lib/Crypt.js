//encrypt
const Cryptr = require('cryptr');
const cryptr = new Cryptr(process.env.SECRET);
const crypto = require('crypto');
const moment = require('moment');

const encrypt = (obj) => {
  const encryptedString = cryptr.encrypt(obj);
  return encryptedString
}

const decrypt = (obj) => {
  const decryptedString = cryptr.decrypt(obj);
  return decryptedString
}

const hash_password = (password) => {
  const salt = process.env.SECRET;
  const hash = crypto.pbkdf2Sync(password, salt, 1000, 64, `sha512`);
  return hash.toString(`hex`);
}

const UUID = (obj) => {
  return obj+"-"+moment().format('YYMMDDHH-mmssSSSS')
}

module.exports = {
  encrypt, decrypt, UUID, hash_password
};