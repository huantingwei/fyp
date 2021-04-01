class authAPI {
    static authenticate(data) {
        return {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            url: '/api/v1/login/authenticate',
            data: data,
        }
    }
    static verifyCode(data) {
        return {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            url: '/api/v1/login/verifyCode',
            data: data,
        }
    }
}

export default authAPI
