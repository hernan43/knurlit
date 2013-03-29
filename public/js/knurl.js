$(document).ready(function(){

	$("#add-params").click(function(){
		add_fields('params');
	});

	$("#add-headers").click(function(){
		add_fields('headers');
	});

	function add_fields(name){
		number_of_items = $('ul#'+name+' li').length;
		new_li = "<li><input placeholder=\"name\" type=\"text\" name=\""+name+"["
		+ number_of_items
		+ "].Name\"/> <input placeholder=\"value\" type=\"text\" name=\""+name+"["
		+ number_of_items
		+ "].Value\"/>"
		+ " <i class=\"icon-trash\"></i>"
		+ "</li>";
		$(new_li).appendTo("#"+name);
	}

	$(document).on('click', '.icon-trash', function(){
		$(this).parent().remove();
	});

	setup_form_for_method();
	$('#methods').change(function(){
		setup_form_for_method();
	});

	function setup_form_for_method(){
		method = $('#methods').find(':selected').text();
		switch(method){
			case 'POST':
			case 'PUT':
			$('#params-block').show();
			break;
			default:
			$('#params-block').hide();
		}
		
	}

  setup_auth_form();   
	function setup_auth_form(){
		auth = $('#auth-radio input:radio:checked').val();
		switch(auth){
			case 'yes':
			$('#auth-block').show();
			break;
			case 'no':
			$('#auth-block').hide();
			break;
		}
	}
	
	$('#auth-radio input:radio').change(function(){
	  setup_auth_form();   
  });

  setup_post_body();	
	function setup_post_body(){
		$("#post-body").hide();
		$("#add-body").text("+ set post body");

    if($('#post-body').val()){
			$("#add-body").text("- set post body");
      $("#post-body").show();
      $("#params").hide();
      $("#add-params").hide();
    }
	}
	
	$("#add-body").click(function(){
		$("#post-body").toggle();
		
		if($("#post-body").is(':visible')) {
			$("#add-body").text("- set post body");
		} else {
			$("#add-body").text("+ set post body");			
		}
		
		$("#params").toggle();
		$("#add-params").toggle();
	});
	
});
