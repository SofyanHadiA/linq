/*global $app $*/

require('cropit');

function saleFormController(endpoint, data) {
    var $modal = $app.$view.$modal;
    var $form = $app.$view.$form;
    var $http = $app.$http;

    var self = {
        load: onLoad,
        modal: $app.$view.$modal,
        formId: "#sale-form",
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
                stock: {
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
            modalId: "sale-modal"
        }

        var input = {
            uidInput: $form.input("uid").setValue(self.data["uid"] || "AUTO"),
            skuInput: $form.input("sku").setValue(self.data["sku"]),
            titleInput: $form.input("title").setValue(self.data["title"]).setClass("required"),
            buyPriceInput: $form.input("buyPrice").setValue(self.data["buyPrice"], 0),
            sellPriceInput: $form.input("sellPrice").setValue(self.data["sellPrice"], 0).setClass("required"),
            stockInput: $form.input("stock").setValue(self.data["stock"], 0).setClass("required"),
            categoryInput: $form.input("categoryId").setClass("required")
        };
        
        if(self.data.category){
            input.categoryInput.setSelected(self.data["categoryId"], self.data["category"]["title"]);
        }

        self.modal = $modal.show(require('./sale.form.template.hbs'), input, self.modalConfig);

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

        $('#sale-photo').cropit({
            onFileChange: function() {
                self.isPhotoChanged = true;
            },
            exportZoom: 3
        });

        if (self.data.image) {
            $('#sale-photo').cropit('imageSrc', './uploads/sale_photos/' + self.data.image);
        };

        $('#select-image-btn').click(function() {
            $("#sale-form.cropit-image-input").prop('disabled', false);
            $('.cropit-image-input').click();
        });

        renderCategoryDropDown();

        return self;
    }

    function uploadPhoto(saleId) {
        if (self.isPhotoChanged) {
            var imageData = $('#sale-photo').cropit('export');
            return $http.put(endpoint + "/" + saleId + "/photo", imageData);
        }
        else {
            return null
        }
    }

    function onDone(data) {
        $.when(uploadPhoto(data.uid)).then(function() {
            self.modal.hide();
            self.defer.resolve();
        }, function() {
            // do nothing
        });
    }

    function renderCategoryDropDown() {
        $('body #categoryId').selectize({
            persist: true,
            valueField: 'uid',
            labelField: 'title',
            searchField: 'title',
            create: false,
            load: function(query, callback) {
                if (!query.length) return callback();
                $.ajax({
                    url: './api/v1/salecategories',
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

module.exports = saleFormController;