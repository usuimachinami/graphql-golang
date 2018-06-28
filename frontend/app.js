const vm = new Vue({
    el: '#app',
    data: {
        results: []
    },
    mounted() {
        getToken().then(token => {
            var configGraphQL = {
                url: 'http://127.0.0.1:5000/query',
                method: 'post',
                headers: { 'Content-Type': 'application/graphql;charset=UTF-8', 'authorization': 'Bearer ' + token },
                data: 'query { titles (order: "") { name, created_at, stories { name } } }'
            };
            axios(configGraphQL).then(response => {
                var titles = response.data.data.titles;
                for (var i in titles) {
                    titles[i].created_at = dateFormat(new Date(titles[i].created_at))
                }
                this.results = titles
            });
        });
    }
});

async function getToken() {
    var configGraphQL = {
        url: 'http://127.0.0.1:5000/login',
        method: 'post',
        headers: { 'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8' },
        data: 'username=testuser&password=testpassword'
    };
    const res = await axios(configGraphQL);

    return res.data.token;
}

function dateFormat(date) {
    var y = date.getFullYear();
    var m = date.getMonth() + 1;
    var d = date.getDate();
    var w = date.getDay();
    var wNames = ['Sun.', 'Mon.', 'Tue.', 'Wed.', 'Thu.', 'Fri.', 'Sat.'];

    m = ('0' + m).slice(-2);
    d = ('0' + d).slice(-2);

    return y + '.' + m + '.' + d + ' ' + wNames[w];
}