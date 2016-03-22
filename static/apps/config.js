'use strict'

var config = {
    route: {
        default: 'home',
        home: {
            template: 'dashboardHome',
            controller: 'dashboardController'
        },
        customers: {
            templateUrl: 'people/customer.html',
            controller: 'customerController',
            model: ''
        },
        items: {
            templateUrl: 'item/item.html',
            controller: 'itemController'
        }
    }
}

module.exports = config;
