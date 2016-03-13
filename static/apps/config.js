'use strict'

var config = {
    route: {
        default: 'home',
        home: {
            template: 'dashboardHome',
            controller: 'dashboardController'
        },
        customers: {
            templateUrl: '../app/people/customer.html',
            controller: 'customerController',
            model: ''
        },
        items: {
            templateUrl: '../app/item/item.html',
            controller: 'itemController'
        }
    }
}

module.exports = config;
