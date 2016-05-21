/*global $app $*/

require('cropit');
var bootbox = require('bootbox');


function productFormController(endpoint, data) {
    var $modal = $app.$view.$modal;
    var $form = $app.$view.$form;
    var $http = $app.$http;
    var $notify = $app.$notify;

    var self = {
        load: onLoad,
        close: onClose,
        modal: $app.$view.$modal,
        formId: "#product-form",
        data: data || {},
        isPhotoChanged: false,
        promise: {},
        defer: $.Deferred(),
        formConfig: {
            rules: {
                product_title: {
                    minlength: 5,
                    required: true
                },
                product_sell_price: {
                    required: true
                },
                product_category: {
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

        self.modal = $modal.show(require('./product.form.template.hbs'), input, modalConfig);

        $form.create(self.formId)
            .config(self.formConfig)
            .onSubmit(function() {
                event.preventDefault();
                var newData = $(self.formId).serializeObject();
                newData.buyPrice = parseFloat(newData.buyPrice);
                newData.sellPrice = parseFloat(newData.sellPrice);
                newData.stock = parseFloat(newData.buyPrice);

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
            bootbox.confirm('Are you sure to delete this product?', function(result) {
                if (result) {
                    doDelete(id);
                }
            });
        });

        $('#product-photo').cropit({
            onFileChange: function() {
                self.isPhotoChanged = true;
            }
        });

        if (self.data.photo) {
            $('#product-photo').cropit('imageSrc', './uploads/product_avatars/' + self.data.photo);
        };

        $('#select-image-btn').click(function() {
            $("#product-form.cropit-image-input").prop('disabled', false);
            $('.cropit-image-input').click();
        });

        return self;
    }

    function uploadPhoto(productId) {
        if (self.isPhotoChanged) {
            var imageData = $('#product-photo').cropit('export');
            return $http.put(endpoint + "/" + productId + "/photo", imageData);
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

module.exports = productFormController;