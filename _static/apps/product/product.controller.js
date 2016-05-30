/*global $app $*/
'use strict'

function productController() {
    var $ = $app.$;
    var $notify = $app.$notify;
    var $tablegrid = $app.$tablegrid;
    var $modal = $app.$modal;
    var $http = $app.$http;
    var $form = $app.$form;
    var productForm = require('./form/product.form.js')($app);

    var self = {
        tableGrid: {},
        table: '#manage-table ',
        form: productForm,
        load: onLoad,
        renderTable: renderTable,
        endpoint: 'api/v1/products'
    };

    self.load();

    return self;

    function renderTable() {
        self.tableGrid = $tablegrid.render("#product-table", self.endpoint, [{
                    sortable: false,
                    data: null,
                    "render": function(data, type, full) {
                        if (full['image']) {
                            return '<img class="table-image" src="./uploads/product_photos/' + full['image'] + '" />'
                        }
                        return ""
                    }
                }, {
                    data: 'sku'
                }, {
                    data: 'title'
                }, {
                    data: 'category.title'
                }, {
                    data: 'sellPrice'
                }, {
                    data: 'stock'
                }

            ],
            'uid');

        self.tableGrid.action.delete = doDelete;
        self.tableGrid.action.deleteBulk = doDeleteBulk;

        $('#product-table').on('click', '.edit-data', function() {
            var productId = $(this).data("id");
            showFormEdit(productId);
        });
    }

    function onLoad() {
        self.renderTable();

        $('body').on('click', '#product-add', function() {
            showFormCreate();
        });
    }

    function showFormCreate() {
        var form = self.form.controller(self.endpoint, null)

        $.when(form.defer.promise()).done(function() {
            self.tableGrid.reload();
        });
    }

    function showFormEdit(id) {
        $http.get(self.endpoint + "/" + id).done(function(model) {
            var form = self.form.controller(self.endpoint, model.data)

            $.when(form.defer.promise()).done(function() {
                self.tableGrid.reload();
            })
        });
    }

    function doDelete(id) {
        $http.delete(self.endpoint + "/" + id).done(function(model) {
            self.tableGrid.reload();
        });
    }

    function doDeleteBulk(ids) {
        $http.post(self.endpoint + "/bulkdelete", {
            ids: ids
        }).done(function(ids) {
            self.tableGrid.reload();
        });
    }
};

module.exports = productController;