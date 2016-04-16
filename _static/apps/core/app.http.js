function httpModule() {
    var cache = {};

    var self = {
        cache: cache,
        getToken: undefined,
        post: post,
        get: get
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
                    $app.$notify.danger("Post Failed to: <b>" + url + "</b> " + jqXHR.responseText);
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
                    self.cache.remove(url)
                },
                fail: function(jqXHR, textStatus, errorThrown) {
                    $app.$notify.danger("Post Failed to: <b>" + url + "</b> " + jqXHR.responseText);
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
                    self.cache.remove(url)
                },
                fail: function(jqXHR, textStatus, errorThrown) {
                    $app.$notify.danger("Post Failed to: <b>" + url + "</b> " + jqXHR.responseText);
                }
            })
        );
    };
};

module.exports = httpModule();