
function customerFormModule ($app) {
	
	global.$app = $app;
	
	return {		
		'controller': require('./sale.form.controller.js'),		
		'template': require('./sale.form.template.hbs'),
	}
};

module.exports = customerFormModule; 