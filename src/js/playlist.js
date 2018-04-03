$(document).ready(function() {
    $('.preSliderBtn').on('click', function() {
        if ($("#playlistSlider").hasClass('active')) {
            $("#playlistSlider").removeClass('active');
        } else {
            $("#playlistSlider").addClass('active');
        }
    });
});