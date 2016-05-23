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
        modal: $app.$view.$modal,
        formId: "#product-form",
        data: data || {},
        isPhotoChanged: false,
        defer: $.Deferred(),
        formConfig: {
            rules: {
                title: {
                    minlength: 5,
                    required: true
                },
                sellPrice: {
                    required: true
                },
                categoryId: {
                    required: true
                }
            }
        }
    }

    self.load();

    return self;

    function onLoad() {
        self.modalConfig = self.modalConfig || {
            size: 'lg',
            modalId: "product-modal"
        }

        var input = {
            uidInput: $form.input("uid").setValue(self.data["uid"] || "AUTO"),
            skuInput: $form.input("sku").setValue(self.data["sku"]),
            titleInput: $form.input("title").setValue(self.data["title"]).setClass("required"),
            buyPriceInput: $form.input("buyPrice").setValue(self.data["buyPrice"], 0),
            sellPriceInput: $form.input("sellPrice").setValue(self.data["sellPrice"], 0).setClass("required"),
            stockInput: $form.input("stock").setValue(self.data["stock"], 0).setClass("required"),
            categoryInput: $form.input("categoryId").setValue(self.data["categoryId"]).setClass("required")
        };

        self.modal = $modal.show(require('./product.form.template.hbs'), input, self.modalConfig);

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

        $('#product-photo').cropit({
            onFileChange: function() {
                self.isPhotoChanged = true;
            },
            exportZoom: 3
        });

        if (self.data.image) {
            $('#product-photo').cropit('imageSrc', './uploads/product_photos/' + self.data.image);
        };

        $('#select-image-btn').click(function() {
            $("#product-form.cropit-image-input").prop('disabled', false);
            $('.cropit-image-input').click();
        });

        renderCategoryDropDown();

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

    function onDone(data) {
        $.when(uploadPhoto(data.uid)).then(function() {
            self.modal.hide();
            $.when(self.defer.promise());
        }, function() {
            // do nothing
        });
    }

    function renderCategoryDropDown() {
        $('body #categoryId').selectize({
            valueField: 'uid',
            labelField: 'title',
            searchField: 'title',
            create: false,
            load: function(query, callback) {
                if (!query.length) return callback();
                $.ajax({
                    url: 'https://linq-sofyan-a-1.c9users.io/api/v1/productcategories',
                    type: 'GET',
                    error: function() {
                        callback();
                    },
                    success: function(res) {
                        callback(res.data.slice(0, 10));
                    }
                });
            }
        });
    }
};

module.exports = productFormController;