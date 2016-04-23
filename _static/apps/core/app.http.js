function httpModule() {
    var self = {
        cache: {},
        getToken: undefined,
        get: get,
        post: post,
        put: put,
        delete: remove
    };
    
    return self;
    
    function get(url) {
        return $.when(self.cache[url] || 
            $.ajax({
                url: url,
                data: {token : ""},
                type: 'get',
                success: function(data, textStatus, jqXHR) {
                    self.cache[url] = data;
                },
                fail: function(jqXHR, textStatus, errorThrown) {
                    $app.$notify.danger("GET Failed: <b>" + url + "</b> " + jqXHR.responseText);
                }
            })
        );
    };

    function post(url, data) {
        var postData = {
            data: data,
            token: ""
        };
        return $.when(
            $.ajax({
                url: url,
                data: JSON.stringify(postData),
                type: 'post',
                fail: function(jqXHR, textStatus, errorThrown) {
                    $app.$notify.danger("POST Failed: <b>" + url + "</b> " + jqXHR.responseText);
                }
            })
        );
    };
    
    function put(url, data) {
        var postData = {
            data: data,
            token: ""
        };
        return $.when(
            $.ajax({
                url: url,
                data: JSON.stringify(postData),
                type: 'put',
                success: function(data, textStatus, jqXHR) {
                    delete self.cache[url]
                },
                fail: function(jqXHR, textStatus, errorThrown) {
                    $app.$notify.danger("PUT Failed: <b>" + url + "</b> " + jqXHR.responseText);
                }
            })
        );
    };
    
    function remove(url, data) {
        var postData = {
            data: data,
            token: ""
        };
        return $.when(
            $.ajax({
                url: url,
                data: JSON.stringify(postData),
                type: 'delete',
                success: function(data, textStatus, jqXHR) {
                    delete self.cache[url]
                },
                fail: function(jqXHR, textStatus, errorThrown) {
                    $app.$notify.danger("DELETE Failed: <b>" + url + "</b> " + jqXHR.responseText);
                }
            })
        );
    };
};

module.exports = httpModule;