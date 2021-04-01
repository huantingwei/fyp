import React, { useState } from 'react'
import store from 'redux/store'
import { Provider } from 'react-redux'
import { BrowserRouter } from 'react-router-dom'
import Root from 'containers'
import Login from 'containers/login'

function App() {
    const [token, setToken] = useState(localStorage.getItem('token'))
    window.addEventListener('storage', (e) => {
        setToken(localStorage.getItem('token'))
    })

    return (
        <Provider store={store}>
            <BrowserRouter>
                {token === '' || token === null || token === undefined ? <Login /> : <Root />}
            </BrowserRouter>
        </Provider>
    )
}

export default App
