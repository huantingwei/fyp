function notNullOrUndefined(value) {
    return !(value === undefined) && !(value === null)
}
const transform = (data, primaryKey = 'name', secondaryKey = '') => {
    console.log(primaryKey, secondaryKey)
    let res = []
    if (Object.keys(data).length > 0) {
        for (let key of Object.keys(data)) {
            // null
            if (data[key] === null) {
                res.push({
                    label: key,
                    type: 'text',
                    primaryKey: primaryKey,
                    content: 'null',
                })
                // array
            } else if (Array.isArray(data[key])) {
                res.push({
                    label: key,
                    type: 'arrayObj',
                    primaryKey: primaryKey,
                    secondaryKey: secondaryKey,
                    content: notNullOrUndefined(data[key]) ? data[key] : 'null',
                })
                // object
            } else if (typeof data[key] === 'object') {
                res.push({
                    label: key,
                    type: 'chip',
                    content: notNullOrUndefined(data[key]) ? data[key] : 'null',
                })
                // long text
            } else if (
                typeof data[key] === 'string' &&
                data[key].length > 100
            ) {
                res.push({
                    label: key,
                    type: 'multiline',
                    content: notNullOrUndefined(data[key])
                        ? data[key].toString()
                        : 'null',
                })
                // short text
            } else {
                res.push({
                    label: key,
                    type: 'text',
                    content: notNullOrUndefined(data[key])
                        ? data[key].toString()
                        : 'null',
                })
            }
        }
    }
    return res
}

const flattenWorkload = (data) => {
    const needFlatten = ['objectmeta']
    const flatten = (item) => {
        let resObj = {}
        for (let key of Object.keys(item)) {
            // [key] needs flatten
            if (needFlatten.includes(key)) {
                // object
                for (let k of Object.keys(item[key])) {
                    resObj[k] = item[key][k]
                }
            } else {
                resObj[key] = item[key]
            }
        }
        return resObj
    }

    // array like - multiple item
    if (Array.isArray(data)) {
        let res = []
        // iterate array of object
        for (let item of data) {
            res.push(flatten(item))
        }
        return res
    }
    // object like - single item
    else if (typeof data === 'object') {
        return flatten(data)
    }
}

export { transform, flattenWorkload, notNullOrUndefined }
