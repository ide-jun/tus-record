import React from 'react'

const Header = () => {
    return (
        <div className="App">    
            <header>
                <h1>tus-record</h1>
            </header>
        </div>
    );
}

const Footer = () => {
    return (
        <div className="App">
            <footer>
                <p>footerだよ</p>
            </footer>
        </div>
    );
}

const AuthErrorForm = () => {
    return (
        <div className='App'>
            <p>ログインしてください</p>
        </div>
    )
}

export { Header, Footer, AuthErrorForm };