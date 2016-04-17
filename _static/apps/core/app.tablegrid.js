var bootbox = require('bootbox');

// load from bower since npm datatables package version does not include dataTables.bootstrap.js
require('./../../vendors/datatables/media/js/jquery.dataTables.js');
require('./../../vendors/datatables/media/js/dataTables.bootstrap.js');

function tableGridModule() {
    
    var tablegrid = {
        table: "",
        dataTable: {},
        render: render,
        get_selected_rows: get_selected_rows,
        delete: {},
        reload: reload
    }

    return tablegrid;

    function render(tableContainer, serviceUrl, tableConfig, columnId = "id") {
        tablegrid.table = tableContainer;

        tableConfig = [{
            sortable: false,
            data: columnId,
            render: function (data, type, row) {
                return '<input type="checkbox" id="rows-' + data + '" value="' + data + '"/>';
            }
        }]
        .concat(tableConfig)
        .concat([{
            sortable: false,
            data: columnId,
            render: function (data, type, row) {
                return '<div class="btn-group"><a class="btn btn-xs btn-default edit-data" data-id="'+data+'" >'
                    + '<i class="fa fa-edit"></i></a> '
                    + '<a class="btn btn-xs btn-default btn-delete" data-id="' + data + '">'
                    +'<i class="fa fa-trash"></i></a></div>';
            }
        }]);

        tablegrid.dataTable = $(tablegrid.table).DataTable({
            "info": true,
            "autoWidth": false,
            columns: tableConfig,
            "pageLength": 25,
            "order": [[1, "asc"]],
            "processing": true,
            "serverSide": true,
            "ajax": {
                url: serviceUrl,
                type: "get",
                error: function (error) {
                    $.notify({icon: 'fa fa-info-circle', message: error.message}, { type: "info" });
                }
            }
        });

        $(tableContainer + ' #select-all').click(function () {
            if ($(this).prop('checked')) {
                $(tableContainer + " tbody :checkbox").each(function () {
                    $(this).prop('checked', true);
                });
            }
            else {
                $(tableContainer + " tbody :checkbox").each(function () {
                    $(this).prop('checked', false);
                });
            }
        });

        $(tableContainer + " tbody").on("click", '.btn-delete', function (event) {
            event.preventDefault();
            var id = $(this).data("id");
            bootbox.confirm('Are you sure to delete this data?', function (result) {
                if (result) {
                    tablegrid.delete(id)
                }
            });
        });

        $('#delete-selected').click(function (event) {
            event.preventDefault();
            if ($(tableContainer + " tbody :checkbox:checked").length > 0) {
                var url = $("#delete-selected").attr('href');
                bootbox.confirm('Are you sure to delete selected data(s)?', function (result) {
                    if (result) {
                        do_delete(url);
                        $(tableContainer + '#select-all').prop('checked', false);
                    }
                });
            }
            else {
                $.notify({
                    icon: 'fa fa-warning',
                    message: 'No data selected',
                }, {
                        type: "danger"
                    });
            }
        });

        $('#import-excel').click(function () {
            event.preventDefault();
            var url = $(this).attr('href');
            $modal(url, 'md');
        });

        return tablegrid;
    }
    
    function reload() {
        tablegrid.dataTable.ajax.reload();
    }

    function get_selected_rows() {
        var selected_rows = new Array();
        $(tablegrid.table + "tbody :checkbox:checked")
            .each(function () {
                selected_rows.push($(this).val());
            });
        return selected_rows;
    }

    function do_delete(url) {
        var row_ids = get_selected_rows();
        $http.post(url, { 'ids[]': row_ids }, tablegrid.dataTable.ajax.reload)
    }
};

module.exports = tableGridModule();