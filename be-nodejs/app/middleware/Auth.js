var jwt = require('jsonwebtoken');
var fs = require('fs');
const { Console } = require('console');

module.exports = (allowedRoles) => {
  return function (req, res, next) {
    const authHeader = req.headers.authorization;
    const cert_public = process.env.JWT_KEY_SECRET;;

    if (!authHeader) {
      return responser.error(res, 401, "Unauthorized", null)
    }
    const token = authHeader.split(' ')[1];

    jwt.verify(token, cert_public, { algorithms: ['HS256'] }, function (err, payload) {
        if (err) {
          return responser.error(res, 401, "Unauthorized", null)
        }

        if (allowedRoles && !allowedRoles.includes(payload.data.role)) {
          return responser.error(res, 401, "You do not have permission to access this resource", null)
        }

        req.user = payload.data;            
        return next(); 
    });
  }
}