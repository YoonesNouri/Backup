//9.2

(function () {
    'use strict';
    function getLocation() {
        if (navigator.geolocation) {
            // Call the geolocation API to get the user's position
            navigator.geolocation.getCurrentPosition(showPosition, showError);
        } else {
            alert("Geolocation is not supported by your browser.");
        }
    }

    function showPosition(position) {
        // Get latitude and longitude from the position object
        const latitude = position.coords.latitude;
        const longitude = position.coords.longitude;

        // Display the latitude and longitude on the page
        const locationOutput = document.getElementById('locationOutput');
        locationOutput.textContent = `Latitude: ${latitude}, Longitude: ${longitude}`;
    }

    function showError(error) {
        switch (error.code) {
            case error.PERMISSION_DENIED:
                alert("User denied the request for Geolocation.");
                break;
            case error.POSITION_UNAVAILABLE:
                alert("Location information is unavailable.");
                break;
            case error.TIMEOUT:
                alert("The request to get user location timed out.");
                break;
            case error.UNKNOWN_ERROR:
                alert("An unknown error occurred.");
                break;
        }
    }

    const getLocationButton = document.getElementById('getLocationButton');
    getLocationButton.addEventListener('click', getLocation);

}());