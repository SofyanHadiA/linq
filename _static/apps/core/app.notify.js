
var $ = jQuery;
require('bootstrap-notify');

function notifyModule() {

    var self = {
        info: info,
        warning: warning,
        danger: danger
    }

    return self;

    function info(message) {
        return $.notify(
            { icon: 'fa fa-info-circle', message: message },
            { type: 'info', z_index: 1100 });
    };

    function warning(message) {
        return $.notify(
            { icon: 'fa fa-exclamation', message: message },
            { type: 'warning', z_index: 1100 });
    };

    function danger(message) {
        return $.notify(
            { icon: 'fa fa-warning', message: message },
            { type: 'danger', z_index: 1100 });
    };
}

module.exports = notifyModule(); 