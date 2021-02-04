class authAPI {
    static login(data) {
        return {
            method: 'POST',
            url: '/',
            data: data,
        }
    }
}

export default authAPI
