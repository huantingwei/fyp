class kubebenchAPI {
    static _new() {
        return {
            method: 'GET',
            url: '/api/v1/kubebench/new',
        }
    }
    static _list() {
        return {
            method: 'GET',
            url: '/api/v1/kubebench/list',
        }
    }
    static _delete(data) {
        return {
            method: 'POST',
            url: '/api/v1/kubebench/delete',
            data: data,
        }
    }
}

export default kubebenchAPI
