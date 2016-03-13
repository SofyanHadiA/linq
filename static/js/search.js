"use strict";

$(document).ready(function () {
    $("#search-form").submit(function (event) {
        event.preventDefault();

        var form = $(this);
        var data = form.find('#search-value').val();

        $.post(form.attr('action'), {search: data}, function (result) {
            // $(".result").html(data);
            updateTableData(result);
        });
    });

    function updateTableData(data) {

        console.log(data);

        if(data.length>0) {
            var result = "";
            data.forEach(function (value) {
                result += '<tr>';
                result += "<td width='5%'><input type='checkbox' id='person_" + value['person_id'] + " value='" + value['person_id'] + "'/></td>";
                result += '<td width="20%">' + value['last_name'] + '</td>';
                result += '<td width="20%">' + value['first_name'] + '</td>';
                result += '<td width="30%"><a href="' + value['email'] + '" target="_top">Send Mail</a></td>';
                result += '<td width="20%">' + value['phone_number'] + '</td>';
                result += '<td width="5%"><a href="customers/view/' + value['person_id'] + '" data-toggle="modal" data-target="#modal-container"><i class="fa fa-edit"></i></a>'
                + ' | <a href="customers/delete/' + value['person_id'] + '"><i class="fa fa-trash"></i></a></td>';
                result += '</tr>';
            });

            $("#manage-table tbody").html(result);
        }
        else {
            $("#manage-table tbody").html("<tr></tr><td colspan='6'>No Data Found</td></tr>");
        }
    }
});
