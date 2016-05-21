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
                productCategory_title: {
                    minlength: 5,
                    required: true
                },
                productCategory_sell_price: {
                    required: true
                },
                productCategory_category: {
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
            uidInput: $form.input("uid").setValue(self.data["uid"] || "AUTO"),
            skuInput: $form.input("sku").setValue(self.data["sku"]),
            titleInput: $form.input("title").setValue(self.data["title"]).setClass("required"),
            buyPriceInput: $form.input("buyPrice").setValue(self.data["buyPrice"], 0),
            sellPriceInput: $form.input("sellPrice").setValue(self.data["sellPrice"], 0).setClass("required"),
            stockInput: $form.input("stock").setValue(self.data["stock"], 0).setClass("required"),
            categoryInput: $form.input("category").setValue(self.data["category"]).setClass("required")
        };

        self.modal = $modal.show(require('./productCategory.form.template.hbs'), input, modalConfig);

        $form.create(self.formId)
            .config(self.formConfig)
            .onSubmit(function() {
                event.preventDefault();
                var newData = $(self.formId).serializeObject();
                newData.buyPrice = parseFloat(newData.buyPrice) || 0;
                newData.sellPrice = parseFloat(newData.sellPrice) || 0;
                newData.stock = parseFloat(newData.stock) || 0;

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

        $('#removeUser').click(function() {
            var id = $(this).data("id");
            bootbox.confirm('Are you sure to delete this productCategory?', function(result) {
                if (result) {
                    doDelete(id);
                }
            });
        });

        $('#productCategory-photo').cropit({
            onFileChange: function() {
                self.isPhotoChanged = true;
            }
        });

        if (self.data.image) {
            $('#productCategory-photo').cropit('imageSrc', './uploads/productCategory_photos/' + self.data.image);
        };

        $('#select-image-btn').click(function() {
            $("#productCategory-form.cropit-image-input").prop('disabled', false);
            $('.cropit-image-input').click();
        });

        return self;
    }

    function uploadPhoto(productCategoryId) {
        if (self.isPhotoChanged) {
            var imageData = $('#productCategory-photo').cropit('export');
            return $http.put(endpoint + "/" + productCategoryId + "/photo", imageData);
        }
        else {
            return null
        }
    }
    
    function doDelete(id) {
        $http.delete(endpoint + "/" + id).success(function(model) {
            self.modal.hide();
            onClose();
        });
    }

    function onDone(data) {
        $.when(uploadPhoto(data.uid)).then(function() {
            self.modal.hide();
            self.defer.resolve();
        }, function() {
            // do nothing
        });
    }

    function onClose() {
        return $.when(self.defer.promise());
    }
};

module.exports = productCategoryFormController;