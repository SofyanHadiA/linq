var $ = jQuery;
var $selectize = require('../../../node_modules/selectize/dist/js/selectize.js');
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
            setSelected: setSelected,
            formGroup: formGroup,
            formGroupPassword: formGroupPassword,
            formGroupNumber: formGroupNumber,
            formGroupTextArea: formGroupTextArea,
            formGroupDropDown: formGroupDropDown,
        }

        return self;

        function setValue(value, defaultValue) {
            self.value = value || defaultValue || "";
            return self;
        }

        function setSelected(value, title) {
            if (value) {
                self.selectedItem = {
                    value: value,
                    title: title || value
                };
            }
            return self;
        }

        function setClass(className) {
            self.className = className;
            return self;
        }

        function formGroup(inputWidth = 8, labelWidth = 4) {
            return '<div class="form-group">' + formLabel(labelWidth) + inputText("text", inputWidth) + '</div>'
        }

        function formGroupNumber(inputWidth = 8, labelWidth = 4) {
            return '<div class="form-group">' + formLabel(labelWidth) + inputText("number", inputWidth) + '</div>'
        }

        function formGroupPassword(inputWidth = 8, labelWidth = 4) {
            return '<div class="form-group">' + formLabel(labelWidth) + inputText("password", inputWidth) + '</div>'
        }

        function formGroupTextArea(inputWidth = 8, labelWidth = 4) {
            return '<div class="form-group">' + formLabel(labelWidth) + inputTextArea() + '</div>'
        }

        function formGroupDropDown(inputWidth = 8, labelWidth = 4) {
            return '<div class="form-group">' + formLabel(labelWidth) + inputDropDown() + '</div>'
        }

        function inputText(type = "text", inputWidth = '8') {
            return '<div class="col-xs-' + inputWidth + '">' +
                '<input type="' + type + '" name="' + self.name + '" id="' + self.name + '" class="form-control" value="' + self.value + '" />' +
                '</div>'
        }

        function inputDropDown(inputWidth = '8') {
            var select = '<div class="col-xs-' + inputWidth + '">' +
                '<select name="' + self.name + '" id="' + self.name + '" class="form-control" value="' + self.value + '" >';
            if (self.selectedItem) {
                select += '<option value="' + self.selectedItem.value + '">' + self.selectedItem.title + '</option>'
            }
            select += '</select></div>'

            return select;
        }

        function inputTextArea(inputWidth = '8') {
            return '<div class="col-xs-' + inputWidth + '">' +
                '<textarea name="' + self.name + '" id="' + self.name + '" class="form-control">' + self.value + '</textarea>' +
                '</div>'
        }

        function formLabel(labelWidth = '4') {
            return '<label for="' + self.name + '" class="col-xs-' + labelWidth + ' control-label ' + self.className + '">' + self.label + '</label>'
        }
    }
};


module.exports = formModule();
