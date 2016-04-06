
function customerFormModule ($app) {
	
	global.$app = $app;
	
	return {		
		'controller': require('./user.form.controller.js'),		
		'template': require('./user.form.template.hbs'),
	}
};

module.exports = customerFormModule; 