$(document).ready(function() {
    $.ajax({
        type: 'GET',
        url: "/message",
        dataType: 'json'
    })
    .done(function(data) {
    	var rows = '<tr><th>Id</th><th>Subject</th><th>Body</th><th>&nbsp</th>';
        $.each(data, function() {
        	rows += "<tr>";
            rows += "<td>" + this.id + "</td>";
            rows += "<td>" + this.subject + "</td>";
            rows += "<td>" + this.body + "</td>";
            rows += "<td onClick='details(" + this.id + ")' style='color:blue;text-decoration: underline;cursor:pointer;'>More Details</td>";
            rows += "</tr>";
        });
        $('.message-table').append(rows);
    })
    .fail(function() {
      alert("Ajax failed to fetch data")
    });
});

function details(id){
    $.ajax({
    	type: 'GET',
        url: '/message/'+ id,
        dataType: 'json',
    	success: function(data) {
	    	$('.message-table').empty();
	    	var userFrom = $.ajax({
				type: 'GET',
			    url: '/user/'+data.fromuser,
			    dataType: 'json'
			    });
	    	var userTo = $.ajax({
				type: 'GET',
			    url: '/user/'+data.touser,
			    dataType: 'json'
			    });
	    	var labels = $.ajax({
				type: 'GET',
				url: '/message/'+data.id+'/labels',
				dataType: 'json'
				});
	    	var text = '<h2 class="well well-sm">';
	    	$.when(userFrom, userTo, labels).done(function (a,b,c) {
	    		//console.log(a, b, c);
	    		text += data.subject + '</h2>';
		    	text += '<p>' + data.body + '</p>';
		    	text += '<p> <b>From: </b>' + a[0].name + '</p>';
		    	text += '<p> <b>To: </b>' + b[0].name + '</p>';
		    	var label = new Array();
				$.each(c[0], function() {
					label.push(this.name);
				});	
		    	text += '<p> <b>Labels: </b>' + label.join(",") + '</p>'
		    	$('.message-content').append(text);	
		    });
	    }
	});
}

