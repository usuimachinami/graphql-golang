const vm = new Vue({
    el: '#app',
    data: {
        results: []
    },
    mounted() {
        // var instance = axios.create({
        //     baseURL: 'http://localhost:5000',
        //     timeout: 1000,
        //     headers: {'content-type': 'text/plain;charset=UTF-8', 'authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTE0MTE0MjcwLCJuYW1lIjoidGVzdHVzZXIifQ.F2uHXNfz8JgdOr_Ey9gkMCApw6mfDtOxHkJZCZJqfro'}
        // });
        var configGraphQL = {
            url: 'http://localhost:5000/query',
            method: 'post',
            headers: { 'Content-Type': 'application/graphql;charset=UTF-8', 'authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTE0MTE0MjcwLCJuYW1lIjoidGVzdHVzZXIifQ.F2uHXNfz8JgdOr_Ey9gkMCApw6mfDtOxHkJZCZJqfro' },
            data: 'query { titles (order: "") { name, created_at, stories { name } } }'
        };
        axios(configGraphQL)
            .then(response => {
                var titles = response.data.data.titles;
                for (var i in titles) {
                    titles[i].created_at = dateFormat(new Date(titles[i].created_at))
                }
                this.results = titles
            });
    }
});

function dateFormat(date) {
    var y = date.getFullYear();
    var m = date.getMonth() + 1;
    var d = date.getDate();
    var w = date.getDay();
    var wNames = ['Sun.', 'Mon.', 'Tue.', 'Wed.', 'Thu.', 'Fri.', 'Sta.'];

    m = ('0' + m).slice(-2);
    d = ('0' + d).slice(-2);

    return y + '.' + m + '.' + d + ' ' + wNames[w];
}