var controllers_router = async function(){
  const route = '/api/v1';
  const route_path = appDir + '/app/controllers/api/v1';
  
  app.use(route+'/', require(route_path+'/AuthControllers'));
  app.use(route+'/product', require(route_path+'/ProductControllers'));
}

module.exports = controllers_router;