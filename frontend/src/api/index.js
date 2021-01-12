const axios = require('axios')
const instance = axios.create({
    baseURL: '',
    timeout: 20000,
})

async function req(apiConfig) {
    return instance(apiConfig)
        .then(async (r) => {
            let result
            try {
                result = await r.data
            } catch (err) {
                throw new Error(
                    'Unknown server response: ' + r + '\nof error: ' + err
                )
            }
            if (!result.Success) {
                throw new Error(
                    `Server error: ${r.status} ${r.statusText} - ${
                        result.Error || ''
                    }`
                )
            }
            return result.Data
        })
        .then((r) => {
            return r
        })
        .catch((err) => {
            throw err
        })
}

export { instance, req }
