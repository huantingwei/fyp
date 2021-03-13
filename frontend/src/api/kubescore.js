class kubescoreAPI {
    static get(id) {
        return {
            method: 'GET',
            url: '/api/v1/kubescore/get',
            params: { id: '5ffdb75dfd8fa33d155bdf4e' },
        }
    }
    static _new() {
        return {
            method: 'GET',
            url: '/api/v1/kubescore/new',
        }
    }
    static _list() {
        return {
            method: 'GET',
            url: '/api/v1/kubescore/list',
        }
    }
    static _delete(data) {
        return {
            method: 'POST',
            url: '/api/v1/kubescore/delete',
            data: data,
        }
    }
}

export default kubescoreAPI
