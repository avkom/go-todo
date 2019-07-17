define([
	'underscore',
	'backbone'
], function (_, Backbone) {

	var TaskModel = Backbone.Model.extend({
		defaults: {
            id: '',
            title: '',
			description: ''
		}
	});

	return TaskModel;
});