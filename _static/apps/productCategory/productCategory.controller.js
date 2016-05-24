/*global $app $*/
'use strict'

function productCategoryController() {
    var $ = $app.$;
    var $notify = $app.$notify;
    var $tablegrid = $app.$tablegrid;
    var $modal = $app.$modal;
    var $http = $app.$http;
    var $form = $app.$form;
    var productCategoryForm = require('./form/productCategory.form.js')($app);

    var self = {
        tableGrid: {},
        table: '#manage-table ',
        form: productCategoryForm,
        renderTable: renderTable,
        load: onLoad,
        endpoint: 'api/v1/productcategories'
    };

    self.load();

    return self;

    function renderTable() {
        self.tableGrid = $tablegrid.render("#productCategory-table", self.endpoint, [{
                data: 'title'
            }, {
                data: 'slug'
            }, {
                data: 'description'
            }],
            'uid');

        self.tableGrid.action.delete = doDelete;
        self.tableGrid.action.deleteBulk = doDeleteBulk;

        $('#productCategory-table').on('click', '.edit-data', function() {
            var productCategoryId = $(this).data("id");
            showFormEdit(productCategoryId);
        });

    }

    function onLoad() {
        self.renderTable();

        $('body').on('click', '#productCategory-add', function() {
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

module.exports = productCategoryController;