const { Op, where,Sequelize } = require("sequelize");
const axios = require("axios");

// Fungsi untuk mendapatkan nilai kurs dari API
const getUSDFromAPI = async () => {
  try {
    const response = await axios.get(
      "https://api.kurs.web.id/api/v1",
      {
        params: {
          token: process.env.EXCHANGE_API_TOKEN,
          matauang: "usd",
          bank: "bca",
        },
      }
    );

    const exchangeRate = response.data.jual;
    if (!exchangeRate) throw new Error("Exchange rate not available.");

    return (exchangeRate).toFixed(2);
  } catch (error) {
    console.error("Error fetching exchange rate:", error.message);
    return 0; // Kembalikan nilai default jika API gagal
  }
};

const get = async(req, res, callback) => {
  try {
    const page = parseInt(req.query.page) || 1;
    const per_page = parseInt(req.query.per_page) || 10;
    
    let search = {}           

    let options = {
      where: search,
      page: page < 1 ? 1 : page,
      paginate: per_page,
      order: [["id", "desc"]],
    };

    const { docs, pages, total } = await model.Product.paginate(options);
    
    const kurs = await getUSDFromAPI()??0;
    await Promise.all(
      //tambah kolom priceIDR dan priceUSD
      docs?.map(async (data) => {
          data.dataValues.priceIDR = data.price; 
          data.dataValues.priceUSD = kurs>0 ? data.price/kurs : 0; 

          delete data.dataValues.price;
          delete data.dataValues.created_at;
          delete data.dataValues.updated_at;
          return await data;
      })
    ).then(async(resData)=>{
      let data = {
        data: await resData,
        current_page: page,
        total_pages: pages,
        total_items: resData?.length,
        per_page: per_page,
        next: page >= pages ? null : page + 1,
        previous: (page === 1) ? null : (page <= pages ? page - 1 : null),
      }
      callback(data, '');
    })
  } catch (error) {
    log.error(error)
    callback('', error);
  }
};
  
const getAggregation = async(req, res, callback) => {
  try {
      const aggregatedData = await model.Product.findAll({
        attributes: [
            'department',
            'product',
            'price'
        ],
        group: ['department', 'product', 'price'],
        order: [['price', 'ASC']],
    });
    aggregatedData ? callback(aggregatedData, '') : callback('', error);
  } catch (error) {
    log.error(error)
    callback('', error);
  }
};
  
module.exports = { get, getAggregation };