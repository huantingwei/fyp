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
    isLoggedIn: localStorage.getItem('authenticated') ? true : true,
}

export default function identify(state = initialState, action) {
    // const { payload } = action

    switch (action.type) {
        case Types.LOGIN: {
            localStorage.setItem('authenticated', true)
            return { ...state, isLoggedIn: true }
        }
        case Types.LOGOUT: {
            localStorage.setItem('authenticated', false)
            return { ...state, isLoggedIn: false }
        }
        default:
            return state
    }
}
