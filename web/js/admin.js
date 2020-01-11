var title = document.getElementById("Title");
var article = document.getElementById("Article");
var sendler = document.getElementById("sender");
function fn(url, obj) {
    return fetch(url, {
        method: 'POST', // *GET, POST, PUT, DELETE, etc.
        mode: 'cors', // no-cors, cors, *same-origin
        cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
        credentials: 'same-origin', // include, *same-origin, omit
        headers: {
            'Content-Type': 'application/json',
            // 'Content-Type': 'application/x-www-form-urlencoded',
        },
        redirect: 'follow', // manual, *follow, error
        referrer: 'no-referrer', // no-referrer, *client
        body: JSON.stringify(obj), // тип данных в body должен соответвовать значению заголовка "Content-Type"
    })
}

sendler.addEventListener("click", function () {
    if(String(title.value).length > 2 && String(article.value).length > 2) {
        fn("/new", {
            Title: String(title.value),
            TextArticle: String(article.value)
        })
    }
})

function setCookie(name, value, options = {}) {

  let updatedCookie = encodeURIComponent(name) + "=" + encodeURIComponent(value);

  for (let optionKey in options) {
    updatedCookie += "; " + optionKey;
    let optionValue = options[optionKey];
    if (optionValue !== true) {
      updatedCookie += "=" + optionValue;
    }
  }

  document.cookie = updatedCookie;
}

// var LogOut = document.getElementById("logOut")
// LogOut.addEventListener('click', function(e) {
//     setCookie('auth-session', '')
//     document.location.href = '/admin'
// });
