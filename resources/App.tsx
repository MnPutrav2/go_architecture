import LogoReact from './assets/logo/react.png'
import LogoGolang from './assets/logo/golang.png'
import { JSX, useRef, useState } from 'react';
import { Login, Register, Token } from './types/auth';
import { Request } from './lib/api';
import Modal, { ModalRef } from './components/Modal';
import { Response } from './types/response';

function LoginCard({onClose}: {onClose: () => void}): JSX.Element {
    const [loginPayload, setLoginPayload] = useState<Login>({name: "", password: ""})

    async function login() {
        try {
            const json = await new Request('http://localhost:8080/api/login').POST<Login, Response<Token>>(loginPayload)
            sessionStorage.setItem("token", json.result.token)
            alert("success")
            onClose()
        }catch(e){
            alert(e)
        }
    }

    return (
        <section>
            <input type="text" value={loginPayload.name} onChange={(e) => setLoginPayload({...loginPayload, name: e.target.value})} />
            <input type="text" value={loginPayload.password} onChange={(e) => setLoginPayload({...loginPayload, password: e.target.value})}  />
            <button className='register' onClick={() => login()}>Login</button>
        </section>
    )
}

function RegisterCard({onClose}: {onClose: () => void}): JSX.Element {
    const [registerPayload, setRegisterPayload] = useState<Register>({name: "", password: "", email: ""})

    async function register() {
        try {
            await new Request('http://localhost:8080/api/register').POST<Register>(registerPayload)
            alert("success")
            onClose()
        }catch(e){
            alert(e)
        }
    }

    return (
        <section>
            <input type="text" value={registerPayload.name} onChange={(e) => setRegisterPayload({...registerPayload, name: e.target.value})} />
            <input type="text" value={registerPayload.password} onChange={(e) => setRegisterPayload({...registerPayload, password: e.target.value})}  />
            <input type="email" value={registerPayload.email} onChange={(e) => setRegisterPayload({...registerPayload, email: e.target.value})}  />
            <button className='register' onClick={() => register()}>Register</button>
        </section>
    )
}

export default function App() {

    const registerRef = useRef<ModalRef>(null)
    const loginRef = useRef<ModalRef>(null)
    const [openRegister, setOpenRegister] = useState<boolean>(false)
    const [openLogin, setOpenLogin] = useState<boolean>(false)

    return (
        <main>
            {/* Delete this */}
            <section className="container">
                <div className="header">
                    <h1>DIGO</h1>
                    <p>
                        React + Golang structure project like laravel
                    </p>
                    <div className="logo">
                        <img src={LogoReact} alt="" />
                        <img src={LogoGolang} alt="" />
                    </div>
                    <div style={{display: 'flex', justifyContent: "center", gap: 10}}>
                        <button className='login' onClick={() => setOpenLogin(true)}>Test login</button>
                        <button className='register' onClick={() => setOpenRegister(true)}>Test register</button>
                    </div>
                </div>
            </section>

            <Modal ref={registerRef} open={openRegister} onClose={() => setOpenRegister(false)}>
                <RegisterCard onClose={() => registerRef.current?.close()} />
            </Modal>

            <Modal ref={loginRef} open={openLogin} onClose={() => setOpenLogin(false)}>
                <LoginCard onClose={() => loginRef.current?.close()} />
            </Modal>

            {/* Delete this */}
        </main>
    );
}