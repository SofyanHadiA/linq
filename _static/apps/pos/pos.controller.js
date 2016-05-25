/*global $app $*/
'use strict'

function posController() {
    var $ = $app.$;
    var $notify = $app.$notify;
    var $tablegrid = $app.$tablegrid;
    var $modal = $app.$modal;
    var $http = $app.$http;

    var self = {
        tableGrid: {},
        load: onLoad,
        renderTable: renderTable,
        endpoint: 'api/v1/poss'
    };

    self.load();

    return self;

    function renderTable() {
        self.tableGrid = $tablegrid.render("#pos-table", self.endpoint, [{
                    sortable: false,
                    data: null,
                    "render": function(data, type, full) {
                        if (full['image']) {
                            return '<img class="table-image" src="./uploads/pos_photos/' + full['image'] + '" />'
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

        $('#pos-table').on('click', '.edit-data', function() {
            var posId = $(this).data("id");
            showFormEdit(posId);
        });
    }

    function onLoad() {
        self.renderTable();

        $('body').on('click', '#pos-add', function() {
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

module.exports = posController;