var $ = jQuery;

function httpModule() {

    var self = {
        getToken: undefined,
        post: post,
        get: get
    };

    //TODO: Get token first before doing any request
    self.token = {}; // app.http.get('../token');
    self.cachedScriptPromises = {};
    self.deferFactory = function(requestFunction) {
        var cache = {};
        return function(key, callback) {
            if (!cache[key]) {
                cache[key] = $.Deferred(function(defer) {
                    requestFunction(defer, key);
                }).promise();
            }
            return cache[key].done(callback);
        };
    };

    function get(url) {
        self.deferFactory(function(defer, url) {
            $.get(url, self.http.token).then(
                defer.resolve,
                defer.reject)
        });
    };

    function post(url, data, callback) {
        var postData = {
            data: data,
            token: ""
        };
        
        $.post(url, postData, function(response) {
            callback(response), "json"
        }).fail(function(error) {
            $app.$notify.danger("Post Failed to: <b>" + url + "</b> " + error.responseText);
        })
    };

    return self;
};

module.exports = httpModule();