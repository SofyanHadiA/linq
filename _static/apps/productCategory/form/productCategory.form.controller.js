/*global $app $*/

require('cropit');
var bootbox = require('bootbox');


function productCategoryFormController(endpoint, data) {
    var $modal = $app.$view.$modal;
    var $form = $app.$view.$form;
    var $http = $app.$http;
    var $notify = $app.$notify;

    var self = {
        load: onLoad,
        close: onClose,
        modal: $app.$view.$modal,
        formId: "#productCategory-form",
        data: data || {},
        isPhotoChanged: false,
        promise: {},
        defer: $.Deferred(),
        formConfig: {
            rules: {
                title: {
                    minlength: 4,
                    required: true
                },
                slug: {
                    required: true
                }
            }
        }
    }

    self.load();

    return self;

    function onLoad() {
        var modalConfig = {
            size: 'lg',
            modalId: self.modal.generateId()
        }

        var input = {
            uidInput: $form.input("uid").setValue(self.data["uid"], "AUTO"),
            titleInput: $form.input("title").setValue(self.data["title"]).setClass("required"),
            slugInput: $form.input("slug").setValue(self.data["slug"]).setClass("required"),
            descriptionInput: $form.input("description").setValue(self.data["description"], 0)
        };

        self.modal = $modal.show(require('./productCategory.form.template.hbs'), input, modalConfig);

        $form.create(self.formId)
            .config(self.formConfig)
            .onSubmit(function() {
                event.preventDefault();
                var newData = $(self.formId).serializeObject();

                if (!data) {
                    $http.post(endpoint, newData).success(function(data) {
                        onDone(data.data[0])
                    });
                }
                else {
                    $http.put(endpoint + "/" + self.data["uid"], newData).success(function(data) {
                        onDone(data.data[0])
                    });
                }
            });

        return self;
    }

    function doDelete(id) {
        $http.delete(endpoint + "/" + id).success(function(model) {
            self.modal.hide();
            onClose();
        });
    }

    function onDone(data) {
        self.modal.hide();
        self.defer.resolve();
    }

    function onClose() {
        return $.when(self.defer.promise());
    }
};

module.exports = productCategoryFormController;