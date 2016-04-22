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
            label: $app.$language[name],
            className: className,
            setValue: setValue,
            setClass: setClass,
            formGroup: formGroup,
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
        
        function formGroup(inputWidth = 8, labelWidth = 4){
            return '<div class="form-group">' + formLabel(labelWidth) + inputText(inputWidth) + '</div>'
        }

        function inputText(inputWidth = '8') {
            return '<div class="col-md-'+inputWidth+'">' +
                '<input type="text" name="' + self.name + '" id="' + self.name + '" class="form-control" value="' + self.value + '" />' +
                '</div>'
        }
        
        function formLabel(labelWidth = '4'){
            return '<label for="' + self.name + '" class="col-md-'+labelWidth+' control-label ' + self.className + '">' + self.label + '</label>' 
        }
    }
};


module.exports = formModule();
