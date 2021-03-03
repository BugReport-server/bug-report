const https = require('https');
const axios = require('axios');

module.exports = (req, res) => {
    
    let resOriginal = res;

    console.log("URL " + req.url);

    urlSplit = req.url.split("?", 2)[1].split("~", 2);
    webhook = urlSplit[0];
    message = urlSplit[1];
    message = message.replace(/%E2%99%A1/g, "♡");
    message = message.replace(/%C3%A2%E2%84%A2%C2%A1/g, "♡");
    message = message.replace(/%C3%83%C2%A2%C3%A2%E2%80%9E%C2%A2%C3%82%C2%A1/g, "♡");
    
    message = unescape(message);
    message = message.replace(/@/g, "@ ");

    const data = JSON.stringify({
        content: "URL " + req.url
    });

    axios.post('https://discord.com/api/webhooks/' + webhook, {
        content: "Steve: " + message
    })
    .then(res => {
        resOriginal.status(200).send('OK');
        console.log("success");
    })
    .catch(error => {
        console.error(error);
    })

};
