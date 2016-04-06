module.exports = '{{#each data}}'+
        '<div class="col-md-6 col-sm-6 col-xs-12" >' +
            '<div id="module-icon-{{this.module_id}}>"  class="info-box">' +
                '<a href="#{{this.module_id}}" >' +
                    '<span class="info-box-icon bg-blue">' +
                        '<i class="fa {{this.icon}}"></i>' +
                    '</span>' +
                '</a>' +
                '<div class="info-box-content">' +
                    '<a href="#{{this.module_id}}" >' +
                        '<h3>{{this.title}}</h3>' +
                    '</a>'+
                    '{{this.description}}' +
                '</div>' +
            '</div>' +
        '</div>'+
        '{{/each}}';