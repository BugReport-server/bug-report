const https = require('https');

module.exports = (req, res) => {
    
    console.log("URL " + req.url);
    res.status(200).send('OK');

};
