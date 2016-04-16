var $ = jQuery;
require('../../../node_modules/jquery-validation/dist/jquery.validate.js');

var formModule = function() {

    var self = {
        create: create,
        config: config,
        validation: {},
        onSubmit: onSubmit,
        input: input
    };

    return self;

    function create(formId) {
        self.container = formId
        self.validation = $(self.container).validate({
            errorClass: "error text-red",
            errorPlacement: function(error, element) {
                error.insertBefore(element);
            },
            highlight: function(element) {
                $(element).closest('.control-group').removeClass('success').addClass('error');
            },
            success: function(element) {
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

    function input(name, inputType = "text", className = "", value = "") {
        
        var self = {
            name: name,
            inputType: inputType,
            value: value,
            className: className,
            setValue: setValue,
            setClass: setClass,
            render: render,
        }

        return self;

        function setValue(val) {
            self.value = val || "";
            return self;
        }

        function setClass(className) {
            self.className = className;
            return self;
        }

        function render() {
            return '<div class="form-group">' +
                '<label for="' + self.name + '" class="col-sm-4 control-label ' + self.className + '">' + $app.$language[self.name] + '</label>' +
                '<div class="col-sm-8">' +
                '<input type="' + self.inputType + '" name="' + self.name + '" id="' + self.name + '" class="form-control" value="' + self.value + '" />' +
                '</div></div>'
        }
    }
};


module.exports = formModule();
