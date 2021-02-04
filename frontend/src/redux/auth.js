export const Types = {
    LOGIN: 'auth/LOGIN',
}

export const Actions = {
    login: (data) => ({
        type: Types.LOGIN,
        payload: data,
    }),
}

const initialState = {
    isLoggedIn: false,
}

export default function identify(state = initialState, action) {
    const { payload } = action

    switch (action.type) {
        case Types.LOGIN: {
            return { isLoggedIn: payload, ...state }
        }
        default:
            return state
    }
}
