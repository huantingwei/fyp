class projectAPI {
    static _get() {
        return {
            method: 'GET',
            url: '/project',
        }
    }
    static _new(data) {
        return {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            url: '/project',
            data: data,
        }
    }
}

export default projectAPI
