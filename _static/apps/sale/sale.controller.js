/*global $app $*/
'use strict'

function saleController() {
    var $ = $app.$;
    var $notify = $app.$notify;
    var $tablegrid = $app.$tablegrid;
    var $modal = $app.$modal;
    var $http = $app.$http;
    var $form = $app.$form;
    var saleForm = require('./form/sale.form.js')($app);

    var self = {
        tableGrid: {},
        table: '#manage-table ',
        form: saleForm,
        load: onLoad,
        renderTable: renderTable,
        endpoint: 'api/v1/sales'
    };

    self.load();

    return self;

    function renderTable() {
        self.tableGrid = $tablegrid.render("#sale-table", self.endpoint, [{
                    sortable: false,
                    data: null,
                    "render": function(data, type, full) {
                        if (full['image']) {
                            return '<img class="table-image" src="./uploads/sale_photos/' + full['image'] + '" />'
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

        $('#sale-table').on('click', '.edit-data', function() {
            var saleId = $(this).data("id");
            showFormEdit(saleId);
        });
    }

    function onLoad() {
        self.renderTable();

        $('body').on('click', '#sale-add', function() {
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
            var form = self.form.controller(self.endpoint, model.data[0])

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

module.exports = saleController;