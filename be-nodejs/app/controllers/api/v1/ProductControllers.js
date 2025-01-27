const express = require('express');
const router = express.Router();
const Product = require(appDir + '/app/repository/Product');

router.get("/", auth(), async (req, res, next) => {
  Product.get(req, res, (data, error) =>{
    if(!error){
      responser.success(res, 200, "Get Product Successfully", data)
    }else{
      responser.error(res, 400, "Get Product Error", error.message)
    }
  })
});

router.get("/aggregation", auth('admin'), async (req, res, next) => {
  Product.getAggregation(req, res, (data, error) =>{
    if(!error){
      responser.success(res, 200, "Get Product Aggregation Successfully", data)
    }else{
      responser.error(res, 400, "Get Product Aggregation Error", error.message)
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