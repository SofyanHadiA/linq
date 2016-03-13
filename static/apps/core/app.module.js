/*
 *
 */

function moduleModule() {
    var self = {
        modules: {},
        register: register,
        resolve: resolve
    };

    return self;

    function register(key, value) {
        this.modules[key] = value;
    };
    function resolve(key) {
        return this.modules[key];
    };

};

module.exports = moduleModule();