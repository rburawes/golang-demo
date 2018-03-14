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