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
    }).then(res => res.json());
}

sendler.addEventListener("click", function () {
    fn("/new", {
        Title: String(title.value),
        TextArticle: String(article.value)
    })
})
