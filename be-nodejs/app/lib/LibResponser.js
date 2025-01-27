const success = async (res, code=200, message='success', result={}) => {
    try {
        res.status(200).json({
            code: code, 
            status: 'success', 
            message: message,
            result: result
        })

    } catch (error) {
        log.error(error);
    }  
}

const error = async (res, code=400,  message='error', result={}) => {
    try { 
        res.status(code===401?401:200).json({
            code: code,
            status: 'error',
            message: message,
            result: result
        })

    } catch (error) {
        log.error(error);
    }  
}

module.exports = {
    success, error
};