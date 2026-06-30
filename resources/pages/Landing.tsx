import LogoReact from '../assets/logo/react.png'
import LogoGolang from '../assets/logo/golang.png'
import { JSX, useRef, useState } from 'react';
import { Login, Register, Token } from '../types/auth';
import { Request } from '../lib/api';
import Modal, { ModalRef } from '../components/Modal';
import { Response } from '../types/response';
import { Input } from '../components/Input';
import { useNavigate } from 'react-router';

function LoginCard({onClose}: {onClose: () => void}): JSX.Element {
    const [loginPayload, setLoginPayload] = useState<Login>({name: "", password: ""})
    const navigate = useNavigate()

    async function login() {
        try {
            const json = await new Request('http://localhost:8080/api/login').POST<Login, Response<Token>>(loginPayload)
            sessionStorage.setItem("token", json.result.token)
            sessionStorage.setItem("refresh_token", json.result.refresh_token)
            alert("success")
            onClose()
            navigate("/dashboard")
        }catch(e){
            alert(e)
        }
    }

    return (
        <section style={{width: "15rem"}}>
            <h2 style={{margin: "0.5rem", marginBottom: "1rem", textAlign: "center"}}>Login</h2>
            <Input tipe='text' placeholder='username' value={loginPayload.name} onChange={(e) => setLoginPayload({...loginPayload, name: e.target.value})} />
            <Input tipe='password' placeholder='password' value={loginPayload.password} onChange={(e) => setLoginPayload({...loginPayload, password: e.target.value})} />
            <div style={{margin: "0.5rem", marginTop: "1rem"}}><button className='register' style={{width: "100%", boxSizing: "border-box"}} onClick={() => login()}>Login</button></div>
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
        <section style={{width: "15rem"}}>
            <h2 style={{margin: "0.5rem", marginBottom: "1rem", textAlign: "center"}}>Register</h2>
            <Input tipe='text' placeholder='username' value={registerPayload.name} onChange={(e) => setRegisterPayload({...registerPayload, name: e.target.value})} />
            <Input tipe='password' placeholder='password' value={registerPayload.password} onChange={(e) => setRegisterPayload({...registerPayload, password: e.target.value})} />
            <Input tipe='email' placeholder='email' value={registerPayload.email} onChange={(e) => setRegisterPayload({...registerPayload, email: e.target.value})} />
            <div style={{margin: "0.5rem", marginTop: "1rem"}}><button className='register' style={{width: "100%"}} onClick={() => register()}>Register</button></div>
        </section>
    )
}

export default function Landing() {

    const registerRef = useRef<ModalRef>(null)
    const loginRef = useRef<ModalRef>(null)
    const [openRegister, setOpenRegister] = useState<boolean>(false)
    const [openLogin, setOpenLogin] = useState<boolean>(false)

    return (
        <main>
            {/* Delete this */}

            <section className="container">
                <div className="header">
                    <h2>Welcome</h2>
                    <p style={{marginTop: "0.5rem"}}>in</p>
                    <h1>DIGO</h1>
                    <p>
                        React + Golang structure project like laravel
                    </p>
                    <div className="logo">
                        <img src={LogoReact} alt="" />
                        <img src={LogoGolang} alt="" />
                    </div>
                    <div style={{display: 'flex', justifyContent: "center", gap: 10}}>
                        <button className='login' style={{width: "50%"}} onClick={() => setOpenLogin(true)}>Test login</button>
                        <button className='register' style={{width: "50%"}} onClick={() => setOpenRegister(true)}>Test register</button>
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