console.log("JS file loaded")

var myRequest = new Request('http://localhost:9090/tweets'); // need to make an env var

fetch(myRequest, {
    method: 'GET', // *GET, POST, PUT, DELETE, etc.
    mode: 'cors', // no-cors, *cors, same-origin
    cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
    credentials: 'same-origin', // include, *same-origin, omit
    headers: {
        Accept: 'application/json',
    },
}).then(function (response) {
    console.log(response.status);
    return response.json();
}).then(function (data) {
    displayTweets(data);
});

function displayTweets(tweets) {
    if (!tweets) return;

    var mainContainer = document.getElementById("tweetData");

    for (var i = 0; i < tweets.length; i++) {
        if (!tweets[i].text) {
            continue
        }

        // quick date format
        var d = new Date(tweets[i].created);
        var dFormatted = d.toLocaleDateString("en-US", {
            weekday: 'short', // long, short, narrow
            day: 'numeric', // numeric, 2-digit
            year: 'numeric', // numeric, 2-digit
            month: 'long', // numeric, 2-digit, long, short, narrow
            hour: 'numeric', // numeric, 2-digit
            minute: 'numeric', // numeric, 2-digit
            second: 'numeric', // numeric, 2-digit
        });
    
        // setup elements
        var div = document.createElement("div");
        var h2 = document.createElement("h2");
        var p = document.createElement("p");

        // append data to elements
        h2.innerHTML = tweets[i].text;
        p.innerHTML = 'Created on: ' + dFormatted;
        div.appendChild(h2);
        h2.insertAdjacentElement('afterend', p);
        mainContainer.appendChild(div);
    }
}

function submitTweet() {
    // on form submit, get the data and send a post request

    var formText = document.getElementById('tweet-text').value;
    let text = '{"text": "' + formText + '"}';

    postTweet(myRequest, text);
}

function postTweet(request, data) {
    fetch(request, {
        method: 'POST',
        mode: 'cors', // no-cors, *cors, same-origin
        cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
        headers: {
            Accept: 'application/json',
        },
        body: data
    }).then(function (response) {
        console.log(response.status);
    });
}