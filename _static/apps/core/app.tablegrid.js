var bootbox = require('bootbox');
require('./../../vendors/datatables/media/js/jquery.dataTables.js');
require('./../../vendors/datatables/media/js/dataTables.bootstrap.js');


function tableGridModule() {

    var self = {
        table: "",
        dataTable: {},
        render: render,
        getSelectedRows: getSelectedRows,
        remove: remove,
        reload: reload,
        action: {
            delete: undefined,
            deleteBulk: undefined
        }
    }

    return self;

    function render(tableContainer, serviceUrl, tableConfig, columnId = "id") {
        self.table = tableContainer;

        tableConfig = [{
                sortable: false,
                data: columnId,
                render: function(data, type, row) {
                    return '<input type="checkbox" id="rows-' + data + '" value="' + data + '"/>';
                }
            }]
            .concat(tableConfig)
            .concat([{
                sortable: false,
                data: columnId,
                render: function(data, type, row) {
                    return '<div class="btn-group"><a class="btn btn-xs btn-default edit-data" data-id="' + data + '" >' + '<i class="fa fa-edit"></i></a> ' + '<a class="btn btn-xs btn-default btn-delete" data-id="' + data + '">' + '<i class="fa fa-trash"></i></a></div>';
                }
            }]);

        self.dataTable = $(self.table).DataTable({
            "info": true,
            "autoWidth": false,
            columns: tableConfig,
            "pageLength": 25,
            "order": [
                [1, "asc"]
            ],
            "processing": true,
            "serverSide": true,
            "ajax": {
                url: serviceUrl,
                type: "get",
                error: function(error) {
                    $app.$notify.danger(error.message);
                }
            }
        });

        $(tableContainer + ' #select-all').click(function() {
            if ($(this).prop('checked')) {
                $(tableContainer + " tbody :checkbox").each(function() {
                    $(this).prop('checked', true);
                });
            }
            else {
                $(tableContainer + " tbody :checkbox").each(function() {
                    $(this).prop('checked', false);
                });
            }
        });

        $(tableContainer + " tbody").on("click", '.btn-delete', function(event) {
            event.preventDefault();
            var id = $(this).data("id");
            self.remove(id)
        });

        $('#delete-selected').click(function(event) {
            event.preventDefault();
            if ($(tableContainer + " tbody :checkbox:checked").length > 0) {
                var url = $("#delete-selected").attr('href');
                bootbox.confirm('Are you sure to delete selected data(s)?', function(result) {
                    if (result) {
                        var ids = self.getSelectedRows();
                        self.action.deleteBulk(ids);
                        $(tableContainer + '#select-all').prop('checked', false);
                    }
                });
            }
            else {
                $app.$notify.warning('No data selected')
            }
        });

        $('#import-excel').click(function() {
            event.preventDefault();
            var url = $(this).attr('href');
            $modal(url, 'md');
        });

        return self;
    }

    function getSelectedRows() {
        var selectedRows = [];

        $(self.table + " tbody :checkbox:checked").each(function() {
            selectedRows.push($(this).val());
        });

        return selectedRows;
    }
    
    function remove(id){
        bootbox.confirm('Are you sure to delete this?', function(result) {
            if (result) {
                self.action.delete(id)
            }
        });
    }

    function reload() {
        self.dataTable.ajax.reload();
    }
};

module.exports = tableGridModule;