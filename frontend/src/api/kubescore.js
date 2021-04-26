class kubescoreAPI {
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
    static _upload(data) {
        return {
            method: 'POST',
            headers: {
                'Content-Type': 'multipart/form-data',
            },
            url: '/api/v1/kubescore/interactive/upload',
            data: data,
        }
    }
    static _getInteractive(data) {
        return {
            method: 'GET',
            url: '/api/v1/kubescore/interactive/get',
            data: data,
        }
    }
}

export default kubescoreAPI
