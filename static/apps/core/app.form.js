/*
 * Application Form Core Module 
 */

var $ = jQuery;
require('../../node_modules/jquery-validation/dist/jquery.validate.js');

var formModule = function () {

    var self = {
        create: create,
        config: config,
        validation: {},
        onSubmit: onSubmit
    };

    return self;

    function create(formId) {
        self.container = formId
        self.validatin = $(self.container).validate({
            errorClass: "error text-red",
            errorPlacement: function (error, element) {
                error.insertBefore(element);
            },
            highlight: function (element) {
                $(element).closest('.control-group').removeClass('success').addClass('error');
            },
            success: function (element) {
                element.addClass('valid').closest('.control-group').removeClass('error').addClass('success');
            },
        })

        return self;
    };

    function config(config) {
        $.extend(self.validation.settings, config);
        return self;
    };

    function onSubmit(submitFunc) {
        self.validation.settings.submitHandler = submitFunc;
        return self;
    };
};


module.exports = formModule();
