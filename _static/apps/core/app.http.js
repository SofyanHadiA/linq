function httpModule() {
    var self = {
        cache: {},
        getToken: undefined,
        get: get,
        post: post,
        put: put,
        delete: remove,
        upload: upload
    };

    return self;

    function get(url) {
        return $.when(self.cache[url] ||
            $.ajax({
                url: url,
                data: {
                    token: ""
                },
                type: 'get',
                success: function(data, textStatus, jqXHR) {
                    self.cache[url] = data;
                },
                error: function(jqXHR, textStatus, errorThrown) {
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
        return (
            $.ajax({
                url: url,
                data: JSON.stringify(postData),
                type: 'post',
                error: function(jqXHR, textStatus, errorThrown) {
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
        return (
            $.ajax({
                url: url,
                data: JSON.stringify(postData),
                type: 'put',
                success: function(data, textStatus, jqXHR) {
                    delete self.cache[url]
                },
                error: function(jqXHR, textStatus, errorThrown) {
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
        return (
            $.ajax({
                url: url,
                data: JSON.stringify(postData),
                type: 'delete',
                success: function(data, textStatus, jqXHR) {
                    delete self.cache[url]
                },
                error: function(jqXHR, textStatus, errorThrown) {
                    $app.$notify.danger("DELETE Failed: <b>" + url + "</b> " + jqXHR.responseText);
                }
            })
        );
    };

    function upload(url, data) {
        return (
            $.ajax({
                url: url,
                data: data,
                type: 'post',
                processData: false,
                contentType: false,
                error: function(jqXHR, textStatus, errorThrown) {
                    $app.$notify.danger("Upload Failed: <b>" + url + "</b> " + jqXHR.responseText);
                }
            })
        );
    };
};

module.exports = httpModule;