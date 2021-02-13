export const Types = {
    LOGIN: 'auth/LOGIN',
    LOGOUT: 'auth/LOGOUT',
}

export const Actions = {
    login: () => ({
        type: Types.LOGIN,
    }),
    logout: () => ({
        type: Types.LOGOUT,
    }),
}

const initialState = {
    isLoggedIn: true,
}

export default function identify(state = initialState, action) {
    // const { payload } = action

    switch (action.type) {
        case Types.LOGIN: {
            return { ...state, isLoggedIn: true }
        }
        case Types.LOGOUT: {
            return { ...state, isLoggedIn: false }
        }
        default:
            return state
    }
}
