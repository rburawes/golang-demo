/**
 * Get all the authors and populate author drop down menu.
 */
function loadAuthors(authorId) {
    $(document).ready(function () {
        $.ajax({
            url: '/authors',
            type: 'get',
            dataType: 'json',
            success: function (response) {
                var len = response.length;
                $("#author").empty();
                var authors = '<option value="0">Select author</option>';
                for (var i = 0; i < len; i++) {
                    var id = response[i]['id'];
                    var fname = response[i]['firstname'];
                    var lname = response[i]['lastname'];
                    authors += "<option value='" + id + "'>" + fname + " " + lname + "</option>";
                }
                $("#author").html(authors);
                if (authorId) {
                    $("#author option[value=" + authorId + "]").attr('selected', 'selected');
                }
            }
        });
    });
}

/**
 * Confirms delete event from a button or an anchor.
 * TODO: Not the best practice to delete using GET method, should be ended in POST/DELETE.
 */
$(document).ready(function () {
    $('#confirm-delete').on('show.bs.modal', function (e) {
        $(this).find('.btn-ok').attr('href', $(e.relatedTarget).data('href'));
        $('.debug-url').html('Delete URL: <strong>' + $(this).find('.btn-ok').attr('href') + '</strong>');
    });
});

/**
 * Submits a form.
 *
 */
$(document).ready(function () {
    $("form").submit(function (event) {
        event.preventDefault();
        var post_url = $(this).attr("action");
        var form_data = $(this).serialize();
        $.post(post_url, form_data, function (response) {
            window.location.href = "/";
        }).fail(function (data, textStatus, xhr) {
            console.log(xhr + ": " + textStatus)
            $("#error-msg").html("<span><strong>ERROR:</strong> " +data.responseText+"</span>");
            $("#error-msg").show();
        }).always(function () {
            console.log("Form submission ended");
        });
    });
});

/**
 * Validates signup form.
 */
$(document).ready(function () {
    $flag = 1;
    $("#firstname").focusout(function () {
        if ($(this).val() == '') {
            $(this).css("border-color", "#FF0000");
            $('#submit').attr('disabled', true);
            $("#error_firstname").text("* First name is required.");
        }
        else {
            $(this).css("border-color", "#2eb82e");
            $('#submit').attr('disabled', false);
            $("#error_lastname").text("");

        }
    });
    $("#lastname").focusout(function () {
        if ($(this).val() == '') {
            $(this).css("border-color", "#FF0000");
            $('#submit').attr('disabled', true);
            $("#error_lastname").text("* Last name is required!");
        }
        else {
            $(this).css("border-color", "#2eb82e");
            $('#submit').attr('disabled', false);
            $("#error_lastname").text("");
        }
    });
    $("#email").focusout(function () {
        if ($(this).val() == '') {
            $(this).css("border-color", "#FF0000");
            $('#submit').attr('disabled', true);
            $("#error_email").text("* Email address is required!");
        }
        else {
            $(this).css("border-color", "#2eb82e");
            $('#submit').attr('disabled', false);
            $("#error_email").text("");
        }
    });
    $("#password").focusout(function () {
        if ($(this).val() == '') {
            $(this).css("border-color", "#FF0000");
            $('#submit').attr('disabled', true);
            $("#error_password").text("* Please provide a valid password!");
        }
        else {
            $(this).css({"border-color": "#2eb82e"});
            $('#submit').attr('disabled', false);
            $("#error_password").text("");

        }
    });
    $("#cpassword").focusout(function () {
        if ($(this).val() == '') {
            $(this).css("border-color", "#FF0000");
            $('#submit').attr('disabled', true);
            $("#error_confirm_password").text("* Confirm your password please!");
        } else {
            $(this).css({"border-color": "#2eb82e"});
            $('#submit').attr('disabled', false);
            $("#error_confirm_password").text("");
        }

    });

    $("#submit").click(function () {
        if ($("#firstname").val() == '') {
            $("#firstname").css("border-color", "#FF0000");
            $('#submit').attr('disabled', true);
            $("#error_firstname").text("* First name is required!");
        }
        if ($("#lastname").val() == '') {
            $("#lastname").css("border-color", "#FF0000");
            $('#submit').attr('disabled', true);
            $("#error_lastname").text("* Last name is required!");
        }
        if ($("#email").val() == '') {
            $("#email").css("border-color", "#FF0000");
            $('#submit').attr('disabled', true);
            $("#error_email").text("* Email address is required!");
        }
        if ($("#password").val() == '') {
            $("#password").css("border-color", "#FF0000");
            $('#submit').attr('disabled', true);
            $("#error_password").text("* Please provide a valid password!");
        }
        if ($("#cpassword").val() == '') {
            $("#cpassword").css("border-color", "#FF0000");
            $('#submit').attr('disabled', true);
            $("#error_confirm_password").text("* Confirm your password please!");
        }
    });
});