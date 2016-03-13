
function customerModule ($app) {
	
	global.$app = $app;
	
	return {
		'model': require('./customer.model.js'),
		'controller': require('./customer.controller.js'),		
		'template': require('./customer.template.hbs'),
	}
};

module.exports = customerModule; 