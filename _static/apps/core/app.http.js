var $ = jQuery;

// TODO: Update Token

function httpModule() {

    var self = {
        getToken: undefined,
        post: post,
        get: get
    };
    
    //TODO: Get token first before doing any request
    self.token = {};// app.http.get('../token');
    self.cachedScriptPromises = {};

    return self;

    var deferFactory = function (requestFunction) {
        var cache = {};
        return function (key, callback) {
            if (!cache[key]) {
                cache[key] = $.Deferred(function (defer) {
                    requestFunction(defer, key);
                }).promise();
            }
            return cache[key].done(callback);
        };
    };

    function get(url) {
        deferFactory(function (defer, url) {
            $.get(url, self.http.token).then(
                defer.resolve,
                defer.reject)
        });
    };

    function post(url, data) {
        deferFactory(function (defer, url) {
            // TODO: MERGE DATA WITH TOKEN
            $.post(url, data, function (response) {
                defer.resolve(response)
            }).then(defer.resolve,
                defer.reject);
        });
    };
};

module.exports = httpModule();