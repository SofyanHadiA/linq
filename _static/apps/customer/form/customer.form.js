
function customerFormModule ($app) {
	
	global.$app = $app;
	
	return {		
		'controller': require('./customer.form.controller.js'),		
		'template': require('./customer.form.template.hbs'),
	}
};

module.exports = customerFormModule; 