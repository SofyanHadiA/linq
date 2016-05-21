
function customerFormModule ($app) {
	
	global.$app = $app;
	
	return {		
		'controller': require('./productCategory.form.controller.js'),		
		'template': require('./productCategory.form.template.hbs'),
	}
};

module.exports = customerFormModule; 