define([
	'underscore',
	'backbone'
], function (_, Backbone) {

	var TaskModel = Backbone.Model.extend({
        urlRoot: '/api/tasks',
		defaults: {
            id: '',
            title: '',
			description: ''
		}
	});

	return TaskModel;
});