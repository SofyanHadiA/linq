
function customerFormModule ($app) {
	
	global.$app = $app;
	
	return {		
		'controller': require('./product.form.controller.js'),		
		'template': require('./product.form.template.hbs'),
	}
};

module.exports = customerFormModule; 