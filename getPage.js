var casper = require('casper').create({
    viewportSize: {
        width: 1280,
        height: 720
    },
    onStepTimeout: function () {
        casper.on('resource.received', function (request) {
            casper.echo('852438d026c018c4307b916406f98c62');

        });
        casper.exit();
    }
});
var target = '{{.Url}}';
var end = false;
var count = 0;
var result = null;
casper.start(target);
casper.page.customHeaders = {
    'Accept-Language': 'en',
};
casper.userAgent('Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.170 Safari/537.36');
var getContent = function () {
    casper.wait(3000, function () {});
    casper.then(function () {
        casper.capture('screenshot.png');
    });
};
getContent();